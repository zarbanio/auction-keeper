// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lendingpool_address_provider_registry

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

// LendingpoolAddressProviderRegistryMetaData contains all meta data concerning the LendingpoolAddressProviderRegistry contract.
var LendingpoolAddressProviderRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"AddressesProviderRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"AddressesProviderUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addressesProvider\",\"type\":\"address\"}],\"name\":\"getAddressesProviderIdByAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressesProvidersList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"registerAddressesProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"unregisterAddressesProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LendingpoolAddressProviderRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use LendingpoolAddressProviderRegistryMetaData.ABI instead.
var LendingpoolAddressProviderRegistryABI = LendingpoolAddressProviderRegistryMetaData.ABI

// LendingpoolAddressProviderRegistry is an auto generated Go binding around an Ethereum contract.
type LendingpoolAddressProviderRegistry struct {
	LendingpoolAddressProviderRegistryCaller     // Read-only binding to the contract
	LendingpoolAddressProviderRegistryTransactor // Write-only binding to the contract
	LendingpoolAddressProviderRegistryFilterer   // Log filterer for contract events
}

// LendingpoolAddressProviderRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type LendingpoolAddressProviderRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingpoolAddressProviderRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LendingpoolAddressProviderRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingpoolAddressProviderRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LendingpoolAddressProviderRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingpoolAddressProviderRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LendingpoolAddressProviderRegistrySession struct {
	Contract     *LendingpoolAddressProviderRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                       // Call options to use throughout this session
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// LendingpoolAddressProviderRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LendingpoolAddressProviderRegistryCallerSession struct {
	Contract *LendingpoolAddressProviderRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                             // Call options to use throughout this session
}

// LendingpoolAddressProviderRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LendingpoolAddressProviderRegistryTransactorSession struct {
	Contract     *LendingpoolAddressProviderRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                             // Transaction auth options to use throughout this session
}

// LendingpoolAddressProviderRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type LendingpoolAddressProviderRegistryRaw struct {
	Contract *LendingpoolAddressProviderRegistry // Generic contract binding to access the raw methods on
}

// LendingpoolAddressProviderRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LendingpoolAddressProviderRegistryCallerRaw struct {
	Contract *LendingpoolAddressProviderRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// LendingpoolAddressProviderRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LendingpoolAddressProviderRegistryTransactorRaw struct {
	Contract *LendingpoolAddressProviderRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLendingpoolAddressProviderRegistry creates a new instance of LendingpoolAddressProviderRegistry, bound to a specific deployed contract.
func NewLendingpoolAddressProviderRegistry(address common.Address, backend bind.ContractBackend) (*LendingpoolAddressProviderRegistry, error) {
	contract, err := bindLendingpoolAddressProviderRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LendingpoolAddressProviderRegistry{LendingpoolAddressProviderRegistryCaller: LendingpoolAddressProviderRegistryCaller{contract: contract}, LendingpoolAddressProviderRegistryTransactor: LendingpoolAddressProviderRegistryTransactor{contract: contract}, LendingpoolAddressProviderRegistryFilterer: LendingpoolAddressProviderRegistryFilterer{contract: contract}}, nil
}

// NewLendingpoolAddressProviderRegistryCaller creates a new read-only instance of LendingpoolAddressProviderRegistry, bound to a specific deployed contract.
func NewLendingpoolAddressProviderRegistryCaller(address common.Address, caller bind.ContractCaller) (*LendingpoolAddressProviderRegistryCaller, error) {
	contract, err := bindLendingpoolAddressProviderRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LendingpoolAddressProviderRegistryCaller{contract: contract}, nil
}

// NewLendingpoolAddressProviderRegistryTransactor creates a new write-only instance of LendingpoolAddressProviderRegistry, bound to a specific deployed contract.
func NewLendingpoolAddressProviderRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*LendingpoolAddressProviderRegistryTransactor, error) {
	contract, err := bindLendingpoolAddressProviderRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LendingpoolAddressProviderRegistryTransactor{contract: contract}, nil
}

// NewLendingpoolAddressProviderRegistryFilterer creates a new log filterer instance of LendingpoolAddressProviderRegistry, bound to a specific deployed contract.
func NewLendingpoolAddressProviderRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*LendingpoolAddressProviderRegistryFilterer, error) {
	contract, err := bindLendingpoolAddressProviderRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LendingpoolAddressProviderRegistryFilterer{contract: contract}, nil
}

