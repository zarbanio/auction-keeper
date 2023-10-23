// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package median

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

// MedianMetaData contains all meta data concerning the Median contract.
var MedianMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_wat\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"age\",\"type\":\"uint256\"}],\"name\":\"LogMedianPrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"age\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"bud\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"a\",\"type\":\"address[]\"}],\"name\":\"diss\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"diss\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"a\",\"type\":\"address[]\"}],\"name\":\"drop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"a\",\"type\":\"address[]\"}],\"name\":\"kiss\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"kiss\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"a\",\"type\":\"address[]\"}],\"name\":\"lift\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"orcl\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"peek\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"val_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"age_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"poke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"read\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bar_\",\"type\":\"uint256\"}],\"name\":\"setBar\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"slot\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wat\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MedianABI is the input ABI used to generate the binding from.
// Deprecated: Use MedianMetaData.ABI instead.
var MedianABI = MedianMetaData.ABI

// Median is an auto generated Go binding around an Ethereum contract.
type Median struct {
	MedianCaller     // Read-only binding to the contract
	MedianTransactor // Write-only binding to the contract
	MedianFilterer   // Log filterer for contract events
}

// MedianCaller is an auto generated read-only Go binding around an Ethereum contract.
type MedianCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MedianTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MedianTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MedianFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MedianFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MedianSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MedianSession struct {
	Contract     *Median           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MedianCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MedianCallerSession struct {
	Contract *MedianCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MedianTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MedianTransactorSession struct {
	Contract     *MedianTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MedianRaw is an auto generated low-level Go binding around an Ethereum contract.
type MedianRaw struct {
	Contract *Median // Generic contract binding to access the raw methods on
}

// MedianCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MedianCallerRaw struct {
	Contract *MedianCaller // Generic read-only contract binding to access the raw methods on
}

// MedianTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MedianTransactorRaw struct {
	Contract *MedianTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMedian creates a new instance of Median, bound to a specific deployed contract.
func NewMedian(address common.Address, backend bind.ContractBackend) (*Median, error) {
	contract, err := bindMedian(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Median{MedianCaller: MedianCaller{contract: contract}, MedianTransactor: MedianTransactor{contract: contract}, MedianFilterer: MedianFilterer{contract: contract}}, nil
}

// NewMedianCaller creates a new read-only instance of Median, bound to a specific deployed contract.
func NewMedianCaller(address common.Address, caller bind.ContractCaller) (*MedianCaller, error) {
	contract, err := bindMedian(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MedianCaller{contract: contract}, nil
}

// NewMedianTransactor creates a new write-only instance of Median, bound to a specific deployed contract.
func NewMedianTransactor(address common.Address, transactor bind.ContractTransactor) (*MedianTransactor, error) {
	contract, err := bindMedian(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MedianTransactor{contract: contract}, nil
}

// NewMedianFilterer creates a new log filterer instance of Median, bound to a specific deployed contract.
func NewMedianFilterer(address common.Address, filterer bind.ContractFilterer) (*MedianFilterer, error) {
	contract, err := bindMedian(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MedianFilterer{contract: contract}, nil
}

// bindMedian binds a generic wrapper to an already deployed contract.
func bindMedian(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MedianMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Median *MedianRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Median.Contract.MedianCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Median *MedianRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Median.Contract.MedianTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Median *MedianRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Median.Contract.MedianTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Median *MedianCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Median.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Median *MedianTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Median.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Median *MedianTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Median.Contract.contract.Transact(opts, method, params...)
}

// Age is a free data retrieval call binding the contract method 0x262a9dff.
//
// Solidity: function age() view returns(uint32)
func (_Median *MedianCaller) Age(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Median.contract.Call(opts, &out, "age")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Age is a free data retrieval call binding the contract method 0x262a9dff.
//
// Solidity: function age() view returns(uint32)
func (_Median *MedianSession) Age() (uint32, error) {
	return _Median.Contract.Age(&_Median.CallOpts)
}

// Age is a free data retrieval call binding the contract method 0x262a9dff.
//
// Solidity: function age() view returns(uint32)
func (_Median *MedianCallerSession) Age() (uint32, error) {
	return _Median.Contract.Age(&_Median.CallOpts)
}

// Bar is a free data retrieval call binding the contract method 0xfebb0f7e.
//
// Solidity: function bar() view returns(uint256)
func (_Median *MedianCaller) Bar(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Median.contract.Call(opts, &out, "bar")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Bar is a free data retrieval call binding the contract method 0xfebb0f7e.
//
// Solidity: function bar() view returns(uint256)
func (_Median *MedianSession) Bar() (*big.Int, error) {
	return _Median.Contract.Bar(&_Median.CallOpts)
}

// Bar is a free data retrieval call binding the contract method 0xfebb0f7e.
//
// Solidity: function bar() view returns(uint256)
func (_Median *MedianCallerSession) Bar() (*big.Int, error) {
	return _Median.Contract.Bar(&_Median.CallOpts)
}

// Bud is a free data retrieval call binding the contract method 0x4fce7a2a.
//
// Solidity: function bud(address ) view returns(uint256)
func (_Median *MedianCaller) Bud(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Median.contract.Call(opts, &out, "bud", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Bud is a free data retrieval call binding the contract method 0x4fce7a2a.
//
// Solidity: function bud(address ) view returns(uint256)
func (_Median *MedianSession) Bud(arg0 common.Address) (*big.Int, error) {
	return _Median.Contract.Bud(&_Median.CallOpts, arg0)
}

// Bud is a free data retrieval call binding the contract method 0x4fce7a2a.
//
// Solidity: function bud(address ) view returns(uint256)
func (_Median *MedianCallerSession) Bud(arg0 common.Address) (*big.Int, error) {
	return _Median.Contract.Bud(&_Median.CallOpts, arg0)
}

// Orcl is a free data retrieval call binding the contract method 0x020b2e32.
//
// Solidity: function orcl(address ) view returns(uint256)
func (_Median *MedianCaller) Orcl(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Median.contract.Call(opts, &out, "orcl", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Orcl is a free data retrieval call binding the contract method 0x020b2e32.
//
// Solidity: function orcl(address ) view returns(uint256)
func (_Median *MedianSession) Orcl(arg0 common.Address) (*big.Int, error) {
	return _Median.Contract.Orcl(&_Median.CallOpts, arg0)
}

// Orcl is a free data retrieval call binding the contract method 0x020b2e32.
//
// Solidity: function orcl(address ) view returns(uint256)
func (_Median *MedianCallerSession) Orcl(arg0 common.Address) (*big.Int, error) {
	return _Median.Contract.Orcl(&_Median.CallOpts, arg0)
}

// Peek is a free data retrieval call binding the contract method 0x59e02dd7.
//
// Solidity: function peek() view returns(uint256, bool)
func (_Median *MedianCaller) Peek(opts *bind.CallOpts) (*big.Int, bool, error) {
	var out []interface{}
	err := _Median.contract.Call(opts, &out, "peek")

	if err != nil {
		return *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// Peek is a free data retrieval call binding the contract method 0x59e02dd7.
//
// Solidity: function peek() view returns(uint256, bool)
func (_Median *MedianSession) Peek() (*big.Int, bool, error) {
	return _Median.Contract.Peek(&_Median.CallOpts)
}

// Peek is a free data retrieval call binding the contract method 0x59e02dd7.
//
// Solidity: function peek() view returns(uint256, bool)
func (_Median *MedianCallerSession) Peek() (*big.Int, bool, error) {
	return _Median.Contract.Peek(&_Median.CallOpts)
}

// Read is a free data retrieval call binding the contract method 0x57de26a4.
//
// Solidity: function read() view returns(uint256)
func (_Median *MedianCaller) Read(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Median.contract.Call(opts, &out, "read")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Read is a free data retrieval call binding the contract method 0x57de26a4.
//
// Solidity: function read() view returns(uint256)
func (_Median *MedianSession) Read() (*big.Int, error) {
	return _Median.Contract.Read(&_Median.CallOpts)
}

// Read is a free data retrieval call binding the contract method 0x57de26a4.
//
// Solidity: function read() view returns(uint256)
func (_Median *MedianCallerSession) Read() (*big.Int, error) {
	return _Median.Contract.Read(&_Median.CallOpts)
}

// Slot is a free data retrieval call binding the contract method 0x8d0e5a9a.
//
// Solidity: function slot(uint8 ) view returns(address)
func (_Median *MedianCaller) Slot(opts *bind.CallOpts, arg0 uint8) (common.Address, error) {
	var out []interface{}
	err := _Median.contract.Call(opts, &out, "slot", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Slot is a free data retrieval call binding the contract method 0x8d0e5a9a.
//
// Solidity: function slot(uint8 ) view returns(address)
func (_Median *MedianSession) Slot(arg0 uint8) (common.Address, error) {
	return _Median.Contract.Slot(&_Median.CallOpts, arg0)
}

// Slot is a free data retrieval call binding the contract method 0x8d0e5a9a.
//
// Solidity: function slot(uint8 ) view returns(address)
func (_Median *MedianCallerSession) Slot(arg0 uint8) (common.Address, error) {
	return _Median.Contract.Slot(&_Median.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Median *MedianCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Median.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Median *MedianSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Median.Contract.Wards(&_Median.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Median *MedianCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Median.Contract.Wards(&_Median.CallOpts, arg0)
}

// Wat is a free data retrieval call binding the contract method 0x4ca29923.
//
// Solidity: function wat() view returns(bytes32)
func (_Median *MedianCaller) Wat(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Median.contract.Call(opts, &out, "wat")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Wat is a free data retrieval call binding the contract method 0x4ca29923.
//
// Solidity: function wat() view returns(bytes32)
func (_Median *MedianSession) Wat() ([32]byte, error) {
	return _Median.Contract.Wat(&_Median.CallOpts)
}

// Wat is a free data retrieval call binding the contract method 0x4ca29923.
//
// Solidity: function wat() view returns(bytes32)
func (_Median *MedianCallerSession) Wat() ([32]byte, error) {
	return _Median.Contract.Wat(&_Median.CallOpts)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Median *MedianTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Median.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Median *MedianSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Median.Contract.Deny(&_Median.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Median *MedianTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Median.Contract.Deny(&_Median.TransactOpts, usr)
}

// Diss is a paid mutator transaction binding the contract method 0x46d4577d.
//
// Solidity: function diss(address[] a) returns()
func (_Median *MedianTransactor) Diss(opts *bind.TransactOpts, a []common.Address) (*types.Transaction, error) {
	return _Median.contract.Transact(opts, "diss", a)
}

// Diss is a paid mutator transaction binding the contract method 0x46d4577d.
//
// Solidity: function diss(address[] a) returns()
func (_Median *MedianSession) Diss(a []common.Address) (*types.Transaction, error) {
	return _Median.Contract.Diss(&_Median.TransactOpts, a)
}

// Diss is a paid mutator transaction binding the contract method 0x46d4577d.
//
// Solidity: function diss(address[] a) returns()
func (_Median *MedianTransactorSession) Diss(a []common.Address) (*types.Transaction, error) {
	return _Median.Contract.Diss(&_Median.TransactOpts, a)
}

// Diss0 is a paid mutator transaction binding the contract method 0x65c4ce7a.
//
// Solidity: function diss(address a) returns()
func (_Median *MedianTransactor) Diss0(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _Median.contract.Transact(opts, "diss0", a)
}

// Diss0 is a paid mutator transaction binding the contract method 0x65c4ce7a.
//
// Solidity: function diss(address a) returns()
func (_Median *MedianSession) Diss0(a common.Address) (*types.Transaction, error) {
	return _Median.Contract.Diss0(&_Median.TransactOpts, a)
}

// Diss0 is a paid mutator transaction binding the contract method 0x65c4ce7a.
//
// Solidity: function diss(address a) returns()
func (_Median *MedianTransactorSession) Diss0(a common.Address) (*types.Transaction, error) {
	return _Median.Contract.Diss0(&_Median.TransactOpts, a)
}

// Drop is a paid mutator transaction binding the contract method 0x8ef5eaf0.
//
// Solidity: function drop(address[] a) returns()
func (_Median *MedianTransactor) Drop(opts *bind.TransactOpts, a []common.Address) (*types.Transaction, error) {
	return _Median.contract.Transact(opts, "drop", a)
}

// Drop is a paid mutator transaction binding the contract method 0x8ef5eaf0.
//
// Solidity: function drop(address[] a) returns()
func (_Median *MedianSession) Drop(a []common.Address) (*types.Transaction, error) {
	return _Median.Contract.Drop(&_Median.TransactOpts, a)
}

// Drop is a paid mutator transaction binding the contract method 0x8ef5eaf0.
//
// Solidity: function drop(address[] a) returns()
func (_Median *MedianTransactorSession) Drop(a []common.Address) (*types.Transaction, error) {
	return _Median.Contract.Drop(&_Median.TransactOpts, a)
}

// Kiss is a paid mutator transaction binding the contract method 0x1b25b65f.
//
// Solidity: function kiss(address[] a) returns()
func (_Median *MedianTransactor) Kiss(opts *bind.TransactOpts, a []common.Address) (*types.Transaction, error) {
	return _Median.contract.Transact(opts, "kiss", a)
}

// Kiss is a paid mutator transaction binding the contract method 0x1b25b65f.
//
// Solidity: function kiss(address[] a) returns()
func (_Median *MedianSession) Kiss(a []common.Address) (*types.Transaction, error) {
	return _Median.Contract.Kiss(&_Median.TransactOpts, a)
}

// Kiss is a paid mutator transaction binding the contract method 0x1b25b65f.
//
// Solidity: function kiss(address[] a) returns()
func (_Median *MedianTransactorSession) Kiss(a []common.Address) (*types.Transaction, error) {
	return _Median.Contract.Kiss(&_Median.TransactOpts, a)
}

// Kiss0 is a paid mutator transaction binding the contract method 0xf29c29c4.
//
// Solidity: function kiss(address a) returns()
func (_Median *MedianTransactor) Kiss0(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _Median.contract.Transact(opts, "kiss0", a)
}

// Kiss0 is a paid mutator transaction binding the contract method 0xf29c29c4.
//
// Solidity: function kiss(address a) returns()
func (_Median *MedianSession) Kiss0(a common.Address) (*types.Transaction, error) {
	return _Median.Contract.Kiss0(&_Median.TransactOpts, a)
}

// Kiss0 is a paid mutator transaction binding the contract method 0xf29c29c4.
//
// Solidity: function kiss(address a) returns()
func (_Median *MedianTransactorSession) Kiss0(a common.Address) (*types.Transaction, error) {
	return _Median.Contract.Kiss0(&_Median.TransactOpts, a)
}

// Lift is a paid mutator transaction binding the contract method 0x94318106.
//
// Solidity: function lift(address[] a) returns()
func (_Median *MedianTransactor) Lift(opts *bind.TransactOpts, a []common.Address) (*types.Transaction, error) {
	return _Median.contract.Transact(opts, "lift", a)
}

// Lift is a paid mutator transaction binding the contract method 0x94318106.
//
// Solidity: function lift(address[] a) returns()
func (_Median *MedianSession) Lift(a []common.Address) (*types.Transaction, error) {
	return _Median.Contract.Lift(&_Median.TransactOpts, a)
}

// Lift is a paid mutator transaction binding the contract method 0x94318106.
//
// Solidity: function lift(address[] a) returns()
func (_Median *MedianTransactorSession) Lift(a []common.Address) (*types.Transaction, error) {
	return _Median.Contract.Lift(&_Median.TransactOpts, a)
}

// Poke is a paid mutator transaction binding the contract method 0x89bbb8b2.
//
// Solidity: function poke(uint256[] val_, uint256[] age_, uint8[] v, bytes32[] r, bytes32[] s) returns()
func (_Median *MedianTransactor) Poke(opts *bind.TransactOpts, val_ []*big.Int, age_ []*big.Int, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Median.contract.Transact(opts, "poke", val_, age_, v, r, s)
}

// Poke is a paid mutator transaction binding the contract method 0x89bbb8b2.
//
// Solidity: function poke(uint256[] val_, uint256[] age_, uint8[] v, bytes32[] r, bytes32[] s) returns()
func (_Median *MedianSession) Poke(val_ []*big.Int, age_ []*big.Int, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Median.Contract.Poke(&_Median.TransactOpts, val_, age_, v, r, s)
}

// Poke is a paid mutator transaction binding the contract method 0x89bbb8b2.
//
// Solidity: function poke(uint256[] val_, uint256[] age_, uint8[] v, bytes32[] r, bytes32[] s) returns()
func (_Median *MedianTransactorSession) Poke(val_ []*big.Int, age_ []*big.Int, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Median.Contract.Poke(&_Median.TransactOpts, val_, age_, v, r, s)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Median *MedianTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Median.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Median *MedianSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Median.Contract.Rely(&_Median.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Median *MedianTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Median.Contract.Rely(&_Median.TransactOpts, usr)
}

// SetBar is a paid mutator transaction binding the contract method 0x352d3fba.
//
// Solidity: function setBar(uint256 bar_) returns()
func (_Median *MedianTransactor) SetBar(opts *bind.TransactOpts, bar_ *big.Int) (*types.Transaction, error) {
	return _Median.contract.Transact(opts, "setBar", bar_)
}

// SetBar is a paid mutator transaction binding the contract method 0x352d3fba.
//
// Solidity: function setBar(uint256 bar_) returns()
func (_Median *MedianSession) SetBar(bar_ *big.Int) (*types.Transaction, error) {
	return _Median.Contract.SetBar(&_Median.TransactOpts, bar_)
}

// SetBar is a paid mutator transaction binding the contract method 0x352d3fba.
//
// Solidity: function setBar(uint256 bar_) returns()
func (_Median *MedianTransactorSession) SetBar(bar_ *big.Int) (*types.Transaction, error) {
	return _Median.Contract.SetBar(&_Median.TransactOpts, bar_)
}

// MedianDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the Median contract.
type MedianDenyIterator struct {
	Event *MedianDeny // Event containing the contract specifics and raw log

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
func (it *MedianDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MedianDeny)
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
		it.Event = new(MedianDeny)
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
func (it *MedianDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MedianDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MedianDeny represents a Deny event raised by the Median contract.
type MedianDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_Median *MedianFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*MedianDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Median.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &MedianDenyIterator{contract: _Median.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_Median *MedianFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *MedianDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Median.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MedianDeny)
				if err := _Median.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_Median *MedianFilterer) ParseDeny(log types.Log) (*MedianDeny, error) {
	event := new(MedianDeny)
	if err := _Median.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MedianLogMedianPriceIterator is returned from FilterLogMedianPrice and is used to iterate over the raw logs and unpacked data for LogMedianPrice events raised by the Median contract.
type MedianLogMedianPriceIterator struct {
	Event *MedianLogMedianPrice // Event containing the contract specifics and raw log

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
func (it *MedianLogMedianPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MedianLogMedianPrice)
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
		it.Event = new(MedianLogMedianPrice)
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
func (it *MedianLogMedianPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MedianLogMedianPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MedianLogMedianPrice represents a LogMedianPrice event raised by the Median contract.
type MedianLogMedianPrice struct {
	Val *big.Int
	Age *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogMedianPrice is a free log retrieval operation binding the contract event 0xb78ebc573f1f889ca9e1e0fb62c843c836f3d3a2e1f43ef62940e9b894f4ea4c.
//
// Solidity: event LogMedianPrice(uint256 val, uint256 age)
func (_Median *MedianFilterer) FilterLogMedianPrice(opts *bind.FilterOpts) (*MedianLogMedianPriceIterator, error) {

	logs, sub, err := _Median.contract.FilterLogs(opts, "LogMedianPrice")
	if err != nil {
		return nil, err
	}
	return &MedianLogMedianPriceIterator{contract: _Median.contract, event: "LogMedianPrice", logs: logs, sub: sub}, nil
}

// WatchLogMedianPrice is a free log subscription operation binding the contract event 0xb78ebc573f1f889ca9e1e0fb62c843c836f3d3a2e1f43ef62940e9b894f4ea4c.
//
// Solidity: event LogMedianPrice(uint256 val, uint256 age)
func (_Median *MedianFilterer) WatchLogMedianPrice(opts *bind.WatchOpts, sink chan<- *MedianLogMedianPrice) (event.Subscription, error) {

	logs, sub, err := _Median.contract.WatchLogs(opts, "LogMedianPrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MedianLogMedianPrice)
				if err := _Median.contract.UnpackLog(event, "LogMedianPrice", log); err != nil {
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

// ParseLogMedianPrice is a log parse operation binding the contract event 0xb78ebc573f1f889ca9e1e0fb62c843c836f3d3a2e1f43ef62940e9b894f4ea4c.
//
// Solidity: event LogMedianPrice(uint256 val, uint256 age)
func (_Median *MedianFilterer) ParseLogMedianPrice(log types.Log) (*MedianLogMedianPrice, error) {
	event := new(MedianLogMedianPrice)
	if err := _Median.contract.UnpackLog(event, "LogMedianPrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MedianRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the Median contract.
type MedianRelyIterator struct {
	Event *MedianRely // Event containing the contract specifics and raw log

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
func (it *MedianRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MedianRely)
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
		it.Event = new(MedianRely)
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
func (it *MedianRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MedianRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MedianRely represents a Rely event raised by the Median contract.
type MedianRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_Median *MedianFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*MedianRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Median.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &MedianRelyIterator{contract: _Median.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_Median *MedianFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *MedianRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Median.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MedianRely)
				if err := _Median.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_Median *MedianFilterer) ParseRely(log types.Log) (*MedianRely, error) {
	event := new(MedianRely)
	if err := _Median.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
