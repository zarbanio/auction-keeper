// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package osmregistry

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

// OsmregistryMetaData contains all meta data concerning the Osmregistry contract.
var OsmregistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gem\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"osm\",\"type\":\"address\"}],\"name\":\"Add\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gem\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"osm\",\"type\":\"address\"}],\"name\":\"Remove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gem\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOsm\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOsm\",\"type\":\"address\"}],\"name\":\"Update\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gem\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"osm\",\"type\":\"address\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gem\",\"type\":\"address\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"osms\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gem\",\"type\":\"address\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gem\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOsm\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// OsmregistryABI is the input ABI used to generate the binding from.
// Deprecated: Use OsmregistryMetaData.ABI instead.
var OsmregistryABI = OsmregistryMetaData.ABI

// Osmregistry is an auto generated Go binding around an Ethereum contract.
type Osmregistry struct {
	OsmregistryCaller     // Read-only binding to the contract
	OsmregistryTransactor // Write-only binding to the contract
	OsmregistryFilterer   // Log filterer for contract events
}

// OsmregistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type OsmregistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OsmregistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OsmregistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OsmregistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OsmregistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OsmregistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OsmregistrySession struct {
	Contract     *Osmregistry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OsmregistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OsmregistryCallerSession struct {
	Contract *OsmregistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// OsmregistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OsmregistryTransactorSession struct {
	Contract     *OsmregistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// OsmregistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type OsmregistryRaw struct {
	Contract *Osmregistry // Generic contract binding to access the raw methods on
}

// OsmregistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OsmregistryCallerRaw struct {
	Contract *OsmregistryCaller // Generic read-only contract binding to access the raw methods on
}

// OsmregistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OsmregistryTransactorRaw struct {
	Contract *OsmregistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOsmregistry creates a new instance of Osmregistry, bound to a specific deployed contract.
func NewOsmregistry(address common.Address, backend bind.ContractBackend) (*Osmregistry, error) {
	contract, err := bindOsmregistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Osmregistry{OsmregistryCaller: OsmregistryCaller{contract: contract}, OsmregistryTransactor: OsmregistryTransactor{contract: contract}, OsmregistryFilterer: OsmregistryFilterer{contract: contract}}, nil
}

// NewOsmregistryCaller creates a new read-only instance of Osmregistry, bound to a specific deployed contract.
func NewOsmregistryCaller(address common.Address, caller bind.ContractCaller) (*OsmregistryCaller, error) {
	contract, err := bindOsmregistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OsmregistryCaller{contract: contract}, nil
}

// NewOsmregistryTransactor creates a new write-only instance of Osmregistry, bound to a specific deployed contract.
func NewOsmregistryTransactor(address common.Address, transactor bind.ContractTransactor) (*OsmregistryTransactor, error) {
	contract, err := bindOsmregistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OsmregistryTransactor{contract: contract}, nil
}

// NewOsmregistryFilterer creates a new log filterer instance of Osmregistry, bound to a specific deployed contract.
func NewOsmregistryFilterer(address common.Address, filterer bind.ContractFilterer) (*OsmregistryFilterer, error) {
	contract, err := bindOsmregistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OsmregistryFilterer{contract: contract}, nil
}

