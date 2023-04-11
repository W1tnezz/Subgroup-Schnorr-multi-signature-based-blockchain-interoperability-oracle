// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package test

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
)

// OracleContractEnrollNode is an auto generated low-level Go binding around an user-defined struct.
type OracleContractEnrollNode struct {
	Addr  common.Address
	Index *big.Int
}

// OracleContractMetaData contains all meta data concerning the OracleContract contract.
var OracleContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registryContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ValidationBegin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumOracleContract.ValidationType\",\"name\":\"typ\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"ValidationRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumOracleContract.ValidationType\",\"name\":\"typ\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"ValidationResponse\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASE_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CHANLLENGE_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COMPENSATION_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EnrollOracleNode\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOTAL_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chanllenge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"countEnrollNodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"findBlockValidationResult\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"findEnrollNodeByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structOracleContract.EnrollNode\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"findTransactionValidationResult\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fine\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"oracleNodeIsEnroll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"aggregateAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"submitBlockValidationResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"aggregateAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"submitTransactionValidationResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"validateBlock\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"validateTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// OracleContractABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleContractMetaData.ABI instead.
var OracleContractABI = OracleContractMetaData.ABI

// OracleContract is an auto generated Go binding around an Ethereum contract.
type OracleContract struct {
	OracleContractCaller     // Read-only binding to the contract
	OracleContractTransactor // Write-only binding to the contract
	OracleContractFilterer   // Log filterer for contract events
}

// OracleContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleContractSession struct {
	Contract     *OracleContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleContractCallerSession struct {
	Contract *OracleContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// OracleContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleContractTransactorSession struct {
	Contract     *OracleContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// OracleContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleContractRaw struct {
	Contract *OracleContract // Generic contract binding to access the raw methods on
}

// OracleContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleContractCallerRaw struct {
	Contract *OracleContractCaller // Generic read-only contract binding to access the raw methods on
}

// OracleContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleContractTransactorRaw struct {
	Contract *OracleContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracleContract creates a new instance of OracleContract, bound to a specific deployed contract.
func NewOracleContract(address common.Address, backend bind.ContractBackend) (*OracleContract, error) {
	contract, err := bindOracleContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OracleContract{OracleContractCaller: OracleContractCaller{contract: contract}, OracleContractTransactor: OracleContractTransactor{contract: contract}, OracleContractFilterer: OracleContractFilterer{contract: contract}}, nil
}

// NewOracleContractCaller creates a new read-only instance of OracleContract, bound to a specific deployed contract.
func NewOracleContractCaller(address common.Address, caller bind.ContractCaller) (*OracleContractCaller, error) {
	contract, err := bindOracleContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleContractCaller{contract: contract}, nil
}

// NewOracleContractTransactor creates a new write-only instance of OracleContract, bound to a specific deployed contract.
func NewOracleContractTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleContractTransactor, error) {
	contract, err := bindOracleContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleContractTransactor{contract: contract}, nil
}

// NewOracleContractFilterer creates a new log filterer instance of OracleContract, bound to a specific deployed contract.
func NewOracleContractFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleContractFilterer, error) {
	contract, err := bindOracleContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleContractFilterer{contract: contract}, nil
}

// bindOracleContract binds a generic wrapper to an already deployed contract.
func bindOracleContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleContract *OracleContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleContract.Contract.OracleContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleContract *OracleContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleContract.Contract.OracleContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleContract *OracleContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleContract.Contract.OracleContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleContract *OracleContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleContract *OracleContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleContract *OracleContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleContract.Contract.contract.Transact(opts, method, params...)
}

// BASEFEE is a free data retrieval call binding the contract method 0x3d18651e.
//
// Solidity: function BASE_FEE() view returns(uint256)
func (_OracleContract *OracleContractCaller) BASEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "BASE_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASEFEE is a free data retrieval call binding the contract method 0x3d18651e.
//
// Solidity: function BASE_FEE() view returns(uint256)
func (_OracleContract *OracleContractSession) BASEFEE() (*big.Int, error) {
	return _OracleContract.Contract.BASEFEE(&_OracleContract.CallOpts)
}

