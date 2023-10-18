// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ilkregistry

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

// IlkregistryMetaData contains all meta data concerning the Ilkregistry contract.
var IlkregistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dog_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spot_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"AddIlk\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"NameError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"RemoveIlk\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"SymbolError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"UpdateIlk\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"adapter\",\"type\":\"address\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"class\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"dec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dog\",\"outputs\":[{\"internalType\":\"contractDogLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"gem\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pos\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ilkData\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"pos\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"join\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gem\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"dec\",\"type\":\"uint8\"},{\"internalType\":\"uint96\",\"name\":\"class\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"pip\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"xlip\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"info\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"class\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dec\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gem\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pip\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"join\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"xlip\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"join\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"list\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"list\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"pip\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"pos\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_join\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gem\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dec\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_class\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_pip\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_xlip\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"put\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"removeAuth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spot\",\"outputs\":[{\"internalType\":\"contractSpotLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"xlip\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IlkregistryABI is the input ABI used to generate the binding from.
// Deprecated: Use IlkregistryMetaData.ABI instead.
var IlkregistryABI = IlkregistryMetaData.ABI

// Ilkregistry is an auto generated Go binding around an Ethereum contract.
type Ilkregistry struct {
	IlkregistryCaller     // Read-only binding to the contract
	IlkregistryTransactor // Write-only binding to the contract
	IlkregistryFilterer   // Log filterer for contract events
}

// IlkregistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IlkregistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IlkregistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IlkregistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IlkregistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IlkregistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IlkregistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IlkregistrySession struct {
	Contract     *Ilkregistry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IlkregistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IlkregistryCallerSession struct {
	Contract *IlkregistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IlkregistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IlkregistryTransactorSession struct {
	Contract     *IlkregistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IlkregistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IlkregistryRaw struct {
	Contract *Ilkregistry // Generic contract binding to access the raw methods on
}

// IlkregistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IlkregistryCallerRaw struct {
	Contract *IlkregistryCaller // Generic read-only contract binding to access the raw methods on
}

// IlkregistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IlkregistryTransactorRaw struct {
	Contract *IlkregistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIlkregistry creates a new instance of Ilkregistry, bound to a specific deployed contract.
func NewIlkregistry(address common.Address, backend bind.ContractBackend) (*Ilkregistry, error) {
	contract, err := bindIlkregistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ilkregistry{IlkregistryCaller: IlkregistryCaller{contract: contract}, IlkregistryTransactor: IlkregistryTransactor{contract: contract}, IlkregistryFilterer: IlkregistryFilterer{contract: contract}}, nil
}

// NewIlkregistryCaller creates a new read-only instance of Ilkregistry, bound to a specific deployed contract.
func NewIlkregistryCaller(address common.Address, caller bind.ContractCaller) (*IlkregistryCaller, error) {
	contract, err := bindIlkregistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IlkregistryCaller{contract: contract}, nil
}

// NewIlkregistryTransactor creates a new write-only instance of Ilkregistry, bound to a specific deployed contract.
func NewIlkregistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IlkregistryTransactor, error) {
	contract, err := bindIlkregistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IlkregistryTransactor{contract: contract}, nil
}

// NewIlkregistryFilterer creates a new log filterer instance of Ilkregistry, bound to a specific deployed contract.
func NewIlkregistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IlkregistryFilterer, error) {
	contract, err := bindIlkregistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IlkregistryFilterer{contract: contract}, nil
}

// bindIlkregistry binds a generic wrapper to an already deployed contract.
func bindIlkregistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IlkregistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ilkregistry *IlkregistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ilkregistry.Contract.IlkregistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ilkregistry *IlkregistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ilkregistry.Contract.IlkregistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ilkregistry *IlkregistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ilkregistry.Contract.IlkregistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ilkregistry *IlkregistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ilkregistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ilkregistry *IlkregistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ilkregistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ilkregistry *IlkregistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ilkregistry.Contract.contract.Transact(opts, method, params...)
}

