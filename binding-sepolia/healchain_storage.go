// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindingsepolia

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

// HealChainStorageMetaData contains all meta data concerning the HealChainStorage contract.
var HealChainStorageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"initialOracles\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addOracle\",\"inputs\":[{\"name\":\"oracle\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approvedOracles\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fulfillStore\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"encoded\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"recordId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getEncoded\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMetadata\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"dataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"originalSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"encodedSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dataShards\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"parityShards\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"recordOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"label\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"fulfilled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOracles\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOracle\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPending\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"remove\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeOracle\",\"inputs\":[{\"name\":\"oracle\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"store\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"dataShards\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"parityShards\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"label\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"store\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"label\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalRecords\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"EncodeRequested\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"requester\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"dataShards\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"parityShards\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"label\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OracleAdded\",\"inputs\":[{\"name\":\"oracle\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OracleRemoved\",\"inputs\":[{\"name\":\"oracle\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RecordDeleted\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Stored\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"dataHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"originalSize\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"encodedSize\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"label\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"EmptyData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotContractOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOracle\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOwner\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OracleAlreadyApproved\",\"inputs\":[{\"name\":\"oracle\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OracleNotApproved\",\"inputs\":[{\"name\":\"oracle\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"RecordNotFound\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"RequestNotFound\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
}

// HealChainStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use HealChainStorageMetaData.ABI instead.
var HealChainStorageABI = HealChainStorageMetaData.ABI

// HealChainStorage is an auto generated Go binding around an Ethereum contract.
type HealChainStorage struct {
	HealChainStorageCaller     // Read-only binding to the contract
	HealChainStorageTransactor // Write-only binding to the contract
	HealChainStorageFilterer   // Log filterer for contract events
}

// HealChainStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type HealChainStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HealChainStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HealChainStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HealChainStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HealChainStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HealChainStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HealChainStorageSession struct {
	Contract     *HealChainStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HealChainStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HealChainStorageCallerSession struct {
	Contract *HealChainStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// HealChainStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HealChainStorageTransactorSession struct {
	Contract     *HealChainStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// HealChainStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type HealChainStorageRaw struct {
	Contract *HealChainStorage // Generic contract binding to access the raw methods on
}

// HealChainStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HealChainStorageCallerRaw struct {
	Contract *HealChainStorageCaller // Generic read-only contract binding to access the raw methods on
}

// HealChainStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HealChainStorageTransactorRaw struct {
	Contract *HealChainStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHealChainStorage creates a new instance of HealChainStorage, bound to a specific deployed contract.
func NewHealChainStorage(address common.Address, backend bind.ContractBackend) (*HealChainStorage, error) {
	contract, err := bindHealChainStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HealChainStorage{HealChainStorageCaller: HealChainStorageCaller{contract: contract}, HealChainStorageTransactor: HealChainStorageTransactor{contract: contract}, HealChainStorageFilterer: HealChainStorageFilterer{contract: contract}}, nil
}

// NewHealChainStorageCaller creates a new read-only instance of HealChainStorage, bound to a specific deployed contract.
func NewHealChainStorageCaller(address common.Address, caller bind.ContractCaller) (*HealChainStorageCaller, error) {
	contract, err := bindHealChainStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HealChainStorageCaller{contract: contract}, nil
}

// NewHealChainStorageTransactor creates a new write-only instance of HealChainStorage, bound to a specific deployed contract.
func NewHealChainStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*HealChainStorageTransactor, error) {
	contract, err := bindHealChainStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HealChainStorageTransactor{contract: contract}, nil
}

// NewHealChainStorageFilterer creates a new log filterer instance of HealChainStorage, bound to a specific deployed contract.
func NewHealChainStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*HealChainStorageFilterer, error) {
	contract, err := bindHealChainStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HealChainStorageFilterer{contract: contract}, nil
}

// bindHealChainStorage binds a generic wrapper to an already deployed contract.
func bindHealChainStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HealChainStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HealChainStorage *HealChainStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HealChainStorage.Contract.HealChainStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HealChainStorage *HealChainStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HealChainStorage.Contract.HealChainStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HealChainStorage *HealChainStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HealChainStorage.Contract.HealChainStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HealChainStorage *HealChainStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HealChainStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HealChainStorage *HealChainStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HealChainStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HealChainStorage *HealChainStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HealChainStorage.Contract.contract.Transact(opts, method, params...)
}

