package multicall

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/omnes-tech/abi"
)

func aggregate(
	calls Calls, client *ethclient.Client, signer SignerInterface, to *common.Address,
	funcSignature string, txReturnTypes []string, withValue bool,
) (Result, error) {
	arrayfiedCalls, msgValue, err := calls.ToArray(withValue, false)
	if err != nil {
		return Result{}, err
	}

	receipt, decodedCallResult, err := aggregateTx(
		arrayfiedCalls, msgValue, client, signer, to, funcSignature, txReturnTypes,
	)
	if err != nil {
		return Result{}, err
	}

	return parseResults(decodedCallResult, receipt.Status == 1, receipt)
}

func tryAggregate(
	calls Calls, requireSuccess bool, client *ethclient.Client, signer SignerInterface, to *common.Address,
	funcSignature string, txReturnTypes []string, withValue bool, isMultiCall3Type bool,
) (Result, error) {
	arrayfiedCalls, msgValue, err := calls.ToArray(withValue, isMultiCall3Type)
	if err != nil {
		return Result{}, err
	}

	callData, err := abi.EncodeWithSignature(funcSignature, arrayfiedCalls, requireSuccess)
	if err != nil {
		return Result{}, err
	}

	tx, err := createTransaction(client, signer.GetAddress(), to, msgValue, callData)
	if err != nil {
		return Result{}, err
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return Result{}, err
	}

	signedTx, err := signer.SignTx(tx, chainId)
	if err != nil {
		return Result{}, err
	}

	encodedCallResult, err := client.CallContract(context.Background(), ethereum.CallMsg{
		From: *signer.GetAddress(),
		To:   to,
		Data: callData,
	}, nil)
	if err != nil {
		return Result{}, err
	}

	receipt, err := sendSignedTransaction(client, signedTx)
	if err != nil {
		return Result{}, err
	}

	decodedCallResult, err := abi.Decode(txReturnTypes, encodedCallResult)
	if err != nil {
		return Result{}, err
	}

	return parseResults(decodedCallResult, receipt.Status == 1, receipt)
}

func tryAggregate3(
	calls CallsWithFailure, client *ethclient.Client, signer SignerInterface, to *common.Address,
	funcSignature string, txReturnTypes []string, withValue bool, isMultiCall3Type bool,
) (Result, error) {
	arrayfiedCalls, msgValue, err := calls.ToArray(withValue, isMultiCall3Type)
	if err != nil {
		return Result{}, err
	}

	receipt, decodedCallResult, err := aggregateTx(
		arrayfiedCalls, msgValue, client, signer, to, funcSignature, txReturnTypes,
	)
	if err != nil {
		return Result{}, err
	}

	return parseResults(decodedCallResult, receipt.Status == 1, receipt)
}

func aggregateTx(
	arrayfiedCalls []any, msgValue *big.Int, client *ethclient.Client, signer SignerInterface,
	to *common.Address, funcSignature string, txReturnTypes []string,
) (*types.Receipt, []any, error) {

	callData, err := abi.EncodeWithSignature(funcSignature, arrayfiedCalls)
	if err != nil {
		return nil, nil, err
	}

	tx, err := createTransaction(client, signer.GetAddress(), to, msgValue, callData)
	if err != nil {
		return nil, nil, err
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil, nil, err
	}

	signedTx, err := signer.SignTx(tx, chainId)
	if err != nil {
		return nil, nil, err
	}

	encodedCallResult, err := client.CallContract(context.Background(), ethereum.CallMsg{
		From: *signer.GetAddress(),
		To:   to,
		Data: callData,
	}, nil)
	if err != nil {
		return nil, nil, err
	}

	receipt, err := sendSignedTransaction(client, signedTx)
	if err != nil {
		return nil, nil, err
	}

	decodedCallResult, err := abi.Decode(txReturnTypes, encodedCallResult)
	if err != nil {
		return nil, nil, err
	}

	return receipt, decodedCallResult, nil
}

func aggregateStatic(
	calls Calls, client *ethclient.Client, to *common.Address, funcSignature string, txReturnTypes []string,
	isSimulation bool, multiCallType *MultiCallType, writeAddress *common.Address, blockNumber *big.Int,
) (Result, error) {
	arrayfiedCalls, _, err := calls.ToArray(false, false)
	if err != nil {
		return Result{}, err
	}

	callData, err := abi.EncodeWithSignature(funcSignature, arrayfiedCalls)
	if err != nil {
		return Result{}, err
	}

	decodedCallResult, decodedAggregatedCallsResultVar, err := makeCall(
		calls,
		client,
		to,
		callData,
		txReturnTypes,
		isSimulation,
		multiCallType,
		writeAddress,
		blockNumber,
	)
	if err != nil {
		return Result{}, err
	}

	return parseResults(decodedAggregatedCallsResultVar, true, decodedCallResult)
}

func tryAggregateStatic(
	calls Calls, requireSuccess bool, client *ethclient.Client, to *common.Address, funcSignature string,
	txReturnTypes []string, multiCallType *MultiCallType, writeAddress *common.Address, blockNumber *big.Int,
) (Result, error) {
	arrayfiedCalls, _, err := calls.ToArray(false, false)
	if err != nil {
		return Result{}, err
	}

	callData, err := abi.EncodeWithSignature(funcSignature, arrayfiedCalls, requireSuccess)
	if err != nil {
		return Result{}, err
	}

	decodedCallResult, decodedAggregatedCallsResultVar, err := makeCall(
		calls,
		client,
		to,
		callData,
		txReturnTypes,
		false,
		multiCallType,
		writeAddress,
		blockNumber,
	)
	if err != nil {
		return Result{}, err
	}

	return parseResults(decodedAggregatedCallsResultVar, true, decodedCallResult)
}

