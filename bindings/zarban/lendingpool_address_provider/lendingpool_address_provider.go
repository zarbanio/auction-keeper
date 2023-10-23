// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lendingpool_address_provider

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

// LendingpoolAddressProviderMetaData contains all meta data concerning the LendingpoolAddressProvider contract.
var LendingpoolAddressProviderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"marketId\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"hasProxy\",\"type\":\"bool\"}],\"name\":\"AddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"ConfigurationAdminUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"EmergencyAdminUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingPoolCollateralManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingPoolConfiguratorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingPoolUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingRateOracleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newMarketId\",\"type\":\"string\"}],\"name\":\"MarketIdSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PriceOracleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"ProxyCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEmergencyAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLendingPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLendingPoolCollateralManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLendingPoolConfigurator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLendingRateOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMarketId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriceOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"implementationAddress\",\"type\":\"address\"}],\"name\":\"setAddressAsProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"emergencyAdmin\",\"type\":\"address\"}],\"name\":\"setEmergencyAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"setLendingPoolCollateralManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"configurator\",\"type\":\"address\"}],\"name\":\"setLendingPoolConfiguratorImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"setLendingPoolImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lendingRateOracle\",\"type\":\"address\"}],\"name\":\"setLendingRateOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"marketId\",\"type\":\"string\"}],\"name\":\"setMarketId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"setPoolAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"priceOracle\",\"type\":\"address\"}],\"name\":\"setPriceOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LendingpoolAddressProviderABI is the input ABI used to generate the binding from.
// Deprecated: Use LendingpoolAddressProviderMetaData.ABI instead.
var LendingpoolAddressProviderABI = LendingpoolAddressProviderMetaData.ABI

// LendingpoolAddressProvider is an auto generated Go binding around an Ethereum contract.
type LendingpoolAddressProvider struct {
	LendingpoolAddressProviderCaller     // Read-only binding to the contract
	LendingpoolAddressProviderTransactor // Write-only binding to the contract
	LendingpoolAddressProviderFilterer   // Log filterer for contract events
}

// LendingpoolAddressProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type LendingpoolAddressProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingpoolAddressProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LendingpoolAddressProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingpoolAddressProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LendingpoolAddressProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingpoolAddressProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LendingpoolAddressProviderSession struct {
	Contract     *LendingpoolAddressProvider // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// LendingpoolAddressProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LendingpoolAddressProviderCallerSession struct {
	Contract *LendingpoolAddressProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// LendingpoolAddressProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LendingpoolAddressProviderTransactorSession struct {
	Contract     *LendingpoolAddressProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// LendingpoolAddressProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type LendingpoolAddressProviderRaw struct {
	Contract *LendingpoolAddressProvider // Generic contract binding to access the raw methods on
}

// LendingpoolAddressProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LendingpoolAddressProviderCallerRaw struct {
	Contract *LendingpoolAddressProviderCaller // Generic read-only contract binding to access the raw methods on
}

// LendingpoolAddressProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LendingpoolAddressProviderTransactorRaw struct {
	Contract *LendingpoolAddressProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLendingpoolAddressProvider creates a new instance of LendingpoolAddressProvider, bound to a specific deployed contract.
func NewLendingpoolAddressProvider(address common.Address, backend bind.ContractBackend) (*LendingpoolAddressProvider, error) {
	contract, err := bindLendingpoolAddressProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LendingpoolAddressProvider{LendingpoolAddressProviderCaller: LendingpoolAddressProviderCaller{contract: contract}, LendingpoolAddressProviderTransactor: LendingpoolAddressProviderTransactor{contract: contract}, LendingpoolAddressProviderFilterer: LendingpoolAddressProviderFilterer{contract: contract}}, nil
}

// NewLendingpoolAddressProviderCaller creates a new read-only instance of LendingpoolAddressProvider, bound to a specific deployed contract.
func NewLendingpoolAddressProviderCaller(address common.Address, caller bind.ContractCaller) (*LendingpoolAddressProviderCaller, error) {
	contract, err := bindLendingpoolAddressProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LendingpoolAddressProviderCaller{contract: contract}, nil
}

// NewLendingpoolAddressProviderTransactor creates a new write-only instance of LendingpoolAddressProvider, bound to a specific deployed contract.
func NewLendingpoolAddressProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*LendingpoolAddressProviderTransactor, error) {
	contract, err := bindLendingpoolAddressProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LendingpoolAddressProviderTransactor{contract: contract}, nil
}

// NewLendingpoolAddressProviderFilterer creates a new log filterer instance of LendingpoolAddressProvider, bound to a specific deployed contract.
func NewLendingpoolAddressProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*LendingpoolAddressProviderFilterer, error) {
	contract, err := bindLendingpoolAddressProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LendingpoolAddressProviderFilterer{contract: contract}, nil
}

// bindLendingpoolAddressProvider binds a generic wrapper to an already deployed contract.
func bindLendingpoolAddressProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LendingpoolAddressProviderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LendingpoolAddressProvider *LendingpoolAddressProviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LendingpoolAddressProvider.Contract.LendingpoolAddressProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LendingpoolAddressProvider *LendingpoolAddressProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.Contract.LendingpoolAddressProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LendingpoolAddressProvider *LendingpoolAddressProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.Contract.LendingpoolAddressProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LendingpoolAddressProvider *LendingpoolAddressProviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LendingpoolAddressProvider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LendingpoolAddressProvider *LendingpoolAddressProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LendingpoolAddressProvider *LendingpoolAddressProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) view returns(address)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderCaller) GetAddress(opts *bind.CallOpts, id [32]byte) (common.Address, error) {
	var out []interface{}
	err := _LendingpoolAddressProvider.contract.Call(opts, &out, "getAddress", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) view returns(address)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderSession) GetAddress(id [32]byte) (common.Address, error) {
	return _LendingpoolAddressProvider.Contract.GetAddress(&_LendingpoolAddressProvider.CallOpts, id)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) view returns(address)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderCallerSession) GetAddress(id [32]byte) (common.Address, error) {
	return _LendingpoolAddressProvider.Contract.GetAddress(&_LendingpoolAddressProvider.CallOpts, id)
}

// GetEmergencyAdmin is a free data retrieval call binding the contract method 0xddcaa9ea.
//
// Solidity: function getEmergencyAdmin() view returns(address)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderCaller) GetEmergencyAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingpoolAddressProvider.contract.Call(opts, &out, "getEmergencyAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEmergencyAdmin is a free data retrieval call binding the contract method 0xddcaa9ea.
//
// Solidity: function getEmergencyAdmin() view returns(address)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderSession) GetEmergencyAdmin() (common.Address, error) {
	return _LendingpoolAddressProvider.Contract.GetEmergencyAdmin(&_LendingpoolAddressProvider.CallOpts)
}

