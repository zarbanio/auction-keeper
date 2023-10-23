// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wallet_balance_provider

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

// WalletBalanceProviderMetaData contains all meta data concerning the WalletBalanceProvider contract.
var WalletBalanceProviderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"batchBalanceOf\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserWalletBalances\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// WalletBalanceProviderABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletBalanceProviderMetaData.ABI instead.
var WalletBalanceProviderABI = WalletBalanceProviderMetaData.ABI

// WalletBalanceProvider is an auto generated Go binding around an Ethereum contract.
type WalletBalanceProvider struct {
	WalletBalanceProviderCaller     // Read-only binding to the contract
	WalletBalanceProviderTransactor // Write-only binding to the contract
	WalletBalanceProviderFilterer   // Log filterer for contract events
}

// WalletBalanceProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletBalanceProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletBalanceProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletBalanceProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletBalanceProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletBalanceProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletBalanceProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletBalanceProviderSession struct {
	Contract     *WalletBalanceProvider // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// WalletBalanceProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletBalanceProviderCallerSession struct {
	Contract *WalletBalanceProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// WalletBalanceProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletBalanceProviderTransactorSession struct {
	Contract     *WalletBalanceProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// WalletBalanceProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletBalanceProviderRaw struct {
	Contract *WalletBalanceProvider // Generic contract binding to access the raw methods on
}

// WalletBalanceProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletBalanceProviderCallerRaw struct {
	Contract *WalletBalanceProviderCaller // Generic read-only contract binding to access the raw methods on
}

// WalletBalanceProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletBalanceProviderTransactorRaw struct {
	Contract *WalletBalanceProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletBalanceProvider creates a new instance of WalletBalanceProvider, bound to a specific deployed contract.
func NewWalletBalanceProvider(address common.Address, backend bind.ContractBackend) (*WalletBalanceProvider, error) {
	contract, err := bindWalletBalanceProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletBalanceProvider{WalletBalanceProviderCaller: WalletBalanceProviderCaller{contract: contract}, WalletBalanceProviderTransactor: WalletBalanceProviderTransactor{contract: contract}, WalletBalanceProviderFilterer: WalletBalanceProviderFilterer{contract: contract}}, nil
}

// NewWalletBalanceProviderCaller creates a new read-only instance of WalletBalanceProvider, bound to a specific deployed contract.
func NewWalletBalanceProviderCaller(address common.Address, caller bind.ContractCaller) (*WalletBalanceProviderCaller, error) {
	contract, err := bindWalletBalanceProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletBalanceProviderCaller{contract: contract}, nil
}

// NewWalletBalanceProviderTransactor creates a new write-only instance of WalletBalanceProvider, bound to a specific deployed contract.
func NewWalletBalanceProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletBalanceProviderTransactor, error) {
	contract, err := bindWalletBalanceProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletBalanceProviderTransactor{contract: contract}, nil
}

// NewWalletBalanceProviderFilterer creates a new log filterer instance of WalletBalanceProvider, bound to a specific deployed contract.
func NewWalletBalanceProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletBalanceProviderFilterer, error) {
	contract, err := bindWalletBalanceProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletBalanceProviderFilterer{contract: contract}, nil
}

