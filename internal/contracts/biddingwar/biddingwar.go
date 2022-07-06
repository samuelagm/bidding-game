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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"activeBidRound\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CashOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"name\":\"NewBidRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"name\":\"PotentialWinner\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"getBidAtRound\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"winningBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedBids\",\"type\":\"uint256\"}],\"internalType\":\"structBiddingWar.BidRound\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommissions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentRoundNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastBid\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"winningBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedBids\",\"type\":\"uint256\"}],\"internalType\":\"structBiddingWar.BidRound\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"payWinner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"restart\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a0604052600560ff1660809060ff1660f81b8152503480156200002257600080fd5b506200004362000037620000c260201b60201c565b620000ca60201b60201c565b6000600560006101000a81548160ff021916908360ff1602179055506000600481905550610e1042620000779190620001cc565b6006819055507f7f8c055dd603dd87c6a1f7020514ab0d46591306f5f005f113dbebaf05e8cbb6600454600654604051620000b49291906200019f565b60405180910390a162000262565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b620001998162000229565b82525050565b6000604082019050620001b660008301856200018e565b620001c560208301846200018e565b9392505050565b6000620001d98262000229565b9150620001e68362000229565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156200021e576200021d62000233565b5b828201905092915050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60805160f81c6116f8620002816000396000610b4401526116f86000f3fe60806040526004361061009c5760003560e01c8063715018a611610064578063715018a61461012f5780638da5cb5b14610146578063a5c0449f14610171578063e92108b51461019a578063f2fde38b146101c5578063f50f2889146101ee5761009c565b80631998aeef146100a15780631ef3755d146100ab5780633ccfd60b146100c25780635dc5b365146100d95780636150ef6b14610104575b600080fd5b6100a961022b565b005b3480156100b757600080fd5b506100c0610541565b005b3480156100ce57600080fd5b506100d76105fb565b005b3480156100e557600080fd5b506100ee6106b2565b6040516100fb91906111e7565b60405180910390f35b34801561011057600080fd5b50610119610749565b6040516101269190611202565b60405180910390f35b34801561013b57600080fd5b50610144610753565b005b34801561015257600080fd5b5061015b610767565b604051610168919061108c565b60405180910390f35b34801561017d57600080fd5b5061019860048036038101906101939190610e4f565b610790565b005b3480156101a657600080fd5b506101af6109c8565b6040516101bc9190611202565b60405180910390f35b3480156101d157600080fd5b506101ec60048036038101906101e79190610e26565b6109d2565b005b3480156101fa57600080fd5b5061021560048036038101906102109190610e4f565b610a56565b60405161022291906111e7565b60405180910390f35b600654421061026f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161026690611147565b60405180910390fd5b600034116102b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102a9906110a7565b60405180910390fd5b6000600560009054906101000a900460ff1660ff1614156103e5576001600560008282829054906101000a900460ff166102ec9190611334565b92506101000a81548160ff021916908360ff16021790555061030e3334610b3e565b6040518060600160405280600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600354815260200160035481525060076000600454815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155604082015181600201559050506103e0610bee565b61053f565b60076000600454815260200190815260200160002060010154341161043f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161043690611167565b60405180910390fd5b6104493334610b3e565b6040518060600160405280600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003548152602001600760006004548152602001908152602001600020600201546003546104c091906112de565b81525060076000600454815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101556040820151816002015590505061053e610bee565b5b565b610549610c7b565b60065442101561058e576040517f447490fd000000000000000000000000000000000000000000000000000000008152600401610585906110c7565b60405180910390fd5b610e104261059c91906112de565b6006819055506001600460008282546105b591906112de565b925050819055507f7f8c055dd603dd87c6a1f7020514ab0d46591306f5f005f113dbebaf05e8cbb66004546006546040516105f1929190611299565b60405180910390a1565b610603610c7b565b60003373ffffffffffffffffffffffffffffffffffffffff164760405161062990611077565b60006040518083038185875af1925050503d8060008114610666576040519150601f19603f3d011682016040523d82523d6000602084013e61066b565b606091505b50509050806106af576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106a690611127565b60405180910390fd5b50565b6106ba610dc5565b6007600060045481526020019081526020016000206040518060600160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201548152602001600282015481525050905090565b6000600454905090565b61075b610c7b565b6107656000610cf9565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610798610c7b565b6006544210156107dd576040517f447490fd0000000000000000000000000000000000000000000000000000000081526004016107d4906110c7565b60405180910390fd5b6000600760008381526020019081526020016000206040518060600160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201548152602001600282015481525050905060008160400151116108b0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108a790611187565b60405180910390fd5b600060076000848152602001908152602001600020600201819055506000816000015173ffffffffffffffffffffffffffffffffffffffff1682604001516040516108fa90611077565b60006040518083038185875af1925050503d8060008114610937576040519150601f19603f3d011682016040523d82523d6000602084013e61093c565b606091505b5050905080610980576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610977906111c7565b60405180910390fd5b7f626be2766431e850fa0daafea58846cc985cbde6efd3887eea47fdc618d3800c83836000015184604001516040516109bb9392919061121d565b60405180910390a1505050565b6000600154905090565b6109da610c7b565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610a4a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a41906110e7565b60405180910390fd5b610a5381610cf9565b50565b610a5e610dc5565b600082118015610a7057506004548211155b610aaf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aa690611107565b60405180910390fd5b600760008381526020019081526020016000206040518060600160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820154815250509050919050565b600060647f0000000000000000000000000000000000000000000000000000000000000000610b6d919061136b565b60ff1682610b7b919061139c565b90508082610b8991906113f6565b6003819055508060016000828254610ba191906112de565b9250508190555082600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b61025860066000828254610c0291906112de565b925050819055507f7622c6e07e6dde4814d2d4642049258551877b7ac81d9e13a69e55ba6df91e9b600454600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660035442600654610c6191906113f6565b604051610c719493929190611254565b60405180910390a1565b610c83610dbd565b73ffffffffffffffffffffffffffffffffffffffff16610ca1610767565b73ffffffffffffffffffffffffffffffffffffffff1614610cf7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cee906111a7565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b6040518060600160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081525090565b600081359050610e0b81611694565b92915050565b600081359050610e20816116ab565b92915050565b600060208284031215610e3857600080fd5b6000610e4684828501610dfc565b91505092915050565b600060208284031215610e6157600080fd5b6000610e6f84828501610e11565b91505092915050565b610e818161142a565b82525050565b610e908161142a565b82525050565b6000610ea3600b836112cd565b9150610eae826114d1565b602082019050919050565b6000610ec66017836112cd565b9150610ed1826114fa565b602082019050919050565b6000610ee96026836112cd565b9150610ef482611523565b604082019050919050565b6000610f0c600d836112cd565b9150610f1782611572565b602082019050919050565b6000610f2f601e836112cd565b9150610f3a8261159b565b602082019050919050565b6000610f526013836112cd565b9150610f5d826115c4565b602082019050919050565b6000610f75601a836112cd565b9150610f80826115ed565b602082019050919050565b6000610f986013836112cd565b9150610fa382611616565b602082019050919050565b6000610fbb6020836112cd565b9150610fc68261163f565b602082019050919050565b6000610fde6014836112cd565b9150610fe982611668565b602082019050919050565b60006110016000836112c2565b915061100c82611691565b600082019050919050565b60608201600082015161102d6000850182610e78565b5060208201516110406020850182611059565b5060408201516110536040850182611059565b50505050565b6110628161145c565b82525050565b6110718161145c565b82525050565b600061108282610ff4565b9150819050919050565b60006020820190506110a16000830184610e87565b92915050565b600060208201905081810360008301526110c081610e96565b9050919050565b600060208201905081810360008301526110e081610eb9565b9050919050565b6000602082019050818103600083015261110081610edc565b9050919050565b6000602082019050818103600083015261112081610eff565b9050919050565b6000602082019050818103600083015261114081610f22565b9050919050565b6000602082019050818103600083015261116081610f45565b9050919050565b6000602082019050818103600083015261118081610f68565b9050919050565b600060208201905081810360008301526111a081610f8b565b9050919050565b600060208201905081810360008301526111c081610fae565b9050919050565b600060208201905081810360008301526111e081610fd1565b9050919050565b60006060820190506111fc6000830184611017565b92915050565b60006020820190506112176000830184611068565b92915050565b60006060820190506112326000830186611068565b61123f6020830185610e87565b61124c6040830184611068565b949350505050565b60006080820190506112696000830187611068565b6112766020830186610e87565b6112836040830185611068565b6112906060830184611068565b95945050505050565b60006040820190506112ae6000830185611068565b6112bb6020830184611068565b9392505050565b600081905092915050565b600082825260208201905092915050565b60006112e98261145c565b91506112f48361145c565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561132957611328611473565b5b828201905092915050565b600061133f82611466565b915061134a83611466565b92508260ff038211156113605761135f611473565b5b828201905092915050565b600061137682611466565b915061138183611466565b925082611391576113906114a2565b5b828204905092915050565b60006113a78261145c565b91506113b28361145c565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156113eb576113ea611473565b5b828202905092915050565b60006114018261145c565b915061140c8361145c565b92508282101561141f5761141e611473565b5b828203905092915050565b60006114358261143c565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f696e76616c696420626964000000000000000000000000000000000000000000600082015250565b7f612062696464696e6720776172206973206f6e676f6967000000000000000000600082015250565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f696e76616c696420726f756e6400000000000000000000000000000000000000600082015250565b7f756e61626c6520746f20776974686472617720636f6d6d697373696f6e730000600082015250565b7f62696420726f756e642068617320656e64656400000000000000000000000000600082015250565b7f626964206c6f776572207468616e2063757272656e7420626964000000000000600082015250565b7f77696e6e657220616c7265616479207061696400000000000000000000000000600082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f756e61626c6520746f207061792077696e6e6572000000000000000000000000600082015250565b50565b61169d8161142a565b81146116a857600080fd5b50565b6116b48161145c565b81146116bf57600080fd5b5056fea264697066735822122041ae575e0b3ef34071f4e8127af980185bee9e0415377bf24a5808f4e2f08cac64736f6c63430008040033",
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