// GetEmergencyAdmin is a free data retrieval call binding the contract method 0xddcaa9ea.
//
// Solidity: function getEmergencyAdmin() view returns(address)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderCallerSession) GetEmergencyAdmin() (common.Address, error) {
	return _LendingpoolAddressProvider.Contract.GetEmergencyAdmin(&_LendingpoolAddressProvider.CallOpts)
}

// GetLendingPool is a free data retrieval call binding the contract method 0x0261bf8b.
//
// Solidity: function getLendingPool() view returns(address)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderCaller) GetLendingPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingpoolAddressProvider.contract.Call(opts, &out, "getLendingPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingPool is a free data retrieval call binding the contract method 0x0261bf8b.
//
// Solidity: function getLendingPool() view returns(address)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderSession) GetLendingPool() (common.Address, error) {
	return _LendingpoolAddressProvider.Contract.GetLendingPool(&_LendingpoolAddressProvider.CallOpts)
}

// GetLendingPool is a free data retrieval call binding the contract method 0x0261bf8b.
//
// Solidity: function getLendingPool() view returns(address)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderCallerSession) GetLendingPool() (common.Address, error) {
	return _LendingpoolAddressProvider.Contract.GetLendingPool(&_LendingpoolAddressProvider.CallOpts)
}

// GetLendingPoolCollateralManager is a free data retrieval call binding the contract method 0x712d9171.
//
// Solidity: function getLendingPoolCollateralManager() view returns(address)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderCaller) GetLendingPoolCollateralManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingpoolAddressProvider.contract.Call(opts, &out, "getLendingPoolCollateralManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingPoolCollateralManager is a free data retrieval call binding the contract method 0x712d9171.
//
// Solidity: function getLendingPoolCollateralManager() view returns(address)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderSession) GetLendingPoolCollateralManager() (common.Address, error) {
	return _LendingpoolAddressProvider.Contract.GetLendingPoolCollateralManager(&_LendingpoolAddressProvider.CallOpts)
}

// GetLendingPoolCollateralManager is a free data retrieval call binding the contract method 0x712d9171.
//
// Solidity: function getLendingPoolCollateralManager() view returns(address)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderCallerSession) GetLendingPoolCollateralManager() (common.Address, error) {
	return _LendingpoolAddressProvider.Contract.GetLendingPoolCollateralManager(&_LendingpoolAddressProvider.CallOpts)
}

// GetLendingPoolConfigurator is a free data retrieval call binding the contract method 0x85c858b1.
//
// Solidity: function getLendingPoolConfigurator() view returns(address)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderCaller) GetLendingPoolConfigurator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingpoolAddressProvider.contract.Call(opts, &out, "getLendingPoolConfigurator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingPoolConfigurator is a free data retrieval call binding the contract method 0x85c858b1.
//
// Solidity: function getLendingPoolConfigurator() view returns(address)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderSession) GetLendingPoolConfigurator() (common.Address, error) {
	return _LendingpoolAddressProvider.Contract.GetLendingPoolConfigurator(&_LendingpoolAddressProvider.CallOpts)
}

// GetLendingPoolConfigurator is a free data retrieval call binding the contract method 0x85c858b1.
//
// Solidity: function getLendingPoolConfigurator() view returns(address)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderCallerSession) GetLendingPoolConfigurator() (common.Address, error) {
	return _LendingpoolAddressProvider.Contract.GetLendingPoolConfigurator(&_LendingpoolAddressProvider.CallOpts)
}

// GetLendingRateOracle is a free data retrieval call binding the contract method 0x3618abba.
//
// Solidity: function getLendingRateOracle() view returns(address)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderCaller) GetLendingRateOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingpoolAddressProvider.contract.Call(opts, &out, "getLendingRateOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingRateOracle is a free data retrieval call binding the contract method 0x3618abba.
//
// Solidity: function getLendingRateOracle() view returns(address)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderSession) GetLendingRateOracle() (common.Address, error) {
	return _LendingpoolAddressProvider.Contract.GetLendingRateOracle(&_LendingpoolAddressProvider.CallOpts)
}

// GetLendingRateOracle is a free data retrieval call binding the contract method 0x3618abba.
//
// Solidity: function getLendingRateOracle() view returns(address)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderCallerSession) GetLendingRateOracle() (common.Address, error) {
	return _LendingpoolAddressProvider.Contract.GetLendingRateOracle(&_LendingpoolAddressProvider.CallOpts)
}

// GetMarketId is a free data retrieval call binding the contract method 0x568ef470.
//
// Solidity: function getMarketId() view returns(string)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderCaller) GetMarketId(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LendingpoolAddressProvider.contract.Call(opts, &out, "getMarketId")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMarketId is a free data retrieval call binding the contract method 0x568ef470.
//
// Solidity: function getMarketId() view returns(string)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderSession) GetMarketId() (string, error) {
	return _LendingpoolAddressProvider.Contract.GetMarketId(&_LendingpoolAddressProvider.CallOpts)
}

// GetMarketId is a free data retrieval call binding the contract method 0x568ef470.
//
// Solidity: function getMarketId() view returns(string)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderCallerSession) GetMarketId() (string, error) {
	return _LendingpoolAddressProvider.Contract.GetMarketId(&_LendingpoolAddressProvider.CallOpts)
}

// GetPoolAdmin is a free data retrieval call binding the contract method 0xaecda378.
//
// Solidity: function getPoolAdmin() view returns(address)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderCaller) GetPoolAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingpoolAddressProvider.contract.Call(opts, &out, "getPoolAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolAdmin is a free data retrieval call binding the contract method 0xaecda378.
//
// Solidity: function getPoolAdmin() view returns(address)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderSession) GetPoolAdmin() (common.Address, error) {
	return _LendingpoolAddressProvider.Contract.GetPoolAdmin(&_LendingpoolAddressProvider.CallOpts)
}