// BASEFEE is a free data retrieval call binding the contract method 0x3d18651e.
//
// Solidity: function BASE_FEE() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) BASEFEE() (*big.Int, error) {
	return _OracleContract.Contract.BASEFEE(&_OracleContract.CallOpts)
}

// CHANLLENGEFEE is a free data retrieval call binding the contract method 0xa2ff2ad7.
//
// Solidity: function CHANLLENGE_FEE() view returns(uint256)
func (_OracleContract *OracleContractCaller) CHANLLENGEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "CHANLLENGE_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CHANLLENGEFEE is a free data retrieval call binding the contract method 0xa2ff2ad7.
//
// Solidity: function CHANLLENGE_FEE() view returns(uint256)
func (_OracleContract *OracleContractSession) CHANLLENGEFEE() (*big.Int, error) {
	return _OracleContract.Contract.CHANLLENGEFEE(&_OracleContract.CallOpts)
}

// CHANLLENGEFEE is a free data retrieval call binding the contract method 0xa2ff2ad7.
//
// Solidity: function CHANLLENGE_FEE() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) CHANLLENGEFEE() (*big.Int, error) {
	return _OracleContract.Contract.CHANLLENGEFEE(&_OracleContract.CallOpts)
}

// COMPENSATIONFEE is a free data retrieval call binding the contract method 0xab3be840.
//
// Solidity: function COMPENSATION_FEE() view returns(uint256)
func (_OracleContract *OracleContractCaller) COMPENSATIONFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "COMPENSATION_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COMPENSATIONFEE is a free data retrieval call binding the contract method 0xab3be840.
//
// Solidity: function COMPENSATION_FEE() view returns(uint256)
func (_OracleContract *OracleContractSession) COMPENSATIONFEE() (*big.Int, error) {
	return _OracleContract.Contract.COMPENSATIONFEE(&_OracleContract.CallOpts)
}

// COMPENSATIONFEE is a free data retrieval call binding the contract method 0xab3be840.
//
// Solidity: function COMPENSATION_FEE() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) COMPENSATIONFEE() (*big.Int, error) {
	return _OracleContract.Contract.COMPENSATIONFEE(&_OracleContract.CallOpts)
}

// TOTALFEE is a free data retrieval call binding the contract method 0x63db7eae.
//
// Solidity: function TOTAL_FEE() view returns(uint256)
func (_OracleContract *OracleContractCaller) TOTALFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "TOTAL_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOTALFEE is a free data retrieval call binding the contract method 0x63db7eae.
//
// Solidity: function TOTAL_FEE() view returns(uint256)
func (_OracleContract *OracleContractSession) TOTALFEE() (*big.Int, error) {
	return _OracleContract.Contract.TOTALFEE(&_OracleContract.CallOpts)
}

// TOTALFEE is a free data retrieval call binding the contract method 0x63db7eae.
//
// Solidity: function TOTAL_FEE() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) TOTALFEE() (*big.Int, error) {
	return _OracleContract.Contract.TOTALFEE(&_OracleContract.CallOpts)
}

// VALIDATORFEE is a free data retrieval call binding the contract method 0x7da83e2b.
//
// Solidity: function VALIDATOR_FEE() view returns(uint256)
func (_OracleContract *OracleContractCaller) VALIDATORFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "VALIDATOR_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VALIDATORFEE is a free data retrieval call binding the contract method 0x7da83e2b.
//
// Solidity: function VALIDATOR_FEE() view returns(uint256)
func (_OracleContract *OracleContractSession) VALIDATORFEE() (*big.Int, error) {
	return _OracleContract.Contract.VALIDATORFEE(&_OracleContract.CallOpts)
}

// VALIDATORFEE is a free data retrieval call binding the contract method 0x7da83e2b.
//
// Solidity: function VALIDATOR_FEE() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) VALIDATORFEE() (*big.Int, error) {
	return _OracleContract.Contract.VALIDATORFEE(&_OracleContract.CallOpts)
}

// CountEnrollNodes is a free data retrieval call binding the contract method 0xfc769ee9.
//
// Solidity: function countEnrollNodes() view returns(uint256)
func (_OracleContract *OracleContractCaller) CountEnrollNodes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "countEnrollNodes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountEnrollNodes is a free data retrieval call binding the contract method 0xfc769ee9.
//
// Solidity: function countEnrollNodes() view returns(uint256)
func (_OracleContract *OracleContractSession) CountEnrollNodes() (*big.Int, error) {
	return _OracleContract.Contract.CountEnrollNodes(&_OracleContract.CallOpts)
}

