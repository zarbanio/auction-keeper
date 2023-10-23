// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cdpmanager

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

// CdpmanagerMetaData contains all meta data concerning the Cdpmanager contract.
var CdpmanagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"own\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"}],\"name\":\"NewCdp\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ok\",\"type\":\"uint256\"}],\"name\":\"cdpAllow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"cdpCan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cdpi\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"}],\"name\":\"enter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"first\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"flux\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"flux\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"dink\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"dart\",\"type\":\"int256\"}],\"name\":\"frob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"}],\"name\":\"give\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ilks\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"last\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"list\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"prev\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"next\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"move\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"open\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owns\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"}],\"name\":\"quit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cdpSrc\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cdpDst\",\"type\":\"uint256\"}],\"name\":\"shift\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ok\",\"type\":\"uint256\"}],\"name\":\"urnAllow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"urnCan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"urns\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CdpmanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use CdpmanagerMetaData.ABI instead.
var CdpmanagerABI = CdpmanagerMetaData.ABI

// Cdpmanager is an auto generated Go binding around an Ethereum contract.
type Cdpmanager struct {
	CdpmanagerCaller     // Read-only binding to the contract
	CdpmanagerTransactor // Write-only binding to the contract
	CdpmanagerFilterer   // Log filterer for contract events
}

// CdpmanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type CdpmanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CdpmanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CdpmanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CdpmanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CdpmanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CdpmanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CdpmanagerSession struct {
	Contract     *Cdpmanager       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CdpmanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CdpmanagerCallerSession struct {
	Contract *CdpmanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CdpmanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CdpmanagerTransactorSession struct {
	Contract     *CdpmanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CdpmanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type CdpmanagerRaw struct {
	Contract *Cdpmanager // Generic contract binding to access the raw methods on
}

// CdpmanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CdpmanagerCallerRaw struct {
	Contract *CdpmanagerCaller // Generic read-only contract binding to access the raw methods on
}

// CdpmanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CdpmanagerTransactorRaw struct {
	Contract *CdpmanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCdpmanager creates a new instance of Cdpmanager, bound to a specific deployed contract.
func NewCdpmanager(address common.Address, backend bind.ContractBackend) (*Cdpmanager, error) {
	contract, err := bindCdpmanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cdpmanager{CdpmanagerCaller: CdpmanagerCaller{contract: contract}, CdpmanagerTransactor: CdpmanagerTransactor{contract: contract}, CdpmanagerFilterer: CdpmanagerFilterer{contract: contract}}, nil
}

// NewCdpmanagerCaller creates a new read-only instance of Cdpmanager, bound to a specific deployed contract.
func NewCdpmanagerCaller(address common.Address, caller bind.ContractCaller) (*CdpmanagerCaller, error) {
	contract, err := bindCdpmanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CdpmanagerCaller{contract: contract}, nil
}

// NewCdpmanagerTransactor creates a new write-only instance of Cdpmanager, bound to a specific deployed contract.
func NewCdpmanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*CdpmanagerTransactor, error) {
	contract, err := bindCdpmanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CdpmanagerTransactor{contract: contract}, nil
}

// NewCdpmanagerFilterer creates a new log filterer instance of Cdpmanager, bound to a specific deployed contract.
func NewCdpmanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*CdpmanagerFilterer, error) {
	contract, err := bindCdpmanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CdpmanagerFilterer{contract: contract}, nil
}

// bindCdpmanager binds a generic wrapper to an already deployed contract.
func bindCdpmanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CdpmanagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cdpmanager *CdpmanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cdpmanager.Contract.CdpmanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cdpmanager *CdpmanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cdpmanager.Contract.CdpmanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cdpmanager *CdpmanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cdpmanager.Contract.CdpmanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cdpmanager *CdpmanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cdpmanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cdpmanager *CdpmanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cdpmanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cdpmanager *CdpmanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cdpmanager.Contract.contract.Transact(opts, method, params...)
}