// GetPoolAdmin is a free data retrieval call binding the contract method 0xaecda378.
//
// Solidity: function getPoolAdmin() view returns(address)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderCallerSession) GetPoolAdmin() (common.Address, error) {
	return _LendingpoolAddressProvider.Contract.GetPoolAdmin(&_LendingpoolAddressProvider.CallOpts)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderCaller) GetPriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingpoolAddressProvider.contract.Call(opts, &out, "getPriceOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderSession) GetPriceOracle() (common.Address, error) {
	return _LendingpoolAddressProvider.Contract.GetPriceOracle(&_LendingpoolAddressProvider.CallOpts)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderCallerSession) GetPriceOracle() (common.Address, error) {
	return _LendingpoolAddressProvider.Contract.GetPriceOracle(&_LendingpoolAddressProvider.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingpoolAddressProvider.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderSession) Owner() (common.Address, error) {
	return _LendingpoolAddressProvider.Contract.Owner(&_LendingpoolAddressProvider.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderCallerSession) Owner() (common.Address, error) {
	return _LendingpoolAddressProvider.Contract.Owner(&_LendingpoolAddressProvider.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderSession) RenounceOwnership() (*types.Transaction, error) {
	return _LendingpoolAddressProvider.Contract.RenounceOwnership(&_LendingpoolAddressProvider.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LendingpoolAddressProvider.Contract.RenounceOwnership(&_LendingpoolAddressProvider.TransactOpts)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 id, address newAddress) returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderTransactor) SetAddress(opts *bind.TransactOpts, id [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.contract.Transact(opts, "setAddress", id, newAddress)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 id, address newAddress) returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderSession) SetAddress(id [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.Contract.SetAddress(&_LendingpoolAddressProvider.TransactOpts, id, newAddress)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 id, address newAddress) returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderTransactorSession) SetAddress(id [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.Contract.SetAddress(&_LendingpoolAddressProvider.TransactOpts, id, newAddress)
}

// SetAddressAsProxy is a paid mutator transaction binding the contract method 0x5dcc528c.
//
// Solidity: function setAddressAsProxy(bytes32 id, address implementationAddress) returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderTransactor) SetAddressAsProxy(opts *bind.TransactOpts, id [32]byte, implementationAddress common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.contract.Transact(opts, "setAddressAsProxy", id, implementationAddress)
}

// SetAddressAsProxy is a paid mutator transaction binding the contract method 0x5dcc528c.
//
// Solidity: function setAddressAsProxy(bytes32 id, address implementationAddress) returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderSession) SetAddressAsProxy(id [32]byte, implementationAddress common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.Contract.SetAddressAsProxy(&_LendingpoolAddressProvider.TransactOpts, id, implementationAddress)
}

// SetAddressAsProxy is a paid mutator transaction binding the contract method 0x5dcc528c.
//
// Solidity: function setAddressAsProxy(bytes32 id, address implementationAddress) returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderTransactorSession) SetAddressAsProxy(id [32]byte, implementationAddress common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.Contract.SetAddressAsProxy(&_LendingpoolAddressProvider.TransactOpts, id, implementationAddress)
}

// SetEmergencyAdmin is a paid mutator transaction binding the contract method 0x35da3394.
//
// Solidity: function setEmergencyAdmin(address emergencyAdmin) returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderTransactor) SetEmergencyAdmin(opts *bind.TransactOpts, emergencyAdmin common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.contract.Transact(opts, "setEmergencyAdmin", emergencyAdmin)
}

// SetEmergencyAdmin is a paid mutator transaction binding the contract method 0x35da3394.
//
// Solidity: function setEmergencyAdmin(address emergencyAdmin) returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderSession) SetEmergencyAdmin(emergencyAdmin common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.Contract.SetEmergencyAdmin(&_LendingpoolAddressProvider.TransactOpts, emergencyAdmin)
}

// SetEmergencyAdmin is a paid mutator transaction binding the contract method 0x35da3394.
//
// Solidity: function setEmergencyAdmin(address emergencyAdmin) returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderTransactorSession) SetEmergencyAdmin(emergencyAdmin common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.Contract.SetEmergencyAdmin(&_LendingpoolAddressProvider.TransactOpts, emergencyAdmin)
}

// SetLendingPoolCollateralManager is a paid mutator transaction binding the contract method 0x398e5553.
//
// Solidity: function setLendingPoolCollateralManager(address manager) returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderTransactor) SetLendingPoolCollateralManager(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.contract.Transact(opts, "setLendingPoolCollateralManager", manager)
}

// SetLendingPoolCollateralManager is a paid mutator transaction binding the contract method 0x398e5553.
//
// Solidity: function setLendingPoolCollateralManager(address manager) returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderSession) SetLendingPoolCollateralManager(manager common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.Contract.SetLendingPoolCollateralManager(&_LendingpoolAddressProvider.TransactOpts, manager)
}

// SetLendingPoolCollateralManager is a paid mutator transaction binding the contract method 0x398e5553.
//
// Solidity: function setLendingPoolCollateralManager(address manager) returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderTransactorSession) SetLendingPoolCollateralManager(manager common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.Contract.SetLendingPoolCollateralManager(&_LendingpoolAddressProvider.TransactOpts, manager)
}

// SetLendingPoolConfiguratorImpl is a paid mutator transaction binding the contract method 0xc12542df.
//
// Solidity: function setLendingPoolConfiguratorImpl(address configurator) returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderTransactor) SetLendingPoolConfiguratorImpl(opts *bind.TransactOpts, configurator common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.contract.Transact(opts, "setLendingPoolConfiguratorImpl", configurator)
}

// SetLendingPoolConfiguratorImpl is a paid mutator transaction binding the contract method 0xc12542df.
//
// Solidity: function setLendingPoolConfiguratorImpl(address configurator) returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderSession) SetLendingPoolConfiguratorImpl(configurator common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.Contract.SetLendingPoolConfiguratorImpl(&_LendingpoolAddressProvider.TransactOpts, configurator)
}

// SetLendingPoolConfiguratorImpl is a paid mutator transaction binding the contract method 0xc12542df.
//
// Solidity: function setLendingPoolConfiguratorImpl(address configurator) returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderTransactorSession) SetLendingPoolConfiguratorImpl(configurator common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.Contract.SetLendingPoolConfiguratorImpl(&_LendingpoolAddressProvider.TransactOpts, configurator)
}

// SetLendingPoolImpl is a paid mutator transaction binding the contract method 0x5aef021f.
//
// Solidity: function setLendingPoolImpl(address pool) returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderTransactor) SetLendingPoolImpl(opts *bind.TransactOpts, pool common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.contract.Transact(opts, "setLendingPoolImpl", pool)
}

// SetLendingPoolImpl is a paid mutator transaction binding the contract method 0x5aef021f.
//
// Solidity: function setLendingPoolImpl(address pool) returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderSession) SetLendingPoolImpl(pool common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.Contract.SetLendingPoolImpl(&_LendingpoolAddressProvider.TransactOpts, pool)
}

// SetLendingPoolImpl is a paid mutator transaction binding the contract method 0x5aef021f.
//
// Solidity: function setLendingPoolImpl(address pool) returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderTransactorSession) SetLendingPoolImpl(pool common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.Contract.SetLendingPoolImpl(&_LendingpoolAddressProvider.TransactOpts, pool)
}