// CountEnrollNodes is a free data retrieval call binding the contract method 0xfc769ee9.
//
// Solidity: function countEnrollNodes() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) CountEnrollNodes() (*big.Int, error) {
	return _OracleContract.Contract.CountEnrollNodes(&_OracleContract.CallOpts)
}

// FindBlockValidationResult is a free data retrieval call binding the contract method 0x3d27ef97.
//
// Solidity: function findBlockValidationResult(bytes32 _hash) view returns(bool)
func (_OracleContract *OracleContractCaller) FindBlockValidationResult(opts *bind.CallOpts, _hash [32]byte) (bool, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "findBlockValidationResult", _hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FindBlockValidationResult is a free data retrieval call binding the contract method 0x3d27ef97.
//
// Solidity: function findBlockValidationResult(bytes32 _hash) view returns(bool)
func (_OracleContract *OracleContractSession) FindBlockValidationResult(_hash [32]byte) (bool, error) {
	return _OracleContract.Contract.FindBlockValidationResult(&_OracleContract.CallOpts, _hash)
}

// FindBlockValidationResult is a free data retrieval call binding the contract method 0x3d27ef97.
//
// Solidity: function findBlockValidationResult(bytes32 _hash) view returns(bool)
func (_OracleContract *OracleContractCallerSession) FindBlockValidationResult(_hash [32]byte) (bool, error) {
	return _OracleContract.Contract.FindBlockValidationResult(&_OracleContract.CallOpts, _hash)
}

// FindEnrollNodeByIndex is a free data retrieval call binding the contract method 0x7e985fd0.
//
// Solidity: function findEnrollNodeByIndex(uint256 _index) view returns((address,uint256))
func (_OracleContract *OracleContractCaller) FindEnrollNodeByIndex(opts *bind.CallOpts, _index *big.Int) (OracleContractEnrollNode, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "findEnrollNodeByIndex", _index)

	if err != nil {
		return *new(OracleContractEnrollNode), err
	}

	out0 := *abi.ConvertType(out[0], new(OracleContractEnrollNode)).(*OracleContractEnrollNode)

	return out0, err

}

// FindEnrollNodeByIndex is a free data retrieval call binding the contract method 0x7e985fd0.
//
// Solidity: function findEnrollNodeByIndex(uint256 _index) view returns((address,uint256))
func (_OracleContract *OracleContractSession) FindEnrollNodeByIndex(_index *big.Int) (OracleContractEnrollNode, error) {
	return _OracleContract.Contract.FindEnrollNodeByIndex(&_OracleContract.CallOpts, _index)
}

// FindEnrollNodeByIndex is a free data retrieval call binding the contract method 0x7e985fd0.
//
// Solidity: function findEnrollNodeByIndex(uint256 _index) view returns((address,uint256))
func (_OracleContract *OracleContractCallerSession) FindEnrollNodeByIndex(_index *big.Int) (OracleContractEnrollNode, error) {
	return _OracleContract.Contract.FindEnrollNodeByIndex(&_OracleContract.CallOpts, _index)
}