// ContractPotentialWinnerIterator is returned from FilterPotentialWinner and is used to iterate over the raw logs and unpacked data for PotentialWinner events raised by the Contract contract.
type ContractPotentialWinnerIterator struct {
	Event *ContractPotentialWinner // Event containing the contract specifics and raw log

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
func (it *ContractPotentialWinnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPotentialWinner)
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
		it.Event = new(ContractPotentialWinner)
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
func (it *ContractPotentialWinnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPotentialWinnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPotentialWinner represents a PotentialWinner event raised by the Contract contract.
type ContractPotentialWinner struct {
	Round     *big.Int
	Bidder    common.Address
	BidAmount *big.Int
	TimeLeft  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPotentialWinner is a free log retrieval operation binding the contract event 0x7622c6e07e6dde4814d2d4642049258551877b7ac81d9e13a69e55ba6df91e9b.
//
// Solidity: event PotentialWinner(uint256 round, address bidder, uint256 bidAmount, uint256 timeLeft)
func (_Contract *ContractFilterer) FilterPotentialWinner(opts *bind.FilterOpts) (*ContractPotentialWinnerIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "PotentialWinner")
	if err != nil {
		return nil, err
	}
	return &ContractPotentialWinnerIterator{contract: _Contract.contract, event: "PotentialWinner", logs: logs, sub: sub}, nil
}

// WatchPotentialWinner is a free log subscription operation binding the contract event 0x7622c6e07e6dde4814d2d4642049258551877b7ac81d9e13a69e55ba6df91e9b.
//
// Solidity: event PotentialWinner(uint256 round, address bidder, uint256 bidAmount, uint256 timeLeft)
func (_Contract *ContractFilterer) WatchPotentialWinner(opts *bind.WatchOpts, sink chan<- *ContractPotentialWinner) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "PotentialWinner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPotentialWinner)
				if err := _Contract.contract.UnpackLog(event, "PotentialWinner", log); err != nil {
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

// ParsePotentialWinner is a log parse operation binding the contract event 0x7622c6e07e6dde4814d2d4642049258551877b7ac81d9e13a69e55ba6df91e9b.
//
// Solidity: event PotentialWinner(uint256 round, address bidder, uint256 bidAmount, uint256 timeLeft)
func (_Contract *ContractFilterer) ParsePotentialWinner(log types.Log) (*ContractPotentialWinner, error) {
	event := new(ContractPotentialWinner)
	if err := _Contract.contract.UnpackLog(event, "PotentialWinner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
