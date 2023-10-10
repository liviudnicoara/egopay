// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
	_ = abi.ConvertType
)

// SplitBillMetaData contains all meta data concerning the SplitBill contract.
var SplitBillMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"payers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amountInUSD\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BillAlreadyPayed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPayers\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Completed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Payed\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"i_amountPerPayerInUSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_totalAmountInUSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"}],\"name\":\"payerStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"hasPayed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_amountPayedInUSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60e060405234801561000f575f80fd5b5060405161094838038061094883398101604081905261002e91610102565b81515f0361004f57604051630f95d2d960e11b815260040160405180910390fd5b3360805260a0819052815161006490826101c9565b60c0525f5b82518110156100cb575f60015f858481518110610088576100886101e8565b6020908102919091018101516001600160a01b031682528101919091526040015f20805460ff1916911515919091179055806100c3816101fc565b915050610069565b505050610220565b634e487b7160e01b5f52604160045260245ffd5b80516001600160a01b03811681146100fd575f80fd5b919050565b5f8060408385031215610113575f80fd5b82516001600160401b0380821115610129575f80fd5b818501915085601f83011261013c575f80fd5b8151602082821115610150576101506100d3565b8160051b604051601f19603f83011681018181108682111715610175576101756100d3565b604052928352818301935084810182019289841115610192575f80fd5b948201945b838610156101b7576101a8866100e7565b85529482019493820193610197565b97909101519698969750505050505050565b5f826101e357634e487b7160e01b5f52601260045260245ffd5b500490565b634e487b7160e01b5f52603260045260245ffd5b5f6001820161021957634e487b7160e01b5f52601160045260245ffd5b5060010190565b60805160a05160c0516106ce61027a5f395f8181609b015281816101ad01528181610221015261037001525f818161013b01528181610258015261031701525f818161016e0152818161028001526103cc01526106ce5ff3fe608060405260043610610073575f3560e01c806387d99d451161004d57806387d99d45146100d8578063b61d276a14610116578063b73f94371461012a578063dba6335f1461015d57610082565b80631b9265b81461008257806328c1d3781461008a5780633ccfd60b146100d057610082565b36610082576100806101a8565b005b6100806101a8565b348015610095575f80fd5b506100bd7f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020015b60405180910390f35b6100806103c1565b3480156100e3575f80fd5b506101066100f236600461056f565b60016020525f908152604090205460ff1681565b60405190151581526020016100c7565b348015610121575f80fd5b506100bd5f5481565b348015610135575f80fd5b506100bd7f000000000000000000000000000000000000000000000000000000000000000081565b348015610168575f80fd5b506101907f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100c7565b6101da7f0000000000000000000000000000000000000000000000000000000000000000670de0b6b3a76400006105b0565b6101e334610498565b1015610202576040516334b2073960e11b815260040160405180910390fd5b335f9081526001602081905260408220805460ff1916909117905580547f00000000000000000000000000000000000000000000000000000000000000009190819061024f9084906105cd565b90915550505f547f00000000000000000000000000000000000000000000000000000000000000009003610367575f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316476040515f6040518083038185875af1925050503d805f81146102e6576040519150601f19603f3d011682016040523d82523d5f602084013e6102eb565b606091505b50909150508015155f0361031257604051631d42c86760e21b815260040160405180910390fd5b6040517f000000000000000000000000000000000000000000000000000000000000000081527fdfd517ed69f8a0a57d49fe494e4864fac3cfe3585c14c0bfddf39f72463ec3fd9060200160405180910390a1505b604080513381527f000000000000000000000000000000000000000000000000000000000000000060208201527f7be9078cfa949e4b9e15888282e2a07a7fae4958503c18c81db93c5d41352497910160405180910390a1565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461040a576040516330cd747160e01b815260040160405180910390fd5b475f0361042a57604051634a8e036d60e11b815260040160405180910390fd5b6040515f90339047908381818185875af1925050503d805f8114610469576040519150601f19603f3d011682016040523d82523d5f602084013e61046e565b606091505b50909150508015155f0361049557604051631d42c86760e21b815260040160405180910390fd5b50565b5f806104a26104ca565b90505f670de0b6b3a76400006104b885846105b0565b6104c291906105e0565b949350505050565b5f8073694aa1769357215de4fac081bf1f309adc32530690505f816001600160a01b031663feaf968c6040518163ffffffff1660e01b815260040160a060405180830381865afa158015610520573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610544919061061d565b505050915050805f03610559575f9250505090565b610568816402540be400610669565b9250505090565b5f6020828403121561057f575f80fd5b81356001600160a01b0381168114610595575f80fd5b9392505050565b634e487b7160e01b5f52601160045260245ffd5b80820281158282048414176105c7576105c761059c565b92915050565b808201808211156105c7576105c761059c565b5f826105fa57634e487b7160e01b5f52601260045260245ffd5b500490565b805169ffffffffffffffffffff81168114610618575f80fd5b919050565b5f805f805f60a08688031215610631575f80fd5b61063a866105ff565b945060208601519350604086015192506060860151915061065d608087016105ff565b90509295509295909350565b8082025f8212600160ff1b841416156106845761068461059c565b81810583148215176105c7576105c761059c56fea264697066735822122042fd2b00c553caa387070b1681c4c5aba59ce5e1aaadf8fdd4955f8989ab64ae64736f6c63430008150033",
}