// FindTransactionValidationResult is a free data retrieval call binding the contract method 0x43434590.
//
// Solidity: function findTransactionValidationResult(bytes32 _hash) view returns(bool)
func (_OracleContract *OracleContractCaller) FindTransactionValidationResult(opts *bind.CallOpts, _hash [32]byte) (bool, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "findTransactionValidationResult", _hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FindTransactionValidationResult is a free data retrieval call binding the contract method 0x43434590.
//
// Solidity: function findTransactionValidationResult(bytes32 _hash) view returns(bool)
func (_OracleContract *OracleContractSession) FindTransactionValidationResult(_hash [32]byte) (bool, error) {
	return _OracleContract.Contract.FindTransactionValidationResult(&_OracleContract.CallOpts, _hash)
}

// FindTransactionValidationResult is a free data retrieval call binding the contract method 0x43434590.
//
// Solidity: function findTransactionValidationResult(bytes32 _hash) view returns(bool)
func (_OracleContract *OracleContractCallerSession) FindTransactionValidationResult(_hash [32]byte) (bool, error) {
	return _OracleContract.Contract.FindTransactionValidationResult(&_OracleContract.CallOpts, _hash)
}

// OracleNodeIsEnroll is a free data retrieval call binding the contract method 0xe92c04d1.
//
// Solidity: function oracleNodeIsEnroll(address _addr) view returns(bool)
func (_OracleContract *OracleContractCaller) OracleNodeIsEnroll(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "oracleNodeIsEnroll", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OracleNodeIsEnroll is a free data retrieval call binding the contract method 0xe92c04d1.
//
// Solidity: function oracleNodeIsEnroll(address _addr) view returns(bool)
func (_OracleContract *OracleContractSession) OracleNodeIsEnroll(_addr common.Address) (bool, error) {
	return _OracleContract.Contract.OracleNodeIsEnroll(&_OracleContract.CallOpts, _addr)
}

// OracleNodeIsEnroll is a free data retrieval call binding the contract method 0xe92c04d1.
//
// Solidity: function oracleNodeIsEnroll(address _addr) view returns(bool)
func (_OracleContract *OracleContractCallerSession) OracleNodeIsEnroll(_addr common.Address) (bool, error) {
	return _OracleContract.Contract.OracleNodeIsEnroll(&_OracleContract.CallOpts, _addr)
}

// EnrollOracleNode is a paid mutator transaction binding the contract method 0x12f1e8f4.
//
// Solidity: function EnrollOracleNode() payable returns()
func (_OracleContract *OracleContractTransactor) EnrollOracleNode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "EnrollOracleNode")
}

// EnrollOracleNode is a paid mutator transaction binding the contract method 0x12f1e8f4.
//
// Solidity: function EnrollOracleNode() payable returns()
func (_OracleContract *OracleContractSession) EnrollOracleNode() (*types.Transaction, error) {
	return _OracleContract.Contract.EnrollOracleNode(&_OracleContract.TransactOpts)
}

// EnrollOracleNode is a paid mutator transaction binding the contract method 0x12f1e8f4.
//
// Solidity: function EnrollOracleNode() payable returns()
func (_OracleContract *OracleContractTransactorSession) EnrollOracleNode() (*types.Transaction, error) {
	return _OracleContract.Contract.EnrollOracleNode(&_OracleContract.TransactOpts)
}

// Chanllenge is a paid mutator transaction binding the contract method 0x460e99f4.
//
// Solidity: function chanllenge() payable returns()
func (_OracleContract *OracleContractTransactor) Chanllenge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "chanllenge")
}

// Chanllenge is a paid mutator transaction binding the contract method 0x460e99f4.
//
// Solidity: function chanllenge() payable returns()
func (_OracleContract *OracleContractSession) Chanllenge() (*types.Transaction, error) {
	return _OracleContract.Contract.Chanllenge(&_OracleContract.TransactOpts)
}

// Chanllenge is a paid mutator transaction binding the contract method 0x460e99f4.
//
// Solidity: function chanllenge() payable returns()
func (_OracleContract *OracleContractTransactorSession) Chanllenge() (*types.Transaction, error) {
	return _OracleContract.Contract.Chanllenge(&_OracleContract.TransactOpts)
}

// Fine is a paid mutator transaction binding the contract method 0xcf21c293.
//
// Solidity: function fine() payable returns()
func (_OracleContract *OracleContractTransactor) Fine(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "fine")
}

// Fine is a paid mutator transaction binding the contract method 0xcf21c293.
//
// Solidity: function fine() payable returns()
func (_OracleContract *OracleContractSession) Fine() (*types.Transaction, error) {
	return _OracleContract.Contract.Fine(&_OracleContract.TransactOpts)
}

// Fine is a paid mutator transaction binding the contract method 0xcf21c293.
//
// Solidity: function fine() payable returns()
func (_OracleContract *OracleContractTransactorSession) Fine() (*types.Transaction, error) {
	return _OracleContract.Contract.Fine(&_OracleContract.TransactOpts)
}

// SubmitBlockValidationResult is a paid mutator transaction binding the contract method 0x2b00aaef.
//
// Solidity: function submitBlockValidationResult(bool _result, address aggregateAddress, uint8 v, bytes32 r, bytes32 s, bytes32 _hash) returns()
func (_OracleContract *OracleContractTransactor) SubmitBlockValidationResult(opts *bind.TransactOpts, _result bool, aggregateAddress common.Address, v uint8, r [32]byte, s [32]byte, _hash [32]byte) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "submitBlockValidationResult", _result, aggregateAddress, v, r, s, _hash)
}