// ApprovedOracles is a free data retrieval call binding the contract method 0x2a03fdc9.
//
// Solidity: function approvedOracles(address ) view returns(bool)
func (_HealChainStorage *HealChainStorageCaller) ApprovedOracles(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _HealChainStorage.contract.Call(opts, &out, "approvedOracles", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovedOracles is a free data retrieval call binding the contract method 0x2a03fdc9.
//
// Solidity: function approvedOracles(address ) view returns(bool)
func (_HealChainStorage *HealChainStorageSession) ApprovedOracles(arg0 common.Address) (bool, error) {
	return _HealChainStorage.Contract.ApprovedOracles(&_HealChainStorage.CallOpts, arg0)
}

// ApprovedOracles is a free data retrieval call binding the contract method 0x2a03fdc9.
//
// Solidity: function approvedOracles(address ) view returns(bool)
func (_HealChainStorage *HealChainStorageCallerSession) ApprovedOracles(arg0 common.Address) (bool, error) {
	return _HealChainStorage.Contract.ApprovedOracles(&_HealChainStorage.CallOpts, arg0)
}

// GetEncoded is a free data retrieval call binding the contract method 0xfe2feafe.
//
// Solidity: function getEncoded(uint256 id) view returns(bytes)
func (_HealChainStorage *HealChainStorageCaller) GetEncoded(opts *bind.CallOpts, id *big.Int) ([]byte, error) {
	var out []interface{}
	err := _HealChainStorage.contract.Call(opts, &out, "getEncoded", id)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetEncoded is a free data retrieval call binding the contract method 0xfe2feafe.
//
// Solidity: function getEncoded(uint256 id) view returns(bytes)
func (_HealChainStorage *HealChainStorageSession) GetEncoded(id *big.Int) ([]byte, error) {
	return _HealChainStorage.Contract.GetEncoded(&_HealChainStorage.CallOpts, id)
}

// GetEncoded is a free data retrieval call binding the contract method 0xfe2feafe.
//
// Solidity: function getEncoded(uint256 id) view returns(bytes)
func (_HealChainStorage *HealChainStorageCallerSession) GetEncoded(id *big.Int) ([]byte, error) {
	return _HealChainStorage.Contract.GetEncoded(&_HealChainStorage.CallOpts, id)
}

// GetMetadata is a free data retrieval call binding the contract method 0xa574cea4.
//
// Solidity: function getMetadata(uint256 id) view returns(bytes32 dataHash, uint256 originalSize, uint256 encodedSize, uint8 dataShards, uint8 parityShards, address recordOwner, uint256 timestamp, string label, bool fulfilled)
func (_HealChainStorage *HealChainStorageCaller) GetMetadata(opts *bind.CallOpts, id *big.Int) (struct {
	DataHash     [32]byte
	OriginalSize *big.Int
	EncodedSize  *big.Int
	DataShards   uint8
	ParityShards uint8
	RecordOwner  common.Address
	Timestamp    *big.Int
	Label        string
	Fulfilled    bool
}, error) {
	var out []interface{}
	err := _HealChainStorage.contract.Call(opts, &out, "getMetadata", id)

	outstruct := new(struct {
		DataHash     [32]byte
		OriginalSize *big.Int
		EncodedSize  *big.Int
		DataShards   uint8
		ParityShards uint8
		RecordOwner  common.Address
		Timestamp    *big.Int
		Label        string
		Fulfilled    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DataHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.OriginalSize = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EncodedSize = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DataShards = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.ParityShards = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.RecordOwner = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.Timestamp = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Label = *abi.ConvertType(out[7], new(string)).(*string)
	outstruct.Fulfilled = *abi.ConvertType(out[8], new(bool)).(*bool)

	return *outstruct, err

}

// GetMetadata is a free data retrieval call binding the contract method 0xa574cea4.
//
// Solidity: function getMetadata(uint256 id) view returns(bytes32 dataHash, uint256 originalSize, uint256 encodedSize, uint8 dataShards, uint8 parityShards, address recordOwner, uint256 timestamp, string label, bool fulfilled)
func (_HealChainStorage *HealChainStorageSession) GetMetadata(id *big.Int) (struct {
	DataHash     [32]byte
	OriginalSize *big.Int
	EncodedSize  *big.Int
	DataShards   uint8
	ParityShards uint8
	RecordOwner  common.Address
	Timestamp    *big.Int
	Label        string
	Fulfilled    bool
}, error) {
	return _HealChainStorage.Contract.GetMetadata(&_HealChainStorage.CallOpts, id)
}

// GetMetadata is a free data retrieval call binding the contract method 0xa574cea4.
//
// Solidity: function getMetadata(uint256 id) view returns(bytes32 dataHash, uint256 originalSize, uint256 encodedSize, uint8 dataShards, uint8 parityShards, address recordOwner, uint256 timestamp, string label, bool fulfilled)
func (_HealChainStorage *HealChainStorageCallerSession) GetMetadata(id *big.Int) (struct {
	DataHash     [32]byte
	OriginalSize *big.Int
	EncodedSize  *big.Int
	DataShards   uint8
	ParityShards uint8
	RecordOwner  common.Address
	Timestamp    *big.Int
	Label        string
	Fulfilled    bool
}, error) {
	return _HealChainStorage.Contract.GetMetadata(&_HealChainStorage.CallOpts, id)
}

// GetOracles is a free data retrieval call binding the contract method 0x40884c52.
//
// Solidity: function getOracles() view returns(address[])
func (_HealChainStorage *HealChainStorageCaller) GetOracles(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _HealChainStorage.contract.Call(opts, &out, "getOracles")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOracles is a free data retrieval call binding the contract method 0x40884c52.
//
// Solidity: function getOracles() view returns(address[])
func (_HealChainStorage *HealChainStorageSession) GetOracles() ([]common.Address, error) {
	return _HealChainStorage.Contract.GetOracles(&_HealChainStorage.CallOpts)
}

// GetOracles is a free data retrieval call binding the contract method 0x40884c52.
//
// Solidity: function getOracles() view returns(address[])
func (_HealChainStorage *HealChainStorageCallerSession) GetOracles() ([]common.Address, error) {
	return _HealChainStorage.Contract.GetOracles(&_HealChainStorage.CallOpts)
}

// IsOracle is a free data retrieval call binding the contract method 0xa97e5c93.
//
// Solidity: function isOracle(address addr) view returns(bool)
func (_HealChainStorage *HealChainStorageCaller) IsOracle(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _HealChainStorage.contract.Call(opts, &out, "isOracle", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOracle is a free data retrieval call binding the contract method 0xa97e5c93.
//
// Solidity: function isOracle(address addr) view returns(bool)
func (_HealChainStorage *HealChainStorageSession) IsOracle(addr common.Address) (bool, error) {
	return _HealChainStorage.Contract.IsOracle(&_HealChainStorage.CallOpts, addr)
}

// IsOracle is a free data retrieval call binding the contract method 0xa97e5c93.
//
// Solidity: function isOracle(address addr) view returns(bool)
func (_HealChainStorage *HealChainStorageCallerSession) IsOracle(addr common.Address) (bool, error) {
	return _HealChainStorage.Contract.IsOracle(&_HealChainStorage.CallOpts, addr)
}

// IsPending is a free data retrieval call binding the contract method 0xca8836d2.
//
// Solidity: function isPending(uint256 requestId) view returns(bool)
func (_HealChainStorage *HealChainStorageCaller) IsPending(opts *bind.CallOpts, requestId *big.Int) (bool, error) {
	var out []interface{}
	err := _HealChainStorage.contract.Call(opts, &out, "isPending", requestId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPending is a free data retrieval call binding the contract method 0xca8836d2.
//
// Solidity: function isPending(uint256 requestId) view returns(bool)
func (_HealChainStorage *HealChainStorageSession) IsPending(requestId *big.Int) (bool, error) {
	return _HealChainStorage.Contract.IsPending(&_HealChainStorage.CallOpts, requestId)
}

// IsPending is a free data retrieval call binding the contract method 0xca8836d2.
//
// Solidity: function isPending(uint256 requestId) view returns(bool)
func (_HealChainStorage *HealChainStorageCallerSession) IsPending(requestId *big.Int) (bool, error) {
	return _HealChainStorage.Contract.IsPending(&_HealChainStorage.CallOpts, requestId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HealChainStorage *HealChainStorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HealChainStorage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HealChainStorage *HealChainStorageSession) Owner() (common.Address, error) {
	return _HealChainStorage.Contract.Owner(&_HealChainStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HealChainStorage *HealChainStorageCallerSession) Owner() (common.Address, error) {
	return _HealChainStorage.Contract.Owner(&_HealChainStorage.CallOpts)
}

// TotalRecords is a free data retrieval call binding the contract method 0x125f8974.
//
// Solidity: function totalRecords() view returns(uint256)
func (_HealChainStorage *HealChainStorageCaller) TotalRecords(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HealChainStorage.contract.Call(opts, &out, "totalRecords")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRecords is a free data retrieval call binding the contract method 0x125f8974.
//
// Solidity: function totalRecords() view returns(uint256)
func (_HealChainStorage *HealChainStorageSession) TotalRecords() (*big.Int, error) {
	return _HealChainStorage.Contract.TotalRecords(&_HealChainStorage.CallOpts)
}

// TotalRecords is a free data retrieval call binding the contract method 0x125f8974.
//
// Solidity: function totalRecords() view returns(uint256)
func (_HealChainStorage *HealChainStorageCallerSession) TotalRecords() (*big.Int, error) {
	return _HealChainStorage.Contract.TotalRecords(&_HealChainStorage.CallOpts)
}

// AddOracle is a paid mutator transaction binding the contract method 0xdf5dd1a5.
//
// Solidity: function addOracle(address oracle) returns()
func (_HealChainStorage *HealChainStorageTransactor) AddOracle(opts *bind.TransactOpts, oracle common.Address) (*types.Transaction, error) {
	return _HealChainStorage.contract.Transact(opts, "addOracle", oracle)
}

// AddOracle is a paid mutator transaction binding the contract method 0xdf5dd1a5.
//
// Solidity: function addOracle(address oracle) returns()
func (_HealChainStorage *HealChainStorageSession) AddOracle(oracle common.Address) (*types.Transaction, error) {
	return _HealChainStorage.Contract.AddOracle(&_HealChainStorage.TransactOpts, oracle)
}

// AddOracle is a paid mutator transaction binding the contract method 0xdf5dd1a5.
//
// Solidity: function addOracle(address oracle) returns()
func (_HealChainStorage *HealChainStorageTransactorSession) AddOracle(oracle common.Address) (*types.Transaction, error) {
	return _HealChainStorage.Contract.AddOracle(&_HealChainStorage.TransactOpts, oracle)
}

// FulfillStore is a paid mutator transaction binding the contract method 0xbb37740e.
//
// Solidity: function fulfillStore(uint256 requestId, bytes encoded) returns(uint256 recordId)
func (_HealChainStorage *HealChainStorageTransactor) FulfillStore(opts *bind.TransactOpts, requestId *big.Int, encoded []byte) (*types.Transaction, error) {
	return _HealChainStorage.contract.Transact(opts, "fulfillStore", requestId, encoded)
}

// FulfillStore is a paid mutator transaction binding the contract method 0xbb37740e.
//
// Solidity: function fulfillStore(uint256 requestId, bytes encoded) returns(uint256 recordId)
func (_HealChainStorage *HealChainStorageSession) FulfillStore(requestId *big.Int, encoded []byte) (*types.Transaction, error) {
	return _HealChainStorage.Contract.FulfillStore(&_HealChainStorage.TransactOpts, requestId, encoded)
}

// FulfillStore is a paid mutator transaction binding the contract method 0xbb37740e.
//
// Solidity: function fulfillStore(uint256 requestId, bytes encoded) returns(uint256 recordId)
func (_HealChainStorage *HealChainStorageTransactorSession) FulfillStore(requestId *big.Int, encoded []byte) (*types.Transaction, error) {
	return _HealChainStorage.Contract.FulfillStore(&_HealChainStorage.TransactOpts, requestId, encoded)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 id) returns()
func (_HealChainStorage *HealChainStorageTransactor) Remove(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _HealChainStorage.contract.Transact(opts, "remove", id)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 id) returns()
func (_HealChainStorage *HealChainStorageSession) Remove(id *big.Int) (*types.Transaction, error) {
	return _HealChainStorage.Contract.Remove(&_HealChainStorage.TransactOpts, id)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 id) returns()
func (_HealChainStorage *HealChainStorageTransactorSession) Remove(id *big.Int) (*types.Transaction, error) {
	return _HealChainStorage.Contract.Remove(&_HealChainStorage.TransactOpts, id)
}

// RemoveOracle is a paid mutator transaction binding the contract method 0xfdc85fc4.
//
// Solidity: function removeOracle(address oracle) returns()
func (_HealChainStorage *HealChainStorageTransactor) RemoveOracle(opts *bind.TransactOpts, oracle common.Address) (*types.Transaction, error) {
	return _HealChainStorage.contract.Transact(opts, "removeOracle", oracle)
}

// RemoveOracle is a paid mutator transaction binding the contract method 0xfdc85fc4.
//
// Solidity: function removeOracle(address oracle) returns()
func (_HealChainStorage *HealChainStorageSession) RemoveOracle(oracle common.Address) (*types.Transaction, error) {
	return _HealChainStorage.Contract.RemoveOracle(&_HealChainStorage.TransactOpts, oracle)
}

// RemoveOracle is a paid mutator transaction binding the contract method 0xfdc85fc4.
//
// Solidity: function removeOracle(address oracle) returns()
func (_HealChainStorage *HealChainStorageTransactorSession) RemoveOracle(oracle common.Address) (*types.Transaction, error) {
	return _HealChainStorage.Contract.RemoveOracle(&_HealChainStorage.TransactOpts, oracle)
}

// Store is a paid mutator transaction binding the contract method 0x46bd680e.
//
// Solidity: function store(bytes data, uint8 dataShards, uint8 parityShards, string label) returns(uint256 requestId)
func (_HealChainStorage *HealChainStorageTransactor) Store(opts *bind.TransactOpts, data []byte, dataShards uint8, parityShards uint8, label string) (*types.Transaction, error) {
	return _HealChainStorage.contract.Transact(opts, "store", data, dataShards, parityShards, label)
}

// Store is a paid mutator transaction binding the contract method 0x46bd680e.
//
// Solidity: function store(bytes data, uint8 dataShards, uint8 parityShards, string label) returns(uint256 requestId)
func (_HealChainStorage *HealChainStorageSession) Store(data []byte, dataShards uint8, parityShards uint8, label string) (*types.Transaction, error) {
	return _HealChainStorage.Contract.Store(&_HealChainStorage.TransactOpts, data, dataShards, parityShards, label)
}

// Store is a paid mutator transaction binding the contract method 0x46bd680e.
//
// Solidity: function store(bytes data, uint8 dataShards, uint8 parityShards, string label) returns(uint256 requestId)
func (_HealChainStorage *HealChainStorageTransactorSession) Store(data []byte, dataShards uint8, parityShards uint8, label string) (*types.Transaction, error) {
	return _HealChainStorage.Contract.Store(&_HealChainStorage.TransactOpts, data, dataShards, parityShards, label)
}

// Store0 is a paid mutator transaction binding the contract method 0xee340c74.
//
// Solidity: function store(bytes data, string label) returns(uint256 requestId)
func (_HealChainStorage *HealChainStorageTransactor) Store0(opts *bind.TransactOpts, data []byte, label string) (*types.Transaction, error) {
	return _HealChainStorage.contract.Transact(opts, "store0", data, label)
}

// Store0 is a paid mutator transaction binding the contract method 0xee340c74.
//
// Solidity: function store(bytes data, string label) returns(uint256 requestId)
func (_HealChainStorage *HealChainStorageSession) Store0(data []byte, label string) (*types.Transaction, error) {
	return _HealChainStorage.Contract.Store0(&_HealChainStorage.TransactOpts, data, label)
}

// Store0 is a paid mutator transaction binding the contract method 0xee340c74.
//
// Solidity: function store(bytes data, string label) returns(uint256 requestId)
func (_HealChainStorage *HealChainStorageTransactorSession) Store0(data []byte, label string) (*types.Transaction, error) {
	return _HealChainStorage.Contract.Store0(&_HealChainStorage.TransactOpts, data, label)
}

// HealChainStorageEncodeRequestedIterator is returned from FilterEncodeRequested and is used to iterate over the raw logs and unpacked data for EncodeRequested events raised by the HealChainStorage contract.
type HealChainStorageEncodeRequestedIterator struct {
	Event *HealChainStorageEncodeRequested // Event containing the contract specifics and raw log

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
func (it *HealChainStorageEncodeRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HealChainStorageEncodeRequested)
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
		it.Event = new(HealChainStorageEncodeRequested)
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
func (it *HealChainStorageEncodeRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HealChainStorageEncodeRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HealChainStorageEncodeRequested represents a EncodeRequested event raised by the HealChainStorage contract.
type HealChainStorageEncodeRequested struct {
	RequestId    *big.Int
	Requester    common.Address
	Data         []byte
	DataShards   uint8
	ParityShards uint8
	Label        string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEncodeRequested is a free log retrieval operation binding the contract event 0x49210a6712191520fc283f4e7f9206ec2b512ba2424570cf00e1c917cb097efe.
//
// Solidity: event EncodeRequested(uint256 indexed requestId, address indexed requester, bytes data, uint8 dataShards, uint8 parityShards, string label)
func (_HealChainStorage *HealChainStorageFilterer) FilterEncodeRequested(opts *bind.FilterOpts, requestId []*big.Int, requester []common.Address) (*HealChainStorageEncodeRequestedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _HealChainStorage.contract.FilterLogs(opts, "EncodeRequested", requestIdRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &HealChainStorageEncodeRequestedIterator{contract: _HealChainStorage.contract, event: "EncodeRequested", logs: logs, sub: sub}, nil
}

// WatchEncodeRequested is a free log subscription operation binding the contract event 0x49210a6712191520fc283f4e7f9206ec2b512ba2424570cf00e1c917cb097efe.
//
// Solidity: event EncodeRequested(uint256 indexed requestId, address indexed requester, bytes data, uint8 dataShards, uint8 parityShards, string label)
func (_HealChainStorage *HealChainStorageFilterer) WatchEncodeRequested(opts *bind.WatchOpts, sink chan<- *HealChainStorageEncodeRequested, requestId []*big.Int, requester []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _HealChainStorage.contract.WatchLogs(opts, "EncodeRequested", requestIdRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HealChainStorageEncodeRequested)
				if err := _HealChainStorage.contract.UnpackLog(event, "EncodeRequested", log); err != nil {
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

// ParseEncodeRequested is a log parse operation binding the contract event 0x49210a6712191520fc283f4e7f9206ec2b512ba2424570cf00e1c917cb097efe.
//
// Solidity: event EncodeRequested(uint256 indexed requestId, address indexed requester, bytes data, uint8 dataShards, uint8 parityShards, string label)
func (_HealChainStorage *HealChainStorageFilterer) ParseEncodeRequested(log types.Log) (*HealChainStorageEncodeRequested, error) {
	event := new(HealChainStorageEncodeRequested)
	if err := _HealChainStorage.contract.UnpackLog(event, "EncodeRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HealChainStorageOracleAddedIterator is returned from FilterOracleAdded and is used to iterate over the raw logs and unpacked data for OracleAdded events raised by the HealChainStorage contract.
type HealChainStorageOracleAddedIterator struct {
	Event *HealChainStorageOracleAdded // Event containing the contract specifics and raw log

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
func (it *HealChainStorageOracleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HealChainStorageOracleAdded)
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
		it.Event = new(HealChainStorageOracleAdded)
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
func (it *HealChainStorageOracleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HealChainStorageOracleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HealChainStorageOracleAdded represents a OracleAdded event raised by the HealChainStorage contract.
type HealChainStorageOracleAdded struct {
	Oracle common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOracleAdded is a free log retrieval operation binding the contract event 0x0047706786c922d17b39285dc59d696bafea72c0b003d3841ae1202076f4c2e4.
//
// Solidity: event OracleAdded(address indexed oracle)
func (_HealChainStorage *HealChainStorageFilterer) FilterOracleAdded(opts *bind.FilterOpts, oracle []common.Address) (*HealChainStorageOracleAddedIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _HealChainStorage.contract.FilterLogs(opts, "OracleAdded", oracleRule)
	if err != nil {
		return nil, err
	}
	return &HealChainStorageOracleAddedIterator{contract: _HealChainStorage.contract, event: "OracleAdded", logs: logs, sub: sub}, nil
}

// WatchOracleAdded is a free log subscription operation binding the contract event 0x0047706786c922d17b39285dc59d696bafea72c0b003d3841ae1202076f4c2e4.
//
// Solidity: event OracleAdded(address indexed oracle)
func (_HealChainStorage *HealChainStorageFilterer) WatchOracleAdded(opts *bind.WatchOpts, sink chan<- *HealChainStorageOracleAdded, oracle []common.Address) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _HealChainStorage.contract.WatchLogs(opts, "OracleAdded", oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HealChainStorageOracleAdded)
				if err := _HealChainStorage.contract.UnpackLog(event, "OracleAdded", log); err != nil {
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

// ParseOracleAdded is a log parse operation binding the contract event 0x0047706786c922d17b39285dc59d696bafea72c0b003d3841ae1202076f4c2e4.
//
// Solidity: event OracleAdded(address indexed oracle)
func (_HealChainStorage *HealChainStorageFilterer) ParseOracleAdded(log types.Log) (*HealChainStorageOracleAdded, error) {
	event := new(HealChainStorageOracleAdded)
	if err := _HealChainStorage.contract.UnpackLog(event, "OracleAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HealChainStorageOracleRemovedIterator is returned from FilterOracleRemoved and is used to iterate over the raw logs and unpacked data for OracleRemoved events raised by the HealChainStorage contract.
type HealChainStorageOracleRemovedIterator struct {
	Event *HealChainStorageOracleRemoved // Event containing the contract specifics and raw log

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
func (it *HealChainStorageOracleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HealChainStorageOracleRemoved)
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
		it.Event = new(HealChainStorageOracleRemoved)
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
func (it *HealChainStorageOracleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HealChainStorageOracleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HealChainStorageOracleRemoved represents a OracleRemoved event raised by the HealChainStorage contract.
type HealChainStorageOracleRemoved struct {
	Oracle common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOracleRemoved is a free log retrieval operation binding the contract event 0x9c8e7d83025bef8a04c664b2f753f64b8814bdb7e27291d7e50935f18cc3c712.
//
// Solidity: event OracleRemoved(address indexed oracle)
func (_HealChainStorage *HealChainStorageFilterer) FilterOracleRemoved(opts *bind.FilterOpts, oracle []common.Address) (*HealChainStorageOracleRemovedIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _HealChainStorage.contract.FilterLogs(opts, "OracleRemoved", oracleRule)
	if err != nil {
		return nil, err
	}
	return &HealChainStorageOracleRemovedIterator{contract: _HealChainStorage.contract, event: "OracleRemoved", logs: logs, sub: sub}, nil
}

// WatchOracleRemoved is a free log subscription operation binding the contract event 0x9c8e7d83025bef8a04c664b2f753f64b8814bdb7e27291d7e50935f18cc3c712.
//
// Solidity: event OracleRemoved(address indexed oracle)
func (_HealChainStorage *HealChainStorageFilterer) WatchOracleRemoved(opts *bind.WatchOpts, sink chan<- *HealChainStorageOracleRemoved, oracle []common.Address) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _HealChainStorage.contract.WatchLogs(opts, "OracleRemoved", oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HealChainStorageOracleRemoved)
				if err := _HealChainStorage.contract.UnpackLog(event, "OracleRemoved", log); err != nil {
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

// ParseOracleRemoved is a log parse operation binding the contract event 0x9c8e7d83025bef8a04c664b2f753f64b8814bdb7e27291d7e50935f18cc3c712.
//
// Solidity: event OracleRemoved(address indexed oracle)
func (_HealChainStorage *HealChainStorageFilterer) ParseOracleRemoved(log types.Log) (*HealChainStorageOracleRemoved, error) {
	event := new(HealChainStorageOracleRemoved)
	if err := _HealChainStorage.contract.UnpackLog(event, "OracleRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HealChainStorageRecordDeletedIterator is returned from FilterRecordDeleted and is used to iterate over the raw logs and unpacked data for RecordDeleted events raised by the HealChainStorage contract.
type HealChainStorageRecordDeletedIterator struct {
	Event *HealChainStorageRecordDeleted // Event containing the contract specifics and raw log

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
func (it *HealChainStorageRecordDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HealChainStorageRecordDeleted)
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
		it.Event = new(HealChainStorageRecordDeleted)
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
func (it *HealChainStorageRecordDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HealChainStorageRecordDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HealChainStorageRecordDeleted represents a RecordDeleted event raised by the HealChainStorage contract.
type HealChainStorageRecordDeleted struct {
	Id    *big.Int
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRecordDeleted is a free log retrieval operation binding the contract event 0x7d14a037f6460104a7e358ca6256314e3b495e7d360a781b444e183be65c18b2.
//
// Solidity: event RecordDeleted(uint256 indexed id, address indexed owner)
func (_HealChainStorage *HealChainStorageFilterer) FilterRecordDeleted(opts *bind.FilterOpts, id []*big.Int, owner []common.Address) (*HealChainStorageRecordDeletedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _HealChainStorage.contract.FilterLogs(opts, "RecordDeleted", idRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &HealChainStorageRecordDeletedIterator{contract: _HealChainStorage.contract, event: "RecordDeleted", logs: logs, sub: sub}, nil
}

// WatchRecordDeleted is a free log subscription operation binding the contract event 0x7d14a037f6460104a7e358ca6256314e3b495e7d360a781b444e183be65c18b2.
//
// Solidity: event RecordDeleted(uint256 indexed id, address indexed owner)
func (_HealChainStorage *HealChainStorageFilterer) WatchRecordDeleted(opts *bind.WatchOpts, sink chan<- *HealChainStorageRecordDeleted, id []*big.Int, owner []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _HealChainStorage.contract.WatchLogs(opts, "RecordDeleted", idRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HealChainStorageRecordDeleted)
				if err := _HealChainStorage.contract.UnpackLog(event, "RecordDeleted", log); err != nil {
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

// ParseRecordDeleted is a log parse operation binding the contract event 0x7d14a037f6460104a7e358ca6256314e3b495e7d360a781b444e183be65c18b2.
//
// Solidity: event RecordDeleted(uint256 indexed id, address indexed owner)
func (_HealChainStorage *HealChainStorageFilterer) ParseRecordDeleted(log types.Log) (*HealChainStorageRecordDeleted, error) {
	event := new(HealChainStorageRecordDeleted)
	if err := _HealChainStorage.contract.UnpackLog(event, "RecordDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HealChainStorageStoredIterator is returned from FilterStored and is used to iterate over the raw logs and unpacked data for Stored events raised by the HealChainStorage contract.
type HealChainStorageStoredIterator struct {
	Event *HealChainStorageStored // Event containing the contract specifics and raw log

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
func (it *HealChainStorageStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HealChainStorageStored)
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
		it.Event = new(HealChainStorageStored)
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
func (it *HealChainStorageStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HealChainStorageStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HealChainStorageStored represents a Stored event raised by the HealChainStorage contract.
type HealChainStorageStored struct {
	Id           *big.Int
	Owner        common.Address
	DataHash     [32]byte
	OriginalSize *big.Int
	EncodedSize  *big.Int
	Label        string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStored is a free log retrieval operation binding the contract event 0x47f5d3928a4ed0c5ee46feb9ae2175dbbfff3bbabceb35485848bc63c636cee7.
//
// Solidity: event Stored(uint256 indexed id, address indexed owner, bytes32 dataHash, uint256 originalSize, uint256 encodedSize, string label)
func (_HealChainStorage *HealChainStorageFilterer) FilterStored(opts *bind.FilterOpts, id []*big.Int, owner []common.Address) (*HealChainStorageStoredIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _HealChainStorage.contract.FilterLogs(opts, "Stored", idRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &HealChainStorageStoredIterator{contract: _HealChainStorage.contract, event: "Stored", logs: logs, sub: sub}, nil
}

// WatchStored is a free log subscription operation binding the contract event 0x47f5d3928a4ed0c5ee46feb9ae2175dbbfff3bbabceb35485848bc63c636cee7.
//
// Solidity: event Stored(uint256 indexed id, address indexed owner, bytes32 dataHash, uint256 originalSize, uint256 encodedSize, string label)
func (_HealChainStorage *HealChainStorageFilterer) WatchStored(opts *bind.WatchOpts, sink chan<- *HealChainStorageStored, id []*big.Int, owner []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _HealChainStorage.contract.WatchLogs(opts, "Stored", idRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HealChainStorageStored)
				if err := _HealChainStorage.contract.UnpackLog(event, "Stored", log); err != nil {
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

// ParseStored is a log parse operation binding the contract event 0x47f5d3928a4ed0c5ee46feb9ae2175dbbfff3bbabceb35485848bc63c636cee7.
//
// Solidity: event Stored(uint256 indexed id, address indexed owner, bytes32 dataHash, uint256 originalSize, uint256 encodedSize, string label)
func (_HealChainStorage *HealChainStorageFilterer) ParseStored(log types.Log) (*HealChainStorageStored, error) {
	event := new(HealChainStorageStored)
	if err := _HealChainStorage.contract.UnpackLog(event, "Stored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
