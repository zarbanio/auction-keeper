// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dog

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

// DogMetaData contains all meta data concerning the Dog contract.
var DogMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"urn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ink\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"art\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"due\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"clip\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Bark\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Cage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"Digs\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"clip\",\"type\":\"address\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Dirt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Hole\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"urn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"kpr\",\"type\":\"address\"}],\"name\":\"bark\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"chop\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"digs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"clip\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ilks\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"clip\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chop\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hole\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dirt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"live\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vow\",\"outputs\":[{\"internalType\":\"contractVowLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DogABI is the input ABI used to generate the binding from.
// Deprecated: Use DogMetaData.ABI instead.
var DogABI = DogMetaData.ABI

// Dog is an auto generated Go binding around an Ethereum contract.
type Dog struct {
	DogCaller     // Read-only binding to the contract
	DogTransactor // Write-only binding to the contract
	DogFilterer   // Log filterer for contract events
}

// DogCaller is an auto generated read-only Go binding around an Ethereum contract.
type DogCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DogTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DogTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DogFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DogFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DogSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DogSession struct {
	Contract     *Dog              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DogCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DogCallerSession struct {
	Contract *DogCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DogTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DogTransactorSession struct {
	Contract     *DogTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DogRaw is an auto generated low-level Go binding around an Ethereum contract.
type DogRaw struct {
	Contract *Dog // Generic contract binding to access the raw methods on
}

// DogCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DogCallerRaw struct {
	Contract *DogCaller // Generic read-only contract binding to access the raw methods on
}

// DogTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DogTransactorRaw struct {
	Contract *DogTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDog creates a new instance of Dog, bound to a specific deployed contract.
func NewDog(address common.Address, backend bind.ContractBackend) (*Dog, error) {
	contract, err := bindDog(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dog{DogCaller: DogCaller{contract: contract}, DogTransactor: DogTransactor{contract: contract}, DogFilterer: DogFilterer{contract: contract}}, nil
}

// NewDogCaller creates a new read-only instance of Dog, bound to a specific deployed contract.
func NewDogCaller(address common.Address, caller bind.ContractCaller) (*DogCaller, error) {
	contract, err := bindDog(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DogCaller{contract: contract}, nil
}

// NewDogTransactor creates a new write-only instance of Dog, bound to a specific deployed contract.
func NewDogTransactor(address common.Address, transactor bind.ContractTransactor) (*DogTransactor, error) {
	contract, err := bindDog(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DogTransactor{contract: contract}, nil
}

// NewDogFilterer creates a new log filterer instance of Dog, bound to a specific deployed contract.
func NewDogFilterer(address common.Address, filterer bind.ContractFilterer) (*DogFilterer, error) {
	contract, err := bindDog(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DogFilterer{contract: contract}, nil
}

// bindDog binds a generic wrapper to an already deployed contract.
func bindDog(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DogMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dog *DogRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dog.Contract.DogCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dog *DogRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dog.Contract.DogTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dog *DogRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dog.Contract.DogTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dog *DogCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dog.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dog *DogTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dog.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dog *DogTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dog.Contract.contract.Transact(opts, method, params...)
}

// Dirt is a free data retrieval call binding the contract method 0xeda6e121.
//
// Solidity: function Dirt() view returns(uint256)
func (_Dog *DogCaller) Dirt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dog.contract.Call(opts, &out, "Dirt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dirt is a free data retrieval call binding the contract method 0xeda6e121.
//
// Solidity: function Dirt() view returns(uint256)
func (_Dog *DogSession) Dirt() (*big.Int, error) {
	return _Dog.Contract.Dirt(&_Dog.CallOpts)
}

// Dirt is a free data retrieval call binding the contract method 0xeda6e121.
//
// Solidity: function Dirt() view returns(uint256)
func (_Dog *DogCallerSession) Dirt() (*big.Int, error) {
	return _Dog.Contract.Dirt(&_Dog.CallOpts)
}

// Hole is a free data retrieval call binding the contract method 0xaf7cfeb1.
//
// Solidity: function Hole() view returns(uint256)
func (_Dog *DogCaller) Hole(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dog.contract.Call(opts, &out, "Hole")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Hole is a free data retrieval call binding the contract method 0xaf7cfeb1.
//
// Solidity: function Hole() view returns(uint256)
func (_Dog *DogSession) Hole() (*big.Int, error) {
	return _Dog.Contract.Hole(&_Dog.CallOpts)
}

// Hole is a free data retrieval call binding the contract method 0xaf7cfeb1.
//
// Solidity: function Hole() view returns(uint256)
func (_Dog *DogCallerSession) Hole() (*big.Int, error) {
	return _Dog.Contract.Hole(&_Dog.CallOpts)
}

// Chop is a free data retrieval call binding the contract method 0xd7926538.
//
// Solidity: function chop(bytes32 ilk) view returns(uint256)
func (_Dog *DogCaller) Chop(opts *bind.CallOpts, ilk [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Dog.contract.Call(opts, &out, "chop", ilk)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Chop is a free data retrieval call binding the contract method 0xd7926538.
//
// Solidity: function chop(bytes32 ilk) view returns(uint256)
func (_Dog *DogSession) Chop(ilk [32]byte) (*big.Int, error) {
	return _Dog.Contract.Chop(&_Dog.CallOpts, ilk)
}

// Chop is a free data retrieval call binding the contract method 0xd7926538.
//
// Solidity: function chop(bytes32 ilk) view returns(uint256)
func (_Dog *DogCallerSession) Chop(ilk [32]byte) (*big.Int, error) {
	return _Dog.Contract.Chop(&_Dog.CallOpts, ilk)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(address clip, uint256 chop, uint256 hole, uint256 dirt)
func (_Dog *DogCaller) Ilks(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Clip common.Address
	Chop *big.Int
	Hole *big.Int
	Dirt *big.Int
}, error) {
	var out []interface{}
	err := _Dog.contract.Call(opts, &out, "ilks", arg0)

	outstruct := new(struct {
		Clip common.Address
		Chop *big.Int
		Hole *big.Int
		Dirt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Clip = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Chop = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Hole = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Dirt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(address clip, uint256 chop, uint256 hole, uint256 dirt)
func (_Dog *DogSession) Ilks(arg0 [32]byte) (struct {
	Clip common.Address
	Chop *big.Int
	Hole *big.Int
	Dirt *big.Int
}, error) {
	return _Dog.Contract.Ilks(&_Dog.CallOpts, arg0)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(address clip, uint256 chop, uint256 hole, uint256 dirt)
func (_Dog *DogCallerSession) Ilks(arg0 [32]byte) (struct {
	Clip common.Address
	Chop *big.Int
	Hole *big.Int
	Dirt *big.Int
}, error) {
	return _Dog.Contract.Ilks(&_Dog.CallOpts, arg0)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_Dog *DogCaller) Live(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dog.contract.Call(opts, &out, "live")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_Dog *DogSession) Live() (*big.Int, error) {
	return _Dog.Contract.Live(&_Dog.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_Dog *DogCallerSession) Live() (*big.Int, error) {
	return _Dog.Contract.Live(&_Dog.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Dog *DogCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dog.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Dog *DogSession) Vat() (common.Address, error) {
	return _Dog.Contract.Vat(&_Dog.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Dog *DogCallerSession) Vat() (common.Address, error) {
	return _Dog.Contract.Vat(&_Dog.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_Dog *DogCaller) Vow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dog.contract.Call(opts, &out, "vow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_Dog *DogSession) Vow() (common.Address, error) {
	return _Dog.Contract.Vow(&_Dog.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_Dog *DogCallerSession) Vow() (common.Address, error) {
	return _Dog.Contract.Vow(&_Dog.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Dog *DogCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dog.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Dog *DogSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Dog.Contract.Wards(&_Dog.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Dog *DogCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Dog.Contract.Wards(&_Dog.CallOpts, arg0)
}

// Bark is a paid mutator transaction binding the contract method 0xed998908.
//
// Solidity: function bark(bytes32 ilk, address urn, address kpr) returns(uint256 id)
func (_Dog *DogTransactor) Bark(opts *bind.TransactOpts, ilk [32]byte, urn common.Address, kpr common.Address) (*types.Transaction, error) {
	return _Dog.contract.Transact(opts, "bark", ilk, urn, kpr)
}

// Bark is a paid mutator transaction binding the contract method 0xed998908.
//
// Solidity: function bark(bytes32 ilk, address urn, address kpr) returns(uint256 id)
func (_Dog *DogSession) Bark(ilk [32]byte, urn common.Address, kpr common.Address) (*types.Transaction, error) {
	return _Dog.Contract.Bark(&_Dog.TransactOpts, ilk, urn, kpr)
}

// Bark is a paid mutator transaction binding the contract method 0xed998908.
//
// Solidity: function bark(bytes32 ilk, address urn, address kpr) returns(uint256 id)
func (_Dog *DogTransactorSession) Bark(ilk [32]byte, urn common.Address, kpr common.Address) (*types.Transaction, error) {
	return _Dog.Contract.Bark(&_Dog.TransactOpts, ilk, urn, kpr)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Dog *DogTransactor) Cage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dog.contract.Transact(opts, "cage")
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Dog *DogSession) Cage() (*types.Transaction, error) {
	return _Dog.Contract.Cage(&_Dog.TransactOpts)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Dog *DogTransactorSession) Cage() (*types.Transaction, error) {
	return _Dog.Contract.Cage(&_Dog.TransactOpts)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Dog *DogTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Dog.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Dog *DogSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Dog.Contract.Deny(&_Dog.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Dog *DogTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Dog.Contract.Deny(&_Dog.TransactOpts, usr)
}

// Digs is a paid mutator transaction binding the contract method 0xc87193f4.
//
// Solidity: function digs(bytes32 ilk, uint256 rad) returns()
func (_Dog *DogTransactor) Digs(opts *bind.TransactOpts, ilk [32]byte, rad *big.Int) (*types.Transaction, error) {
	return _Dog.contract.Transact(opts, "digs", ilk, rad)
}

// Digs is a paid mutator transaction binding the contract method 0xc87193f4.
//
// Solidity: function digs(bytes32 ilk, uint256 rad) returns()
func (_Dog *DogSession) Digs(ilk [32]byte, rad *big.Int) (*types.Transaction, error) {
	return _Dog.Contract.Digs(&_Dog.TransactOpts, ilk, rad)
}

// Digs is a paid mutator transaction binding the contract method 0xc87193f4.
//
// Solidity: function digs(bytes32 ilk, uint256 rad) returns()
func (_Dog *DogTransactorSession) Digs(ilk [32]byte, rad *big.Int) (*types.Transaction, error) {
	return _Dog.Contract.Digs(&_Dog.TransactOpts, ilk, rad)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_Dog *DogTransactor) File(opts *bind.TransactOpts, ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Dog.contract.Transact(opts, "file", ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_Dog *DogSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Dog.Contract.File(&_Dog.TransactOpts, ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_Dog *DogTransactorSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Dog.Contract.File(&_Dog.TransactOpts, ilk, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Dog *DogTransactor) File0(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Dog.contract.Transact(opts, "file0", what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Dog *DogSession) File0(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Dog.Contract.File0(&_Dog.TransactOpts, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Dog *DogTransactorSession) File0(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Dog.Contract.File0(&_Dog.TransactOpts, what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_Dog *DogTransactor) File1(opts *bind.TransactOpts, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _Dog.contract.Transact(opts, "file1", what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_Dog *DogSession) File1(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _Dog.Contract.File1(&_Dog.TransactOpts, what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_Dog *DogTransactorSession) File1(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _Dog.Contract.File1(&_Dog.TransactOpts, what, data)
}

// File2 is a paid mutator transaction binding the contract method 0xebecb39d.
//
// Solidity: function file(bytes32 ilk, bytes32 what, address clip) returns()
func (_Dog *DogTransactor) File2(opts *bind.TransactOpts, ilk [32]byte, what [32]byte, clip common.Address) (*types.Transaction, error) {
	return _Dog.contract.Transact(opts, "file2", ilk, what, clip)
}

// File2 is a paid mutator transaction binding the contract method 0xebecb39d.
//
// Solidity: function file(bytes32 ilk, bytes32 what, address clip) returns()
func (_Dog *DogSession) File2(ilk [32]byte, what [32]byte, clip common.Address) (*types.Transaction, error) {
	return _Dog.Contract.File2(&_Dog.TransactOpts, ilk, what, clip)
}

// File2 is a paid mutator transaction binding the contract method 0xebecb39d.
//
// Solidity: function file(bytes32 ilk, bytes32 what, address clip) returns()
func (_Dog *DogTransactorSession) File2(ilk [32]byte, what [32]byte, clip common.Address) (*types.Transaction, error) {
	return _Dog.Contract.File2(&_Dog.TransactOpts, ilk, what, clip)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Dog *DogTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Dog.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Dog *DogSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Dog.Contract.Rely(&_Dog.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Dog *DogTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Dog.Contract.Rely(&_Dog.TransactOpts, usr)
}

// DogBarkIterator is returned from FilterBark and is used to iterate over the raw logs and unpacked data for Bark events raised by the Dog contract.
type DogBarkIterator struct {
	Event *DogBark // Event containing the contract specifics and raw log

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
func (it *DogBarkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DogBark)
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
		it.Event = new(DogBark)
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
func (it *DogBarkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DogBarkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DogBark represents a Bark event raised by the Dog contract.
type DogBark struct {
	Ilk  [32]byte
	Urn  common.Address
	Ink  *big.Int
	Art  *big.Int
	Due  *big.Int
	Clip common.Address
	Id   *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBark is a free log retrieval operation binding the contract event 0x85258d09e1e4ef299ff3fc11e74af99563f022d21f3f940db982229dc2a3358c.
//
// Solidity: event Bark(bytes32 indexed ilk, address indexed urn, uint256 ink, uint256 art, uint256 due, address clip, uint256 indexed id)
func (_Dog *DogFilterer) FilterBark(opts *bind.FilterOpts, ilk [][32]byte, urn []common.Address, id []*big.Int) (*DogBarkIterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var urnRule []interface{}
	for _, urnItem := range urn {
		urnRule = append(urnRule, urnItem)
	}

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Dog.contract.FilterLogs(opts, "Bark", ilkRule, urnRule, idRule)
	if err != nil {
		return nil, err
	}
	return &DogBarkIterator{contract: _Dog.contract, event: "Bark", logs: logs, sub: sub}, nil
}

// WatchBark is a free log subscription operation binding the contract event 0x85258d09e1e4ef299ff3fc11e74af99563f022d21f3f940db982229dc2a3358c.
//
// Solidity: event Bark(bytes32 indexed ilk, address indexed urn, uint256 ink, uint256 art, uint256 due, address clip, uint256 indexed id)
func (_Dog *DogFilterer) WatchBark(opts *bind.WatchOpts, sink chan<- *DogBark, ilk [][32]byte, urn []common.Address, id []*big.Int) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var urnRule []interface{}
	for _, urnItem := range urn {
		urnRule = append(urnRule, urnItem)
	}

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Dog.contract.WatchLogs(opts, "Bark", ilkRule, urnRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DogBark)
				if err := _Dog.contract.UnpackLog(event, "Bark", log); err != nil {
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

// ParseBark is a log parse operation binding the contract event 0x85258d09e1e4ef299ff3fc11e74af99563f022d21f3f940db982229dc2a3358c.
//
// Solidity: event Bark(bytes32 indexed ilk, address indexed urn, uint256 ink, uint256 art, uint256 due, address clip, uint256 indexed id)
func (_Dog *DogFilterer) ParseBark(log types.Log) (*DogBark, error) {
	event := new(DogBark)
	if err := _Dog.contract.UnpackLog(event, "Bark", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DogCageIterator is returned from FilterCage and is used to iterate over the raw logs and unpacked data for Cage events raised by the Dog contract.
type DogCageIterator struct {
	Event *DogCage // Event containing the contract specifics and raw log

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
func (it *DogCageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DogCage)
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
		it.Event = new(DogCage)
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
func (it *DogCageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DogCageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DogCage represents a Cage event raised by the Dog contract.
type DogCage struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCage is a free log retrieval operation binding the contract event 0x2308ed18a14e800c39b86eb6ea43270105955ca385b603b64eca89f98ae8fbda.
//
// Solidity: event Cage()
func (_Dog *DogFilterer) FilterCage(opts *bind.FilterOpts) (*DogCageIterator, error) {

	logs, sub, err := _Dog.contract.FilterLogs(opts, "Cage")
	if err != nil {
		return nil, err
	}
	return &DogCageIterator{contract: _Dog.contract, event: "Cage", logs: logs, sub: sub}, nil
}

// WatchCage is a free log subscription operation binding the contract event 0x2308ed18a14e800c39b86eb6ea43270105955ca385b603b64eca89f98ae8fbda.
//
// Solidity: event Cage()
func (_Dog *DogFilterer) WatchCage(opts *bind.WatchOpts, sink chan<- *DogCage) (event.Subscription, error) {

	logs, sub, err := _Dog.contract.WatchLogs(opts, "Cage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DogCage)
				if err := _Dog.contract.UnpackLog(event, "Cage", log); err != nil {
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
func (_Dog *DogFilterer) ParseCage(log types.Log) (*DogCage, error) {
	event := new(DogCage)
	if err := _Dog.contract.UnpackLog(event, "Cage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DogDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the Dog contract.
type DogDenyIterator struct {
	Event *DogDeny // Event containing the contract specifics and raw log

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
func (it *DogDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DogDeny)
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
		it.Event = new(DogDeny)
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
func (it *DogDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DogDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DogDeny represents a Deny event raised by the Dog contract.
type DogDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_Dog *DogFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*DogDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Dog.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &DogDenyIterator{contract: _Dog.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_Dog *DogFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *DogDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Dog.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DogDeny)
				if err := _Dog.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_Dog *DogFilterer) ParseDeny(log types.Log) (*DogDeny, error) {
	event := new(DogDeny)
	if err := _Dog.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DogDigsIterator is returned from FilterDigs and is used to iterate over the raw logs and unpacked data for Digs events raised by the Dog contract.
type DogDigsIterator struct {
	Event *DogDigs // Event containing the contract specifics and raw log

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
func (it *DogDigsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DogDigs)
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
		it.Event = new(DogDigs)
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
func (it *DogDigsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DogDigsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DogDigs represents a Digs event raised by the Dog contract.
type DogDigs struct {
	Ilk [32]byte
	Rad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDigs is a free log retrieval operation binding the contract event 0x54f095dc7308776bf01e8580e4dd40fd959ea4bf50b069975768320ef8d77d8a.
//
// Solidity: event Digs(bytes32 indexed ilk, uint256 rad)
func (_Dog *DogFilterer) FilterDigs(opts *bind.FilterOpts, ilk [][32]byte) (*DogDigsIterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _Dog.contract.FilterLogs(opts, "Digs", ilkRule)
	if err != nil {
		return nil, err
	}
	return &DogDigsIterator{contract: _Dog.contract, event: "Digs", logs: logs, sub: sub}, nil
}

// WatchDigs is a free log subscription operation binding the contract event 0x54f095dc7308776bf01e8580e4dd40fd959ea4bf50b069975768320ef8d77d8a.
//
// Solidity: event Digs(bytes32 indexed ilk, uint256 rad)
func (_Dog *DogFilterer) WatchDigs(opts *bind.WatchOpts, sink chan<- *DogDigs, ilk [][32]byte) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _Dog.contract.WatchLogs(opts, "Digs", ilkRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DogDigs)
				if err := _Dog.contract.UnpackLog(event, "Digs", log); err != nil {
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

// ParseDigs is a log parse operation binding the contract event 0x54f095dc7308776bf01e8580e4dd40fd959ea4bf50b069975768320ef8d77d8a.
//
// Solidity: event Digs(bytes32 indexed ilk, uint256 rad)
func (_Dog *DogFilterer) ParseDigs(log types.Log) (*DogDigs, error) {
	event := new(DogDigs)
	if err := _Dog.contract.UnpackLog(event, "Digs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DogFileIterator is returned from FilterFile and is used to iterate over the raw logs and unpacked data for File events raised by the Dog contract.
type DogFileIterator struct {
	Event *DogFile // Event containing the contract specifics and raw log

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
func (it *DogFileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DogFile)
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
		it.Event = new(DogFile)
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
func (it *DogFileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DogFileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DogFile represents a File event raised by the Dog contract.
type DogFile struct {
	What [32]byte
	Data *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile is a free log retrieval operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_Dog *DogFilterer) FilterFile(opts *bind.FilterOpts, what [][32]byte) (*DogFileIterator, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Dog.contract.FilterLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return &DogFileIterator{contract: _Dog.contract, event: "File", logs: logs, sub: sub}, nil
}

// WatchFile is a free log subscription operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_Dog *DogFilterer) WatchFile(opts *bind.WatchOpts, sink chan<- *DogFile, what [][32]byte) (event.Subscription, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Dog.contract.WatchLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DogFile)
				if err := _Dog.contract.UnpackLog(event, "File", log); err != nil {
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
func (_Dog *DogFilterer) ParseFile(log types.Log) (*DogFile, error) {
	event := new(DogFile)
	if err := _Dog.contract.UnpackLog(event, "File", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DogFile0Iterator is returned from FilterFile0 and is used to iterate over the raw logs and unpacked data for File0 events raised by the Dog contract.
type DogFile0Iterator struct {
	Event *DogFile0 // Event containing the contract specifics and raw log

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
func (it *DogFile0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DogFile0)
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
		it.Event = new(DogFile0)
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
func (it *DogFile0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DogFile0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DogFile0 represents a File0 event raised by the Dog contract.
type DogFile0 struct {
	What [32]byte
	Data common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile0 is a free log retrieval operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 indexed what, address data)
func (_Dog *DogFilterer) FilterFile0(opts *bind.FilterOpts, what [][32]byte) (*DogFile0Iterator, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Dog.contract.FilterLogs(opts, "File0", whatRule)
	if err != nil {
		return nil, err
	}
	return &DogFile0Iterator{contract: _Dog.contract, event: "File0", logs: logs, sub: sub}, nil
}

// WatchFile0 is a free log subscription operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 indexed what, address data)
func (_Dog *DogFilterer) WatchFile0(opts *bind.WatchOpts, sink chan<- *DogFile0, what [][32]byte) (event.Subscription, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Dog.contract.WatchLogs(opts, "File0", whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DogFile0)
				if err := _Dog.contract.UnpackLog(event, "File0", log); err != nil {
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
func (_Dog *DogFilterer) ParseFile0(log types.Log) (*DogFile0, error) {
	event := new(DogFile0)
	if err := _Dog.contract.UnpackLog(event, "File0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DogFile1Iterator is returned from FilterFile1 and is used to iterate over the raw logs and unpacked data for File1 events raised by the Dog contract.
type DogFile1Iterator struct {
	Event *DogFile1 // Event containing the contract specifics and raw log

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
func (it *DogFile1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DogFile1)
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
		it.Event = new(DogFile1)
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
func (it *DogFile1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DogFile1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DogFile1 represents a File1 event raised by the Dog contract.
type DogFile1 struct {
	Ilk  [32]byte
	What [32]byte
	Data *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile1 is a free log retrieval operation binding the contract event 0x851aa1caf4888170ad8875449d18f0f512fd6deb2a6571ea1a41fb9f95acbcd1.
//
// Solidity: event File(bytes32 indexed ilk, bytes32 indexed what, uint256 data)
func (_Dog *DogFilterer) FilterFile1(opts *bind.FilterOpts, ilk [][32]byte, what [][32]byte) (*DogFile1Iterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Dog.contract.FilterLogs(opts, "File1", ilkRule, whatRule)
	if err != nil {
		return nil, err
	}
	return &DogFile1Iterator{contract: _Dog.contract, event: "File1", logs: logs, sub: sub}, nil
}

// WatchFile1 is a free log subscription operation binding the contract event 0x851aa1caf4888170ad8875449d18f0f512fd6deb2a6571ea1a41fb9f95acbcd1.
//
// Solidity: event File(bytes32 indexed ilk, bytes32 indexed what, uint256 data)
func (_Dog *DogFilterer) WatchFile1(opts *bind.WatchOpts, sink chan<- *DogFile1, ilk [][32]byte, what [][32]byte) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Dog.contract.WatchLogs(opts, "File1", ilkRule, whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DogFile1)
				if err := _Dog.contract.UnpackLog(event, "File1", log); err != nil {
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
func (_Dog *DogFilterer) ParseFile1(log types.Log) (*DogFile1, error) {
	event := new(DogFile1)
	if err := _Dog.contract.UnpackLog(event, "File1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DogFile2Iterator is returned from FilterFile2 and is used to iterate over the raw logs and unpacked data for File2 events raised by the Dog contract.
type DogFile2Iterator struct {
	Event *DogFile2 // Event containing the contract specifics and raw log

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
func (it *DogFile2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DogFile2)
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
		it.Event = new(DogFile2)
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
func (it *DogFile2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DogFile2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DogFile2 represents a File2 event raised by the Dog contract.
type DogFile2 struct {
	Ilk  [32]byte
	What [32]byte
	Clip common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile2 is a free log retrieval operation binding the contract event 0x4ff2caaa972a7c6629ea01fae9c93d73cc307d13ea4c369f9bbbb7f9b7e9461d.
//
// Solidity: event File(bytes32 indexed ilk, bytes32 indexed what, address clip)
func (_Dog *DogFilterer) FilterFile2(opts *bind.FilterOpts, ilk [][32]byte, what [][32]byte) (*DogFile2Iterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Dog.contract.FilterLogs(opts, "File2", ilkRule, whatRule)
	if err != nil {
		return nil, err
	}
	return &DogFile2Iterator{contract: _Dog.contract, event: "File2", logs: logs, sub: sub}, nil
}

// WatchFile2 is a free log subscription operation binding the contract event 0x4ff2caaa972a7c6629ea01fae9c93d73cc307d13ea4c369f9bbbb7f9b7e9461d.
//
// Solidity: event File(bytes32 indexed ilk, bytes32 indexed what, address clip)
func (_Dog *DogFilterer) WatchFile2(opts *bind.WatchOpts, sink chan<- *DogFile2, ilk [][32]byte, what [][32]byte) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Dog.contract.WatchLogs(opts, "File2", ilkRule, whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DogFile2)
				if err := _Dog.contract.UnpackLog(event, "File2", log); err != nil {
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

// ParseFile2 is a log parse operation binding the contract event 0x4ff2caaa972a7c6629ea01fae9c93d73cc307d13ea4c369f9bbbb7f9b7e9461d.
//
// Solidity: event File(bytes32 indexed ilk, bytes32 indexed what, address clip)
func (_Dog *DogFilterer) ParseFile2(log types.Log) (*DogFile2, error) {
	event := new(DogFile2)
	if err := _Dog.contract.UnpackLog(event, "File2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DogRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the Dog contract.
type DogRelyIterator struct {
	Event *DogRely // Event containing the contract specifics and raw log

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
func (it *DogRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DogRely)
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
		it.Event = new(DogRely)
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
func (it *DogRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DogRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DogRely represents a Rely event raised by the Dog contract.
type DogRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_Dog *DogFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*DogRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Dog.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &DogRelyIterator{contract: _Dog.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_Dog *DogFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *DogRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Dog.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DogRely)
				if err := _Dog.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_Dog *DogFilterer) ParseRely(log types.Log) (*DogRely, error) {
	event := new(DogRely)
	if err := _Dog.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