// SubmitBlockValidationResult is a paid mutator transaction binding the contract method 0x2b00aaef.
//
// Solidity: function submitBlockValidationResult(bool _result, address aggregateAddress, uint8 v, bytes32 r, bytes32 s, bytes32 _hash) returns()
func (_OracleContract *OracleContractSession) SubmitBlockValidationResult(_result bool, aggregateAddress common.Address, v uint8, r [32]byte, s [32]byte, _hash [32]byte) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitBlockValidationResult(&_OracleContract.TransactOpts, _result, aggregateAddress, v, r, s, _hash)
}

// SubmitBlockValidationResult is a paid mutator transaction binding the contract method 0x2b00aaef.
//
// Solidity: function submitBlockValidationResult(bool _result, address aggregateAddress, uint8 v, bytes32 r, bytes32 s, bytes32 _hash) returns()
func (_OracleContract *OracleContractTransactorSession) SubmitBlockValidationResult(_result bool, aggregateAddress common.Address, v uint8, r [32]byte, s [32]byte, _hash [32]byte) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitBlockValidationResult(&_OracleContract.TransactOpts, _result, aggregateAddress, v, r, s, _hash)
}

// SubmitTransactionValidationResult is a paid mutator transaction binding the contract method 0xc6d6e3d5.
//
// Solidity: function submitTransactionValidationResult(bool _result, address aggregateAddress, uint8 v, bytes32 r, bytes32 s, bytes32 _hash) returns()
func (_OracleContract *OracleContractTransactor) SubmitTransactionValidationResult(opts *bind.TransactOpts, _result bool, aggregateAddress common.Address, v uint8, r [32]byte, s [32]byte, _hash [32]byte) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "submitTransactionValidationResult", _result, aggregateAddress, v, r, s, _hash)
}

// SubmitTransactionValidationResult is a paid mutator transaction binding the contract method 0xc6d6e3d5.
//
// Solidity: function submitTransactionValidationResult(bool _result, address aggregateAddress, uint8 v, bytes32 r, bytes32 s, bytes32 _hash) returns()
func (_OracleContract *OracleContractSession) SubmitTransactionValidationResult(_result bool, aggregateAddress common.Address, v uint8, r [32]byte, s [32]byte, _hash [32]byte) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitTransactionValidationResult(&_OracleContract.TransactOpts, _result, aggregateAddress, v, r, s, _hash)
}

// SubmitTransactionValidationResult is a paid mutator transaction binding the contract method 0xc6d6e3d5.
//
// Solidity: function submitTransactionValidationResult(bool _result, address aggregateAddress, uint8 v, bytes32 r, bytes32 s, bytes32 _hash) returns()
func (_OracleContract *OracleContractTransactorSession) SubmitTransactionValidationResult(_result bool, aggregateAddress common.Address, v uint8, r [32]byte, s [32]byte, _hash [32]byte) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitTransactionValidationResult(&_OracleContract.TransactOpts, _result, aggregateAddress, v, r, s, _hash)
}

// ValidateBlock is a paid mutator transaction binding the contract method 0xdd5e22df.
//
// Solidity: function validateBlock(bytes32 _hash) payable returns()
func (_OracleContract *OracleContractTransactor) ValidateBlock(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "validateBlock", _hash)
}

// ValidateBlock is a paid mutator transaction binding the contract method 0xdd5e22df.
//
// Solidity: function validateBlock(bytes32 _hash) payable returns()
func (_OracleContract *OracleContractSession) ValidateBlock(_hash [32]byte) (*types.Transaction, error) {
	return _OracleContract.Contract.ValidateBlock(&_OracleContract.TransactOpts, _hash)
}

