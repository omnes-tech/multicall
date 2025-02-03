package multicall_test

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/omnes-tech/multicall"
)

func ExampleNewClient() {
	rpc := "https://eth.llamarpc.com"
	client, err := ethclient.Dial(rpc)
	if err != nil {
		panic(err)
	}

	mcall, err := multicall.NewMultiCall(multicall.GENERAL, client, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("MultiCallType: %d WriteAddress: %s ReadAddress: %s",
		mcall.MultiCallType, mcall.WriteAddress, mcall.ReadAddress)

	// Output: MultiCallType: 0 WriteAddress: 0xcA11bde05977b3631167028862bE2a173976CA11 ReadAddress: <nil>
}

func ExampleMultiCall_SimulateCall() {
	rpc := "https://eth.llamarpc.com"
	client, err := ethclient.Dial(rpc)
	if err != nil {
		panic(err)
	}

	mcall, err := multicall.NewMultiCall(multicall.GENERAL, client, nil)
	if err != nil {
		panic(err)
	}

	targets := []common.Address{
		common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
		common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
	}
	funcSigs := []string{
		"deposit()",
		"deposit()",
	}
	values := []*big.Int{
		big.NewInt(1000000000000000000),
		big.NewInt(1000000000000000000),
	}

	calls := multicall.NewCalls(targets, funcSigs, nil, nil, nil, values)

	results, err := mcall.SimulateCall(calls, client, nil)

	fmt.Println(results)

	// Output: {true [[true  33921] [true  9521]] <nil>}
}

func ExampleMultiCall_AggregateStatic() {
	rpc := "https://eth.llamarpc.com"
	client, err := ethclient.Dial(rpc)
	if err != nil {
		panic(err)
	}

	mcall, err := multicall.NewMultiCall(multicall.GENERAL, client, nil)
	if err != nil {
		panic(err)
	}

	targets := []common.Address{
		common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
		common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
	}
	funcSigs := []string{
		"balanceOf(address)",
		"balanceOf(address)",
	}

	address := common.Address{}
	argss := [][]any{
		{&address},
		{&address},
	}
	returnTypes := [][]string{
		{"uint256"},
		{"uint256"},
	}

	calls := multicall.NewCalls(targets, funcSigs, argss, nil, returnTypes, nil)

	results, err := mcall.AggregateStatic(calls, client, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(results)

	// Output: {true [[1085420955917931147422] [1085420955917931147422]] <nil>}
}

func ExampleMultiCall_TryAggregateStatic() {
	rpc := "https://eth.llamarpc.com"
	client, err := ethclient.Dial(rpc)
	if err != nil {
		panic(err)
	}

	mcall, err := multicall.NewMultiCall(multicall.GENERAL, client, nil)
	if err != nil {
		panic(err)
	}

	targets := []common.Address{
		common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
		common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
	}
	funcSigs := []string{
		"balanceOf(address)",
		"balanceOf(address)",
	}

	address := common.Address{}
	argss := [][]any{
		{&address},
		{&address},
	}
	returnTypes := [][]string{
		{"uint256"},
		{"uint256"},
	}

	calls := multicall.NewCalls(targets, funcSigs, argss, nil, returnTypes, nil)

	results, err := mcall.TryAggregateStatic(calls, true, client, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(results)

	// Output: {true [[true [1085420955917931147422]] [true [1085420955917931147422]]] <nil>}
}

func ExampleMultiCall_TryAggregateStatic3() {
	rpc := "https://eth.llamarpc.com"
	client, err := ethclient.Dial(rpc)
	if err != nil {
		panic(err)
	}

	mcall, err := multicall.NewMultiCall(multicall.GENERAL, client, nil)
	if err != nil {
		panic(err)
	}

	targets := []common.Address{
		common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
		common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
	}
	funcSigs := []string{
		"balanceOf(address)",
		"balanceOf(address)",
	}

	address := common.Address{}
	argss := [][]any{
		{&address},
		{&address},
	}
	returnTypes := [][]string{
		{"uint256"},
		{"uint256"},
	}
	requireSuccess := []bool{true, true}

	calls := multicall.NewCallsWithFailure(targets, funcSigs, argss, nil, returnTypes, nil, requireSuccess)

	results, err := mcall.TryAggregateStatic3(calls, client, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(results)

	// Output: {true [[true [1085420955917931147422]] [true [1085420955917931147422]]] <nil>}
}

func ExampleMultiCall_CodeLengths() {
	rpc := "https://eth.llamarpc.com"
	client, err := ethclient.Dial(rpc)
	if err != nil {
		panic(err)
	}

	mcall, err := multicall.NewMultiCall(multicall.GENERAL, client, nil)
	if err != nil {
		panic(err)
	}

	address := common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")

	targets := []*common.Address{
		&address,
		&address,
	}

	results, err := mcall.CodeLengths(targets, client, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(results)

	// Output: {true [3124 3124] <nil>}
}

func ExampleMultiCall_Balances() {
	rpc := "https://eth.llamarpc.com"
	client, err := ethclient.Dial(rpc)
	if err != nil {
		panic(err)
	}

	mcall, err := multicall.NewMultiCall(multicall.GENERAL, client, nil)
	if err != nil {
		panic(err)
	}

	address := common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")

	targets := []*common.Address{
		&address,
		&address,
	}

	results, err := mcall.Balances(targets, client, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(results)
}

func ExampleMultiCall_AddressesData() {
	rpc := "https://eth.llamarpc.com"
	client, err := ethclient.Dial(rpc)
	if err != nil {
		panic(err)
	}

	mcall, err := multicall.NewMultiCall(multicall.GENERAL, client, nil)
	if err != nil {
		panic(err)
	}

	address := common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")

	targets := []*common.Address{
		&address,
		&address,
	}

	results, err := mcall.AddressesData(targets, client, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(results)
}

func ExampleMultiCall_ChainData() {
	rpc := "https://eth.llamarpc.com"
	client, err := ethclient.Dial(rpc)
	if err != nil {
		panic(err)
	}

	mcall, err := multicall.NewMultiCall(multicall.GENERAL, client, nil)
	if err != nil {
		panic(err)
	}

	results, err := mcall.ChainData(client, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(results)
}