// bindOsmregistry binds a generic wrapper to an already deployed contract.
func bindOsmregistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OsmregistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Osmregistry *OsmregistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Osmregistry.Contract.OsmregistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Osmregistry *OsmregistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Osmregistry.Contract.OsmregistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Osmregistry *OsmregistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Osmregistry.Contract.OsmregistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Osmregistry *OsmregistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Osmregistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Osmregistry *OsmregistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Osmregistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Osmregistry *OsmregistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Osmregistry.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address gem) view returns(address, address)
func (_Osmregistry *OsmregistryCaller) Get(opts *bind.CallOpts, gem common.Address) (common.Address, common.Address, error) {
	var out []interface{}
	err := _Osmregistry.contract.Call(opts, &out, "get", gem)

	if err != nil {
		return *new(common.Address), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address gem) view returns(address, address)
func (_Osmregistry *OsmregistrySession) Get(gem common.Address) (common.Address, common.Address, error) {
	return _Osmregistry.Contract.Get(&_Osmregistry.CallOpts, gem)
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address gem) view returns(address, address)
func (_Osmregistry *OsmregistryCallerSession) Get(gem common.Address) (common.Address, common.Address, error) {
	return _Osmregistry.Contract.Get(&_Osmregistry.CallOpts, gem)
}

// Osms is a free data retrieval call binding the contract method 0xd8e7ff1b.
//
// Solidity: function osms(address ) view returns(address)
func (_Osmregistry *OsmregistryCaller) Osms(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Osmregistry.contract.Call(opts, &out, "osms", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Osms is a free data retrieval call binding the contract method 0xd8e7ff1b.
//
// Solidity: function osms(address ) view returns(address)
func (_Osmregistry *OsmregistrySession) Osms(arg0 common.Address) (common.Address, error) {
	return _Osmregistry.Contract.Osms(&_Osmregistry.CallOpts, arg0)
}

// Osms is a free data retrieval call binding the contract method 0xd8e7ff1b.
//
// Solidity: function osms(address ) view returns(address)
func (_Osmregistry *OsmregistryCallerSession) Osms(arg0 common.Address) (common.Address, error) {
	return _Osmregistry.Contract.Osms(&_Osmregistry.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Osmregistry *OsmregistryCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Osmregistry.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Osmregistry *OsmregistrySession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Osmregistry.Contract.Wards(&_Osmregistry.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Osmregistry *OsmregistryCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Osmregistry.Contract.Wards(&_Osmregistry.CallOpts, arg0)
}

// Add is a paid mutator transaction binding the contract method 0x52c28fab.
//
// Solidity: function add(address gem, address osm) returns()
func (_Osmregistry *OsmregistryTransactor) Add(opts *bind.TransactOpts, gem common.Address, osm common.Address) (*types.Transaction, error) {
	return _Osmregistry.contract.Transact(opts, "add", gem, osm)
}

// Add is a paid mutator transaction binding the contract method 0x52c28fab.
//
// Solidity: function add(address gem, address osm) returns()
func (_Osmregistry *OsmregistrySession) Add(gem common.Address, osm common.Address) (*types.Transaction, error) {
	return _Osmregistry.Contract.Add(&_Osmregistry.TransactOpts, gem, osm)
}

// Add is a paid mutator transaction binding the contract method 0x52c28fab.
//
// Solidity: function add(address gem, address osm) returns()
func (_Osmregistry *OsmregistryTransactorSession) Add(gem common.Address, osm common.Address) (*types.Transaction, error) {
	return _Osmregistry.Contract.Add(&_Osmregistry.TransactOpts, gem, osm)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Osmregistry *OsmregistryTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Osmregistry.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Osmregistry *OsmregistrySession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Osmregistry.Contract.Deny(&_Osmregistry.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Osmregistry *OsmregistryTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Osmregistry.Contract.Deny(&_Osmregistry.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Osmregistry *OsmregistryTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Osmregistry.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Osmregistry *OsmregistrySession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Osmregistry.Contract.Rely(&_Osmregistry.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Osmregistry *OsmregistryTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Osmregistry.Contract.Rely(&_Osmregistry.TransactOpts, usr)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(address gem) returns()
func (_Osmregistry *OsmregistryTransactor) Remove(opts *bind.TransactOpts, gem common.Address) (*types.Transaction, error) {
	return _Osmregistry.contract.Transact(opts, "remove", gem)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(address gem) returns()
func (_Osmregistry *OsmregistrySession) Remove(gem common.Address) (*types.Transaction, error) {
	return _Osmregistry.Contract.Remove(&_Osmregistry.TransactOpts, gem)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(address gem) returns()
func (_Osmregistry *OsmregistryTransactorSession) Remove(gem common.Address) (*types.Transaction, error) {
	return _Osmregistry.Contract.Remove(&_Osmregistry.TransactOpts, gem)
}

// Update is a paid mutator transaction binding the contract method 0xc640752d.
//
// Solidity: function update(address gem, address newOsm) returns()
func (_Osmregistry *OsmregistryTransactor) Update(opts *bind.TransactOpts, gem common.Address, newOsm common.Address) (*types.Transaction, error) {
	return _Osmregistry.contract.Transact(opts, "update", gem, newOsm)
}

// Update is a paid mutator transaction binding the contract method 0xc640752d.
//
// Solidity: function update(address gem, address newOsm) returns()
func (_Osmregistry *OsmregistrySession) Update(gem common.Address, newOsm common.Address) (*types.Transaction, error) {
	return _Osmregistry.Contract.Update(&_Osmregistry.TransactOpts, gem, newOsm)
}

// Update is a paid mutator transaction binding the contract method 0xc640752d.
//
// Solidity: function update(address gem, address newOsm) returns()
func (_Osmregistry *OsmregistryTransactorSession) Update(gem common.Address, newOsm common.Address) (*types.Transaction, error) {
	return _Osmregistry.Contract.Update(&_Osmregistry.TransactOpts, gem, newOsm)
}

// OsmregistryAddIterator is returned from FilterAdd and is used to iterate over the raw logs and unpacked data for Add events raised by the Osmregistry contract.
type OsmregistryAddIterator struct {
	Event *OsmregistryAdd // Event containing the contract specifics and raw log

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
func (it *OsmregistryAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OsmregistryAdd)
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
		it.Event = new(OsmregistryAdd)
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
func (it *OsmregistryAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OsmregistryAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OsmregistryAdd represents a Add event raised by the Osmregistry contract.
type OsmregistryAdd struct {
	Gem common.Address
	Osm common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAdd is a free log retrieval operation binding the contract event 0x473b736fe95295e8fbc851ca8acdc12a750976edad27a92f666b3d888eb895d3.
//
// Solidity: event Add(address indexed gem, address indexed osm)
func (_Osmregistry *OsmregistryFilterer) FilterAdd(opts *bind.FilterOpts, gem []common.Address, osm []common.Address) (*OsmregistryAddIterator, error) {

	var gemRule []interface{}
	for _, gemItem := range gem {
		gemRule = append(gemRule, gemItem)
	}
	var osmRule []interface{}
	for _, osmItem := range osm {
		osmRule = append(osmRule, osmItem)
	}

	logs, sub, err := _Osmregistry.contract.FilterLogs(opts, "Add", gemRule, osmRule)
	if err != nil {
		return nil, err
	}
	return &OsmregistryAddIterator{contract: _Osmregistry.contract, event: "Add", logs: logs, sub: sub}, nil
}

// WatchAdd is a free log subscription operation binding the contract event 0x473b736fe95295e8fbc851ca8acdc12a750976edad27a92f666b3d888eb895d3.
//
// Solidity: event Add(address indexed gem, address indexed osm)
func (_Osmregistry *OsmregistryFilterer) WatchAdd(opts *bind.WatchOpts, sink chan<- *OsmregistryAdd, gem []common.Address, osm []common.Address) (event.Subscription, error) {

	var gemRule []interface{}
	for _, gemItem := range gem {
		gemRule = append(gemRule, gemItem)
	}
	var osmRule []interface{}
	for _, osmItem := range osm {
		osmRule = append(osmRule, osmItem)
	}

	logs, sub, err := _Osmregistry.contract.WatchLogs(opts, "Add", gemRule, osmRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OsmregistryAdd)
				if err := _Osmregistry.contract.UnpackLog(event, "Add", log); err != nil {
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

// ParseAdd is a log parse operation binding the contract event 0x473b736fe95295e8fbc851ca8acdc12a750976edad27a92f666b3d888eb895d3.
//
// Solidity: event Add(address indexed gem, address indexed osm)
func (_Osmregistry *OsmregistryFilterer) ParseAdd(log types.Log) (*OsmregistryAdd, error) {
	event := new(OsmregistryAdd)
	if err := _Osmregistry.contract.UnpackLog(event, "Add", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OsmregistryDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the Osmregistry contract.
type OsmregistryDenyIterator struct {
	Event *OsmregistryDeny // Event containing the contract specifics and raw log

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
func (it *OsmregistryDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OsmregistryDeny)
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
		it.Event = new(OsmregistryDeny)
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
func (it *OsmregistryDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OsmregistryDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OsmregistryDeny represents a Deny event raised by the Osmregistry contract.
type OsmregistryDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_Osmregistry *OsmregistryFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*OsmregistryDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Osmregistry.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &OsmregistryDenyIterator{contract: _Osmregistry.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_Osmregistry *OsmregistryFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *OsmregistryDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Osmregistry.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OsmregistryDeny)
				if err := _Osmregistry.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_Osmregistry *OsmregistryFilterer) ParseDeny(log types.Log) (*OsmregistryDeny, error) {
	event := new(OsmregistryDeny)
	if err := _Osmregistry.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OsmregistryRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the Osmregistry contract.
type OsmregistryRelyIterator struct {
	Event *OsmregistryRely // Event containing the contract specifics and raw log

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
func (it *OsmregistryRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OsmregistryRely)
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
		it.Event = new(OsmregistryRely)
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
func (it *OsmregistryRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OsmregistryRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OsmregistryRely represents a Rely event raised by the Osmregistry contract.
type OsmregistryRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_Osmregistry *OsmregistryFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*OsmregistryRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Osmregistry.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &OsmregistryRelyIterator{contract: _Osmregistry.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_Osmregistry *OsmregistryFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *OsmregistryRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Osmregistry.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OsmregistryRely)
				if err := _Osmregistry.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_Osmregistry *OsmregistryFilterer) ParseRely(log types.Log) (*OsmregistryRely, error) {
	event := new(OsmregistryRely)
	if err := _Osmregistry.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OsmregistryRemoveIterator is returned from FilterRemove and is used to iterate over the raw logs and unpacked data for Remove events raised by the Osmregistry contract.
type OsmregistryRemoveIterator struct {
	Event *OsmregistryRemove // Event containing the contract specifics and raw log

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
func (it *OsmregistryRemoveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OsmregistryRemove)
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
		it.Event = new(OsmregistryRemove)
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
func (it *OsmregistryRemoveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OsmregistryRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OsmregistryRemove represents a Remove event raised by the Osmregistry contract.
type OsmregistryRemove struct {
	Gem common.Address
	Osm common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRemove is a free log retrieval operation binding the contract event 0xdf0aeeaee96d1d44914c04fae90b6bd2595c79e3a78f9692d4e975759cdb6e3f.
//
// Solidity: event Remove(address indexed gem, address indexed osm)
func (_Osmregistry *OsmregistryFilterer) FilterRemove(opts *bind.FilterOpts, gem []common.Address, osm []common.Address) (*OsmregistryRemoveIterator, error) {

	var gemRule []interface{}
	for _, gemItem := range gem {
		gemRule = append(gemRule, gemItem)
	}
	var osmRule []interface{}
	for _, osmItem := range osm {
		osmRule = append(osmRule, osmItem)
	}

	logs, sub, err := _Osmregistry.contract.FilterLogs(opts, "Remove", gemRule, osmRule)
	if err != nil {
		return nil, err
	}
	return &OsmregistryRemoveIterator{contract: _Osmregistry.contract, event: "Remove", logs: logs, sub: sub}, nil
}

// WatchRemove is a free log subscription operation binding the contract event 0xdf0aeeaee96d1d44914c04fae90b6bd2595c79e3a78f9692d4e975759cdb6e3f.
//
// Solidity: event Remove(address indexed gem, address indexed osm)
func (_Osmregistry *OsmregistryFilterer) WatchRemove(opts *bind.WatchOpts, sink chan<- *OsmregistryRemove, gem []common.Address, osm []common.Address) (event.Subscription, error) {

	var gemRule []interface{}
	for _, gemItem := range gem {
		gemRule = append(gemRule, gemItem)
	}
	var osmRule []interface{}
	for _, osmItem := range osm {
		osmRule = append(osmRule, osmItem)
	}

	logs, sub, err := _Osmregistry.contract.WatchLogs(opts, "Remove", gemRule, osmRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OsmregistryRemove)
				if err := _Osmregistry.contract.UnpackLog(event, "Remove", log); err != nil {
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

// ParseRemove is a log parse operation binding the contract event 0xdf0aeeaee96d1d44914c04fae90b6bd2595c79e3a78f9692d4e975759cdb6e3f.
//
// Solidity: event Remove(address indexed gem, address indexed osm)
func (_Osmregistry *OsmregistryFilterer) ParseRemove(log types.Log) (*OsmregistryRemove, error) {
	event := new(OsmregistryRemove)
	if err := _Osmregistry.contract.UnpackLog(event, "Remove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OsmregistryUpdateIterator is returned from FilterUpdate and is used to iterate over the raw logs and unpacked data for Update events raised by the Osmregistry contract.
type OsmregistryUpdateIterator struct {
	Event *OsmregistryUpdate // Event containing the contract specifics and raw log

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
func (it *OsmregistryUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OsmregistryUpdate)
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
		it.Event = new(OsmregistryUpdate)
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
func (it *OsmregistryUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OsmregistryUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OsmregistryUpdate represents a Update event raised by the Osmregistry contract.
type OsmregistryUpdate struct {
	Gem    common.Address
	OldOsm common.Address
	NewOsm common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdate is a free log retrieval operation binding the contract event 0x6e00a2c9e27b94ec0e5043bb11a037e6bda7c76b53f51e0782ec81825350882b.
//
// Solidity: event Update(address indexed gem, address indexed oldOsm, address indexed newOsm)
func (_Osmregistry *OsmregistryFilterer) FilterUpdate(opts *bind.FilterOpts, gem []common.Address, oldOsm []common.Address, newOsm []common.Address) (*OsmregistryUpdateIterator, error) {

	var gemRule []interface{}
	for _, gemItem := range gem {
		gemRule = append(gemRule, gemItem)
	}
	var oldOsmRule []interface{}
	for _, oldOsmItem := range oldOsm {
		oldOsmRule = append(oldOsmRule, oldOsmItem)
	}
	var newOsmRule []interface{}
	for _, newOsmItem := range newOsm {
		newOsmRule = append(newOsmRule, newOsmItem)
	}

	logs, sub, err := _Osmregistry.contract.FilterLogs(opts, "Update", gemRule, oldOsmRule, newOsmRule)
	if err != nil {
		return nil, err
	}
	return &OsmregistryUpdateIterator{contract: _Osmregistry.contract, event: "Update", logs: logs, sub: sub}, nil
}

// WatchUpdate is a free log subscription operation binding the contract event 0x6e00a2c9e27b94ec0e5043bb11a037e6bda7c76b53f51e0782ec81825350882b.
//
// Solidity: event Update(address indexed gem, address indexed oldOsm, address indexed newOsm)
func (_Osmregistry *OsmregistryFilterer) WatchUpdate(opts *bind.WatchOpts, sink chan<- *OsmregistryUpdate, gem []common.Address, oldOsm []common.Address, newOsm []common.Address) (event.Subscription, error) {

	var gemRule []interface{}
	for _, gemItem := range gem {
		gemRule = append(gemRule, gemItem)
	}
	var oldOsmRule []interface{}
	for _, oldOsmItem := range oldOsm {
		oldOsmRule = append(oldOsmRule, oldOsmItem)
	}
	var newOsmRule []interface{}
	for _, newOsmItem := range newOsm {
		newOsmRule = append(newOsmRule, newOsmItem)
	}

	logs, sub, err := _Osmregistry.contract.WatchLogs(opts, "Update", gemRule, oldOsmRule, newOsmRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OsmregistryUpdate)
				if err := _Osmregistry.contract.UnpackLog(event, "Update", log); err != nil {
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

// ParseUpdate is a log parse operation binding the contract event 0x6e00a2c9e27b94ec0e5043bb11a037e6bda7c76b53f51e0782ec81825350882b.
//
// Solidity: event Update(address indexed gem, address indexed oldOsm, address indexed newOsm)
func (_Osmregistry *OsmregistryFilterer) ParseUpdate(log types.Log) (*OsmregistryUpdate, error) {
	event := new(OsmregistryUpdate)
	if err := _Osmregistry.contract.UnpackLog(event, "Update", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