// bindLendingpoolAddressProviderRegistry binds a generic wrapper to an already deployed contract.
func bindLendingpoolAddressProviderRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LendingpoolAddressProviderRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LendingpoolAddressProviderRegistry.Contract.LendingpoolAddressProviderRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingpoolAddressProviderRegistry.Contract.LendingpoolAddressProviderRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LendingpoolAddressProviderRegistry.Contract.LendingpoolAddressProviderRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LendingpoolAddressProviderRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingpoolAddressProviderRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LendingpoolAddressProviderRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetAddressesProviderIdByAddress is a free data retrieval call binding the contract method 0xd0267be7.
//
// Solidity: function getAddressesProviderIdByAddress(address addressesProvider) view returns(uint256)
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistryCaller) GetAddressesProviderIdByAddress(opts *bind.CallOpts, addressesProvider common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LendingpoolAddressProviderRegistry.contract.Call(opts, &out, "getAddressesProviderIdByAddress", addressesProvider)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAddressesProviderIdByAddress is a free data retrieval call binding the contract method 0xd0267be7.
//
// Solidity: function getAddressesProviderIdByAddress(address addressesProvider) view returns(uint256)
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistrySession) GetAddressesProviderIdByAddress(addressesProvider common.Address) (*big.Int, error) {
	return _LendingpoolAddressProviderRegistry.Contract.GetAddressesProviderIdByAddress(&_LendingpoolAddressProviderRegistry.CallOpts, addressesProvider)
}

// GetAddressesProviderIdByAddress is a free data retrieval call binding the contract method 0xd0267be7.
//
// Solidity: function getAddressesProviderIdByAddress(address addressesProvider) view returns(uint256)
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistryCallerSession) GetAddressesProviderIdByAddress(addressesProvider common.Address) (*big.Int, error) {
	return _LendingpoolAddressProviderRegistry.Contract.GetAddressesProviderIdByAddress(&_LendingpoolAddressProviderRegistry.CallOpts, addressesProvider)
}

// GetAddressesProvidersList is a free data retrieval call binding the contract method 0x365ccbbf.
//
// Solidity: function getAddressesProvidersList() view returns(address[])
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistryCaller) GetAddressesProvidersList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _LendingpoolAddressProviderRegistry.contract.Call(opts, &out, "getAddressesProvidersList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAddressesProvidersList is a free data retrieval call binding the contract method 0x365ccbbf.
//
// Solidity: function getAddressesProvidersList() view returns(address[])
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistrySession) GetAddressesProvidersList() ([]common.Address, error) {
	return _LendingpoolAddressProviderRegistry.Contract.GetAddressesProvidersList(&_LendingpoolAddressProviderRegistry.CallOpts)
}

// GetAddressesProvidersList is a free data retrieval call binding the contract method 0x365ccbbf.
//
// Solidity: function getAddressesProvidersList() view returns(address[])
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistryCallerSession) GetAddressesProvidersList() ([]common.Address, error) {
	return _LendingpoolAddressProviderRegistry.Contract.GetAddressesProvidersList(&_LendingpoolAddressProviderRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingpoolAddressProviderRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistrySession) Owner() (common.Address, error) {
	return _LendingpoolAddressProviderRegistry.Contract.Owner(&_LendingpoolAddressProviderRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistryCallerSession) Owner() (common.Address, error) {
	return _LendingpoolAddressProviderRegistry.Contract.Owner(&_LendingpoolAddressProviderRegistry.CallOpts)
}

// RegisterAddressesProvider is a paid mutator transaction binding the contract method 0xd258191e.
//
// Solidity: function registerAddressesProvider(address provider, uint256 id) returns()
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistryTransactor) RegisterAddressesProvider(opts *bind.TransactOpts, provider common.Address, id *big.Int) (*types.Transaction, error) {
	return _LendingpoolAddressProviderRegistry.contract.Transact(opts, "registerAddressesProvider", provider, id)
}

// RegisterAddressesProvider is a paid mutator transaction binding the contract method 0xd258191e.
//
// Solidity: function registerAddressesProvider(address provider, uint256 id) returns()
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistrySession) RegisterAddressesProvider(provider common.Address, id *big.Int) (*types.Transaction, error) {
	return _LendingpoolAddressProviderRegistry.Contract.RegisterAddressesProvider(&_LendingpoolAddressProviderRegistry.TransactOpts, provider, id)
}

