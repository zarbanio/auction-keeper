// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ui_pool_data_provider

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

// IUiPoolDataProviderV2AggregatedReserveData is an auto generated low-level Go binding around an user-defined struct.
type IUiPoolDataProviderV2AggregatedReserveData struct {
	UnderlyingAsset                common.Address
	Name                           string
	Symbol                         string
	Decimals                       *big.Int
	BaseLTVasCollateral            *big.Int
	ReserveLiquidationThreshold    *big.Int
	ReserveLiquidationBonus        *big.Int
	ReserveFactor                  *big.Int
	UsageAsCollateralEnabled       bool
	BorrowingEnabled               bool
	StableBorrowRateEnabled        bool
	IsActive                       bool
	IsFrozen                       bool
	LiquidityIndex                 *big.Int
	VariableBorrowIndex            *big.Int
	LiquidityRate                  *big.Int
	VariableBorrowRate             *big.Int
	StableBorrowRate               *big.Int
	LastUpdateTimestamp            *big.Int
	ZTokenAddress                  common.Address
	StableDebtTokenAddress         common.Address
	VariableDebtTokenAddress       common.Address
	InterestRateStrategyAddress    common.Address
	AvailableLiquidity             *big.Int
	TotalPrincipalStableDebt       *big.Int
	AverageStableRate              *big.Int
	StableDebtLastUpdateTimestamp  *big.Int
	TotalScaledVariableDebt        *big.Int
	PriceInMarketReferenceCurrency *big.Int
	VariableRateSlope1             *big.Int
	VariableRateSlope2             *big.Int
	StableRateSlope1               *big.Int
	StableRateSlope2               *big.Int
}

// IUiPoolDataProviderV2BaseCurrencyInfo is an auto generated low-level Go binding around an user-defined struct.
type IUiPoolDataProviderV2BaseCurrencyInfo struct {
	MarketReferenceCurrencyUnit       *big.Int
	MarketReferenceCurrencyPriceInUsd *big.Int
	NetworkBaseTokenPriceInUsd        *big.Int
	NetworkBaseTokenPriceDecimals     uint8
}

// IUiPoolDataProviderV2UserReserveData is an auto generated low-level Go binding around an user-defined struct.
type IUiPoolDataProviderV2UserReserveData struct {
	UnderlyingAsset                 common.Address
	ScaledZTokenBalance             *big.Int
	UsageAsCollateralEnabledOnUser  bool
	StableBorrowRate                *big.Int
	ScaledVariableDebt              *big.Int
	PrincipalStableDebt             *big.Int
	StableBorrowLastUpdateTimestamp *big.Int
}

// UiPoolDataProviderMetaData contains all meta data concerning the UiPoolDataProvider contract.
var UiPoolDataProviderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIChainlinkAggregator\",\"name\":\"_networkBaseTokenPriceInUsdProxyAggregator\",\"type\":\"address\"},{\"internalType\":\"contractIChainlinkAggregator\",\"name\":\"_marketReferenceCurrencyPriceInUsdProxyAggregator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ETH_CURRENCY_UNIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_bytes32\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILendingPoolAddressesProvider\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"getReservesData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"underlyingAsset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseLTVasCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveLiquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveLiquidationBonus\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveFactor\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"borrowingEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"stableBorrowRateEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isFrozen\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"liquidityIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"variableBorrowIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"liquidityRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"variableBorrowRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"stableBorrowRate\",\"type\":\"uint128\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint40\"},{\"internalType\":\"address\",\"name\":\"zTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stableDebtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"variableDebtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestRateStrategyAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"availableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPrincipalStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"averageStableRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableDebtLastUpdateTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalScaledVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceInMarketReferenceCurrency\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableRateSlope1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableRateSlope2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableRateSlope1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableRateSlope2\",\"type\":\"uint256\"}],\"internalType\":\"structIUiPoolDataProviderV2.AggregatedReserveData[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"marketReferenceCurrencyUnit\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"marketReferenceCurrencyPriceInUsd\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"networkBaseTokenPriceInUsd\",\"type\":\"int256\"},{\"internalType\":\"uint8\",\"name\":\"networkBaseTokenPriceDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structIUiPoolDataProviderV2.BaseCurrencyInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILendingPoolAddressesProvider\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"getReservesList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILendingPoolAddressesProvider\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserReservesData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"underlyingAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"scaledZTokenBalance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabledOnUser\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"stableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scaledVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principalStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableBorrowLastUpdateTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structIUiPoolDataProviderV2.UserReserveData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketReferenceCurrencyPriceInUsdProxyAggregator\",\"outputs\":[{\"internalType\":\"contractIChainlinkAggregator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkBaseTokenPriceInUsdProxyAggregator\",\"outputs\":[{\"internalType\":\"contractIChainlinkAggregator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// UiPoolDataProviderABI is the input ABI used to generate the binding from.
// Deprecated: Use UiPoolDataProviderMetaData.ABI instead.
var UiPoolDataProviderABI = UiPoolDataProviderMetaData.ABI

