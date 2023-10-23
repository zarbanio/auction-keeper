// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package proxy

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

// ProxyMetaData contains all meta data concerning the Proxy contract.
var ProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cacheAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"internalType\":\"contractDSAuthority\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cache\",\"outputs\":[{\"internalType\":\"contractDSProxyCache\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_code\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractDSAuthority\",\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cacheAddr\",\"type\":\"address\"}],\"name\":\"setCache\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use ProxyMetaData.ABI instead.
var ProxyABI = ProxyMetaData.ABI

// Proxy is an auto generated Go binding around an Ethereum contract.
type Proxy struct {
	ProxyCaller     // Read-only binding to the contract
	ProxyTransactor // Write-only binding to the contract
	ProxyFilterer   // Log filterer for contract events
}

// ProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxySession struct {
	Contract     *Proxy            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyCallerSession struct {
	Contract *ProxyCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyTransactorSession struct {
	Contract     *ProxyTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyRaw struct {
	Contract *Proxy // Generic contract binding to access the raw methods on
}

// ProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyCallerRaw struct {
	Contract *ProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyTransactorRaw struct {
	Contract *ProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxy creates a new instance of Proxy, bound to a specific deployed contract.
func NewProxy(address common.Address, backend bind.ContractBackend) (*Proxy, error) {
	contract, err := bindProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proxy{ProxyCaller: ProxyCaller{contract: contract}, ProxyTransactor: ProxyTransactor{contract: contract}, ProxyFilterer: ProxyFilterer{contract: contract}}, nil
}

// NewProxyCaller creates a new read-only instance of Proxy, bound to a specific deployed contract.
func NewProxyCaller(address common.Address, caller bind.ContractCaller) (*ProxyCaller, error) {
	contract, err := bindProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyCaller{contract: contract}, nil
}

// NewProxyTransactor creates a new write-only instance of Proxy, bound to a specific deployed contract.
func NewProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyTransactor, error) {
	contract, err := bindProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyTransactor{contract: contract}, nil
}

// NewProxyFilterer creates a new log filterer instance of Proxy, bound to a specific deployed contract.
func NewProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyFilterer, error) {
	contract, err := bindProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyFilterer{contract: contract}, nil
}

// bindProxy binds a generic wrapper to an already deployed contract.
func bindProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.ProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transact(opts, method, params...)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_Proxy *ProxyCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "authority")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_Proxy *ProxySession) Authority() (common.Address, error) {
	return _Proxy.Contract.Authority(&_Proxy.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_Proxy *ProxyCallerSession) Authority() (common.Address, error) {
	return _Proxy.Contract.Authority(&_Proxy.CallOpts)
}

// Cache is a free data retrieval call binding the contract method 0x60c7d295.
//
// Solidity: function cache() view returns(address)
func (_Proxy *ProxyCaller) Cache(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "cache")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cache is a free data retrieval call binding the contract method 0x60c7d295.
//
// Solidity: function cache() view returns(address)
func (_Proxy *ProxySession) Cache() (common.Address, error) {
	return _Proxy.Contract.Cache(&_Proxy.CallOpts)
}

// Cache is a free data retrieval call binding the contract method 0x60c7d295.
//
// Solidity: function cache() view returns(address)
func (_Proxy *ProxyCallerSession) Cache() (common.Address, error) {
	return _Proxy.Contract.Cache(&_Proxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Proxy *ProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Proxy *ProxySession) Owner() (common.Address, error) {
	return _Proxy.Contract.Owner(&_Proxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Proxy *ProxyCallerSession) Owner() (common.Address, error) {
	return _Proxy.Contract.Owner(&_Proxy.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address _target, bytes _data) payable returns(bytes response)
func (_Proxy *ProxyTransactor) Execute(opts *bind.TransactOpts, _target common.Address, _data []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "execute", _target, _data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address _target, bytes _data) payable returns(bytes response)
func (_Proxy *ProxySession) Execute(_target common.Address, _data []byte) (*types.Transaction, error) {
	return _Proxy.Contract.Execute(&_Proxy.TransactOpts, _target, _data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address _target, bytes _data) payable returns(bytes response)
func (_Proxy *ProxyTransactorSession) Execute(_target common.Address, _data []byte) (*types.Transaction, error) {
	return _Proxy.Contract.Execute(&_Proxy.TransactOpts, _target, _data)
}

// Execute0 is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _code, bytes _data) payable returns(address target, bytes response)
func (_Proxy *ProxyTransactor) Execute0(opts *bind.TransactOpts, _code []byte, _data []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "execute0", _code, _data)
}

// Execute0 is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _code, bytes _data) payable returns(address target, bytes response)
func (_Proxy *ProxySession) Execute0(_code []byte, _data []byte) (*types.Transaction, error) {
	return _Proxy.Contract.Execute0(&_Proxy.TransactOpts, _code, _data)
}

// Execute0 is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _code, bytes _data) payable returns(address target, bytes response)
func (_Proxy *ProxyTransactorSession) Execute0(_code []byte, _data []byte) (*types.Transaction, error) {
	return _Proxy.Contract.Execute0(&_Proxy.TransactOpts, _code, _data)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_Proxy *ProxyTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_Proxy *ProxySession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.SetAuthority(&_Proxy.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_Proxy *ProxyTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.SetAuthority(&_Proxy.TransactOpts, authority_)
}

// SetCache is a paid mutator transaction binding the contract method 0x948f5076.
//
// Solidity: function setCache(address _cacheAddr) returns(bool)
func (_Proxy *ProxyTransactor) SetCache(opts *bind.TransactOpts, _cacheAddr common.Address) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "setCache", _cacheAddr)
}

// SetCache is a paid mutator transaction binding the contract method 0x948f5076.
//
// Solidity: function setCache(address _cacheAddr) returns(bool)
func (_Proxy *ProxySession) SetCache(_cacheAddr common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.SetCache(&_Proxy.TransactOpts, _cacheAddr)
}

// SetCache is a paid mutator transaction binding the contract method 0x948f5076.
//
// Solidity: function setCache(address _cacheAddr) returns(bool)
func (_Proxy *ProxyTransactorSession) SetCache(_cacheAddr common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.SetCache(&_Proxy.TransactOpts, _cacheAddr)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_Proxy *ProxyTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_Proxy *ProxySession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.SetOwner(&_Proxy.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_Proxy *ProxyTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.SetOwner(&_Proxy.TransactOpts, owner_)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Proxy *ProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Proxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Proxy *ProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Proxy.Contract.Fallback(&_Proxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Proxy *ProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Proxy.Contract.Fallback(&_Proxy.TransactOpts, calldata)
}

// ProxyLogSetAuthorityIterator is returned from FilterLogSetAuthority and is used to iterate over the raw logs and unpacked data for LogSetAuthority events raised by the Proxy contract.
type ProxyLogSetAuthorityIterator struct {
	Event *ProxyLogSetAuthority // Event containing the contract specifics and raw log

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
func (it *ProxyLogSetAuthorityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyLogSetAuthority)
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
		it.Event = new(ProxyLogSetAuthority)
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
func (it *ProxyLogSetAuthorityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyLogSetAuthorityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyLogSetAuthority represents a LogSetAuthority event raised by the Proxy contract.
type ProxyLogSetAuthority struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogSetAuthority is a free log retrieval operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_Proxy *ProxyFilterer) FilterLogSetAuthority(opts *bind.FilterOpts, authority []common.Address) (*ProxyLogSetAuthorityIterator, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return &ProxyLogSetAuthorityIterator{contract: _Proxy.contract, event: "LogSetAuthority", logs: logs, sub: sub}, nil
}

// WatchLogSetAuthority is a free log subscription operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_Proxy *ProxyFilterer) WatchLogSetAuthority(opts *bind.WatchOpts, sink chan<- *ProxyLogSetAuthority, authority []common.Address) (event.Subscription, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyLogSetAuthority)
				if err := _Proxy.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
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

// ParseLogSetAuthority is a log parse operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_Proxy *ProxyFilterer) ParseLogSetAuthority(log types.Log) (*ProxyLogSetAuthority, error) {
	event := new(ProxyLogSetAuthority)
	if err := _Proxy.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyLogSetOwnerIterator is returned from FilterLogSetOwner and is used to iterate over the raw logs and unpacked data for LogSetOwner events raised by the Proxy contract.
type ProxyLogSetOwnerIterator struct {
	Event *ProxyLogSetOwner // Event containing the contract specifics and raw log

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
func (it *ProxyLogSetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyLogSetOwner)
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
		it.Event = new(ProxyLogSetOwner)
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
func (it *ProxyLogSetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyLogSetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyLogSetOwner represents a LogSetOwner event raised by the Proxy contract.
type ProxyLogSetOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogSetOwner is a free log retrieval operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_Proxy *ProxyFilterer) FilterLogSetOwner(opts *bind.FilterOpts, owner []common.Address) (*ProxyLogSetOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &ProxyLogSetOwnerIterator{contract: _Proxy.contract, event: "LogSetOwner", logs: logs, sub: sub}, nil
}

// WatchLogSetOwner is a free log subscription operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_Proxy *ProxyFilterer) WatchLogSetOwner(opts *bind.WatchOpts, sink chan<- *ProxyLogSetOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyLogSetOwner)
				if err := _Proxy.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
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

// ParseLogSetOwner is a log parse operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_Proxy *ProxyFilterer) ParseLogSetOwner(log types.Log) (*ProxyLogSetOwner, error) {
	event := new(ProxyLogSetOwner)
	if err := _Proxy.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
