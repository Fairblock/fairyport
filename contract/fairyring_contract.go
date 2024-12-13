// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// FairyringContractMetaData contains all meta data concerning the FairyringContract contract.
var FairyringContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"encryptionKeyExists\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRandomnessByHeight\",\"inputs\":[{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestEncryptionKey\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestRandomness\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestRandomnessHashOnly\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitDecryptionKey\",\"inputs\":[{\"name\":\"encryptionKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"decryptionKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitEncryptionKey\",\"inputs\":[{\"name\":\"encryptionKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// FairyringContractABI is the input ABI used to generate the binding from.
// Deprecated: Use FairyringContractMetaData.ABI instead.
var FairyringContractABI = FairyringContractMetaData.ABI

// FairyringContract is an auto generated Go binding around an Ethereum contract.
type FairyringContract struct {
	FairyringContractCaller     // Read-only binding to the contract
	FairyringContractTransactor // Write-only binding to the contract
	FairyringContractFilterer   // Log filterer for contract events
}

// FairyringContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type FairyringContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FairyringContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FairyringContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FairyringContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FairyringContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FairyringContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FairyringContractSession struct {
	Contract     *FairyringContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FairyringContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FairyringContractCallerSession struct {
	Contract *FairyringContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// FairyringContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FairyringContractTransactorSession struct {
	Contract     *FairyringContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// FairyringContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type FairyringContractRaw struct {
	Contract *FairyringContract // Generic contract binding to access the raw methods on
}

// FairyringContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FairyringContractCallerRaw struct {
	Contract *FairyringContractCaller // Generic read-only contract binding to access the raw methods on
}

// FairyringContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FairyringContractTransactorRaw struct {
	Contract *FairyringContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFairyringContract creates a new instance of FairyringContract, bound to a specific deployed contract.
func NewFairyringContract(address common.Address, backend bind.ContractBackend) (*FairyringContract, error) {
	contract, err := bindFairyringContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FairyringContract{FairyringContractCaller: FairyringContractCaller{contract: contract}, FairyringContractTransactor: FairyringContractTransactor{contract: contract}, FairyringContractFilterer: FairyringContractFilterer{contract: contract}}, nil
}

// NewFairyringContractCaller creates a new read-only instance of FairyringContract, bound to a specific deployed contract.
func NewFairyringContractCaller(address common.Address, caller bind.ContractCaller) (*FairyringContractCaller, error) {
	contract, err := bindFairyringContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FairyringContractCaller{contract: contract}, nil
}

// NewFairyringContractTransactor creates a new write-only instance of FairyringContract, bound to a specific deployed contract.
func NewFairyringContractTransactor(address common.Address, transactor bind.ContractTransactor) (*FairyringContractTransactor, error) {
	contract, err := bindFairyringContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FairyringContractTransactor{contract: contract}, nil
}

// NewFairyringContractFilterer creates a new log filterer instance of FairyringContract, bound to a specific deployed contract.
func NewFairyringContractFilterer(address common.Address, filterer bind.ContractFilterer) (*FairyringContractFilterer, error) {
	contract, err := bindFairyringContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FairyringContractFilterer{contract: contract}, nil
}

// bindFairyringContract binds a generic wrapper to an already deployed contract.
func bindFairyringContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FairyringContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FairyringContract *FairyringContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FairyringContract.Contract.FairyringContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FairyringContract *FairyringContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FairyringContract.Contract.FairyringContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FairyringContract *FairyringContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FairyringContract.Contract.FairyringContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FairyringContract *FairyringContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FairyringContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FairyringContract *FairyringContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FairyringContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FairyringContract *FairyringContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FairyringContract.Contract.contract.Transact(opts, method, params...)
}