// SetLendingRateOracle is a paid mutator transaction binding the contract method 0x820d1274.
//
// Solidity: function setLendingRateOracle(address lendingRateOracle) returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderTransactor) SetLendingRateOracle(opts *bind.TransactOpts, lendingRateOracle common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.contract.Transact(opts, "setLendingRateOracle", lendingRateOracle)
}

// SetLendingRateOracle is a paid mutator transaction binding the contract method 0x820d1274.
//
// Solidity: function setLendingRateOracle(address lendingRateOracle) returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderSession) SetLendingRateOracle(lendingRateOracle common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.Contract.SetLendingRateOracle(&_LendingpoolAddressProvider.TransactOpts, lendingRateOracle)
}

// SetLendingRateOracle is a paid mutator transaction binding the contract method 0x820d1274.
//
// Solidity: function setLendingRateOracle(address lendingRateOracle) returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderTransactorSession) SetLendingRateOracle(lendingRateOracle common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.Contract.SetLendingRateOracle(&_LendingpoolAddressProvider.TransactOpts, lendingRateOracle)
}

// SetMarketId is a paid mutator transaction binding the contract method 0xf67b1847.
//
// Solidity: function setMarketId(string marketId) returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderTransactor) SetMarketId(opts *bind.TransactOpts, marketId string) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.contract.Transact(opts, "setMarketId", marketId)
}

// SetMarketId is a paid mutator transaction binding the contract method 0xf67b1847.
//
// Solidity: function setMarketId(string marketId) returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderSession) SetMarketId(marketId string) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.Contract.SetMarketId(&_LendingpoolAddressProvider.TransactOpts, marketId)
}

// SetMarketId is a paid mutator transaction binding the contract method 0xf67b1847.
//
// Solidity: function setMarketId(string marketId) returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderTransactorSession) SetMarketId(marketId string) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.Contract.SetMarketId(&_LendingpoolAddressProvider.TransactOpts, marketId)
}

// SetPoolAdmin is a paid mutator transaction binding the contract method 0x283d62ad.
//
// Solidity: function setPoolAdmin(address admin) returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderTransactor) SetPoolAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.contract.Transact(opts, "setPoolAdmin", admin)
}

// SetPoolAdmin is a paid mutator transaction binding the contract method 0x283d62ad.
//
// Solidity: function setPoolAdmin(address admin) returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderSession) SetPoolAdmin(admin common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.Contract.SetPoolAdmin(&_LendingpoolAddressProvider.TransactOpts, admin)
}

// SetPoolAdmin is a paid mutator transaction binding the contract method 0x283d62ad.
//
// Solidity: function setPoolAdmin(address admin) returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderTransactorSession) SetPoolAdmin(admin common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.Contract.SetPoolAdmin(&_LendingpoolAddressProvider.TransactOpts, admin)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address priceOracle) returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderTransactor) SetPriceOracle(opts *bind.TransactOpts, priceOracle common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.contract.Transact(opts, "setPriceOracle", priceOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address priceOracle) returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderSession) SetPriceOracle(priceOracle common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.Contract.SetPriceOracle(&_LendingpoolAddressProvider.TransactOpts, priceOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address priceOracle) returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderTransactorSession) SetPriceOracle(priceOracle common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.Contract.SetPriceOracle(&_LendingpoolAddressProvider.TransactOpts, priceOracle)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.Contract.TransferOwnership(&_LendingpoolAddressProvider.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LendingpoolAddressProvider *LendingpoolAddressProviderTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProvider.Contract.TransferOwnership(&_LendingpoolAddressProvider.TransactOpts, newOwner)
}

