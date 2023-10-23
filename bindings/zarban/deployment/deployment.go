// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package deployment

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

// DeploymentMetaData contains all meta data concerning the Deployment contract.
var DeploymentMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractVatFab\",\"name\":\"vatFab_\",\"type\":\"address\"},{\"internalType\":\"contractJugFab\",\"name\":\"jugFab_\",\"type\":\"address\"},{\"internalType\":\"contractVowFab\",\"name\":\"vowFab_\",\"type\":\"address\"},{\"internalType\":\"contractDogFab\",\"name\":\"dogFab_\",\"type\":\"address\"},{\"internalType\":\"contractZarFab\",\"name\":\"zarFab_\",\"type\":\"address\"},{\"internalType\":\"contractZarJoinFab\",\"name\":\"zarJoinFab_\",\"type\":\"address\"}],\"name\":\"addFabs1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractClipFab\",\"name\":\"clipFab_\",\"type\":\"address\"},{\"internalType\":\"contractCalcFab\",\"name\":\"calcFab_\",\"type\":\"address\"},{\"internalType\":\"contractSpotFab\",\"name\":\"spotFab_\",\"type\":\"address\"},{\"internalType\":\"contractEndFab\",\"name\":\"endFab_\",\"type\":\"address\"}],\"name\":\"addFabs2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"internalType\":\"contractAuthority\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calcFab\",\"outputs\":[{\"internalType\":\"contractCalcFab\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clipFab\",\"outputs\":[{\"internalType\":\"contractClipFab\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gov\",\"type\":\"address\"}],\"name\":\"deployAuctions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"join\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pip\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"calc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"deployCollateralClip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployEnd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployLiquidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployTaxation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployVat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployZar\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dog\",\"outputs\":[{\"internalType\":\"contractDog\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dogFab\",\"outputs\":[{\"internalType\":\"contractDogFab\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"end\",\"outputs\":[{\"internalType\":\"contractEnd\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endFab\",\"outputs\":[{\"internalType\":\"contractEndFab\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ilks\",\"outputs\":[{\"internalType\":\"contractClipper\",\"name\":\"clip\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"join\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"jug\",\"outputs\":[{\"internalType\":\"contractJug\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"jugFab\",\"outputs\":[{\"internalType\":\"contractJugFab\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"releaseAuth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"releaseAuthClip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"relyAuthority\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAuthority\",\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spotFab\",\"outputs\":[{\"internalType\":\"contractSpotFab\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spotter\",\"outputs\":[{\"internalType\":\"contractSpotter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"step\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVat\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vatFab\",\"outputs\":[{\"internalType\":\"contractVatFab\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vow\",\"outputs\":[{\"internalType\":\"contractVow\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vowFab\",\"outputs\":[{\"internalType\":\"contractVowFab\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zar\",\"outputs\":[{\"internalType\":\"contractZar\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zarFab\",\"outputs\":[{\"internalType\":\"contractZarFab\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zarJoin\",\"outputs\":[{\"internalType\":\"contractZarJoin\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zarJoinFab\",\"outputs\":[{\"internalType\":\"contractZarJoinFab\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DeploymentABI is the input ABI used to generate the binding from.
// Deprecated: Use DeploymentMetaData.ABI instead.
var DeploymentABI = DeploymentMetaData.ABI

// Deployment is an auto generated Go binding around an Ethereum contract.
type Deployment struct {
	DeploymentCaller     // Read-only binding to the contract
	DeploymentTransactor // Write-only binding to the contract
	DeploymentFilterer   // Log filterer for contract events
}

