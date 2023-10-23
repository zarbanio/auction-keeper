// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package osm

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

// OsmMetaData contains all meta data concerning the Osm contract.
var OsmMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"val\",\"type\":\"bytes32\"}],\"name\":\"LogValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"bud\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src_\",\"type\":\"address\"}],\"name\":\"change\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"a\",\"type\":\"address[]\"}],\"name\":\"diss\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"diss\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hop\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"a\",\"type\":\"address[]\"}],\"name\":\"kiss\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"kiss\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pass\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"ok\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"peek\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"peep\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"read\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"src\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"ts\",\"type\":\"uint16\"}],\"name\":\"step\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"void\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zzz\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// OsmABI is the input ABI used to generate the binding from.
// Deprecated: Use OsmMetaData.ABI instead.
var OsmABI = OsmMetaData.ABI

// Osm is an auto generated Go binding around an Ethereum contract.
type Osm struct {
	OsmCaller     // Read-only binding to the contract
	OsmTransactor // Write-only binding to the contract
	OsmFilterer   // Log filterer for contract events
}

// OsmCaller is an auto generated read-only Go binding around an Ethereum contract.
type OsmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OsmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OsmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OsmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OsmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OsmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OsmSession struct {
	Contract     *Osm              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OsmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OsmCallerSession struct {
	Contract *OsmCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OsmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OsmTransactorSession struct {
	Contract     *OsmTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OsmRaw is an auto generated low-level Go binding around an Ethereum contract.
type OsmRaw struct {
	Contract *Osm // Generic contract binding to access the raw methods on
}

// OsmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OsmCallerRaw struct {
	Contract *OsmCaller // Generic read-only contract binding to access the raw methods on
}

// OsmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OsmTransactorRaw struct {
	Contract *OsmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOsm creates a new instance of Osm, bound to a specific deployed contract.
func NewOsm(address common.Address, backend bind.ContractBackend) (*Osm, error) {
	contract, err := bindOsm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Osm{OsmCaller: OsmCaller{contract: contract}, OsmTransactor: OsmTransactor{contract: contract}, OsmFilterer: OsmFilterer{contract: contract}}, nil
}

// NewOsmCaller creates a new read-only instance of Osm, bound to a specific deployed contract.
func NewOsmCaller(address common.Address, caller bind.ContractCaller) (*OsmCaller, error) {
	contract, err := bindOsm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OsmCaller{contract: contract}, nil
}

// NewOsmTransactor creates a new write-only instance of Osm, bound to a specific deployed contract.
func NewOsmTransactor(address common.Address, transactor bind.ContractTransactor) (*OsmTransactor, error) {
	contract, err := bindOsm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OsmTransactor{contract: contract}, nil
}

// NewOsmFilterer creates a new log filterer instance of Osm, bound to a specific deployed contract.
func NewOsmFilterer(address common.Address, filterer bind.ContractFilterer) (*OsmFilterer, error) {
	contract, err := bindOsm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OsmFilterer{contract: contract}, nil
}

// bindOsm binds a generic wrapper to an already deployed contract.
func bindOsm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OsmMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Osm *OsmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Osm.Contract.OsmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Osm *OsmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Osm.Contract.OsmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Osm *OsmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Osm.Contract.OsmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Osm *OsmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Osm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Osm *OsmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Osm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Osm *OsmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Osm.Contract.contract.Transact(opts, method, params...)
}

// Bud is a free data retrieval call binding the contract method 0x4fce7a2a.
//
// Solidity: function bud(address ) view returns(uint256)
func (_Osm *OsmCaller) Bud(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Osm.contract.Call(opts, &out, "bud", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Bud is a free data retrieval call binding the contract method 0x4fce7a2a.
//
// Solidity: function bud(address ) view returns(uint256)
func (_Osm *OsmSession) Bud(arg0 common.Address) (*big.Int, error) {
	return _Osm.Contract.Bud(&_Osm.CallOpts, arg0)
}

// Bud is a free data retrieval call binding the contract method 0x4fce7a2a.
//
// Solidity: function bud(address ) view returns(uint256)
func (_Osm *OsmCallerSession) Bud(arg0 common.Address) (*big.Int, error) {
	return _Osm.Contract.Bud(&_Osm.CallOpts, arg0)
}

// Hop is a free data retrieval call binding the contract method 0xb0b8579b.
//
// Solidity: function hop() view returns(uint16)
func (_Osm *OsmCaller) Hop(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Osm.contract.Call(opts, &out, "hop")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Hop is a free data retrieval call binding the contract method 0xb0b8579b.
//
// Solidity: function hop() view returns(uint16)
func (_Osm *OsmSession) Hop() (uint16, error) {
	return _Osm.Contract.Hop(&_Osm.CallOpts)
}

// Hop is a free data retrieval call binding the contract method 0xb0b8579b.
//
// Solidity: function hop() view returns(uint16)
func (_Osm *OsmCallerSession) Hop() (uint16, error) {
	return _Osm.Contract.Hop(&_Osm.CallOpts)
}

// Pass is a free data retrieval call binding the contract method 0xa7a1ed72.
//
// Solidity: function pass() view returns(bool ok)
func (_Osm *OsmCaller) Pass(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Osm.contract.Call(opts, &out, "pass")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Pass is a free data retrieval call binding the contract method 0xa7a1ed72.
//
// Solidity: function pass() view returns(bool ok)
func (_Osm *OsmSession) Pass() (bool, error) {
	return _Osm.Contract.Pass(&_Osm.CallOpts)
}

// Pass is a free data retrieval call binding the contract method 0xa7a1ed72.
//
// Solidity: function pass() view returns(bool ok)
func (_Osm *OsmCallerSession) Pass() (bool, error) {
	return _Osm.Contract.Pass(&_Osm.CallOpts)
}

// Peek is a free data retrieval call binding the contract method 0x59e02dd7.
//
// Solidity: function peek() view returns(bytes32, bool)
func (_Osm *OsmCaller) Peek(opts *bind.CallOpts) ([32]byte, bool, error) {
	var out []interface{}
	err := _Osm.contract.Call(opts, &out, "peek")

	if err != nil {
		return *new([32]byte), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// Peek is a free data retrieval call binding the contract method 0x59e02dd7.
//
// Solidity: function peek() view returns(bytes32, bool)
func (_Osm *OsmSession) Peek() ([32]byte, bool, error) {
	return _Osm.Contract.Peek(&_Osm.CallOpts)
}

// Peek is a free data retrieval call binding the contract method 0x59e02dd7.
//
// Solidity: function peek() view returns(bytes32, bool)
func (_Osm *OsmCallerSession) Peek() ([32]byte, bool, error) {
	return _Osm.Contract.Peek(&_Osm.CallOpts)
}

// Peep is a free data retrieval call binding the contract method 0x0e5a6c70.
//
// Solidity: function peep() view returns(bytes32, bool)
func (_Osm *OsmCaller) Peep(opts *bind.CallOpts) ([32]byte, bool, error) {
	var out []interface{}
	err := _Osm.contract.Call(opts, &out, "peep")

	if err != nil {
		return *new([32]byte), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// Peep is a free data retrieval call binding the contract method 0x0e5a6c70.
//
// Solidity: function peep() view returns(bytes32, bool)
func (_Osm *OsmSession) Peep() ([32]byte, bool, error) {
	return _Osm.Contract.Peep(&_Osm.CallOpts)
}

// Peep is a free data retrieval call binding the contract method 0x0e5a6c70.
//
// Solidity: function peep() view returns(bytes32, bool)
func (_Osm *OsmCallerSession) Peep() ([32]byte, bool, error) {
	return _Osm.Contract.Peep(&_Osm.CallOpts)
}

// Read is a free data retrieval call binding the contract method 0x57de26a4.
//
// Solidity: function read() view returns(bytes32)
func (_Osm *OsmCaller) Read(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Osm.contract.Call(opts, &out, "read")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Read is a free data retrieval call binding the contract method 0x57de26a4.
//
// Solidity: function read() view returns(bytes32)
func (_Osm *OsmSession) Read() ([32]byte, error) {
	return _Osm.Contract.Read(&_Osm.CallOpts)
}

// Read is a free data retrieval call binding the contract method 0x57de26a4.
//
// Solidity: function read() view returns(bytes32)
func (_Osm *OsmCallerSession) Read() ([32]byte, error) {
	return _Osm.Contract.Read(&_Osm.CallOpts)
}

// Src is a free data retrieval call binding the contract method 0x2e7dc6af.
//
// Solidity: function src() view returns(address)
func (_Osm *OsmCaller) Src(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Osm.contract.Call(opts, &out, "src")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Src is a free data retrieval call binding the contract method 0x2e7dc6af.
//
// Solidity: function src() view returns(address)
func (_Osm *OsmSession) Src() (common.Address, error) {
	return _Osm.Contract.Src(&_Osm.CallOpts)
}

// Src is a free data retrieval call binding the contract method 0x2e7dc6af.
//
// Solidity: function src() view returns(address)
func (_Osm *OsmCallerSession) Src() (common.Address, error) {
	return _Osm.Contract.Src(&_Osm.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() view returns(uint256)
func (_Osm *OsmCaller) Stopped(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Osm.contract.Call(opts, &out, "stopped")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() view returns(uint256)
func (_Osm *OsmSession) Stopped() (*big.Int, error) {
	return _Osm.Contract.Stopped(&_Osm.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() view returns(uint256)
func (_Osm *OsmCallerSession) Stopped() (*big.Int, error) {
	return _Osm.Contract.Stopped(&_Osm.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Osm *OsmCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Osm.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Osm *OsmSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Osm.Contract.Wards(&_Osm.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Osm *OsmCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Osm.Contract.Wards(&_Osm.CallOpts, arg0)
}

// Zzz is a free data retrieval call binding the contract method 0xa4dff0a2.
//
// Solidity: function zzz() view returns(uint64)
func (_Osm *OsmCaller) Zzz(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Osm.contract.Call(opts, &out, "zzz")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Zzz is a free data retrieval call binding the contract method 0xa4dff0a2.
//
// Solidity: function zzz() view returns(uint64)
func (_Osm *OsmSession) Zzz() (uint64, error) {
	return _Osm.Contract.Zzz(&_Osm.CallOpts)
}

// Zzz is a free data retrieval call binding the contract method 0xa4dff0a2.
//
// Solidity: function zzz() view returns(uint64)
func (_Osm *OsmCallerSession) Zzz() (uint64, error) {
	return _Osm.Contract.Zzz(&_Osm.CallOpts)
}

// Change is a paid mutator transaction binding the contract method 0x1e77933e.
//
// Solidity: function change(address src_) returns()
func (_Osm *OsmTransactor) Change(opts *bind.TransactOpts, src_ common.Address) (*types.Transaction, error) {
	return _Osm.contract.Transact(opts, "change", src_)
}

// Change is a paid mutator transaction binding the contract method 0x1e77933e.
//
// Solidity: function change(address src_) returns()
func (_Osm *OsmSession) Change(src_ common.Address) (*types.Transaction, error) {
	return _Osm.Contract.Change(&_Osm.TransactOpts, src_)
}

// Change is a paid mutator transaction binding the contract method 0x1e77933e.
//
// Solidity: function change(address src_) returns()
func (_Osm *OsmTransactorSession) Change(src_ common.Address) (*types.Transaction, error) {
	return _Osm.Contract.Change(&_Osm.TransactOpts, src_)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Osm *OsmTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Osm.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Osm *OsmSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Osm.Contract.Deny(&_Osm.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Osm *OsmTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Osm.Contract.Deny(&_Osm.TransactOpts, usr)
}

// Diss is a paid mutator transaction binding the contract method 0x46d4577d.
//
// Solidity: function diss(address[] a) returns()
func (_Osm *OsmTransactor) Diss(opts *bind.TransactOpts, a []common.Address) (*types.Transaction, error) {
	return _Osm.contract.Transact(opts, "diss", a)
}

// Diss is a paid mutator transaction binding the contract method 0x46d4577d.
//
// Solidity: function diss(address[] a) returns()
func (_Osm *OsmSession) Diss(a []common.Address) (*types.Transaction, error) {
	return _Osm.Contract.Diss(&_Osm.TransactOpts, a)
}

// Diss is a paid mutator transaction binding the contract method 0x46d4577d.
//
// Solidity: function diss(address[] a) returns()
func (_Osm *OsmTransactorSession) Diss(a []common.Address) (*types.Transaction, error) {
	return _Osm.Contract.Diss(&_Osm.TransactOpts, a)
}

// Diss0 is a paid mutator transaction binding the contract method 0x65c4ce7a.
//
// Solidity: function diss(address a) returns()
func (_Osm *OsmTransactor) Diss0(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _Osm.contract.Transact(opts, "diss0", a)
}

// Diss0 is a paid mutator transaction binding the contract method 0x65c4ce7a.
//
// Solidity: function diss(address a) returns()
func (_Osm *OsmSession) Diss0(a common.Address) (*types.Transaction, error) {
	return _Osm.Contract.Diss0(&_Osm.TransactOpts, a)
}

// Diss0 is a paid mutator transaction binding the contract method 0x65c4ce7a.
//
// Solidity: function diss(address a) returns()
func (_Osm *OsmTransactorSession) Diss0(a common.Address) (*types.Transaction, error) {
	return _Osm.Contract.Diss0(&_Osm.TransactOpts, a)
}

// Kiss is a paid mutator transaction binding the contract method 0x1b25b65f.
//
// Solidity: function kiss(address[] a) returns()
func (_Osm *OsmTransactor) Kiss(opts *bind.TransactOpts, a []common.Address) (*types.Transaction, error) {
	return _Osm.contract.Transact(opts, "kiss", a)
}

// Kiss is a paid mutator transaction binding the contract method 0x1b25b65f.
//
// Solidity: function kiss(address[] a) returns()
func (_Osm *OsmSession) Kiss(a []common.Address) (*types.Transaction, error) {
	return _Osm.Contract.Kiss(&_Osm.TransactOpts, a)
}

// Kiss is a paid mutator transaction binding the contract method 0x1b25b65f.
//
// Solidity: function kiss(address[] a) returns()
func (_Osm *OsmTransactorSession) Kiss(a []common.Address) (*types.Transaction, error) {
	return _Osm.Contract.Kiss(&_Osm.TransactOpts, a)
}

// Kiss0 is a paid mutator transaction binding the contract method 0xf29c29c4.
//
// Solidity: function kiss(address a) returns()
func (_Osm *OsmTransactor) Kiss0(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _Osm.contract.Transact(opts, "kiss0", a)
}

// Kiss0 is a paid mutator transaction binding the contract method 0xf29c29c4.
//
// Solidity: function kiss(address a) returns()
func (_Osm *OsmSession) Kiss0(a common.Address) (*types.Transaction, error) {
	return _Osm.Contract.Kiss0(&_Osm.TransactOpts, a)
}

// Kiss0 is a paid mutator transaction binding the contract method 0xf29c29c4.
//
// Solidity: function kiss(address a) returns()
func (_Osm *OsmTransactorSession) Kiss0(a common.Address) (*types.Transaction, error) {
	return _Osm.Contract.Kiss0(&_Osm.TransactOpts, a)
}

// Poke is a paid mutator transaction binding the contract method 0x18178358.
//
// Solidity: function poke() returns()
func (_Osm *OsmTransactor) Poke(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Osm.contract.Transact(opts, "poke")
}

// Poke is a paid mutator transaction binding the contract method 0x18178358.
//
// Solidity: function poke() returns()
func (_Osm *OsmSession) Poke() (*types.Transaction, error) {
	return _Osm.Contract.Poke(&_Osm.TransactOpts)
}

// Poke is a paid mutator transaction binding the contract method 0x18178358.
//
// Solidity: function poke() returns()
func (_Osm *OsmTransactorSession) Poke() (*types.Transaction, error) {
	return _Osm.Contract.Poke(&_Osm.TransactOpts)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Osm *OsmTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Osm.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Osm *OsmSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Osm.Contract.Rely(&_Osm.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Osm *OsmTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Osm.Contract.Rely(&_Osm.TransactOpts, usr)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Osm *OsmTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Osm.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Osm *OsmSession) Start() (*types.Transaction, error) {
	return _Osm.Contract.Start(&_Osm.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Osm *OsmTransactorSession) Start() (*types.Transaction, error) {
	return _Osm.Contract.Start(&_Osm.TransactOpts)
}

// Step is a paid mutator transaction binding the contract method 0xe38e2cfb.
//
// Solidity: function step(uint16 ts) returns()
func (_Osm *OsmTransactor) Step(opts *bind.TransactOpts, ts uint16) (*types.Transaction, error) {
	return _Osm.contract.Transact(opts, "step", ts)
}

// Step is a paid mutator transaction binding the contract method 0xe38e2cfb.
//
// Solidity: function step(uint16 ts) returns()
func (_Osm *OsmSession) Step(ts uint16) (*types.Transaction, error) {
	return _Osm.Contract.Step(&_Osm.TransactOpts, ts)
}

// Step is a paid mutator transaction binding the contract method 0xe38e2cfb.
//
// Solidity: function step(uint16 ts) returns()
func (_Osm *OsmTransactorSession) Step(ts uint16) (*types.Transaction, error) {
	return _Osm.Contract.Step(&_Osm.TransactOpts, ts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Osm *OsmTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Osm.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Osm *OsmSession) Stop() (*types.Transaction, error) {
	return _Osm.Contract.Stop(&_Osm.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Osm *OsmTransactorSession) Stop() (*types.Transaction, error) {
	return _Osm.Contract.Stop(&_Osm.TransactOpts)
}

// Void is a paid mutator transaction binding the contract method 0xac4c25b2.
//
// Solidity: function void() returns()
func (_Osm *OsmTransactor) Void(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Osm.contract.Transact(opts, "void")
}

// Void is a paid mutator transaction binding the contract method 0xac4c25b2.
//
// Solidity: function void() returns()
func (_Osm *OsmSession) Void() (*types.Transaction, error) {
	return _Osm.Contract.Void(&_Osm.TransactOpts)
}

// Void is a paid mutator transaction binding the contract method 0xac4c25b2.
//
// Solidity: function void() returns()
func (_Osm *OsmTransactorSession) Void() (*types.Transaction, error) {
	return _Osm.Contract.Void(&_Osm.TransactOpts)
}

// OsmDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the Osm contract.
type OsmDenyIterator struct {
	Event *OsmDeny // Event containing the contract specifics and raw log

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
func (it *OsmDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OsmDeny)
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
		it.Event = new(OsmDeny)
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
func (it *OsmDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OsmDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OsmDeny represents a Deny event raised by the Osm contract.
type OsmDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_Osm *OsmFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*OsmDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Osm.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &OsmDenyIterator{contract: _Osm.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_Osm *OsmFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *OsmDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Osm.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OsmDeny)
				if err := _Osm.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_Osm *OsmFilterer) ParseDeny(log types.Log) (*OsmDeny, error) {
	event := new(OsmDeny)
	if err := _Osm.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OsmLogValueIterator is returned from FilterLogValue and is used to iterate over the raw logs and unpacked data for LogValue events raised by the Osm contract.
type OsmLogValueIterator struct {
	Event *OsmLogValue // Event containing the contract specifics and raw log

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
func (it *OsmLogValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OsmLogValue)
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
		it.Event = new(OsmLogValue)
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
func (it *OsmLogValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OsmLogValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OsmLogValue represents a LogValue event raised by the Osm contract.
type OsmLogValue struct {
	Val [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogValue is a free log retrieval operation binding the contract event 0x296ba4ca62c6c21c95e828080cb8aec7481b71390585605300a8a76f9e95b527.
//
// Solidity: event LogValue(bytes32 val)
func (_Osm *OsmFilterer) FilterLogValue(opts *bind.FilterOpts) (*OsmLogValueIterator, error) {

	logs, sub, err := _Osm.contract.FilterLogs(opts, "LogValue")
	if err != nil {
		return nil, err
	}
	return &OsmLogValueIterator{contract: _Osm.contract, event: "LogValue", logs: logs, sub: sub}, nil
}

// WatchLogValue is a free log subscription operation binding the contract event 0x296ba4ca62c6c21c95e828080cb8aec7481b71390585605300a8a76f9e95b527.
//
// Solidity: event LogValue(bytes32 val)
func (_Osm *OsmFilterer) WatchLogValue(opts *bind.WatchOpts, sink chan<- *OsmLogValue) (event.Subscription, error) {

	logs, sub, err := _Osm.contract.WatchLogs(opts, "LogValue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OsmLogValue)
				if err := _Osm.contract.UnpackLog(event, "LogValue", log); err != nil {
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

// ParseLogValue is a log parse operation binding the contract event 0x296ba4ca62c6c21c95e828080cb8aec7481b71390585605300a8a76f9e95b527.
//
// Solidity: event LogValue(bytes32 val)
func (_Osm *OsmFilterer) ParseLogValue(log types.Log) (*OsmLogValue, error) {
	event := new(OsmLogValue)
	if err := _Osm.contract.UnpackLog(event, "LogValue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OsmRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the Osm contract.
type OsmRelyIterator struct {
	Event *OsmRely // Event containing the contract specifics and raw log

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
func (it *OsmRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OsmRely)
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
		it.Event = new(OsmRely)
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
func (it *OsmRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OsmRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OsmRely represents a Rely event raised by the Osm contract.
type OsmRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_Osm *OsmFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*OsmRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Osm.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &OsmRelyIterator{contract: _Osm.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_Osm *OsmFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *OsmRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Osm.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OsmRely)
				if err := _Osm.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_Osm *OsmFilterer) ParseRely(log types.Log) (*OsmRely, error) {
	event := new(OsmRely)
	if err := _Osm.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
