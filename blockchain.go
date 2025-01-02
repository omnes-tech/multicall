package multicall

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// ReadContract makes a call to a contract and returns the returned bytecode.
func ReadContract(rpc string, from, to *common.Address, encodedCall []byte) ([]byte, error) {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return []byte{}, fmt.Errorf("error establishing RPC connection: %v", err)
	}
	defer client.Close()

	if from == nil {
		from = &ZERO_ADDRESS
	}

	result, err := client.CallContract(context.Background(),
		ethereum.CallMsg{
			From: *from,
			To:   to,
			Data: encodedCall,
		},
		nil,
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// CreateTransaction creates a new transaction object.
func CreateTransaction(
	rpc string,
	from *common.Address,
	to *common.Address,
	msgValue *big.Int,
	callData []byte,
) (*types.Transaction, error) {

	client, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, fmt.Errorf("error establishing RPC connection: %v", err)
	}
	defer client.Close()

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	gasLimit, err := client.EstimateGas(
		context.Background(),
		ethereum.CallMsg{
			From: *from, // the sender of the 'transaction'
			To:   to,    // the destination contract (nil for contract creation)
			Gas:  0,     // if 0, the call executes with near-infinite gas
			// GasPrice:  gasPrice        // wei <-> gas exchange ratio
			GasFeeCap: gasPrice, // EIP-1559 fee cap per gas.
			GasTipCap: gasPrice, // EIP-1559 tip per gas.
			Value:     msgValue,
			Data:      callData,
		},
	)

	if err != nil {
		return nil, err
	}

	nonce, err := client.PendingNonceAt(context.Background(), *from)
	if err != nil {
		return nil, err
	}

	return types.NewTransaction(nonce, *to, msgValue, gasLimit, gasPrice, nil), nil
}

// SendSignedTransaction sends a signed transaction
func SendSignedTransaction(rpc string, tx *types.Transaction) (*types.Receipt, error) {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, fmt.Errorf("error establishing RPC connection: %v", err)
	}
	defer client.Close()

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		return nil, fmt.Errorf("error sending transaction (txHash=%v): %v", tx.Hash(), err)
	}

	// @note implement retry to bump gas
	d := time.Now().Add(MINING_WAIT_DURATION)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		return nil, fmt.Errorf("error while waiting for receipt (txHash=%v): %v", tx.Hash(), err)
	}

	return receipt, nil
}

func parseRevertData(err error) ([]byte, bool) {

	var ec rpc.Error
	var ed rpc.DataError
	if errors.As(err, &ec) && errors.As(err, &ed) && ec.ErrorCode() == 3 {
		revertData := hexutil.MustDecode(ed.ErrorData().(string))

		return revertData, true

	}
	return nil, false
}
