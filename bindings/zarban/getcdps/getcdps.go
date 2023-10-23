// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package getcdps

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

// GetcdpsMetaData contains all meta data concerning the Getcdps contract.
var GetcdpsMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"getCdpsAsc\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"urns\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ilks\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"getCdpsDesc\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"urns\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ilks\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// GetcdpsABI is the input ABI used to generate the binding from.
// Deprecated: Use GetcdpsMetaData.ABI instead.
var GetcdpsABI = GetcdpsMetaData.ABI

// Getcdps is an auto generated Go binding around an Ethereum contract.
type Getcdps struct {
	GetcdpsCaller     // Read-only binding to the contract
	GetcdpsTransactor // Write-only binding to the contract
	GetcdpsFilterer   // Log filterer for contract events
}

// GetcdpsCaller is an auto generated read-only Go binding around an Ethereum contract.
type GetcdpsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GetcdpsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GetcdpsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GetcdpsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GetcdpsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GetcdpsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GetcdpsSession struct {
	Contract     *Getcdps          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GetcdpsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GetcdpsCallerSession struct {
	Contract *GetcdpsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// GetcdpsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GetcdpsTransactorSession struct {
	Contract     *GetcdpsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GetcdpsRaw is an auto generated low-level Go binding around an Ethereum contract.
type GetcdpsRaw struct {
	Contract *Getcdps // Generic contract binding to access the raw methods on
}

// GetcdpsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GetcdpsCallerRaw struct {
	Contract *GetcdpsCaller // Generic read-only contract binding to access the raw methods on
}

// GetcdpsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GetcdpsTransactorRaw struct {
	Contract *GetcdpsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGetcdps creates a new instance of Getcdps, bound to a specific deployed contract.
func NewGetcdps(address common.Address, backend bind.ContractBackend) (*Getcdps, error) {
	contract, err := bindGetcdps(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Getcdps{GetcdpsCaller: GetcdpsCaller{contract: contract}, GetcdpsTransactor: GetcdpsTransactor{contract: contract}, GetcdpsFilterer: GetcdpsFilterer{contract: contract}}, nil
}

// NewGetcdpsCaller creates a new read-only instance of Getcdps, bound to a specific deployed contract.
func NewGetcdpsCaller(address common.Address, caller bind.ContractCaller) (*GetcdpsCaller, error) {
	contract, err := bindGetcdps(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GetcdpsCaller{contract: contract}, nil
}

// NewGetcdpsTransactor creates a new write-only instance of Getcdps, bound to a specific deployed contract.
func NewGetcdpsTransactor(address common.Address, transactor bind.ContractTransactor) (*GetcdpsTransactor, error) {
	contract, err := bindGetcdps(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GetcdpsTransactor{contract: contract}, nil
}

// NewGetcdpsFilterer creates a new log filterer instance of Getcdps, bound to a specific deployed contract.
func NewGetcdpsFilterer(address common.Address, filterer bind.ContractFilterer) (*GetcdpsFilterer, error) {
	contract, err := bindGetcdps(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GetcdpsFilterer{contract: contract}, nil
}

// bindGetcdps binds a generic wrapper to an already deployed contract.
func bindGetcdps(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GetcdpsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Getcdps *GetcdpsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Getcdps.Contract.GetcdpsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Getcdps *GetcdpsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Getcdps.Contract.GetcdpsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Getcdps *GetcdpsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Getcdps.Contract.GetcdpsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Getcdps *GetcdpsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Getcdps.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Getcdps *GetcdpsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Getcdps.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Getcdps *GetcdpsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Getcdps.Contract.contract.Transact(opts, method, params...)
}

// GetCdpsAsc is a free data retrieval call binding the contract method 0x1ce03f38.
//
// Solidity: function getCdpsAsc(address manager, address guy) view returns(uint256[] ids, address[] urns, bytes32[] ilks)
func (_Getcdps *GetcdpsCaller) GetCdpsAsc(opts *bind.CallOpts, manager common.Address, guy common.Address) (struct {
	Ids  []*big.Int
	Urns []common.Address
	Ilks [][32]byte
}, error) {
	var out []interface{}
	err := _Getcdps.contract.Call(opts, &out, "getCdpsAsc", manager, guy)

	outstruct := new(struct {
		Ids  []*big.Int
		Urns []common.Address
		Ilks [][32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ids = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Urns = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Ilks = *abi.ConvertType(out[2], new([][32]byte)).(*[][32]byte)

	return *outstruct, err

}

// GetCdpsAsc is a free data retrieval call binding the contract method 0x1ce03f38.
//
// Solidity: function getCdpsAsc(address manager, address guy) view returns(uint256[] ids, address[] urns, bytes32[] ilks)
func (_Getcdps *GetcdpsSession) GetCdpsAsc(manager common.Address, guy common.Address) (struct {
	Ids  []*big.Int
	Urns []common.Address
	Ilks [][32]byte
}, error) {
	return _Getcdps.Contract.GetCdpsAsc(&_Getcdps.CallOpts, manager, guy)
}

// GetCdpsAsc is a free data retrieval call binding the contract method 0x1ce03f38.
//
// Solidity: function getCdpsAsc(address manager, address guy) view returns(uint256[] ids, address[] urns, bytes32[] ilks)
func (_Getcdps *GetcdpsCallerSession) GetCdpsAsc(manager common.Address, guy common.Address) (struct {
	Ids  []*big.Int
	Urns []common.Address
	Ilks [][32]byte
}, error) {
	return _Getcdps.Contract.GetCdpsAsc(&_Getcdps.CallOpts, manager, guy)
}

// GetCdpsDesc is a free data retrieval call binding the contract method 0x38f7acb4.
//
// Solidity: function getCdpsDesc(address manager, address guy) view returns(uint256[] ids, address[] urns, bytes32[] ilks)
func (_Getcdps *GetcdpsCaller) GetCdpsDesc(opts *bind.CallOpts, manager common.Address, guy common.Address) (struct {
	Ids  []*big.Int
	Urns []common.Address
	Ilks [][32]byte
}, error) {
	var out []interface{}
	err := _Getcdps.contract.Call(opts, &out, "getCdpsDesc", manager, guy)

	outstruct := new(struct {
		Ids  []*big.Int
		Urns []common.Address
		Ilks [][32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ids = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Urns = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Ilks = *abi.ConvertType(out[2], new([][32]byte)).(*[][32]byte)

	return *outstruct, err

}

// GetCdpsDesc is a free data retrieval call binding the contract method 0x38f7acb4.
//
// Solidity: function getCdpsDesc(address manager, address guy) view returns(uint256[] ids, address[] urns, bytes32[] ilks)
func (_Getcdps *GetcdpsSession) GetCdpsDesc(manager common.Address, guy common.Address) (struct {
	Ids  []*big.Int
	Urns []common.Address
	Ilks [][32]byte
}, error) {
	return _Getcdps.Contract.GetCdpsDesc(&_Getcdps.CallOpts, manager, guy)
}

// GetCdpsDesc is a free data retrieval call binding the contract method 0x38f7acb4.
//
// Solidity: function getCdpsDesc(address manager, address guy) view returns(uint256[] ids, address[] urns, bytes32[] ilks)
func (_Getcdps *GetcdpsCallerSession) GetCdpsDesc(manager common.Address, guy common.Address) (struct {
	Ids  []*big.Int
	Urns []common.Address
	Ilks [][32]byte
}, error) {
	return _Getcdps.Contract.GetCdpsDesc(&_Getcdps.CallOpts, manager, guy)
}