// UiPoolDataProvider is an auto generated Go binding around an Ethereum contract.
type UiPoolDataProvider struct {
	UiPoolDataProviderCaller     // Read-only binding to the contract
	UiPoolDataProviderTransactor // Write-only binding to the contract
	UiPoolDataProviderFilterer   // Log filterer for contract events
}

// UiPoolDataProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type UiPoolDataProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UiPoolDataProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UiPoolDataProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UiPoolDataProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UiPoolDataProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UiPoolDataProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UiPoolDataProviderSession struct {
	Contract     *UiPoolDataProvider // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// UiPoolDataProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UiPoolDataProviderCallerSession struct {
	Contract *UiPoolDataProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// UiPoolDataProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UiPoolDataProviderTransactorSession struct {
	Contract     *UiPoolDataProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// UiPoolDataProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type UiPoolDataProviderRaw struct {
	Contract *UiPoolDataProvider // Generic contract binding to access the raw methods on
}

// UiPoolDataProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UiPoolDataProviderCallerRaw struct {
	Contract *UiPoolDataProviderCaller // Generic read-only contract binding to access the raw methods on
}

// UiPoolDataProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UiPoolDataProviderTransactorRaw struct {
	Contract *UiPoolDataProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUiPoolDataProvider creates a new instance of UiPoolDataProvider, bound to a specific deployed contract.
func NewUiPoolDataProvider(address common.Address, backend bind.ContractBackend) (*UiPoolDataProvider, error) {
	contract, err := bindUiPoolDataProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UiPoolDataProvider{UiPoolDataProviderCaller: UiPoolDataProviderCaller{contract: contract}, UiPoolDataProviderTransactor: UiPoolDataProviderTransactor{contract: contract}, UiPoolDataProviderFilterer: UiPoolDataProviderFilterer{contract: contract}}, nil
}

// NewUiPoolDataProviderCaller creates a new read-only instance of UiPoolDataProvider, bound to a specific deployed contract.
func NewUiPoolDataProviderCaller(address common.Address, caller bind.ContractCaller) (*UiPoolDataProviderCaller, error) {
	contract, err := bindUiPoolDataProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UiPoolDataProviderCaller{contract: contract}, nil
}

// NewUiPoolDataProviderTransactor creates a new write-only instance of UiPoolDataProvider, bound to a specific deployed contract.
func NewUiPoolDataProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*UiPoolDataProviderTransactor, error) {
	contract, err := bindUiPoolDataProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UiPoolDataProviderTransactor{contract: contract}, nil
}

// NewUiPoolDataProviderFilterer creates a new log filterer instance of UiPoolDataProvider, bound to a specific deployed contract.
func NewUiPoolDataProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*UiPoolDataProviderFilterer, error) {
	contract, err := bindUiPoolDataProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UiPoolDataProviderFilterer{contract: contract}, nil
}

