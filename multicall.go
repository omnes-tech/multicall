package multicall

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type MultiCall struct {
	MultiCallType MultiCallType
	WriteAddress  *common.Address
	ReadAddress   *common.Address
	Signer        *SignerInterface
}

func NewMultiCall(multiCallType MultiCallType, client *ethclient.Client, signer *SignerInterface) (*MultiCall, error) {
	if multiCallType > 1 {
		return nil, fmt.Errorf("invalid multi call type %d", multiCallType)

	}
	if (multiCallType == OMNES && OMNES_MULTICALL_ADDRESS == common.Address{}) {
		log.Printf("no OMNES address found. Using GENERAL address\n\n")
		multiCallType = GENERAL
	}

	var writeAddress *common.Address
	var readAddress *common.Address
	if multiCallType == GENERAL {
		writeAddress = &GENERAL_MULTICALL_ADDRESS
		readAddress = &OMNES_MULTICALL_ADDRESS

	} else if multiCallType == OMNES {
		writeAddress = &OMNES_MULTICALL_ADDRESS
		readAddress = &OMNES_MULTICALL_ADDRESS
	}

	if writeAddress.Cmp(*readAddress) == 0 {
		bytecode, err := client.CodeAt(context.Background(), *writeAddress, nil)
		if err != nil {
			return nil, fmt.Errorf("error getting bytecode: %v", err)
		}

		if len(bytecode) == 0 {
			log.Printf("no deployed contract found. Using deployless method\n\n")

			writeAddress = nil
			multiCallType = DEPLOYLESS
		}

		toDeployless := writeAddress.Cmp(OMNES_MULTICALL_ADDRESS) == 0
		contractDeployed, newAddress, err := isContract(client, writeAddress, toDeployless, false)
		if err != nil {
			return nil, fmt.Errorf("error checking contract: %v", err)
		}

		if !contractDeployed {
			writeAddress = newAddress
			if toDeployless {
				multiCallType = DEPLOYLESS
			} else {
				multiCallType = OMNES
			}
		}

		return &MultiCall{
			MultiCallType: multiCallType,
			WriteAddress:  writeAddress,
			ReadAddress:   writeAddress,
			Signer:        signer,
		}, nil
	} else {
		toDeployless := writeAddress.Cmp(OMNES_MULTICALL_ADDRESS) == 0
		contractDeployed, newAddress, err := isContract(client, writeAddress, toDeployless, false)
		if err != nil {
			return nil, fmt.Errorf("error checking contract: %v", err)
		}

		if !contractDeployed {
			writeAddress = newAddress
			if toDeployless {
				multiCallType = DEPLOYLESS
			} else {
				multiCallType = OMNES
			}
		}

		toDeployless = readAddress.Cmp(OMNES_MULTICALL_ADDRESS) == 0
		contractDeployed, newAddress, err = isContract(client, readAddress, toDeployless, true)
		if err != nil {
			return nil, fmt.Errorf("error checking contract: %v", err)
		}

		if !contractDeployed {
			readAddress = newAddress
		}

		return &MultiCall{
			MultiCallType: multiCallType,
			WriteAddress:  writeAddress,
			ReadAddress:   readAddress,
			Signer:        signer,
		}, nil
	}
}

func (m *MultiCall) AggregateCalls(
	calls []Call, client *ethclient.Client, blockNumber *big.Int,
) (Result, error) {
	if m.Signer == nil {
		return Result{}, fmt.Errorf("no signer configured")
	}

	if m.MultiCallType == GENERAL {
		return aggregate(
			calls,
			client,
			*m.Signer,
			m.WriteAddress,
			"aggregate((address,bytes)[])",
			[]string{"bytes[]"},
			false,
		)
	} else if m.MultiCallType == OMNES {
		return aggregate(
			calls,
			client,
			*m.Signer,
			m.WriteAddress,
			"aggregateCalls((address,bytes,uint256)[])",
			[]string{"bytes[]"},
			true,
		)
	} else {
		return Result{}, fmt.Errorf("cannot do call with multi call type %d", m.MultiCallType)
	}
}

func (m *MultiCall) TryAggregateCalls(
	calls []Call, requireSuccess bool, client *ethclient.Client, blockNumber *big.Int,
) (Result, error) {
	if m.Signer == nil {
		return Result{}, fmt.Errorf("no signer configured")
	}

	if m.MultiCallType == GENERAL {
		return Result{}, fmt.Errorf("cannot do call with multi call type %d", m.MultiCallType)
	} else if m.MultiCallType == OMNES {
		return tryAggregate(
			calls,
			requireSuccess,
			client,
			*m.Signer,
			m.WriteAddress,
			"tryAggregateCalls((address,bytes,uint256)[],bool)",
			[]string{"(bool,bytes)[]"},
			true,
			false,
		)
	} else {
		return Result{}, fmt.Errorf("cannot do call with multi call type %d", m.MultiCallType)
	}
}

func (m *MultiCall) TryAggregateCalls3(
	calls []CallWithFailure, client *ethclient.Client, blockNumber *big.Int,
) (Result, error) {
	if m.Signer == nil {
		return Result{}, fmt.Errorf("no signer configured")
	}

	if m.MultiCallType == GENERAL {
		withValue, funcSignature := isWithValue(calls)

		return tryAggregate3(
			calls,
			client,
			*m.Signer,
			m.WriteAddress,
			funcSignature,
			[]string{"(bool,bytes)[]"},
			withValue,
			true,
		)
	} else if m.MultiCallType == OMNES {
		return tryAggregate3(
			calls,
			client,
			*m.Signer,
			m.WriteAddress,
			"tryAggregateCalls((address,bytes,uint256,bool)[])",
			[]string{"(bool,bytes)[]"},
			true,
			false,
		)
	} else {
		return Result{}, fmt.Errorf("cannot do call with multi call type %d", m.MultiCallType)
	}
}

