// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zar_price_oracle

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

// ZarPriceOracleMetaData contains all meta data concerning the ZarPriceOracle contract.
var ZarPriceOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_decimals\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zar\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_secondsAgo\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_daiUsdPriceFeed\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"secondsAgo\",\"type\":\"uint32\"}],\"name\":\"SecondsAgoUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"baseAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daiUsdDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daiUsdPriceFeed\",\"outputs\":[{\"internalType\":\"contractIChainlinkAggregator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAssetPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"secondsAgo\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_daiUsdPriceFeed\",\"type\":\"address\"}],\"name\":\"setChainlinkDaiUsdPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_secondsAgo\",\"type\":\"uint32\"}],\"name\":\"setSecondsAgo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zar\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zarIsToken0\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ZarPriceOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use ZarPriceOracleMetaData.ABI instead.
var ZarPriceOracleABI = ZarPriceOracleMetaData.ABI

// ZarPriceOracle is an auto generated Go binding around an Ethereum contract.
type ZarPriceOracle struct {
	ZarPriceOracleCaller     // Read-only binding to the contract
	ZarPriceOracleTransactor // Write-only binding to the contract
	ZarPriceOracleFilterer   // Log filterer for contract events
}

// ZarPriceOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZarPriceOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZarPriceOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZarPriceOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZarPriceOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZarPriceOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZarPriceOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZarPriceOracleSession struct {
	Contract     *ZarPriceOracle   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZarPriceOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZarPriceOracleCallerSession struct {
	Contract *ZarPriceOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ZarPriceOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZarPriceOracleTransactorSession struct {
	Contract     *ZarPriceOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ZarPriceOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZarPriceOracleRaw struct {
	Contract *ZarPriceOracle // Generic contract binding to access the raw methods on
}

// ZarPriceOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZarPriceOracleCallerRaw struct {
	Contract *ZarPriceOracleCaller // Generic read-only contract binding to access the raw methods on
}

// ZarPriceOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZarPriceOracleTransactorRaw struct {
	Contract *ZarPriceOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZarPriceOracle creates a new instance of ZarPriceOracle, bound to a specific deployed contract.
func NewZarPriceOracle(address common.Address, backend bind.ContractBackend) (*ZarPriceOracle, error) {
	contract, err := bindZarPriceOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZarPriceOracle{ZarPriceOracleCaller: ZarPriceOracleCaller{contract: contract}, ZarPriceOracleTransactor: ZarPriceOracleTransactor{contract: contract}, ZarPriceOracleFilterer: ZarPriceOracleFilterer{contract: contract}}, nil
}

// NewZarPriceOracleCaller creates a new read-only instance of ZarPriceOracle, bound to a specific deployed contract.
func NewZarPriceOracleCaller(address common.Address, caller bind.ContractCaller) (*ZarPriceOracleCaller, error) {
	contract, err := bindZarPriceOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZarPriceOracleCaller{contract: contract}, nil
}

// NewZarPriceOracleTransactor creates a new write-only instance of ZarPriceOracle, bound to a specific deployed contract.
func NewZarPriceOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*ZarPriceOracleTransactor, error) {
	contract, err := bindZarPriceOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZarPriceOracleTransactor{contract: contract}, nil
}

// NewZarPriceOracleFilterer creates a new log filterer instance of ZarPriceOracle, bound to a specific deployed contract.
func NewZarPriceOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*ZarPriceOracleFilterer, error) {
	contract, err := bindZarPriceOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZarPriceOracleFilterer{contract: contract}, nil
}

// bindZarPriceOracle binds a generic wrapper to an already deployed contract.
func bindZarPriceOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZarPriceOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZarPriceOracle *ZarPriceOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZarPriceOracle.Contract.ZarPriceOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZarPriceOracle *ZarPriceOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZarPriceOracle.Contract.ZarPriceOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZarPriceOracle *ZarPriceOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZarPriceOracle.Contract.ZarPriceOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZarPriceOracle *ZarPriceOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZarPriceOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZarPriceOracle *ZarPriceOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZarPriceOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZarPriceOracle *ZarPriceOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZarPriceOracle.Contract.contract.Transact(opts, method, params...)
}

