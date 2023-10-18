// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package proxy_registry

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

// ProxyRegistryMetaData contains all meta data concerning the ProxyRegistry contract.
var ProxyRegistryMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[],\"name\":\"build\",\"outputs\":[{\"name\":\"proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"proxies\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"build\",\"outputs\":[{\"name\":\"proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"factory_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]",
}

// ProxyRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ProxyRegistryMetaData.ABI instead.
var ProxyRegistryABI = ProxyRegistryMetaData.ABI

// ProxyRegistry is an auto generated Go binding around an Ethereum contract.
type ProxyRegistry struct {
	ProxyRegistryCaller     // Read-only binding to the contract
	ProxyRegistryTransactor // Write-only binding to the contract
	ProxyRegistryFilterer   // Log filterer for contract events
}

// ProxyRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxyRegistrySession struct {
	Contract     *ProxyRegistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyRegistryCallerSession struct {
	Contract *ProxyRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ProxyRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyRegistryTransactorSession struct {
	Contract     *ProxyRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ProxyRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyRegistryRaw struct {
	Contract *ProxyRegistry // Generic contract binding to access the raw methods on
}

// ProxyRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyRegistryCallerRaw struct {
	Contract *ProxyRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyRegistryTransactorRaw struct {
	Contract *ProxyRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxyRegistry creates a new instance of ProxyRegistry, bound to a specific deployed contract.
func NewProxyRegistry(address common.Address, backend bind.ContractBackend) (*ProxyRegistry, error) {
	contract, err := bindProxyRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProxyRegistry{ProxyRegistryCaller: ProxyRegistryCaller{contract: contract}, ProxyRegistryTransactor: ProxyRegistryTransactor{contract: contract}, ProxyRegistryFilterer: ProxyRegistryFilterer{contract: contract}}, nil
}

// NewProxyRegistryCaller creates a new read-only instance of ProxyRegistry, bound to a specific deployed contract.
func NewProxyRegistryCaller(address common.Address, caller bind.ContractCaller) (*ProxyRegistryCaller, error) {
	contract, err := bindProxyRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyRegistryCaller{contract: contract}, nil
}

// NewProxyRegistryTransactor creates a new write-only instance of ProxyRegistry, bound to a specific deployed contract.
func NewProxyRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyRegistryTransactor, error) {
	contract, err := bindProxyRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyRegistryTransactor{contract: contract}, nil
}

// NewProxyRegistryFilterer creates a new log filterer instance of ProxyRegistry, bound to a specific deployed contract.
func NewProxyRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyRegistryFilterer, error) {
	contract, err := bindProxyRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyRegistryFilterer{contract: contract}, nil
}

// bindProxyRegistry binds a generic wrapper to an already deployed contract.
func bindProxyRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProxyRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProxyRegistry *ProxyRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProxyRegistry.Contract.ProxyRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProxyRegistry *ProxyRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyRegistry.Contract.ProxyRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProxyRegistry *ProxyRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProxyRegistry.Contract.ProxyRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProxyRegistry *ProxyRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProxyRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProxyRegistry *ProxyRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProxyRegistry *ProxyRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProxyRegistry.Contract.contract.Transact(opts, method, params...)
}

// Proxies is a free data retrieval call binding the contract method 0xc4552791.
//
// Solidity: function proxies(address ) view returns(address)
func (_ProxyRegistry *ProxyRegistryCaller) Proxies(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _ProxyRegistry.contract.Call(opts, &out, "proxies", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Proxies is a free data retrieval call binding the contract method 0xc4552791.
//
// Solidity: function proxies(address ) view returns(address)
func (_ProxyRegistry *ProxyRegistrySession) Proxies(arg0 common.Address) (common.Address, error) {
	return _ProxyRegistry.Contract.Proxies(&_ProxyRegistry.CallOpts, arg0)
}

// Proxies is a free data retrieval call binding the contract method 0xc4552791.
//
// Solidity: function proxies(address ) view returns(address)
func (_ProxyRegistry *ProxyRegistryCallerSession) Proxies(arg0 common.Address) (common.Address, error) {
	return _ProxyRegistry.Contract.Proxies(&_ProxyRegistry.CallOpts, arg0)
}

// Build is a paid mutator transaction binding the contract method 0x8e1a55fc.
//
// Solidity: function build() returns(address proxy)
func (_ProxyRegistry *ProxyRegistryTransactor) Build(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyRegistry.contract.Transact(opts, "build")
}

// Build is a paid mutator transaction binding the contract method 0x8e1a55fc.
//
// Solidity: function build() returns(address proxy)
func (_ProxyRegistry *ProxyRegistrySession) Build() (*types.Transaction, error) {
	return _ProxyRegistry.Contract.Build(&_ProxyRegistry.TransactOpts)
}

// Build is a paid mutator transaction binding the contract method 0x8e1a55fc.
//
// Solidity: function build() returns(address proxy)
func (_ProxyRegistry *ProxyRegistryTransactorSession) Build() (*types.Transaction, error) {
	return _ProxyRegistry.Contract.Build(&_ProxyRegistry.TransactOpts)
}

// Build0 is a paid mutator transaction binding the contract method 0xf3701da2.
//
// Solidity: function build(address owner) returns(address proxy)
func (_ProxyRegistry *ProxyRegistryTransactor) Build0(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _ProxyRegistry.contract.Transact(opts, "build0", owner)
}

// Build0 is a paid mutator transaction binding the contract method 0xf3701da2.
//
// Solidity: function build(address owner) returns(address proxy)
func (_ProxyRegistry *ProxyRegistrySession) Build0(owner common.Address) (*types.Transaction, error) {
	return _ProxyRegistry.Contract.Build0(&_ProxyRegistry.TransactOpts, owner)
}

// Build0 is a paid mutator transaction binding the contract method 0xf3701da2.
//
// Solidity: function build(address owner) returns(address proxy)
func (_ProxyRegistry *ProxyRegistryTransactorSession) Build0(owner common.Address) (*types.Transaction, error) {
	return _ProxyRegistry.Contract.Build0(&_ProxyRegistry.TransactOpts, owner)
}