func tryAggregateStatic3(
	calls CallsWithFailure, client *ethclient.Client, to *common.Address, funcSignature string,
	txReturnTypes []string, multiCallType *MultiCallType, writeAddress *common.Address, blockNumber *big.Int,
) (Result, error) {
	arrayfiedCalls, _, err := calls.ToArray(false, false)
	if err != nil {
		return Result{}, err
	}

	callData, err := abi.EncodeWithSignature(funcSignature, arrayfiedCalls)
	if err != nil {
		return Result{}, err
	}

	decodedCallResult, decodedAggregatedCallsResultVar, err := makeCall(
		calls,
		client,
		to,
		callData,
		txReturnTypes,
		false,
		multiCallType,
		writeAddress,
		blockNumber,
	)
	if err != nil {
		return Result{}, err
	}

	return parseResults(decodedAggregatedCallsResultVar, true, decodedCallResult)
}

func getCodeLengths(
	addresses []*common.Address, client *ethclient.Client, to *common.Address, blockNumber *big.Int,
) (Result, error) {
	return getData(
		addresses, client, to, "getCodeLengths(address[])", []string{"uint256[]"}, blockNumber,
	)
}

func getBalances(
	addresses []*common.Address, client *ethclient.Client, to *common.Address, blockNumber *big.Int,
) (Result, error) {
	return getData(
		addresses, client, to, "getBalances(address[])", []string{"uint256[]"}, blockNumber,
	)
}

func getAddressesData(
	addresses []*common.Address, client *ethclient.Client, to *common.Address, blockNumber *big.Int,
) (Result, error) {
	return getData(
		addresses,
		client,
		to,
		"getAddressesData(address[])",
		[]string{"uint256[]", "uint256[]"},
		blockNumber,
	)
}

func getChainData(
	client *ethclient.Client, to *common.Address, blockNumber *big.Int,
) (Result, error) {
	return getData(nil, client, to, "getChainData()", []string{
		"uint256",
		"uint256",
		"bytes32",
		"uint256",
		"address",
		"uint256",
		"uint256",
		"uint256",
		"uint256",
	},
		blockNumber,
	)
}

func getData(
	addresses []*common.Address, client *ethclient.Client, to *common.Address,
	funcSignature string, returnTypes []string, blockNumber *big.Int,
) (Result, error) {

	var callData []byte
	var err error
	if addresses != nil {
		callData, err = abi.EncodeWithSignature(funcSignature, toAnyArray(addresses))
	} else {
		callData, err = abi.EncodeWithSignature(funcSignature)
	}
	if err != nil {
		return Result{}, err
	}

	encodedCallResult, err := readContract(client, &ZERO_ADDRESS, to, callData, blockNumber)
	if err != nil {
		return Result{}, err
	}

	decodedCallResult, err := abi.Decode(returnTypes, encodedCallResult)
	if err != nil {
		return Result{}, err
	}

	return Result{Success: true, Result: decodedCallResult}, nil
}

func makeCall(
	calls CallsInterface, client *ethclient.Client, to *common.Address, callData []byte, txReturnTypes []string,
	isSimulation bool, multiCallType *MultiCallType, writeAddress *common.Address, blockNumber *big.Int,
) ([]any, []any, error) {
	if !true {
		log.Println(writeAddress)
	}
	encodedCallResult, err := readContract(client, &ZERO_ADDRESS, to, callData, blockNumber)
	if err != nil && !isSimulation {
		return nil, nil, err
	} else if isSimulation {
		log.Println(err)
	}
	if len(encodedCallResult) == 0 {
		*multiCallType = DEPLOYLESS
		writeAddress = nil
	}

	decodedCallResult, err := abi.Decode(txReturnTypes, encodedCallResult)
	if err != nil {
		return nil, nil, err
	}

	for len(decodedCallResult) != calls.Len() {
		decodedCallResult = decodedCallResult[0].([]any)
	}

	decodedAggregatedCallsResultVar, err := decodeAggregateCallsResult(decodedCallResult, calls)
	if err != nil {
		return nil, nil, err
	}

	return decodedCallResult, decodedAggregatedCallsResultVar, nil
}

func parseResults(
	decodedCallResult []any, status bool, callOrTxResult any,
) (Result, error) {
	var result Result
	if len(decodedCallResult) > 0 {
		result = Result{
			Success: status,
			Result:  decodedCallResult,
			Error:   nil,
		}
	} else {
		result = Result{
			Success: status,
			Result:  callOrTxResult,
			Error:   nil,
		}
	}

	return result, nil
}

func decodeAggregateCallsResult(result []any, calls CallsInterface) ([]any, error) {
	var decodedResult []any
	for i, res := range result {
		returnTypes := calls.GetReturnTypes(i)
		if returnTypes != nil {
			r, ok := res.([]byte)
			if ok {
				decodedR, err := abi.Decode(returnTypes, r)
				if err != nil {
					return nil, err
				}

				decodedResult = append(decodedResult, decodedR)
			} else {
				decodedResult = append(decodedResult, res)
			}
		} else {
			decodedResult = append(decodedResult, res.([]byte))
		}
	}

	return decodedResult, nil
}
