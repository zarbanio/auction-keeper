// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gemjoin

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

// GemjoinMetaData contains all meta data concerning the Gemjoin contract.
var GemjoinMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"ilk_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"gem_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Cage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Exit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Join\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"cage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gem\",\"outputs\":[{\"internalType\":\"contractGemLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ilk\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"join\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"live\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// GemjoinABI is the input ABI used to generate the binding from.
// Deprecated: Use GemjoinMetaData.ABI instead.
var GemjoinABI = GemjoinMetaData.ABI

// Gemjoin is an auto generated Go binding around an Ethereum contract.
type Gemjoin struct {
	GemjoinCaller     // Read-only binding to the contract
	GemjoinTransactor // Write-only binding to the contract
	GemjoinFilterer   // Log filterer for contract events
}

// GemjoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type GemjoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GemjoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GemjoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GemjoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GemjoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GemjoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GemjoinSession struct {
	Contract     *Gemjoin          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GemjoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GemjoinCallerSession struct {
	Contract *GemjoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// GemjoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GemjoinTransactorSession struct {
	Contract     *GemjoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GemjoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type GemjoinRaw struct {
	Contract *Gemjoin // Generic contract binding to access the raw methods on
}

// GemjoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GemjoinCallerRaw struct {
	Contract *GemjoinCaller // Generic read-only contract binding to access the raw methods on
}

// GemjoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GemjoinTransactorRaw struct {
	Contract *GemjoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGemjoin creates a new instance of Gemjoin, bound to a specific deployed contract.
func NewGemjoin(address common.Address, backend bind.ContractBackend) (*Gemjoin, error) {
	contract, err := bindGemjoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gemjoin{GemjoinCaller: GemjoinCaller{contract: contract}, GemjoinTransactor: GemjoinTransactor{contract: contract}, GemjoinFilterer: GemjoinFilterer{contract: contract}}, nil
}

// NewGemjoinCaller creates a new read-only instance of Gemjoin, bound to a specific deployed contract.
func NewGemjoinCaller(address common.Address, caller bind.ContractCaller) (*GemjoinCaller, error) {
	contract, err := bindGemjoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GemjoinCaller{contract: contract}, nil
}

// NewGemjoinTransactor creates a new write-only instance of Gemjoin, bound to a specific deployed contract.
func NewGemjoinTransactor(address common.Address, transactor bind.ContractTransactor) (*GemjoinTransactor, error) {
	contract, err := bindGemjoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GemjoinTransactor{contract: contract}, nil
}

// NewGemjoinFilterer creates a new log filterer instance of Gemjoin, bound to a specific deployed contract.
func NewGemjoinFilterer(address common.Address, filterer bind.ContractFilterer) (*GemjoinFilterer, error) {
	contract, err := bindGemjoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GemjoinFilterer{contract: contract}, nil
}

// bindGemjoin binds a generic wrapper to an already deployed contract.
func bindGemjoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GemjoinMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gemjoin *GemjoinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gemjoin.Contract.GemjoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gemjoin *GemjoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gemjoin.Contract.GemjoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gemjoin *GemjoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gemjoin.Contract.GemjoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gemjoin *GemjoinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gemjoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gemjoin *GemjoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gemjoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gemjoin *GemjoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gemjoin.Contract.contract.Transact(opts, method, params...)
}