// CdpCan is a free data retrieval call binding the contract method 0x5aebb460.
//
// Solidity: function cdpCan(address , uint256 , address ) view returns(uint256)
func (_Cdpmanager *CdpmanagerCaller) CdpCan(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cdpmanager.contract.Call(opts, &out, "cdpCan", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CdpCan is a free data retrieval call binding the contract method 0x5aebb460.
//
// Solidity: function cdpCan(address , uint256 , address ) view returns(uint256)
func (_Cdpmanager *CdpmanagerSession) CdpCan(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (*big.Int, error) {
	return _Cdpmanager.Contract.CdpCan(&_Cdpmanager.CallOpts, arg0, arg1, arg2)
}

// CdpCan is a free data retrieval call binding the contract method 0x5aebb460.
//
// Solidity: function cdpCan(address , uint256 , address ) view returns(uint256)
func (_Cdpmanager *CdpmanagerCallerSession) CdpCan(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (*big.Int, error) {
	return _Cdpmanager.Contract.CdpCan(&_Cdpmanager.CallOpts, arg0, arg1, arg2)
}

// Cdpi is a free data retrieval call binding the contract method 0xb3d178f2.
//
// Solidity: function cdpi() view returns(uint256)
func (_Cdpmanager *CdpmanagerCaller) Cdpi(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cdpmanager.contract.Call(opts, &out, "cdpi")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cdpi is a free data retrieval call binding the contract method 0xb3d178f2.
//
// Solidity: function cdpi() view returns(uint256)
func (_Cdpmanager *CdpmanagerSession) Cdpi() (*big.Int, error) {
	return _Cdpmanager.Contract.Cdpi(&_Cdpmanager.CallOpts)
}

// Cdpi is a free data retrieval call binding the contract method 0xb3d178f2.
//
// Solidity: function cdpi() view returns(uint256)
func (_Cdpmanager *CdpmanagerCallerSession) Cdpi() (*big.Int, error) {
	return _Cdpmanager.Contract.Cdpi(&_Cdpmanager.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x05d85eda.
//
// Solidity: function count(address ) view returns(uint256)
func (_Cdpmanager *CdpmanagerCaller) Count(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cdpmanager.contract.Call(opts, &out, "count", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x05d85eda.
//
// Solidity: function count(address ) view returns(uint256)
func (_Cdpmanager *CdpmanagerSession) Count(arg0 common.Address) (*big.Int, error) {
	return _Cdpmanager.Contract.Count(&_Cdpmanager.CallOpts, arg0)
}

// Count is a free data retrieval call binding the contract method 0x05d85eda.
//
// Solidity: function count(address ) view returns(uint256)
func (_Cdpmanager *CdpmanagerCallerSession) Count(arg0 common.Address) (*big.Int, error) {
	return _Cdpmanager.Contract.Count(&_Cdpmanager.CallOpts, arg0)
}

// First is a free data retrieval call binding the contract method 0xfc73d771.
//
// Solidity: function first(address ) view returns(uint256)
func (_Cdpmanager *CdpmanagerCaller) First(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cdpmanager.contract.Call(opts, &out, "first", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// First is a free data retrieval call binding the contract method 0xfc73d771.
//
// Solidity: function first(address ) view returns(uint256)
func (_Cdpmanager *CdpmanagerSession) First(arg0 common.Address) (*big.Int, error) {
	return _Cdpmanager.Contract.First(&_Cdpmanager.CallOpts, arg0)
}

// First is a free data retrieval call binding the contract method 0xfc73d771.
//
// Solidity: function first(address ) view returns(uint256)
func (_Cdpmanager *CdpmanagerCallerSession) First(arg0 common.Address) (*big.Int, error) {
	return _Cdpmanager.Contract.First(&_Cdpmanager.CallOpts, arg0)
}

// Ilks is a free data retrieval call binding the contract method 0x2c2cb9fd.
//
// Solidity: function ilks(uint256 ) view returns(bytes32)
func (_Cdpmanager *CdpmanagerCaller) Ilks(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Cdpmanager.contract.Call(opts, &out, "ilks", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Ilks is a free data retrieval call binding the contract method 0x2c2cb9fd.
//
// Solidity: function ilks(uint256 ) view returns(bytes32)
func (_Cdpmanager *CdpmanagerSession) Ilks(arg0 *big.Int) ([32]byte, error) {
	return _Cdpmanager.Contract.Ilks(&_Cdpmanager.CallOpts, arg0)
}

// Ilks is a free data retrieval call binding the contract method 0x2c2cb9fd.
//
// Solidity: function ilks(uint256 ) view returns(bytes32)
func (_Cdpmanager *CdpmanagerCallerSession) Ilks(arg0 *big.Int) ([32]byte, error) {
	return _Cdpmanager.Contract.Ilks(&_Cdpmanager.CallOpts, arg0)
}

// Last is a free data retrieval call binding the contract method 0x9a816f7d.
//
// Solidity: function last(address ) view returns(uint256)
func (_Cdpmanager *CdpmanagerCaller) Last(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cdpmanager.contract.Call(opts, &out, "last", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Last is a free data retrieval call binding the contract method 0x9a816f7d.
//
// Solidity: function last(address ) view returns(uint256)
func (_Cdpmanager *CdpmanagerSession) Last(arg0 common.Address) (*big.Int, error) {
	return _Cdpmanager.Contract.Last(&_Cdpmanager.CallOpts, arg0)
}

// Last is a free data retrieval call binding the contract method 0x9a816f7d.
//
// Solidity: function last(address ) view returns(uint256)
func (_Cdpmanager *CdpmanagerCallerSession) Last(arg0 common.Address) (*big.Int, error) {
	return _Cdpmanager.Contract.Last(&_Cdpmanager.CallOpts, arg0)
}

// List is a free data retrieval call binding the contract method 0x80c9419e.
//
// Solidity: function list(uint256 ) view returns(uint256 prev, uint256 next)
func (_Cdpmanager *CdpmanagerCaller) List(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Prev *big.Int
	Next *big.Int
}, error) {
	var out []interface{}
	err := _Cdpmanager.contract.Call(opts, &out, "list", arg0)

	outstruct := new(struct {
		Prev *big.Int
		Next *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Prev = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Next = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// List is a free data retrieval call binding the contract method 0x80c9419e.
//
// Solidity: function list(uint256 ) view returns(uint256 prev, uint256 next)
func (_Cdpmanager *CdpmanagerSession) List(arg0 *big.Int) (struct {
	Prev *big.Int
	Next *big.Int
}, error) {
	return _Cdpmanager.Contract.List(&_Cdpmanager.CallOpts, arg0)
}

// List is a free data retrieval call binding the contract method 0x80c9419e.
//
// Solidity: function list(uint256 ) view returns(uint256 prev, uint256 next)
func (_Cdpmanager *CdpmanagerCallerSession) List(arg0 *big.Int) (struct {
	Prev *big.Int
	Next *big.Int
}, error) {
	return _Cdpmanager.Contract.List(&_Cdpmanager.CallOpts, arg0)
}

// Owns is a free data retrieval call binding the contract method 0x8161b120.
//
// Solidity: function owns(uint256 ) view returns(address)
func (_Cdpmanager *CdpmanagerCaller) Owns(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Cdpmanager.contract.Call(opts, &out, "owns", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owns is a free data retrieval call binding the contract method 0x8161b120.
//
// Solidity: function owns(uint256 ) view returns(address)
func (_Cdpmanager *CdpmanagerSession) Owns(arg0 *big.Int) (common.Address, error) {
	return _Cdpmanager.Contract.Owns(&_Cdpmanager.CallOpts, arg0)
}

// Owns is a free data retrieval call binding the contract method 0x8161b120.
//
// Solidity: function owns(uint256 ) view returns(address)
func (_Cdpmanager *CdpmanagerCallerSession) Owns(arg0 *big.Int) (common.Address, error) {
	return _Cdpmanager.Contract.Owns(&_Cdpmanager.CallOpts, arg0)
}

// UrnCan is a free data retrieval call binding the contract method 0xb2b192e6.
//
// Solidity: function urnCan(address , address ) view returns(uint256)
func (_Cdpmanager *CdpmanagerCaller) UrnCan(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cdpmanager.contract.Call(opts, &out, "urnCan", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UrnCan is a free data retrieval call binding the contract method 0xb2b192e6.
//
// Solidity: function urnCan(address , address ) view returns(uint256)
func (_Cdpmanager *CdpmanagerSession) UrnCan(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Cdpmanager.Contract.UrnCan(&_Cdpmanager.CallOpts, arg0, arg1)
}

// UrnCan is a free data retrieval call binding the contract method 0xb2b192e6.
//
// Solidity: function urnCan(address , address ) view returns(uint256)
func (_Cdpmanager *CdpmanagerCallerSession) UrnCan(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Cdpmanager.Contract.UrnCan(&_Cdpmanager.CallOpts, arg0, arg1)
}

// Urns is a free data retrieval call binding the contract method 0x2726b073.
//
// Solidity: function urns(uint256 ) view returns(address)
func (_Cdpmanager *CdpmanagerCaller) Urns(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Cdpmanager.contract.Call(opts, &out, "urns", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Urns is a free data retrieval call binding the contract method 0x2726b073.
//
// Solidity: function urns(uint256 ) view returns(address)
func (_Cdpmanager *CdpmanagerSession) Urns(arg0 *big.Int) (common.Address, error) {
	return _Cdpmanager.Contract.Urns(&_Cdpmanager.CallOpts, arg0)
}

// Urns is a free data retrieval call binding the contract method 0x2726b073.
//
// Solidity: function urns(uint256 ) view returns(address)
func (_Cdpmanager *CdpmanagerCallerSession) Urns(arg0 *big.Int) (common.Address, error) {
	return _Cdpmanager.Contract.Urns(&_Cdpmanager.CallOpts, arg0)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Cdpmanager *CdpmanagerCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cdpmanager.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Cdpmanager *CdpmanagerSession) Vat() (common.Address, error) {
	return _Cdpmanager.Contract.Vat(&_Cdpmanager.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Cdpmanager *CdpmanagerCallerSession) Vat() (common.Address, error) {
	return _Cdpmanager.Contract.Vat(&_Cdpmanager.CallOpts)
}

// CdpAllow is a paid mutator transaction binding the contract method 0x0b63fb62.
//
// Solidity: function cdpAllow(uint256 cdp, address usr, uint256 ok) returns()
func (_Cdpmanager *CdpmanagerTransactor) CdpAllow(opts *bind.TransactOpts, cdp *big.Int, usr common.Address, ok *big.Int) (*types.Transaction, error) {
	return _Cdpmanager.contract.Transact(opts, "cdpAllow", cdp, usr, ok)
}

// CdpAllow is a paid mutator transaction binding the contract method 0x0b63fb62.
//
// Solidity: function cdpAllow(uint256 cdp, address usr, uint256 ok) returns()
func (_Cdpmanager *CdpmanagerSession) CdpAllow(cdp *big.Int, usr common.Address, ok *big.Int) (*types.Transaction, error) {
	return _Cdpmanager.Contract.CdpAllow(&_Cdpmanager.TransactOpts, cdp, usr, ok)
}

// CdpAllow is a paid mutator transaction binding the contract method 0x0b63fb62.
//
// Solidity: function cdpAllow(uint256 cdp, address usr, uint256 ok) returns()
func (_Cdpmanager *CdpmanagerTransactorSession) CdpAllow(cdp *big.Int, usr common.Address, ok *big.Int) (*types.Transaction, error) {
	return _Cdpmanager.Contract.CdpAllow(&_Cdpmanager.TransactOpts, cdp, usr, ok)
}

// Enter is a paid mutator transaction binding the contract method 0x7e348b7d.
//
// Solidity: function enter(address src, uint256 cdp) returns()
func (_Cdpmanager *CdpmanagerTransactor) Enter(opts *bind.TransactOpts, src common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _Cdpmanager.contract.Transact(opts, "enter", src, cdp)
}

// Enter is a paid mutator transaction binding the contract method 0x7e348b7d.
//
// Solidity: function enter(address src, uint256 cdp) returns()
func (_Cdpmanager *CdpmanagerSession) Enter(src common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _Cdpmanager.Contract.Enter(&_Cdpmanager.TransactOpts, src, cdp)
}

// Enter is a paid mutator transaction binding the contract method 0x7e348b7d.
//
// Solidity: function enter(address src, uint256 cdp) returns()
func (_Cdpmanager *CdpmanagerTransactorSession) Enter(src common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _Cdpmanager.Contract.Enter(&_Cdpmanager.TransactOpts, src, cdp)
}

// Flux is a paid mutator transaction binding the contract method 0x18af4d60.
//
// Solidity: function flux(bytes32 ilk, uint256 cdp, address dst, uint256 wad) returns()
func (_Cdpmanager *CdpmanagerTransactor) Flux(opts *bind.TransactOpts, ilk [32]byte, cdp *big.Int, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Cdpmanager.contract.Transact(opts, "flux", ilk, cdp, dst, wad)
}

// Flux is a paid mutator transaction binding the contract method 0x18af4d60.
//
// Solidity: function flux(bytes32 ilk, uint256 cdp, address dst, uint256 wad) returns()
func (_Cdpmanager *CdpmanagerSession) Flux(ilk [32]byte, cdp *big.Int, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Cdpmanager.Contract.Flux(&_Cdpmanager.TransactOpts, ilk, cdp, dst, wad)
}

// Flux is a paid mutator transaction binding the contract method 0x18af4d60.
//
// Solidity: function flux(bytes32 ilk, uint256 cdp, address dst, uint256 wad) returns()
func (_Cdpmanager *CdpmanagerTransactorSession) Flux(ilk [32]byte, cdp *big.Int, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Cdpmanager.Contract.Flux(&_Cdpmanager.TransactOpts, ilk, cdp, dst, wad)
}

// Flux0 is a paid mutator transaction binding the contract method 0x9bb8f838.
//
// Solidity: function flux(uint256 cdp, address dst, uint256 wad) returns()
func (_Cdpmanager *CdpmanagerTransactor) Flux0(opts *bind.TransactOpts, cdp *big.Int, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Cdpmanager.contract.Transact(opts, "flux0", cdp, dst, wad)
}

// Flux0 is a paid mutator transaction binding the contract method 0x9bb8f838.
//
// Solidity: function flux(uint256 cdp, address dst, uint256 wad) returns()
func (_Cdpmanager *CdpmanagerSession) Flux0(cdp *big.Int, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Cdpmanager.Contract.Flux0(&_Cdpmanager.TransactOpts, cdp, dst, wad)
}

// Flux0 is a paid mutator transaction binding the contract method 0x9bb8f838.
//
// Solidity: function flux(uint256 cdp, address dst, uint256 wad) returns()
func (_Cdpmanager *CdpmanagerTransactorSession) Flux0(cdp *big.Int, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Cdpmanager.Contract.Flux0(&_Cdpmanager.TransactOpts, cdp, dst, wad)
}

// Frob is a paid mutator transaction binding the contract method 0x45e6bdcd.
//
// Solidity: function frob(uint256 cdp, int256 dink, int256 dart) returns()
func (_Cdpmanager *CdpmanagerTransactor) Frob(opts *bind.TransactOpts, cdp *big.Int, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Cdpmanager.contract.Transact(opts, "frob", cdp, dink, dart)
}

// Frob is a paid mutator transaction binding the contract method 0x45e6bdcd.
//
// Solidity: function frob(uint256 cdp, int256 dink, int256 dart) returns()
func (_Cdpmanager *CdpmanagerSession) Frob(cdp *big.Int, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Cdpmanager.Contract.Frob(&_Cdpmanager.TransactOpts, cdp, dink, dart)
}

// Frob is a paid mutator transaction binding the contract method 0x45e6bdcd.
//
// Solidity: function frob(uint256 cdp, int256 dink, int256 dart) returns()
func (_Cdpmanager *CdpmanagerTransactorSession) Frob(cdp *big.Int, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Cdpmanager.Contract.Frob(&_Cdpmanager.TransactOpts, cdp, dink, dart)
}

// Give is a paid mutator transaction binding the contract method 0xfcafcc68.
//
// Solidity: function give(uint256 cdp, address dst) returns()
func (_Cdpmanager *CdpmanagerTransactor) Give(opts *bind.TransactOpts, cdp *big.Int, dst common.Address) (*types.Transaction, error) {
	return _Cdpmanager.contract.Transact(opts, "give", cdp, dst)
}

// Give is a paid mutator transaction binding the contract method 0xfcafcc68.
//
// Solidity: function give(uint256 cdp, address dst) returns()
func (_Cdpmanager *CdpmanagerSession) Give(cdp *big.Int, dst common.Address) (*types.Transaction, error) {
	return _Cdpmanager.Contract.Give(&_Cdpmanager.TransactOpts, cdp, dst)
}

// Give is a paid mutator transaction binding the contract method 0xfcafcc68.
//
// Solidity: function give(uint256 cdp, address dst) returns()
func (_Cdpmanager *CdpmanagerTransactorSession) Give(cdp *big.Int, dst common.Address) (*types.Transaction, error) {
	return _Cdpmanager.Contract.Give(&_Cdpmanager.TransactOpts, cdp, dst)
}

// Move is a paid mutator transaction binding the contract method 0xf9f30db6.
//
// Solidity: function move(uint256 cdp, address dst, uint256 rad) returns()
func (_Cdpmanager *CdpmanagerTransactor) Move(opts *bind.TransactOpts, cdp *big.Int, dst common.Address, rad *big.Int) (*types.Transaction, error) {
	return _Cdpmanager.contract.Transact(opts, "move", cdp, dst, rad)
}

// Move is a paid mutator transaction binding the contract method 0xf9f30db6.
//
// Solidity: function move(uint256 cdp, address dst, uint256 rad) returns()
func (_Cdpmanager *CdpmanagerSession) Move(cdp *big.Int, dst common.Address, rad *big.Int) (*types.Transaction, error) {
	return _Cdpmanager.Contract.Move(&_Cdpmanager.TransactOpts, cdp, dst, rad)
}

// Move is a paid mutator transaction binding the contract method 0xf9f30db6.
//
// Solidity: function move(uint256 cdp, address dst, uint256 rad) returns()
func (_Cdpmanager *CdpmanagerTransactorSession) Move(cdp *big.Int, dst common.Address, rad *big.Int) (*types.Transaction, error) {
	return _Cdpmanager.Contract.Move(&_Cdpmanager.TransactOpts, cdp, dst, rad)
}

// Open is a paid mutator transaction binding the contract method 0x6090dec5.
//
// Solidity: function open(bytes32 ilk, address usr) returns(uint256)
func (_Cdpmanager *CdpmanagerTransactor) Open(opts *bind.TransactOpts, ilk [32]byte, usr common.Address) (*types.Transaction, error) {
	return _Cdpmanager.contract.Transact(opts, "open", ilk, usr)
}

// Open is a paid mutator transaction binding the contract method 0x6090dec5.
//
// Solidity: function open(bytes32 ilk, address usr) returns(uint256)
func (_Cdpmanager *CdpmanagerSession) Open(ilk [32]byte, usr common.Address) (*types.Transaction, error) {
	return _Cdpmanager.Contract.Open(&_Cdpmanager.TransactOpts, ilk, usr)
}

// Open is a paid mutator transaction binding the contract method 0x6090dec5.
//
// Solidity: function open(bytes32 ilk, address usr) returns(uint256)
func (_Cdpmanager *CdpmanagerTransactorSession) Open(ilk [32]byte, usr common.Address) (*types.Transaction, error) {
	return _Cdpmanager.Contract.Open(&_Cdpmanager.TransactOpts, ilk, usr)
}

// Quit is a paid mutator transaction binding the contract method 0x1b0dbf72.
//
// Solidity: function quit(uint256 cdp, address dst) returns()
func (_Cdpmanager *CdpmanagerTransactor) Quit(opts *bind.TransactOpts, cdp *big.Int, dst common.Address) (*types.Transaction, error) {
	return _Cdpmanager.contract.Transact(opts, "quit", cdp, dst)
}

// Quit is a paid mutator transaction binding the contract method 0x1b0dbf72.
//
// Solidity: function quit(uint256 cdp, address dst) returns()
func (_Cdpmanager *CdpmanagerSession) Quit(cdp *big.Int, dst common.Address) (*types.Transaction, error) {
	return _Cdpmanager.Contract.Quit(&_Cdpmanager.TransactOpts, cdp, dst)
}

// Quit is a paid mutator transaction binding the contract method 0x1b0dbf72.
//
// Solidity: function quit(uint256 cdp, address dst) returns()
func (_Cdpmanager *CdpmanagerTransactorSession) Quit(cdp *big.Int, dst common.Address) (*types.Transaction, error) {
	return _Cdpmanager.Contract.Quit(&_Cdpmanager.TransactOpts, cdp, dst)
}

// Shift is a paid mutator transaction binding the contract method 0xe50322a2.
//
// Solidity: function shift(uint256 cdpSrc, uint256 cdpDst) returns()
func (_Cdpmanager *CdpmanagerTransactor) Shift(opts *bind.TransactOpts, cdpSrc *big.Int, cdpDst *big.Int) (*types.Transaction, error) {
	return _Cdpmanager.contract.Transact(opts, "shift", cdpSrc, cdpDst)
}

// Shift is a paid mutator transaction binding the contract method 0xe50322a2.
//
// Solidity: function shift(uint256 cdpSrc, uint256 cdpDst) returns()
func (_Cdpmanager *CdpmanagerSession) Shift(cdpSrc *big.Int, cdpDst *big.Int) (*types.Transaction, error) {
	return _Cdpmanager.Contract.Shift(&_Cdpmanager.TransactOpts, cdpSrc, cdpDst)
}

// Shift is a paid mutator transaction binding the contract method 0xe50322a2.
//
// Solidity: function shift(uint256 cdpSrc, uint256 cdpDst) returns()
func (_Cdpmanager *CdpmanagerTransactorSession) Shift(cdpSrc *big.Int, cdpDst *big.Int) (*types.Transaction, error) {
	return _Cdpmanager.Contract.Shift(&_Cdpmanager.TransactOpts, cdpSrc, cdpDst)
}

// UrnAllow is a paid mutator transaction binding the contract method 0xb68f4004.
//
// Solidity: function urnAllow(address usr, uint256 ok) returns()
func (_Cdpmanager *CdpmanagerTransactor) UrnAllow(opts *bind.TransactOpts, usr common.Address, ok *big.Int) (*types.Transaction, error) {
	return _Cdpmanager.contract.Transact(opts, "urnAllow", usr, ok)
}

// UrnAllow is a paid mutator transaction binding the contract method 0xb68f4004.
//
// Solidity: function urnAllow(address usr, uint256 ok) returns()
func (_Cdpmanager *CdpmanagerSession) UrnAllow(usr common.Address, ok *big.Int) (*types.Transaction, error) {
	return _Cdpmanager.Contract.UrnAllow(&_Cdpmanager.TransactOpts, usr, ok)
}

// UrnAllow is a paid mutator transaction binding the contract method 0xb68f4004.
//
// Solidity: function urnAllow(address usr, uint256 ok) returns()
func (_Cdpmanager *CdpmanagerTransactorSession) UrnAllow(usr common.Address, ok *big.Int) (*types.Transaction, error) {
	return _Cdpmanager.Contract.UrnAllow(&_Cdpmanager.TransactOpts, usr, ok)
}

// CdpmanagerNewCdpIterator is returned from FilterNewCdp and is used to iterate over the raw logs and unpacked data for NewCdp events raised by the Cdpmanager contract.
type CdpmanagerNewCdpIterator struct {
	Event *CdpmanagerNewCdp // Event containing the contract specifics and raw log

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
func (it *CdpmanagerNewCdpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdpmanagerNewCdp)
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
		it.Event = new(CdpmanagerNewCdp)
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
func (it *CdpmanagerNewCdpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdpmanagerNewCdpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdpmanagerNewCdp represents a NewCdp event raised by the Cdpmanager contract.
type CdpmanagerNewCdp struct {
	Usr common.Address
	Own common.Address
	Cdp *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNewCdp is a free log retrieval operation binding the contract event 0xd6be0bc178658a382ff4f91c8c68b542aa6b71685b8fe427966b87745c3ea7a2.
//
// Solidity: event NewCdp(address indexed usr, address indexed own, uint256 indexed cdp)
func (_Cdpmanager *CdpmanagerFilterer) FilterNewCdp(opts *bind.FilterOpts, usr []common.Address, own []common.Address, cdp []*big.Int) (*CdpmanagerNewCdpIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}
	var ownRule []interface{}
	for _, ownItem := range own {
		ownRule = append(ownRule, ownItem)
	}
	var cdpRule []interface{}
	for _, cdpItem := range cdp {
		cdpRule = append(cdpRule, cdpItem)
	}

	logs, sub, err := _Cdpmanager.contract.FilterLogs(opts, "NewCdp", usrRule, ownRule, cdpRule)
	if err != nil {
		return nil, err
	}
	return &CdpmanagerNewCdpIterator{contract: _Cdpmanager.contract, event: "NewCdp", logs: logs, sub: sub}, nil
}

// WatchNewCdp is a free log subscription operation binding the contract event 0xd6be0bc178658a382ff4f91c8c68b542aa6b71685b8fe427966b87745c3ea7a2.
//
// Solidity: event NewCdp(address indexed usr, address indexed own, uint256 indexed cdp)
func (_Cdpmanager *CdpmanagerFilterer) WatchNewCdp(opts *bind.WatchOpts, sink chan<- *CdpmanagerNewCdp, usr []common.Address, own []common.Address, cdp []*big.Int) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}
	var ownRule []interface{}
	for _, ownItem := range own {
		ownRule = append(ownRule, ownItem)
	}
	var cdpRule []interface{}
	for _, cdpItem := range cdp {
		cdpRule = append(cdpRule, cdpItem)
	}

	logs, sub, err := _Cdpmanager.contract.WatchLogs(opts, "NewCdp", usrRule, ownRule, cdpRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdpmanagerNewCdp)
				if err := _Cdpmanager.contract.UnpackLog(event, "NewCdp", log); err != nil {
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

// ParseNewCdp is a log parse operation binding the contract event 0xd6be0bc178658a382ff4f91c8c68b542aa6b71685b8fe427966b87745c3ea7a2.
//
// Solidity: event NewCdp(address indexed usr, address indexed own, uint256 indexed cdp)
func (_Cdpmanager *CdpmanagerFilterer) ParseNewCdp(log types.Log) (*CdpmanagerNewCdp, error) {
	event := new(CdpmanagerNewCdp)
	if err := _Cdpmanager.contract.UnpackLog(event, "NewCdp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