// Class is a free data retrieval call binding the contract method 0x217cf12b.
//
// Solidity: function class(bytes32 ilk) view returns(uint256)
func (_Ilkregistry *IlkregistryCaller) Class(opts *bind.CallOpts, ilk [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Ilkregistry.contract.Call(opts, &out, "class", ilk)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Class is a free data retrieval call binding the contract method 0x217cf12b.
//
// Solidity: function class(bytes32 ilk) view returns(uint256)
func (_Ilkregistry *IlkregistrySession) Class(ilk [32]byte) (*big.Int, error) {
	return _Ilkregistry.Contract.Class(&_Ilkregistry.CallOpts, ilk)
}

// Class is a free data retrieval call binding the contract method 0x217cf12b.
//
// Solidity: function class(bytes32 ilk) view returns(uint256)
func (_Ilkregistry *IlkregistryCallerSession) Class(ilk [32]byte) (*big.Int, error) {
	return _Ilkregistry.Contract.Class(&_Ilkregistry.CallOpts, ilk)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Ilkregistry *IlkregistryCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ilkregistry.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Ilkregistry *IlkregistrySession) Count() (*big.Int, error) {
	return _Ilkregistry.Contract.Count(&_Ilkregistry.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Ilkregistry *IlkregistryCallerSession) Count() (*big.Int, error) {
	return _Ilkregistry.Contract.Count(&_Ilkregistry.CallOpts)
}

// Dec is a free data retrieval call binding the contract method 0x3017a54d.
//
// Solidity: function dec(bytes32 ilk) view returns(uint256)
func (_Ilkregistry *IlkregistryCaller) Dec(opts *bind.CallOpts, ilk [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Ilkregistry.contract.Call(opts, &out, "dec", ilk)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dec is a free data retrieval call binding the contract method 0x3017a54d.
//
// Solidity: function dec(bytes32 ilk) view returns(uint256)
func (_Ilkregistry *IlkregistrySession) Dec(ilk [32]byte) (*big.Int, error) {
	return _Ilkregistry.Contract.Dec(&_Ilkregistry.CallOpts, ilk)
}

// Dec is a free data retrieval call binding the contract method 0x3017a54d.
//
// Solidity: function dec(bytes32 ilk) view returns(uint256)
func (_Ilkregistry *IlkregistryCallerSession) Dec(ilk [32]byte) (*big.Int, error) {
	return _Ilkregistry.Contract.Dec(&_Ilkregistry.CallOpts, ilk)
}

// Dog is a free data retrieval call binding the contract method 0xc3b3ad7f.
//
// Solidity: function dog() view returns(address)
func (_Ilkregistry *IlkregistryCaller) Dog(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ilkregistry.contract.Call(opts, &out, "dog")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dog is a free data retrieval call binding the contract method 0xc3b3ad7f.
//
// Solidity: function dog() view returns(address)
func (_Ilkregistry *IlkregistrySession) Dog() (common.Address, error) {
	return _Ilkregistry.Contract.Dog(&_Ilkregistry.CallOpts)
}

// Dog is a free data retrieval call binding the contract method 0xc3b3ad7f.
//
// Solidity: function dog() view returns(address)
func (_Ilkregistry *IlkregistryCallerSession) Dog() (common.Address, error) {
	return _Ilkregistry.Contract.Dog(&_Ilkregistry.CallOpts)
}

// Gem is a free data retrieval call binding the contract method 0x41f0b723.
//
// Solidity: function gem(bytes32 ilk) view returns(address)
func (_Ilkregistry *IlkregistryCaller) Gem(opts *bind.CallOpts, ilk [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Ilkregistry.contract.Call(opts, &out, "gem", ilk)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gem is a free data retrieval call binding the contract method 0x41f0b723.
//
// Solidity: function gem(bytes32 ilk) view returns(address)
func (_Ilkregistry *IlkregistrySession) Gem(ilk [32]byte) (common.Address, error) {
	return _Ilkregistry.Contract.Gem(&_Ilkregistry.CallOpts, ilk)
}

// Gem is a free data retrieval call binding the contract method 0x41f0b723.
//
// Solidity: function gem(bytes32 ilk) view returns(address)
func (_Ilkregistry *IlkregistryCallerSession) Gem(ilk [32]byte) (common.Address, error) {
	return _Ilkregistry.Contract.Gem(&_Ilkregistry.CallOpts, ilk)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 pos) view returns(bytes32)
func (_Ilkregistry *IlkregistryCaller) Get(opts *bind.CallOpts, pos *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Ilkregistry.contract.Call(opts, &out, "get", pos)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 pos) view returns(bytes32)
func (_Ilkregistry *IlkregistrySession) Get(pos *big.Int) ([32]byte, error) {
	return _Ilkregistry.Contract.Get(&_Ilkregistry.CallOpts, pos)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 pos) view returns(bytes32)
func (_Ilkregistry *IlkregistryCallerSession) Get(pos *big.Int) ([32]byte, error) {
	return _Ilkregistry.Contract.Get(&_Ilkregistry.CallOpts, pos)
}

// IlkData is a free data retrieval call binding the contract method 0xa53a42b5.
//
// Solidity: function ilkData(bytes32 ) view returns(uint96 pos, address join, address gem, uint8 dec, uint96 class, address pip, address xlip, string name, string symbol)
func (_Ilkregistry *IlkregistryCaller) IlkData(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Pos    *big.Int
	Join   common.Address
	Gem    common.Address
	Dec    uint8
	Class  *big.Int
	Pip    common.Address
	Xlip   common.Address
	Name   string
	Symbol string
}, error) {
	var out []interface{}
	err := _Ilkregistry.contract.Call(opts, &out, "ilkData", arg0)

	outstruct := new(struct {
		Pos    *big.Int
		Join   common.Address
		Gem    common.Address
		Dec    uint8
		Class  *big.Int
		Pip    common.Address
		Xlip   common.Address
		Name   string
		Symbol string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pos = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Join = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Gem = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Dec = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.Class = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Pip = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.Xlip = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.Name = *abi.ConvertType(out[7], new(string)).(*string)
	outstruct.Symbol = *abi.ConvertType(out[8], new(string)).(*string)

	return *outstruct, err

}

// IlkData is a free data retrieval call binding the contract method 0xa53a42b5.
//
// Solidity: function ilkData(bytes32 ) view returns(uint96 pos, address join, address gem, uint8 dec, uint96 class, address pip, address xlip, string name, string symbol)
func (_Ilkregistry *IlkregistrySession) IlkData(arg0 [32]byte) (struct {
	Pos    *big.Int
	Join   common.Address
	Gem    common.Address
	Dec    uint8
	Class  *big.Int
	Pip    common.Address
	Xlip   common.Address
	Name   string
	Symbol string
}, error) {
	return _Ilkregistry.Contract.IlkData(&_Ilkregistry.CallOpts, arg0)
}

// IlkData is a free data retrieval call binding the contract method 0xa53a42b5.
//
// Solidity: function ilkData(bytes32 ) view returns(uint96 pos, address join, address gem, uint8 dec, uint96 class, address pip, address xlip, string name, string symbol)
func (_Ilkregistry *IlkregistryCallerSession) IlkData(arg0 [32]byte) (struct {
	Pos    *big.Int
	Join   common.Address
	Gem    common.Address
	Dec    uint8
	Class  *big.Int
	Pip    common.Address
	Xlip   common.Address
	Name   string
	Symbol string
}, error) {
	return _Ilkregistry.Contract.IlkData(&_Ilkregistry.CallOpts, arg0)
}

// Info is a free data retrieval call binding the contract method 0xb64a097e.
//
// Solidity: function info(bytes32 ilk) view returns(string name, string symbol, uint256 class, uint256 dec, address gem, address pip, address join, address xlip)
func (_Ilkregistry *IlkregistryCaller) Info(opts *bind.CallOpts, ilk [32]byte) (struct {
	Name   string
	Symbol string
	Class  *big.Int
	Dec    *big.Int
	Gem    common.Address
	Pip    common.Address
	Join   common.Address
	Xlip   common.Address
}, error) {
	var out []interface{}
	err := _Ilkregistry.contract.Call(opts, &out, "info", ilk)

	outstruct := new(struct {
		Name   string
		Symbol string
		Class  *big.Int
		Dec    *big.Int
		Gem    common.Address
		Pip    common.Address
		Join   common.Address
		Xlip   common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Symbol = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Class = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Dec = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Gem = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Pip = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.Join = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.Xlip = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Info is a free data retrieval call binding the contract method 0xb64a097e.
//
// Solidity: function info(bytes32 ilk) view returns(string name, string symbol, uint256 class, uint256 dec, address gem, address pip, address join, address xlip)
func (_Ilkregistry *IlkregistrySession) Info(ilk [32]byte) (struct {
	Name   string
	Symbol string
	Class  *big.Int
	Dec    *big.Int
	Gem    common.Address
	Pip    common.Address
	Join   common.Address
	Xlip   common.Address
}, error) {
	return _Ilkregistry.Contract.Info(&_Ilkregistry.CallOpts, ilk)
}

// Info is a free data retrieval call binding the contract method 0xb64a097e.
//
// Solidity: function info(bytes32 ilk) view returns(string name, string symbol, uint256 class, uint256 dec, address gem, address pip, address join, address xlip)
func (_Ilkregistry *IlkregistryCallerSession) Info(ilk [32]byte) (struct {
	Name   string
	Symbol string
	Class  *big.Int
	Dec    *big.Int
	Gem    common.Address
	Pip    common.Address
	Join   common.Address
	Xlip   common.Address
}, error) {
	return _Ilkregistry.Contract.Info(&_Ilkregistry.CallOpts, ilk)
}

// Join is a free data retrieval call binding the contract method 0xad677d0b.
//
// Solidity: function join(bytes32 ilk) view returns(address)
func (_Ilkregistry *IlkregistryCaller) Join(opts *bind.CallOpts, ilk [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Ilkregistry.contract.Call(opts, &out, "join", ilk)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Join is a free data retrieval call binding the contract method 0xad677d0b.
//
// Solidity: function join(bytes32 ilk) view returns(address)
func (_Ilkregistry *IlkregistrySession) Join(ilk [32]byte) (common.Address, error) {
	return _Ilkregistry.Contract.Join(&_Ilkregistry.CallOpts, ilk)
}

// Join is a free data retrieval call binding the contract method 0xad677d0b.
//
// Solidity: function join(bytes32 ilk) view returns(address)
func (_Ilkregistry *IlkregistryCallerSession) Join(ilk [32]byte) (common.Address, error) {
	return _Ilkregistry.Contract.Join(&_Ilkregistry.CallOpts, ilk)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns(bytes32[])
func (_Ilkregistry *IlkregistryCaller) List(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _Ilkregistry.contract.Call(opts, &out, "list")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns(bytes32[])
func (_Ilkregistry *IlkregistrySession) List() ([][32]byte, error) {
	return _Ilkregistry.Contract.List(&_Ilkregistry.CallOpts)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns(bytes32[])
func (_Ilkregistry *IlkregistryCallerSession) List() ([][32]byte, error) {
	return _Ilkregistry.Contract.List(&_Ilkregistry.CallOpts)
}

// List0 is a free data retrieval call binding the contract method 0x50fd7367.
//
// Solidity: function list(uint256 start, uint256 end) view returns(bytes32[])
func (_Ilkregistry *IlkregistryCaller) List0(opts *bind.CallOpts, start *big.Int, end *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _Ilkregistry.contract.Call(opts, &out, "list0", start, end)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// List0 is a free data retrieval call binding the contract method 0x50fd7367.
//
// Solidity: function list(uint256 start, uint256 end) view returns(bytes32[])
func (_Ilkregistry *IlkregistrySession) List0(start *big.Int, end *big.Int) ([][32]byte, error) {
	return _Ilkregistry.Contract.List0(&_Ilkregistry.CallOpts, start, end)
}

// List0 is a free data retrieval call binding the contract method 0x50fd7367.
//
// Solidity: function list(uint256 start, uint256 end) view returns(bytes32[])
func (_Ilkregistry *IlkregistryCallerSession) List0(start *big.Int, end *big.Int) ([][32]byte, error) {
	return _Ilkregistry.Contract.List0(&_Ilkregistry.CallOpts, start, end)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 ilk) view returns(string)
func (_Ilkregistry *IlkregistryCaller) Name(opts *bind.CallOpts, ilk [32]byte) (string, error) {
	var out []interface{}
	err := _Ilkregistry.contract.Call(opts, &out, "name", ilk)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 ilk) view returns(string)
func (_Ilkregistry *IlkregistrySession) Name(ilk [32]byte) (string, error) {
	return _Ilkregistry.Contract.Name(&_Ilkregistry.CallOpts, ilk)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 ilk) view returns(string)
func (_Ilkregistry *IlkregistryCallerSession) Name(ilk [32]byte) (string, error) {
	return _Ilkregistry.Contract.Name(&_Ilkregistry.CallOpts, ilk)
}

// Pip is a free data retrieval call binding the contract method 0xa4903036.
//
// Solidity: function pip(bytes32 ilk) view returns(address)
func (_Ilkregistry *IlkregistryCaller) Pip(opts *bind.CallOpts, ilk [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Ilkregistry.contract.Call(opts, &out, "pip", ilk)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pip is a free data retrieval call binding the contract method 0xa4903036.
//
// Solidity: function pip(bytes32 ilk) view returns(address)
func (_Ilkregistry *IlkregistrySession) Pip(ilk [32]byte) (common.Address, error) {
	return _Ilkregistry.Contract.Pip(&_Ilkregistry.CallOpts, ilk)
}

// Pip is a free data retrieval call binding the contract method 0xa4903036.
//
// Solidity: function pip(bytes32 ilk) view returns(address)
func (_Ilkregistry *IlkregistryCallerSession) Pip(ilk [32]byte) (common.Address, error) {
	return _Ilkregistry.Contract.Pip(&_Ilkregistry.CallOpts, ilk)
}

// Pos is a free data retrieval call binding the contract method 0x56eac7dc.
//
// Solidity: function pos(bytes32 ilk) view returns(uint256)
func (_Ilkregistry *IlkregistryCaller) Pos(opts *bind.CallOpts, ilk [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Ilkregistry.contract.Call(opts, &out, "pos", ilk)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Pos is a free data retrieval call binding the contract method 0x56eac7dc.
//
// Solidity: function pos(bytes32 ilk) view returns(uint256)
func (_Ilkregistry *IlkregistrySession) Pos(ilk [32]byte) (*big.Int, error) {
	return _Ilkregistry.Contract.Pos(&_Ilkregistry.CallOpts, ilk)
}

// Pos is a free data retrieval call binding the contract method 0x56eac7dc.
//
// Solidity: function pos(bytes32 ilk) view returns(uint256)
func (_Ilkregistry *IlkregistryCallerSession) Pos(ilk [32]byte) (*big.Int, error) {
	return _Ilkregistry.Contract.Pos(&_Ilkregistry.CallOpts, ilk)
}

// Spot is a free data retrieval call binding the contract method 0x6f265b93.
//
// Solidity: function spot() view returns(address)
func (_Ilkregistry *IlkregistryCaller) Spot(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ilkregistry.contract.Call(opts, &out, "spot")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Spot is a free data retrieval call binding the contract method 0x6f265b93.
//
// Solidity: function spot() view returns(address)
func (_Ilkregistry *IlkregistrySession) Spot() (common.Address, error) {
	return _Ilkregistry.Contract.Spot(&_Ilkregistry.CallOpts)
}

// Spot is a free data retrieval call binding the contract method 0x6f265b93.
//
// Solidity: function spot() view returns(address)
func (_Ilkregistry *IlkregistryCallerSession) Spot() (common.Address, error) {
	return _Ilkregistry.Contract.Spot(&_Ilkregistry.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x6baa0330.
//
// Solidity: function symbol(bytes32 ilk) view returns(string)
func (_Ilkregistry *IlkregistryCaller) Symbol(opts *bind.CallOpts, ilk [32]byte) (string, error) {
	var out []interface{}
	err := _Ilkregistry.contract.Call(opts, &out, "symbol", ilk)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x6baa0330.
//
// Solidity: function symbol(bytes32 ilk) view returns(string)
func (_Ilkregistry *IlkregistrySession) Symbol(ilk [32]byte) (string, error) {
	return _Ilkregistry.Contract.Symbol(&_Ilkregistry.CallOpts, ilk)
}

// Symbol is a free data retrieval call binding the contract method 0x6baa0330.
//
// Solidity: function symbol(bytes32 ilk) view returns(string)
func (_Ilkregistry *IlkregistryCallerSession) Symbol(ilk [32]byte) (string, error) {
	return _Ilkregistry.Contract.Symbol(&_Ilkregistry.CallOpts, ilk)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Ilkregistry *IlkregistryCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ilkregistry.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Ilkregistry *IlkregistrySession) Vat() (common.Address, error) {
	return _Ilkregistry.Contract.Vat(&_Ilkregistry.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Ilkregistry *IlkregistryCallerSession) Vat() (common.Address, error) {
	return _Ilkregistry.Contract.Vat(&_Ilkregistry.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Ilkregistry *IlkregistryCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ilkregistry.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Ilkregistry *IlkregistrySession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Ilkregistry.Contract.Wards(&_Ilkregistry.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Ilkregistry *IlkregistryCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Ilkregistry.Contract.Wards(&_Ilkregistry.CallOpts, arg0)
}

// Xlip is a free data retrieval call binding the contract method 0x247c803f.
//
// Solidity: function xlip(bytes32 ilk) view returns(address)
func (_Ilkregistry *IlkregistryCaller) Xlip(opts *bind.CallOpts, ilk [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Ilkregistry.contract.Call(opts, &out, "xlip", ilk)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Xlip is a free data retrieval call binding the contract method 0x247c803f.
//
// Solidity: function xlip(bytes32 ilk) view returns(address)
func (_Ilkregistry *IlkregistrySession) Xlip(ilk [32]byte) (common.Address, error) {
	return _Ilkregistry.Contract.Xlip(&_Ilkregistry.CallOpts, ilk)
}

// Xlip is a free data retrieval call binding the contract method 0x247c803f.
//
// Solidity: function xlip(bytes32 ilk) view returns(address)
func (_Ilkregistry *IlkregistryCallerSession) Xlip(ilk [32]byte) (common.Address, error) {
	return _Ilkregistry.Contract.Xlip(&_Ilkregistry.CallOpts, ilk)
}

// Add is a paid mutator transaction binding the contract method 0x0a3b0a4f.
//
// Solidity: function add(address adapter) returns()
func (_Ilkregistry *IlkregistryTransactor) Add(opts *bind.TransactOpts, adapter common.Address) (*types.Transaction, error) {
	return _Ilkregistry.contract.Transact(opts, "add", adapter)
}

// Add is a paid mutator transaction binding the contract method 0x0a3b0a4f.
//
// Solidity: function add(address adapter) returns()
func (_Ilkregistry *IlkregistrySession) Add(adapter common.Address) (*types.Transaction, error) {
	return _Ilkregistry.Contract.Add(&_Ilkregistry.TransactOpts, adapter)
}

// Add is a paid mutator transaction binding the contract method 0x0a3b0a4f.
//
// Solidity: function add(address adapter) returns()
func (_Ilkregistry *IlkregistryTransactorSession) Add(adapter common.Address) (*types.Transaction, error) {
	return _Ilkregistry.Contract.Add(&_Ilkregistry.TransactOpts, adapter)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Ilkregistry *IlkregistryTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Ilkregistry.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Ilkregistry *IlkregistrySession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Ilkregistry.Contract.Deny(&_Ilkregistry.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Ilkregistry *IlkregistryTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Ilkregistry.Contract.Deny(&_Ilkregistry.TransactOpts, usr)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_Ilkregistry *IlkregistryTransactor) File(opts *bind.TransactOpts, ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Ilkregistry.contract.Transact(opts, "file", ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_Ilkregistry *IlkregistrySession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Ilkregistry.Contract.File(&_Ilkregistry.TransactOpts, ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_Ilkregistry *IlkregistryTransactorSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Ilkregistry.Contract.File(&_Ilkregistry.TransactOpts, ilk, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xc8b97f71.
//
// Solidity: function file(bytes32 ilk, bytes32 what, string data) returns()
func (_Ilkregistry *IlkregistryTransactor) File0(opts *bind.TransactOpts, ilk [32]byte, what [32]byte, data string) (*types.Transaction, error) {
	return _Ilkregistry.contract.Transact(opts, "file0", ilk, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xc8b97f71.
//
// Solidity: function file(bytes32 ilk, bytes32 what, string data) returns()
func (_Ilkregistry *IlkregistrySession) File0(ilk [32]byte, what [32]byte, data string) (*types.Transaction, error) {
	return _Ilkregistry.Contract.File0(&_Ilkregistry.TransactOpts, ilk, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xc8b97f71.
//
// Solidity: function file(bytes32 ilk, bytes32 what, string data) returns()
func (_Ilkregistry *IlkregistryTransactorSession) File0(ilk [32]byte, what [32]byte, data string) (*types.Transaction, error) {
	return _Ilkregistry.Contract.File0(&_Ilkregistry.TransactOpts, ilk, what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_Ilkregistry *IlkregistryTransactor) File1(opts *bind.TransactOpts, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _Ilkregistry.contract.Transact(opts, "file1", what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_Ilkregistry *IlkregistrySession) File1(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _Ilkregistry.Contract.File1(&_Ilkregistry.TransactOpts, what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_Ilkregistry *IlkregistryTransactorSession) File1(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _Ilkregistry.Contract.File1(&_Ilkregistry.TransactOpts, what, data)
}

// File2 is a paid mutator transaction binding the contract method 0xebecb39d.
//
// Solidity: function file(bytes32 ilk, bytes32 what, address data) returns()
func (_Ilkregistry *IlkregistryTransactor) File2(opts *bind.TransactOpts, ilk [32]byte, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _Ilkregistry.contract.Transact(opts, "file2", ilk, what, data)
}

// File2 is a paid mutator transaction binding the contract method 0xebecb39d.
//
// Solidity: function file(bytes32 ilk, bytes32 what, address data) returns()
func (_Ilkregistry *IlkregistrySession) File2(ilk [32]byte, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _Ilkregistry.Contract.File2(&_Ilkregistry.TransactOpts, ilk, what, data)
}

// File2 is a paid mutator transaction binding the contract method 0xebecb39d.
//
// Solidity: function file(bytes32 ilk, bytes32 what, address data) returns()
func (_Ilkregistry *IlkregistryTransactorSession) File2(ilk [32]byte, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _Ilkregistry.Contract.File2(&_Ilkregistry.TransactOpts, ilk, what, data)
}

// Put is a paid mutator transaction binding the contract method 0x4d8835e6.
//
// Solidity: function put(bytes32 _ilk, address _join, address _gem, uint256 _dec, uint256 _class, address _pip, address _xlip, string _name, string _symbol) returns()
func (_Ilkregistry *IlkregistryTransactor) Put(opts *bind.TransactOpts, _ilk [32]byte, _join common.Address, _gem common.Address, _dec *big.Int, _class *big.Int, _pip common.Address, _xlip common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _Ilkregistry.contract.Transact(opts, "put", _ilk, _join, _gem, _dec, _class, _pip, _xlip, _name, _symbol)
}

// Put is a paid mutator transaction binding the contract method 0x4d8835e6.
//
// Solidity: function put(bytes32 _ilk, address _join, address _gem, uint256 _dec, uint256 _class, address _pip, address _xlip, string _name, string _symbol) returns()
func (_Ilkregistry *IlkregistrySession) Put(_ilk [32]byte, _join common.Address, _gem common.Address, _dec *big.Int, _class *big.Int, _pip common.Address, _xlip common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _Ilkregistry.Contract.Put(&_Ilkregistry.TransactOpts, _ilk, _join, _gem, _dec, _class, _pip, _xlip, _name, _symbol)
}

// Put is a paid mutator transaction binding the contract method 0x4d8835e6.
//
// Solidity: function put(bytes32 _ilk, address _join, address _gem, uint256 _dec, uint256 _class, address _pip, address _xlip, string _name, string _symbol) returns()
func (_Ilkregistry *IlkregistryTransactorSession) Put(_ilk [32]byte, _join common.Address, _gem common.Address, _dec *big.Int, _class *big.Int, _pip common.Address, _xlip common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _Ilkregistry.Contract.Put(&_Ilkregistry.TransactOpts, _ilk, _join, _gem, _dec, _class, _pip, _xlip, _name, _symbol)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Ilkregistry *IlkregistryTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Ilkregistry.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Ilkregistry *IlkregistrySession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Ilkregistry.Contract.Rely(&_Ilkregistry.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Ilkregistry *IlkregistryTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Ilkregistry.Contract.Rely(&_Ilkregistry.TransactOpts, usr)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(bytes32 ilk) returns()
func (_Ilkregistry *IlkregistryTransactor) Remove(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _Ilkregistry.contract.Transact(opts, "remove", ilk)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(bytes32 ilk) returns()
func (_Ilkregistry *IlkregistrySession) Remove(ilk [32]byte) (*types.Transaction, error) {
	return _Ilkregistry.Contract.Remove(&_Ilkregistry.TransactOpts, ilk)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(bytes32 ilk) returns()
func (_Ilkregistry *IlkregistryTransactorSession) Remove(ilk [32]byte) (*types.Transaction, error) {
	return _Ilkregistry.Contract.Remove(&_Ilkregistry.TransactOpts, ilk)
}

// RemoveAuth is a paid mutator transaction binding the contract method 0xa19555d9.
//
// Solidity: function removeAuth(bytes32 ilk) returns()
func (_Ilkregistry *IlkregistryTransactor) RemoveAuth(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _Ilkregistry.contract.Transact(opts, "removeAuth", ilk)
}

// RemoveAuth is a paid mutator transaction binding the contract method 0xa19555d9.
//
// Solidity: function removeAuth(bytes32 ilk) returns()
func (_Ilkregistry *IlkregistrySession) RemoveAuth(ilk [32]byte) (*types.Transaction, error) {
	return _Ilkregistry.Contract.RemoveAuth(&_Ilkregistry.TransactOpts, ilk)
}

// RemoveAuth is a paid mutator transaction binding the contract method 0xa19555d9.
//
// Solidity: function removeAuth(bytes32 ilk) returns()
func (_Ilkregistry *IlkregistryTransactorSession) RemoveAuth(ilk [32]byte) (*types.Transaction, error) {
	return _Ilkregistry.Contract.RemoveAuth(&_Ilkregistry.TransactOpts, ilk)
}

// Update is a paid mutator transaction binding the contract method 0x8b147245.
//
// Solidity: function update(bytes32 ilk) returns()
func (_Ilkregistry *IlkregistryTransactor) Update(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _Ilkregistry.contract.Transact(opts, "update", ilk)
}

// Update is a paid mutator transaction binding the contract method 0x8b147245.
//
// Solidity: function update(bytes32 ilk) returns()
func (_Ilkregistry *IlkregistrySession) Update(ilk [32]byte) (*types.Transaction, error) {
	return _Ilkregistry.Contract.Update(&_Ilkregistry.TransactOpts, ilk)
}

// Update is a paid mutator transaction binding the contract method 0x8b147245.
//
// Solidity: function update(bytes32 ilk) returns()
func (_Ilkregistry *IlkregistryTransactorSession) Update(ilk [32]byte) (*types.Transaction, error) {
	return _Ilkregistry.Contract.Update(&_Ilkregistry.TransactOpts, ilk)
}

// IlkregistryAddIlkIterator is returned from FilterAddIlk and is used to iterate over the raw logs and unpacked data for AddIlk events raised by the Ilkregistry contract.
type IlkregistryAddIlkIterator struct {
	Event *IlkregistryAddIlk // Event containing the contract specifics and raw log

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
func (it *IlkregistryAddIlkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IlkregistryAddIlk)
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
		it.Event = new(IlkregistryAddIlk)
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
func (it *IlkregistryAddIlkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IlkregistryAddIlkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IlkregistryAddIlk represents a AddIlk event raised by the Ilkregistry contract.
type IlkregistryAddIlk struct {
	Ilk [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAddIlk is a free log retrieval operation binding the contract event 0x74ceb2982b813d6b690af89638316706e6acb9a48fced388741b61b510f165b7.
//
// Solidity: event AddIlk(bytes32 indexed ilk)
func (_Ilkregistry *IlkregistryFilterer) FilterAddIlk(opts *bind.FilterOpts, ilk [][32]byte) (*IlkregistryAddIlkIterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _Ilkregistry.contract.FilterLogs(opts, "AddIlk", ilkRule)
	if err != nil {
		return nil, err
	}
	return &IlkregistryAddIlkIterator{contract: _Ilkregistry.contract, event: "AddIlk", logs: logs, sub: sub}, nil
}

// WatchAddIlk is a free log subscription operation binding the contract event 0x74ceb2982b813d6b690af89638316706e6acb9a48fced388741b61b510f165b7.
//
// Solidity: event AddIlk(bytes32 indexed ilk)
func (_Ilkregistry *IlkregistryFilterer) WatchAddIlk(opts *bind.WatchOpts, sink chan<- *IlkregistryAddIlk, ilk [][32]byte) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _Ilkregistry.contract.WatchLogs(opts, "AddIlk", ilkRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IlkregistryAddIlk)
				if err := _Ilkregistry.contract.UnpackLog(event, "AddIlk", log); err != nil {
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

// ParseAddIlk is a log parse operation binding the contract event 0x74ceb2982b813d6b690af89638316706e6acb9a48fced388741b61b510f165b7.
//
// Solidity: event AddIlk(bytes32 indexed ilk)
func (_Ilkregistry *IlkregistryFilterer) ParseAddIlk(log types.Log) (*IlkregistryAddIlk, error) {
	event := new(IlkregistryAddIlk)
	if err := _Ilkregistry.contract.UnpackLog(event, "AddIlk", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IlkregistryDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the Ilkregistry contract.
type IlkregistryDenyIterator struct {
	Event *IlkregistryDeny // Event containing the contract specifics and raw log

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
func (it *IlkregistryDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IlkregistryDeny)
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
		it.Event = new(IlkregistryDeny)
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
func (it *IlkregistryDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IlkregistryDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IlkregistryDeny represents a Deny event raised by the Ilkregistry contract.
type IlkregistryDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_Ilkregistry *IlkregistryFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*IlkregistryDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Ilkregistry.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &IlkregistryDenyIterator{contract: _Ilkregistry.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_Ilkregistry *IlkregistryFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *IlkregistryDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Ilkregistry.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IlkregistryDeny)
				if err := _Ilkregistry.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_Ilkregistry *IlkregistryFilterer) ParseDeny(log types.Log) (*IlkregistryDeny, error) {
	event := new(IlkregistryDeny)
	if err := _Ilkregistry.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IlkregistryFileIterator is returned from FilterFile and is used to iterate over the raw logs and unpacked data for File events raised by the Ilkregistry contract.
type IlkregistryFileIterator struct {
	Event *IlkregistryFile // Event containing the contract specifics and raw log

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
func (it *IlkregistryFileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IlkregistryFile)
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
		it.Event = new(IlkregistryFile)
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
func (it *IlkregistryFileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IlkregistryFileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IlkregistryFile represents a File event raised by the Ilkregistry contract.
type IlkregistryFile struct {
	What [32]byte
	Data common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile is a free log retrieval operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 indexed what, address data)
func (_Ilkregistry *IlkregistryFilterer) FilterFile(opts *bind.FilterOpts, what [][32]byte) (*IlkregistryFileIterator, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Ilkregistry.contract.FilterLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return &IlkregistryFileIterator{contract: _Ilkregistry.contract, event: "File", logs: logs, sub: sub}, nil
}

// WatchFile is a free log subscription operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 indexed what, address data)
func (_Ilkregistry *IlkregistryFilterer) WatchFile(opts *bind.WatchOpts, sink chan<- *IlkregistryFile, what [][32]byte) (event.Subscription, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Ilkregistry.contract.WatchLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IlkregistryFile)
				if err := _Ilkregistry.contract.UnpackLog(event, "File", log); err != nil {
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

// ParseFile is a log parse operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 indexed what, address data)
func (_Ilkregistry *IlkregistryFilterer) ParseFile(log types.Log) (*IlkregistryFile, error) {
	event := new(IlkregistryFile)
	if err := _Ilkregistry.contract.UnpackLog(event, "File", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IlkregistryFile0Iterator is returned from FilterFile0 and is used to iterate over the raw logs and unpacked data for File0 events raised by the Ilkregistry contract.
type IlkregistryFile0Iterator struct {
	Event *IlkregistryFile0 // Event containing the contract specifics and raw log

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
func (it *IlkregistryFile0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IlkregistryFile0)
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
		it.Event = new(IlkregistryFile0)
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
func (it *IlkregistryFile0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IlkregistryFile0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IlkregistryFile0 represents a File0 event raised by the Ilkregistry contract.
type IlkregistryFile0 struct {
	Ilk  [32]byte
	What [32]byte
	Data common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile0 is a free log retrieval operation binding the contract event 0x4ff2caaa972a7c6629ea01fae9c93d73cc307d13ea4c369f9bbbb7f9b7e9461d.
//
// Solidity: event File(bytes32 indexed ilk, bytes32 indexed what, address data)
func (_Ilkregistry *IlkregistryFilterer) FilterFile0(opts *bind.FilterOpts, ilk [][32]byte, what [][32]byte) (*IlkregistryFile0Iterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Ilkregistry.contract.FilterLogs(opts, "File0", ilkRule, whatRule)
	if err != nil {
		return nil, err
	}
	return &IlkregistryFile0Iterator{contract: _Ilkregistry.contract, event: "File0", logs: logs, sub: sub}, nil
}

// WatchFile0 is a free log subscription operation binding the contract event 0x4ff2caaa972a7c6629ea01fae9c93d73cc307d13ea4c369f9bbbb7f9b7e9461d.
//
// Solidity: event File(bytes32 indexed ilk, bytes32 indexed what, address data)
func (_Ilkregistry *IlkregistryFilterer) WatchFile0(opts *bind.WatchOpts, sink chan<- *IlkregistryFile0, ilk [][32]byte, what [][32]byte) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Ilkregistry.contract.WatchLogs(opts, "File0", ilkRule, whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IlkregistryFile0)
				if err := _Ilkregistry.contract.UnpackLog(event, "File0", log); err != nil {
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

// ParseFile0 is a log parse operation binding the contract event 0x4ff2caaa972a7c6629ea01fae9c93d73cc307d13ea4c369f9bbbb7f9b7e9461d.
//
// Solidity: event File(bytes32 indexed ilk, bytes32 indexed what, address data)
func (_Ilkregistry *IlkregistryFilterer) ParseFile0(log types.Log) (*IlkregistryFile0, error) {
	event := new(IlkregistryFile0)
	if err := _Ilkregistry.contract.UnpackLog(event, "File0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IlkregistryFile1Iterator is returned from FilterFile1 and is used to iterate over the raw logs and unpacked data for File1 events raised by the Ilkregistry contract.
type IlkregistryFile1Iterator struct {
	Event *IlkregistryFile1 // Event containing the contract specifics and raw log

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
func (it *IlkregistryFile1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IlkregistryFile1)
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
		it.Event = new(IlkregistryFile1)
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
func (it *IlkregistryFile1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IlkregistryFile1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IlkregistryFile1 represents a File1 event raised by the Ilkregistry contract.
type IlkregistryFile1 struct {
	Ilk  [32]byte
	What [32]byte
	Data *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile1 is a free log retrieval operation binding the contract event 0x851aa1caf4888170ad8875449d18f0f512fd6deb2a6571ea1a41fb9f95acbcd1.
//
// Solidity: event File(bytes32 indexed ilk, bytes32 indexed what, uint256 data)
func (_Ilkregistry *IlkregistryFilterer) FilterFile1(opts *bind.FilterOpts, ilk [][32]byte, what [][32]byte) (*IlkregistryFile1Iterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Ilkregistry.contract.FilterLogs(opts, "File1", ilkRule, whatRule)
	if err != nil {
		return nil, err
	}
	return &IlkregistryFile1Iterator{contract: _Ilkregistry.contract, event: "File1", logs: logs, sub: sub}, nil
}

// WatchFile1 is a free log subscription operation binding the contract event 0x851aa1caf4888170ad8875449d18f0f512fd6deb2a6571ea1a41fb9f95acbcd1.
//
// Solidity: event File(bytes32 indexed ilk, bytes32 indexed what, uint256 data)
func (_Ilkregistry *IlkregistryFilterer) WatchFile1(opts *bind.WatchOpts, sink chan<- *IlkregistryFile1, ilk [][32]byte, what [][32]byte) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Ilkregistry.contract.WatchLogs(opts, "File1", ilkRule, whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IlkregistryFile1)
				if err := _Ilkregistry.contract.UnpackLog(event, "File1", log); err != nil {
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

// ParseFile1 is a log parse operation binding the contract event 0x851aa1caf4888170ad8875449d18f0f512fd6deb2a6571ea1a41fb9f95acbcd1.
//
// Solidity: event File(bytes32 indexed ilk, bytes32 indexed what, uint256 data)
func (_Ilkregistry *IlkregistryFilterer) ParseFile1(log types.Log) (*IlkregistryFile1, error) {
	event := new(IlkregistryFile1)
	if err := _Ilkregistry.contract.UnpackLog(event, "File1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IlkregistryFile2Iterator is returned from FilterFile2 and is used to iterate over the raw logs and unpacked data for File2 events raised by the Ilkregistry contract.
type IlkregistryFile2Iterator struct {
	Event *IlkregistryFile2 // Event containing the contract specifics and raw log

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
func (it *IlkregistryFile2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IlkregistryFile2)
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
		it.Event = new(IlkregistryFile2)
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
func (it *IlkregistryFile2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IlkregistryFile2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IlkregistryFile2 represents a File2 event raised by the Ilkregistry contract.
type IlkregistryFile2 struct {
	Ilk  [32]byte
	What [32]byte
	Data string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile2 is a free log retrieval operation binding the contract event 0x6a04c0a277676f3a4d382fc6167bf871235d53006834505ea2d2c6101041eda8.
//
// Solidity: event File(bytes32 indexed ilk, bytes32 indexed what, string data)
func (_Ilkregistry *IlkregistryFilterer) FilterFile2(opts *bind.FilterOpts, ilk [][32]byte, what [][32]byte) (*IlkregistryFile2Iterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Ilkregistry.contract.FilterLogs(opts, "File2", ilkRule, whatRule)
	if err != nil {
		return nil, err
	}
	return &IlkregistryFile2Iterator{contract: _Ilkregistry.contract, event: "File2", logs: logs, sub: sub}, nil
}

// WatchFile2 is a free log subscription operation binding the contract event 0x6a04c0a277676f3a4d382fc6167bf871235d53006834505ea2d2c6101041eda8.
//
// Solidity: event File(bytes32 indexed ilk, bytes32 indexed what, string data)
func (_Ilkregistry *IlkregistryFilterer) WatchFile2(opts *bind.WatchOpts, sink chan<- *IlkregistryFile2, ilk [][32]byte, what [][32]byte) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Ilkregistry.contract.WatchLogs(opts, "File2", ilkRule, whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IlkregistryFile2)
				if err := _Ilkregistry.contract.UnpackLog(event, "File2", log); err != nil {
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

// ParseFile2 is a log parse operation binding the contract event 0x6a04c0a277676f3a4d382fc6167bf871235d53006834505ea2d2c6101041eda8.
//
// Solidity: event File(bytes32 indexed ilk, bytes32 indexed what, string data)
func (_Ilkregistry *IlkregistryFilterer) ParseFile2(log types.Log) (*IlkregistryFile2, error) {
	event := new(IlkregistryFile2)
	if err := _Ilkregistry.contract.UnpackLog(event, "File2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IlkregistryNameErrorIterator is returned from FilterNameError and is used to iterate over the raw logs and unpacked data for NameError events raised by the Ilkregistry contract.
type IlkregistryNameErrorIterator struct {
	Event *IlkregistryNameError // Event containing the contract specifics and raw log

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
func (it *IlkregistryNameErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IlkregistryNameError)
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
		it.Event = new(IlkregistryNameError)
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
func (it *IlkregistryNameErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IlkregistryNameErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IlkregistryNameError represents a NameError event raised by the Ilkregistry contract.
type IlkregistryNameError struct {
	Ilk [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNameError is a free log retrieval operation binding the contract event 0x93272f551c7dd0dd38e4c01ae7b4eeef80d2557b4460caa3ee96697d93bc809a.
//
// Solidity: event NameError(bytes32 indexed ilk)
func (_Ilkregistry *IlkregistryFilterer) FilterNameError(opts *bind.FilterOpts, ilk [][32]byte) (*IlkregistryNameErrorIterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _Ilkregistry.contract.FilterLogs(opts, "NameError", ilkRule)
	if err != nil {
		return nil, err
	}
	return &IlkregistryNameErrorIterator{contract: _Ilkregistry.contract, event: "NameError", logs: logs, sub: sub}, nil
}

// WatchNameError is a free log subscription operation binding the contract event 0x93272f551c7dd0dd38e4c01ae7b4eeef80d2557b4460caa3ee96697d93bc809a.
//
// Solidity: event NameError(bytes32 indexed ilk)
func (_Ilkregistry *IlkregistryFilterer) WatchNameError(opts *bind.WatchOpts, sink chan<- *IlkregistryNameError, ilk [][32]byte) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _Ilkregistry.contract.WatchLogs(opts, "NameError", ilkRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IlkregistryNameError)
				if err := _Ilkregistry.contract.UnpackLog(event, "NameError", log); err != nil {
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

// ParseNameError is a log parse operation binding the contract event 0x93272f551c7dd0dd38e4c01ae7b4eeef80d2557b4460caa3ee96697d93bc809a.
//
// Solidity: event NameError(bytes32 indexed ilk)
func (_Ilkregistry *IlkregistryFilterer) ParseNameError(log types.Log) (*IlkregistryNameError, error) {
	event := new(IlkregistryNameError)
	if err := _Ilkregistry.contract.UnpackLog(event, "NameError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IlkregistryRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the Ilkregistry contract.
type IlkregistryRelyIterator struct {
	Event *IlkregistryRely // Event containing the contract specifics and raw log

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
func (it *IlkregistryRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IlkregistryRely)
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
		it.Event = new(IlkregistryRely)
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
func (it *IlkregistryRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IlkregistryRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IlkregistryRely represents a Rely event raised by the Ilkregistry contract.
type IlkregistryRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_Ilkregistry *IlkregistryFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*IlkregistryRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Ilkregistry.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &IlkregistryRelyIterator{contract: _Ilkregistry.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_Ilkregistry *IlkregistryFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *IlkregistryRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Ilkregistry.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IlkregistryRely)
				if err := _Ilkregistry.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_Ilkregistry *IlkregistryFilterer) ParseRely(log types.Log) (*IlkregistryRely, error) {
	event := new(IlkregistryRely)
	if err := _Ilkregistry.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IlkregistryRemoveIlkIterator is returned from FilterRemoveIlk and is used to iterate over the raw logs and unpacked data for RemoveIlk events raised by the Ilkregistry contract.
type IlkregistryRemoveIlkIterator struct {
	Event *IlkregistryRemoveIlk // Event containing the contract specifics and raw log

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
func (it *IlkregistryRemoveIlkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IlkregistryRemoveIlk)
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
		it.Event = new(IlkregistryRemoveIlk)
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
func (it *IlkregistryRemoveIlkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IlkregistryRemoveIlkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IlkregistryRemoveIlk represents a RemoveIlk event raised by the Ilkregistry contract.
type IlkregistryRemoveIlk struct {
	Ilk [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRemoveIlk is a free log retrieval operation binding the contract event 0x42f3b824eb9d522b949ff3d8f70db1872c46f3fc68b6df1a4c8d6aaebfcb6796.
//
// Solidity: event RemoveIlk(bytes32 indexed ilk)
func (_Ilkregistry *IlkregistryFilterer) FilterRemoveIlk(opts *bind.FilterOpts, ilk [][32]byte) (*IlkregistryRemoveIlkIterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _Ilkregistry.contract.FilterLogs(opts, "RemoveIlk", ilkRule)
	if err != nil {
		return nil, err
	}
	return &IlkregistryRemoveIlkIterator{contract: _Ilkregistry.contract, event: "RemoveIlk", logs: logs, sub: sub}, nil
}

// WatchRemoveIlk is a free log subscription operation binding the contract event 0x42f3b824eb9d522b949ff3d8f70db1872c46f3fc68b6df1a4c8d6aaebfcb6796.
//
// Solidity: event RemoveIlk(bytes32 indexed ilk)
func (_Ilkregistry *IlkregistryFilterer) WatchRemoveIlk(opts *bind.WatchOpts, sink chan<- *IlkregistryRemoveIlk, ilk [][32]byte) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _Ilkregistry.contract.WatchLogs(opts, "RemoveIlk", ilkRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IlkregistryRemoveIlk)
				if err := _Ilkregistry.contract.UnpackLog(event, "RemoveIlk", log); err != nil {
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

// ParseRemoveIlk is a log parse operation binding the contract event 0x42f3b824eb9d522b949ff3d8f70db1872c46f3fc68b6df1a4c8d6aaebfcb6796.
//
// Solidity: event RemoveIlk(bytes32 indexed ilk)
func (_Ilkregistry *IlkregistryFilterer) ParseRemoveIlk(log types.Log) (*IlkregistryRemoveIlk, error) {
	event := new(IlkregistryRemoveIlk)
	if err := _Ilkregistry.contract.UnpackLog(event, "RemoveIlk", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IlkregistrySymbolErrorIterator is returned from FilterSymbolError and is used to iterate over the raw logs and unpacked data for SymbolError events raised by the Ilkregistry contract.
type IlkregistrySymbolErrorIterator struct {
	Event *IlkregistrySymbolError // Event containing the contract specifics and raw log

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
func (it *IlkregistrySymbolErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IlkregistrySymbolError)
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
		it.Event = new(IlkregistrySymbolError)
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
func (it *IlkregistrySymbolErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IlkregistrySymbolErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IlkregistrySymbolError represents a SymbolError event raised by the Ilkregistry contract.
type IlkregistrySymbolError struct {
	Ilk [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSymbolError is a free log retrieval operation binding the contract event 0xd4596cfd8cc9635c5a006e070f5c23e1af9b5d2e65665a8d73958c9e6cc17b4d.
//
// Solidity: event SymbolError(bytes32 indexed ilk)
func (_Ilkregistry *IlkregistryFilterer) FilterSymbolError(opts *bind.FilterOpts, ilk [][32]byte) (*IlkregistrySymbolErrorIterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _Ilkregistry.contract.FilterLogs(opts, "SymbolError", ilkRule)
	if err != nil {
		return nil, err
	}
	return &IlkregistrySymbolErrorIterator{contract: _Ilkregistry.contract, event: "SymbolError", logs: logs, sub: sub}, nil
}

// WatchSymbolError is a free log subscription operation binding the contract event 0xd4596cfd8cc9635c5a006e070f5c23e1af9b5d2e65665a8d73958c9e6cc17b4d.
//
// Solidity: event SymbolError(bytes32 indexed ilk)
func (_Ilkregistry *IlkregistryFilterer) WatchSymbolError(opts *bind.WatchOpts, sink chan<- *IlkregistrySymbolError, ilk [][32]byte) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _Ilkregistry.contract.WatchLogs(opts, "SymbolError", ilkRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IlkregistrySymbolError)
				if err := _Ilkregistry.contract.UnpackLog(event, "SymbolError", log); err != nil {
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

// ParseSymbolError is a log parse operation binding the contract event 0xd4596cfd8cc9635c5a006e070f5c23e1af9b5d2e65665a8d73958c9e6cc17b4d.
//
// Solidity: event SymbolError(bytes32 indexed ilk)
func (_Ilkregistry *IlkregistryFilterer) ParseSymbolError(log types.Log) (*IlkregistrySymbolError, error) {
	event := new(IlkregistrySymbolError)
	if err := _Ilkregistry.contract.UnpackLog(event, "SymbolError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IlkregistryUpdateIlkIterator is returned from FilterUpdateIlk and is used to iterate over the raw logs and unpacked data for UpdateIlk events raised by the Ilkregistry contract.
type IlkregistryUpdateIlkIterator struct {
	Event *IlkregistryUpdateIlk // Event containing the contract specifics and raw log

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
func (it *IlkregistryUpdateIlkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IlkregistryUpdateIlk)
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
		it.Event = new(IlkregistryUpdateIlk)
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
func (it *IlkregistryUpdateIlkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IlkregistryUpdateIlkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IlkregistryUpdateIlk represents a UpdateIlk event raised by the Ilkregistry contract.
type IlkregistryUpdateIlk struct {
	Ilk [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpdateIlk is a free log retrieval operation binding the contract event 0x176e1433f84712b82b982cc7a7b738797bd98e17b0882a6edc1a9a89e3dcbdfa.
//
// Solidity: event UpdateIlk(bytes32 indexed ilk)
func (_Ilkregistry *IlkregistryFilterer) FilterUpdateIlk(opts *bind.FilterOpts, ilk [][32]byte) (*IlkregistryUpdateIlkIterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _Ilkregistry.contract.FilterLogs(opts, "UpdateIlk", ilkRule)
	if err != nil {
		return nil, err
	}
	return &IlkregistryUpdateIlkIterator{contract: _Ilkregistry.contract, event: "UpdateIlk", logs: logs, sub: sub}, nil
}

// WatchUpdateIlk is a free log subscription operation binding the contract event 0x176e1433f84712b82b982cc7a7b738797bd98e17b0882a6edc1a9a89e3dcbdfa.
//
// Solidity: event UpdateIlk(bytes32 indexed ilk)
func (_Ilkregistry *IlkregistryFilterer) WatchUpdateIlk(opts *bind.WatchOpts, sink chan<- *IlkregistryUpdateIlk, ilk [][32]byte) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _Ilkregistry.contract.WatchLogs(opts, "UpdateIlk", ilkRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IlkregistryUpdateIlk)
				if err := _Ilkregistry.contract.UnpackLog(event, "UpdateIlk", log); err != nil {
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

// ParseUpdateIlk is a log parse operation binding the contract event 0x176e1433f84712b82b982cc7a7b738797bd98e17b0882a6edc1a9a89e3dcbdfa.
//
// Solidity: event UpdateIlk(bytes32 indexed ilk)
func (_Ilkregistry *IlkregistryFilterer) ParseUpdateIlk(log types.Log) (*IlkregistryUpdateIlk, error) {
	event := new(IlkregistryUpdateIlk)
	if err := _Ilkregistry.contract.UnpackLog(event, "UpdateIlk", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