// BaseAmount is a free data retrieval call binding the contract method 0x4864d140.
//
// Solidity: function baseAmount() view returns(uint256)
func (_ZarPriceOracle *ZarPriceOracleCaller) BaseAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZarPriceOracle.contract.Call(opts, &out, "baseAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseAmount is a free data retrieval call binding the contract method 0x4864d140.
//
// Solidity: function baseAmount() view returns(uint256)
func (_ZarPriceOracle *ZarPriceOracleSession) BaseAmount() (*big.Int, error) {
	return _ZarPriceOracle.Contract.BaseAmount(&_ZarPriceOracle.CallOpts)
}

// BaseAmount is a free data retrieval call binding the contract method 0x4864d140.
//
// Solidity: function baseAmount() view returns(uint256)
func (_ZarPriceOracle *ZarPriceOracleCallerSession) BaseAmount() (*big.Int, error) {
	return _ZarPriceOracle.Contract.BaseAmount(&_ZarPriceOracle.CallOpts)
}

// DaiUsdDecimals is a free data retrieval call binding the contract method 0xf04f56b9.
//
// Solidity: function daiUsdDecimals() view returns(uint8)
func (_ZarPriceOracle *ZarPriceOracleCaller) DaiUsdDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ZarPriceOracle.contract.Call(opts, &out, "daiUsdDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DaiUsdDecimals is a free data retrieval call binding the contract method 0xf04f56b9.
//
// Solidity: function daiUsdDecimals() view returns(uint8)
func (_ZarPriceOracle *ZarPriceOracleSession) DaiUsdDecimals() (uint8, error) {
	return _ZarPriceOracle.Contract.DaiUsdDecimals(&_ZarPriceOracle.CallOpts)
}

// DaiUsdDecimals is a free data retrieval call binding the contract method 0xf04f56b9.
//
// Solidity: function daiUsdDecimals() view returns(uint8)
func (_ZarPriceOracle *ZarPriceOracleCallerSession) DaiUsdDecimals() (uint8, error) {
	return _ZarPriceOracle.Contract.DaiUsdDecimals(&_ZarPriceOracle.CallOpts)
}

// DaiUsdPriceFeed is a free data retrieval call binding the contract method 0x15dfe1e6.
//
// Solidity: function daiUsdPriceFeed() view returns(address)
func (_ZarPriceOracle *ZarPriceOracleCaller) DaiUsdPriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZarPriceOracle.contract.Call(opts, &out, "daiUsdPriceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaiUsdPriceFeed is a free data retrieval call binding the contract method 0x15dfe1e6.
//
// Solidity: function daiUsdPriceFeed() view returns(address)
func (_ZarPriceOracle *ZarPriceOracleSession) DaiUsdPriceFeed() (common.Address, error) {
	return _ZarPriceOracle.Contract.DaiUsdPriceFeed(&_ZarPriceOracle.CallOpts)
}

// DaiUsdPriceFeed is a free data retrieval call binding the contract method 0x15dfe1e6.
//
// Solidity: function daiUsdPriceFeed() view returns(address)
func (_ZarPriceOracle *ZarPriceOracleCallerSession) DaiUsdPriceFeed() (common.Address, error) {
	return _ZarPriceOracle.Contract.DaiUsdPriceFeed(&_ZarPriceOracle.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_ZarPriceOracle *ZarPriceOracleCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZarPriceOracle.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_ZarPriceOracle *ZarPriceOracleSession) Decimals() (*big.Int, error) {
	return _ZarPriceOracle.Contract.Decimals(&_ZarPriceOracle.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_ZarPriceOracle *ZarPriceOracleCallerSession) Decimals() (*big.Int, error) {
	return _ZarPriceOracle.Contract.Decimals(&_ZarPriceOracle.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_ZarPriceOracle *ZarPriceOracleCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZarPriceOracle.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_ZarPriceOracle *ZarPriceOracleSession) Description() (string, error) {
	return _ZarPriceOracle.Contract.Description(&_ZarPriceOracle.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_ZarPriceOracle *ZarPriceOracleCallerSession) Description() (string, error) {
	return _ZarPriceOracle.Contract.Description(&_ZarPriceOracle.CallOpts)
}

// GetAssetPrice is a free data retrieval call binding the contract method 0xe54f0880.
//
// Solidity: function getAssetPrice() view returns(uint256)
func (_ZarPriceOracle *ZarPriceOracleCaller) GetAssetPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZarPriceOracle.contract.Call(opts, &out, "getAssetPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAssetPrice is a free data retrieval call binding the contract method 0xe54f0880.
//
// Solidity: function getAssetPrice() view returns(uint256)
func (_ZarPriceOracle *ZarPriceOracleSession) GetAssetPrice() (*big.Int, error) {
	return _ZarPriceOracle.Contract.GetAssetPrice(&_ZarPriceOracle.CallOpts)
}

// GetAssetPrice is a free data retrieval call binding the contract method 0xe54f0880.
//
// Solidity: function getAssetPrice() view returns(uint256)
func (_ZarPriceOracle *ZarPriceOracleCallerSession) GetAssetPrice() (*big.Int, error) {
	return _ZarPriceOracle.Contract.GetAssetPrice(&_ZarPriceOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZarPriceOracle *ZarPriceOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZarPriceOracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZarPriceOracle *ZarPriceOracleSession) Owner() (common.Address, error) {
	return _ZarPriceOracle.Contract.Owner(&_ZarPriceOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZarPriceOracle *ZarPriceOracleCallerSession) Owner() (common.Address, error) {
	return _ZarPriceOracle.Contract.Owner(&_ZarPriceOracle.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_ZarPriceOracle *ZarPriceOracleCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZarPriceOracle.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_ZarPriceOracle *ZarPriceOracleSession) Pool() (common.Address, error) {
	return _ZarPriceOracle.Contract.Pool(&_ZarPriceOracle.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_ZarPriceOracle *ZarPriceOracleCallerSession) Pool() (common.Address, error) {
	return _ZarPriceOracle.Contract.Pool(&_ZarPriceOracle.CallOpts)
}

// SecondsAgo is a free data retrieval call binding the contract method 0x633dd145.
//
// Solidity: function secondsAgo() view returns(uint32)
func (_ZarPriceOracle *ZarPriceOracleCaller) SecondsAgo(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ZarPriceOracle.contract.Call(opts, &out, "secondsAgo")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SecondsAgo is a free data retrieval call binding the contract method 0x633dd145.
//
// Solidity: function secondsAgo() view returns(uint32)
func (_ZarPriceOracle *ZarPriceOracleSession) SecondsAgo() (uint32, error) {
	return _ZarPriceOracle.Contract.SecondsAgo(&_ZarPriceOracle.CallOpts)
}

// SecondsAgo is a free data retrieval call binding the contract method 0x633dd145.
//
// Solidity: function secondsAgo() view returns(uint32)
func (_ZarPriceOracle *ZarPriceOracleCallerSession) SecondsAgo() (uint32, error) {
	return _ZarPriceOracle.Contract.SecondsAgo(&_ZarPriceOracle.CallOpts)
}

// Zar is a free data retrieval call binding the contract method 0x7d510ab5.
//
// Solidity: function zar() view returns(address)
func (_ZarPriceOracle *ZarPriceOracleCaller) Zar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZarPriceOracle.contract.Call(opts, &out, "zar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Zar is a free data retrieval call binding the contract method 0x7d510ab5.
//
// Solidity: function zar() view returns(address)
func (_ZarPriceOracle *ZarPriceOracleSession) Zar() (common.Address, error) {
	return _ZarPriceOracle.Contract.Zar(&_ZarPriceOracle.CallOpts)
}

// Zar is a free data retrieval call binding the contract method 0x7d510ab5.
//
// Solidity: function zar() view returns(address)
func (_ZarPriceOracle *ZarPriceOracleCallerSession) Zar() (common.Address, error) {
	return _ZarPriceOracle.Contract.Zar(&_ZarPriceOracle.CallOpts)
}

// ZarIsToken0 is a free data retrieval call binding the contract method 0xe2096c90.
//
// Solidity: function zarIsToken0() view returns(bool)
func (_ZarPriceOracle *ZarPriceOracleCaller) ZarIsToken0(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZarPriceOracle.contract.Call(opts, &out, "zarIsToken0")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ZarIsToken0 is a free data retrieval call binding the contract method 0xe2096c90.
//
// Solidity: function zarIsToken0() view returns(bool)
func (_ZarPriceOracle *ZarPriceOracleSession) ZarIsToken0() (bool, error) {
	return _ZarPriceOracle.Contract.ZarIsToken0(&_ZarPriceOracle.CallOpts)
}

// ZarIsToken0 is a free data retrieval call binding the contract method 0xe2096c90.
//
// Solidity: function zarIsToken0() view returns(bool)
func (_ZarPriceOracle *ZarPriceOracleCallerSession) ZarIsToken0() (bool, error) {
	return _ZarPriceOracle.Contract.ZarIsToken0(&_ZarPriceOracle.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZarPriceOracle *ZarPriceOracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZarPriceOracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZarPriceOracle *ZarPriceOracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _ZarPriceOracle.Contract.RenounceOwnership(&_ZarPriceOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZarPriceOracle *ZarPriceOracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ZarPriceOracle.Contract.RenounceOwnership(&_ZarPriceOracle.TransactOpts)
}

// SetChainlinkDaiUsdPriceFeed is a paid mutator transaction binding the contract method 0x8b6d533d.
//
// Solidity: function setChainlinkDaiUsdPriceFeed(address _daiUsdPriceFeed) returns()
func (_ZarPriceOracle *ZarPriceOracleTransactor) SetChainlinkDaiUsdPriceFeed(opts *bind.TransactOpts, _daiUsdPriceFeed common.Address) (*types.Transaction, error) {
	return _ZarPriceOracle.contract.Transact(opts, "setChainlinkDaiUsdPriceFeed", _daiUsdPriceFeed)
}

// SetChainlinkDaiUsdPriceFeed is a paid mutator transaction binding the contract method 0x8b6d533d.
//
// Solidity: function setChainlinkDaiUsdPriceFeed(address _daiUsdPriceFeed) returns()
func (_ZarPriceOracle *ZarPriceOracleSession) SetChainlinkDaiUsdPriceFeed(_daiUsdPriceFeed common.Address) (*types.Transaction, error) {
	return _ZarPriceOracle.Contract.SetChainlinkDaiUsdPriceFeed(&_ZarPriceOracle.TransactOpts, _daiUsdPriceFeed)
}

// SetChainlinkDaiUsdPriceFeed is a paid mutator transaction binding the contract method 0x8b6d533d.
//
// Solidity: function setChainlinkDaiUsdPriceFeed(address _daiUsdPriceFeed) returns()
func (_ZarPriceOracle *ZarPriceOracleTransactorSession) SetChainlinkDaiUsdPriceFeed(_daiUsdPriceFeed common.Address) (*types.Transaction, error) {
	return _ZarPriceOracle.Contract.SetChainlinkDaiUsdPriceFeed(&_ZarPriceOracle.TransactOpts, _daiUsdPriceFeed)
}

// SetSecondsAgo is a paid mutator transaction binding the contract method 0x1ad7b127.
//
// Solidity: function setSecondsAgo(uint32 _secondsAgo) returns()
func (_ZarPriceOracle *ZarPriceOracleTransactor) SetSecondsAgo(opts *bind.TransactOpts, _secondsAgo uint32) (*types.Transaction, error) {
	return _ZarPriceOracle.contract.Transact(opts, "setSecondsAgo", _secondsAgo)
}

// SetSecondsAgo is a paid mutator transaction binding the contract method 0x1ad7b127.
//
// Solidity: function setSecondsAgo(uint32 _secondsAgo) returns()
func (_ZarPriceOracle *ZarPriceOracleSession) SetSecondsAgo(_secondsAgo uint32) (*types.Transaction, error) {
	return _ZarPriceOracle.Contract.SetSecondsAgo(&_ZarPriceOracle.TransactOpts, _secondsAgo)
}

// SetSecondsAgo is a paid mutator transaction binding the contract method 0x1ad7b127.
//
// Solidity: function setSecondsAgo(uint32 _secondsAgo) returns()
func (_ZarPriceOracle *ZarPriceOracleTransactorSession) SetSecondsAgo(_secondsAgo uint32) (*types.Transaction, error) {
	return _ZarPriceOracle.Contract.SetSecondsAgo(&_ZarPriceOracle.TransactOpts, _secondsAgo)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZarPriceOracle *ZarPriceOracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ZarPriceOracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZarPriceOracle *ZarPriceOracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ZarPriceOracle.Contract.TransferOwnership(&_ZarPriceOracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZarPriceOracle *ZarPriceOracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ZarPriceOracle.Contract.TransferOwnership(&_ZarPriceOracle.TransactOpts, newOwner)
}

// ZarPriceOracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ZarPriceOracle contract.
type ZarPriceOracleOwnershipTransferredIterator struct {
	Event *ZarPriceOracleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ZarPriceOracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZarPriceOracleOwnershipTransferred)
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
		it.Event = new(ZarPriceOracleOwnershipTransferred)
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
func (it *ZarPriceOracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZarPriceOracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZarPriceOracleOwnershipTransferred represents a OwnershipTransferred event raised by the ZarPriceOracle contract.
type ZarPriceOracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ZarPriceOracle *ZarPriceOracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ZarPriceOracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZarPriceOracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ZarPriceOracleOwnershipTransferredIterator{contract: _ZarPriceOracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ZarPriceOracle *ZarPriceOracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ZarPriceOracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZarPriceOracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZarPriceOracleOwnershipTransferred)
				if err := _ZarPriceOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ZarPriceOracle *ZarPriceOracleFilterer) ParseOwnershipTransferred(log types.Log) (*ZarPriceOracleOwnershipTransferred, error) {
	event := new(ZarPriceOracleOwnershipTransferred)
	if err := _ZarPriceOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZarPriceOracleSecondsAgoUpdatedIterator is returned from FilterSecondsAgoUpdated and is used to iterate over the raw logs and unpacked data for SecondsAgoUpdated events raised by the ZarPriceOracle contract.
type ZarPriceOracleSecondsAgoUpdatedIterator struct {
	Event *ZarPriceOracleSecondsAgoUpdated // Event containing the contract specifics and raw log

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
func (it *ZarPriceOracleSecondsAgoUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZarPriceOracleSecondsAgoUpdated)
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
		it.Event = new(ZarPriceOracleSecondsAgoUpdated)
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
func (it *ZarPriceOracleSecondsAgoUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZarPriceOracleSecondsAgoUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZarPriceOracleSecondsAgoUpdated represents a SecondsAgoUpdated event raised by the ZarPriceOracle contract.
type ZarPriceOracleSecondsAgoUpdated struct {
	SecondsAgo uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSecondsAgoUpdated is a free log retrieval operation binding the contract event 0xc36c65fb93b8b85e687c326e1c5ffdd0e8d84ee62bdfe71c9236e874f0161c3c.
//
// Solidity: event SecondsAgoUpdated(uint32 secondsAgo)
func (_ZarPriceOracle *ZarPriceOracleFilterer) FilterSecondsAgoUpdated(opts *bind.FilterOpts) (*ZarPriceOracleSecondsAgoUpdatedIterator, error) {

	logs, sub, err := _ZarPriceOracle.contract.FilterLogs(opts, "SecondsAgoUpdated")
	if err != nil {
		return nil, err
	}
	return &ZarPriceOracleSecondsAgoUpdatedIterator{contract: _ZarPriceOracle.contract, event: "SecondsAgoUpdated", logs: logs, sub: sub}, nil
}

// WatchSecondsAgoUpdated is a free log subscription operation binding the contract event 0xc36c65fb93b8b85e687c326e1c5ffdd0e8d84ee62bdfe71c9236e874f0161c3c.
//
// Solidity: event SecondsAgoUpdated(uint32 secondsAgo)
func (_ZarPriceOracle *ZarPriceOracleFilterer) WatchSecondsAgoUpdated(opts *bind.WatchOpts, sink chan<- *ZarPriceOracleSecondsAgoUpdated) (event.Subscription, error) {

	logs, sub, err := _ZarPriceOracle.contract.WatchLogs(opts, "SecondsAgoUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZarPriceOracleSecondsAgoUpdated)
				if err := _ZarPriceOracle.contract.UnpackLog(event, "SecondsAgoUpdated", log); err != nil {
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

// ParseSecondsAgoUpdated is a log parse operation binding the contract event 0xc36c65fb93b8b85e687c326e1c5ffdd0e8d84ee62bdfe71c9236e874f0161c3c.
//
// Solidity: event SecondsAgoUpdated(uint32 secondsAgo)
func (_ZarPriceOracle *ZarPriceOracleFilterer) ParseSecondsAgoUpdated(log types.Log) (*ZarPriceOracleSecondsAgoUpdated, error) {
	event := new(ZarPriceOracleSecondsAgoUpdated)
	if err := _ZarPriceOracle.contract.UnpackLog(event, "SecondsAgoUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