// SplitBillABI is the input ABI used to generate the binding from.
// Deprecated: Use SplitBillMetaData.ABI instead.
var SplitBillABI = SplitBillMetaData.ABI

// SplitBillBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SplitBillMetaData.Bin instead.
var SplitBillBin = SplitBillMetaData.Bin

// DeploySplitBill deploys a new Ethereum contract, binding an instance of SplitBill to it.
func DeploySplitBill(auth *bind.TransactOpts, backend bind.ContractBackend, payers []common.Address, amountInUSD *big.Int) (common.Address, *types.Transaction, *SplitBill, error) {
	parsed, err := SplitBillMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SplitBillBin), backend, payers, amountInUSD)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SplitBill{SplitBillCaller: SplitBillCaller{contract: contract}, SplitBillTransactor: SplitBillTransactor{contract: contract}, SplitBillFilterer: SplitBillFilterer{contract: contract}}, nil
}

// SplitBill is an auto generated Go binding around an Ethereum contract.
type SplitBill struct {
	SplitBillCaller     // Read-only binding to the contract
	SplitBillTransactor // Write-only binding to the contract
	SplitBillFilterer   // Log filterer for contract events
}

// SplitBillCaller is an auto generated read-only Go binding around an Ethereum contract.
type SplitBillCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SplitBillTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SplitBillTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SplitBillFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SplitBillFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SplitBillSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SplitBillSession struct {
	Contract     *SplitBill        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SplitBillCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SplitBillCallerSession struct {
	Contract *SplitBillCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SplitBillTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SplitBillTransactorSession struct {
	Contract     *SplitBillTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SplitBillRaw is an auto generated low-level Go binding around an Ethereum contract.
type SplitBillRaw struct {
	Contract *SplitBill // Generic contract binding to access the raw methods on
}

// SplitBillCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SplitBillCallerRaw struct {
	Contract *SplitBillCaller // Generic read-only contract binding to access the raw methods on
}

// SplitBillTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SplitBillTransactorRaw struct {
	Contract *SplitBillTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSplitBill creates a new instance of SplitBill, bound to a specific deployed contract.
func NewSplitBill(address common.Address, backend bind.ContractBackend) (*SplitBill, error) {
	contract, err := bindSplitBill(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SplitBill{SplitBillCaller: SplitBillCaller{contract: contract}, SplitBillTransactor: SplitBillTransactor{contract: contract}, SplitBillFilterer: SplitBillFilterer{contract: contract}}, nil
}

// NewSplitBillCaller creates a new read-only instance of SplitBill, bound to a specific deployed contract.
func NewSplitBillCaller(address common.Address, caller bind.ContractCaller) (*SplitBillCaller, error) {
	contract, err := bindSplitBill(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SplitBillCaller{contract: contract}, nil
}

// NewSplitBillTransactor creates a new write-only instance of SplitBill, bound to a specific deployed contract.
func NewSplitBillTransactor(address common.Address, transactor bind.ContractTransactor) (*SplitBillTransactor, error) {
	contract, err := bindSplitBill(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SplitBillTransactor{contract: contract}, nil
}

// NewSplitBillFilterer creates a new log filterer instance of SplitBill, bound to a specific deployed contract.
func NewSplitBillFilterer(address common.Address, filterer bind.ContractFilterer) (*SplitBillFilterer, error) {
	contract, err := bindSplitBill(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SplitBillFilterer{contract: contract}, nil
}

// bindSplitBill binds a generic wrapper to an already deployed contract.
func bindSplitBill(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SplitBillMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SplitBill *SplitBillRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SplitBill.Contract.SplitBillCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SplitBill *SplitBillRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SplitBill.Contract.SplitBillTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SplitBill *SplitBillRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SplitBill.Contract.SplitBillTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SplitBill *SplitBillCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SplitBill.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SplitBill *SplitBillTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SplitBill.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SplitBill *SplitBillTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SplitBill.Contract.contract.Transact(opts, method, params...)
}

// IAmountPerPayerInUSD is a free data retrieval call binding the contract method 0x28c1d378.
//
// Solidity: function i_amountPerPayerInUSD() view returns(uint256)
func (_SplitBill *SplitBillCaller) IAmountPerPayerInUSD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SplitBill.contract.Call(opts, &out, "i_amountPerPayerInUSD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IAmountPerPayerInUSD is a free data retrieval call binding the contract method 0x28c1d378.
//
// Solidity: function i_amountPerPayerInUSD() view returns(uint256)
func (_SplitBill *SplitBillSession) IAmountPerPayerInUSD() (*big.Int, error) {
	return _SplitBill.Contract.IAmountPerPayerInUSD(&_SplitBill.CallOpts)
}

// IAmountPerPayerInUSD is a free data retrieval call binding the contract method 0x28c1d378.
//
// Solidity: function i_amountPerPayerInUSD() view returns(uint256)
func (_SplitBill *SplitBillCallerSession) IAmountPerPayerInUSD() (*big.Int, error) {
	return _SplitBill.Contract.IAmountPerPayerInUSD(&_SplitBill.CallOpts)
}

// IOwner is a free data retrieval call binding the contract method 0xdba6335f.
//
// Solidity: function i_owner() view returns(address)
func (_SplitBill *SplitBillCaller) IOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SplitBill.contract.Call(opts, &out, "i_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IOwner is a free data retrieval call binding the contract method 0xdba6335f.
//
// Solidity: function i_owner() view returns(address)
func (_SplitBill *SplitBillSession) IOwner() (common.Address, error) {
	return _SplitBill.Contract.IOwner(&_SplitBill.CallOpts)
}

// IOwner is a free data retrieval call binding the contract method 0xdba6335f.
//
// Solidity: function i_owner() view returns(address)
func (_SplitBill *SplitBillCallerSession) IOwner() (common.Address, error) {
	return _SplitBill.Contract.IOwner(&_SplitBill.CallOpts)
}

// ITotalAmountInUSD is a free data retrieval call binding the contract method 0xb73f9437.
//
// Solidity: function i_totalAmountInUSD() view returns(uint256)
func (_SplitBill *SplitBillCaller) ITotalAmountInUSD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SplitBill.contract.Call(opts, &out, "i_totalAmountInUSD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ITotalAmountInUSD is a free data retrieval call binding the contract method 0xb73f9437.
//
// Solidity: function i_totalAmountInUSD() view returns(uint256)
func (_SplitBill *SplitBillSession) ITotalAmountInUSD() (*big.Int, error) {
	return _SplitBill.Contract.ITotalAmountInUSD(&_SplitBill.CallOpts)
}

// ITotalAmountInUSD is a free data retrieval call binding the contract method 0xb73f9437.
//
// Solidity: function i_totalAmountInUSD() view returns(uint256)
func (_SplitBill *SplitBillCallerSession) ITotalAmountInUSD() (*big.Int, error) {
	return _SplitBill.Contract.ITotalAmountInUSD(&_SplitBill.CallOpts)
}

// PayerStatus is a free data retrieval call binding the contract method 0x87d99d45.
//
// Solidity: function payerStatus(address payer) view returns(bool hasPayed)
func (_SplitBill *SplitBillCaller) PayerStatus(opts *bind.CallOpts, payer common.Address) (bool, error) {
	var out []interface{}
	err := _SplitBill.contract.Call(opts, &out, "payerStatus", payer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PayerStatus is a free data retrieval call binding the contract method 0x87d99d45.
//
// Solidity: function payerStatus(address payer) view returns(bool hasPayed)
func (_SplitBill *SplitBillSession) PayerStatus(payer common.Address) (bool, error) {
	return _SplitBill.Contract.PayerStatus(&_SplitBill.CallOpts, payer)
}

// PayerStatus is a free data retrieval call binding the contract method 0x87d99d45.
//
// Solidity: function payerStatus(address payer) view returns(bool hasPayed)
func (_SplitBill *SplitBillCallerSession) PayerStatus(payer common.Address) (bool, error) {
	return _SplitBill.Contract.PayerStatus(&_SplitBill.CallOpts, payer)
}

// SAmountPayedInUSD is a free data retrieval call binding the contract method 0xb61d276a.
//
// Solidity: function s_amountPayedInUSD() view returns(uint256)
func (_SplitBill *SplitBillCaller) SAmountPayedInUSD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SplitBill.contract.Call(opts, &out, "s_amountPayedInUSD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SAmountPayedInUSD is a free data retrieval call binding the contract method 0xb61d276a.
//
// Solidity: function s_amountPayedInUSD() view returns(uint256)
func (_SplitBill *SplitBillSession) SAmountPayedInUSD() (*big.Int, error) {
	return _SplitBill.Contract.SAmountPayedInUSD(&_SplitBill.CallOpts)
}

// SAmountPayedInUSD is a free data retrieval call binding the contract method 0xb61d276a.
//
// Solidity: function s_amountPayedInUSD() view returns(uint256)
func (_SplitBill *SplitBillCallerSession) SAmountPayedInUSD() (*big.Int, error) {
	return _SplitBill.Contract.SAmountPayedInUSD(&_SplitBill.CallOpts)
}

// Pay is a paid mutator transaction binding the contract method 0x1b9265b8.
//
// Solidity: function pay() payable returns()
func (_SplitBill *SplitBillTransactor) Pay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SplitBill.contract.Transact(opts, "pay")
}

// Pay is a paid mutator transaction binding the contract method 0x1b9265b8.
//
// Solidity: function pay() payable returns()
func (_SplitBill *SplitBillSession) Pay() (*types.Transaction, error) {
	return _SplitBill.Contract.Pay(&_SplitBill.TransactOpts)
}

// Pay is a paid mutator transaction binding the contract method 0x1b9265b8.
//
// Solidity: function pay() payable returns()
func (_SplitBill *SplitBillTransactorSession) Pay() (*types.Transaction, error) {
	return _SplitBill.Contract.Pay(&_SplitBill.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_SplitBill *SplitBillTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SplitBill.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_SplitBill *SplitBillSession) Withdraw() (*types.Transaction, error) {
	return _SplitBill.Contract.Withdraw(&_SplitBill.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_SplitBill *SplitBillTransactorSession) Withdraw() (*types.Transaction, error) {
	return _SplitBill.Contract.Withdraw(&_SplitBill.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SplitBill *SplitBillTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _SplitBill.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SplitBill *SplitBillSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SplitBill.Contract.Fallback(&_SplitBill.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SplitBill *SplitBillTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SplitBill.Contract.Fallback(&_SplitBill.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SplitBill *SplitBillTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SplitBill.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SplitBill *SplitBillSession) Receive() (*types.Transaction, error) {
	return _SplitBill.Contract.Receive(&_SplitBill.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SplitBill *SplitBillTransactorSession) Receive() (*types.Transaction, error) {
	return _SplitBill.Contract.Receive(&_SplitBill.TransactOpts)
}

// SplitBillCompletedIterator is returned from FilterCompleted and is used to iterate over the raw logs and unpacked data for Completed events raised by the SplitBill contract.
type SplitBillCompletedIterator struct {
	Event *SplitBillCompleted // Event containing the contract specifics and raw log

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
func (it *SplitBillCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SplitBillCompleted)
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
		it.Event = new(SplitBillCompleted)
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
func (it *SplitBillCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SplitBillCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SplitBillCompleted represents a Completed event raised by the SplitBill contract.
type SplitBillCompleted struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCompleted is a free log retrieval operation binding the contract event 0xdfd517ed69f8a0a57d49fe494e4864fac3cfe3585c14c0bfddf39f72463ec3fd.
//
// Solidity: event Completed(uint256 amount)
func (_SplitBill *SplitBillFilterer) FilterCompleted(opts *bind.FilterOpts) (*SplitBillCompletedIterator, error) {

	logs, sub, err := _SplitBill.contract.FilterLogs(opts, "Completed")
	if err != nil {
		return nil, err
	}
	return &SplitBillCompletedIterator{contract: _SplitBill.contract, event: "Completed", logs: logs, sub: sub}, nil
}

// WatchCompleted is a free log subscription operation binding the contract event 0xdfd517ed69f8a0a57d49fe494e4864fac3cfe3585c14c0bfddf39f72463ec3fd.
//
// Solidity: event Completed(uint256 amount)
func (_SplitBill *SplitBillFilterer) WatchCompleted(opts *bind.WatchOpts, sink chan<- *SplitBillCompleted) (event.Subscription, error) {

	logs, sub, err := _SplitBill.contract.WatchLogs(opts, "Completed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SplitBillCompleted)
				if err := _SplitBill.contract.UnpackLog(event, "Completed", log); err != nil {
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

// ParseCompleted is a log parse operation binding the contract event 0xdfd517ed69f8a0a57d49fe494e4864fac3cfe3585c14c0bfddf39f72463ec3fd.
//
// Solidity: event Completed(uint256 amount)
func (_SplitBill *SplitBillFilterer) ParseCompleted(log types.Log) (*SplitBillCompleted, error) {
	event := new(SplitBillCompleted)
	if err := _SplitBill.contract.UnpackLog(event, "Completed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SplitBillPayedIterator is returned from FilterPayed and is used to iterate over the raw logs and unpacked data for Payed events raised by the SplitBill contract.
type SplitBillPayedIterator struct {
	Event *SplitBillPayed // Event containing the contract specifics and raw log

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
func (it *SplitBillPayedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SplitBillPayed)
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
		it.Event = new(SplitBillPayed)
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
func (it *SplitBillPayedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SplitBillPayedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SplitBillPayed represents a Payed event raised by the SplitBill contract.
type SplitBillPayed struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPayed is a free log retrieval operation binding the contract event 0x7be9078cfa949e4b9e15888282e2a07a7fae4958503c18c81db93c5d41352497.
//
// Solidity: event Payed(address from, uint256 amount)
func (_SplitBill *SplitBillFilterer) FilterPayed(opts *bind.FilterOpts) (*SplitBillPayedIterator, error) {

	logs, sub, err := _SplitBill.contract.FilterLogs(opts, "Payed")
	if err != nil {
		return nil, err
	}
	return &SplitBillPayedIterator{contract: _SplitBill.contract, event: "Payed", logs: logs, sub: sub}, nil
}

// WatchPayed is a free log subscription operation binding the contract event 0x7be9078cfa949e4b9e15888282e2a07a7fae4958503c18c81db93c5d41352497.
//
// Solidity: event Payed(address from, uint256 amount)
func (_SplitBill *SplitBillFilterer) WatchPayed(opts *bind.WatchOpts, sink chan<- *SplitBillPayed) (event.Subscription, error) {

	logs, sub, err := _SplitBill.contract.WatchLogs(opts, "Payed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SplitBillPayed)
				if err := _SplitBill.contract.UnpackLog(event, "Payed", log); err != nil {
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

// ParsePayed is a log parse operation binding the contract event 0x7be9078cfa949e4b9e15888282e2a07a7fae4958503c18c81db93c5d41352497.
//
// Solidity: event Payed(address from, uint256 amount)
func (_SplitBill *SplitBillFilterer) ParsePayed(log types.Log) (*SplitBillPayed, error) {
	event := new(SplitBillPayed)
	if err := _SplitBill.contract.UnpackLog(event, "Payed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
