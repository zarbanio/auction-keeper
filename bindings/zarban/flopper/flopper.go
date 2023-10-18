// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package flopper

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

// FlopperMetaData contains all meta data concerning the Flopper contract.
var FlopperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stake_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Cage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gal\",\"type\":\"address\"}],\"name\":\"Kick\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"cage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gal\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"}],\"name\":\"kick\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"live\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stake\",\"outputs\":[{\"internalType\":\"contractStakeLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FlopperABI is the input ABI used to generate the binding from.
// Deprecated: Use FlopperMetaData.ABI instead.
var FlopperABI = FlopperMetaData.ABI

// Flopper is an auto generated Go binding around an Ethereum contract.
type Flopper struct {
	FlopperCaller     // Read-only binding to the contract
	FlopperTransactor // Write-only binding to the contract
	FlopperFilterer   // Log filterer for contract events
}

// FlopperCaller is an auto generated read-only Go binding around an Ethereum contract.
type FlopperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlopperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FlopperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlopperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FlopperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlopperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FlopperSession struct {
	Contract     *Flopper          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FlopperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FlopperCallerSession struct {
	Contract *FlopperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FlopperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FlopperTransactorSession struct {
	Contract     *FlopperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FlopperRaw is an auto generated low-level Go binding around an Ethereum contract.
type FlopperRaw struct {
	Contract *Flopper // Generic contract binding to access the raw methods on
}

// FlopperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FlopperCallerRaw struct {
	Contract *FlopperCaller // Generic read-only contract binding to access the raw methods on
}

// FlopperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FlopperTransactorRaw struct {
	Contract *FlopperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFlopper creates a new instance of Flopper, bound to a specific deployed contract.
func NewFlopper(address common.Address, backend bind.ContractBackend) (*Flopper, error) {
	contract, err := bindFlopper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Flopper{FlopperCaller: FlopperCaller{contract: contract}, FlopperTransactor: FlopperTransactor{contract: contract}, FlopperFilterer: FlopperFilterer{contract: contract}}, nil
}

// NewFlopperCaller creates a new read-only instance of Flopper, bound to a specific deployed contract.
func NewFlopperCaller(address common.Address, caller bind.ContractCaller) (*FlopperCaller, error) {
	contract, err := bindFlopper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FlopperCaller{contract: contract}, nil
}

// NewFlopperTransactor creates a new write-only instance of Flopper, bound to a specific deployed contract.
func NewFlopperTransactor(address common.Address, transactor bind.ContractTransactor) (*FlopperTransactor, error) {
	contract, err := bindFlopper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FlopperTransactor{contract: contract}, nil
}

// NewFlopperFilterer creates a new log filterer instance of Flopper, bound to a specific deployed contract.
func NewFlopperFilterer(address common.Address, filterer bind.ContractFilterer) (*FlopperFilterer, error) {
	contract, err := bindFlopper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FlopperFilterer{contract: contract}, nil
}

// bindFlopper binds a generic wrapper to an already deployed contract.
func bindFlopper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FlopperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Flopper *FlopperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Flopper.Contract.FlopperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Flopper *FlopperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Flopper.Contract.FlopperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Flopper *FlopperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Flopper.Contract.FlopperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Flopper *FlopperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Flopper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Flopper *FlopperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Flopper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Flopper *FlopperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Flopper.Contract.contract.Transact(opts, method, params...)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_Flopper *FlopperCaller) Live(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Flopper.contract.Call(opts, &out, "live")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_Flopper *FlopperSession) Live() (*big.Int, error) {
	return _Flopper.Contract.Live(&_Flopper.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_Flopper *FlopperCallerSession) Live() (*big.Int, error) {
	return _Flopper.Contract.Live(&_Flopper.CallOpts)
}

// Stake is a free data retrieval call binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() view returns(address)
func (_Flopper *FlopperCaller) Stake(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Flopper.contract.Call(opts, &out, "stake")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Stake is a free data retrieval call binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() view returns(address)
func (_Flopper *FlopperSession) Stake() (common.Address, error) {
	return _Flopper.Contract.Stake(&_Flopper.CallOpts)
}

// Stake is a free data retrieval call binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() view returns(address)
func (_Flopper *FlopperCallerSession) Stake() (common.Address, error) {
	return _Flopper.Contract.Stake(&_Flopper.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Flopper *FlopperCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Flopper.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Flopper *FlopperSession) Vat() (common.Address, error) {
	return _Flopper.Contract.Vat(&_Flopper.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Flopper *FlopperCallerSession) Vat() (common.Address, error) {
	return _Flopper.Contract.Vat(&_Flopper.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_Flopper *FlopperCaller) Vow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Flopper.contract.Call(opts, &out, "vow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_Flopper *FlopperSession) Vow() (common.Address, error) {
	return _Flopper.Contract.Vow(&_Flopper.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_Flopper *FlopperCallerSession) Vow() (common.Address, error) {
	return _Flopper.Contract.Vow(&_Flopper.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Flopper *FlopperCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Flopper.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Flopper *FlopperSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Flopper.Contract.Wards(&_Flopper.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Flopper *FlopperCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Flopper.Contract.Wards(&_Flopper.CallOpts, arg0)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Flopper *FlopperTransactor) Cage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Flopper.contract.Transact(opts, "cage")
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Flopper *FlopperSession) Cage() (*types.Transaction, error) {
	return _Flopper.Contract.Cage(&_Flopper.TransactOpts)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Flopper *FlopperTransactorSession) Cage() (*types.Transaction, error) {
	return _Flopper.Contract.Cage(&_Flopper.TransactOpts)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Flopper *FlopperTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Flopper.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Flopper *FlopperSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Flopper.Contract.Deny(&_Flopper.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Flopper *FlopperTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Flopper.Contract.Deny(&_Flopper.TransactOpts, usr)
}

// Kick is a paid mutator transaction binding the contract method 0x0203d8fb.
//
// Solidity: function kick(address gal, uint256 lot) returns()
func (_Flopper *FlopperTransactor) Kick(opts *bind.TransactOpts, gal common.Address, lot *big.Int) (*types.Transaction, error) {
	return _Flopper.contract.Transact(opts, "kick", gal, lot)
}

// Kick is a paid mutator transaction binding the contract method 0x0203d8fb.
//
// Solidity: function kick(address gal, uint256 lot) returns()
func (_Flopper *FlopperSession) Kick(gal common.Address, lot *big.Int) (*types.Transaction, error) {
	return _Flopper.Contract.Kick(&_Flopper.TransactOpts, gal, lot)
}

// Kick is a paid mutator transaction binding the contract method 0x0203d8fb.
//
// Solidity: function kick(address gal, uint256 lot) returns()
func (_Flopper *FlopperTransactorSession) Kick(gal common.Address, lot *big.Int) (*types.Transaction, error) {
	return _Flopper.Contract.Kick(&_Flopper.TransactOpts, gal, lot)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Flopper *FlopperTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Flopper.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Flopper *FlopperSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Flopper.Contract.Rely(&_Flopper.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Flopper *FlopperTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Flopper.Contract.Rely(&_Flopper.TransactOpts, usr)
}

// FlopperCageIterator is returned from FilterCage and is used to iterate over the raw logs and unpacked data for Cage events raised by the Flopper contract.
type FlopperCageIterator struct {
	Event *FlopperCage // Event containing the contract specifics and raw log

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
func (it *FlopperCageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlopperCage)
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
		it.Event = new(FlopperCage)
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
func (it *FlopperCageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlopperCageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlopperCage represents a Cage event raised by the Flopper contract.
type FlopperCage struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCage is a free log retrieval operation binding the contract event 0x2308ed18a14e800c39b86eb6ea43270105955ca385b603b64eca89f98ae8fbda.
//
// Solidity: event Cage()
func (_Flopper *FlopperFilterer) FilterCage(opts *bind.FilterOpts) (*FlopperCageIterator, error) {

	logs, sub, err := _Flopper.contract.FilterLogs(opts, "Cage")
	if err != nil {
		return nil, err
	}
	return &FlopperCageIterator{contract: _Flopper.contract, event: "Cage", logs: logs, sub: sub}, nil
}

// WatchCage is a free log subscription operation binding the contract event 0x2308ed18a14e800c39b86eb6ea43270105955ca385b603b64eca89f98ae8fbda.
//
// Solidity: event Cage()
func (_Flopper *FlopperFilterer) WatchCage(opts *bind.WatchOpts, sink chan<- *FlopperCage) (event.Subscription, error) {

	logs, sub, err := _Flopper.contract.WatchLogs(opts, "Cage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlopperCage)
				if err := _Flopper.contract.UnpackLog(event, "Cage", log); err != nil {
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
func (_Flopper *FlopperFilterer) ParseCage(log types.Log) (*FlopperCage, error) {
	event := new(FlopperCage)
	if err := _Flopper.contract.UnpackLog(event, "Cage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlopperDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the Flopper contract.
type FlopperDenyIterator struct {
	Event *FlopperDeny // Event containing the contract specifics and raw log

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
func (it *FlopperDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlopperDeny)
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
		it.Event = new(FlopperDeny)
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
func (it *FlopperDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlopperDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlopperDeny represents a Deny event raised by the Flopper contract.
type FlopperDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_Flopper *FlopperFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*FlopperDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Flopper.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &FlopperDenyIterator{contract: _Flopper.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_Flopper *FlopperFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *FlopperDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Flopper.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlopperDeny)
				if err := _Flopper.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_Flopper *FlopperFilterer) ParseDeny(log types.Log) (*FlopperDeny, error) {
	event := new(FlopperDeny)
	if err := _Flopper.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlopperKickIterator is returned from FilterKick and is used to iterate over the raw logs and unpacked data for Kick events raised by the Flopper contract.
type FlopperKickIterator struct {
	Event *FlopperKick // Event containing the contract specifics and raw log

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
func (it *FlopperKickIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlopperKick)
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
		it.Event = new(FlopperKick)
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
func (it *FlopperKickIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlopperKickIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlopperKick represents a Kick event raised by the Flopper contract.
type FlopperKick struct {
	Lot *big.Int
	Gal common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterKick is a free log retrieval operation binding the contract event 0x0a90cdede34e48dec267bc60cbbd90898951b2fc42aa23620672bc93d73d8185.
//
// Solidity: event Kick(uint256 lot, address indexed gal)
func (_Flopper *FlopperFilterer) FilterKick(opts *bind.FilterOpts, gal []common.Address) (*FlopperKickIterator, error) {

	var galRule []interface{}
	for _, galItem := range gal {
		galRule = append(galRule, galItem)
	}

	logs, sub, err := _Flopper.contract.FilterLogs(opts, "Kick", galRule)
	if err != nil {
		return nil, err
	}
	return &FlopperKickIterator{contract: _Flopper.contract, event: "Kick", logs: logs, sub: sub}, nil
}

// WatchKick is a free log subscription operation binding the contract event 0x0a90cdede34e48dec267bc60cbbd90898951b2fc42aa23620672bc93d73d8185.
//
// Solidity: event Kick(uint256 lot, address indexed gal)
func (_Flopper *FlopperFilterer) WatchKick(opts *bind.WatchOpts, sink chan<- *FlopperKick, gal []common.Address) (event.Subscription, error) {

	var galRule []interface{}
	for _, galItem := range gal {
		galRule = append(galRule, galItem)
	}

	logs, sub, err := _Flopper.contract.WatchLogs(opts, "Kick", galRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlopperKick)
				if err := _Flopper.contract.UnpackLog(event, "Kick", log); err != nil {
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

// ParseKick is a log parse operation binding the contract event 0x0a90cdede34e48dec267bc60cbbd90898951b2fc42aa23620672bc93d73d8185.
//
// Solidity: event Kick(uint256 lot, address indexed gal)
func (_Flopper *FlopperFilterer) ParseKick(log types.Log) (*FlopperKick, error) {
	event := new(FlopperKick)
	if err := _Flopper.contract.UnpackLog(event, "Kick", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlopperRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the Flopper contract.
type FlopperRelyIterator struct {
	Event *FlopperRely // Event containing the contract specifics and raw log

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
func (it *FlopperRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlopperRely)
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
		it.Event = new(FlopperRely)
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
func (it *FlopperRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlopperRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlopperRely represents a Rely event raised by the Flopper contract.
type FlopperRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_Flopper *FlopperFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*FlopperRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Flopper.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &FlopperRelyIterator{contract: _Flopper.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_Flopper *FlopperFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *FlopperRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Flopper.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlopperRely)
				if err := _Flopper.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_Flopper *FlopperFilterer) ParseRely(log types.Log) (*FlopperRely, error) {
	event := new(FlopperRely)
	if err := _Flopper.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
