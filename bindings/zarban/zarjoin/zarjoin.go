// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zarjoin

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

// ZarjoinMetaData contains all meta data concerning the Zarjoin contract.
var ZarjoinMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zar_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Cage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Exit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Join\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"cage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"join\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"live\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zar\",\"outputs\":[{\"internalType\":\"contractTokenLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ZarjoinABI is the input ABI used to generate the binding from.
// Deprecated: Use ZarjoinMetaData.ABI instead.
var ZarjoinABI = ZarjoinMetaData.ABI

// Zarjoin is an auto generated Go binding around an Ethereum contract.
type Zarjoin struct {
	ZarjoinCaller     // Read-only binding to the contract
	ZarjoinTransactor // Write-only binding to the contract
	ZarjoinFilterer   // Log filterer for contract events
}

// ZarjoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZarjoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZarjoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZarjoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZarjoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZarjoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZarjoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZarjoinSession struct {
	Contract     *Zarjoin          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZarjoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZarjoinCallerSession struct {
	Contract *ZarjoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ZarjoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZarjoinTransactorSession struct {
	Contract     *ZarjoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ZarjoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZarjoinRaw struct {
	Contract *Zarjoin // Generic contract binding to access the raw methods on
}

// ZarjoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZarjoinCallerRaw struct {
	Contract *ZarjoinCaller // Generic read-only contract binding to access the raw methods on
}

// ZarjoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZarjoinTransactorRaw struct {
	Contract *ZarjoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZarjoin creates a new instance of Zarjoin, bound to a specific deployed contract.
func NewZarjoin(address common.Address, backend bind.ContractBackend) (*Zarjoin, error) {
	contract, err := bindZarjoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Zarjoin{ZarjoinCaller: ZarjoinCaller{contract: contract}, ZarjoinTransactor: ZarjoinTransactor{contract: contract}, ZarjoinFilterer: ZarjoinFilterer{contract: contract}}, nil
}

// NewZarjoinCaller creates a new read-only instance of Zarjoin, bound to a specific deployed contract.
func NewZarjoinCaller(address common.Address, caller bind.ContractCaller) (*ZarjoinCaller, error) {
	contract, err := bindZarjoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZarjoinCaller{contract: contract}, nil
}

// NewZarjoinTransactor creates a new write-only instance of Zarjoin, bound to a specific deployed contract.
func NewZarjoinTransactor(address common.Address, transactor bind.ContractTransactor) (*ZarjoinTransactor, error) {
	contract, err := bindZarjoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZarjoinTransactor{contract: contract}, nil
}

// NewZarjoinFilterer creates a new log filterer instance of Zarjoin, bound to a specific deployed contract.
func NewZarjoinFilterer(address common.Address, filterer bind.ContractFilterer) (*ZarjoinFilterer, error) {
	contract, err := bindZarjoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZarjoinFilterer{contract: contract}, nil
}

// bindZarjoin binds a generic wrapper to an already deployed contract.
func bindZarjoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZarjoinMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zarjoin *ZarjoinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Zarjoin.Contract.ZarjoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zarjoin *ZarjoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zarjoin.Contract.ZarjoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zarjoin *ZarjoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zarjoin.Contract.ZarjoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zarjoin *ZarjoinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Zarjoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zarjoin *ZarjoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zarjoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zarjoin *ZarjoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zarjoin.Contract.contract.Transact(opts, method, params...)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_Zarjoin *ZarjoinCaller) Live(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zarjoin.contract.Call(opts, &out, "live")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_Zarjoin *ZarjoinSession) Live() (*big.Int, error) {
	return _Zarjoin.Contract.Live(&_Zarjoin.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_Zarjoin *ZarjoinCallerSession) Live() (*big.Int, error) {
	return _Zarjoin.Contract.Live(&_Zarjoin.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Zarjoin *ZarjoinCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Zarjoin.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Zarjoin *ZarjoinSession) Vat() (common.Address, error) {
	return _Zarjoin.Contract.Vat(&_Zarjoin.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Zarjoin *ZarjoinCallerSession) Vat() (common.Address, error) {
	return _Zarjoin.Contract.Vat(&_Zarjoin.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Zarjoin *ZarjoinCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Zarjoin.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Zarjoin *ZarjoinSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Zarjoin.Contract.Wards(&_Zarjoin.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Zarjoin *ZarjoinCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Zarjoin.Contract.Wards(&_Zarjoin.CallOpts, arg0)
}

// Zar is a free data retrieval call binding the contract method 0x7d510ab5.
//
// Solidity: function zar() view returns(address)
func (_Zarjoin *ZarjoinCaller) Zar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Zarjoin.contract.Call(opts, &out, "zar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Zar is a free data retrieval call binding the contract method 0x7d510ab5.
//
// Solidity: function zar() view returns(address)
func (_Zarjoin *ZarjoinSession) Zar() (common.Address, error) {
	return _Zarjoin.Contract.Zar(&_Zarjoin.CallOpts)
}

// Zar is a free data retrieval call binding the contract method 0x7d510ab5.
//
// Solidity: function zar() view returns(address)
func (_Zarjoin *ZarjoinCallerSession) Zar() (common.Address, error) {
	return _Zarjoin.Contract.Zar(&_Zarjoin.CallOpts)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Zarjoin *ZarjoinTransactor) Cage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zarjoin.contract.Transact(opts, "cage")
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Zarjoin *ZarjoinSession) Cage() (*types.Transaction, error) {
	return _Zarjoin.Contract.Cage(&_Zarjoin.TransactOpts)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Zarjoin *ZarjoinTransactorSession) Cage() (*types.Transaction, error) {
	return _Zarjoin.Contract.Cage(&_Zarjoin.TransactOpts)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Zarjoin *ZarjoinTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Zarjoin.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Zarjoin *ZarjoinSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Zarjoin.Contract.Deny(&_Zarjoin.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Zarjoin *ZarjoinTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Zarjoin.Contract.Deny(&_Zarjoin.TransactOpts, usr)
}

// Exit is a paid mutator transaction binding the contract method 0xef693bed.
//
// Solidity: function exit(address usr, uint256 wad) returns()
func (_Zarjoin *ZarjoinTransactor) Exit(opts *bind.TransactOpts, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Zarjoin.contract.Transact(opts, "exit", usr, wad)
}

// Exit is a paid mutator transaction binding the contract method 0xef693bed.
//
// Solidity: function exit(address usr, uint256 wad) returns()
func (_Zarjoin *ZarjoinSession) Exit(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Zarjoin.Contract.Exit(&_Zarjoin.TransactOpts, usr, wad)
}

// Exit is a paid mutator transaction binding the contract method 0xef693bed.
//
// Solidity: function exit(address usr, uint256 wad) returns()
func (_Zarjoin *ZarjoinTransactorSession) Exit(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Zarjoin.Contract.Exit(&_Zarjoin.TransactOpts, usr, wad)
}

// Join is a paid mutator transaction binding the contract method 0x3b4da69f.
//
// Solidity: function join(address usr, uint256 wad) returns()
func (_Zarjoin *ZarjoinTransactor) Join(opts *bind.TransactOpts, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Zarjoin.contract.Transact(opts, "join", usr, wad)
}

// Join is a paid mutator transaction binding the contract method 0x3b4da69f.
//
// Solidity: function join(address usr, uint256 wad) returns()
func (_Zarjoin *ZarjoinSession) Join(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Zarjoin.Contract.Join(&_Zarjoin.TransactOpts, usr, wad)
}

// Join is a paid mutator transaction binding the contract method 0x3b4da69f.
//
// Solidity: function join(address usr, uint256 wad) returns()
func (_Zarjoin *ZarjoinTransactorSession) Join(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Zarjoin.Contract.Join(&_Zarjoin.TransactOpts, usr, wad)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Zarjoin *ZarjoinTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Zarjoin.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Zarjoin *ZarjoinSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Zarjoin.Contract.Rely(&_Zarjoin.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Zarjoin *ZarjoinTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Zarjoin.Contract.Rely(&_Zarjoin.TransactOpts, usr)
}

// ZarjoinCageIterator is returned from FilterCage and is used to iterate over the raw logs and unpacked data for Cage events raised by the Zarjoin contract.
type ZarjoinCageIterator struct {
	Event *ZarjoinCage // Event containing the contract specifics and raw log

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
func (it *ZarjoinCageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZarjoinCage)
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
		it.Event = new(ZarjoinCage)
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
func (it *ZarjoinCageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZarjoinCageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZarjoinCage represents a Cage event raised by the Zarjoin contract.
type ZarjoinCage struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCage is a free log retrieval operation binding the contract event 0x2308ed18a14e800c39b86eb6ea43270105955ca385b603b64eca89f98ae8fbda.
//
// Solidity: event Cage()
func (_Zarjoin *ZarjoinFilterer) FilterCage(opts *bind.FilterOpts) (*ZarjoinCageIterator, error) {

	logs, sub, err := _Zarjoin.contract.FilterLogs(opts, "Cage")
	if err != nil {
		return nil, err
	}
	return &ZarjoinCageIterator{contract: _Zarjoin.contract, event: "Cage", logs: logs, sub: sub}, nil
}

// WatchCage is a free log subscription operation binding the contract event 0x2308ed18a14e800c39b86eb6ea43270105955ca385b603b64eca89f98ae8fbda.
//
// Solidity: event Cage()
func (_Zarjoin *ZarjoinFilterer) WatchCage(opts *bind.WatchOpts, sink chan<- *ZarjoinCage) (event.Subscription, error) {

	logs, sub, err := _Zarjoin.contract.WatchLogs(opts, "Cage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZarjoinCage)
				if err := _Zarjoin.contract.UnpackLog(event, "Cage", log); err != nil {
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
func (_Zarjoin *ZarjoinFilterer) ParseCage(log types.Log) (*ZarjoinCage, error) {
	event := new(ZarjoinCage)
	if err := _Zarjoin.contract.UnpackLog(event, "Cage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZarjoinDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the Zarjoin contract.
type ZarjoinDenyIterator struct {
	Event *ZarjoinDeny // Event containing the contract specifics and raw log

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
func (it *ZarjoinDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZarjoinDeny)
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
		it.Event = new(ZarjoinDeny)
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
func (it *ZarjoinDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZarjoinDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZarjoinDeny represents a Deny event raised by the Zarjoin contract.
type ZarjoinDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_Zarjoin *ZarjoinFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*ZarjoinDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Zarjoin.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &ZarjoinDenyIterator{contract: _Zarjoin.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_Zarjoin *ZarjoinFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *ZarjoinDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Zarjoin.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZarjoinDeny)
				if err := _Zarjoin.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_Zarjoin *ZarjoinFilterer) ParseDeny(log types.Log) (*ZarjoinDeny, error) {
	event := new(ZarjoinDeny)
	if err := _Zarjoin.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZarjoinExitIterator is returned from FilterExit and is used to iterate over the raw logs and unpacked data for Exit events raised by the Zarjoin contract.
type ZarjoinExitIterator struct {
	Event *ZarjoinExit // Event containing the contract specifics and raw log

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
func (it *ZarjoinExitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZarjoinExit)
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
		it.Event = new(ZarjoinExit)
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
func (it *ZarjoinExitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZarjoinExitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZarjoinExit represents a Exit event raised by the Zarjoin contract.
type ZarjoinExit struct {
	Usr common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterExit is a free log retrieval operation binding the contract event 0x22d324652c93739755cf4581508b60875ebdd78c20c0cff5cf8e23452b299631.
//
// Solidity: event Exit(address indexed usr, uint256 wad)
func (_Zarjoin *ZarjoinFilterer) FilterExit(opts *bind.FilterOpts, usr []common.Address) (*ZarjoinExitIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Zarjoin.contract.FilterLogs(opts, "Exit", usrRule)
	if err != nil {
		return nil, err
	}
	return &ZarjoinExitIterator{contract: _Zarjoin.contract, event: "Exit", logs: logs, sub: sub}, nil
}

// WatchExit is a free log subscription operation binding the contract event 0x22d324652c93739755cf4581508b60875ebdd78c20c0cff5cf8e23452b299631.
//
// Solidity: event Exit(address indexed usr, uint256 wad)
func (_Zarjoin *ZarjoinFilterer) WatchExit(opts *bind.WatchOpts, sink chan<- *ZarjoinExit, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Zarjoin.contract.WatchLogs(opts, "Exit", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZarjoinExit)
				if err := _Zarjoin.contract.UnpackLog(event, "Exit", log); err != nil {
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
func (_Zarjoin *ZarjoinFilterer) ParseExit(log types.Log) (*ZarjoinExit, error) {
	event := new(ZarjoinExit)
	if err := _Zarjoin.contract.UnpackLog(event, "Exit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZarjoinJoinIterator is returned from FilterJoin and is used to iterate over the raw logs and unpacked data for Join events raised by the Zarjoin contract.
type ZarjoinJoinIterator struct {
	Event *ZarjoinJoin // Event containing the contract specifics and raw log

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
func (it *ZarjoinJoinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZarjoinJoin)
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
		it.Event = new(ZarjoinJoin)
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
func (it *ZarjoinJoinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZarjoinJoinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZarjoinJoin represents a Join event raised by the Zarjoin contract.
type ZarjoinJoin struct {
	Usr common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterJoin is a free log retrieval operation binding the contract event 0xb4e09949657f21548b58afe74e7b86cd2295da5ff1598ae1e5faecb1cf19ca95.
//
// Solidity: event Join(address indexed usr, uint256 wad)
func (_Zarjoin *ZarjoinFilterer) FilterJoin(opts *bind.FilterOpts, usr []common.Address) (*ZarjoinJoinIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Zarjoin.contract.FilterLogs(opts, "Join", usrRule)
	if err != nil {
		return nil, err
	}
	return &ZarjoinJoinIterator{contract: _Zarjoin.contract, event: "Join", logs: logs, sub: sub}, nil
}

// WatchJoin is a free log subscription operation binding the contract event 0xb4e09949657f21548b58afe74e7b86cd2295da5ff1598ae1e5faecb1cf19ca95.
//
// Solidity: event Join(address indexed usr, uint256 wad)
func (_Zarjoin *ZarjoinFilterer) WatchJoin(opts *bind.WatchOpts, sink chan<- *ZarjoinJoin, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Zarjoin.contract.WatchLogs(opts, "Join", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZarjoinJoin)
				if err := _Zarjoin.contract.UnpackLog(event, "Join", log); err != nil {
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
func (_Zarjoin *ZarjoinFilterer) ParseJoin(log types.Log) (*ZarjoinJoin, error) {
	event := new(ZarjoinJoin)
	if err := _Zarjoin.contract.UnpackLog(event, "Join", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZarjoinRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the Zarjoin contract.
type ZarjoinRelyIterator struct {
	Event *ZarjoinRely // Event containing the contract specifics and raw log

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
func (it *ZarjoinRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZarjoinRely)
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
		it.Event = new(ZarjoinRely)
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
func (it *ZarjoinRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZarjoinRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZarjoinRely represents a Rely event raised by the Zarjoin contract.
type ZarjoinRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_Zarjoin *ZarjoinFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*ZarjoinRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Zarjoin.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &ZarjoinRelyIterator{contract: _Zarjoin.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_Zarjoin *ZarjoinFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *ZarjoinRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Zarjoin.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZarjoinRely)
				if err := _Zarjoin.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_Zarjoin *ZarjoinFilterer) ParseRely(log types.Log) (*ZarjoinRely, error) {
	event := new(ZarjoinRely)
	if err := _Zarjoin.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