// LendingpoolAddressProviderAddressSetIterator is returned from FilterAddressSet and is used to iterate over the raw logs and unpacked data for AddressSet events raised by the LendingpoolAddressProvider contract.
type LendingpoolAddressProviderAddressSetIterator struct {
	Event *LendingpoolAddressProviderAddressSet // Event containing the contract specifics and raw log

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
func (it *LendingpoolAddressProviderAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingpoolAddressProviderAddressSet)
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
		it.Event = new(LendingpoolAddressProviderAddressSet)
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
func (it *LendingpoolAddressProviderAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingpoolAddressProviderAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingpoolAddressProviderAddressSet represents a AddressSet event raised by the LendingpoolAddressProvider contract.
type LendingpoolAddressProviderAddressSet struct {
	Id         [32]byte
	NewAddress common.Address
	HasProxy   bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddressSet is a free log retrieval operation binding the contract event 0xf2689d5d5cd0c639e137642cae5d40afced201a1a0327e7ac9358461dc9fff31.
//
// Solidity: event AddressSet(bytes32 id, address indexed newAddress, bool hasProxy)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderFilterer) FilterAddressSet(opts *bind.FilterOpts, newAddress []common.Address) (*LendingpoolAddressProviderAddressSetIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingpoolAddressProvider.contract.FilterLogs(opts, "AddressSet", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LendingpoolAddressProviderAddressSetIterator{contract: _LendingpoolAddressProvider.contract, event: "AddressSet", logs: logs, sub: sub}, nil
}

// WatchAddressSet is a free log subscription operation binding the contract event 0xf2689d5d5cd0c639e137642cae5d40afced201a1a0327e7ac9358461dc9fff31.
//
// Solidity: event AddressSet(bytes32 id, address indexed newAddress, bool hasProxy)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderFilterer) WatchAddressSet(opts *bind.WatchOpts, sink chan<- *LendingpoolAddressProviderAddressSet, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingpoolAddressProvider.contract.WatchLogs(opts, "AddressSet", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingpoolAddressProviderAddressSet)
				if err := _LendingpoolAddressProvider.contract.UnpackLog(event, "AddressSet", log); err != nil {
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

// ParseAddressSet is a log parse operation binding the contract event 0xf2689d5d5cd0c639e137642cae5d40afced201a1a0327e7ac9358461dc9fff31.
//
// Solidity: event AddressSet(bytes32 id, address indexed newAddress, bool hasProxy)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderFilterer) ParseAddressSet(log types.Log) (*LendingpoolAddressProviderAddressSet, error) {
	event := new(LendingpoolAddressProviderAddressSet)
	if err := _LendingpoolAddressProvider.contract.UnpackLog(event, "AddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingpoolAddressProviderConfigurationAdminUpdatedIterator is returned from FilterConfigurationAdminUpdated and is used to iterate over the raw logs and unpacked data for ConfigurationAdminUpdated events raised by the LendingpoolAddressProvider contract.
type LendingpoolAddressProviderConfigurationAdminUpdatedIterator struct {
	Event *LendingpoolAddressProviderConfigurationAdminUpdated // Event containing the contract specifics and raw log

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
func (it *LendingpoolAddressProviderConfigurationAdminUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingpoolAddressProviderConfigurationAdminUpdated)
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
		it.Event = new(LendingpoolAddressProviderConfigurationAdminUpdated)
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
func (it *LendingpoolAddressProviderConfigurationAdminUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingpoolAddressProviderConfigurationAdminUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingpoolAddressProviderConfigurationAdminUpdated represents a ConfigurationAdminUpdated event raised by the LendingpoolAddressProvider contract.
type LendingpoolAddressProviderConfigurationAdminUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterConfigurationAdminUpdated is a free log retrieval operation binding the contract event 0xc20a317155a9e7d84e06b716b4b355d47742ab9f8c5d630e7f556553f582430d.
//
// Solidity: event ConfigurationAdminUpdated(address indexed newAddress)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderFilterer) FilterConfigurationAdminUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*LendingpoolAddressProviderConfigurationAdminUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingpoolAddressProvider.contract.FilterLogs(opts, "ConfigurationAdminUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LendingpoolAddressProviderConfigurationAdminUpdatedIterator{contract: _LendingpoolAddressProvider.contract, event: "ConfigurationAdminUpdated", logs: logs, sub: sub}, nil
}

// WatchConfigurationAdminUpdated is a free log subscription operation binding the contract event 0xc20a317155a9e7d84e06b716b4b355d47742ab9f8c5d630e7f556553f582430d.
//
// Solidity: event ConfigurationAdminUpdated(address indexed newAddress)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderFilterer) WatchConfigurationAdminUpdated(opts *bind.WatchOpts, sink chan<- *LendingpoolAddressProviderConfigurationAdminUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingpoolAddressProvider.contract.WatchLogs(opts, "ConfigurationAdminUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingpoolAddressProviderConfigurationAdminUpdated)
				if err := _LendingpoolAddressProvider.contract.UnpackLog(event, "ConfigurationAdminUpdated", log); err != nil {
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

// ParseConfigurationAdminUpdated is a log parse operation binding the contract event 0xc20a317155a9e7d84e06b716b4b355d47742ab9f8c5d630e7f556553f582430d.
//
// Solidity: event ConfigurationAdminUpdated(address indexed newAddress)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderFilterer) ParseConfigurationAdminUpdated(log types.Log) (*LendingpoolAddressProviderConfigurationAdminUpdated, error) {
	event := new(LendingpoolAddressProviderConfigurationAdminUpdated)
	if err := _LendingpoolAddressProvider.contract.UnpackLog(event, "ConfigurationAdminUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingpoolAddressProviderEmergencyAdminUpdatedIterator is returned from FilterEmergencyAdminUpdated and is used to iterate over the raw logs and unpacked data for EmergencyAdminUpdated events raised by the LendingpoolAddressProvider contract.
type LendingpoolAddressProviderEmergencyAdminUpdatedIterator struct {
	Event *LendingpoolAddressProviderEmergencyAdminUpdated // Event containing the contract specifics and raw log

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
func (it *LendingpoolAddressProviderEmergencyAdminUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingpoolAddressProviderEmergencyAdminUpdated)
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
		it.Event = new(LendingpoolAddressProviderEmergencyAdminUpdated)
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
func (it *LendingpoolAddressProviderEmergencyAdminUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingpoolAddressProviderEmergencyAdminUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingpoolAddressProviderEmergencyAdminUpdated represents a EmergencyAdminUpdated event raised by the LendingpoolAddressProvider contract.
type LendingpoolAddressProviderEmergencyAdminUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEmergencyAdminUpdated is a free log retrieval operation binding the contract event 0xe19673fc861bfeb894cf2d6b7662505497ef31c0f489b742db24ee3310826916.
//
// Solidity: event EmergencyAdminUpdated(address indexed newAddress)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderFilterer) FilterEmergencyAdminUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*LendingpoolAddressProviderEmergencyAdminUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingpoolAddressProvider.contract.FilterLogs(opts, "EmergencyAdminUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LendingpoolAddressProviderEmergencyAdminUpdatedIterator{contract: _LendingpoolAddressProvider.contract, event: "EmergencyAdminUpdated", logs: logs, sub: sub}, nil
}

// WatchEmergencyAdminUpdated is a free log subscription operation binding the contract event 0xe19673fc861bfeb894cf2d6b7662505497ef31c0f489b742db24ee3310826916.
//
// Solidity: event EmergencyAdminUpdated(address indexed newAddress)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderFilterer) WatchEmergencyAdminUpdated(opts *bind.WatchOpts, sink chan<- *LendingpoolAddressProviderEmergencyAdminUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingpoolAddressProvider.contract.WatchLogs(opts, "EmergencyAdminUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingpoolAddressProviderEmergencyAdminUpdated)
				if err := _LendingpoolAddressProvider.contract.UnpackLog(event, "EmergencyAdminUpdated", log); err != nil {
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

// ParseEmergencyAdminUpdated is a log parse operation binding the contract event 0xe19673fc861bfeb894cf2d6b7662505497ef31c0f489b742db24ee3310826916.
//
// Solidity: event EmergencyAdminUpdated(address indexed newAddress)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderFilterer) ParseEmergencyAdminUpdated(log types.Log) (*LendingpoolAddressProviderEmergencyAdminUpdated, error) {
	event := new(LendingpoolAddressProviderEmergencyAdminUpdated)
	if err := _LendingpoolAddressProvider.contract.UnpackLog(event, "EmergencyAdminUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingpoolAddressProviderLendingPoolCollateralManagerUpdatedIterator is returned from FilterLendingPoolCollateralManagerUpdated and is used to iterate over the raw logs and unpacked data for LendingPoolCollateralManagerUpdated events raised by the LendingpoolAddressProvider contract.
type LendingpoolAddressProviderLendingPoolCollateralManagerUpdatedIterator struct {
	Event *LendingpoolAddressProviderLendingPoolCollateralManagerUpdated // Event containing the contract specifics and raw log

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
func (it *LendingpoolAddressProviderLendingPoolCollateralManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingpoolAddressProviderLendingPoolCollateralManagerUpdated)
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
		it.Event = new(LendingpoolAddressProviderLendingPoolCollateralManagerUpdated)
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
func (it *LendingpoolAddressProviderLendingPoolCollateralManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingpoolAddressProviderLendingPoolCollateralManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingpoolAddressProviderLendingPoolCollateralManagerUpdated represents a LendingPoolCollateralManagerUpdated event raised by the LendingpoolAddressProvider contract.
type LendingpoolAddressProviderLendingPoolCollateralManagerUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingPoolCollateralManagerUpdated is a free log retrieval operation binding the contract event 0x991888326f0eab3df6084aadb82bee6781b5c9aa75379e8bc50ae86934541638.
//
// Solidity: event LendingPoolCollateralManagerUpdated(address indexed newAddress)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderFilterer) FilterLendingPoolCollateralManagerUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*LendingpoolAddressProviderLendingPoolCollateralManagerUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingpoolAddressProvider.contract.FilterLogs(opts, "LendingPoolCollateralManagerUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LendingpoolAddressProviderLendingPoolCollateralManagerUpdatedIterator{contract: _LendingpoolAddressProvider.contract, event: "LendingPoolCollateralManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingPoolCollateralManagerUpdated is a free log subscription operation binding the contract event 0x991888326f0eab3df6084aadb82bee6781b5c9aa75379e8bc50ae86934541638.
//
// Solidity: event LendingPoolCollateralManagerUpdated(address indexed newAddress)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderFilterer) WatchLendingPoolCollateralManagerUpdated(opts *bind.WatchOpts, sink chan<- *LendingpoolAddressProviderLendingPoolCollateralManagerUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingpoolAddressProvider.contract.WatchLogs(opts, "LendingPoolCollateralManagerUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingpoolAddressProviderLendingPoolCollateralManagerUpdated)
				if err := _LendingpoolAddressProvider.contract.UnpackLog(event, "LendingPoolCollateralManagerUpdated", log); err != nil {
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

// ParseLendingPoolCollateralManagerUpdated is a log parse operation binding the contract event 0x991888326f0eab3df6084aadb82bee6781b5c9aa75379e8bc50ae86934541638.
//
// Solidity: event LendingPoolCollateralManagerUpdated(address indexed newAddress)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderFilterer) ParseLendingPoolCollateralManagerUpdated(log types.Log) (*LendingpoolAddressProviderLendingPoolCollateralManagerUpdated, error) {
	event := new(LendingpoolAddressProviderLendingPoolCollateralManagerUpdated)
	if err := _LendingpoolAddressProvider.contract.UnpackLog(event, "LendingPoolCollateralManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingpoolAddressProviderLendingPoolConfiguratorUpdatedIterator is returned from FilterLendingPoolConfiguratorUpdated and is used to iterate over the raw logs and unpacked data for LendingPoolConfiguratorUpdated events raised by the LendingpoolAddressProvider contract.
type LendingpoolAddressProviderLendingPoolConfiguratorUpdatedIterator struct {
	Event *LendingpoolAddressProviderLendingPoolConfiguratorUpdated // Event containing the contract specifics and raw log

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
func (it *LendingpoolAddressProviderLendingPoolConfiguratorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingpoolAddressProviderLendingPoolConfiguratorUpdated)
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
		it.Event = new(LendingpoolAddressProviderLendingPoolConfiguratorUpdated)
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
func (it *LendingpoolAddressProviderLendingPoolConfiguratorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingpoolAddressProviderLendingPoolConfiguratorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingpoolAddressProviderLendingPoolConfiguratorUpdated represents a LendingPoolConfiguratorUpdated event raised by the LendingpoolAddressProvider contract.
type LendingpoolAddressProviderLendingPoolConfiguratorUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingPoolConfiguratorUpdated is a free log retrieval operation binding the contract event 0xdfabe479bad36782fb1e77fbfddd4e382671713527e4786cfc93a022ae763729.
//
// Solidity: event LendingPoolConfiguratorUpdated(address indexed newAddress)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderFilterer) FilterLendingPoolConfiguratorUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*LendingpoolAddressProviderLendingPoolConfiguratorUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingpoolAddressProvider.contract.FilterLogs(opts, "LendingPoolConfiguratorUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LendingpoolAddressProviderLendingPoolConfiguratorUpdatedIterator{contract: _LendingpoolAddressProvider.contract, event: "LendingPoolConfiguratorUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingPoolConfiguratorUpdated is a free log subscription operation binding the contract event 0xdfabe479bad36782fb1e77fbfddd4e382671713527e4786cfc93a022ae763729.
//
// Solidity: event LendingPoolConfiguratorUpdated(address indexed newAddress)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderFilterer) WatchLendingPoolConfiguratorUpdated(opts *bind.WatchOpts, sink chan<- *LendingpoolAddressProviderLendingPoolConfiguratorUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingpoolAddressProvider.contract.WatchLogs(opts, "LendingPoolConfiguratorUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingpoolAddressProviderLendingPoolConfiguratorUpdated)
				if err := _LendingpoolAddressProvider.contract.UnpackLog(event, "LendingPoolConfiguratorUpdated", log); err != nil {
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

// ParseLendingPoolConfiguratorUpdated is a log parse operation binding the contract event 0xdfabe479bad36782fb1e77fbfddd4e382671713527e4786cfc93a022ae763729.
//
// Solidity: event LendingPoolConfiguratorUpdated(address indexed newAddress)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderFilterer) ParseLendingPoolConfiguratorUpdated(log types.Log) (*LendingpoolAddressProviderLendingPoolConfiguratorUpdated, error) {
	event := new(LendingpoolAddressProviderLendingPoolConfiguratorUpdated)
	if err := _LendingpoolAddressProvider.contract.UnpackLog(event, "LendingPoolConfiguratorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingpoolAddressProviderLendingPoolUpdatedIterator is returned from FilterLendingPoolUpdated and is used to iterate over the raw logs and unpacked data for LendingPoolUpdated events raised by the LendingpoolAddressProvider contract.
type LendingpoolAddressProviderLendingPoolUpdatedIterator struct {
	Event *LendingpoolAddressProviderLendingPoolUpdated // Event containing the contract specifics and raw log

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
func (it *LendingpoolAddressProviderLendingPoolUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingpoolAddressProviderLendingPoolUpdated)
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
		it.Event = new(LendingpoolAddressProviderLendingPoolUpdated)
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
func (it *LendingpoolAddressProviderLendingPoolUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingpoolAddressProviderLendingPoolUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingpoolAddressProviderLendingPoolUpdated represents a LendingPoolUpdated event raised by the LendingpoolAddressProvider contract.
type LendingpoolAddressProviderLendingPoolUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingPoolUpdated is a free log retrieval operation binding the contract event 0xc4e6c6cdf28d0edbd8bcf071d724d33cc2e7a30be7d06443925656e9cb492aa4.
//
// Solidity: event LendingPoolUpdated(address indexed newAddress)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderFilterer) FilterLendingPoolUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*LendingpoolAddressProviderLendingPoolUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingpoolAddressProvider.contract.FilterLogs(opts, "LendingPoolUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LendingpoolAddressProviderLendingPoolUpdatedIterator{contract: _LendingpoolAddressProvider.contract, event: "LendingPoolUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingPoolUpdated is a free log subscription operation binding the contract event 0xc4e6c6cdf28d0edbd8bcf071d724d33cc2e7a30be7d06443925656e9cb492aa4.
//
// Solidity: event LendingPoolUpdated(address indexed newAddress)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderFilterer) WatchLendingPoolUpdated(opts *bind.WatchOpts, sink chan<- *LendingpoolAddressProviderLendingPoolUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingpoolAddressProvider.contract.WatchLogs(opts, "LendingPoolUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingpoolAddressProviderLendingPoolUpdated)
				if err := _LendingpoolAddressProvider.contract.UnpackLog(event, "LendingPoolUpdated", log); err != nil {
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

// ParseLendingPoolUpdated is a log parse operation binding the contract event 0xc4e6c6cdf28d0edbd8bcf071d724d33cc2e7a30be7d06443925656e9cb492aa4.
//
// Solidity: event LendingPoolUpdated(address indexed newAddress)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderFilterer) ParseLendingPoolUpdated(log types.Log) (*LendingpoolAddressProviderLendingPoolUpdated, error) {
	event := new(LendingpoolAddressProviderLendingPoolUpdated)
	if err := _LendingpoolAddressProvider.contract.UnpackLog(event, "LendingPoolUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingpoolAddressProviderLendingRateOracleUpdatedIterator is returned from FilterLendingRateOracleUpdated and is used to iterate over the raw logs and unpacked data for LendingRateOracleUpdated events raised by the LendingpoolAddressProvider contract.
type LendingpoolAddressProviderLendingRateOracleUpdatedIterator struct {
	Event *LendingpoolAddressProviderLendingRateOracleUpdated // Event containing the contract specifics and raw log

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
func (it *LendingpoolAddressProviderLendingRateOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingpoolAddressProviderLendingRateOracleUpdated)
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
		it.Event = new(LendingpoolAddressProviderLendingRateOracleUpdated)
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
func (it *LendingpoolAddressProviderLendingRateOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingpoolAddressProviderLendingRateOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingpoolAddressProviderLendingRateOracleUpdated represents a LendingRateOracleUpdated event raised by the LendingpoolAddressProvider contract.
type LendingpoolAddressProviderLendingRateOracleUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingRateOracleUpdated is a free log retrieval operation binding the contract event 0x5c29179aba6942020a8a2d38f65de02fb6b7f784e7f049ed3a3cab97621859b5.
//
// Solidity: event LendingRateOracleUpdated(address indexed newAddress)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderFilterer) FilterLendingRateOracleUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*LendingpoolAddressProviderLendingRateOracleUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingpoolAddressProvider.contract.FilterLogs(opts, "LendingRateOracleUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LendingpoolAddressProviderLendingRateOracleUpdatedIterator{contract: _LendingpoolAddressProvider.contract, event: "LendingRateOracleUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingRateOracleUpdated is a free log subscription operation binding the contract event 0x5c29179aba6942020a8a2d38f65de02fb6b7f784e7f049ed3a3cab97621859b5.
//
// Solidity: event LendingRateOracleUpdated(address indexed newAddress)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderFilterer) WatchLendingRateOracleUpdated(opts *bind.WatchOpts, sink chan<- *LendingpoolAddressProviderLendingRateOracleUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingpoolAddressProvider.contract.WatchLogs(opts, "LendingRateOracleUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingpoolAddressProviderLendingRateOracleUpdated)
				if err := _LendingpoolAddressProvider.contract.UnpackLog(event, "LendingRateOracleUpdated", log); err != nil {
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

// ParseLendingRateOracleUpdated is a log parse operation binding the contract event 0x5c29179aba6942020a8a2d38f65de02fb6b7f784e7f049ed3a3cab97621859b5.
//
// Solidity: event LendingRateOracleUpdated(address indexed newAddress)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderFilterer) ParseLendingRateOracleUpdated(log types.Log) (*LendingpoolAddressProviderLendingRateOracleUpdated, error) {
	event := new(LendingpoolAddressProviderLendingRateOracleUpdated)
	if err := _LendingpoolAddressProvider.contract.UnpackLog(event, "LendingRateOracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingpoolAddressProviderMarketIdSetIterator is returned from FilterMarketIdSet and is used to iterate over the raw logs and unpacked data for MarketIdSet events raised by the LendingpoolAddressProvider contract.
type LendingpoolAddressProviderMarketIdSetIterator struct {
	Event *LendingpoolAddressProviderMarketIdSet // Event containing the contract specifics and raw log

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
func (it *LendingpoolAddressProviderMarketIdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingpoolAddressProviderMarketIdSet)
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
		it.Event = new(LendingpoolAddressProviderMarketIdSet)
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
func (it *LendingpoolAddressProviderMarketIdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingpoolAddressProviderMarketIdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingpoolAddressProviderMarketIdSet represents a MarketIdSet event raised by the LendingpoolAddressProvider contract.
type LendingpoolAddressProviderMarketIdSet struct {
	NewMarketId string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMarketIdSet is a free log retrieval operation binding the contract event 0x5e667c32fd847cf8bce48ab3400175cbf107bdc82b2dea62e3364909dfaee799.
//
// Solidity: event MarketIdSet(string newMarketId)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderFilterer) FilterMarketIdSet(opts *bind.FilterOpts) (*LendingpoolAddressProviderMarketIdSetIterator, error) {

	logs, sub, err := _LendingpoolAddressProvider.contract.FilterLogs(opts, "MarketIdSet")
	if err != nil {
		return nil, err
	}
	return &LendingpoolAddressProviderMarketIdSetIterator{contract: _LendingpoolAddressProvider.contract, event: "MarketIdSet", logs: logs, sub: sub}, nil
}

// WatchMarketIdSet is a free log subscription operation binding the contract event 0x5e667c32fd847cf8bce48ab3400175cbf107bdc82b2dea62e3364909dfaee799.
//
// Solidity: event MarketIdSet(string newMarketId)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderFilterer) WatchMarketIdSet(opts *bind.WatchOpts, sink chan<- *LendingpoolAddressProviderMarketIdSet) (event.Subscription, error) {

	logs, sub, err := _LendingpoolAddressProvider.contract.WatchLogs(opts, "MarketIdSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingpoolAddressProviderMarketIdSet)
				if err := _LendingpoolAddressProvider.contract.UnpackLog(event, "MarketIdSet", log); err != nil {
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

// ParseMarketIdSet is a log parse operation binding the contract event 0x5e667c32fd847cf8bce48ab3400175cbf107bdc82b2dea62e3364909dfaee799.
//
// Solidity: event MarketIdSet(string newMarketId)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderFilterer) ParseMarketIdSet(log types.Log) (*LendingpoolAddressProviderMarketIdSet, error) {
	event := new(LendingpoolAddressProviderMarketIdSet)
	if err := _LendingpoolAddressProvider.contract.UnpackLog(event, "MarketIdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingpoolAddressProviderOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LendingpoolAddressProvider contract.
type LendingpoolAddressProviderOwnershipTransferredIterator struct {
	Event *LendingpoolAddressProviderOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LendingpoolAddressProviderOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingpoolAddressProviderOwnershipTransferred)
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
		it.Event = new(LendingpoolAddressProviderOwnershipTransferred)
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
func (it *LendingpoolAddressProviderOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingpoolAddressProviderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingpoolAddressProviderOwnershipTransferred represents a OwnershipTransferred event raised by the LendingpoolAddressProvider contract.
type LendingpoolAddressProviderOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LendingpoolAddressProviderOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LendingpoolAddressProvider.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LendingpoolAddressProviderOwnershipTransferredIterator{contract: _LendingpoolAddressProvider.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LendingpoolAddressProviderOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LendingpoolAddressProvider.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingpoolAddressProviderOwnershipTransferred)
				if err := _LendingpoolAddressProvider.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderFilterer) ParseOwnershipTransferred(log types.Log) (*LendingpoolAddressProviderOwnershipTransferred, error) {
	event := new(LendingpoolAddressProviderOwnershipTransferred)
	if err := _LendingpoolAddressProvider.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingpoolAddressProviderPriceOracleUpdatedIterator is returned from FilterPriceOracleUpdated and is used to iterate over the raw logs and unpacked data for PriceOracleUpdated events raised by the LendingpoolAddressProvider contract.
type LendingpoolAddressProviderPriceOracleUpdatedIterator struct {
	Event *LendingpoolAddressProviderPriceOracleUpdated // Event containing the contract specifics and raw log

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
func (it *LendingpoolAddressProviderPriceOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingpoolAddressProviderPriceOracleUpdated)
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
		it.Event = new(LendingpoolAddressProviderPriceOracleUpdated)
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
func (it *LendingpoolAddressProviderPriceOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingpoolAddressProviderPriceOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingpoolAddressProviderPriceOracleUpdated represents a PriceOracleUpdated event raised by the LendingpoolAddressProvider contract.
type LendingpoolAddressProviderPriceOracleUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPriceOracleUpdated is a free log retrieval operation binding the contract event 0xefe8ab924ca486283a79dc604baa67add51afb82af1db8ac386ebbba643cdffd.
//
// Solidity: event PriceOracleUpdated(address indexed newAddress)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderFilterer) FilterPriceOracleUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*LendingpoolAddressProviderPriceOracleUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingpoolAddressProvider.contract.FilterLogs(opts, "PriceOracleUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LendingpoolAddressProviderPriceOracleUpdatedIterator{contract: _LendingpoolAddressProvider.contract, event: "PriceOracleUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceOracleUpdated is a free log subscription operation binding the contract event 0xefe8ab924ca486283a79dc604baa67add51afb82af1db8ac386ebbba643cdffd.
//
// Solidity: event PriceOracleUpdated(address indexed newAddress)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderFilterer) WatchPriceOracleUpdated(opts *bind.WatchOpts, sink chan<- *LendingpoolAddressProviderPriceOracleUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingpoolAddressProvider.contract.WatchLogs(opts, "PriceOracleUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingpoolAddressProviderPriceOracleUpdated)
				if err := _LendingpoolAddressProvider.contract.UnpackLog(event, "PriceOracleUpdated", log); err != nil {
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

// ParsePriceOracleUpdated is a log parse operation binding the contract event 0xefe8ab924ca486283a79dc604baa67add51afb82af1db8ac386ebbba643cdffd.
//
// Solidity: event PriceOracleUpdated(address indexed newAddress)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderFilterer) ParsePriceOracleUpdated(log types.Log) (*LendingpoolAddressProviderPriceOracleUpdated, error) {
	event := new(LendingpoolAddressProviderPriceOracleUpdated)
	if err := _LendingpoolAddressProvider.contract.UnpackLog(event, "PriceOracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingpoolAddressProviderProxyCreatedIterator is returned from FilterProxyCreated and is used to iterate over the raw logs and unpacked data for ProxyCreated events raised by the LendingpoolAddressProvider contract.
type LendingpoolAddressProviderProxyCreatedIterator struct {
	Event *LendingpoolAddressProviderProxyCreated // Event containing the contract specifics and raw log

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
func (it *LendingpoolAddressProviderProxyCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingpoolAddressProviderProxyCreated)
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
		it.Event = new(LendingpoolAddressProviderProxyCreated)
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
func (it *LendingpoolAddressProviderProxyCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingpoolAddressProviderProxyCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingpoolAddressProviderProxyCreated represents a ProxyCreated event raised by the LendingpoolAddressProvider contract.
type LendingpoolAddressProviderProxyCreated struct {
	Id         [32]byte
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProxyCreated is a free log retrieval operation binding the contract event 0x1eb35cb4b5bbb23d152f3b4016a5a46c37a07ae930ed0956aba951e231142438.
//
// Solidity: event ProxyCreated(bytes32 id, address indexed newAddress)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderFilterer) FilterProxyCreated(opts *bind.FilterOpts, newAddress []common.Address) (*LendingpoolAddressProviderProxyCreatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingpoolAddressProvider.contract.FilterLogs(opts, "ProxyCreated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LendingpoolAddressProviderProxyCreatedIterator{contract: _LendingpoolAddressProvider.contract, event: "ProxyCreated", logs: logs, sub: sub}, nil
}

// WatchProxyCreated is a free log subscription operation binding the contract event 0x1eb35cb4b5bbb23d152f3b4016a5a46c37a07ae930ed0956aba951e231142438.
//
// Solidity: event ProxyCreated(bytes32 id, address indexed newAddress)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderFilterer) WatchProxyCreated(opts *bind.WatchOpts, sink chan<- *LendingpoolAddressProviderProxyCreated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingpoolAddressProvider.contract.WatchLogs(opts, "ProxyCreated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingpoolAddressProviderProxyCreated)
				if err := _LendingpoolAddressProvider.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
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

// ParseProxyCreated is a log parse operation binding the contract event 0x1eb35cb4b5bbb23d152f3b4016a5a46c37a07ae930ed0956aba951e231142438.
//
// Solidity: event ProxyCreated(bytes32 id, address indexed newAddress)
func (_LendingpoolAddressProvider *LendingpoolAddressProviderFilterer) ParseProxyCreated(log types.Log) (*LendingpoolAddressProviderProxyCreated, error) {
	event := new(LendingpoolAddressProviderProxyCreated)
	if err := _LendingpoolAddressProvider.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
