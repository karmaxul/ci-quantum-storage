// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package binding

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
	ABI: "[{\"type\":\"function\",\"name\":\"getEncoded\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMetadata\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"dataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"originalSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"encodedSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dataShards\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"parityShards\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"label\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"remove\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"retrieve\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"retrieveVerified\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"store\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"dataShards\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"parityShards\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"label\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"store\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"label\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalRecords\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"RecordDeleted\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Retrieved\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"verified\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Stored\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"dataHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"originalSize\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"encodedSize\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"label\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"DecodeCallFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyInput\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EncodeCallFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOwner\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"RecordNotFound\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"VerificationFailed\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"got\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
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
// Solidity: function getMetadata(uint256 id) view returns(bytes32 dataHash, uint256 originalSize, uint256 encodedSize, uint8 dataShards, uint8 parityShards, address owner, uint256 timestamp, string label)
func (_HealChainStorage *HealChainStorageCaller) GetMetadata(opts *bind.CallOpts, id *big.Int) (struct {
	DataHash     [32]byte
	OriginalSize *big.Int
	EncodedSize  *big.Int
	DataShards   uint8
	ParityShards uint8
	Owner        common.Address
	Timestamp    *big.Int
	Label        string
}, error) {
	var out []interface{}
	err := _HealChainStorage.contract.Call(opts, &out, "getMetadata", id)

	outstruct := new(struct {
		DataHash     [32]byte
		OriginalSize *big.Int
		EncodedSize  *big.Int
		DataShards   uint8
		ParityShards uint8
		Owner        common.Address
		Timestamp    *big.Int
		Label        string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DataHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.OriginalSize = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EncodedSize = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DataShards = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.ParityShards = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.Owner = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.Timestamp = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Label = *abi.ConvertType(out[7], new(string)).(*string)

	return *outstruct, err

}

// GetMetadata is a free data retrieval call binding the contract method 0xa574cea4.
//
// Solidity: function getMetadata(uint256 id) view returns(bytes32 dataHash, uint256 originalSize, uint256 encodedSize, uint8 dataShards, uint8 parityShards, address owner, uint256 timestamp, string label)
func (_HealChainStorage *HealChainStorageSession) GetMetadata(id *big.Int) (struct {
	DataHash     [32]byte
	OriginalSize *big.Int
	EncodedSize  *big.Int
	DataShards   uint8
	ParityShards uint8
	Owner        common.Address
	Timestamp    *big.Int
	Label        string
}, error) {
	return _HealChainStorage.Contract.GetMetadata(&_HealChainStorage.CallOpts, id)
}

// GetMetadata is a free data retrieval call binding the contract method 0xa574cea4.
//
// Solidity: function getMetadata(uint256 id) view returns(bytes32 dataHash, uint256 originalSize, uint256 encodedSize, uint8 dataShards, uint8 parityShards, address owner, uint256 timestamp, string label)
func (_HealChainStorage *HealChainStorageCallerSession) GetMetadata(id *big.Int) (struct {
	DataHash     [32]byte
	OriginalSize *big.Int
	EncodedSize  *big.Int
	DataShards   uint8
	ParityShards uint8
	Owner        common.Address
	Timestamp    *big.Int
	Label        string
}, error) {
	return _HealChainStorage.Contract.GetMetadata(&_HealChainStorage.CallOpts, id)
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

// Retrieve is a paid mutator transaction binding the contract method 0x8f88708b.
//
// Solidity: function retrieve(uint256 id) returns(bytes)
func (_HealChainStorage *HealChainStorageTransactor) Retrieve(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _HealChainStorage.contract.Transact(opts, "retrieve", id)
}

// Retrieve is a paid mutator transaction binding the contract method 0x8f88708b.
//
// Solidity: function retrieve(uint256 id) returns(bytes)
func (_HealChainStorage *HealChainStorageSession) Retrieve(id *big.Int) (*types.Transaction, error) {
	return _HealChainStorage.Contract.Retrieve(&_HealChainStorage.TransactOpts, id)
}

// Retrieve is a paid mutator transaction binding the contract method 0x8f88708b.
//
// Solidity: function retrieve(uint256 id) returns(bytes)
func (_HealChainStorage *HealChainStorageTransactorSession) Retrieve(id *big.Int) (*types.Transaction, error) {
	return _HealChainStorage.Contract.Retrieve(&_HealChainStorage.TransactOpts, id)
}

// RetrieveVerified is a paid mutator transaction binding the contract method 0x12a92c17.
//
// Solidity: function retrieveVerified(uint256 id) returns(bytes)
func (_HealChainStorage *HealChainStorageTransactor) RetrieveVerified(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _HealChainStorage.contract.Transact(opts, "retrieveVerified", id)
}

// RetrieveVerified is a paid mutator transaction binding the contract method 0x12a92c17.
//
// Solidity: function retrieveVerified(uint256 id) returns(bytes)
func (_HealChainStorage *HealChainStorageSession) RetrieveVerified(id *big.Int) (*types.Transaction, error) {
	return _HealChainStorage.Contract.RetrieveVerified(&_HealChainStorage.TransactOpts, id)
}

// RetrieveVerified is a paid mutator transaction binding the contract method 0x12a92c17.
//
// Solidity: function retrieveVerified(uint256 id) returns(bytes)
func (_HealChainStorage *HealChainStorageTransactorSession) RetrieveVerified(id *big.Int) (*types.Transaction, error) {
	return _HealChainStorage.Contract.RetrieveVerified(&_HealChainStorage.TransactOpts, id)
}

// Store is a paid mutator transaction binding the contract method 0x46bd680e.
//
// Solidity: function store(bytes data, uint8 dataShards, uint8 parityShards, string label) returns(uint256 id)
func (_HealChainStorage *HealChainStorageTransactor) Store(opts *bind.TransactOpts, data []byte, dataShards uint8, parityShards uint8, label string) (*types.Transaction, error) {
	return _HealChainStorage.contract.Transact(opts, "store", data, dataShards, parityShards, label)
}

// Store is a paid mutator transaction binding the contract method 0x46bd680e.
//
// Solidity: function store(bytes data, uint8 dataShards, uint8 parityShards, string label) returns(uint256 id)
func (_HealChainStorage *HealChainStorageSession) Store(data []byte, dataShards uint8, parityShards uint8, label string) (*types.Transaction, error) {
	return _HealChainStorage.Contract.Store(&_HealChainStorage.TransactOpts, data, dataShards, parityShards, label)
}

// Store is a paid mutator transaction binding the contract method 0x46bd680e.
//
// Solidity: function store(bytes data, uint8 dataShards, uint8 parityShards, string label) returns(uint256 id)
func (_HealChainStorage *HealChainStorageTransactorSession) Store(data []byte, dataShards uint8, parityShards uint8, label string) (*types.Transaction, error) {
	return _HealChainStorage.Contract.Store(&_HealChainStorage.TransactOpts, data, dataShards, parityShards, label)
}

// Store0 is a paid mutator transaction binding the contract method 0xee340c74.
//
// Solidity: function store(bytes data, string label) returns(uint256 id)
func (_HealChainStorage *HealChainStorageTransactor) Store0(opts *bind.TransactOpts, data []byte, label string) (*types.Transaction, error) {
	return _HealChainStorage.contract.Transact(opts, "store0", data, label)
}

// Store0 is a paid mutator transaction binding the contract method 0xee340c74.
//
// Solidity: function store(bytes data, string label) returns(uint256 id)
func (_HealChainStorage *HealChainStorageSession) Store0(data []byte, label string) (*types.Transaction, error) {
	return _HealChainStorage.Contract.Store0(&_HealChainStorage.TransactOpts, data, label)
}

// Store0 is a paid mutator transaction binding the contract method 0xee340c74.
//
// Solidity: function store(bytes data, string label) returns(uint256 id)
func (_HealChainStorage *HealChainStorageTransactorSession) Store0(data []byte, label string) (*types.Transaction, error) {
	return _HealChainStorage.Contract.Store0(&_HealChainStorage.TransactOpts, data, label)
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

// HealChainStorageRetrievedIterator is returned from FilterRetrieved and is used to iterate over the raw logs and unpacked data for Retrieved events raised by the HealChainStorage contract.
type HealChainStorageRetrievedIterator struct {
	Event *HealChainStorageRetrieved // Event containing the contract specifics and raw log

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
func (it *HealChainStorageRetrievedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HealChainStorageRetrieved)
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
		it.Event = new(HealChainStorageRetrieved)
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
func (it *HealChainStorageRetrievedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HealChainStorageRetrievedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HealChainStorageRetrieved represents a Retrieved event raised by the HealChainStorage contract.
type HealChainStorageRetrieved struct {
	Id       *big.Int
	Caller   common.Address
	Verified bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRetrieved is a free log retrieval operation binding the contract event 0xe3086fe6321b51142821be491f79abf73cd9046082dd1d6b6bd7ee7c732fc453.
//
// Solidity: event Retrieved(uint256 indexed id, address indexed caller, bool verified)
func (_HealChainStorage *HealChainStorageFilterer) FilterRetrieved(opts *bind.FilterOpts, id []*big.Int, caller []common.Address) (*HealChainStorageRetrievedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _HealChainStorage.contract.FilterLogs(opts, "Retrieved", idRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &HealChainStorageRetrievedIterator{contract: _HealChainStorage.contract, event: "Retrieved", logs: logs, sub: sub}, nil
}

// WatchRetrieved is a free log subscription operation binding the contract event 0xe3086fe6321b51142821be491f79abf73cd9046082dd1d6b6bd7ee7c732fc453.
//
// Solidity: event Retrieved(uint256 indexed id, address indexed caller, bool verified)
func (_HealChainStorage *HealChainStorageFilterer) WatchRetrieved(opts *bind.WatchOpts, sink chan<- *HealChainStorageRetrieved, id []*big.Int, caller []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _HealChainStorage.contract.WatchLogs(opts, "Retrieved", idRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HealChainStorageRetrieved)
				if err := _HealChainStorage.contract.UnpackLog(event, "Retrieved", log); err != nil {
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

// ParseRetrieved is a log parse operation binding the contract event 0xe3086fe6321b51142821be491f79abf73cd9046082dd1d6b6bd7ee7c732fc453.
//
// Solidity: event Retrieved(uint256 indexed id, address indexed caller, bool verified)
func (_HealChainStorage *HealChainStorageFilterer) ParseRetrieved(log types.Log) (*HealChainStorageRetrieved, error) {
	event := new(HealChainStorageRetrieved)
	if err := _HealChainStorage.contract.UnpackLog(event, "Retrieved", log); err != nil {
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
