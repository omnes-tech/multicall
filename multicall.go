package multicall

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type MultiCallClient struct {
	RPC           string
	MultiCallType MultiCallType
	WriteAddress  *common.Address
	ReadAddress   *common.Address
	Signer        *SignerInterface
}

func NewClient(multiCallType MultiCallType, rpc string, signer *SignerInterface) (*MultiCallClient, error) {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	if multiCallType > 1 {
		return nil, fmt.Errorf("invalid multi call type %d", multiCallType)

	}
	if (multiCallType == OMNES && OMNES_MULTICALL_ADDRESS == common.Address{}) {
		log.Println("no OMNES address found. Using GENERAL address")
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
			log.Println("no deployed contract found. Using deployless method")

			writeAddress = nil
			multiCallType = DEPLOYLESS
		}

		toDeployless := writeAddress.Cmp(OMNES_MULTICALL_ADDRESS) == 0
		contractDeployed, newAddress, err := isContract(client, writeAddress, toDeployless)
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

		return &MultiCallClient{
			RPC:           rpc,
			MultiCallType: multiCallType,
			WriteAddress:  writeAddress,
			ReadAddress:   writeAddress,
			Signer:        signer,
		}, nil
	} else {
		toDeployless := writeAddress.Cmp(OMNES_MULTICALL_ADDRESS) == 0
		contractDeployed, newAddress, err := isContract(client, writeAddress, toDeployless)
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
		contractDeployed, newAddress, err = isContract(client, readAddress, toDeployless)
		if err != nil {
			return nil, fmt.Errorf("error checking contract: %v", err)
		}

		if !contractDeployed {
			readAddress = newAddress
		}

		return &MultiCallClient{
			RPC:           rpc,
			MultiCallType: multiCallType,
			WriteAddress:  writeAddress,
			ReadAddress:   readAddress,
			Signer:        signer,
		}, nil
	}
}

func (m *MultiCallClient) AggregateCalls(calls []Call) (Result, error) {

	if m.MultiCallType == GENERAL {
		return aggregate(
			calls,
			m.RPC,
			*m.Signer,
			m.WriteAddress,
			"aggregate((address,bytes)[])",
			[]string{"bytes[]"},
			false,
		)
	} else if m.MultiCallType == OMNES {
		return aggregate(
			calls,
			m.RPC,
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

func (m *MultiCallClient) TryAggregateCalls(calls []Call, requireSuccess bool) (Result, error) {

	if m.MultiCallType == GENERAL {
		return Result{}, fmt.Errorf("cannot do call with multi call type %d", m.MultiCallType)
	} else if m.MultiCallType == OMNES {
		return tryAggregate(
			calls,
			requireSuccess,
			m.RPC,
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

func (m *MultiCallClient) TryAggregateCalls3(calls []CallWithFailure) (Result, error) {

	if m.MultiCallType == GENERAL {
		withValue, funcSignature := isWithValue(calls)

		return tryAggregate3(
			calls,
			m.RPC,
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
			m.RPC,
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

func (m *MultiCallClient) SimulateCall(calls []Call) (Result, error) {

	if m.MultiCallType == GENERAL {
		return deploylessSimulation(calls, m.RPC)
	} else if m.MultiCallType == OMNES {
		return aggregateStatic(
			calls,
			m.RPC,
			m.WriteAddress,
			"simulateCalls((address,bytes)[])",
			nil,
			true,
			&m.MultiCallType,
			m.WriteAddress,
		)
	} else {
		return deploylessSimulation(calls, m.RPC)
	}
}

func (m *MultiCallClient) AggregateStatic(calls []Call) (Result, error) {

	if m.MultiCallType == GENERAL {
		return deploylessAggregateStatic(calls, m.RPC)
	} else if m.MultiCallType == OMNES {
		return aggregateStatic(
			calls,
			m.RPC,
			m.WriteAddress,
			"aggregateStatic((address,bytes)[])",
			[]string{"bytes[]"},
			false,
			&m.MultiCallType,
			m.WriteAddress,
		)
	} else {
		return deploylessAggregateStatic(calls, m.RPC)
	}
}

func (m *MultiCallClient) TryAggregateStatic(calls []Call, requireSuccess bool) (Result, error) {

	if m.MultiCallType == GENERAL {
		return deploylessTryAggregateStatic(calls, requireSuccess, m.RPC)
	} else if m.MultiCallType == OMNES {
		return tryAggregateStatic(
			calls,
			requireSuccess,
			m.RPC,
			m.WriteAddress,
			"aggregateStatic((address,bytes)[])",
			[]string{"bytes[]"},
			&m.MultiCallType,
			m.WriteAddress,
		)
	} else {
		return deploylessTryAggregateStatic(calls, requireSuccess, m.RPC)
	}
}

func (m *MultiCallClient) TryAggregateStatic3(calls []CallWithFailure) (Result, error) {

	if m.MultiCallType == GENERAL {
		return deploylessTryAggregateStatic3(calls, m.RPC)
	} else if m.MultiCallType == OMNES {
		return tryAggregateStatic3(
			calls,
			m.RPC,
			m.WriteAddress,
			"aggregateStatic((address,bytes)[])",
			[]string{"bytes[]"},
			&m.MultiCallType,
			m.WriteAddress,
		)
	} else {
		return deploylessTryAggregateStatic3(calls, m.RPC)
	}
}

func (m *MultiCallClient) CodeLengths(addresses []*common.Address) (Result, error) {

	if m.MultiCallType == GENERAL {
		return deploylessGetCodeLengths(addresses, m.RPC)
	} else if m.MultiCallType == OMNES {
		return getCodeLengths(addresses, m.RPC, m.ReadAddress)
	} else {
		return deploylessGetCodeLengths(addresses, m.RPC)
	}
}

func (m *MultiCallClient) Balances(addresses []*common.Address) (Result, error) {

	if m.MultiCallType == GENERAL {
		return deploylessGetBalances(addresses, m.RPC)
	} else if m.MultiCallType == OMNES {
		return getBalances(addresses, m.RPC, m.ReadAddress)
	} else {
		return deploylessGetBalances(addresses, m.RPC)
	}
}

func (m *MultiCallClient) AddressesData(addresses []*common.Address) (Result, error) {

	if m.MultiCallType == GENERAL {
		return deploylessGetAddressesData(addresses, m.RPC)
	} else if m.MultiCallType == OMNES {
		return getAddressesData(addresses, m.RPC, m.ReadAddress)
	} else {
		return deploylessGetAddressesData(addresses, m.RPC)
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

func isContract(client *ethclient.Client, address *common.Address, toDeployless bool) (bool, *common.Address, error) {
	bytecode, err := client.CodeAt(context.Background(), *address, nil)
	if err != nil {
		return false, nil, fmt.Errorf("error getting bytecode: %v", err)
	}

	if len(bytecode) == 0 {
		if toDeployless {
			log.Println("no deployed contract found. Using deployless method")

			return false, nil, nil
		} else {
			log.Println("no deployed contract found. Checking Omnes contract")

			return isContract(client, &OMNES_MULTICALL_ADDRESS, true)
		}
	}

	return true, address, nil

}
