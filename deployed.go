package multicall

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/omnes-tech/abi"
)

func transactWithFailure(
	calls CallsWithFailure, requireSuccess bool, client *ethclient.Client,
	signer SignerInterface, to *common.Address, funcSignature string, txReturnTypes []string,
	withValue bool, isMultiCall3Type bool,
) Result {
	return write(
		calls,
		requireSuccess,
		client,
		signer,
		to,
		funcSignature,
		txReturnTypes,
		withValue,
		isMultiCall3Type,
	)
}

func transact(
	calls Calls, requireSuccess bool, client *ethclient.Client,
	signer SignerInterface, to *common.Address, funcSignature string, txReturnTypes []string,
	withValue bool, isMultiCall3Type bool,
) Result {
	return write(
		calls,
		requireSuccess,
		client,
		signer,
		to,
		funcSignature,
		txReturnTypes,
		withValue,
		isMultiCall3Type,
	)
}

func write(
	calls CallsInterface, requireSuccess bool, client *ethclient.Client, signer SignerInterface,
	to *common.Address, funcSignature string, txReturnTypes []string, withValue bool, isMultiCall3Type bool,
) Result {
	arrayfiedCalls, msgValue, err := calls.ToArray(withValue, isMultiCall3Type)
	if err != nil {
		return Result{Success: false, Error: err}
	}

	var callData []byte
	if funcSignature == "tryAggregateCalls((address,bytes,uint256)[],bool)" {
		callData, err = abi.EncodeWithSignature(funcSignature, arrayfiedCalls, requireSuccess)
	} else {
		callData, err = abi.EncodeWithSignature(funcSignature, arrayfiedCalls)
	}
	if err != nil {
		return Result{Success: false, Error: err}
	}

	tx, err := createTransaction(client, signer.GetAddress(), to, msgValue, callData)
	if err != nil {
		return Result{Success: false, Error: err}
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return Result{Success: false, Error: err}
	}

	signedTx, err := signer.SignTx(tx, chainId)
	if err != nil {
		return Result{Success: false, Error: err}
	}

	encodedCallResult, err := client.CallContract(context.Background(), ethereum.CallMsg{
		From: *signer.GetAddress(),
		To:   to,
		Data: callData,
	}, nil)
	if err != nil {
		return Result{Success: false, Error: fmt.Errorf("error calling contract: %w, with data: %s", err, common.Bytes2Hex(callData))}
	}

	receipt, err := sendSignedTransaction(client, signedTx)
	if err != nil {
		return Result{Success: false, Error: fmt.Errorf("error sending signed transaction: %w", err)}
	}

	decodedCallResult, err := abi.Decode(txReturnTypes, encodedCallResult)
	if err != nil {
		return Result{Success: false, Error: err}
	}

	return parseResults(decodedCallResult, receipt.Status == 1, receipt)
}

func txAsReadWithFailure(
	calls CallsWithFailure, requireSuccess bool, client *ethclient.Client, to *common.Address,
	funcSignature string, txReturnTypes []string, multiCallType *MultiCallType,
) Result {
	return asRead(
		calls,
		requireSuccess,
		client,
		to,
		funcSignature,
		txReturnTypes,
		multiCallType,
	)
}

func txAsRead(
	calls Calls, requireSuccess bool, client *ethclient.Client, to *common.Address,
	funcSignature string, txReturnTypes []string, multiCallType *MultiCallType,
) Result {
	return asRead(
		calls,
		requireSuccess,
		client,
		to,
		funcSignature,
		txReturnTypes,
		multiCallType,
	)
}

func asRead(
	calls CallsInterface, requireSuccess bool, client *ethclient.Client, to *common.Address,
	funcSignature string, txReturnTypes []string, multiCallType *MultiCallType,
) Result {
	arrayfiedCalls, _, err := calls.ToArray(true, false)
	if err != nil {
		return Result{Success: false, Error: err}
	}

	var callData []byte
	if funcSignature == "tryAggregateCalls((address,bytes,uint256)[],bool)" {
		callData, err = abi.EncodeWithSignature(funcSignature, arrayfiedCalls, requireSuccess)
	} else {
		callData, err = abi.EncodeWithSignature(funcSignature, arrayfiedCalls)
	}
	if err != nil {
		return Result{Success: false, Error: err}
	}

	decodedCallResult, decodedAggregatedCallsResultVar, err := makeCall(
		calls,
		client,
		to,
		callData,
		txReturnTypes,
		false,
		multiCallType,
		nil,
		nil,
	)
	if err != nil {
		return Result{Success: false, Error: err}
	}

	return parseResults(decodedAggregatedCallsResultVar, true, decodedCallResult)
}

