// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package jug

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

// JugMetaData contains all meta data concerning the Jug contract.
var JugMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"Drip\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"Init\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"base\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"drip\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ilks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"duty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rho\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// JugABI is the input ABI used to generate the binding from.
// Deprecated: Use JugMetaData.ABI instead.
var JugABI = JugMetaData.ABI

// Jug is an auto generated Go binding around an Ethereum contract.
type Jug struct {
	JugCaller     // Read-only binding to the contract
	JugTransactor // Write-only binding to the contract
	JugFilterer   // Log filterer for contract events
}

// JugCaller is an auto generated read-only Go binding around an Ethereum contract.
type JugCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JugTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JugTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JugFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JugFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JugSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JugSession struct {
	Contract     *Jug              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JugCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JugCallerSession struct {
	Contract *JugCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// JugTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JugTransactorSession struct {
	Contract     *JugTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JugRaw is an auto generated low-level Go binding around an Ethereum contract.
type JugRaw struct {
	Contract *Jug // Generic contract binding to access the raw methods on
}

// JugCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JugCallerRaw struct {
	Contract *JugCaller // Generic read-only contract binding to access the raw methods on
}

// JugTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JugTransactorRaw struct {
	Contract *JugTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJug creates a new instance of Jug, bound to a specific deployed contract.
func NewJug(address common.Address, backend bind.ContractBackend) (*Jug, error) {
	contract, err := bindJug(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Jug{JugCaller: JugCaller{contract: contract}, JugTransactor: JugTransactor{contract: contract}, JugFilterer: JugFilterer{contract: contract}}, nil
}

// NewJugCaller creates a new read-only instance of Jug, bound to a specific deployed contract.
func NewJugCaller(address common.Address, caller bind.ContractCaller) (*JugCaller, error) {
	contract, err := bindJug(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JugCaller{contract: contract}, nil
}

// NewJugTransactor creates a new write-only instance of Jug, bound to a specific deployed contract.
func NewJugTransactor(address common.Address, transactor bind.ContractTransactor) (*JugTransactor, error) {
	contract, err := bindJug(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JugTransactor{contract: contract}, nil
}

// NewJugFilterer creates a new log filterer instance of Jug, bound to a specific deployed contract.
func NewJugFilterer(address common.Address, filterer bind.ContractFilterer) (*JugFilterer, error) {
	contract, err := bindJug(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JugFilterer{contract: contract}, nil
}

// bindJug binds a generic wrapper to an already deployed contract.
func bindJug(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := JugMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Jug *JugRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Jug.Contract.JugCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Jug *JugRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jug.Contract.JugTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Jug *JugRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Jug.Contract.JugTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Jug *JugCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Jug.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Jug *JugTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jug.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Jug *JugTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Jug.Contract.contract.Transact(opts, method, params...)
}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(uint256)
func (_Jug *JugCaller) Base(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Jug.contract.Call(opts, &out, "base")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(uint256)
func (_Jug *JugSession) Base() (*big.Int, error) {
	return _Jug.Contract.Base(&_Jug.CallOpts)
}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(uint256)
func (_Jug *JugCallerSession) Base() (*big.Int, error) {
	return _Jug.Contract.Base(&_Jug.CallOpts)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(uint256 duty, uint256 rho)
func (_Jug *JugCaller) Ilks(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Duty *big.Int
	Rho  *big.Int
}, error) {
	var out []interface{}
	err := _Jug.contract.Call(opts, &out, "ilks", arg0)

	outstruct := new(struct {
		Duty *big.Int
		Rho  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Duty = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Rho = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(uint256 duty, uint256 rho)
func (_Jug *JugSession) Ilks(arg0 [32]byte) (struct {
	Duty *big.Int
	Rho  *big.Int
}, error) {
	return _Jug.Contract.Ilks(&_Jug.CallOpts, arg0)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(uint256 duty, uint256 rho)
func (_Jug *JugCallerSession) Ilks(arg0 [32]byte) (struct {
	Duty *big.Int
	Rho  *big.Int
}, error) {
	return _Jug.Contract.Ilks(&_Jug.CallOpts, arg0)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Jug *JugCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Jug.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Jug *JugSession) Vat() (common.Address, error) {
	return _Jug.Contract.Vat(&_Jug.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Jug *JugCallerSession) Vat() (common.Address, error) {
	return _Jug.Contract.Vat(&_Jug.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_Jug *JugCaller) Vow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Jug.contract.Call(opts, &out, "vow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_Jug *JugSession) Vow() (common.Address, error) {
	return _Jug.Contract.Vow(&_Jug.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_Jug *JugCallerSession) Vow() (common.Address, error) {
	return _Jug.Contract.Vow(&_Jug.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Jug *JugCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Jug.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Jug *JugSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Jug.Contract.Wards(&_Jug.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Jug *JugCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Jug.Contract.Wards(&_Jug.CallOpts, arg0)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Jug *JugTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Jug.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Jug *JugSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Jug.Contract.Deny(&_Jug.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Jug *JugTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Jug.Contract.Deny(&_Jug.TransactOpts, usr)
}

// Drip is a paid mutator transaction binding the contract method 0x44e2a5a8.
//
// Solidity: function drip(bytes32 ilk) returns(uint256 rate)
func (_Jug *JugTransactor) Drip(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _Jug.contract.Transact(opts, "drip", ilk)
}

// Drip is a paid mutator transaction binding the contract method 0x44e2a5a8.
//
// Solidity: function drip(bytes32 ilk) returns(uint256 rate)
func (_Jug *JugSession) Drip(ilk [32]byte) (*types.Transaction, error) {
	return _Jug.Contract.Drip(&_Jug.TransactOpts, ilk)
}

// Drip is a paid mutator transaction binding the contract method 0x44e2a5a8.
//
// Solidity: function drip(bytes32 ilk) returns(uint256 rate)
func (_Jug *JugTransactorSession) Drip(ilk [32]byte) (*types.Transaction, error) {
	return _Jug.Contract.Drip(&_Jug.TransactOpts, ilk)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_Jug *JugTransactor) File(opts *bind.TransactOpts, ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Jug.contract.Transact(opts, "file", ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_Jug *JugSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Jug.Contract.File(&_Jug.TransactOpts, ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_Jug *JugTransactorSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Jug.Contract.File(&_Jug.TransactOpts, ilk, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Jug *JugTransactor) File0(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Jug.contract.Transact(opts, "file0", what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Jug *JugSession) File0(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Jug.Contract.File0(&_Jug.TransactOpts, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Jug *JugTransactorSession) File0(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Jug.Contract.File0(&_Jug.TransactOpts, what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_Jug *JugTransactor) File1(opts *bind.TransactOpts, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _Jug.contract.Transact(opts, "file1", what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_Jug *JugSession) File1(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _Jug.Contract.File1(&_Jug.TransactOpts, what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_Jug *JugTransactorSession) File1(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _Jug.Contract.File1(&_Jug.TransactOpts, what, data)
}

// Init is a paid mutator transaction binding the contract method 0x3b663195.
//
// Solidity: function init(bytes32 ilk) returns()
func (_Jug *JugTransactor) Init(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _Jug.contract.Transact(opts, "init", ilk)
}

// Init is a paid mutator transaction binding the contract method 0x3b663195.
//
// Solidity: function init(bytes32 ilk) returns()
func (_Jug *JugSession) Init(ilk [32]byte) (*types.Transaction, error) {
	return _Jug.Contract.Init(&_Jug.TransactOpts, ilk)
}

// Init is a paid mutator transaction binding the contract method 0x3b663195.
//
// Solidity: function init(bytes32 ilk) returns()
func (_Jug *JugTransactorSession) Init(ilk [32]byte) (*types.Transaction, error) {
	return _Jug.Contract.Init(&_Jug.TransactOpts, ilk)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Jug *JugTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Jug.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Jug *JugSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Jug.Contract.Rely(&_Jug.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Jug *JugTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Jug.Contract.Rely(&_Jug.TransactOpts, usr)
}

// JugDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the Jug contract.
type JugDenyIterator struct {
	Event *JugDeny // Event containing the contract specifics and raw log

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
func (it *JugDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JugDeny)
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
		it.Event = new(JugDeny)
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
func (it *JugDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JugDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JugDeny represents a Deny event raised by the Jug contract.
type JugDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_Jug *JugFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*JugDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Jug.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &JugDenyIterator{contract: _Jug.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_Jug *JugFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *JugDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Jug.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JugDeny)
				if err := _Jug.contract.UnpackLog(event, "Deny", log); err != nil {
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

// ParseDeny is a log parse operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_Jug *JugFilterer) ParseDeny(log types.Log) (*JugDeny, error) {
	event := new(JugDeny)
	if err := _Jug.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JugDripIterator is returned from FilterDrip and is used to iterate over the raw logs and unpacked data for Drip events raised by the Jug contract.
type JugDripIterator struct {
	Event *JugDrip // Event containing the contract specifics and raw log

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
func (it *JugDripIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JugDrip)
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
		it.Event = new(JugDrip)
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
func (it *JugDripIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JugDripIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JugDrip represents a Drip event raised by the Jug contract.
type JugDrip struct {
	Ilk  [32]byte
	Rate *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDrip is a free log retrieval operation binding the contract event 0x2fbd8559330c6faf38d4ee435b8463427bd58cd6bc7edd1a4e6f3ab6b57de71e.
//
// Solidity: event Drip(bytes32 indexed ilk, uint256 indexed rate)
func (_Jug *JugFilterer) FilterDrip(opts *bind.FilterOpts, ilk [][32]byte, rate []*big.Int) (*JugDripIterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var rateRule []interface{}
	for _, rateItem := range rate {
		rateRule = append(rateRule, rateItem)
	}

	logs, sub, err := _Jug.contract.FilterLogs(opts, "Drip", ilkRule, rateRule)
	if err != nil {
		return nil, err
	}
	return &JugDripIterator{contract: _Jug.contract, event: "Drip", logs: logs, sub: sub}, nil
}

// WatchDrip is a free log subscription operation binding the contract event 0x2fbd8559330c6faf38d4ee435b8463427bd58cd6bc7edd1a4e6f3ab6b57de71e.
//
// Solidity: event Drip(bytes32 indexed ilk, uint256 indexed rate)
func (_Jug *JugFilterer) WatchDrip(opts *bind.WatchOpts, sink chan<- *JugDrip, ilk [][32]byte, rate []*big.Int) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var rateRule []interface{}
	for _, rateItem := range rate {
		rateRule = append(rateRule, rateItem)
	}

	logs, sub, err := _Jug.contract.WatchLogs(opts, "Drip", ilkRule, rateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JugDrip)
				if err := _Jug.contract.UnpackLog(event, "Drip", log); err != nil {
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

// ParseDrip is a log parse operation binding the contract event 0x2fbd8559330c6faf38d4ee435b8463427bd58cd6bc7edd1a4e6f3ab6b57de71e.
//
// Solidity: event Drip(bytes32 indexed ilk, uint256 indexed rate)
func (_Jug *JugFilterer) ParseDrip(log types.Log) (*JugDrip, error) {
	event := new(JugDrip)
	if err := _Jug.contract.UnpackLog(event, "Drip", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JugFileIterator is returned from FilterFile and is used to iterate over the raw logs and unpacked data for File events raised by the Jug contract.
type JugFileIterator struct {
	Event *JugFile // Event containing the contract specifics and raw log

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
func (it *JugFileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JugFile)
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
		it.Event = new(JugFile)
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
func (it *JugFileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JugFileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JugFile represents a File event raised by the Jug contract.
type JugFile struct {
	Ilk  [32]byte
	What [32]byte
	Data *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile is a free log retrieval operation binding the contract event 0x851aa1caf4888170ad8875449d18f0f512fd6deb2a6571ea1a41fb9f95acbcd1.
//
// Solidity: event File(bytes32 indexed ilk, bytes32 indexed what, uint256 indexed data)
func (_Jug *JugFilterer) FilterFile(opts *bind.FilterOpts, ilk [][32]byte, what [][32]byte, data []*big.Int) (*JugFileIterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}
	var dataRule []interface{}
	for _, dataItem := range data {
		dataRule = append(dataRule, dataItem)
	}

	logs, sub, err := _Jug.contract.FilterLogs(opts, "File", ilkRule, whatRule, dataRule)
	if err != nil {
		return nil, err
	}
	return &JugFileIterator{contract: _Jug.contract, event: "File", logs: logs, sub: sub}, nil
}

// WatchFile is a free log subscription operation binding the contract event 0x851aa1caf4888170ad8875449d18f0f512fd6deb2a6571ea1a41fb9f95acbcd1.
//
// Solidity: event File(bytes32 indexed ilk, bytes32 indexed what, uint256 indexed data)
func (_Jug *JugFilterer) WatchFile(opts *bind.WatchOpts, sink chan<- *JugFile, ilk [][32]byte, what [][32]byte, data []*big.Int) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}
	var dataRule []interface{}
	for _, dataItem := range data {
		dataRule = append(dataRule, dataItem)
	}

	logs, sub, err := _Jug.contract.WatchLogs(opts, "File", ilkRule, whatRule, dataRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JugFile)
				if err := _Jug.contract.UnpackLog(event, "File", log); err != nil {
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

// ParseFile is a log parse operation binding the contract event 0x851aa1caf4888170ad8875449d18f0f512fd6deb2a6571ea1a41fb9f95acbcd1.
//
// Solidity: event File(bytes32 indexed ilk, bytes32 indexed what, uint256 indexed data)
func (_Jug *JugFilterer) ParseFile(log types.Log) (*JugFile, error) {
	event := new(JugFile)
	if err := _Jug.contract.UnpackLog(event, "File", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JugFile0Iterator is returned from FilterFile0 and is used to iterate over the raw logs and unpacked data for File0 events raised by the Jug contract.
type JugFile0Iterator struct {
	Event *JugFile0 // Event containing the contract specifics and raw log

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
func (it *JugFile0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JugFile0)
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
		it.Event = new(JugFile0)
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
func (it *JugFile0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JugFile0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JugFile0 represents a File0 event raised by the Jug contract.
type JugFile0 struct {
	What [32]byte
	Data *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile0 is a free log retrieval operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_Jug *JugFilterer) FilterFile0(opts *bind.FilterOpts, what [][32]byte) (*JugFile0Iterator, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Jug.contract.FilterLogs(opts, "File0", whatRule)
	if err != nil {
		return nil, err
	}
	return &JugFile0Iterator{contract: _Jug.contract, event: "File0", logs: logs, sub: sub}, nil
}

// WatchFile0 is a free log subscription operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_Jug *JugFilterer) WatchFile0(opts *bind.WatchOpts, sink chan<- *JugFile0, what [][32]byte) (event.Subscription, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Jug.contract.WatchLogs(opts, "File0", whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JugFile0)
				if err := _Jug.contract.UnpackLog(event, "File0", log); err != nil {
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

// ParseFile0 is a log parse operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_Jug *JugFilterer) ParseFile0(log types.Log) (*JugFile0, error) {
	event := new(JugFile0)
	if err := _Jug.contract.UnpackLog(event, "File0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JugFile1Iterator is returned from FilterFile1 and is used to iterate over the raw logs and unpacked data for File1 events raised by the Jug contract.
type JugFile1Iterator struct {
	Event *JugFile1 // Event containing the contract specifics and raw log

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
func (it *JugFile1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JugFile1)
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
		it.Event = new(JugFile1)
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
func (it *JugFile1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JugFile1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JugFile1 represents a File1 event raised by the Jug contract.
type JugFile1 struct {
	What [32]byte
	Data common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile1 is a free log retrieval operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 indexed what, address data)
func (_Jug *JugFilterer) FilterFile1(opts *bind.FilterOpts, what [][32]byte) (*JugFile1Iterator, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Jug.contract.FilterLogs(opts, "File1", whatRule)
	if err != nil {
		return nil, err
	}
	return &JugFile1Iterator{contract: _Jug.contract, event: "File1", logs: logs, sub: sub}, nil
}

// WatchFile1 is a free log subscription operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 indexed what, address data)
func (_Jug *JugFilterer) WatchFile1(opts *bind.WatchOpts, sink chan<- *JugFile1, what [][32]byte) (event.Subscription, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Jug.contract.WatchLogs(opts, "File1", whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JugFile1)
				if err := _Jug.contract.UnpackLog(event, "File1", log); err != nil {
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

// ParseFile1 is a log parse operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 indexed what, address data)
func (_Jug *JugFilterer) ParseFile1(log types.Log) (*JugFile1, error) {
	event := new(JugFile1)
	if err := _Jug.contract.UnpackLog(event, "File1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JugInitIterator is returned from FilterInit and is used to iterate over the raw logs and unpacked data for Init events raised by the Jug contract.
type JugInitIterator struct {
	Event *JugInit // Event containing the contract specifics and raw log

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
func (it *JugInitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JugInit)
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
		it.Event = new(JugInit)
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
func (it *JugInitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JugInitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JugInit represents a Init event raised by the Jug contract.
type JugInit struct {
	Ilk [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterInit is a free log retrieval operation binding the contract event 0xeeb45f27c5b399a603237b10d4803743d494bfc24c3a004cadb716c41033a555.
//
// Solidity: event Init(bytes32 indexed ilk)
func (_Jug *JugFilterer) FilterInit(opts *bind.FilterOpts, ilk [][32]byte) (*JugInitIterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _Jug.contract.FilterLogs(opts, "Init", ilkRule)
	if err != nil {
		return nil, err
	}
	return &JugInitIterator{contract: _Jug.contract, event: "Init", logs: logs, sub: sub}, nil
}

// WatchInit is a free log subscription operation binding the contract event 0xeeb45f27c5b399a603237b10d4803743d494bfc24c3a004cadb716c41033a555.
//
// Solidity: event Init(bytes32 indexed ilk)
func (_Jug *JugFilterer) WatchInit(opts *bind.WatchOpts, sink chan<- *JugInit, ilk [][32]byte) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _Jug.contract.WatchLogs(opts, "Init", ilkRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JugInit)
				if err := _Jug.contract.UnpackLog(event, "Init", log); err != nil {
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

// ParseInit is a log parse operation binding the contract event 0xeeb45f27c5b399a603237b10d4803743d494bfc24c3a004cadb716c41033a555.
//
// Solidity: event Init(bytes32 indexed ilk)
func (_Jug *JugFilterer) ParseInit(log types.Log) (*JugInit, error) {
	event := new(JugInit)
	if err := _Jug.contract.UnpackLog(event, "Init", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JugRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the Jug contract.
type JugRelyIterator struct {
	Event *JugRely // Event containing the contract specifics and raw log

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
func (it *JugRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JugRely)
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
		it.Event = new(JugRely)
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
func (it *JugRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JugRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JugRely represents a Rely event raised by the Jug contract.
type JugRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_Jug *JugFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*JugRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Jug.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &JugRelyIterator{contract: _Jug.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_Jug *JugFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *JugRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Jug.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JugRely)
				if err := _Jug.contract.UnpackLog(event, "Rely", log); err != nil {
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

// ParseRely is a log parse operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_Jug *JugFilterer) ParseRely(log types.Log) (*JugRely, error) {
	event := new(JugRely)
	if err := _Jug.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