// DeploymentCaller is an auto generated read-only Go binding around an Ethereum contract.
type DeploymentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeploymentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DeploymentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeploymentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DeploymentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeploymentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DeploymentSession struct {
	Contract     *Deployment       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DeploymentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DeploymentCallerSession struct {
	Contract *DeploymentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DeploymentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DeploymentTransactorSession struct {
	Contract     *DeploymentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DeploymentRaw is an auto generated low-level Go binding around an Ethereum contract.
type DeploymentRaw struct {
	Contract *Deployment // Generic contract binding to access the raw methods on
}

// DeploymentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DeploymentCallerRaw struct {
	Contract *DeploymentCaller // Generic read-only contract binding to access the raw methods on
}

// DeploymentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DeploymentTransactorRaw struct {
	Contract *DeploymentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeployment creates a new instance of Deployment, bound to a specific deployed contract.
func NewDeployment(address common.Address, backend bind.ContractBackend) (*Deployment, error) {
	contract, err := bindDeployment(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Deployment{DeploymentCaller: DeploymentCaller{contract: contract}, DeploymentTransactor: DeploymentTransactor{contract: contract}, DeploymentFilterer: DeploymentFilterer{contract: contract}}, nil
}

// NewDeploymentCaller creates a new read-only instance of Deployment, bound to a specific deployed contract.
func NewDeploymentCaller(address common.Address, caller bind.ContractCaller) (*DeploymentCaller, error) {
	contract, err := bindDeployment(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DeploymentCaller{contract: contract}, nil
}

// NewDeploymentTransactor creates a new write-only instance of Deployment, bound to a specific deployed contract.
func NewDeploymentTransactor(address common.Address, transactor bind.ContractTransactor) (*DeploymentTransactor, error) {
	contract, err := bindDeployment(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DeploymentTransactor{contract: contract}, nil
}

// NewDeploymentFilterer creates a new log filterer instance of Deployment, bound to a specific deployed contract.
func NewDeploymentFilterer(address common.Address, filterer bind.ContractFilterer) (*DeploymentFilterer, error) {
	contract, err := bindDeployment(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DeploymentFilterer{contract: contract}, nil
}

// bindDeployment binds a generic wrapper to an already deployed contract.
func bindDeployment(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DeploymentMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deployment *DeploymentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Deployment.Contract.DeploymentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deployment *DeploymentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deployment.Contract.DeploymentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deployment *DeploymentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deployment.Contract.DeploymentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deployment *DeploymentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Deployment.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deployment *DeploymentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deployment.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deployment *DeploymentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deployment.Contract.contract.Transact(opts, method, params...)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_Deployment *DeploymentCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deployment.contract.Call(opts, &out, "authority")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_Deployment *DeploymentSession) Authority() (common.Address, error) {
	return _Deployment.Contract.Authority(&_Deployment.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_Deployment *DeploymentCallerSession) Authority() (common.Address, error) {
	return _Deployment.Contract.Authority(&_Deployment.CallOpts)
}

// CalcFab is a free data retrieval call binding the contract method 0x82febc43.
//
// Solidity: function calcFab() view returns(address)
func (_Deployment *DeploymentCaller) CalcFab(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deployment.contract.Call(opts, &out, "calcFab")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CalcFab is a free data retrieval call binding the contract method 0x82febc43.
//
// Solidity: function calcFab() view returns(address)
func (_Deployment *DeploymentSession) CalcFab() (common.Address, error) {
	return _Deployment.Contract.CalcFab(&_Deployment.CallOpts)
}

// CalcFab is a free data retrieval call binding the contract method 0x82febc43.
//
// Solidity: function calcFab() view returns(address)
func (_Deployment *DeploymentCallerSession) CalcFab() (common.Address, error) {
	return _Deployment.Contract.CalcFab(&_Deployment.CallOpts)
}

// ClipFab is a free data retrieval call binding the contract method 0x2ffcd955.
//
// Solidity: function clipFab() view returns(address)
func (_Deployment *DeploymentCaller) ClipFab(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deployment.contract.Call(opts, &out, "clipFab")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ClipFab is a free data retrieval call binding the contract method 0x2ffcd955.
//
// Solidity: function clipFab() view returns(address)
func (_Deployment *DeploymentSession) ClipFab() (common.Address, error) {
	return _Deployment.Contract.ClipFab(&_Deployment.CallOpts)
}

// ClipFab is a free data retrieval call binding the contract method 0x2ffcd955.
//
// Solidity: function clipFab() view returns(address)
func (_Deployment *DeploymentCallerSession) ClipFab() (common.Address, error) {
	return _Deployment.Contract.ClipFab(&_Deployment.CallOpts)
}

// Dog is a free data retrieval call binding the contract method 0xc3b3ad7f.
//
// Solidity: function dog() view returns(address)
func (_Deployment *DeploymentCaller) Dog(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deployment.contract.Call(opts, &out, "dog")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dog is a free data retrieval call binding the contract method 0xc3b3ad7f.
//
// Solidity: function dog() view returns(address)
func (_Deployment *DeploymentSession) Dog() (common.Address, error) {
	return _Deployment.Contract.Dog(&_Deployment.CallOpts)
}

// Dog is a free data retrieval call binding the contract method 0xc3b3ad7f.
//
// Solidity: function dog() view returns(address)
func (_Deployment *DeploymentCallerSession) Dog() (common.Address, error) {
	return _Deployment.Contract.Dog(&_Deployment.CallOpts)
}

// DogFab is a free data retrieval call binding the contract method 0x0dfe2811.
//
// Solidity: function dogFab() view returns(address)
func (_Deployment *DeploymentCaller) DogFab(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deployment.contract.Call(opts, &out, "dogFab")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DogFab is a free data retrieval call binding the contract method 0x0dfe2811.
//
// Solidity: function dogFab() view returns(address)
func (_Deployment *DeploymentSession) DogFab() (common.Address, error) {
	return _Deployment.Contract.DogFab(&_Deployment.CallOpts)
}

// DogFab is a free data retrieval call binding the contract method 0x0dfe2811.
//
// Solidity: function dogFab() view returns(address)
func (_Deployment *DeploymentCallerSession) DogFab() (common.Address, error) {
	return _Deployment.Contract.DogFab(&_Deployment.CallOpts)
}

// End is a free data retrieval call binding the contract method 0xefbe1c1c.
//
// Solidity: function end() view returns(address)
func (_Deployment *DeploymentCaller) End(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deployment.contract.Call(opts, &out, "end")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// End is a free data retrieval call binding the contract method 0xefbe1c1c.
//
// Solidity: function end() view returns(address)
func (_Deployment *DeploymentSession) End() (common.Address, error) {
	return _Deployment.Contract.End(&_Deployment.CallOpts)
}

// End is a free data retrieval call binding the contract method 0xefbe1c1c.
//
// Solidity: function end() view returns(address)
func (_Deployment *DeploymentCallerSession) End() (common.Address, error) {
	return _Deployment.Contract.End(&_Deployment.CallOpts)
}

// EndFab is a free data retrieval call binding the contract method 0xef0add5b.
//
// Solidity: function endFab() view returns(address)
func (_Deployment *DeploymentCaller) EndFab(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deployment.contract.Call(opts, &out, "endFab")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EndFab is a free data retrieval call binding the contract method 0xef0add5b.
//
// Solidity: function endFab() view returns(address)
func (_Deployment *DeploymentSession) EndFab() (common.Address, error) {
	return _Deployment.Contract.EndFab(&_Deployment.CallOpts)
}

// EndFab is a free data retrieval call binding the contract method 0xef0add5b.
//
// Solidity: function endFab() view returns(address)
func (_Deployment *DeploymentCallerSession) EndFab() (common.Address, error) {
	return _Deployment.Contract.EndFab(&_Deployment.CallOpts)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(address clip, address join)
func (_Deployment *DeploymentCaller) Ilks(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Clip common.Address
	Join common.Address
}, error) {
	var out []interface{}
	err := _Deployment.contract.Call(opts, &out, "ilks", arg0)

	outstruct := new(struct {
		Clip common.Address
		Join common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Clip = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Join = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(address clip, address join)
func (_Deployment *DeploymentSession) Ilks(arg0 [32]byte) (struct {
	Clip common.Address
	Join common.Address
}, error) {
	return _Deployment.Contract.Ilks(&_Deployment.CallOpts, arg0)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(address clip, address join)
func (_Deployment *DeploymentCallerSession) Ilks(arg0 [32]byte) (struct {
	Clip common.Address
	Join common.Address
}, error) {
	return _Deployment.Contract.Ilks(&_Deployment.CallOpts, arg0)
}

// Jug is a free data retrieval call binding the contract method 0x84718d89.
//
// Solidity: function jug() view returns(address)
func (_Deployment *DeploymentCaller) Jug(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deployment.contract.Call(opts, &out, "jug")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Jug is a free data retrieval call binding the contract method 0x84718d89.
//
// Solidity: function jug() view returns(address)
func (_Deployment *DeploymentSession) Jug() (common.Address, error) {
	return _Deployment.Contract.Jug(&_Deployment.CallOpts)
}

// Jug is a free data retrieval call binding the contract method 0x84718d89.
//
// Solidity: function jug() view returns(address)
func (_Deployment *DeploymentCallerSession) Jug() (common.Address, error) {
	return _Deployment.Contract.Jug(&_Deployment.CallOpts)
}

// JugFab is a free data retrieval call binding the contract method 0x5f35a49d.
//
// Solidity: function jugFab() view returns(address)
func (_Deployment *DeploymentCaller) JugFab(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deployment.contract.Call(opts, &out, "jugFab")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// JugFab is a free data retrieval call binding the contract method 0x5f35a49d.
//
// Solidity: function jugFab() view returns(address)
func (_Deployment *DeploymentSession) JugFab() (common.Address, error) {
	return _Deployment.Contract.JugFab(&_Deployment.CallOpts)
}

// JugFab is a free data retrieval call binding the contract method 0x5f35a49d.
//
// Solidity: function jugFab() view returns(address)
func (_Deployment *DeploymentCallerSession) JugFab() (common.Address, error) {
	return _Deployment.Contract.JugFab(&_Deployment.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Deployment *DeploymentCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deployment.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Deployment *DeploymentSession) Owner() (common.Address, error) {
	return _Deployment.Contract.Owner(&_Deployment.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Deployment *DeploymentCallerSession) Owner() (common.Address, error) {
	return _Deployment.Contract.Owner(&_Deployment.CallOpts)
}

// SpotFab is a free data retrieval call binding the contract method 0x92ed5460.
//
// Solidity: function spotFab() view returns(address)
func (_Deployment *DeploymentCaller) SpotFab(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deployment.contract.Call(opts, &out, "spotFab")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SpotFab is a free data retrieval call binding the contract method 0x92ed5460.
//
// Solidity: function spotFab() view returns(address)
func (_Deployment *DeploymentSession) SpotFab() (common.Address, error) {
	return _Deployment.Contract.SpotFab(&_Deployment.CallOpts)
}

// SpotFab is a free data retrieval call binding the contract method 0x92ed5460.
//
// Solidity: function spotFab() view returns(address)
func (_Deployment *DeploymentCallerSession) SpotFab() (common.Address, error) {
	return _Deployment.Contract.SpotFab(&_Deployment.CallOpts)
}

// Spotter is a free data retrieval call binding the contract method 0x2e77468d.
//
// Solidity: function spotter() view returns(address)
func (_Deployment *DeploymentCaller) Spotter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deployment.contract.Call(opts, &out, "spotter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Spotter is a free data retrieval call binding the contract method 0x2e77468d.
//
// Solidity: function spotter() view returns(address)
func (_Deployment *DeploymentSession) Spotter() (common.Address, error) {
	return _Deployment.Contract.Spotter(&_Deployment.CallOpts)
}

// Spotter is a free data retrieval call binding the contract method 0x2e77468d.
//
// Solidity: function spotter() view returns(address)
func (_Deployment *DeploymentCallerSession) Spotter() (common.Address, error) {
	return _Deployment.Contract.Spotter(&_Deployment.CallOpts)
}

// Step is a free data retrieval call binding the contract method 0xe25fe175.
//
// Solidity: function step() view returns(uint8)
func (_Deployment *DeploymentCaller) Step(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Deployment.contract.Call(opts, &out, "step")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Step is a free data retrieval call binding the contract method 0xe25fe175.
//
// Solidity: function step() view returns(uint8)
func (_Deployment *DeploymentSession) Step() (uint8, error) {
	return _Deployment.Contract.Step(&_Deployment.CallOpts)
}

// Step is a free data retrieval call binding the contract method 0xe25fe175.
//
// Solidity: function step() view returns(uint8)
func (_Deployment *DeploymentCallerSession) Step() (uint8, error) {
	return _Deployment.Contract.Step(&_Deployment.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Deployment *DeploymentCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deployment.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Deployment *DeploymentSession) Vat() (common.Address, error) {
	return _Deployment.Contract.Vat(&_Deployment.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Deployment *DeploymentCallerSession) Vat() (common.Address, error) {
	return _Deployment.Contract.Vat(&_Deployment.CallOpts)
}

// VatFab is a free data retrieval call binding the contract method 0x4a847390.
//
// Solidity: function vatFab() view returns(address)
func (_Deployment *DeploymentCaller) VatFab(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deployment.contract.Call(opts, &out, "vatFab")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VatFab is a free data retrieval call binding the contract method 0x4a847390.
//
// Solidity: function vatFab() view returns(address)
func (_Deployment *DeploymentSession) VatFab() (common.Address, error) {
	return _Deployment.Contract.VatFab(&_Deployment.CallOpts)
}

// VatFab is a free data retrieval call binding the contract method 0x4a847390.
//
// Solidity: function vatFab() view returns(address)
func (_Deployment *DeploymentCallerSession) VatFab() (common.Address, error) {
	return _Deployment.Contract.VatFab(&_Deployment.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_Deployment *DeploymentCaller) Vow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deployment.contract.Call(opts, &out, "vow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_Deployment *DeploymentSession) Vow() (common.Address, error) {
	return _Deployment.Contract.Vow(&_Deployment.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_Deployment *DeploymentCallerSession) Vow() (common.Address, error) {
	return _Deployment.Contract.Vow(&_Deployment.CallOpts)
}

// VowFab is a free data retrieval call binding the contract method 0xb94fdb8e.
//
// Solidity: function vowFab() view returns(address)
func (_Deployment *DeploymentCaller) VowFab(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deployment.contract.Call(opts, &out, "vowFab")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VowFab is a free data retrieval call binding the contract method 0xb94fdb8e.
//
// Solidity: function vowFab() view returns(address)
func (_Deployment *DeploymentSession) VowFab() (common.Address, error) {
	return _Deployment.Contract.VowFab(&_Deployment.CallOpts)
}

// VowFab is a free data retrieval call binding the contract method 0xb94fdb8e.
//
// Solidity: function vowFab() view returns(address)
func (_Deployment *DeploymentCallerSession) VowFab() (common.Address, error) {
	return _Deployment.Contract.VowFab(&_Deployment.CallOpts)
}

// Zar is a free data retrieval call binding the contract method 0x7d510ab5.
//
// Solidity: function zar() view returns(address)
func (_Deployment *DeploymentCaller) Zar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deployment.contract.Call(opts, &out, "zar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Zar is a free data retrieval call binding the contract method 0x7d510ab5.
//
// Solidity: function zar() view returns(address)
func (_Deployment *DeploymentSession) Zar() (common.Address, error) {
	return _Deployment.Contract.Zar(&_Deployment.CallOpts)
}

// Zar is a free data retrieval call binding the contract method 0x7d510ab5.
//
// Solidity: function zar() view returns(address)
func (_Deployment *DeploymentCallerSession) Zar() (common.Address, error) {
	return _Deployment.Contract.Zar(&_Deployment.CallOpts)
}

// ZarFab is a free data retrieval call binding the contract method 0x9e6a4ff3.
//
// Solidity: function zarFab() view returns(address)
func (_Deployment *DeploymentCaller) ZarFab(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deployment.contract.Call(opts, &out, "zarFab")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZarFab is a free data retrieval call binding the contract method 0x9e6a4ff3.
//
// Solidity: function zarFab() view returns(address)
func (_Deployment *DeploymentSession) ZarFab() (common.Address, error) {
	return _Deployment.Contract.ZarFab(&_Deployment.CallOpts)
}

// ZarFab is a free data retrieval call binding the contract method 0x9e6a4ff3.
//
// Solidity: function zarFab() view returns(address)
func (_Deployment *DeploymentCallerSession) ZarFab() (common.Address, error) {
	return _Deployment.Contract.ZarFab(&_Deployment.CallOpts)
}

// ZarJoin is a free data retrieval call binding the contract method 0x3cc56376.
//
// Solidity: function zarJoin() view returns(address)
func (_Deployment *DeploymentCaller) ZarJoin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deployment.contract.Call(opts, &out, "zarJoin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZarJoin is a free data retrieval call binding the contract method 0x3cc56376.
//
// Solidity: function zarJoin() view returns(address)
func (_Deployment *DeploymentSession) ZarJoin() (common.Address, error) {
	return _Deployment.Contract.ZarJoin(&_Deployment.CallOpts)
}

// ZarJoin is a free data retrieval call binding the contract method 0x3cc56376.
//
// Solidity: function zarJoin() view returns(address)
func (_Deployment *DeploymentCallerSession) ZarJoin() (common.Address, error) {
	return _Deployment.Contract.ZarJoin(&_Deployment.CallOpts)
}

// ZarJoinFab is a free data retrieval call binding the contract method 0x0643e09f.
//
// Solidity: function zarJoinFab() view returns(address)
func (_Deployment *DeploymentCaller) ZarJoinFab(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deployment.contract.Call(opts, &out, "zarJoinFab")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZarJoinFab is a free data retrieval call binding the contract method 0x0643e09f.
//
// Solidity: function zarJoinFab() view returns(address)
func (_Deployment *DeploymentSession) ZarJoinFab() (common.Address, error) {
	return _Deployment.Contract.ZarJoinFab(&_Deployment.CallOpts)
}

// ZarJoinFab is a free data retrieval call binding the contract method 0x0643e09f.
//
// Solidity: function zarJoinFab() view returns(address)
func (_Deployment *DeploymentCallerSession) ZarJoinFab() (common.Address, error) {
	return _Deployment.Contract.ZarJoinFab(&_Deployment.CallOpts)
}

// AddFabs1 is a paid mutator transaction binding the contract method 0x27c60747.
//
// Solidity: function addFabs1(address vatFab_, address jugFab_, address vowFab_, address dogFab_, address zarFab_, address zarJoinFab_) returns()
func (_Deployment *DeploymentTransactor) AddFabs1(opts *bind.TransactOpts, vatFab_ common.Address, jugFab_ common.Address, vowFab_ common.Address, dogFab_ common.Address, zarFab_ common.Address, zarJoinFab_ common.Address) (*types.Transaction, error) {
	return _Deployment.contract.Transact(opts, "addFabs1", vatFab_, jugFab_, vowFab_, dogFab_, zarFab_, zarJoinFab_)
}

// AddFabs1 is a paid mutator transaction binding the contract method 0x27c60747.
//
// Solidity: function addFabs1(address vatFab_, address jugFab_, address vowFab_, address dogFab_, address zarFab_, address zarJoinFab_) returns()
func (_Deployment *DeploymentSession) AddFabs1(vatFab_ common.Address, jugFab_ common.Address, vowFab_ common.Address, dogFab_ common.Address, zarFab_ common.Address, zarJoinFab_ common.Address) (*types.Transaction, error) {
	return _Deployment.Contract.AddFabs1(&_Deployment.TransactOpts, vatFab_, jugFab_, vowFab_, dogFab_, zarFab_, zarJoinFab_)
}

// AddFabs1 is a paid mutator transaction binding the contract method 0x27c60747.
//
// Solidity: function addFabs1(address vatFab_, address jugFab_, address vowFab_, address dogFab_, address zarFab_, address zarJoinFab_) returns()
func (_Deployment *DeploymentTransactorSession) AddFabs1(vatFab_ common.Address, jugFab_ common.Address, vowFab_ common.Address, dogFab_ common.Address, zarFab_ common.Address, zarJoinFab_ common.Address) (*types.Transaction, error) {
	return _Deployment.Contract.AddFabs1(&_Deployment.TransactOpts, vatFab_, jugFab_, vowFab_, dogFab_, zarFab_, zarJoinFab_)
}

// AddFabs2 is a paid mutator transaction binding the contract method 0x825f4573.
//
// Solidity: function addFabs2(address clipFab_, address calcFab_, address spotFab_, address endFab_) returns()
func (_Deployment *DeploymentTransactor) AddFabs2(opts *bind.TransactOpts, clipFab_ common.Address, calcFab_ common.Address, spotFab_ common.Address, endFab_ common.Address) (*types.Transaction, error) {
	return _Deployment.contract.Transact(opts, "addFabs2", clipFab_, calcFab_, spotFab_, endFab_)
}

// AddFabs2 is a paid mutator transaction binding the contract method 0x825f4573.
//
// Solidity: function addFabs2(address clipFab_, address calcFab_, address spotFab_, address endFab_) returns()
func (_Deployment *DeploymentSession) AddFabs2(clipFab_ common.Address, calcFab_ common.Address, spotFab_ common.Address, endFab_ common.Address) (*types.Transaction, error) {
	return _Deployment.Contract.AddFabs2(&_Deployment.TransactOpts, clipFab_, calcFab_, spotFab_, endFab_)
}

// AddFabs2 is a paid mutator transaction binding the contract method 0x825f4573.
//
// Solidity: function addFabs2(address clipFab_, address calcFab_, address spotFab_, address endFab_) returns()
func (_Deployment *DeploymentTransactorSession) AddFabs2(clipFab_ common.Address, calcFab_ common.Address, spotFab_ common.Address, endFab_ common.Address) (*types.Transaction, error) {
	return _Deployment.Contract.AddFabs2(&_Deployment.TransactOpts, clipFab_, calcFab_, spotFab_, endFab_)
}

// DeployAuctions is a paid mutator transaction binding the contract method 0xcc3f5ec1.
//
// Solidity: function deployAuctions(address gov) returns()
func (_Deployment *DeploymentTransactor) DeployAuctions(opts *bind.TransactOpts, gov common.Address) (*types.Transaction, error) {
	return _Deployment.contract.Transact(opts, "deployAuctions", gov)
}

// DeployAuctions is a paid mutator transaction binding the contract method 0xcc3f5ec1.
//
// Solidity: function deployAuctions(address gov) returns()
func (_Deployment *DeploymentSession) DeployAuctions(gov common.Address) (*types.Transaction, error) {
	return _Deployment.Contract.DeployAuctions(&_Deployment.TransactOpts, gov)
}

// DeployAuctions is a paid mutator transaction binding the contract method 0xcc3f5ec1.
//
// Solidity: function deployAuctions(address gov) returns()
func (_Deployment *DeploymentTransactorSession) DeployAuctions(gov common.Address) (*types.Transaction, error) {
	return _Deployment.Contract.DeployAuctions(&_Deployment.TransactOpts, gov)
}

// DeployCollateralClip is a paid mutator transaction binding the contract method 0x5e9d66a8.
//
// Solidity: function deployCollateralClip(bytes32 ilk, address join, address pip, address calc, address authority) returns()
func (_Deployment *DeploymentTransactor) DeployCollateralClip(opts *bind.TransactOpts, ilk [32]byte, join common.Address, pip common.Address, calc common.Address, authority common.Address) (*types.Transaction, error) {
	return _Deployment.contract.Transact(opts, "deployCollateralClip", ilk, join, pip, calc, authority)
}

// DeployCollateralClip is a paid mutator transaction binding the contract method 0x5e9d66a8.
//
// Solidity: function deployCollateralClip(bytes32 ilk, address join, address pip, address calc, address authority) returns()
func (_Deployment *DeploymentSession) DeployCollateralClip(ilk [32]byte, join common.Address, pip common.Address, calc common.Address, authority common.Address) (*types.Transaction, error) {
	return _Deployment.Contract.DeployCollateralClip(&_Deployment.TransactOpts, ilk, join, pip, calc, authority)
}

// DeployCollateralClip is a paid mutator transaction binding the contract method 0x5e9d66a8.
//
// Solidity: function deployCollateralClip(bytes32 ilk, address join, address pip, address calc, address authority) returns()
func (_Deployment *DeploymentTransactorSession) DeployCollateralClip(ilk [32]byte, join common.Address, pip common.Address, calc common.Address, authority common.Address) (*types.Transaction, error) {
	return _Deployment.Contract.DeployCollateralClip(&_Deployment.TransactOpts, ilk, join, pip, calc, authority)
}

// DeployEnd is a paid mutator transaction binding the contract method 0x59f1cb8a.
//
// Solidity: function deployEnd() returns()
func (_Deployment *DeploymentTransactor) DeployEnd(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deployment.contract.Transact(opts, "deployEnd")
}

// DeployEnd is a paid mutator transaction binding the contract method 0x59f1cb8a.
//
// Solidity: function deployEnd() returns()
func (_Deployment *DeploymentSession) DeployEnd() (*types.Transaction, error) {
	return _Deployment.Contract.DeployEnd(&_Deployment.TransactOpts)
}

// DeployEnd is a paid mutator transaction binding the contract method 0x59f1cb8a.
//
// Solidity: function deployEnd() returns()
func (_Deployment *DeploymentTransactorSession) DeployEnd() (*types.Transaction, error) {
	return _Deployment.Contract.DeployEnd(&_Deployment.TransactOpts)
}

// DeployLiquidator is a paid mutator transaction binding the contract method 0xdcc91c9b.
//
// Solidity: function deployLiquidator() returns()
func (_Deployment *DeploymentTransactor) DeployLiquidator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deployment.contract.Transact(opts, "deployLiquidator")
}

// DeployLiquidator is a paid mutator transaction binding the contract method 0xdcc91c9b.
//
// Solidity: function deployLiquidator() returns()
func (_Deployment *DeploymentSession) DeployLiquidator() (*types.Transaction, error) {
	return _Deployment.Contract.DeployLiquidator(&_Deployment.TransactOpts)
}

// DeployLiquidator is a paid mutator transaction binding the contract method 0xdcc91c9b.
//
// Solidity: function deployLiquidator() returns()
func (_Deployment *DeploymentTransactorSession) DeployLiquidator() (*types.Transaction, error) {
	return _Deployment.Contract.DeployLiquidator(&_Deployment.TransactOpts)
}

// DeployTaxation is a paid mutator transaction binding the contract method 0x80e7cd89.
//
// Solidity: function deployTaxation() returns()
func (_Deployment *DeploymentTransactor) DeployTaxation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deployment.contract.Transact(opts, "deployTaxation")
}

// DeployTaxation is a paid mutator transaction binding the contract method 0x80e7cd89.
//
// Solidity: function deployTaxation() returns()
func (_Deployment *DeploymentSession) DeployTaxation() (*types.Transaction, error) {
	return _Deployment.Contract.DeployTaxation(&_Deployment.TransactOpts)
}

// DeployTaxation is a paid mutator transaction binding the contract method 0x80e7cd89.
//
// Solidity: function deployTaxation() returns()
func (_Deployment *DeploymentTransactorSession) DeployTaxation() (*types.Transaction, error) {
	return _Deployment.Contract.DeployTaxation(&_Deployment.TransactOpts)
}

// DeployVat is a paid mutator transaction binding the contract method 0x2800a568.
//
// Solidity: function deployVat() returns()
func (_Deployment *DeploymentTransactor) DeployVat(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deployment.contract.Transact(opts, "deployVat")
}

// DeployVat is a paid mutator transaction binding the contract method 0x2800a568.
//
// Solidity: function deployVat() returns()
func (_Deployment *DeploymentSession) DeployVat() (*types.Transaction, error) {
	return _Deployment.Contract.DeployVat(&_Deployment.TransactOpts)
}

// DeployVat is a paid mutator transaction binding the contract method 0x2800a568.
//
// Solidity: function deployVat() returns()
func (_Deployment *DeploymentTransactorSession) DeployVat() (*types.Transaction, error) {
	return _Deployment.Contract.DeployVat(&_Deployment.TransactOpts)
}

// DeployZar is a paid mutator transaction binding the contract method 0xf4da4ab2.
//
// Solidity: function deployZar() returns()
func (_Deployment *DeploymentTransactor) DeployZar(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deployment.contract.Transact(opts, "deployZar")
}

// DeployZar is a paid mutator transaction binding the contract method 0xf4da4ab2.
//
// Solidity: function deployZar() returns()
func (_Deployment *DeploymentSession) DeployZar() (*types.Transaction, error) {
	return _Deployment.Contract.DeployZar(&_Deployment.TransactOpts)
}

// DeployZar is a paid mutator transaction binding the contract method 0xf4da4ab2.
//
// Solidity: function deployZar() returns()
func (_Deployment *DeploymentTransactorSession) DeployZar() (*types.Transaction, error) {
	return _Deployment.Contract.DeployZar(&_Deployment.TransactOpts)
}

// ReleaseAuth is a paid mutator transaction binding the contract method 0x70402bb9.
//
// Solidity: function releaseAuth() returns()
func (_Deployment *DeploymentTransactor) ReleaseAuth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deployment.contract.Transact(opts, "releaseAuth")
}

// ReleaseAuth is a paid mutator transaction binding the contract method 0x70402bb9.
//
// Solidity: function releaseAuth() returns()
func (_Deployment *DeploymentSession) ReleaseAuth() (*types.Transaction, error) {
	return _Deployment.Contract.ReleaseAuth(&_Deployment.TransactOpts)
}

// ReleaseAuth is a paid mutator transaction binding the contract method 0x70402bb9.
//
// Solidity: function releaseAuth() returns()
func (_Deployment *DeploymentTransactorSession) ReleaseAuth() (*types.Transaction, error) {
	return _Deployment.Contract.ReleaseAuth(&_Deployment.TransactOpts)
}

// ReleaseAuthClip is a paid mutator transaction binding the contract method 0x45b09ba5.
//
// Solidity: function releaseAuthClip(bytes32 ilk) returns()
func (_Deployment *DeploymentTransactor) ReleaseAuthClip(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _Deployment.contract.Transact(opts, "releaseAuthClip", ilk)
}

// ReleaseAuthClip is a paid mutator transaction binding the contract method 0x45b09ba5.
//
// Solidity: function releaseAuthClip(bytes32 ilk) returns()
func (_Deployment *DeploymentSession) ReleaseAuthClip(ilk [32]byte) (*types.Transaction, error) {
	return _Deployment.Contract.ReleaseAuthClip(&_Deployment.TransactOpts, ilk)
}

// ReleaseAuthClip is a paid mutator transaction binding the contract method 0x45b09ba5.
//
// Solidity: function releaseAuthClip(bytes32 ilk) returns()
func (_Deployment *DeploymentTransactorSession) ReleaseAuthClip(ilk [32]byte) (*types.Transaction, error) {
	return _Deployment.Contract.ReleaseAuthClip(&_Deployment.TransactOpts, ilk)
}

// RelyAuthority is a paid mutator transaction binding the contract method 0x420411f4.
//
// Solidity: function relyAuthority(address authority) returns()
func (_Deployment *DeploymentTransactor) RelyAuthority(opts *bind.TransactOpts, authority common.Address) (*types.Transaction, error) {
	return _Deployment.contract.Transact(opts, "relyAuthority", authority)
}

// RelyAuthority is a paid mutator transaction binding the contract method 0x420411f4.
//
// Solidity: function relyAuthority(address authority) returns()
func (_Deployment *DeploymentSession) RelyAuthority(authority common.Address) (*types.Transaction, error) {
	return _Deployment.Contract.RelyAuthority(&_Deployment.TransactOpts, authority)
}

// RelyAuthority is a paid mutator transaction binding the contract method 0x420411f4.
//
// Solidity: function relyAuthority(address authority) returns()
func (_Deployment *DeploymentTransactorSession) RelyAuthority(authority common.Address) (*types.Transaction, error) {
	return _Deployment.Contract.RelyAuthority(&_Deployment.TransactOpts, authority)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_Deployment *DeploymentTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _Deployment.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_Deployment *DeploymentSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Deployment.Contract.SetAuthority(&_Deployment.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_Deployment *DeploymentTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Deployment.Contract.SetAuthority(&_Deployment.TransactOpts, authority_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_Deployment *DeploymentTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _Deployment.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_Deployment *DeploymentSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Deployment.Contract.SetOwner(&_Deployment.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_Deployment *DeploymentTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Deployment.Contract.SetOwner(&_Deployment.TransactOpts, owner_)
}

// DeploymentLogSetAuthorityIterator is returned from FilterLogSetAuthority and is used to iterate over the raw logs and unpacked data for LogSetAuthority events raised by the Deployment contract.
type DeploymentLogSetAuthorityIterator struct {
	Event *DeploymentLogSetAuthority // Event containing the contract specifics and raw log

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
func (it *DeploymentLogSetAuthorityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeploymentLogSetAuthority)
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
		it.Event = new(DeploymentLogSetAuthority)
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
func (it *DeploymentLogSetAuthorityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeploymentLogSetAuthorityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeploymentLogSetAuthority represents a LogSetAuthority event raised by the Deployment contract.
type DeploymentLogSetAuthority struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogSetAuthority is a free log retrieval operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_Deployment *DeploymentFilterer) FilterLogSetAuthority(opts *bind.FilterOpts, authority []common.Address) (*DeploymentLogSetAuthorityIterator, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _Deployment.contract.FilterLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return &DeploymentLogSetAuthorityIterator{contract: _Deployment.contract, event: "LogSetAuthority", logs: logs, sub: sub}, nil
}

// WatchLogSetAuthority is a free log subscription operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_Deployment *DeploymentFilterer) WatchLogSetAuthority(opts *bind.WatchOpts, sink chan<- *DeploymentLogSetAuthority, authority []common.Address) (event.Subscription, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _Deployment.contract.WatchLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeploymentLogSetAuthority)
				if err := _Deployment.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
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
func (_Deployment *DeploymentFilterer) ParseLogSetAuthority(log types.Log) (*DeploymentLogSetAuthority, error) {
	event := new(DeploymentLogSetAuthority)
	if err := _Deployment.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeploymentLogSetOwnerIterator is returned from FilterLogSetOwner and is used to iterate over the raw logs and unpacked data for LogSetOwner events raised by the Deployment contract.
type DeploymentLogSetOwnerIterator struct {
	Event *DeploymentLogSetOwner // Event containing the contract specifics and raw log

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
func (it *DeploymentLogSetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeploymentLogSetOwner)
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
		it.Event = new(DeploymentLogSetOwner)
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
func (it *DeploymentLogSetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeploymentLogSetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeploymentLogSetOwner represents a LogSetOwner event raised by the Deployment contract.
type DeploymentLogSetOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogSetOwner is a free log retrieval operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_Deployment *DeploymentFilterer) FilterLogSetOwner(opts *bind.FilterOpts, owner []common.Address) (*DeploymentLogSetOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Deployment.contract.FilterLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &DeploymentLogSetOwnerIterator{contract: _Deployment.contract, event: "LogSetOwner", logs: logs, sub: sub}, nil
}

// WatchLogSetOwner is a free log subscription operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_Deployment *DeploymentFilterer) WatchLogSetOwner(opts *bind.WatchOpts, sink chan<- *DeploymentLogSetOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Deployment.contract.WatchLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeploymentLogSetOwner)
				if err := _Deployment.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
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
func (_Deployment *DeploymentFilterer) ParseLogSetOwner(log types.Log) (*DeploymentLogSetOwner, error) {
	event := new(DeploymentLogSetOwner)
	if err := _Deployment.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