// RegisterAddressesProvider is a paid mutator transaction binding the contract method 0xd258191e.
//
// Solidity: function registerAddressesProvider(address provider, uint256 id) returns()
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistryTransactorSession) RegisterAddressesProvider(provider common.Address, id *big.Int) (*types.Transaction, error) {
	return _LendingpoolAddressProviderRegistry.Contract.RegisterAddressesProvider(&_LendingpoolAddressProviderRegistry.TransactOpts, provider, id)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingpoolAddressProviderRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _LendingpoolAddressProviderRegistry.Contract.RenounceOwnership(&_LendingpoolAddressProviderRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LendingpoolAddressProviderRegistry.Contract.RenounceOwnership(&_LendingpoolAddressProviderRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProviderRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProviderRegistry.Contract.TransferOwnership(&_LendingpoolAddressProviderRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProviderRegistry.Contract.TransferOwnership(&_LendingpoolAddressProviderRegistry.TransactOpts, newOwner)
}

// UnregisterAddressesProvider is a paid mutator transaction binding the contract method 0x0de26707.
//
// Solidity: function unregisterAddressesProvider(address provider) returns()
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistryTransactor) UnregisterAddressesProvider(opts *bind.TransactOpts, provider common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProviderRegistry.contract.Transact(opts, "unregisterAddressesProvider", provider)
}

// UnregisterAddressesProvider is a paid mutator transaction binding the contract method 0x0de26707.
//
// Solidity: function unregisterAddressesProvider(address provider) returns()
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistrySession) UnregisterAddressesProvider(provider common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProviderRegistry.Contract.UnregisterAddressesProvider(&_LendingpoolAddressProviderRegistry.TransactOpts, provider)
}

// UnregisterAddressesProvider is a paid mutator transaction binding the contract method 0x0de26707.
//
// Solidity: function unregisterAddressesProvider(address provider) returns()
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistryTransactorSession) UnregisterAddressesProvider(provider common.Address) (*types.Transaction, error) {
	return _LendingpoolAddressProviderRegistry.Contract.UnregisterAddressesProvider(&_LendingpoolAddressProviderRegistry.TransactOpts, provider)
}

