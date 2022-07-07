// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BiddingWarBidRound is an auto generated low-level Go binding around an user-defined struct.
type BiddingWarBidRound struct {
	Winner          common.Address
	WinningBid      *big.Int
	AccumulatedBids *big.Int
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"activeBidRound\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CashOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"name\":\"NewBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"name\":\"NewBidRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gameIsRunning\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"getBidAtRound\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"winningBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedBids\",\"type\":\"uint256\"}],\"internalType\":\"structBiddingWar.BidRound\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommissions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentRoundNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastBid\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"winningBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedBids\",\"type\":\"uint256\"}],\"internalType\":\"structBiddingWar.BidRound\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"payWinner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"restart\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a0604052600560ff1660809060ff1660f81b8152503480156200002257600080fd5b506200004362000037620000ca60201b60201c565b620000d260201b60201c565b6000600560006101000a81548160ff021916908360ff1602179055506000600481905550610e1042620000779190620001d4565b60068190555060006001819055507f7f8c055dd603dd87c6a1f7020514ab0d46591306f5f005f113dbebaf05e8cbb6600454600654604051620000bc929190620001a7565b60405180910390a16200026a565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b620001a18162000231565b82525050565b6000604082019050620001be600083018562000196565b620001cd602083018462000196565b9392505050565b6000620001e18262000231565b9150620001ee8362000231565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156200022657620002256200023b565b5b828201905092915050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60805160f81c611781620002896000396000610b9701526117816000f3fe6080604052600436106100a75760003560e01c80638da5cb5b116100645780638da5cb5b14610151578063a5c0449f1461017c578063bca91d7f146101a5578063e92108b5146101d0578063f2fde38b146101fb578063f50f288914610224576100a7565b80631998aeef146100ac5780631ef3755d146100b65780633ccfd60b146100cd5780635dc5b365146100e45780636150ef6b1461010f578063715018a61461013a575b600080fd5b6100b4610261565b005b3480156100c257600080fd5b506100cb610578565b005b3480156100d957600080fd5b506100e2610631565b005b3480156100f057600080fd5b506100f96106ea565b6040516101069190611264565b60405180910390f35b34801561011b57600080fd5b50610124610781565b604051610131919061127f565b60405180910390f35b34801561014657600080fd5b5061014f61078b565b005b34801561015d57600080fd5b5061016661079f565b60405161017391906110ee565b60405180910390f35b34801561018857600080fd5b506101a3600480360381019061019e9190610ea2565b6107c8565b005b3480156101b157600080fd5b506101ba6109ff565b6040516101c79190611109565b60405180910390f35b3480156101dc57600080fd5b506101e5610a1b565b6040516101f2919061127f565b60405180910390f35b34801561020757600080fd5b50610222600480360381019061021d9190610e79565b610a25565b005b34801561023057600080fd5b5061024b60048036038101906102469190610ea2565b610aa9565b6040516102589190611264565b60405180910390f35b6006544211156102a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161029d906111a4565b60405180910390fd5b600034116102e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102e090611124565b60405180910390fd5b6000600560009054906101000a900460ff1660ff16141561041c576001600560008282829054906101000a900460ff1661032391906113b1565b92506101000a81548160ff021916908360ff1602179055506103453334610b91565b6040518060600160405280600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600354815260200160035481525060076000600454815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020155905050610417610c41565b610576565b600760006004548152602001908152602001600020600101543411610476576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161046d906111c4565b60405180910390fd5b6104803334610b91565b6040518060600160405280600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003548152602001600760006004548152602001908152602001600020600201546003546104f7919061135b565b81525060076000600454815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020155905050610575610c41565b5b565b610580610cce565b60065442116105c4576040517f447490fd0000000000000000000000000000000000000000000000000000000081526004016105bb906111e4565b60405180910390fd5b610e10426105d2919061135b565b6006819055506001600460008282546105eb919061135b565b925050819055507f7f8c055dd603dd87c6a1f7020514ab0d46591306f5f005f113dbebaf05e8cbb6600454600654604051610627929190611316565b60405180910390a1565b610639610cce565b60003373ffffffffffffffffffffffffffffffffffffffff16600154604051610661906110d9565b60006040518083038185875af1925050503d806000811461069e576040519150601f19603f3d011682016040523d82523d6000602084013e6106a3565b606091505b50509050806106e7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106de90611184565b60405180910390fd5b50565b6106f2610e18565b6007600060045481526020019081526020016000206040518060600160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201548152602001600282015481525050905090565b6000600454905090565b610793610cce565b61079d6000610d4c565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6107d0610cce565b6006544211610814576040517f447490fd00000000000000000000000000000000000000000000000000000000815260040161080b906111e4565b60405180910390fd5b6000600760008381526020019081526020016000206040518060600160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201548152602001600282015481525050905060008160400151116108e7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108de90611204565b60405180910390fd5b600060076000848152602001908152602001600020600201819055506000816000015173ffffffffffffffffffffffffffffffffffffffff168260400151604051610931906110d9565b60006040518083038185875af1925050503d806000811461096e576040519150601f19603f3d011682016040523d82523d6000602084013e610973565b606091505b50509050806109b7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109ae90611244565b60405180910390fd5b7f626be2766431e850fa0daafea58846cc985cbde6efd3887eea47fdc618d3800c83836000015184604001516040516109f29392919061129a565b60405180910390a1505050565b60006006544211610a135760019050610a18565b600090505b90565b6000600154905090565b610a2d610cce565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610a9d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a9490611144565b60405180910390fd5b610aa681610d4c565b50565b610ab1610e18565b600082118015610ac357506004548211155b610b02576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610af990611164565b60405180910390fd5b600760008381526020019081526020016000206040518060600160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820154815250509050919050565b600060647f0000000000000000000000000000000000000000000000000000000000000000610bc091906113e8565b60ff1682610bce9190611419565b90508082610bdc9190611473565b6003819055508060016000828254610bf4919061135b565b9250508190555082600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b61025860066000828254610c55919061135b565b925050819055507fcde18333b910a25ba5421e992ecee5c314e557346e94ab2ab424658bbb3707e6600454600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660035442600654610cb49190611473565b604051610cc494939291906112d1565b60405180910390a1565b610cd6610e10565b73ffffffffffffffffffffffffffffffffffffffff16610cf461079f565b73ffffffffffffffffffffffffffffffffffffffff1614610d4a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d4190611224565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b6040518060600160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081525090565b600081359050610e5e8161171d565b92915050565b600081359050610e7381611734565b92915050565b600060208284031215610e8b57600080fd5b6000610e9984828501610e4f565b91505092915050565b600060208284031215610eb457600080fd5b6000610ec284828501610e64565b91505092915050565b610ed4816114a7565b82525050565b610ee3816114a7565b82525050565b610ef2816114b9565b82525050565b6000610f05600b8361134a565b9150610f108261155a565b602082019050919050565b6000610f2860268361134a565b9150610f3382611583565b604082019050919050565b6000610f4b600d8361134a565b9150610f56826115d2565b602082019050919050565b6000610f6e601e8361134a565b9150610f79826115fb565b602082019050919050565b6000610f9160138361134a565b9150610f9c82611624565b602082019050919050565b6000610fb4601a8361134a565b9150610fbf8261164d565b602082019050919050565b6000610fd760188361134a565b9150610fe282611676565b602082019050919050565b6000610ffa60138361134a565b91506110058261169f565b602082019050919050565b600061101d60208361134a565b9150611028826116c8565b602082019050919050565b600061104060148361134a565b915061104b826116f1565b602082019050919050565b600061106360008361133f565b915061106e8261171a565b600082019050919050565b60608201600082015161108f6000850182610ecb565b5060208201516110a260208501826110bb565b5060408201516110b560408501826110bb565b50505050565b6110c4816114e5565b82525050565b6110d3816114e5565b82525050565b60006110e482611056565b9150819050919050565b60006020820190506111036000830184610eda565b92915050565b600060208201905061111e6000830184610ee9565b92915050565b6000602082019050818103600083015261113d81610ef8565b9050919050565b6000602082019050818103600083015261115d81610f1b565b9050919050565b6000602082019050818103600083015261117d81610f3e565b9050919050565b6000602082019050818103600083015261119d81610f61565b9050919050565b600060208201905081810360008301526111bd81610f84565b9050919050565b600060208201905081810360008301526111dd81610fa7565b9050919050565b600060208201905081810360008301526111fd81610fca565b9050919050565b6000602082019050818103600083015261121d81610fed565b9050919050565b6000602082019050818103600083015261123d81611010565b9050919050565b6000602082019050818103600083015261125d81611033565b9050919050565b60006060820190506112796000830184611079565b92915050565b600060208201905061129460008301846110ca565b92915050565b60006060820190506112af60008301866110ca565b6112bc6020830185610eda565b6112c960408301846110ca565b949350505050565b60006080820190506112e660008301876110ca565b6112f36020830186610eda565b61130060408301856110ca565b61130d60608301846110ca565b95945050505050565b600060408201905061132b60008301856110ca565b61133860208301846110ca565b9392505050565b600081905092915050565b600082825260208201905092915050565b6000611366826114e5565b9150611371836114e5565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156113a6576113a56114fc565b5b828201905092915050565b60006113bc826114ef565b91506113c7836114ef565b92508260ff038211156113dd576113dc6114fc565b5b828201905092915050565b60006113f3826114ef565b91506113fe836114ef565b92508261140e5761140d61152b565b5b828204905092915050565b6000611424826114e5565b915061142f836114e5565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611468576114676114fc565b5b828202905092915050565b600061147e826114e5565b9150611489836114e5565b92508282101561149c5761149b6114fc565b5b828203905092915050565b60006114b2826114c5565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f696e76616c696420626964000000000000000000000000000000000000000000600082015250565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f696e76616c696420726f756e6400000000000000000000000000000000000000600082015250565b7f756e61626c6520746f20776974686472617720636f6d6d697373696f6e730000600082015250565b7f62696420726f756e642068617320656e64656400000000000000000000000000600082015250565b7f626964206c6f776572207468616e2063757272656e7420626964000000000000600082015250565b7f612062696464696e6720776172206973206f6e676f696e670000000000000000600082015250565b7f77696e6e657220616c7265616479207061696400000000000000000000000000600082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f756e61626c6520746f207061792077696e6e6572000000000000000000000000600082015250565b50565b611726816114a7565b811461173157600080fd5b50565b61173d816114e5565b811461174857600080fd5b5056fea264697066735822122007e2e4956eca0ba86bdc02f72f78dbad481c282f9a3a0bf680593ff21683c65264736f6c63430008040033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// GameIsRunning is a free data retrieval call binding the contract method 0xbca91d7f.
//
// Solidity: function gameIsRunning() view returns(bool)
func (_Contract *ContractCaller) GameIsRunning(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "gameIsRunning")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GameIsRunning is a free data retrieval call binding the contract method 0xbca91d7f.
//
// Solidity: function gameIsRunning() view returns(bool)
func (_Contract *ContractSession) GameIsRunning() (bool, error) {
	return _Contract.Contract.GameIsRunning(&_Contract.CallOpts)
}

// GameIsRunning is a free data retrieval call binding the contract method 0xbca91d7f.
//
// Solidity: function gameIsRunning() view returns(bool)
func (_Contract *ContractCallerSession) GameIsRunning() (bool, error) {
	return _Contract.Contract.GameIsRunning(&_Contract.CallOpts)
}

// GetBidAtRound is a free data retrieval call binding the contract method 0xf50f2889.
//
// Solidity: function getBidAtRound(uint256 round) view returns((address,uint256,uint256))
func (_Contract *ContractCaller) GetBidAtRound(opts *bind.CallOpts, round *big.Int) (BiddingWarBidRound, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getBidAtRound", round)

	if err != nil {
		return *new(BiddingWarBidRound), err
	}

	out0 := *abi.ConvertType(out[0], new(BiddingWarBidRound)).(*BiddingWarBidRound)

	return out0, err

}

// GetBidAtRound is a free data retrieval call binding the contract method 0xf50f2889.
//
// Solidity: function getBidAtRound(uint256 round) view returns((address,uint256,uint256))
func (_Contract *ContractSession) GetBidAtRound(round *big.Int) (BiddingWarBidRound, error) {
	return _Contract.Contract.GetBidAtRound(&_Contract.CallOpts, round)
}

// GetBidAtRound is a free data retrieval call binding the contract method 0xf50f2889.
//
// Solidity: function getBidAtRound(uint256 round) view returns((address,uint256,uint256))
func (_Contract *ContractCallerSession) GetBidAtRound(round *big.Int) (BiddingWarBidRound, error) {
	return _Contract.Contract.GetBidAtRound(&_Contract.CallOpts, round)
}

// GetCommissions is a free data retrieval call binding the contract method 0xe92108b5.
//
// Solidity: function getCommissions() view returns(uint256)
func (_Contract *ContractCaller) GetCommissions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getCommissions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCommissions is a free data retrieval call binding the contract method 0xe92108b5.
//
// Solidity: function getCommissions() view returns(uint256)
func (_Contract *ContractSession) GetCommissions() (*big.Int, error) {
	return _Contract.Contract.GetCommissions(&_Contract.CallOpts)
}

// GetCommissions is a free data retrieval call binding the contract method 0xe92108b5.
//
// Solidity: function getCommissions() view returns(uint256)
func (_Contract *ContractCallerSession) GetCommissions() (*big.Int, error) {
	return _Contract.Contract.GetCommissions(&_Contract.CallOpts)
}

// GetCurrentRoundNumber is a free data retrieval call binding the contract method 0x6150ef6b.
//
// Solidity: function getCurrentRoundNumber() view returns(uint256)
func (_Contract *ContractCaller) GetCurrentRoundNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getCurrentRoundNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentRoundNumber is a free data retrieval call binding the contract method 0x6150ef6b.
//
// Solidity: function getCurrentRoundNumber() view returns(uint256)
func (_Contract *ContractSession) GetCurrentRoundNumber() (*big.Int, error) {
	return _Contract.Contract.GetCurrentRoundNumber(&_Contract.CallOpts)
}

// GetCurrentRoundNumber is a free data retrieval call binding the contract method 0x6150ef6b.
//
// Solidity: function getCurrentRoundNumber() view returns(uint256)
func (_Contract *ContractCallerSession) GetCurrentRoundNumber() (*big.Int, error) {
	return _Contract.Contract.GetCurrentRoundNumber(&_Contract.CallOpts)
}

// GetLastBid is a free data retrieval call binding the contract method 0x5dc5b365.
//
// Solidity: function getLastBid() view returns((address,uint256,uint256))
func (_Contract *ContractCaller) GetLastBid(opts *bind.CallOpts) (BiddingWarBidRound, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getLastBid")

	if err != nil {
		return *new(BiddingWarBidRound), err
	}

	out0 := *abi.ConvertType(out[0], new(BiddingWarBidRound)).(*BiddingWarBidRound)

	return out0, err

}

// GetLastBid is a free data retrieval call binding the contract method 0x5dc5b365.
//
// Solidity: function getLastBid() view returns((address,uint256,uint256))
func (_Contract *ContractSession) GetLastBid() (BiddingWarBidRound, error) {
	return _Contract.Contract.GetLastBid(&_Contract.CallOpts)
}

// GetLastBid is a free data retrieval call binding the contract method 0x5dc5b365.
//
// Solidity: function getLastBid() view returns((address,uint256,uint256))
func (_Contract *ContractCallerSession) GetLastBid() (BiddingWarBidRound, error) {
	return _Contract.Contract.GetLastBid(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Contract *ContractTransactor) Bid(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "bid")
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Contract *ContractSession) Bid() (*types.Transaction, error) {
	return _Contract.Contract.Bid(&_Contract.TransactOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Contract *ContractTransactorSession) Bid() (*types.Transaction, error) {
	return _Contract.Contract.Bid(&_Contract.TransactOpts)
}

// PayWinner is a paid mutator transaction binding the contract method 0xa5c0449f.
//
// Solidity: function payWinner(uint256 round) returns()
func (_Contract *ContractTransactor) PayWinner(opts *bind.TransactOpts, round *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "payWinner", round)
}

// PayWinner is a paid mutator transaction binding the contract method 0xa5c0449f.
//
// Solidity: function payWinner(uint256 round) returns()
func (_Contract *ContractSession) PayWinner(round *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PayWinner(&_Contract.TransactOpts, round)
}

// PayWinner is a paid mutator transaction binding the contract method 0xa5c0449f.
//
// Solidity: function payWinner(uint256 round) returns()
func (_Contract *ContractTransactorSession) PayWinner(round *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PayWinner(&_Contract.TransactOpts, round)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_Contract *ContractTransactor) Restart(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "restart")
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_Contract *ContractSession) Restart() (*types.Transaction, error) {
	return _Contract.Contract.Restart(&_Contract.TransactOpts)
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_Contract *ContractTransactorSession) Restart() (*types.Transaction, error) {
	return _Contract.Contract.Restart(&_Contract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Contract *ContractTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Contract *ContractSession) Withdraw() (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Contract *ContractTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts)
}