func call(
	calls Calls, requireSuccess bool, client *ethclient.Client, to *common.Address, funcSignature string,
	txReturnTypes []string, multiCallType *MultiCallType, writeAddress *common.Address,
	blockNumber *big.Int, isSimulation bool,
) Result {
	return read(
		calls,
		requireSuccess,
		client,
		to,
		funcSignature,
		txReturnTypes,
		multiCallType,
		writeAddress,
		blockNumber,
		isSimulation,
	)
}

func callWithFailure(
	calls CallsWithFailure, client *ethclient.Client, to *common.Address, funcSignature string,
	txReturnTypes []string, multiCallType *MultiCallType, writeAddress *common.Address, blockNumber *big.Int,
) Result {
	return read(
		calls,
		false,
		client,
		to,
		funcSignature,
		txReturnTypes,
		multiCallType,
		writeAddress,
		blockNumber,
		false,
	)
}

func read(
	calls CallsInterface, requireSuccess bool, client *ethclient.Client, to *common.Address, funcSignature string,
	txReturnTypes []string, multiCallType *MultiCallType, writeAddress *common.Address, blockNumber *big.Int,
	isSimulation bool,
) Result {
	arrayfiedCalls, _, err := calls.ToArray(false, false)
	if err != nil {
		return Result{Success: false, Error: err}
	}

	if funcSignature == "tryAggregateStatic((address,bytes,bool)[])" {
		isSimulation = false
	}

	var callData []byte
	if funcSignature == "tryAggregateStatic((address,bytes)[],bool)" {
		isSimulation = false
		callData, err = abi.EncodeWithSignature(funcSignature, arrayfiedCalls, requireSuccess)
	} else {
		callData, err = abi.EncodeWithSignature(funcSignature, arrayfiedCalls)
	}
	if err != nil {
		return Result{Success: false, Error: err}
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
		return Result{Success: false, Error: err}
	}

	return parseResults(decodedAggregatedCallsResultVar, true, decodedCallResult)
}

func getData(
	addresses []*common.Address, client *ethclient.Client, to *common.Address,
	funcSignature string, returnTypes []string, blockNumber *big.Int,
) Result {

	var callData []byte
	var err error
	if addresses != nil {
		callData, err = abi.EncodeWithSignature(funcSignature, toAnyArray(addresses))
	} else {
		callData, err = abi.EncodeWithSignature(funcSignature)
	}
	if err != nil {
		return Result{Success: false, Error: err}
	}

	encodedCallResult, err := readContract(client, &ZERO_ADDRESS, to, callData, blockNumber)
	if err != nil {
		return Result{Success: false, Error: err}
	}

	decodedCallResult, err := abi.Decode(returnTypes, encodedCallResult)
	if err != nil {
		return Result{Success: false, Error: err}
	}

	return Result{Success: true, Result: decodedCallResult}
}

func makeCall(
	calls CallsInterface, client *ethclient.Client, to *common.Address, callData []byte, txReturnTypes []string,
	isSimulation bool, multiCallType *MultiCallType, writeAddress *common.Address, blockNumber *big.Int,
) ([]any, []any, error) {
	if !true {
		log.Println(writeAddress)
	}

	var decodedCallResult []any
	encodedCallResult, err := readContract(client, &ZERO_ADDRESS, to, callData, blockNumber)
	if err != nil && !isSimulation {
		return nil, nil, err
	} else if isSimulation {
		if strings.Contains(err.Error(), "execution reverted") {
			encodedRevert, ok := parseRevertData(err)
			if ok {
				decodedCallResult, err := abi.DecodeWithSignature(
					"MultiCall__Simulation((bool,bytes,uint256)[])",
					encodedRevert,
				)
				if err != nil {
					return nil, nil, err
				}
				decodedCallResult = decodedCallResult[0].([]any)

				for i, result := range decodedCallResult {
					decodedCallResult[i].([]any)[1] = common.Bytes2Hex(result.([]any)[1].([]byte))
				}
			}
		}
	}
	if len(encodedCallResult) == 0 {
		*multiCallType = DEPLOYLESS
		writeAddress = nil
	}

	if !isSimulation {
		decodedCallResult, err = abi.Decode(txReturnTypes, encodedCallResult)
		if err != nil {
			return nil, nil, err
		}
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
) Result {
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

	return result
}

func decodeAggregateCallsResult(result []any, calls CallsInterface) ([]any, error) {
	var decodedResult []any
	for i, res := range result {
		returnTypes := calls.GetReturnTypes(i)
		if returnTypes != nil || len(returnTypes) > 0 {
			r, ok := res.([]byte)
			if ok {
				decodedR, err := abi.Decode(returnTypes, r)
				if err != nil {
					return nil, err
				}

				decodedResult = append(decodedResult, decodedR)
			} else {
				decodedResult = append(decodedResult, res.([]any))
			}
		} else {
			decodedResult = append(decodedResult, res.([]any))
		}
	}

	return decodedResult, nil
}