// LendingpoolAddressProviderRegistryAddressesProviderRegisteredIterator is returned from FilterAddressesProviderRegistered and is used to iterate over the raw logs and unpacked data for AddressesProviderRegistered events raised by the LendingpoolAddressProviderRegistry contract.
type LendingpoolAddressProviderRegistryAddressesProviderRegisteredIterator struct {
	Event *LendingpoolAddressProviderRegistryAddressesProviderRegistered // Event containing the contract specifics and raw log

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
func (it *LendingpoolAddressProviderRegistryAddressesProviderRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingpoolAddressProviderRegistryAddressesProviderRegistered)
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
		it.Event = new(LendingpoolAddressProviderRegistryAddressesProviderRegistered)
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
func (it *LendingpoolAddressProviderRegistryAddressesProviderRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingpoolAddressProviderRegistryAddressesProviderRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingpoolAddressProviderRegistryAddressesProviderRegistered represents a AddressesProviderRegistered event raised by the LendingpoolAddressProviderRegistry contract.
type LendingpoolAddressProviderRegistryAddressesProviderRegistered struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddressesProviderRegistered is a free log retrieval operation binding the contract event 0x2db38786c10176b033a1608361716b0ca992e3af55dc05b6dc710969790beeda.
//
// Solidity: event AddressesProviderRegistered(address indexed newAddress)
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistryFilterer) FilterAddressesProviderRegistered(opts *bind.FilterOpts, newAddress []common.Address) (*LendingpoolAddressProviderRegistryAddressesProviderRegisteredIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingpoolAddressProviderRegistry.contract.FilterLogs(opts, "AddressesProviderRegistered", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LendingpoolAddressProviderRegistryAddressesProviderRegisteredIterator{contract: _LendingpoolAddressProviderRegistry.contract, event: "AddressesProviderRegistered", logs: logs, sub: sub}, nil
}

// WatchAddressesProviderRegistered is a free log subscription operation binding the contract event 0x2db38786c10176b033a1608361716b0ca992e3af55dc05b6dc710969790beeda.
//
// Solidity: event AddressesProviderRegistered(address indexed newAddress)
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistryFilterer) WatchAddressesProviderRegistered(opts *bind.WatchOpts, sink chan<- *LendingpoolAddressProviderRegistryAddressesProviderRegistered, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingpoolAddressProviderRegistry.contract.WatchLogs(opts, "AddressesProviderRegistered", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingpoolAddressProviderRegistryAddressesProviderRegistered)
				if err := _LendingpoolAddressProviderRegistry.contract.UnpackLog(event, "AddressesProviderRegistered", log); err != nil {
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

// ParseAddressesProviderRegistered is a log parse operation binding the contract event 0x2db38786c10176b033a1608361716b0ca992e3af55dc05b6dc710969790beeda.
//
// Solidity: event AddressesProviderRegistered(address indexed newAddress)
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistryFilterer) ParseAddressesProviderRegistered(log types.Log) (*LendingpoolAddressProviderRegistryAddressesProviderRegistered, error) {
	event := new(LendingpoolAddressProviderRegistryAddressesProviderRegistered)
	if err := _LendingpoolAddressProviderRegistry.contract.UnpackLog(event, "AddressesProviderRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingpoolAddressProviderRegistryAddressesProviderUnregisteredIterator is returned from FilterAddressesProviderUnregistered and is used to iterate over the raw logs and unpacked data for AddressesProviderUnregistered events raised by the LendingpoolAddressProviderRegistry contract.
type LendingpoolAddressProviderRegistryAddressesProviderUnregisteredIterator struct {
	Event *LendingpoolAddressProviderRegistryAddressesProviderUnregistered // Event containing the contract specifics and raw log

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
func (it *LendingpoolAddressProviderRegistryAddressesProviderUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingpoolAddressProviderRegistryAddressesProviderUnregistered)
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
		it.Event = new(LendingpoolAddressProviderRegistryAddressesProviderUnregistered)
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
func (it *LendingpoolAddressProviderRegistryAddressesProviderUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingpoolAddressProviderRegistryAddressesProviderUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingpoolAddressProviderRegistryAddressesProviderUnregistered represents a AddressesProviderUnregistered event raised by the LendingpoolAddressProviderRegistry contract.
type LendingpoolAddressProviderRegistryAddressesProviderUnregistered struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddressesProviderUnregistered is a free log retrieval operation binding the contract event 0x851e5971c053e6b76e3a1e0b8ffa81430df738007fad86e195c409a757faccd2.
//
// Solidity: event AddressesProviderUnregistered(address indexed newAddress)
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistryFilterer) FilterAddressesProviderUnregistered(opts *bind.FilterOpts, newAddress []common.Address) (*LendingpoolAddressProviderRegistryAddressesProviderUnregisteredIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingpoolAddressProviderRegistry.contract.FilterLogs(opts, "AddressesProviderUnregistered", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LendingpoolAddressProviderRegistryAddressesProviderUnregisteredIterator{contract: _LendingpoolAddressProviderRegistry.contract, event: "AddressesProviderUnregistered", logs: logs, sub: sub}, nil
}

// WatchAddressesProviderUnregistered is a free log subscription operation binding the contract event 0x851e5971c053e6b76e3a1e0b8ffa81430df738007fad86e195c409a757faccd2.
//
// Solidity: event AddressesProviderUnregistered(address indexed newAddress)
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistryFilterer) WatchAddressesProviderUnregistered(opts *bind.WatchOpts, sink chan<- *LendingpoolAddressProviderRegistryAddressesProviderUnregistered, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingpoolAddressProviderRegistry.contract.WatchLogs(opts, "AddressesProviderUnregistered", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingpoolAddressProviderRegistryAddressesProviderUnregistered)
				if err := _LendingpoolAddressProviderRegistry.contract.UnpackLog(event, "AddressesProviderUnregistered", log); err != nil {
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

// ParseAddressesProviderUnregistered is a log parse operation binding the contract event 0x851e5971c053e6b76e3a1e0b8ffa81430df738007fad86e195c409a757faccd2.
//
// Solidity: event AddressesProviderUnregistered(address indexed newAddress)
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistryFilterer) ParseAddressesProviderUnregistered(log types.Log) (*LendingpoolAddressProviderRegistryAddressesProviderUnregistered, error) {
	event := new(LendingpoolAddressProviderRegistryAddressesProviderUnregistered)
	if err := _LendingpoolAddressProviderRegistry.contract.UnpackLog(event, "AddressesProviderUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingpoolAddressProviderRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LendingpoolAddressProviderRegistry contract.
type LendingpoolAddressProviderRegistryOwnershipTransferredIterator struct {
	Event *LendingpoolAddressProviderRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LendingpoolAddressProviderRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingpoolAddressProviderRegistryOwnershipTransferred)
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
		it.Event = new(LendingpoolAddressProviderRegistryOwnershipTransferred)
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
func (it *LendingpoolAddressProviderRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingpoolAddressProviderRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingpoolAddressProviderRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the LendingpoolAddressProviderRegistry contract.
type LendingpoolAddressProviderRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LendingpoolAddressProviderRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LendingpoolAddressProviderRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LendingpoolAddressProviderRegistryOwnershipTransferredIterator{contract: _LendingpoolAddressProviderRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LendingpoolAddressProviderRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LendingpoolAddressProviderRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingpoolAddressProviderRegistryOwnershipTransferred)
				if err := _LendingpoolAddressProviderRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LendingpoolAddressProviderRegistry *LendingpoolAddressProviderRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*LendingpoolAddressProviderRegistryOwnershipTransferred, error) {
	event := new(LendingpoolAddressProviderRegistryOwnershipTransferred)
	if err := _LendingpoolAddressProviderRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