// ContractCashOutIterator is returned from FilterCashOut and is used to iterate over the raw logs and unpacked data for CashOut events raised by the Contract contract.
type ContractCashOutIterator struct {
	Event *ContractCashOut // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractCashOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCashOut)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractCashOut)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractCashOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCashOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCashOut represents a CashOut event raised by the Contract contract.
type ContractCashOut struct {
	Round  *big.Int
	Winner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCashOut is a free log retrieval operation binding the contract event 0x626be2766431e850fa0daafea58846cc985cbde6efd3887eea47fdc618d3800c.
//
// Solidity: event CashOut(uint256 round, address winner, uint256 amount)
func (_Contract *ContractFilterer) FilterCashOut(opts *bind.FilterOpts) (*ContractCashOutIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "CashOut")
	if err != nil {
		return nil, err
	}
	return &ContractCashOutIterator{contract: _Contract.contract, event: "CashOut", logs: logs, sub: sub}, nil
}

// WatchCashOut is a free log subscription operation binding the contract event 0x626be2766431e850fa0daafea58846cc985cbde6efd3887eea47fdc618d3800c.
//
// Solidity: event CashOut(uint256 round, address winner, uint256 amount)
func (_Contract *ContractFilterer) WatchCashOut(opts *bind.WatchOpts, sink chan<- *ContractCashOut) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "CashOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCashOut)
				if err := _Contract.contract.UnpackLog(event, "CashOut", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCashOut is a log parse operation binding the contract event 0x626be2766431e850fa0daafea58846cc985cbde6efd3887eea47fdc618d3800c.
//
// Solidity: event CashOut(uint256 round, address winner, uint256 amount)
func (_Contract *ContractFilterer) ParseCashOut(log types.Log) (*ContractCashOut, error) {
	event := new(ContractCashOut)
	if err := _Contract.contract.UnpackLog(event, "CashOut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNewBidIterator is returned from FilterNewBid and is used to iterate over the raw logs and unpacked data for NewBid events raised by the Contract contract.
type ContractNewBidIterator struct {
	Event *ContractNewBid // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractNewBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewBid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractNewBid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractNewBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewBid represents a NewBid event raised by the Contract contract.
type ContractNewBid struct {
	Round     *big.Int
	Bidder    common.Address
	BidAmount *big.Int
	TimeLeft  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewBid is a free log retrieval operation binding the contract event 0xcde18333b910a25ba5421e992ecee5c314e557346e94ab2ab424658bbb3707e6.
//
// Solidity: event NewBid(uint256 round, address bidder, uint256 bidAmount, uint256 timeLeft)
func (_Contract *ContractFilterer) FilterNewBid(opts *bind.FilterOpts) (*ContractNewBidIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewBid")
	if err != nil {
		return nil, err
	}
	return &ContractNewBidIterator{contract: _Contract.contract, event: "NewBid", logs: logs, sub: sub}, nil
}

// WatchNewBid is a free log subscription operation binding the contract event 0xcde18333b910a25ba5421e992ecee5c314e557346e94ab2ab424658bbb3707e6.
//
// Solidity: event NewBid(uint256 round, address bidder, uint256 bidAmount, uint256 timeLeft)
func (_Contract *ContractFilterer) WatchNewBid(opts *bind.WatchOpts, sink chan<- *ContractNewBid) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewBid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewBid)
				if err := _Contract.contract.UnpackLog(event, "NewBid", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewBid is a log parse operation binding the contract event 0xcde18333b910a25ba5421e992ecee5c314e557346e94ab2ab424658bbb3707e6.
//
// Solidity: event NewBid(uint256 round, address bidder, uint256 bidAmount, uint256 timeLeft)
func (_Contract *ContractFilterer) ParseNewBid(log types.Log) (*ContractNewBid, error) {
	event := new(ContractNewBid)
	if err := _Contract.contract.UnpackLog(event, "NewBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNewBidRoundIterator is returned from FilterNewBidRound and is used to iterate over the raw logs and unpacked data for NewBidRound events raised by the Contract contract.
type ContractNewBidRoundIterator struct {
	Event *ContractNewBidRound // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractNewBidRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewBidRound)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractNewBidRound)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractNewBidRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewBidRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewBidRound represents a NewBidRound event raised by the Contract contract.
type ContractNewBidRound struct {
	Round  *big.Int
	Expiry *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewBidRound is a free log retrieval operation binding the contract event 0x7f8c055dd603dd87c6a1f7020514ab0d46591306f5f005f113dbebaf05e8cbb6.
//
// Solidity: event NewBidRound(uint256 round, uint256 expiry)
func (_Contract *ContractFilterer) FilterNewBidRound(opts *bind.FilterOpts) (*ContractNewBidRoundIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewBidRound")
	if err != nil {
		return nil, err
	}
	return &ContractNewBidRoundIterator{contract: _Contract.contract, event: "NewBidRound", logs: logs, sub: sub}, nil
}

// WatchNewBidRound is a free log subscription operation binding the contract event 0x7f8c055dd603dd87c6a1f7020514ab0d46591306f5f005f113dbebaf05e8cbb6.
//
// Solidity: event NewBidRound(uint256 round, uint256 expiry)
func (_Contract *ContractFilterer) WatchNewBidRound(opts *bind.WatchOpts, sink chan<- *ContractNewBidRound) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewBidRound")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewBidRound)
				if err := _Contract.contract.UnpackLog(event, "NewBidRound", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewBidRound is a log parse operation binding the contract event 0x7f8c055dd603dd87c6a1f7020514ab0d46591306f5f005f113dbebaf05e8cbb6.
//
// Solidity: event NewBidRound(uint256 round, uint256 expiry)
func (_Contract *ContractFilterer) ParseNewBidRound(log types.Log) (*ContractNewBidRound, error) {
	event := new(ContractNewBidRound)
	if err := _Contract.contract.UnpackLog(event, "NewBidRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