// EncryptionKeyExists is a free data retrieval call binding the contract method 0x9c533b8a.
//
// Solidity: function encryptionKeyExists(bytes ) view returns(bool)
func (_FairyringContract *FairyringContractCaller) EncryptionKeyExists(opts *bind.CallOpts, arg0 []byte) (bool, error) {
	var out []interface{}
	err := _FairyringContract.contract.Call(opts, &out, "encryptionKeyExists", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EncryptionKeyExists is a free data retrieval call binding the contract method 0x9c533b8a.
//
// Solidity: function encryptionKeyExists(bytes ) view returns(bool)
func (_FairyringContract *FairyringContractSession) EncryptionKeyExists(arg0 []byte) (bool, error) {
	return _FairyringContract.Contract.EncryptionKeyExists(&_FairyringContract.CallOpts, arg0)
}

// EncryptionKeyExists is a free data retrieval call binding the contract method 0x9c533b8a.
//
// Solidity: function encryptionKeyExists(bytes ) view returns(bool)
func (_FairyringContract *FairyringContractCallerSession) EncryptionKeyExists(arg0 []byte) (bool, error) {
	return _FairyringContract.Contract.EncryptionKeyExists(&_FairyringContract.CallOpts, arg0)
}

// GetRandomnessByHeight is a free data retrieval call binding the contract method 0xee4d43cf.
//
// Solidity: function getRandomnessByHeight(uint256 height) view returns(bytes32)
func (_FairyringContract *FairyringContractCaller) GetRandomnessByHeight(opts *bind.CallOpts, height *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _FairyringContract.contract.Call(opts, &out, "getRandomnessByHeight", height)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRandomnessByHeight is a free data retrieval call binding the contract method 0xee4d43cf.
//
// Solidity: function getRandomnessByHeight(uint256 height) view returns(bytes32)
func (_FairyringContract *FairyringContractSession) GetRandomnessByHeight(height *big.Int) ([32]byte, error) {
	return _FairyringContract.Contract.GetRandomnessByHeight(&_FairyringContract.CallOpts, height)
}

// GetRandomnessByHeight is a free data retrieval call binding the contract method 0xee4d43cf.
//
// Solidity: function getRandomnessByHeight(uint256 height) view returns(bytes32)
func (_FairyringContract *FairyringContractCallerSession) GetRandomnessByHeight(height *big.Int) ([32]byte, error) {
	return _FairyringContract.Contract.GetRandomnessByHeight(&_FairyringContract.CallOpts, height)
}

// LatestEncryptionKey is a free data retrieval call binding the contract method 0x74236c86.
//
// Solidity: function latestEncryptionKey() view returns(bytes)
func (_FairyringContract *FairyringContractCaller) LatestEncryptionKey(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _FairyringContract.contract.Call(opts, &out, "latestEncryptionKey")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// LatestEncryptionKey is a free data retrieval call binding the contract method 0x74236c86.
//
// Solidity: function latestEncryptionKey() view returns(bytes)
func (_FairyringContract *FairyringContractSession) LatestEncryptionKey() ([]byte, error) {
	return _FairyringContract.Contract.LatestEncryptionKey(&_FairyringContract.CallOpts)
}

// LatestEncryptionKey is a free data retrieval call binding the contract method 0x74236c86.
//
// Solidity: function latestEncryptionKey() view returns(bytes)
func (_FairyringContract *FairyringContractCallerSession) LatestEncryptionKey() ([]byte, error) {
	return _FairyringContract.Contract.LatestEncryptionKey(&_FairyringContract.CallOpts)
}

// LatestRandomness is a free data retrieval call binding the contract method 0x3b5cd3a8.
//
// Solidity: function latestRandomness() view returns(bytes32, uint256)
func (_FairyringContract *FairyringContractCaller) LatestRandomness(opts *bind.CallOpts) ([32]byte, *big.Int, error) {
	var out []interface{}
	err := _FairyringContract.contract.Call(opts, &out, "latestRandomness")

	if err != nil {
		return *new([32]byte), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// LatestRandomness is a free data retrieval call binding the contract method 0x3b5cd3a8.
//
// Solidity: function latestRandomness() view returns(bytes32, uint256)
func (_FairyringContract *FairyringContractSession) LatestRandomness() ([32]byte, *big.Int, error) {
	return _FairyringContract.Contract.LatestRandomness(&_FairyringContract.CallOpts)
}

// LatestRandomness is a free data retrieval call binding the contract method 0x3b5cd3a8.
//
// Solidity: function latestRandomness() view returns(bytes32, uint256)
func (_FairyringContract *FairyringContractCallerSession) LatestRandomness() ([32]byte, *big.Int, error) {
	return _FairyringContract.Contract.LatestRandomness(&_FairyringContract.CallOpts)
}

// LatestRandomnessHashOnly is a free data retrieval call binding the contract method 0xd0d343a7.
//
// Solidity: function latestRandomnessHashOnly() view returns(bytes32)
func (_FairyringContract *FairyringContractCaller) LatestRandomnessHashOnly(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FairyringContract.contract.Call(opts, &out, "latestRandomnessHashOnly")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LatestRandomnessHashOnly is a free data retrieval call binding the contract method 0xd0d343a7.
//
// Solidity: function latestRandomnessHashOnly() view returns(bytes32)
func (_FairyringContract *FairyringContractSession) LatestRandomnessHashOnly() ([32]byte, error) {
	return _FairyringContract.Contract.LatestRandomnessHashOnly(&_FairyringContract.CallOpts)
}

// LatestRandomnessHashOnly is a free data retrieval call binding the contract method 0xd0d343a7.
//
// Solidity: function latestRandomnessHashOnly() view returns(bytes32)
func (_FairyringContract *FairyringContractCallerSession) LatestRandomnessHashOnly() ([32]byte, error) {
	return _FairyringContract.Contract.LatestRandomnessHashOnly(&_FairyringContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FairyringContract *FairyringContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FairyringContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FairyringContract *FairyringContractSession) Owner() (common.Address, error) {
	return _FairyringContract.Contract.Owner(&_FairyringContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FairyringContract *FairyringContractCallerSession) Owner() (common.Address, error) {
	return _FairyringContract.Contract.Owner(&_FairyringContract.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FairyringContract *FairyringContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FairyringContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FairyringContract *FairyringContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _FairyringContract.Contract.RenounceOwnership(&_FairyringContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FairyringContract *FairyringContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FairyringContract.Contract.RenounceOwnership(&_FairyringContract.TransactOpts)
}

// SubmitDecryptionKey is a paid mutator transaction binding the contract method 0xaa878cf2.
//
// Solidity: function submitDecryptionKey(bytes encryptionKey, bytes decryptionKey, uint256 height) returns()
func (_FairyringContract *FairyringContractTransactor) SubmitDecryptionKey(opts *bind.TransactOpts, encryptionKey []byte, decryptionKey []byte, height *big.Int) (*types.Transaction, error) {
	return _FairyringContract.contract.Transact(opts, "submitDecryptionKey", encryptionKey, decryptionKey, height)
}

// SubmitDecryptionKey is a paid mutator transaction binding the contract method 0xaa878cf2.
//
// Solidity: function submitDecryptionKey(bytes encryptionKey, bytes decryptionKey, uint256 height) returns()
func (_FairyringContract *FairyringContractSession) SubmitDecryptionKey(encryptionKey []byte, decryptionKey []byte, height *big.Int) (*types.Transaction, error) {
	return _FairyringContract.Contract.SubmitDecryptionKey(&_FairyringContract.TransactOpts, encryptionKey, decryptionKey, height)
}

// SubmitDecryptionKey is a paid mutator transaction binding the contract method 0xaa878cf2.
//
// Solidity: function submitDecryptionKey(bytes encryptionKey, bytes decryptionKey, uint256 height) returns()
func (_FairyringContract *FairyringContractTransactorSession) SubmitDecryptionKey(encryptionKey []byte, decryptionKey []byte, height *big.Int) (*types.Transaction, error) {
	return _FairyringContract.Contract.SubmitDecryptionKey(&_FairyringContract.TransactOpts, encryptionKey, decryptionKey, height)
}

// SubmitEncryptionKey is a paid mutator transaction binding the contract method 0x0d006aa2.
//
// Solidity: function submitEncryptionKey(bytes encryptionKey) returns()
func (_FairyringContract *FairyringContractTransactor) SubmitEncryptionKey(opts *bind.TransactOpts, encryptionKey []byte) (*types.Transaction, error) {
	return _FairyringContract.contract.Transact(opts, "submitEncryptionKey", encryptionKey)
}

// SubmitEncryptionKey is a paid mutator transaction binding the contract method 0x0d006aa2.
//
// Solidity: function submitEncryptionKey(bytes encryptionKey) returns()
func (_FairyringContract *FairyringContractSession) SubmitEncryptionKey(encryptionKey []byte) (*types.Transaction, error) {
	return _FairyringContract.Contract.SubmitEncryptionKey(&_FairyringContract.TransactOpts, encryptionKey)
}

// SubmitEncryptionKey is a paid mutator transaction binding the contract method 0x0d006aa2.
//
// Solidity: function submitEncryptionKey(bytes encryptionKey) returns()
func (_FairyringContract *FairyringContractTransactorSession) SubmitEncryptionKey(encryptionKey []byte) (*types.Transaction, error) {
	return _FairyringContract.Contract.SubmitEncryptionKey(&_FairyringContract.TransactOpts, encryptionKey)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FairyringContract *FairyringContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FairyringContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FairyringContract *FairyringContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FairyringContract.Contract.TransferOwnership(&_FairyringContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FairyringContract *FairyringContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FairyringContract.Contract.TransferOwnership(&_FairyringContract.TransactOpts, newOwner)
}

// FairyringContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FairyringContract contract.
type FairyringContractOwnershipTransferredIterator struct {
	Event *FairyringContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FairyringContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FairyringContractOwnershipTransferred)
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
		it.Event = new(FairyringContractOwnershipTransferred)
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
func (it *FairyringContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FairyringContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FairyringContractOwnershipTransferred represents a OwnershipTransferred event raised by the FairyringContract contract.
type FairyringContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FairyringContract *FairyringContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FairyringContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FairyringContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FairyringContractOwnershipTransferredIterator{contract: _FairyringContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FairyringContract *FairyringContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FairyringContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FairyringContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FairyringContractOwnershipTransferred)
				if err := _FairyringContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FairyringContract *FairyringContractFilterer) ParseOwnershipTransferred(log types.Log) (*FairyringContractOwnershipTransferred, error) {
	event := new(FairyringContractOwnershipTransferred)
	if err := _FairyringContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
