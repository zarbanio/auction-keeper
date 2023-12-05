// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abacus

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

// AbacusMetaData contains all meta data concerning the Abacus contract.
var AbacusMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AbacusABI is the input ABI used to generate the binding from.
// Deprecated: Use AbacusMetaData.ABI instead.
var AbacusABI = AbacusMetaData.ABI

// Abacus is an auto generated Go binding around an Ethereum contract.
type Abacus struct {
	AbacusCaller     // Read-only binding to the contract
	AbacusTransactor // Write-only binding to the contract
	AbacusFilterer   // Log filterer for contract events
}

// AbacusCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbacusCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbacusTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbacusTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbacusFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbacusFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbacusSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbacusSession struct {
	Contract     *Abacus           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbacusCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbacusCallerSession struct {
	Contract *AbacusCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AbacusTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbacusTransactorSession struct {
	Contract     *AbacusTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbacusRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbacusRaw struct {
	Contract *Abacus // Generic contract binding to access the raw methods on
}

// AbacusCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbacusCallerRaw struct {
	Contract *AbacusCaller // Generic read-only contract binding to access the raw methods on
}

// AbacusTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbacusTransactorRaw struct {
	Contract *AbacusTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbacus creates a new instance of Abacus, bound to a specific deployed contract.
func NewAbacus(address common.Address, backend bind.ContractBackend) (*Abacus, error) {
	contract, err := bindAbacus(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abacus{AbacusCaller: AbacusCaller{contract: contract}, AbacusTransactor: AbacusTransactor{contract: contract}, AbacusFilterer: AbacusFilterer{contract: contract}}, nil
}

// NewAbacusCaller creates a new read-only instance of Abacus, bound to a specific deployed contract.
func NewAbacusCaller(address common.Address, caller bind.ContractCaller) (*AbacusCaller, error) {
	contract, err := bindAbacus(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbacusCaller{contract: contract}, nil
}

// NewAbacusTransactor creates a new write-only instance of Abacus, bound to a specific deployed contract.
func NewAbacusTransactor(address common.Address, transactor bind.ContractTransactor) (*AbacusTransactor, error) {
	contract, err := bindAbacus(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbacusTransactor{contract: contract}, nil
}

// NewAbacusFilterer creates a new log filterer instance of Abacus, bound to a specific deployed contract.
func NewAbacusFilterer(address common.Address, filterer bind.ContractFilterer) (*AbacusFilterer, error) {
	contract, err := bindAbacus(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbacusFilterer{contract: contract}, nil
}

// bindAbacus binds a generic wrapper to an already deployed contract.
func bindAbacus(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AbacusMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abacus *AbacusRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abacus.Contract.AbacusCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abacus *AbacusRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abacus.Contract.AbacusTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abacus *AbacusRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abacus.Contract.AbacusTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abacus *AbacusCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abacus.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abacus *AbacusTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abacus.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abacus *AbacusTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abacus.Contract.contract.Transact(opts, method, params...)
}

// Price is a free data retrieval call binding the contract method 0x487a2395.
//
// Solidity: function price(uint256 , uint256 ) view returns(uint256)
func (_Abacus *AbacusCaller) Price(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Abacus.contract.Call(opts, &out, "price", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0x487a2395.
//
// Solidity: function price(uint256 , uint256 ) view returns(uint256)
func (_Abacus *AbacusSession) Price(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Abacus.Contract.Price(&_Abacus.CallOpts, arg0, arg1)
}

// Price is a free data retrieval call binding the contract method 0x487a2395.
//
// Solidity: function price(uint256 , uint256 ) view returns(uint256)
func (_Abacus *AbacusCallerSession) Price(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Abacus.Contract.Price(&_Abacus.CallOpts, arg0, arg1)
}
