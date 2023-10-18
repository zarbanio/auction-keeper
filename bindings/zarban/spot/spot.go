// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package spot

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

// SpotMetaData contains all meta data concerning the Spot contract.
var SpotMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pip_\",\"type\":\"address\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"val\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"spot\",\"type\":\"uint256\"}],\"name\":\"Poke\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"cage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"pip_\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ilks\",\"outputs\":[{\"internalType\":\"contractPipLike\",\"name\":\"pip\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mat\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"live\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"par\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"poke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SpotABI is the input ABI used to generate the binding from.
// Deprecated: Use SpotMetaData.ABI instead.
var SpotABI = SpotMetaData.ABI

// Spot is an auto generated Go binding around an Ethereum contract.
type Spot struct {
	SpotCaller     // Read-only binding to the contract
	SpotTransactor // Write-only binding to the contract
	SpotFilterer   // Log filterer for contract events
}

// SpotCaller is an auto generated read-only Go binding around an Ethereum contract.
type SpotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SpotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SpotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SpotSession struct {
	Contract     *Spot             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SpotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SpotCallerSession struct {
	Contract *SpotCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SpotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SpotTransactorSession struct {
	Contract     *SpotTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SpotRaw is an auto generated low-level Go binding around an Ethereum contract.
type SpotRaw struct {
	Contract *Spot // Generic contract binding to access the raw methods on
}

// SpotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SpotCallerRaw struct {
	Contract *SpotCaller // Generic read-only contract binding to access the raw methods on
}

// SpotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SpotTransactorRaw struct {
	Contract *SpotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSpot creates a new instance of Spot, bound to a specific deployed contract.
func NewSpot(address common.Address, backend bind.ContractBackend) (*Spot, error) {
	contract, err := bindSpot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Spot{SpotCaller: SpotCaller{contract: contract}, SpotTransactor: SpotTransactor{contract: contract}, SpotFilterer: SpotFilterer{contract: contract}}, nil
}

// NewSpotCaller creates a new read-only instance of Spot, bound to a specific deployed contract.
func NewSpotCaller(address common.Address, caller bind.ContractCaller) (*SpotCaller, error) {
	contract, err := bindSpot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SpotCaller{contract: contract}, nil
}

// NewSpotTransactor creates a new write-only instance of Spot, bound to a specific deployed contract.
func NewSpotTransactor(address common.Address, transactor bind.ContractTransactor) (*SpotTransactor, error) {
	contract, err := bindSpot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SpotTransactor{contract: contract}, nil
}

// NewSpotFilterer creates a new log filterer instance of Spot, bound to a specific deployed contract.
func NewSpotFilterer(address common.Address, filterer bind.ContractFilterer) (*SpotFilterer, error) {
	contract, err := bindSpot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SpotFilterer{contract: contract}, nil
}

// bindSpot binds a generic wrapper to an already deployed contract.
func bindSpot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SpotMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Spot *SpotRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Spot.Contract.SpotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Spot *SpotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spot.Contract.SpotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Spot *SpotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Spot.Contract.SpotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Spot *SpotCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Spot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Spot *SpotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Spot *SpotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Spot.Contract.contract.Transact(opts, method, params...)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(address pip, uint256 mat)
func (_Spot *SpotCaller) Ilks(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Pip common.Address
	Mat *big.Int
}, error) {
	var out []interface{}
	err := _Spot.contract.Call(opts, &out, "ilks", arg0)

	outstruct := new(struct {
		Pip common.Address
		Mat *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pip = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Mat = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(address pip, uint256 mat)
func (_Spot *SpotSession) Ilks(arg0 [32]byte) (struct {
	Pip common.Address
	Mat *big.Int
}, error) {
	return _Spot.Contract.Ilks(&_Spot.CallOpts, arg0)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(address pip, uint256 mat)
func (_Spot *SpotCallerSession) Ilks(arg0 [32]byte) (struct {
	Pip common.Address
	Mat *big.Int
}, error) {
	return _Spot.Contract.Ilks(&_Spot.CallOpts, arg0)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_Spot *SpotCaller) Live(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Spot.contract.Call(opts, &out, "live")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_Spot *SpotSession) Live() (*big.Int, error) {
	return _Spot.Contract.Live(&_Spot.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_Spot *SpotCallerSession) Live() (*big.Int, error) {
	return _Spot.Contract.Live(&_Spot.CallOpts)
}

// Par is a free data retrieval call binding the contract method 0x495d32cb.
//
// Solidity: function par() view returns(uint256)
func (_Spot *SpotCaller) Par(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Spot.contract.Call(opts, &out, "par")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Par is a free data retrieval call binding the contract method 0x495d32cb.
//
// Solidity: function par() view returns(uint256)
func (_Spot *SpotSession) Par() (*big.Int, error) {
	return _Spot.Contract.Par(&_Spot.CallOpts)
}

// Par is a free data retrieval call binding the contract method 0x495d32cb.
//
// Solidity: function par() view returns(uint256)
func (_Spot *SpotCallerSession) Par() (*big.Int, error) {
	return _Spot.Contract.Par(&_Spot.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Spot *SpotCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Spot.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Spot *SpotSession) Vat() (common.Address, error) {
	return _Spot.Contract.Vat(&_Spot.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Spot *SpotCallerSession) Vat() (common.Address, error) {
	return _Spot.Contract.Vat(&_Spot.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Spot *SpotCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Spot.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Spot *SpotSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Spot.Contract.Wards(&_Spot.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Spot *SpotCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Spot.Contract.Wards(&_Spot.CallOpts, arg0)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Spot *SpotTransactor) Cage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spot.contract.Transact(opts, "cage")
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Spot *SpotSession) Cage() (*types.Transaction, error) {
	return _Spot.Contract.Cage(&_Spot.TransactOpts)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Spot *SpotTransactorSession) Cage() (*types.Transaction, error) {
	return _Spot.Contract.Cage(&_Spot.TransactOpts)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address guy) returns()
func (_Spot *SpotTransactor) Deny(opts *bind.TransactOpts, guy common.Address) (*types.Transaction, error) {
	return _Spot.contract.Transact(opts, "deny", guy)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address guy) returns()
func (_Spot *SpotSession) Deny(guy common.Address) (*types.Transaction, error) {
	return _Spot.Contract.Deny(&_Spot.TransactOpts, guy)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address guy) returns()
func (_Spot *SpotTransactorSession) Deny(guy common.Address) (*types.Transaction, error) {
	return _Spot.Contract.Deny(&_Spot.TransactOpts, guy)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_Spot *SpotTransactor) File(opts *bind.TransactOpts, ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Spot.contract.Transact(opts, "file", ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_Spot *SpotSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Spot.Contract.File(&_Spot.TransactOpts, ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_Spot *SpotTransactorSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Spot.Contract.File(&_Spot.TransactOpts, ilk, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Spot *SpotTransactor) File0(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Spot.contract.Transact(opts, "file0", what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Spot *SpotSession) File0(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Spot.Contract.File0(&_Spot.TransactOpts, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Spot *SpotTransactorSession) File0(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Spot.Contract.File0(&_Spot.TransactOpts, what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xebecb39d.
//
// Solidity: function file(bytes32 ilk, bytes32 what, address pip_) returns()
func (_Spot *SpotTransactor) File1(opts *bind.TransactOpts, ilk [32]byte, what [32]byte, pip_ common.Address) (*types.Transaction, error) {
	return _Spot.contract.Transact(opts, "file1", ilk, what, pip_)
}

// File1 is a paid mutator transaction binding the contract method 0xebecb39d.
//
// Solidity: function file(bytes32 ilk, bytes32 what, address pip_) returns()
func (_Spot *SpotSession) File1(ilk [32]byte, what [32]byte, pip_ common.Address) (*types.Transaction, error) {
	return _Spot.Contract.File1(&_Spot.TransactOpts, ilk, what, pip_)
}

// File1 is a paid mutator transaction binding the contract method 0xebecb39d.
//
// Solidity: function file(bytes32 ilk, bytes32 what, address pip_) returns()
func (_Spot *SpotTransactorSession) File1(ilk [32]byte, what [32]byte, pip_ common.Address) (*types.Transaction, error) {
	return _Spot.Contract.File1(&_Spot.TransactOpts, ilk, what, pip_)
}

// Poke is a paid mutator transaction binding the contract method 0x1504460f.
//
// Solidity: function poke(bytes32 ilk) returns()
func (_Spot *SpotTransactor) Poke(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _Spot.contract.Transact(opts, "poke", ilk)
}

// Poke is a paid mutator transaction binding the contract method 0x1504460f.
//
// Solidity: function poke(bytes32 ilk) returns()
func (_Spot *SpotSession) Poke(ilk [32]byte) (*types.Transaction, error) {
	return _Spot.Contract.Poke(&_Spot.TransactOpts, ilk)
}

// Poke is a paid mutator transaction binding the contract method 0x1504460f.
//
// Solidity: function poke(bytes32 ilk) returns()
func (_Spot *SpotTransactorSession) Poke(ilk [32]byte) (*types.Transaction, error) {
	return _Spot.Contract.Poke(&_Spot.TransactOpts, ilk)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address guy) returns()
func (_Spot *SpotTransactor) Rely(opts *bind.TransactOpts, guy common.Address) (*types.Transaction, error) {
	return _Spot.contract.Transact(opts, "rely", guy)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address guy) returns()
func (_Spot *SpotSession) Rely(guy common.Address) (*types.Transaction, error) {
	return _Spot.Contract.Rely(&_Spot.TransactOpts, guy)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address guy) returns()
func (_Spot *SpotTransactorSession) Rely(guy common.Address) (*types.Transaction, error) {
	return _Spot.Contract.Rely(&_Spot.TransactOpts, guy)
}

// SpotDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the Spot contract.
type SpotDenyIterator struct {
	Event *SpotDeny // Event containing the contract specifics and raw log

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
func (it *SpotDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotDeny)
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
		it.Event = new(SpotDeny)
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
func (it *SpotDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotDeny represents a Deny event raised by the Spot contract.
type SpotDeny struct {
	Guy common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed guy)
func (_Spot *SpotFilterer) FilterDeny(opts *bind.FilterOpts, guy []common.Address) (*SpotDenyIterator, error) {

	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Spot.contract.FilterLogs(opts, "Deny", guyRule)
	if err != nil {
		return nil, err
	}
	return &SpotDenyIterator{contract: _Spot.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed guy)
func (_Spot *SpotFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *SpotDeny, guy []common.Address) (event.Subscription, error) {

	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Spot.contract.WatchLogs(opts, "Deny", guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotDeny)
				if err := _Spot.contract.UnpackLog(event, "Deny", log); err != nil {
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
// Solidity: event Deny(address indexed guy)
func (_Spot *SpotFilterer) ParseDeny(log types.Log) (*SpotDeny, error) {
	event := new(SpotDeny)
	if err := _Spot.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotFileIterator is returned from FilterFile and is used to iterate over the raw logs and unpacked data for File events raised by the Spot contract.
type SpotFileIterator struct {
	Event *SpotFile // Event containing the contract specifics and raw log

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
func (it *SpotFileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotFile)
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
		it.Event = new(SpotFile)
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
func (it *SpotFileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotFileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotFile represents a File event raised by the Spot contract.
type SpotFile struct {
	Ilk  [32]byte
	What [32]byte
	Data *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile is a free log retrieval operation binding the contract event 0x851aa1caf4888170ad8875449d18f0f512fd6deb2a6571ea1a41fb9f95acbcd1.
//
// Solidity: event File(bytes32 indexed ilk, bytes32 indexed what, uint256 data)
func (_Spot *SpotFilterer) FilterFile(opts *bind.FilterOpts, ilk [][32]byte, what [][32]byte) (*SpotFileIterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Spot.contract.FilterLogs(opts, "File", ilkRule, whatRule)
	if err != nil {
		return nil, err
	}
	return &SpotFileIterator{contract: _Spot.contract, event: "File", logs: logs, sub: sub}, nil
}

// WatchFile is a free log subscription operation binding the contract event 0x851aa1caf4888170ad8875449d18f0f512fd6deb2a6571ea1a41fb9f95acbcd1.
//
// Solidity: event File(bytes32 indexed ilk, bytes32 indexed what, uint256 data)
func (_Spot *SpotFilterer) WatchFile(opts *bind.WatchOpts, sink chan<- *SpotFile, ilk [][32]byte, what [][32]byte) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Spot.contract.WatchLogs(opts, "File", ilkRule, whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotFile)
				if err := _Spot.contract.UnpackLog(event, "File", log); err != nil {
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

// ParseFile is a log parse operation binding the contract event 0x851aa1caf4888170ad8875449d18f0f512fd6deb2a6571ea1a41fb9f95acbcd1.
//
// Solidity: event File(bytes32 indexed ilk, bytes32 indexed what, uint256 data)
func (_Spot *SpotFilterer) ParseFile(log types.Log) (*SpotFile, error) {
	event := new(SpotFile)
	if err := _Spot.contract.UnpackLog(event, "File", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotFile0Iterator is returned from FilterFile0 and is used to iterate over the raw logs and unpacked data for File0 events raised by the Spot contract.
type SpotFile0Iterator struct {
	Event *SpotFile0 // Event containing the contract specifics and raw log

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
func (it *SpotFile0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotFile0)
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
		it.Event = new(SpotFile0)
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
func (it *SpotFile0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotFile0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotFile0 represents a File0 event raised by the Spot contract.
type SpotFile0 struct {
	Ilk  [32]byte
	What [32]byte
	Pip  common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile0 is a free log retrieval operation binding the contract event 0x4ff2caaa972a7c6629ea01fae9c93d73cc307d13ea4c369f9bbbb7f9b7e9461d.
//
// Solidity: event File(bytes32 indexed ilk, bytes32 indexed what, address pip_)
func (_Spot *SpotFilterer) FilterFile0(opts *bind.FilterOpts, ilk [][32]byte, what [][32]byte) (*SpotFile0Iterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Spot.contract.FilterLogs(opts, "File0", ilkRule, whatRule)
	if err != nil {
		return nil, err
	}
	return &SpotFile0Iterator{contract: _Spot.contract, event: "File0", logs: logs, sub: sub}, nil
}

// WatchFile0 is a free log subscription operation binding the contract event 0x4ff2caaa972a7c6629ea01fae9c93d73cc307d13ea4c369f9bbbb7f9b7e9461d.
//
// Solidity: event File(bytes32 indexed ilk, bytes32 indexed what, address pip_)
func (_Spot *SpotFilterer) WatchFile0(opts *bind.WatchOpts, sink chan<- *SpotFile0, ilk [][32]byte, what [][32]byte) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Spot.contract.WatchLogs(opts, "File0", ilkRule, whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotFile0)
				if err := _Spot.contract.UnpackLog(event, "File0", log); err != nil {
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

// ParseFile0 is a log parse operation binding the contract event 0x4ff2caaa972a7c6629ea01fae9c93d73cc307d13ea4c369f9bbbb7f9b7e9461d.
//
// Solidity: event File(bytes32 indexed ilk, bytes32 indexed what, address pip_)
func (_Spot *SpotFilterer) ParseFile0(log types.Log) (*SpotFile0, error) {
	event := new(SpotFile0)
	if err := _Spot.contract.UnpackLog(event, "File0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotFile1Iterator is returned from FilterFile1 and is used to iterate over the raw logs and unpacked data for File1 events raised by the Spot contract.
type SpotFile1Iterator struct {
	Event *SpotFile1 // Event containing the contract specifics and raw log

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
func (it *SpotFile1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotFile1)
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
		it.Event = new(SpotFile1)
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
func (it *SpotFile1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotFile1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotFile1 represents a File1 event raised by the Spot contract.
type SpotFile1 struct {
	What [32]byte
	Data *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile1 is a free log retrieval operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 what, uint256 data)
func (_Spot *SpotFilterer) FilterFile1(opts *bind.FilterOpts) (*SpotFile1Iterator, error) {

	logs, sub, err := _Spot.contract.FilterLogs(opts, "File1")
	if err != nil {
		return nil, err
	}
	return &SpotFile1Iterator{contract: _Spot.contract, event: "File1", logs: logs, sub: sub}, nil
}

// WatchFile1 is a free log subscription operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 what, uint256 data)
func (_Spot *SpotFilterer) WatchFile1(opts *bind.WatchOpts, sink chan<- *SpotFile1) (event.Subscription, error) {

	logs, sub, err := _Spot.contract.WatchLogs(opts, "File1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotFile1)
				if err := _Spot.contract.UnpackLog(event, "File1", log); err != nil {
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

// ParseFile1 is a log parse operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 what, uint256 data)
func (_Spot *SpotFilterer) ParseFile1(log types.Log) (*SpotFile1, error) {
	event := new(SpotFile1)
	if err := _Spot.contract.UnpackLog(event, "File1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotPokeIterator is returned from FilterPoke and is used to iterate over the raw logs and unpacked data for Poke events raised by the Spot contract.
type SpotPokeIterator struct {
	Event *SpotPoke // Event containing the contract specifics and raw log

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
func (it *SpotPokeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotPoke)
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
		it.Event = new(SpotPoke)
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
func (it *SpotPokeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotPokeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotPoke represents a Poke event raised by the Spot contract.
type SpotPoke struct {
	Ilk  [32]byte
	Val  [32]byte
	Spot *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPoke is a free log retrieval operation binding the contract event 0xdfd7467e425a8107cfd368d159957692c25085aacbcf5228ce08f10f2146486e.
//
// Solidity: event Poke(bytes32 ilk, bytes32 val, uint256 spot)
func (_Spot *SpotFilterer) FilterPoke(opts *bind.FilterOpts) (*SpotPokeIterator, error) {

	logs, sub, err := _Spot.contract.FilterLogs(opts, "Poke")
	if err != nil {
		return nil, err
	}
	return &SpotPokeIterator{contract: _Spot.contract, event: "Poke", logs: logs, sub: sub}, nil
}

// WatchPoke is a free log subscription operation binding the contract event 0xdfd7467e425a8107cfd368d159957692c25085aacbcf5228ce08f10f2146486e.
//
// Solidity: event Poke(bytes32 ilk, bytes32 val, uint256 spot)
func (_Spot *SpotFilterer) WatchPoke(opts *bind.WatchOpts, sink chan<- *SpotPoke) (event.Subscription, error) {

	logs, sub, err := _Spot.contract.WatchLogs(opts, "Poke")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotPoke)
				if err := _Spot.contract.UnpackLog(event, "Poke", log); err != nil {
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

// ParsePoke is a log parse operation binding the contract event 0xdfd7467e425a8107cfd368d159957692c25085aacbcf5228ce08f10f2146486e.
//
// Solidity: event Poke(bytes32 ilk, bytes32 val, uint256 spot)
func (_Spot *SpotFilterer) ParsePoke(log types.Log) (*SpotPoke, error) {
	event := new(SpotPoke)
	if err := _Spot.contract.UnpackLog(event, "Poke", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the Spot contract.
type SpotRelyIterator struct {
	Event *SpotRely // Event containing the contract specifics and raw log

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
func (it *SpotRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotRely)
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
		it.Event = new(SpotRely)
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
func (it *SpotRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotRely represents a Rely event raised by the Spot contract.
type SpotRely struct {
	Guy common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed guy)
func (_Spot *SpotFilterer) FilterRely(opts *bind.FilterOpts, guy []common.Address) (*SpotRelyIterator, error) {

	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Spot.contract.FilterLogs(opts, "Rely", guyRule)
	if err != nil {
		return nil, err
	}
	return &SpotRelyIterator{contract: _Spot.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed guy)
func (_Spot *SpotFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *SpotRely, guy []common.Address) (event.Subscription, error) {

	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Spot.contract.WatchLogs(opts, "Rely", guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotRely)
				if err := _Spot.contract.UnpackLog(event, "Rely", log); err != nil {
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
// Solidity: event Rely(address indexed guy)
func (_Spot *SpotFilterer) ParseRely(log types.Log) (*SpotRely, error) {
	event := new(SpotRely)
	if err := _Spot.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