// bindWalletBalanceProvider binds a generic wrapper to an already deployed contract.
func bindWalletBalanceProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletBalanceProviderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletBalanceProvider *WalletBalanceProviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletBalanceProvider.Contract.WalletBalanceProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletBalanceProvider *WalletBalanceProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletBalanceProvider.Contract.WalletBalanceProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletBalanceProvider *WalletBalanceProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletBalanceProvider.Contract.WalletBalanceProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletBalanceProvider *WalletBalanceProviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletBalanceProvider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletBalanceProvider *WalletBalanceProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletBalanceProvider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletBalanceProvider *WalletBalanceProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletBalanceProvider.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address user, address token) view returns(uint256)
func (_WalletBalanceProvider *WalletBalanceProviderCaller) BalanceOf(opts *bind.CallOpts, user common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WalletBalanceProvider.contract.Call(opts, &out, "balanceOf", user, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address user, address token) view returns(uint256)
func (_WalletBalanceProvider *WalletBalanceProviderSession) BalanceOf(user common.Address, token common.Address) (*big.Int, error) {
	return _WalletBalanceProvider.Contract.BalanceOf(&_WalletBalanceProvider.CallOpts, user, token)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address user, address token) view returns(uint256)
func (_WalletBalanceProvider *WalletBalanceProviderCallerSession) BalanceOf(user common.Address, token common.Address) (*big.Int, error) {
	return _WalletBalanceProvider.Contract.BalanceOf(&_WalletBalanceProvider.CallOpts, user, token)
}

// BatchBalanceOf is a free data retrieval call binding the contract method 0xb59b28ef.
//
// Solidity: function batchBalanceOf(address[] users, address[] tokens) view returns(uint256[])
func (_WalletBalanceProvider *WalletBalanceProviderCaller) BatchBalanceOf(opts *bind.CallOpts, users []common.Address, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _WalletBalanceProvider.contract.Call(opts, &out, "batchBalanceOf", users, tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BatchBalanceOf is a free data retrieval call binding the contract method 0xb59b28ef.
//
// Solidity: function batchBalanceOf(address[] users, address[] tokens) view returns(uint256[])
func (_WalletBalanceProvider *WalletBalanceProviderSession) BatchBalanceOf(users []common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _WalletBalanceProvider.Contract.BatchBalanceOf(&_WalletBalanceProvider.CallOpts, users, tokens)
}

// BatchBalanceOf is a free data retrieval call binding the contract method 0xb59b28ef.
//
// Solidity: function batchBalanceOf(address[] users, address[] tokens) view returns(uint256[])
func (_WalletBalanceProvider *WalletBalanceProviderCallerSession) BatchBalanceOf(users []common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _WalletBalanceProvider.Contract.BatchBalanceOf(&_WalletBalanceProvider.CallOpts, users, tokens)
}

// GetUserWalletBalances is a free data retrieval call binding the contract method 0x02405343.
//
// Solidity: function getUserWalletBalances(address provider, address user) view returns(address[], uint256[])
func (_WalletBalanceProvider *WalletBalanceProviderCaller) GetUserWalletBalances(opts *bind.CallOpts, provider common.Address, user common.Address) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _WalletBalanceProvider.contract.Call(opts, &out, "getUserWalletBalances", provider, user)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetUserWalletBalances is a free data retrieval call binding the contract method 0x02405343.
//
// Solidity: function getUserWalletBalances(address provider, address user) view returns(address[], uint256[])
func (_WalletBalanceProvider *WalletBalanceProviderSession) GetUserWalletBalances(provider common.Address, user common.Address) ([]common.Address, []*big.Int, error) {
	return _WalletBalanceProvider.Contract.GetUserWalletBalances(&_WalletBalanceProvider.CallOpts, provider, user)
}

// GetUserWalletBalances is a free data retrieval call binding the contract method 0x02405343.
//
// Solidity: function getUserWalletBalances(address provider, address user) view returns(address[], uint256[])
func (_WalletBalanceProvider *WalletBalanceProviderCallerSession) GetUserWalletBalances(provider common.Address, user common.Address) ([]common.Address, []*big.Int, error) {
	return _WalletBalanceProvider.Contract.GetUserWalletBalances(&_WalletBalanceProvider.CallOpts, provider, user)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WalletBalanceProvider *WalletBalanceProviderTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletBalanceProvider.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WalletBalanceProvider *WalletBalanceProviderSession) Receive() (*types.Transaction, error) {
	return _WalletBalanceProvider.Contract.Receive(&_WalletBalanceProvider.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WalletBalanceProvider *WalletBalanceProviderTransactorSession) Receive() (*types.Transaction, error) {
	return _WalletBalanceProvider.Contract.Receive(&_WalletBalanceProvider.TransactOpts)
}