// ValidateBlock is a paid mutator transaction binding the contract method 0xdd5e22df.
//
// Solidity: function validateBlock(bytes32 _hash) payable returns()
func (_OracleContract *OracleContractTransactorSession) ValidateBlock(_hash [32]byte) (*types.Transaction, error) {
	return _OracleContract.Contract.ValidateBlock(&_OracleContract.TransactOpts, _hash)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0x344829c8.
//
// Solidity: function validateTransaction(bytes32 _hash) payable returns()
func (_OracleContract *OracleContractTransactor) ValidateTransaction(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "validateTransaction", _hash)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0x344829c8.
//
// Solidity: function validateTransaction(bytes32 _hash) payable returns()
func (_OracleContract *OracleContractSession) ValidateTransaction(_hash [32]byte) (*types.Transaction, error) {
	return _OracleContract.Contract.ValidateTransaction(&_OracleContract.TransactOpts, _hash)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0x344829c8.
//
// Solidity: function validateTransaction(bytes32 _hash) payable returns()
func (_OracleContract *OracleContractTransactorSession) ValidateTransaction(_hash [32]byte) (*types.Transaction, error) {
	return _OracleContract.Contract.ValidateTransaction(&_OracleContract.TransactOpts, _hash)
}

// OracleContractValidationBeginIterator is returned from FilterValidationBegin and is used to iterate over the raw logs and unpacked data for ValidationBegin events raised by the OracleContract contract.
type OracleContractValidationBeginIterator struct {
	Event *OracleContractValidationBegin // Event containing the contract specifics and raw log

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
func (it *OracleContractValidationBeginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleContractValidationBegin)
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
		it.Event = new(OracleContractValidationBegin)
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
func (it *OracleContractValidationBeginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleContractValidationBeginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleContractValidationBegin represents a ValidationBegin event raised by the OracleContract contract.
type OracleContractValidationBegin struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterValidationBegin is a free log retrieval operation binding the contract event 0x8e00f41384bb0046e201ec383e79140e9a0484a5db044984cbd072e4fbe6d7e2.
//
// Solidity: event ValidationBegin()
func (_OracleContract *OracleContractFilterer) FilterValidationBegin(opts *bind.FilterOpts) (*OracleContractValidationBeginIterator, error) {

	logs, sub, err := _OracleContract.contract.FilterLogs(opts, "ValidationBegin")
	if err != nil {
		return nil, err
	}
	return &OracleContractValidationBeginIterator{contract: _OracleContract.contract, event: "ValidationBegin", logs: logs, sub: sub}, nil
}

// WatchValidationBegin is a free log subscription operation binding the contract event 0x8e00f41384bb0046e201ec383e79140e9a0484a5db044984cbd072e4fbe6d7e2.
//
// Solidity: event ValidationBegin()
func (_OracleContract *OracleContractFilterer) WatchValidationBegin(opts *bind.WatchOpts, sink chan<- *OracleContractValidationBegin) (event.Subscription, error) {

	logs, sub, err := _OracleContract.contract.WatchLogs(opts, "ValidationBegin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleContractValidationBegin)
				if err := _OracleContract.contract.UnpackLog(event, "ValidationBegin", log); err != nil {
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

// ParseValidationBegin is a log parse operation binding the contract event 0x8e00f41384bb0046e201ec383e79140e9a0484a5db044984cbd072e4fbe6d7e2.
//
// Solidity: event ValidationBegin()
func (_OracleContract *OracleContractFilterer) ParseValidationBegin(log types.Log) (*OracleContractValidationBegin, error) {
	event := new(OracleContractValidationBegin)
	if err := _OracleContract.contract.UnpackLog(event, "ValidationBegin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleContractValidationRequestIterator is returned from FilterValidationRequest and is used to iterate over the raw logs and unpacked data for ValidationRequest events raised by the OracleContract contract.
type OracleContractValidationRequestIterator struct {
	Event *OracleContractValidationRequest // Event containing the contract specifics and raw log

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
func (it *OracleContractValidationRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleContractValidationRequest)
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
		it.Event = new(OracleContractValidationRequest)
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
func (it *OracleContractValidationRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleContractValidationRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleContractValidationRequest represents a ValidationRequest event raised by the OracleContract contract.
type OracleContractValidationRequest struct {
	Typ  uint8
	From common.Address
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterValidationRequest is a free log retrieval operation binding the contract event 0x3706933cbfd265e74e347f4c40263753cecc292080a1bfd0e9fd6ce994c08396.
//
// Solidity: event ValidationRequest(uint8 typ, address indexed from, bytes32 hash)
func (_OracleContract *OracleContractFilterer) FilterValidationRequest(opts *bind.FilterOpts, from []common.Address) (*OracleContractValidationRequestIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _OracleContract.contract.FilterLogs(opts, "ValidationRequest", fromRule)
	if err != nil {
		return nil, err
	}
	return &OracleContractValidationRequestIterator{contract: _OracleContract.contract, event: "ValidationRequest", logs: logs, sub: sub}, nil
}

// WatchValidationRequest is a free log subscription operation binding the contract event 0x3706933cbfd265e74e347f4c40263753cecc292080a1bfd0e9fd6ce994c08396.
//
// Solidity: event ValidationRequest(uint8 typ, address indexed from, bytes32 hash)
func (_OracleContract *OracleContractFilterer) WatchValidationRequest(opts *bind.WatchOpts, sink chan<- *OracleContractValidationRequest, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _OracleContract.contract.WatchLogs(opts, "ValidationRequest", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleContractValidationRequest)
				if err := _OracleContract.contract.UnpackLog(event, "ValidationRequest", log); err != nil {
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

// ParseValidationRequest is a log parse operation binding the contract event 0x3706933cbfd265e74e347f4c40263753cecc292080a1bfd0e9fd6ce994c08396.
//
// Solidity: event ValidationRequest(uint8 typ, address indexed from, bytes32 hash)
func (_OracleContract *OracleContractFilterer) ParseValidationRequest(log types.Log) (*OracleContractValidationRequest, error) {
	event := new(OracleContractValidationRequest)
	if err := _OracleContract.contract.UnpackLog(event, "ValidationRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleContractValidationResponseIterator is returned from FilterValidationResponse and is used to iterate over the raw logs and unpacked data for ValidationResponse events raised by the OracleContract contract.
type OracleContractValidationResponseIterator struct {
	Event *OracleContractValidationResponse // Event containing the contract specifics and raw log

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
func (it *OracleContractValidationResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleContractValidationResponse)
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
		it.Event = new(OracleContractValidationResponse)
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
func (it *OracleContractValidationResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleContractValidationResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleContractValidationResponse represents a ValidationResponse event raised by the OracleContract contract.
type OracleContractValidationResponse struct {
	Typ        uint8
	Aggregator common.Address
	Hash       [32]byte
	Valid      bool
	Fee        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidationResponse is a free log retrieval operation binding the contract event 0x9739d27192db56a131c86c297677afccac89aa4f39cbc5bcd62a8d10ce559675.
//
// Solidity: event ValidationResponse(uint8 typ, address indexed aggregator, bytes32 hash, bool valid, uint256 fee)
func (_OracleContract *OracleContractFilterer) FilterValidationResponse(opts *bind.FilterOpts, aggregator []common.Address) (*OracleContractValidationResponseIterator, error) {

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _OracleContract.contract.FilterLogs(opts, "ValidationResponse", aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &OracleContractValidationResponseIterator{contract: _OracleContract.contract, event: "ValidationResponse", logs: logs, sub: sub}, nil
}

// WatchValidationResponse is a free log subscription operation binding the contract event 0x9739d27192db56a131c86c297677afccac89aa4f39cbc5bcd62a8d10ce559675.
//
// Solidity: event ValidationResponse(uint8 typ, address indexed aggregator, bytes32 hash, bool valid, uint256 fee)
func (_OracleContract *OracleContractFilterer) WatchValidationResponse(opts *bind.WatchOpts, sink chan<- *OracleContractValidationResponse, aggregator []common.Address) (event.Subscription, error) {

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _OracleContract.contract.WatchLogs(opts, "ValidationResponse", aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleContractValidationResponse)
				if err := _OracleContract.contract.UnpackLog(event, "ValidationResponse", log); err != nil {
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

// ParseValidationResponse is a log parse operation binding the contract event 0x9739d27192db56a131c86c297677afccac89aa4f39cbc5bcd62a8d10ce559675.
//
// Solidity: event ValidationResponse(uint8 typ, address indexed aggregator, bytes32 hash, bool valid, uint256 fee)
func (_OracleContract *OracleContractFilterer) ParseValidationResponse(log types.Log) (*OracleContractValidationResponse, error) {
	event := new(OracleContractValidationResponse)
	if err := _OracleContract.contract.UnpackLog(event, "ValidationResponse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}