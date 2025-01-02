package multicall

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/omnes-tech/abi"
)

type MultiCallType uint8

const (
	GENERAL = iota
	OMNES
	DEPLOYLESS
)

type Result struct {
	Success bool
	Result  any
	Error   error
}

type commonCall struct {
	Target        common.Address
	FuncSignature string
	Args          []interface{}
	ReturnTypes   []string
}

type Call struct {
	commonCall
	Value *big.Int
}

type CallWithFailure struct {
	Call
	RequireSuccess bool
}

type StaticCall commonCall

type StaticCallWithFailure struct {
	StaticCall
	RequireSuccess bool
}

type CallsInterface interface {
	GetTarget(i int) *common.Address
	GetFuncSignature(i int) string
	GetArgs(i int) []interface{}
	GetReturnTypes(i int) []string
	GetValue(i int) *big.Int
	Len() int
	ToArray(withValue bool, isMultiCall3Type bool) ([]any, *big.Int, error)
}

type CallsWithFailureInterface interface {
	CallsInterface
	GetRequireSuccess(i int) bool
}

type Calls []Call
type CallsWithFailure []CallWithFailure

func NewCall(
	target common.Address, funcSignature string,
	args []interface{}, returnTypes []string, value *big.Int,
) Call {
	return Call{
		commonCall: commonCall{
			Target:        target,
			FuncSignature: funcSignature,
			Args:          args,
			ReturnTypes:   returnTypes,
		},
		Value: value,
	}
}

func NewCalls(
	targets []common.Address, funcSignatures []string,
	argss [][]interface{}, returnTypess [][]string, values []*big.Int,
) Calls {
	calls := make(Calls, len(targets))
	for i, target := range targets {
		calls[i] = NewCall(target, funcSignatures[i], argss[i], returnTypess[i], values[i])
	}

	return calls
}

func NewCallWithFailure(
	target common.Address, funcSignature string, args []interface{},
	returnTypes []string, value *big.Int, requireSuccess bool,
) CallWithFailure {
	return CallWithFailure{
		Call: Call{
			commonCall: commonCall{
				Target:        target,
				FuncSignature: funcSignature,
				Args:          args,
				ReturnTypes:   returnTypes,
			},
			Value: value,
		},
		RequireSuccess: requireSuccess,
	}
}

func NewCallsWithFailure(
	targets []common.Address, funcSignatures []string, argss [][]interface{},
	returnTypess [][]string, values []*big.Int, requireSuccesss []bool,
) CallsWithFailure {
	calls := make(CallsWithFailure, len(targets))
	for i, target := range targets {
		calls[i] = NewCallWithFailure(target, funcSignatures[i], argss[i], returnTypess[i], values[i], requireSuccesss[i])
	}

	return calls
}

func ParseCallToCalls(calls []Call) Calls {
	result := make(Calls, len(calls))
	for i, c := range calls {
		result[i] = Call{
			commonCall: commonCall{
				Target:        c.Target,
				FuncSignature: c.FuncSignature,
				Args:          c.Args,
				ReturnTypes:   c.ReturnTypes,
			},
			Value: c.Value,
		}
	}
	return result
}

func ParseCallWithFailureToCallsWithFailure(calls []CallWithFailure) CallsWithFailure {
	result := make(CallsWithFailure, len(calls))
	for i, c := range calls {
		result[i] = CallWithFailure{
			Call: Call{
				commonCall: commonCall{
					Target:        c.Target,
					FuncSignature: c.FuncSignature,
					Args:          c.Args,
					ReturnTypes:   c.ReturnTypes,
				},
				Value: c.Value,
			},
			RequireSuccess: c.RequireSuccess,
		}
	}
	return result
}

func (c Calls) GetTarget(i int) *common.Address {
	return &c[i].Target
}

func (c Calls) GetFuncSignature(i int) string {
	return c[i].FuncSignature
}

func (c Calls) GetArgs(i int) []interface{} {
	return c[i].Args
}

func (c Calls) GetReturnTypes(i int) []string {
	return c[i].ReturnTypes
}

func (c Calls) GetValue(i int) *big.Int {
	return c[i].Value
}

func (c Calls) Len() int {
	return len(c)
}

func (c Calls) ToArray(withValue bool, isMultiCall3Type bool) ([]any, *big.Int, error) {
	var result []any
	summed := big.NewInt(0)
	for i := 0; i < c.Len(); i++ {
		var args []any
		args = append(args, c.GetTarget(i))

		var callData []byte
		var err error
		if c.GetArgs(i) != nil || len(c.GetArgs(i)) > 0 {
			callData, err = abi.EncodeWithSignature(c.GetFuncSignature(i), c.GetArgs(i)...)
		} else {
			callData, err = abi.EncodeWithSignature(c.GetFuncSignature(i))
		}
		if err != nil {
			return nil, nil, err
		}
		args = append(args, callData)

		if withValue {
			value := big.NewInt(0)
			if c.GetValue(i) != nil {
				value.Add(value, c.GetValue(i))
			}
			summed.Add(summed, value)

			args = append(args, value)
		}

		result = append(result, args)
	}

	return result, summed, nil
}

func (c CallsWithFailure) GetTarget(i int) *common.Address {
	return &c[i].Target
}

func (c CallsWithFailure) GetFuncSignature(i int) string {
	return c[i].FuncSignature
}

func (c CallsWithFailure) GetArgs(i int) []interface{} {
	return c[i].Args
}

func (c CallsWithFailure) GetReturnTypes(i int) []string {
	return c[i].ReturnTypes
}

func (c CallsWithFailure) GetValue(i int) *big.Int {
	return c[i].Value
}

func (c CallsWithFailure) GetRequireSuccess(i int) bool {
	return c[i].RequireSuccess
}

func (c CallsWithFailure) Len() int {
	return len(c)
}

func (c CallsWithFailure) ToArray(withValue bool, isMultiCall3Type bool) ([]any, *big.Int, error) {
	// Omnes: (address,bytes,uint256,bool)
	// MultiCall3: (address,bool,bytes) or (address,bool,uint256,bytes)

	var result []any
	summed := big.NewInt(0)
	for i := 0; i < c.Len(); i++ {
		var args []any

		var callData []byte
		var err error
		if c.GetArgs(i) != nil || len(c.GetArgs(i)) > 0 {
			callData, err = abi.EncodeWithSignature(c.GetFuncSignature(i), c.GetArgs(i)...)
		} else {
			callData, err = abi.EncodeWithSignature(c.GetFuncSignature(i))
		}
		if err != nil {
			return nil, nil, err
		}

		if isMultiCall3Type {
			args = append(args, c.GetTarget(i))
			args = append(args, c.GetRequireSuccess(i))
			if withValue {
				value := big.NewInt(0)
				if c.GetValue(i) != nil {
					value.Add(value, c.GetValue(i))
				}
				summed.Add(summed, value)

				args = append(args, value)
			}
			args = append(args, callData)
		} else {
			args = append(args, c.GetTarget(i))
			args = append(args, callData)
			if withValue {
				value := big.NewInt(0)
				if c.GetValue(i) != nil {
					value.Add(value, c.GetValue(i))
				}
				summed.Add(summed, value)

				args = append(args, value)
			}
			args = append(args, c.GetRequireSuccess(i))
		}

		result = append(result, args)
	}

	return result, summed, nil
}
