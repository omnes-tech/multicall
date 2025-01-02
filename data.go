package multicall

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

var GENERAL_MULTICALL_ADDRESS = common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11")

const GENERAL_MULTICALL_ABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"struct Multicall3.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"returnData\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"struct Multicall3.Call3[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate3\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"struct Multicall3.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"struct Multicall3.Call3Value[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate3Value\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"struct Multicall3.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"struct Multicall3.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"blockAndAggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"struct Multicall3.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBasefee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"basefee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chainid\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockCoinbase\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockDifficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gaslimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getEthBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"requireSuccess\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"struct Multicall3.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"tryAggregate\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"struct Multicall3.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"requireSuccess\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"struct Multicall3.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"tryBlockAndAggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"struct Multicall3.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

var OMNES_MULTICALL_ADDRESS = common.Address{}

const OMNES_MULTICALL_ABI = "[{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"aggregateCalls\",\"inputs\":[{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"struct Call[]\",\"components\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"returnDatas\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"aggregateStatic\",\"inputs\":[{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"struct StaticCall[]\",\"components\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"returnData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAddressesData\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"balances\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"codeLengths\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBalances\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"balances\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCodeLengths\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"lengths\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"simulateCalls\",\"inputs\":[{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"struct Call[]\",\"components\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"tryAggregateCalls\",\"inputs\":[{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"struct MultiCall.CallWithFailure[]\",\"components\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requireSuccess\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[{\"name\":\"results\",\"type\":\"tuple[]\",\"internalType\":\"struct MultiCall.Result[]\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"tryAggregateCalls\",\"inputs\":[{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"struct Call[]\",\"components\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"requireSuccess\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"results\",\"type\":\"tuple[]\",\"internalType\":\"struct MultiCall.Result[]\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"tryAggregateStatic\",\"inputs\":[{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"struct StaticCall[]\",\"components\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"requireSuccess\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"returnData\",\"type\":\"tuple[]\",\"internalType\":\"struct MultiCall.Result[]\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tryAggregateStatic\",\"inputs\":[{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"struct StaticCallWithFailure[]\",\"components\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"requireSuccess\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[{\"name\":\"returnData\",\"type\":\"tuple[]\",\"internalType\":\"struct MultiCall.Result[]\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"MultiCall__CallFailed\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MultiCall__SendingValueNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MultiCall__Simulation\",\"inputs\":[{\"name\":\"results\",\"type\":\"tuple[]\",\"internalType\":\"struct MultiCall.SimulatedResult[]\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"type\":\"error\",\"name\":\"MultiCall__StaticCallFailed\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]"

const DEPLOYLESS_MULTICALL_ABI = "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"error\",\"name\":\"MultiCallCodec__InvalidStaticCallType\",\"inputs\":[{\"name\":\"type_\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"MultiCall__SendingValueNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MultiCall__Simulation\",\"inputs\":[{\"name\":\"results\",\"type\":\"tuple[]\",\"internalType\":\"struct DeploylessMultiCall.SimulatedResult[]\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"type\":\"error\",\"name\":\"MultiCall__StaticCallFailed\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]"
const DEPLOYLESS_MULTICALL_BYTECODE = "0x608060405260405161206b38038061206b83398181016040528101906100259190611224565b5f5f5f5f5f5f61003a876102fb60201b60201c565b95509550955095509550955060605f600681111561005b5761005a61126b565b5b87600681111561006e5761006d61126b565b5b03610087576100828361051f60201b60201c565b6102f3565b6001600681111561009b5761009a61126b565b5b8760068111156100ae576100ad61126b565b5b036100ed575f6100c3876106e060201b60201c565b9050806040516020016100d691906113a5565b6040516020818303038152906040529150506102f2565b600260068111156101015761010061126b565b5b8760068111156101145761011361126b565b5b03610154575f61012a878761088960201b60201c565b90508060405160200161013d91906114d4565b6040516020818303038152906040529150506102f1565b600360068111156101685761016761126b565b5b87600681111561017b5761017a61126b565b5b036101ba575f61019085610a5260201b60201c565b9050806040516020016101a391906114d4565b6040516020818303038152906040529150506102f0565b600460068111156101ce576101cd61126b565b5b8760068111156101e1576101e061126b565b5b03610220575f6101f683610c4460201b60201c565b90508060405160200161020991906115b4565b6040516020818303038152906040529150506102ef565b600560068111156102345761023361126b565b5b8760068111156102475761024661126b565b5b03610286575f61025c83610d1b60201b60201c565b90508060405160200161026f91906115b4565b6040516020818303038152906040529150506102ee565b6006808111156102995761029861126b565b5b8760068111156102ac576102ab61126b565b5b036102ed575f5f6102c284610df260201b60201c565b9150915081816040516020016102d99291906115d4565b604051602081830303815290604052925050505b5b5b5b5b5b5b805160208201f35b5f60605f60608060605f875f8151811061031857610317611609565b5b602001015160f81c60f81b60f81c90505f610347896001808c5161033c9190611663565b610f7f60201b60201c565b90506001600681111561035d5761035c61126b565b5b60ff168260ff1603610384578080602001905181019061037d9190611843565b96506104fc565b600260068111156103985761039761126b565b5b60ff168260ff16036103c557808060200190518101906103b891906118b4565b80975081985050506104fb565b600360068111156103d9576103d861126b565b5b60ff168260ff160361040057808060200190518101906103f99190611a69565b94506104fa565b5f60068111156104135761041261126b565b5b60ff168260ff160361043a57808060200190518101906104339190611c35565b93506104f9565b6004600681111561044e5761044d61126b565b5b60ff168260ff1614806104795750600560068111156104705761046f61126b565b5b60ff168260ff16145b8061049b57506006808111156104925761049161126b565b5b60ff168260ff16145b156104bb57808060200190518101906104b49190611d3c565b92506104f8565b816040517f654f0df70000000000000000000000000000000000000000000000000000000081526004016104ef9190611d9e565b60405180910390fd5b5b5b5b5b8160ff1660068111156105125761051161126b565b5b9750505091939550919395565b5f815167ffffffffffffffff81111561053b5761053a611100565b5b60405190808252806020026020018201604052801561057457816020015b61056161109b565b8152602001906001900390816105595790505b5090505f5f90505b82518110156106a2575f5a90505f5f85848151811061059e5761059d611609565b5b60200260200101515f015173ffffffffffffffffffffffffffffffffffffffff168685815181106105d2576105d1611609565b5b6020026020010151604001518786815181106105f1576105f0611609565b5b60200260200101516020015160405161060a9190611df1565b5f6040518083038185875af1925050503d805f8114610644576040519150601f19603f3d011682016040523d82523d5f602084013e610649565b606091505b5091509150604051806060016040528083151581526020018281526020015a856106739190611663565b81525085858151811061068957610688611609565b5b602002602001018190525083600101935050505061057c565b50806040517fc6da632c0000000000000000000000000000000000000000000000000000000081526004016106d79190611f0f565b60405180910390fd5b6060815167ffffffffffffffff8111156106fd576106fc611100565b5b60405190808252806020026020018201604052801561073057816020015b606081526020019060019003908161071b5790505b5090505f825190505b5f811115610883575f836001836107509190611663565b8151811061076157610760611609565b5b60200260200101515f015173ffffffffffffffffffffffffffffffffffffffff16846001846107909190611663565b815181106107a1576107a0611609565b5b6020026020010151602001516040516107ba9190611df1565b5f60405180830381855afa9150503d805f81146107f2576040519150601f19603f3d011682016040523d82523d5f602084013e6107f7565b606091505b50846001856108069190611663565b8151811061081757610816611609565b5b602002602001018190528192505050806001836108349190611663565b90610875576040517f9eadd6c500000000000000000000000000000000000000000000000000000000815260040161086c9190611f3e565b60405180910390fd5b508160019003915050610739565b50919050565b6060825167ffffffffffffffff8111156108a6576108a5611100565b5b6040519080825280602002602001820160405280156108df57816020015b6108cc6110bc565b8152602001906001900390816108c45790505b5090505f835190505b5f811115610a4b575f826001836108ff9190611663565b815181106109105761090f611609565b5b60200260200101519050846001836109289190611663565b8151811061093957610938611609565b5b60200260200101515f015173ffffffffffffffffffffffffffffffffffffffff16856001846109689190611663565b8151811061097957610978611609565b5b6020026020010151602001516040516109929190611df1565b5f60405180830381855afa9150503d805f81146109ca576040519150601f19603f3d011682016040523d82523d5f602084013e6109cf565b606091505b50825f0183602001829052821515151581525050508315610a3e57805f01516001836109fb9190611663565b90610a3c576040517f9eadd6c5000000000000000000000000000000000000000000000000000000008152600401610a339190611f3e565b60405180910390fd5b505b81600190039150506108e8565b5092915050565b6060815167ffffffffffffffff811115610a6f57610a6e611100565b5b604051908082528060200260200182016040528015610aa857816020015b610a956110bc565b815260200190600190039081610a8d5790505b5090505f825190505b5f811115610c3e575f82600183610ac89190611663565b81518110610ad957610ad8611609565b5b6020026020010151905083600183610af19190611663565b81518110610b0257610b01611609565b5b60200260200101515f015173ffffffffffffffffffffffffffffffffffffffff1684600184610b319190611663565b81518110610b4257610b41611609565b5b602002602001015160200151604051610b5b9190611df1565b5f60405180830381855afa9150503d805f8114610b93576040519150601f19603f3d011682016040523d82523d5f602084013e610b98565b606091505b50825f01836020018290528215151515815250505083600183610bbb9190611663565b81518110610bcc57610bcb611609565b5b60200260200101516040015115610c3157805f0151600183610bee9190611663565b90610c2f576040517f9eadd6c5000000000000000000000000000000000000000000000000000000008152600401610c269190611f3e565b60405180910390fd5b505b8160019003915050610ab1565b50919050565b6060815167ffffffffffffffff811115610c6157610c60611100565b5b604051908082528060200260200182016040528015610c8f5781602001602082028036833780820191505090505b5090505f825190505b5f811115610d155782600182610cae9190611663565b81518110610cbf57610cbe611609565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff163b82600183610cec9190611663565b81518110610cfd57610cfc611609565b5b60200260200101818152505080600190039050610c98565b50919050565b6060815167ffffffffffffffff811115610d3857610d37611100565b5b604051908082528060200260200182016040528015610d665781602001602082028036833780820191505090505b5090505f825190505b5f811115610dec5782600182610d859190611663565b81518110610d9657610d95611609565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff163182600183610dc39190611663565b81518110610dd457610dd3611609565b5b60200260200101818152505080600190039050610d6f565b50919050565b606080825167ffffffffffffffff811115610e1057610e0f611100565b5b604051908082528060200260200182016040528015610e3e5781602001602082028036833780820191505090505b509150825167ffffffffffffffff811115610e5c57610e5b611100565b5b604051908082528060200260200182016040528015610e8a5781602001602082028036833780820191505090505b5090505f835190505b5f811115610f795783600182610ea99190611663565b81518110610eba57610eb9611609565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff163183600183610ee79190611663565b81518110610ef857610ef7611609565b5b60200260200101818152505083600182610f129190611663565b81518110610f2357610f22611609565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff163b82600183610f509190611663565b81518110610f6157610f60611609565b5b60200260200101818152505080600190039050610e93565b50915091565b606081601f83610f8f9190611f57565b1015610fd0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fc790611fe4565b60405180910390fd5b8183610fdc9190611f57565b8451101561101f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110169061204c565b60405180910390fd5b606082155f811461103e5760405191505f82526020820160405261108f565b6040519150601f8416801560200281840101858101878315602002848b0101015b8183101561107c578051835260208301925060208101905061105f565b50868552601f19601f8301166040525050505b50809150509392505050565b60405180606001604052805f15158152602001606081526020015f81525090565b60405180604001604052805f15158152602001606081525090565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611136826110f0565b810181811067ffffffffffffffff8211171561115557611154611100565b5b80604052505050565b5f6111676110d7565b9050611173828261112d565b919050565b5f67ffffffffffffffff82111561119257611191611100565b5b61119b826110f0565b9050602081019050919050565b8281835e5f83830152505050565b5f6111c86111c384611178565b61115e565b9050828152602081018484840111156111e4576111e36110ec565b5b6111ef8482856111a8565b509392505050565b5f82601f83011261120b5761120a6110e8565b5b815161121b8482602086016111b6565b91505092915050565b5f60208284031215611239576112386110e0565b5b5f82015167ffffffffffffffff811115611256576112556110e4565b5b611262848285016111f7565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f6112e5826112c1565b6112ef81856112cb565b93506112ff8185602086016111a8565b611308816110f0565b840191505092915050565b5f61131e83836112db565b905092915050565b5f602082019050919050565b5f61133c82611298565b61134681856112a2565b935083602082028501611358856112b2565b805f5b8581101561139357848403895281516113748582611313565b945061137f83611326565b925060208a0199505060018101905061135b565b50829750879550505050505092915050565b5f6020820190508181035f8301526113bd8184611332565b905092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f8115159050919050565b611402816113ee565b82525050565b5f604083015f83015161141d5f8601826113f9565b506020830151848203602086015261143582826112db565b9150508091505092915050565b5f61144d8383611408565b905092915050565b5f602082019050919050565b5f61146b826113c5565b61147581856113cf565b935083602082028501611487856113df565b805f5b858110156114c257848403895281516114a38582611442565b94506114ae83611455565b925060208a0199505060018101905061148a565b50829750879550505050505092915050565b5f6020820190508181035f8301526114ec8184611461565b905092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f819050919050565b61152f8161151d565b82525050565b5f6115408383611526565b60208301905092915050565b5f602082019050919050565b5f611562826114f4565b61156c81856114fe565b93506115778361150e565b805f5b838110156115a757815161158e8882611535565b97506115998361154c565b92505060018101905061157a565b5085935050505092915050565b5f6020820190508181035f8301526115cc8184611558565b905092915050565b5f6040820190508181035f8301526115ec8185611558565b905081810360208301526116008184611558565b90509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61166d8261151d565b91506116788361151d565b92508282039050818111156116905761168f611636565b5b92915050565b5f67ffffffffffffffff8211156116b0576116af611100565b5b602082029050602081019050919050565b5f5ffd5b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6116f6826116cd565b9050919050565b611706816116ec565b8114611710575f5ffd5b50565b5f81519050611721816116fd565b92915050565b5f6040828403121561173c5761173b6116c5565b5b611746604061115e565b90505f61175584828501611713565b5f83015250602082015167ffffffffffffffff811115611778576117776116c9565b5b611784848285016111f7565b60208301525092915050565b5f6117a261179d84611696565b61115e565b905080838252602082019050602084028301858111156117c5576117c46116c1565b5b835b8181101561180c57805167ffffffffffffffff8111156117ea576117e96110e8565b5b8086016117f78982611727565b855260208501945050506020810190506117c7565b5050509392505050565b5f82601f83011261182a576118296110e8565b5b815161183a848260208601611790565b91505092915050565b5f60208284031215611858576118576110e0565b5b5f82015167ffffffffffffffff811115611875576118746110e4565b5b61188184828501611816565b91505092915050565b611893816113ee565b811461189d575f5ffd5b50565b5f815190506118ae8161188a565b92915050565b5f5f604083850312156118ca576118c96110e0565b5b5f83015167ffffffffffffffff8111156118e7576118e66110e4565b5b6118f385828601611816565b9250506020611904858286016118a0565b9150509250929050565b5f67ffffffffffffffff82111561192857611927611100565b5b602082029050602081019050919050565b5f6060828403121561194e5761194d6116c5565b5b611958606061115e565b90505f61196784828501611713565b5f83015250602082015167ffffffffffffffff81111561198a576119896116c9565b5b611996848285016111f7565b60208301525060406119aa848285016118a0565b60408301525092915050565b5f6119c86119c38461190e565b61115e565b905080838252602082019050602084028301858111156119eb576119ea6116c1565b5b835b81811015611a3257805167ffffffffffffffff811115611a1057611a0f6110e8565b5b808601611a1d8982611939565b855260208501945050506020810190506119ed565b5050509392505050565b5f82601f830112611a5057611a4f6110e8565b5b8151611a608482602086016119b6565b91505092915050565b5f60208284031215611a7e57611a7d6110e0565b5b5f82015167ffffffffffffffff811115611a9b57611a9a6110e4565b5b611aa784828501611a3c565b91505092915050565b5f67ffffffffffffffff821115611aca57611ac9611100565b5b602082029050602081019050919050565b611ae48161151d565b8114611aee575f5ffd5b50565b5f81519050611aff81611adb565b92915050565b5f60608284031215611b1a57611b196116c5565b5b611b24606061115e565b90505f611b3384828501611713565b5f83015250602082015167ffffffffffffffff811115611b5657611b556116c9565b5b611b62848285016111f7565b6020830152506040611b7684828501611af1565b60408301525092915050565b5f611b94611b8f84611ab0565b61115e565b90508083825260208201905060208402830185811115611bb757611bb66116c1565b5b835b81811015611bfe57805167ffffffffffffffff811115611bdc57611bdb6110e8565b5b808601611be98982611b05565b85526020850194505050602081019050611bb9565b5050509392505050565b5f82601f830112611c1c57611c1b6110e8565b5b8151611c2c848260208601611b82565b91505092915050565b5f60208284031215611c4a57611c496110e0565b5b5f82015167ffffffffffffffff811115611c6757611c666110e4565b5b611c7384828501611c08565b91505092915050565b5f67ffffffffffffffff821115611c9657611c95611100565b5b602082029050602081019050919050565b5f611cb9611cb484611c7c565b61115e565b90508083825260208201905060208402830185811115611cdc57611cdb6116c1565b5b835b81811015611d055780611cf18882611713565b845260208401935050602081019050611cde565b5050509392505050565b5f82601f830112611d2357611d226110e8565b5b8151611d33848260208601611ca7565b91505092915050565b5f60208284031215611d5157611d506110e0565b5b5f82015167ffffffffffffffff811115611d6e57611d6d6110e4565b5b611d7a84828501611d0f565b91505092915050565b5f60ff82169050919050565b611d9881611d83565b82525050565b5f602082019050611db15f830184611d8f565b92915050565b5f81905092915050565b5f611dcb826112c1565b611dd58185611db7565b9350611de58185602086016111a8565b80840191505092915050565b5f611dfc8284611dc1565b915081905092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f606083015f830151611e455f8601826113f9565b5060208301518482036020860152611e5d82826112db565b9150506040830151611e726040860182611526565b508091505092915050565b5f611e888383611e30565b905092915050565b5f602082019050919050565b5f611ea682611e07565b611eb08185611e11565b935083602082028501611ec285611e21565b805f5b85811015611efd5784840389528151611ede8582611e7d565b9450611ee983611e90565b925060208a01995050600181019050611ec5565b50829750879550505050505092915050565b5f6020820190508181035f830152611f278184611e9c565b905092915050565b611f388161151d565b82525050565b5f602082019050611f515f830184611f2f565b92915050565b5f611f618261151d565b9150611f6c8361151d565b9250828201905080821115611f8457611f83611636565b5b92915050565b5f82825260208201905092915050565b7f736c6963655f6f766572666c6f770000000000000000000000000000000000005f82015250565b5f611fce600e83611f8a565b9150611fd982611f9a565b602082019050919050565b5f6020820190508181035f830152611ffb81611fc2565b9050919050565b7f736c6963655f6f75744f66426f756e64730000000000000000000000000000005f82015250565b5f612036601183611f8a565b915061204182612002565b602082019050919050565b5f6020820190508181035f8301526120638161202a565b905091905056fe"

var ZERO_ADDRESS = common.Address{}

const MINING_WAIT_DURATION = 600 * time.Second