// bindUiPoolDataProvider binds a generic wrapper to an already deployed contract.
func bindUiPoolDataProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UiPoolDataProviderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UiPoolDataProvider *UiPoolDataProviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UiPoolDataProvider.Contract.UiPoolDataProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UiPoolDataProvider *UiPoolDataProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UiPoolDataProvider.Contract.UiPoolDataProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UiPoolDataProvider *UiPoolDataProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UiPoolDataProvider.Contract.UiPoolDataProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UiPoolDataProvider *UiPoolDataProviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UiPoolDataProvider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UiPoolDataProvider *UiPoolDataProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UiPoolDataProvider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UiPoolDataProvider *UiPoolDataProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UiPoolDataProvider.Contract.contract.Transact(opts, method, params...)
}

// ETHCURRENCYUNIT is a free data retrieval call binding the contract method 0x0496f53a.
//
// Solidity: function ETH_CURRENCY_UNIT() view returns(uint256)
func (_UiPoolDataProvider *UiPoolDataProviderCaller) ETHCURRENCYUNIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UiPoolDataProvider.contract.Call(opts, &out, "ETH_CURRENCY_UNIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ETHCURRENCYUNIT is a free data retrieval call binding the contract method 0x0496f53a.
//
// Solidity: function ETH_CURRENCY_UNIT() view returns(uint256)
func (_UiPoolDataProvider *UiPoolDataProviderSession) ETHCURRENCYUNIT() (*big.Int, error) {
	return _UiPoolDataProvider.Contract.ETHCURRENCYUNIT(&_UiPoolDataProvider.CallOpts)
}

// ETHCURRENCYUNIT is a free data retrieval call binding the contract method 0x0496f53a.
//
// Solidity: function ETH_CURRENCY_UNIT() view returns(uint256)
func (_UiPoolDataProvider *UiPoolDataProviderCallerSession) ETHCURRENCYUNIT() (*big.Int, error) {
	return _UiPoolDataProvider.Contract.ETHCURRENCYUNIT(&_UiPoolDataProvider.CallOpts)
}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 _bytes32) pure returns(string)
func (_UiPoolDataProvider *UiPoolDataProviderCaller) Bytes32ToString(opts *bind.CallOpts, _bytes32 [32]byte) (string, error) {
	var out []interface{}
	err := _UiPoolDataProvider.contract.Call(opts, &out, "bytes32ToString", _bytes32)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 _bytes32) pure returns(string)
func (_UiPoolDataProvider *UiPoolDataProviderSession) Bytes32ToString(_bytes32 [32]byte) (string, error) {
	return _UiPoolDataProvider.Contract.Bytes32ToString(&_UiPoolDataProvider.CallOpts, _bytes32)
}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 _bytes32) pure returns(string)
func (_UiPoolDataProvider *UiPoolDataProviderCallerSession) Bytes32ToString(_bytes32 [32]byte) (string, error) {
	return _UiPoolDataProvider.Contract.Bytes32ToString(&_UiPoolDataProvider.CallOpts, _bytes32)
}

// GetReservesData is a free data retrieval call binding the contract method 0xec489c21.
//
// Solidity: function getReservesData(address provider) view returns((address,string,string,uint256,uint256,uint256,uint256,uint256,bool,bool,bool,bool,bool,uint128,uint128,uint128,uint128,uint128,uint40,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256)[], (uint256,int256,int256,uint8))
func (_UiPoolDataProvider *UiPoolDataProviderCaller) GetReservesData(opts *bind.CallOpts, provider common.Address) ([]IUiPoolDataProviderV2AggregatedReserveData, IUiPoolDataProviderV2BaseCurrencyInfo, error) {
	var out []interface{}
	err := _UiPoolDataProvider.contract.Call(opts, &out, "getReservesData", provider)

	if err != nil {
		return *new([]IUiPoolDataProviderV2AggregatedReserveData), *new(IUiPoolDataProviderV2BaseCurrencyInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IUiPoolDataProviderV2AggregatedReserveData)).(*[]IUiPoolDataProviderV2AggregatedReserveData)
	out1 := *abi.ConvertType(out[1], new(IUiPoolDataProviderV2BaseCurrencyInfo)).(*IUiPoolDataProviderV2BaseCurrencyInfo)

	return out0, out1, err

}