// Dec is a free data retrieval call binding the contract method 0xb3bcfa82.
//
// Solidity: function dec() view returns(uint256)
func (_Gemjoin *GemjoinCaller) Dec(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gemjoin.contract.Call(opts, &out, "dec")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dec is a free data retrieval call binding the contract method 0xb3bcfa82.
//
// Solidity: function dec() view returns(uint256)
func (_Gemjoin *GemjoinSession) Dec() (*big.Int, error) {
	return _Gemjoin.Contract.Dec(&_Gemjoin.CallOpts)
}

// Dec is a free data retrieval call binding the contract method 0xb3bcfa82.
//
// Solidity: function dec() view returns(uint256)
func (_Gemjoin *GemjoinCallerSession) Dec() (*big.Int, error) {
	return _Gemjoin.Contract.Dec(&_Gemjoin.CallOpts)
}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() view returns(address)
func (_Gemjoin *GemjoinCaller) Gem(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gemjoin.contract.Call(opts, &out, "gem")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() view returns(address)
func (_Gemjoin *GemjoinSession) Gem() (common.Address, error) {
	return _Gemjoin.Contract.Gem(&_Gemjoin.CallOpts)
}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() view returns(address)
func (_Gemjoin *GemjoinCallerSession) Gem() (common.Address, error) {
	return _Gemjoin.Contract.Gem(&_Gemjoin.CallOpts)
}

// Ilk is a free data retrieval call binding the contract method 0xc5ce281e.
//
// Solidity: function ilk() view returns(bytes32)
func (_Gemjoin *GemjoinCaller) Ilk(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Gemjoin.contract.Call(opts, &out, "ilk")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Ilk is a free data retrieval call binding the contract method 0xc5ce281e.
//
// Solidity: function ilk() view returns(bytes32)
func (_Gemjoin *GemjoinSession) Ilk() ([32]byte, error) {
	return _Gemjoin.Contract.Ilk(&_Gemjoin.CallOpts)
}

// Ilk is a free data retrieval call binding the contract method 0xc5ce281e.
//
// Solidity: function ilk() view returns(bytes32)
func (_Gemjoin *GemjoinCallerSession) Ilk() ([32]byte, error) {
	return _Gemjoin.Contract.Ilk(&_Gemjoin.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_Gemjoin *GemjoinCaller) Live(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gemjoin.contract.Call(opts, &out, "live")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_Gemjoin *GemjoinSession) Live() (*big.Int, error) {
	return _Gemjoin.Contract.Live(&_Gemjoin.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_Gemjoin *GemjoinCallerSession) Live() (*big.Int, error) {
	return _Gemjoin.Contract.Live(&_Gemjoin.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Gemjoin *GemjoinCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gemjoin.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Gemjoin *GemjoinSession) Vat() (common.Address, error) {
	return _Gemjoin.Contract.Vat(&_Gemjoin.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Gemjoin *GemjoinCallerSession) Vat() (common.Address, error) {
	return _Gemjoin.Contract.Vat(&_Gemjoin.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Gemjoin *GemjoinCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gemjoin.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Gemjoin *GemjoinSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Gemjoin.Contract.Wards(&_Gemjoin.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Gemjoin *GemjoinCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Gemjoin.Contract.Wards(&_Gemjoin.CallOpts, arg0)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Gemjoin *GemjoinTransactor) Cage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gemjoin.contract.Transact(opts, "cage")
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Gemjoin *GemjoinSession) Cage() (*types.Transaction, error) {
	return _Gemjoin.Contract.Cage(&_Gemjoin.TransactOpts)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Gemjoin *GemjoinTransactorSession) Cage() (*types.Transaction, error) {
	return _Gemjoin.Contract.Cage(&_Gemjoin.TransactOpts)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Gemjoin *GemjoinTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Gemjoin.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Gemjoin *GemjoinSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Gemjoin.Contract.Deny(&_Gemjoin.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Gemjoin *GemjoinTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Gemjoin.Contract.Deny(&_Gemjoin.TransactOpts, usr)
}

// Exit is a paid mutator transaction binding the contract method 0xef693bed.
//
// Solidity: function exit(address usr, uint256 wad) returns()
func (_Gemjoin *GemjoinTransactor) Exit(opts *bind.TransactOpts, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Gemjoin.contract.Transact(opts, "exit", usr, wad)
}

// Exit is a paid mutator transaction binding the contract method 0xef693bed.
//
// Solidity: function exit(address usr, uint256 wad) returns()
func (_Gemjoin *GemjoinSession) Exit(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Gemjoin.Contract.Exit(&_Gemjoin.TransactOpts, usr, wad)
}

// Exit is a paid mutator transaction binding the contract method 0xef693bed.
//
// Solidity: function exit(address usr, uint256 wad) returns()
func (_Gemjoin *GemjoinTransactorSession) Exit(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Gemjoin.Contract.Exit(&_Gemjoin.TransactOpts, usr, wad)
}

// Join is a paid mutator transaction binding the contract method 0x3b4da69f.
//
// Solidity: function join(address usr, uint256 wad) returns()
func (_Gemjoin *GemjoinTransactor) Join(opts *bind.TransactOpts, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Gemjoin.contract.Transact(opts, "join", usr, wad)
}

// Join is a paid mutator transaction binding the contract method 0x3b4da69f.
//
// Solidity: function join(address usr, uint256 wad) returns()
func (_Gemjoin *GemjoinSession) Join(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Gemjoin.Contract.Join(&_Gemjoin.TransactOpts, usr, wad)
}

// Join is a paid mutator transaction binding the contract method 0x3b4da69f.
//
// Solidity: function join(address usr, uint256 wad) returns()
func (_Gemjoin *GemjoinTransactorSession) Join(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Gemjoin.Contract.Join(&_Gemjoin.TransactOpts, usr, wad)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Gemjoin *GemjoinTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Gemjoin.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Gemjoin *GemjoinSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Gemjoin.Contract.Rely(&_Gemjoin.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Gemjoin *GemjoinTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Gemjoin.Contract.Rely(&_Gemjoin.TransactOpts, usr)
}

// GemjoinCageIterator is returned from FilterCage and is used to iterate over the raw logs and unpacked data for Cage events raised by the Gemjoin contract.
type GemjoinCageIterator struct {
	Event *GemjoinCage // Event containing the contract specifics and raw log

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
func (it *GemjoinCageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GemjoinCage)
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
		it.Event = new(GemjoinCage)
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
func (it *GemjoinCageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GemjoinCageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GemjoinCage represents a Cage event raised by the Gemjoin contract.
type GemjoinCage struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCage is a free log retrieval operation binding the contract event 0x2308ed18a14e800c39b86eb6ea43270105955ca385b603b64eca89f98ae8fbda.
//
// Solidity: event Cage()
func (_Gemjoin *GemjoinFilterer) FilterCage(opts *bind.FilterOpts) (*GemjoinCageIterator, error) {

	logs, sub, err := _Gemjoin.contract.FilterLogs(opts, "Cage")
	if err != nil {
		return nil, err
	}
	return &GemjoinCageIterator{contract: _Gemjoin.contract, event: "Cage", logs: logs, sub: sub}, nil
}

// WatchCage is a free log subscription operation binding the contract event 0x2308ed18a14e800c39b86eb6ea43270105955ca385b603b64eca89f98ae8fbda.
//
// Solidity: event Cage()
func (_Gemjoin *GemjoinFilterer) WatchCage(opts *bind.WatchOpts, sink chan<- *GemjoinCage) (event.Subscription, error) {

	logs, sub, err := _Gemjoin.contract.WatchLogs(opts, "Cage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GemjoinCage)
				if err := _Gemjoin.contract.UnpackLog(event, "Cage", log); err != nil {
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
func (_Gemjoin *GemjoinFilterer) ParseCage(log types.Log) (*GemjoinCage, error) {
	event := new(GemjoinCage)
	if err := _Gemjoin.contract.UnpackLog(event, "Cage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GemjoinDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the Gemjoin contract.
type GemjoinDenyIterator struct {
	Event *GemjoinDeny // Event containing the contract specifics and raw log

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
func (it *GemjoinDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GemjoinDeny)
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
		it.Event = new(GemjoinDeny)
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
func (it *GemjoinDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GemjoinDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GemjoinDeny represents a Deny event raised by the Gemjoin contract.
type GemjoinDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_Gemjoin *GemjoinFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*GemjoinDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Gemjoin.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &GemjoinDenyIterator{contract: _Gemjoin.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_Gemjoin *GemjoinFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *GemjoinDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Gemjoin.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GemjoinDeny)
				if err := _Gemjoin.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_Gemjoin *GemjoinFilterer) ParseDeny(log types.Log) (*GemjoinDeny, error) {
	event := new(GemjoinDeny)
	if err := _Gemjoin.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GemjoinExitIterator is returned from FilterExit and is used to iterate over the raw logs and unpacked data for Exit events raised by the Gemjoin contract.
type GemjoinExitIterator struct {
	Event *GemjoinExit // Event containing the contract specifics and raw log

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
func (it *GemjoinExitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GemjoinExit)
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
		it.Event = new(GemjoinExit)
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
func (it *GemjoinExitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GemjoinExitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GemjoinExit represents a Exit event raised by the Gemjoin contract.
type GemjoinExit struct {
	Usr common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterExit is a free log retrieval operation binding the contract event 0x22d324652c93739755cf4581508b60875ebdd78c20c0cff5cf8e23452b299631.
//
// Solidity: event Exit(address indexed usr, uint256 wad)
func (_Gemjoin *GemjoinFilterer) FilterExit(opts *bind.FilterOpts, usr []common.Address) (*GemjoinExitIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Gemjoin.contract.FilterLogs(opts, "Exit", usrRule)
	if err != nil {
		return nil, err
	}
	return &GemjoinExitIterator{contract: _Gemjoin.contract, event: "Exit", logs: logs, sub: sub}, nil
}

// WatchExit is a free log subscription operation binding the contract event 0x22d324652c93739755cf4581508b60875ebdd78c20c0cff5cf8e23452b299631.
//
// Solidity: event Exit(address indexed usr, uint256 wad)
func (_Gemjoin *GemjoinFilterer) WatchExit(opts *bind.WatchOpts, sink chan<- *GemjoinExit, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Gemjoin.contract.WatchLogs(opts, "Exit", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GemjoinExit)
				if err := _Gemjoin.contract.UnpackLog(event, "Exit", log); err != nil {
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

// ParseExit is a log parse operation binding the contract event 0x22d324652c93739755cf4581508b60875ebdd78c20c0cff5cf8e23452b299631.
//
// Solidity: event Exit(address indexed usr, uint256 wad)
func (_Gemjoin *GemjoinFilterer) ParseExit(log types.Log) (*GemjoinExit, error) {
	event := new(GemjoinExit)
	if err := _Gemjoin.contract.UnpackLog(event, "Exit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GemjoinJoinIterator is returned from FilterJoin and is used to iterate over the raw logs and unpacked data for Join events raised by the Gemjoin contract.
type GemjoinJoinIterator struct {
	Event *GemjoinJoin // Event containing the contract specifics and raw log

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
func (it *GemjoinJoinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GemjoinJoin)
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
		it.Event = new(GemjoinJoin)
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
func (it *GemjoinJoinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GemjoinJoinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GemjoinJoin represents a Join event raised by the Gemjoin contract.
type GemjoinJoin struct {
	Usr common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterJoin is a free log retrieval operation binding the contract event 0xb4e09949657f21548b58afe74e7b86cd2295da5ff1598ae1e5faecb1cf19ca95.
//
// Solidity: event Join(address indexed usr, uint256 wad)
func (_Gemjoin *GemjoinFilterer) FilterJoin(opts *bind.FilterOpts, usr []common.Address) (*GemjoinJoinIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Gemjoin.contract.FilterLogs(opts, "Join", usrRule)
	if err != nil {
		return nil, err
	}
	return &GemjoinJoinIterator{contract: _Gemjoin.contract, event: "Join", logs: logs, sub: sub}, nil
}

// WatchJoin is a free log subscription operation binding the contract event 0xb4e09949657f21548b58afe74e7b86cd2295da5ff1598ae1e5faecb1cf19ca95.
//
// Solidity: event Join(address indexed usr, uint256 wad)
func (_Gemjoin *GemjoinFilterer) WatchJoin(opts *bind.WatchOpts, sink chan<- *GemjoinJoin, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Gemjoin.contract.WatchLogs(opts, "Join", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GemjoinJoin)
				if err := _Gemjoin.contract.UnpackLog(event, "Join", log); err != nil {
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

// ParseJoin is a log parse operation binding the contract event 0xb4e09949657f21548b58afe74e7b86cd2295da5ff1598ae1e5faecb1cf19ca95.
//
// Solidity: event Join(address indexed usr, uint256 wad)
func (_Gemjoin *GemjoinFilterer) ParseJoin(log types.Log) (*GemjoinJoin, error) {
	event := new(GemjoinJoin)
	if err := _Gemjoin.contract.UnpackLog(event, "Join", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GemjoinRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the Gemjoin contract.
type GemjoinRelyIterator struct {
	Event *GemjoinRely // Event containing the contract specifics and raw log

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
func (it *GemjoinRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GemjoinRely)
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
		it.Event = new(GemjoinRely)
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
func (it *GemjoinRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GemjoinRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GemjoinRely represents a Rely event raised by the Gemjoin contract.
type GemjoinRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_Gemjoin *GemjoinFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*GemjoinRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Gemjoin.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &GemjoinRelyIterator{contract: _Gemjoin.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_Gemjoin *GemjoinFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *GemjoinRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Gemjoin.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GemjoinRely)
				if err := _Gemjoin.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_Gemjoin *GemjoinFilterer) ParseRely(log types.Log) (*GemjoinRely, error) {
	event := new(GemjoinRely)
	if err := _Gemjoin.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
