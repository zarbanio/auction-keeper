// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vow

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

// VowMetaData contains all meta data concerning the Vow contract.
var VowMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"flopper_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Cage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tab\",\"type\":\"uint256\"}],\"name\":\"Fess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"era\",\"type\":\"uint256\"}],\"name\":\"Flog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Flop\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"Heal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"Kiss\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Ash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Sin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tab\",\"type\":\"uint256\"}],\"name\":\"fess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"era\",\"type\":\"uint256\"}],\"name\":\"flog\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flopper\",\"outputs\":[{\"internalType\":\"contractFlopLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"heal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hump\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"kiss\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"live\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sump\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wait\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VowABI is the input ABI used to generate the binding from.
// Deprecated: Use VowMetaData.ABI instead.
var VowABI = VowMetaData.ABI

// Vow is an auto generated Go binding around an Ethereum contract.
type Vow struct {
	VowCaller     // Read-only binding to the contract
	VowTransactor // Write-only binding to the contract
	VowFilterer   // Log filterer for contract events
}

// VowCaller is an auto generated read-only Go binding around an Ethereum contract.
type VowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VowSession struct {
	Contract     *Vow              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VowCallerSession struct {
	Contract *VowCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VowTransactorSession struct {
	Contract     *VowTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VowRaw is an auto generated low-level Go binding around an Ethereum contract.
type VowRaw struct {
	Contract *Vow // Generic contract binding to access the raw methods on
}

// VowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VowCallerRaw struct {
	Contract *VowCaller // Generic read-only contract binding to access the raw methods on
}

// VowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VowTransactorRaw struct {
	Contract *VowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVow creates a new instance of Vow, bound to a specific deployed contract.
func NewVow(address common.Address, backend bind.ContractBackend) (*Vow, error) {
	contract, err := bindVow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vow{VowCaller: VowCaller{contract: contract}, VowTransactor: VowTransactor{contract: contract}, VowFilterer: VowFilterer{contract: contract}}, nil
}

// NewVowCaller creates a new read-only instance of Vow, bound to a specific deployed contract.
func NewVowCaller(address common.Address, caller bind.ContractCaller) (*VowCaller, error) {
	contract, err := bindVow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VowCaller{contract: contract}, nil
}

// NewVowTransactor creates a new write-only instance of Vow, bound to a specific deployed contract.
func NewVowTransactor(address common.Address, transactor bind.ContractTransactor) (*VowTransactor, error) {
	contract, err := bindVow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VowTransactor{contract: contract}, nil
}

// NewVowFilterer creates a new log filterer instance of Vow, bound to a specific deployed contract.
func NewVowFilterer(address common.Address, filterer bind.ContractFilterer) (*VowFilterer, error) {
	contract, err := bindVow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VowFilterer{contract: contract}, nil
}

// bindVow binds a generic wrapper to an already deployed contract.
func bindVow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VowMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vow *VowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vow.Contract.VowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vow *VowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vow.Contract.VowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vow *VowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vow.Contract.VowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vow *VowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vow *VowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vow *VowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vow.Contract.contract.Transact(opts, method, params...)
}

// Ash is a free data retrieval call binding the contract method 0x2a1d2b3c.
//
// Solidity: function Ash() view returns(uint256)
func (_Vow *VowCaller) Ash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vow.contract.Call(opts, &out, "Ash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Ash is a free data retrieval call binding the contract method 0x2a1d2b3c.
//
// Solidity: function Ash() view returns(uint256)
func (_Vow *VowSession) Ash() (*big.Int, error) {
	return _Vow.Contract.Ash(&_Vow.CallOpts)
}

// Ash is a free data retrieval call binding the contract method 0x2a1d2b3c.
//
// Solidity: function Ash() view returns(uint256)
func (_Vow *VowCallerSession) Ash() (*big.Int, error) {
	return _Vow.Contract.Ash(&_Vow.CallOpts)
}

// QueuedDebt is a free data retrieval call binding the contract method 0xd0adc35f.
//
// Solidity: function Sin() view returns(uint256)
func (_Vow *VowCaller) QueuedDebt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vow.contract.Call(opts, &out, "Sin")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueuedDebt is a free data retrieval call binding the contract method 0xd0adc35f.
//
// Solidity: function Sin() view returns(uint256)
func (_Vow *VowSession) QueuedDebt() (*big.Int, error) {
	return _Vow.Contract.QueuedDebt(&_Vow.CallOpts)
}

// QueuedDebt is a free data retrieval call binding the contract method 0xd0adc35f.
//
// Solidity: function Sin() view returns(uint256)
func (_Vow *VowCallerSession) QueuedDebt() (*big.Int, error) {
	return _Vow.Contract.QueuedDebt(&_Vow.CallOpts)
}

// Flopper is a free data retrieval call binding the contract method 0x4081d73a.
//
// Solidity: function flopper() view returns(address)
func (_Vow *VowCaller) Flopper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vow.contract.Call(opts, &out, "flopper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Flopper is a free data retrieval call binding the contract method 0x4081d73a.
//
// Solidity: function flopper() view returns(address)
func (_Vow *VowSession) Flopper() (common.Address, error) {
	return _Vow.Contract.Flopper(&_Vow.CallOpts)
}

// Flopper is a free data retrieval call binding the contract method 0x4081d73a.
//
// Solidity: function flopper() view returns(address)
func (_Vow *VowCallerSession) Flopper() (common.Address, error) {
	return _Vow.Contract.Flopper(&_Vow.CallOpts)
}

// Hump is a free data retrieval call binding the contract method 0x1b8e8cfa.
//
// Solidity: function hump() view returns(uint256)
func (_Vow *VowCaller) Hump(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vow.contract.Call(opts, &out, "hump")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Hump is a free data retrieval call binding the contract method 0x1b8e8cfa.
//
// Solidity: function hump() view returns(uint256)
func (_Vow *VowSession) Hump() (*big.Int, error) {
	return _Vow.Contract.Hump(&_Vow.CallOpts)
}

// Hump is a free data retrieval call binding the contract method 0x1b8e8cfa.
//
// Solidity: function hump() view returns(uint256)
func (_Vow *VowCallerSession) Hump() (*big.Int, error) {
	return _Vow.Contract.Hump(&_Vow.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_Vow *VowCaller) Live(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vow.contract.Call(opts, &out, "live")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_Vow *VowSession) Live() (*big.Int, error) {
	return _Vow.Contract.Live(&_Vow.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_Vow *VowCallerSession) Live() (*big.Int, error) {
	return _Vow.Contract.Live(&_Vow.CallOpts)
}

// Sin is a free data retrieval call binding the contract method 0xcb5cc109.
//
// Solidity: function sin(uint256 ) view returns(uint256)
func (_Vow *VowCaller) Sin(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vow.contract.Call(opts, &out, "sin", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sin is a free data retrieval call binding the contract method 0xcb5cc109.
//
// Solidity: function sin(uint256 ) view returns(uint256)
func (_Vow *VowSession) Sin(arg0 *big.Int) (*big.Int, error) {
	return _Vow.Contract.Sin(&_Vow.CallOpts, arg0)
}

// Sin is a free data retrieval call binding the contract method 0xcb5cc109.
//
// Solidity: function sin(uint256 ) view returns(uint256)
func (_Vow *VowCallerSession) Sin(arg0 *big.Int) (*big.Int, error) {
	return _Vow.Contract.Sin(&_Vow.CallOpts, arg0)
}

// Sump is a free data retrieval call binding the contract method 0xc349d362.
//
// Solidity: function sump() view returns(uint256)
func (_Vow *VowCaller) Sump(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vow.contract.Call(opts, &out, "sump")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sump is a free data retrieval call binding the contract method 0xc349d362.
//
// Solidity: function sump() view returns(uint256)
func (_Vow *VowSession) Sump() (*big.Int, error) {
	return _Vow.Contract.Sump(&_Vow.CallOpts)
}

// Sump is a free data retrieval call binding the contract method 0xc349d362.
//
// Solidity: function sump() view returns(uint256)
func (_Vow *VowCallerSession) Sump() (*big.Int, error) {
	return _Vow.Contract.Sump(&_Vow.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Vow *VowCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vow.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Vow *VowSession) Vat() (common.Address, error) {
	return _Vow.Contract.Vat(&_Vow.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Vow *VowCallerSession) Vat() (common.Address, error) {
	return _Vow.Contract.Vat(&_Vow.CallOpts)
}

// Wait is a free data retrieval call binding the contract method 0x64bd7013.
//
// Solidity: function wait() view returns(uint256)
func (_Vow *VowCaller) Wait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vow.contract.Call(opts, &out, "wait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wait is a free data retrieval call binding the contract method 0x64bd7013.
//
// Solidity: function wait() view returns(uint256)
func (_Vow *VowSession) Wait() (*big.Int, error) {
	return _Vow.Contract.Wait(&_Vow.CallOpts)
}

// Wait is a free data retrieval call binding the contract method 0x64bd7013.
//
// Solidity: function wait() view returns(uint256)
func (_Vow *VowCallerSession) Wait() (*big.Int, error) {
	return _Vow.Contract.Wait(&_Vow.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Vow *VowCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vow.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Vow *VowSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Vow.Contract.Wards(&_Vow.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Vow *VowCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Vow.Contract.Wards(&_Vow.CallOpts, arg0)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Vow *VowTransactor) Cage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vow.contract.Transact(opts, "cage")
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Vow *VowSession) Cage() (*types.Transaction, error) {
	return _Vow.Contract.Cage(&_Vow.TransactOpts)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Vow *VowTransactorSession) Cage() (*types.Transaction, error) {
	return _Vow.Contract.Cage(&_Vow.TransactOpts)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Vow *VowTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Vow.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Vow *VowSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Vow.Contract.Deny(&_Vow.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Vow *VowTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Vow.Contract.Deny(&_Vow.TransactOpts, usr)
}

// Fess is a paid mutator transaction binding the contract method 0x697efb78.
//
// Solidity: function fess(uint256 tab) returns()
func (_Vow *VowTransactor) Fess(opts *bind.TransactOpts, tab *big.Int) (*types.Transaction, error) {
	return _Vow.contract.Transact(opts, "fess", tab)
}

// Fess is a paid mutator transaction binding the contract method 0x697efb78.
//
// Solidity: function fess(uint256 tab) returns()
func (_Vow *VowSession) Fess(tab *big.Int) (*types.Transaction, error) {
	return _Vow.Contract.Fess(&_Vow.TransactOpts, tab)
}

// Fess is a paid mutator transaction binding the contract method 0x697efb78.
//
// Solidity: function fess(uint256 tab) returns()
func (_Vow *VowTransactorSession) Fess(tab *big.Int) (*types.Transaction, error) {
	return _Vow.Contract.Fess(&_Vow.TransactOpts, tab)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Vow *VowTransactor) File(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Vow.contract.Transact(opts, "file", what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Vow *VowSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Vow.Contract.File(&_Vow.TransactOpts, what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Vow *VowTransactorSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Vow.Contract.File(&_Vow.TransactOpts, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_Vow *VowTransactor) File0(opts *bind.TransactOpts, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _Vow.contract.Transact(opts, "file0", what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_Vow *VowSession) File0(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _Vow.Contract.File0(&_Vow.TransactOpts, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_Vow *VowTransactorSession) File0(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _Vow.Contract.File0(&_Vow.TransactOpts, what, data)
}

// Flog is a paid mutator transaction binding the contract method 0xd7ee674b.
//
// Solidity: function flog(uint256 era) returns()
func (_Vow *VowTransactor) Flog(opts *bind.TransactOpts, era *big.Int) (*types.Transaction, error) {
	return _Vow.contract.Transact(opts, "flog", era)
}

// Flog is a paid mutator transaction binding the contract method 0xd7ee674b.
//
// Solidity: function flog(uint256 era) returns()
func (_Vow *VowSession) Flog(era *big.Int) (*types.Transaction, error) {
	return _Vow.Contract.Flog(&_Vow.TransactOpts, era)
}

// Flog is a paid mutator transaction binding the contract method 0xd7ee674b.
//
// Solidity: function flog(uint256 era) returns()
func (_Vow *VowTransactorSession) Flog(era *big.Int) (*types.Transaction, error) {
	return _Vow.Contract.Flog(&_Vow.TransactOpts, era)
}

// Flop is a paid mutator transaction binding the contract method 0xbbbb0d7b.
//
// Solidity: function flop() returns()
func (_Vow *VowTransactor) Flop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vow.contract.Transact(opts, "flop")
}

// Flop is a paid mutator transaction binding the contract method 0xbbbb0d7b.
//
// Solidity: function flop() returns()
func (_Vow *VowSession) Flop() (*types.Transaction, error) {
	return _Vow.Contract.Flop(&_Vow.TransactOpts)
}

// Flop is a paid mutator transaction binding the contract method 0xbbbb0d7b.
//
// Solidity: function flop() returns()
func (_Vow *VowTransactorSession) Flop() (*types.Transaction, error) {
	return _Vow.Contract.Flop(&_Vow.TransactOpts)
}

// Heal is a paid mutator transaction binding the contract method 0xf37ac61c.
//
// Solidity: function heal(uint256 rad) returns()
func (_Vow *VowTransactor) Heal(opts *bind.TransactOpts, rad *big.Int) (*types.Transaction, error) {
	return _Vow.contract.Transact(opts, "heal", rad)
}

// Heal is a paid mutator transaction binding the contract method 0xf37ac61c.
//
// Solidity: function heal(uint256 rad) returns()
func (_Vow *VowSession) Heal(rad *big.Int) (*types.Transaction, error) {
	return _Vow.Contract.Heal(&_Vow.TransactOpts, rad)
}

// Heal is a paid mutator transaction binding the contract method 0xf37ac61c.
//
// Solidity: function heal(uint256 rad) returns()
func (_Vow *VowTransactorSession) Heal(rad *big.Int) (*types.Transaction, error) {
	return _Vow.Contract.Heal(&_Vow.TransactOpts, rad)
}

// Kiss is a paid mutator transaction binding the contract method 0x2506855a.
//
// Solidity: function kiss(uint256 rad) returns()
func (_Vow *VowTransactor) Kiss(opts *bind.TransactOpts, rad *big.Int) (*types.Transaction, error) {
	return _Vow.contract.Transact(opts, "kiss", rad)
}

// Kiss is a paid mutator transaction binding the contract method 0x2506855a.
//
// Solidity: function kiss(uint256 rad) returns()
func (_Vow *VowSession) Kiss(rad *big.Int) (*types.Transaction, error) {
	return _Vow.Contract.Kiss(&_Vow.TransactOpts, rad)
}

// Kiss is a paid mutator transaction binding the contract method 0x2506855a.
//
// Solidity: function kiss(uint256 rad) returns()
func (_Vow *VowTransactorSession) Kiss(rad *big.Int) (*types.Transaction, error) {
	return _Vow.Contract.Kiss(&_Vow.TransactOpts, rad)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Vow *VowTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Vow.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Vow *VowSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Vow.Contract.Rely(&_Vow.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Vow *VowTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Vow.Contract.Rely(&_Vow.TransactOpts, usr)
}

// VowCageIterator is returned from FilterCage and is used to iterate over the raw logs and unpacked data for Cage events raised by the Vow contract.
type VowCageIterator struct {
	Event *VowCage // Event containing the contract specifics and raw log

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
func (it *VowCageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VowCage)
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
		it.Event = new(VowCage)
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
func (it *VowCageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VowCageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VowCage represents a Cage event raised by the Vow contract.
type VowCage struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCage is a free log retrieval operation binding the contract event 0x2308ed18a14e800c39b86eb6ea43270105955ca385b603b64eca89f98ae8fbda.
//
// Solidity: event Cage()
func (_Vow *VowFilterer) FilterCage(opts *bind.FilterOpts) (*VowCageIterator, error) {

	logs, sub, err := _Vow.contract.FilterLogs(opts, "Cage")
	if err != nil {
		return nil, err
	}
	return &VowCageIterator{contract: _Vow.contract, event: "Cage", logs: logs, sub: sub}, nil
}

// WatchCage is a free log subscription operation binding the contract event 0x2308ed18a14e800c39b86eb6ea43270105955ca385b603b64eca89f98ae8fbda.
//
// Solidity: event Cage()
func (_Vow *VowFilterer) WatchCage(opts *bind.WatchOpts, sink chan<- *VowCage) (event.Subscription, error) {

	logs, sub, err := _Vow.contract.WatchLogs(opts, "Cage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VowCage)
				if err := _Vow.contract.UnpackLog(event, "Cage", log); err != nil {
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

// ParseCage is a log parse operation binding the contract event 0x2308ed18a14e800c39b86eb6ea43270105955ca385b603b64eca89f98ae8fbda.
//
// Solidity: event Cage()
func (_Vow *VowFilterer) ParseCage(log types.Log) (*VowCage, error) {
	event := new(VowCage)
	if err := _Vow.contract.UnpackLog(event, "Cage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VowDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the Vow contract.
type VowDenyIterator struct {
	Event *VowDeny // Event containing the contract specifics and raw log

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
func (it *VowDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VowDeny)
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
		it.Event = new(VowDeny)
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
func (it *VowDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VowDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VowDeny represents a Deny event raised by the Vow contract.
type VowDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_Vow *VowFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*VowDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Vow.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &VowDenyIterator{contract: _Vow.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_Vow *VowFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *VowDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Vow.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VowDeny)
				if err := _Vow.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_Vow *VowFilterer) ParseDeny(log types.Log) (*VowDeny, error) {
	event := new(VowDeny)
	if err := _Vow.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VowFessIterator is returned from FilterFess and is used to iterate over the raw logs and unpacked data for Fess events raised by the Vow contract.
type VowFessIterator struct {
	Event *VowFess // Event containing the contract specifics and raw log

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
func (it *VowFessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VowFess)
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
		it.Event = new(VowFess)
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
func (it *VowFessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VowFessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VowFess represents a Fess event raised by the Vow contract.
type VowFess struct {
	Tab *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFess is a free log retrieval operation binding the contract event 0x7a3f1a1ebf14b193365bc7468b58eb3b80ae1638635424aae4eec386da2f02ba.
//
// Solidity: event Fess(uint256 indexed tab)
func (_Vow *VowFilterer) FilterFess(opts *bind.FilterOpts, tab []*big.Int) (*VowFessIterator, error) {

	var tabRule []interface{}
	for _, tabItem := range tab {
		tabRule = append(tabRule, tabItem)
	}

	logs, sub, err := _Vow.contract.FilterLogs(opts, "Fess", tabRule)
	if err != nil {
		return nil, err
	}
	return &VowFessIterator{contract: _Vow.contract, event: "Fess", logs: logs, sub: sub}, nil
}

// WatchFess is a free log subscription operation binding the contract event 0x7a3f1a1ebf14b193365bc7468b58eb3b80ae1638635424aae4eec386da2f02ba.
//
// Solidity: event Fess(uint256 indexed tab)
func (_Vow *VowFilterer) WatchFess(opts *bind.WatchOpts, sink chan<- *VowFess, tab []*big.Int) (event.Subscription, error) {

	var tabRule []interface{}
	for _, tabItem := range tab {
		tabRule = append(tabRule, tabItem)
	}

	logs, sub, err := _Vow.contract.WatchLogs(opts, "Fess", tabRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VowFess)
				if err := _Vow.contract.UnpackLog(event, "Fess", log); err != nil {
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

// ParseFess is a log parse operation binding the contract event 0x7a3f1a1ebf14b193365bc7468b58eb3b80ae1638635424aae4eec386da2f02ba.
//
// Solidity: event Fess(uint256 indexed tab)
func (_Vow *VowFilterer) ParseFess(log types.Log) (*VowFess, error) {
	event := new(VowFess)
	if err := _Vow.contract.UnpackLog(event, "Fess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VowFileIterator is returned from FilterFile and is used to iterate over the raw logs and unpacked data for File events raised by the Vow contract.
type VowFileIterator struct {
	Event *VowFile // Event containing the contract specifics and raw log

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
func (it *VowFileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VowFile)
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
		it.Event = new(VowFile)
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
func (it *VowFileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VowFileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VowFile represents a File event raised by the Vow contract.
type VowFile struct {
	What [32]byte
	Data *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile is a free log retrieval operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_Vow *VowFilterer) FilterFile(opts *bind.FilterOpts, what [][32]byte) (*VowFileIterator, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Vow.contract.FilterLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return &VowFileIterator{contract: _Vow.contract, event: "File", logs: logs, sub: sub}, nil
}

// WatchFile is a free log subscription operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_Vow *VowFilterer) WatchFile(opts *bind.WatchOpts, sink chan<- *VowFile, what [][32]byte) (event.Subscription, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Vow.contract.WatchLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VowFile)
				if err := _Vow.contract.UnpackLog(event, "File", log); err != nil {
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

// ParseFile is a log parse operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_Vow *VowFilterer) ParseFile(log types.Log) (*VowFile, error) {
	event := new(VowFile)
	if err := _Vow.contract.UnpackLog(event, "File", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VowFile0Iterator is returned from FilterFile0 and is used to iterate over the raw logs and unpacked data for File0 events raised by the Vow contract.
type VowFile0Iterator struct {
	Event *VowFile0 // Event containing the contract specifics and raw log

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
func (it *VowFile0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VowFile0)
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
		it.Event = new(VowFile0)
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
func (it *VowFile0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VowFile0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VowFile0 represents a File0 event raised by the Vow contract.
type VowFile0 struct {
	What [32]byte
	Data common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile0 is a free log retrieval operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 indexed what, address data)
func (_Vow *VowFilterer) FilterFile0(opts *bind.FilterOpts, what [][32]byte) (*VowFile0Iterator, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Vow.contract.FilterLogs(opts, "File0", whatRule)
	if err != nil {
		return nil, err
	}
	return &VowFile0Iterator{contract: _Vow.contract, event: "File0", logs: logs, sub: sub}, nil
}

// WatchFile0 is a free log subscription operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 indexed what, address data)
func (_Vow *VowFilterer) WatchFile0(opts *bind.WatchOpts, sink chan<- *VowFile0, what [][32]byte) (event.Subscription, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Vow.contract.WatchLogs(opts, "File0", whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VowFile0)
				if err := _Vow.contract.UnpackLog(event, "File0", log); err != nil {
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

// ParseFile0 is a log parse operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 indexed what, address data)
func (_Vow *VowFilterer) ParseFile0(log types.Log) (*VowFile0, error) {
	event := new(VowFile0)
	if err := _Vow.contract.UnpackLog(event, "File0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VowFlogIterator is returned from FilterFlog and is used to iterate over the raw logs and unpacked data for Flog events raised by the Vow contract.
type VowFlogIterator struct {
	Event *VowFlog // Event containing the contract specifics and raw log

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
func (it *VowFlogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VowFlog)
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
		it.Event = new(VowFlog)
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
func (it *VowFlogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VowFlogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VowFlog represents a Flog event raised by the Vow contract.
type VowFlog struct {
	Era *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFlog is a free log retrieval operation binding the contract event 0x5aa14c9b66239d17e56d0724b7e90d8d82f28fcbdfb0d39e60614bd1d01dc753.
//
// Solidity: event Flog(uint256 indexed era)
func (_Vow *VowFilterer) FilterFlog(opts *bind.FilterOpts, era []*big.Int) (*VowFlogIterator, error) {

	var eraRule []interface{}
	for _, eraItem := range era {
		eraRule = append(eraRule, eraItem)
	}

	logs, sub, err := _Vow.contract.FilterLogs(opts, "Flog", eraRule)
	if err != nil {
		return nil, err
	}
	return &VowFlogIterator{contract: _Vow.contract, event: "Flog", logs: logs, sub: sub}, nil
}

// WatchFlog is a free log subscription operation binding the contract event 0x5aa14c9b66239d17e56d0724b7e90d8d82f28fcbdfb0d39e60614bd1d01dc753.
//
// Solidity: event Flog(uint256 indexed era)
func (_Vow *VowFilterer) WatchFlog(opts *bind.WatchOpts, sink chan<- *VowFlog, era []*big.Int) (event.Subscription, error) {

	var eraRule []interface{}
	for _, eraItem := range era {
		eraRule = append(eraRule, eraItem)
	}

	logs, sub, err := _Vow.contract.WatchLogs(opts, "Flog", eraRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VowFlog)
				if err := _Vow.contract.UnpackLog(event, "Flog", log); err != nil {
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

// ParseFlog is a log parse operation binding the contract event 0x5aa14c9b66239d17e56d0724b7e90d8d82f28fcbdfb0d39e60614bd1d01dc753.
//
// Solidity: event Flog(uint256 indexed era)
func (_Vow *VowFilterer) ParseFlog(log types.Log) (*VowFlog, error) {
	event := new(VowFlog)
	if err := _Vow.contract.UnpackLog(event, "Flog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VowFlopIterator is returned from FilterFlop and is used to iterate over the raw logs and unpacked data for Flop events raised by the Vow contract.
type VowFlopIterator struct {
	Event *VowFlop // Event containing the contract specifics and raw log

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
func (it *VowFlopIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VowFlop)
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
		it.Event = new(VowFlop)
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
func (it *VowFlopIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VowFlopIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VowFlop represents a Flop event raised by the Vow contract.
type VowFlop struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFlop is a free log retrieval operation binding the contract event 0x2a6fd69fb4120226e0d810a9039fc9a40879738ce61c80379d760f4614f6680d.
//
// Solidity: event Flop()
func (_Vow *VowFilterer) FilterFlop(opts *bind.FilterOpts) (*VowFlopIterator, error) {

	logs, sub, err := _Vow.contract.FilterLogs(opts, "Flop")
	if err != nil {
		return nil, err
	}
	return &VowFlopIterator{contract: _Vow.contract, event: "Flop", logs: logs, sub: sub}, nil
}

// WatchFlop is a free log subscription operation binding the contract event 0x2a6fd69fb4120226e0d810a9039fc9a40879738ce61c80379d760f4614f6680d.
//
// Solidity: event Flop()
func (_Vow *VowFilterer) WatchFlop(opts *bind.WatchOpts, sink chan<- *VowFlop) (event.Subscription, error) {

	logs, sub, err := _Vow.contract.WatchLogs(opts, "Flop")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VowFlop)
				if err := _Vow.contract.UnpackLog(event, "Flop", log); err != nil {
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

// ParseFlop is a log parse operation binding the contract event 0x2a6fd69fb4120226e0d810a9039fc9a40879738ce61c80379d760f4614f6680d.
//
// Solidity: event Flop()
func (_Vow *VowFilterer) ParseFlop(log types.Log) (*VowFlop, error) {
	event := new(VowFlop)
	if err := _Vow.contract.UnpackLog(event, "Flop", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VowHealIterator is returned from FilterHeal and is used to iterate over the raw logs and unpacked data for Heal events raised by the Vow contract.
type VowHealIterator struct {
	Event *VowHeal // Event containing the contract specifics and raw log

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
func (it *VowHealIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VowHeal)
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
		it.Event = new(VowHeal)
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
func (it *VowHealIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VowHealIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VowHeal represents a Heal event raised by the Vow contract.
type VowHeal struct {
	Rad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterHeal is a free log retrieval operation binding the contract event 0x917d6982889419f491488c036c2e6abe788b07222064ab462158ec64ca2c4db7.
//
// Solidity: event Heal(uint256 indexed rad)
func (_Vow *VowFilterer) FilterHeal(opts *bind.FilterOpts, rad []*big.Int) (*VowHealIterator, error) {

	var radRule []interface{}
	for _, radItem := range rad {
		radRule = append(radRule, radItem)
	}

	logs, sub, err := _Vow.contract.FilterLogs(opts, "Heal", radRule)
	if err != nil {
		return nil, err
	}
	return &VowHealIterator{contract: _Vow.contract, event: "Heal", logs: logs, sub: sub}, nil
}

// WatchHeal is a free log subscription operation binding the contract event 0x917d6982889419f491488c036c2e6abe788b07222064ab462158ec64ca2c4db7.
//
// Solidity: event Heal(uint256 indexed rad)
func (_Vow *VowFilterer) WatchHeal(opts *bind.WatchOpts, sink chan<- *VowHeal, rad []*big.Int) (event.Subscription, error) {

	var radRule []interface{}
	for _, radItem := range rad {
		radRule = append(radRule, radItem)
	}

	logs, sub, err := _Vow.contract.WatchLogs(opts, "Heal", radRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VowHeal)
				if err := _Vow.contract.UnpackLog(event, "Heal", log); err != nil {
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

// ParseHeal is a log parse operation binding the contract event 0x917d6982889419f491488c036c2e6abe788b07222064ab462158ec64ca2c4db7.
//
// Solidity: event Heal(uint256 indexed rad)
func (_Vow *VowFilterer) ParseHeal(log types.Log) (*VowHeal, error) {
	event := new(VowHeal)
	if err := _Vow.contract.UnpackLog(event, "Heal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VowKissIterator is returned from FilterKiss and is used to iterate over the raw logs and unpacked data for Kiss events raised by the Vow contract.
type VowKissIterator struct {
	Event *VowKiss // Event containing the contract specifics and raw log

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
func (it *VowKissIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VowKiss)
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
		it.Event = new(VowKiss)
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
func (it *VowKissIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VowKissIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VowKiss represents a Kiss event raised by the Vow contract.
type VowKiss struct {
	Rad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterKiss is a free log retrieval operation binding the contract event 0xdf1d0254f949dd4607095c8a45ed43a96d548776dbb1d6e8347513d07b109e9b.
//
// Solidity: event Kiss(uint256 indexed rad)
func (_Vow *VowFilterer) FilterKiss(opts *bind.FilterOpts, rad []*big.Int) (*VowKissIterator, error) {

	var radRule []interface{}
	for _, radItem := range rad {
		radRule = append(radRule, radItem)
	}

	logs, sub, err := _Vow.contract.FilterLogs(opts, "Kiss", radRule)
	if err != nil {
		return nil, err
	}
	return &VowKissIterator{contract: _Vow.contract, event: "Kiss", logs: logs, sub: sub}, nil
}

// WatchKiss is a free log subscription operation binding the contract event 0xdf1d0254f949dd4607095c8a45ed43a96d548776dbb1d6e8347513d07b109e9b.
//
// Solidity: event Kiss(uint256 indexed rad)
func (_Vow *VowFilterer) WatchKiss(opts *bind.WatchOpts, sink chan<- *VowKiss, rad []*big.Int) (event.Subscription, error) {

	var radRule []interface{}
	for _, radItem := range rad {
		radRule = append(radRule, radItem)
	}

	logs, sub, err := _Vow.contract.WatchLogs(opts, "Kiss", radRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VowKiss)
				if err := _Vow.contract.UnpackLog(event, "Kiss", log); err != nil {
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

// ParseKiss is a log parse operation binding the contract event 0xdf1d0254f949dd4607095c8a45ed43a96d548776dbb1d6e8347513d07b109e9b.
//
// Solidity: event Kiss(uint256 indexed rad)
func (_Vow *VowFilterer) ParseKiss(log types.Log) (*VowKiss, error) {
	event := new(VowKiss)
	if err := _Vow.contract.UnpackLog(event, "Kiss", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VowRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the Vow contract.
type VowRelyIterator struct {
	Event *VowRely // Event containing the contract specifics and raw log

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
func (it *VowRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VowRely)
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
		it.Event = new(VowRely)
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
func (it *VowRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VowRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VowRely represents a Rely event raised by the Vow contract.
type VowRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_Vow *VowFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*VowRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Vow.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &VowRelyIterator{contract: _Vow.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_Vow *VowFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *VowRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Vow.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VowRely)
				if err := _Vow.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_Vow *VowFilterer) ParseRely(log types.Log) (*VowRely, error) {
	event := new(VowRely)
	if err := _Vow.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