// GetReservesData is a free data retrieval call binding the contract method 0xec489c21.
//
// Solidity: function getReservesData(address provider) view returns((address,string,string,uint256,uint256,uint256,uint256,uint256,bool,bool,bool,bool,bool,uint128,uint128,uint128,uint128,uint128,uint40,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256)[], (uint256,int256,int256,uint8))
func (_UiPoolDataProvider *UiPoolDataProviderSession) GetReservesData(provider common.Address) ([]IUiPoolDataProviderV2AggregatedReserveData, IUiPoolDataProviderV2BaseCurrencyInfo, error) {
	return _UiPoolDataProvider.Contract.GetReservesData(&_UiPoolDataProvider.CallOpts, provider)
}

// GetReservesData is a free data retrieval call binding the contract method 0xec489c21.
//
// Solidity: function getReservesData(address provider) view returns((address,string,string,uint256,uint256,uint256,uint256,uint256,bool,bool,bool,bool,bool,uint128,uint128,uint128,uint128,uint128,uint40,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256)[], (uint256,int256,int256,uint8))
func (_UiPoolDataProvider *UiPoolDataProviderCallerSession) GetReservesData(provider common.Address) ([]IUiPoolDataProviderV2AggregatedReserveData, IUiPoolDataProviderV2BaseCurrencyInfo, error) {
	return _UiPoolDataProvider.Contract.GetReservesData(&_UiPoolDataProvider.CallOpts, provider)
}

