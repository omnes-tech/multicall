package multicall_test

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/omnes-tech/multicall"
)

func ExampleNewClient() {
	rpc := "https://eth.llamarpc.com"
	client, err := multicall.NewClient(multicall.GENERAL, rpc, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("RPC: %s MultiCallType: %d WriteAddress: %s ReadAddress: %s",
		client.RPC, client.MultiCallType, client.WriteAddress, client.ReadAddress)

	// Output: RPC: https://eth.llamarpc.com MultiCallType: 0 WriteAddress: 0xcA11bde05977b3631167028862bE2a173976CA11 ReadAddress: <nil>
}

func ExampleMultiCallClient_SimulateCall() {
	rpc := "https://eth.llamarpc.com"
	client, err := multicall.NewClient(multicall.GENERAL, rpc, nil)
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

	calls := multicall.NewCalls(targets, funcSigs, nil, nil, values)

	results, err := client.SimulateCall(calls)

	fmt.Println(results)

	// Output: {true [[true  33921] [true  9521]] <nil>}
}

func ExampleMultiCallClient_AggregateStatic() {
	rpc := "https://eth.llamarpc.com"
	client, err := multicall.NewClient(multicall.GENERAL, rpc, nil)
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

	calls := multicall.NewCalls(targets, funcSigs, argss, returnTypes, nil)

	results, err := client.AggregateStatic(calls)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(results)

	// Output: {true [[1085420955917931147422] [1085420955917931147422]] <nil>}
}

func ExampleMultiCallClient_TryAggregateStatic() {
	rpc := "https://eth.llamarpc.com"
	client, err := multicall.NewClient(multicall.GENERAL, rpc, nil)
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

	calls := multicall.NewCalls(targets, funcSigs, argss, returnTypes, nil)

	results, err := client.TryAggregateStatic(calls, true)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(results)

	// Output: {true [[true [1085420955917931147422]] [true [1085420955917931147422]]] <nil>}
}

func ExampleMultiCallClient_TryAggregateStatic3() {
	rpc := "https://eth.llamarpc.com"
	client, err := multicall.NewClient(multicall.GENERAL, rpc, nil)
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

	calls := multicall.NewCallsWithFailure(targets, funcSigs, argss, returnTypes, nil, requireSuccess)

	results, err := client.TryAggregateStatic3(calls)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(results)

	// Output: {true [[true [1085420955917931147422]] [true [1085420955917931147422]]] <nil>}
}

func ExampleMultiCallClient_CodeLengths() {
	rpc := "https://eth.llamarpc.com"
	client, err := multicall.NewClient(multicall.GENERAL, rpc, nil)
	if err != nil {
		panic(err)
	}

	address := common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")

	targets := []*common.Address{
		&address,
		&address,
	}

	results, err := client.CodeLengths(targets)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(results)

	// Output: {true [3124 3124] <nil>}
}

func ExampleMultiCallClient_Balances() {
	rpc := "https://eth.llamarpc.com"
	client, err := multicall.NewClient(multicall.GENERAL, rpc, nil)
	if err != nil {
		panic(err)
	}

	address := common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")

	targets := []*common.Address{
		&address,
		&address,
	}

	results, err := client.Balances(targets)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(results)

	// Output: {true [2990859034749149558049965 2990859034749149558049965] <nil>}
}

func ExampleMultiCallClient_AddressesData() {
	rpc := "https://eth.llamarpc.com"
	client, err := multicall.NewClient(multicall.GENERAL, rpc, nil)
	if err != nil {
		panic(err)
	}

	address := common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")

	targets := []*common.Address{
		&address,
		&address,
	}

	results, err := client.AddressesData(targets)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(results)

	// Output: {true [[2990903258548144038500852 3124] [2990903258548144038500852 3124]] <nil>}
}