func (m *MultiCall) SimulateCall(
	calls []Call, client *ethclient.Client, blockNumber *big.Int,
) (Result, error) {

	if m.MultiCallType == GENERAL {
		return deploylessSimulation(calls, client, blockNumber)
	} else if m.MultiCallType == OMNES {
		return aggregateStatic(
			calls,
			client,
			m.WriteAddress,
			"simulateCalls((address,bytes)[])",
			nil,
			true,
			&m.MultiCallType,
			m.WriteAddress,
			blockNumber,
		)
	} else {
		return deploylessSimulation(calls, client, blockNumber)
	}
}

func (m *MultiCall) AggregateStatic(
	calls []Call, client *ethclient.Client, blockNumber *big.Int,
) (Result, error) {

	if m.MultiCallType == GENERAL {
		return deploylessAggregateStatic(calls, client, blockNumber)
	} else if m.MultiCallType == OMNES {
		return aggregateStatic(
			calls,
			client,
			m.WriteAddress,
			"aggregateStatic((address,bytes)[])",
			[]string{"bytes[]"},
			false,
			&m.MultiCallType,
			m.WriteAddress,
			blockNumber,
		)
	} else {
		return deploylessAggregateStatic(calls, client, blockNumber)
	}
}

func (m *MultiCall) TryAggregateStatic(
	calls []Call, requireSuccess bool, client *ethclient.Client, blockNumber *big.Int,
) (Result, error) {

	if m.MultiCallType == GENERAL {
		return deploylessTryAggregateStatic(calls, requireSuccess, client, blockNumber)
	} else if m.MultiCallType == OMNES {
		return tryAggregateStatic(
			calls,
			requireSuccess,
			client,
			m.WriteAddress,
			"aggregateStatic((address,bytes)[],bool)",
			[]string{"bytes[]"},
			&m.MultiCallType,
			m.WriteAddress,
			blockNumber,
		)
	} else {
		return deploylessTryAggregateStatic(calls, requireSuccess, client, blockNumber)
	}
}

func (m *MultiCall) TryAggregateStatic3(
	calls []CallWithFailure, client *ethclient.Client, blockNumber *big.Int,
) (Result, error) {

	if m.MultiCallType == GENERAL {
		return deploylessTryAggregateStatic3(calls, client, blockNumber)
	} else if m.MultiCallType == OMNES {
		return tryAggregateStatic3(
			calls,
			client,
			m.WriteAddress,
			"aggregateStatic((address,bytes,bool)[])",
			[]string{"bytes[]"},
			&m.MultiCallType,
			m.WriteAddress,
			blockNumber,
		)
	} else {
		return deploylessTryAggregateStatic3(calls, client, blockNumber)
	}
}

func (m *MultiCall) CodeLengths(
	addresses []*common.Address, client *ethclient.Client, blockNumber *big.Int,
) (Result, error) {

	if m.MultiCallType == GENERAL {
		return deploylessGetCodeLengths(addresses, client, blockNumber)
	} else if m.MultiCallType == OMNES {
		return getCodeLengths(addresses, client, m.ReadAddress, blockNumber)
	} else {
		return deploylessGetCodeLengths(addresses, client, blockNumber)
	}
}

func (m *MultiCall) Balances(
	addresses []*common.Address, client *ethclient.Client, blockNumber *big.Int,
) (Result, error) {

	if m.MultiCallType == GENERAL {
		return deploylessGetBalances(addresses, client, blockNumber)
	} else if m.MultiCallType == OMNES {
		return getBalances(addresses, client, m.ReadAddress, blockNumber)
	} else {
		return deploylessGetBalances(addresses, client, blockNumber)
	}
}

func (m *MultiCall) AddressesData(
	addresses []*common.Address, client *ethclient.Client, blockNumber *big.Int,
) (Result, error) {

	if m.MultiCallType == GENERAL {
		return deploylessGetAddressesData(addresses, client, blockNumber)
	} else if m.MultiCallType == OMNES {
		return getAddressesData(addresses, client, m.ReadAddress, blockNumber)
	} else {
		return deploylessGetAddressesData(addresses, client, blockNumber)
	}
}

func (m *MultiCall) ChainData(client *ethclient.Client, blockNumber *big.Int) (Result, error) {

	if m.MultiCallType == GENERAL {
		return deploylessGetChainData(client, blockNumber)
	} else if m.MultiCallType == OMNES {
		return getChainData(client, m.ReadAddress, blockNumber)
	} else {
		return deploylessGetChainData(client, blockNumber)
	}
}

func isWithValue(calls []CallWithFailure) (bool, string) {
	for _, call := range calls {
		if call.Value != nil {
			return true, "aggregate3Value((address,bool,uint256,bytes)[])"
		}
	}

	return false, "aggregate3((address,bool,bytes)[])"
}

func isContract(client *ethclient.Client, address *common.Address, toDeployless bool, justForReading bool) (bool, *common.Address, error) {
	bytecode, err := client.CodeAt(context.Background(), *address, nil)
	if err != nil {
		return false, nil, fmt.Errorf("error getting bytecode: %v", err)
	}

	if len(bytecode) == 0 {
		var logMsg string

		if toDeployless {
			logMsg = "no deployed contract found. Using deployless method"
		} else {
			logMsg = "no deployed contract found. Checking Omnes contract"
		}
		if justForReading {
			logMsg += " (reading)"
		}
		logMsg += "\n\n"

		if toDeployless {
			return false, nil, nil
		}

		return isContract(client, &OMNES_MULTICALL_ADDRESS, true, justForReading)
	}

	return true, address, nil

}