// GetReservesList is a free data retrieval call binding the contract method 0x586c1442.
//
// Solidity: function getReservesList(address provider) view returns(address[])
func (_UiPoolDataProvider *UiPoolDataProviderCaller) GetReservesList(opts *bind.CallOpts, provider common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _UiPoolDataProvider.contract.Call(opts, &out, "getReservesList", provider)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetReservesList is a free data retrieval call binding the contract method 0x586c1442.
//
// Solidity: function getReservesList(address provider) view returns(address[])
func (_UiPoolDataProvider *UiPoolDataProviderSession) GetReservesList(provider common.Address) ([]common.Address, error) {
	return _UiPoolDataProvider.Contract.GetReservesList(&_UiPoolDataProvider.CallOpts, provider)
}

// GetReservesList is a free data retrieval call binding the contract method 0x586c1442.
//
// Solidity: function getReservesList(address provider) view returns(address[])
func (_UiPoolDataProvider *UiPoolDataProviderCallerSession) GetReservesList(provider common.Address) ([]common.Address, error) {
	return _UiPoolDataProvider.Contract.GetReservesList(&_UiPoolDataProvider.CallOpts, provider)
}

// GetUserReservesData is a free data retrieval call binding the contract method 0x51974cc0.
//
// Solidity: function getUserReservesData(address provider, address user) view returns((address,uint256,bool,uint256,uint256,uint256,uint256)[])
func (_UiPoolDataProvider *UiPoolDataProviderCaller) GetUserReservesData(opts *bind.CallOpts, provider common.Address, user common.Address) ([]IUiPoolDataProviderV2UserReserveData, error) {
	var out []interface{}
	err := _UiPoolDataProvider.contract.Call(opts, &out, "getUserReservesData", provider, user)

	if err != nil {
		return *new([]IUiPoolDataProviderV2UserReserveData), err
	}

	out0 := *abi.ConvertType(out[0], new([]IUiPoolDataProviderV2UserReserveData)).(*[]IUiPoolDataProviderV2UserReserveData)

	return out0, err

}

// GetUserReservesData is a free data retrieval call binding the contract method 0x51974cc0.
//
// Solidity: function getUserReservesData(address provider, address user) view returns((address,uint256,bool,uint256,uint256,uint256,uint256)[])
func (_UiPoolDataProvider *UiPoolDataProviderSession) GetUserReservesData(provider common.Address, user common.Address) ([]IUiPoolDataProviderV2UserReserveData, error) {
	return _UiPoolDataProvider.Contract.GetUserReservesData(&_UiPoolDataProvider.CallOpts, provider, user)
}

// GetUserReservesData is a free data retrieval call binding the contract method 0x51974cc0.
//
// Solidity: function getUserReservesData(address provider, address user) view returns((address,uint256,bool,uint256,uint256,uint256,uint256)[])
func (_UiPoolDataProvider *UiPoolDataProviderCallerSession) GetUserReservesData(provider common.Address, user common.Address) ([]IUiPoolDataProviderV2UserReserveData, error) {
	return _UiPoolDataProvider.Contract.GetUserReservesData(&_UiPoolDataProvider.CallOpts, provider, user)
}

// MarketReferenceCurrencyPriceInUsdProxyAggregator is a free data retrieval call binding the contract method 0xd22cf68a.
//
// Solidity: function marketReferenceCurrencyPriceInUsdProxyAggregator() view returns(address)
func (_UiPoolDataProvider *UiPoolDataProviderCaller) MarketReferenceCurrencyPriceInUsdProxyAggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UiPoolDataProvider.contract.Call(opts, &out, "marketReferenceCurrencyPriceInUsdProxyAggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketReferenceCurrencyPriceInUsdProxyAggregator is a free data retrieval call binding the contract method 0xd22cf68a.
//
// Solidity: function marketReferenceCurrencyPriceInUsdProxyAggregator() view returns(address)
func (_UiPoolDataProvider *UiPoolDataProviderSession) MarketReferenceCurrencyPriceInUsdProxyAggregator() (common.Address, error) {
	return _UiPoolDataProvider.Contract.MarketReferenceCurrencyPriceInUsdProxyAggregator(&_UiPoolDataProvider.CallOpts)
}

// MarketReferenceCurrencyPriceInUsdProxyAggregator is a free data retrieval call binding the contract method 0xd22cf68a.
//
// Solidity: function marketReferenceCurrencyPriceInUsdProxyAggregator() view returns(address)
func (_UiPoolDataProvider *UiPoolDataProviderCallerSession) MarketReferenceCurrencyPriceInUsdProxyAggregator() (common.Address, error) {
	return _UiPoolDataProvider.Contract.MarketReferenceCurrencyPriceInUsdProxyAggregator(&_UiPoolDataProvider.CallOpts)
}

// NetworkBaseTokenPriceInUsdProxyAggregator is a free data retrieval call binding the contract method 0x3c1740ed.
//
// Solidity: function networkBaseTokenPriceInUsdProxyAggregator() view returns(address)
func (_UiPoolDataProvider *UiPoolDataProviderCaller) NetworkBaseTokenPriceInUsdProxyAggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UiPoolDataProvider.contract.Call(opts, &out, "networkBaseTokenPriceInUsdProxyAggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NetworkBaseTokenPriceInUsdProxyAggregator is a free data retrieval call binding the contract method 0x3c1740ed.
//
// Solidity: function networkBaseTokenPriceInUsdProxyAggregator() view returns(address)
func (_UiPoolDataProvider *UiPoolDataProviderSession) NetworkBaseTokenPriceInUsdProxyAggregator() (common.Address, error) {
	return _UiPoolDataProvider.Contract.NetworkBaseTokenPriceInUsdProxyAggregator(&_UiPoolDataProvider.CallOpts)
}

// NetworkBaseTokenPriceInUsdProxyAggregator is a free data retrieval call binding the contract method 0x3c1740ed.
//
// Solidity: function networkBaseTokenPriceInUsdProxyAggregator() view returns(address)
func (_UiPoolDataProvider *UiPoolDataProviderCallerSession) NetworkBaseTokenPriceInUsdProxyAggregator() (common.Address, error) {
	return _UiPoolDataProvider.Contract.NetworkBaseTokenPriceInUsdProxyAggregator(&_UiPoolDataProvider.CallOpts)
}
