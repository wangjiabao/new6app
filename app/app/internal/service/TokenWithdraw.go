// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package service

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

// TokenWithdrawMetaData contains all meta data concerning the TokenWithdraw contract.
var TokenWithdrawMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"nextWithdrawTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeLimit_\",\"type\":\"uint256\"}],\"name\":\"setTimeLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"withdrawAddress_\",\"type\":\"address\"}],\"name\":\"setWithdrawAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"setWithdrawAmountLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sx\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAmountLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawSx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TokenWithdrawABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenWithdrawMetaData.ABI instead.
var TokenWithdrawABI = TokenWithdrawMetaData.ABI

// TokenWithdraw is an auto generated Go binding around an Ethereum contract.
type TokenWithdraw struct {
	TokenWithdrawCaller     // Read-only binding to the contract
	TokenWithdrawTransactor // Write-only binding to the contract
	TokenWithdrawFilterer   // Log filterer for contract events
}

// TokenWithdrawCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenWithdrawCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenWithdrawTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenWithdrawTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenWithdrawFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenWithdrawFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenWithdrawSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenWithdrawSession struct {
	Contract     *TokenWithdraw    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenWithdrawCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenWithdrawCallerSession struct {
	Contract *TokenWithdrawCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TokenWithdrawTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenWithdrawTransactorSession struct {
	Contract     *TokenWithdrawTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TokenWithdrawRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenWithdrawRaw struct {
	Contract *TokenWithdraw // Generic contract binding to access the raw methods on
}

// TokenWithdrawCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenWithdrawCallerRaw struct {
	Contract *TokenWithdrawCaller // Generic read-only contract binding to access the raw methods on
}

// TokenWithdrawTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenWithdrawTransactorRaw struct {
	Contract *TokenWithdrawTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenWithdraw creates a new instance of TokenWithdraw, bound to a specific deployed contract.
func NewTokenWithdraw(address common.Address, backend bind.ContractBackend) (*TokenWithdraw, error) {
	contract, err := bindTokenWithdraw(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenWithdraw{TokenWithdrawCaller: TokenWithdrawCaller{contract: contract}, TokenWithdrawTransactor: TokenWithdrawTransactor{contract: contract}, TokenWithdrawFilterer: TokenWithdrawFilterer{contract: contract}}, nil
}

// NewTokenWithdrawCaller creates a new read-only instance of TokenWithdraw, bound to a specific deployed contract.
func NewTokenWithdrawCaller(address common.Address, caller bind.ContractCaller) (*TokenWithdrawCaller, error) {
	contract, err := bindTokenWithdraw(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenWithdrawCaller{contract: contract}, nil
}

// NewTokenWithdrawTransactor creates a new write-only instance of TokenWithdraw, bound to a specific deployed contract.
func NewTokenWithdrawTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenWithdrawTransactor, error) {
	contract, err := bindTokenWithdraw(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenWithdrawTransactor{contract: contract}, nil
}

// NewTokenWithdrawFilterer creates a new log filterer instance of TokenWithdraw, bound to a specific deployed contract.
func NewTokenWithdrawFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenWithdrawFilterer, error) {
	contract, err := bindTokenWithdraw(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenWithdrawFilterer{contract: contract}, nil
}

// bindTokenWithdraw binds a generic wrapper to an already deployed contract.
func bindTokenWithdraw(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenWithdrawABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenWithdraw *TokenWithdrawRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenWithdraw.Contract.TokenWithdrawCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenWithdraw *TokenWithdrawRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenWithdraw.Contract.TokenWithdrawTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenWithdraw *TokenWithdrawRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenWithdraw.Contract.TokenWithdrawTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenWithdraw *TokenWithdrawCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenWithdraw.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenWithdraw *TokenWithdrawTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenWithdraw.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenWithdraw *TokenWithdrawTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenWithdraw.Contract.contract.Transact(opts, method, params...)
}

// NextWithdrawTime is a free data retrieval call binding the contract method 0xd6b142b2.
//
// Solidity: function nextWithdrawTime() view returns(uint256)
func (_TokenWithdraw *TokenWithdrawCaller) NextWithdrawTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenWithdraw.contract.Call(opts, &out, "nextWithdrawTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextWithdrawTime is a free data retrieval call binding the contract method 0xd6b142b2.
//
// Solidity: function nextWithdrawTime() view returns(uint256)
func (_TokenWithdraw *TokenWithdrawSession) NextWithdrawTime() (*big.Int, error) {
	return _TokenWithdraw.Contract.NextWithdrawTime(&_TokenWithdraw.CallOpts)
}

// NextWithdrawTime is a free data retrieval call binding the contract method 0xd6b142b2.
//
// Solidity: function nextWithdrawTime() view returns(uint256)
func (_TokenWithdraw *TokenWithdrawCallerSession) NextWithdrawTime() (*big.Int, error) {
	return _TokenWithdraw.Contract.NextWithdrawTime(&_TokenWithdraw.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenWithdraw *TokenWithdrawCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenWithdraw.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenWithdraw *TokenWithdrawSession) Owner() (common.Address, error) {
	return _TokenWithdraw.Contract.Owner(&_TokenWithdraw.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenWithdraw *TokenWithdrawCallerSession) Owner() (common.Address, error) {
	return _TokenWithdraw.Contract.Owner(&_TokenWithdraw.CallOpts)
}

// Sx is a free data retrieval call binding the contract method 0x90b82161.
//
// Solidity: function sx() view returns(address)
func (_TokenWithdraw *TokenWithdrawCaller) Sx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenWithdraw.contract.Call(opts, &out, "sx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sx is a free data retrieval call binding the contract method 0x90b82161.
//
// Solidity: function sx() view returns(address)
func (_TokenWithdraw *TokenWithdrawSession) Sx() (common.Address, error) {
	return _TokenWithdraw.Contract.Sx(&_TokenWithdraw.CallOpts)
}

// Sx is a free data retrieval call binding the contract method 0x90b82161.
//
// Solidity: function sx() view returns(address)
func (_TokenWithdraw *TokenWithdrawCallerSession) Sx() (common.Address, error) {
	return _TokenWithdraw.Contract.Sx(&_TokenWithdraw.CallOpts)
}

// TimeLimit is a free data retrieval call binding the contract method 0xc08d1fe5.
//
// Solidity: function timeLimit() view returns(uint256)
func (_TokenWithdraw *TokenWithdrawCaller) TimeLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenWithdraw.contract.Call(opts, &out, "timeLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeLimit is a free data retrieval call binding the contract method 0xc08d1fe5.
//
// Solidity: function timeLimit() view returns(uint256)
func (_TokenWithdraw *TokenWithdrawSession) TimeLimit() (*big.Int, error) {
	return _TokenWithdraw.Contract.TimeLimit(&_TokenWithdraw.CallOpts)
}

// TimeLimit is a free data retrieval call binding the contract method 0xc08d1fe5.
//
// Solidity: function timeLimit() view returns(uint256)
func (_TokenWithdraw *TokenWithdrawCallerSession) TimeLimit() (*big.Int, error) {
	return _TokenWithdraw.Contract.TimeLimit(&_TokenWithdraw.CallOpts)
}

// WithdrawAddress is a free data retrieval call binding the contract method 0x1581b600.
//
// Solidity: function withdrawAddress() view returns(address)
func (_TokenWithdraw *TokenWithdrawCaller) WithdrawAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenWithdraw.contract.Call(opts, &out, "withdrawAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WithdrawAddress is a free data retrieval call binding the contract method 0x1581b600.
//
// Solidity: function withdrawAddress() view returns(address)
func (_TokenWithdraw *TokenWithdrawSession) WithdrawAddress() (common.Address, error) {
	return _TokenWithdraw.Contract.WithdrawAddress(&_TokenWithdraw.CallOpts)
}

// WithdrawAddress is a free data retrieval call binding the contract method 0x1581b600.
//
// Solidity: function withdrawAddress() view returns(address)
func (_TokenWithdraw *TokenWithdrawCallerSession) WithdrawAddress() (common.Address, error) {
	return _TokenWithdraw.Contract.WithdrawAddress(&_TokenWithdraw.CallOpts)
}

// WithdrawAmountLimit is a free data retrieval call binding the contract method 0x2f698c62.
//
// Solidity: function withdrawAmountLimit() view returns(uint256)
func (_TokenWithdraw *TokenWithdrawCaller) WithdrawAmountLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenWithdraw.contract.Call(opts, &out, "withdrawAmountLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawAmountLimit is a free data retrieval call binding the contract method 0x2f698c62.
//
// Solidity: function withdrawAmountLimit() view returns(uint256)
func (_TokenWithdraw *TokenWithdrawSession) WithdrawAmountLimit() (*big.Int, error) {
	return _TokenWithdraw.Contract.WithdrawAmountLimit(&_TokenWithdraw.CallOpts)
}

// WithdrawAmountLimit is a free data retrieval call binding the contract method 0x2f698c62.
//
// Solidity: function withdrawAmountLimit() view returns(uint256)
func (_TokenWithdraw *TokenWithdrawCallerSession) WithdrawAmountLimit() (*big.Int, error) {
	return _TokenWithdraw.Contract.WithdrawAmountLimit(&_TokenWithdraw.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenWithdraw *TokenWithdrawTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenWithdraw.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenWithdraw *TokenWithdrawSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenWithdraw.Contract.RenounceOwnership(&_TokenWithdraw.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenWithdraw *TokenWithdrawTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenWithdraw.Contract.RenounceOwnership(&_TokenWithdraw.TransactOpts)
}

// SetTimeLimit is a paid mutator transaction binding the contract method 0xe2889c82.
//
// Solidity: function setTimeLimit(uint256 timeLimit_) returns()
func (_TokenWithdraw *TokenWithdrawTransactor) SetTimeLimit(opts *bind.TransactOpts, timeLimit_ *big.Int) (*types.Transaction, error) {
	return _TokenWithdraw.contract.Transact(opts, "setTimeLimit", timeLimit_)
}

// SetTimeLimit is a paid mutator transaction binding the contract method 0xe2889c82.
//
// Solidity: function setTimeLimit(uint256 timeLimit_) returns()
func (_TokenWithdraw *TokenWithdrawSession) SetTimeLimit(timeLimit_ *big.Int) (*types.Transaction, error) {
	return _TokenWithdraw.Contract.SetTimeLimit(&_TokenWithdraw.TransactOpts, timeLimit_)
}

// SetTimeLimit is a paid mutator transaction binding the contract method 0xe2889c82.
//
// Solidity: function setTimeLimit(uint256 timeLimit_) returns()
func (_TokenWithdraw *TokenWithdrawTransactorSession) SetTimeLimit(timeLimit_ *big.Int) (*types.Transaction, error) {
	return _TokenWithdraw.Contract.SetTimeLimit(&_TokenWithdraw.TransactOpts, timeLimit_)
}

// SetWithdrawAddress is a paid mutator transaction binding the contract method 0x3ab1a494.
//
// Solidity: function setWithdrawAddress(address withdrawAddress_) returns()
func (_TokenWithdraw *TokenWithdrawTransactor) SetWithdrawAddress(opts *bind.TransactOpts, withdrawAddress_ common.Address) (*types.Transaction, error) {
	return _TokenWithdraw.contract.Transact(opts, "setWithdrawAddress", withdrawAddress_)
}

// SetWithdrawAddress is a paid mutator transaction binding the contract method 0x3ab1a494.
//
// Solidity: function setWithdrawAddress(address withdrawAddress_) returns()
func (_TokenWithdraw *TokenWithdrawSession) SetWithdrawAddress(withdrawAddress_ common.Address) (*types.Transaction, error) {
	return _TokenWithdraw.Contract.SetWithdrawAddress(&_TokenWithdraw.TransactOpts, withdrawAddress_)
}

// SetWithdrawAddress is a paid mutator transaction binding the contract method 0x3ab1a494.
//
// Solidity: function setWithdrawAddress(address withdrawAddress_) returns()
func (_TokenWithdraw *TokenWithdrawTransactorSession) SetWithdrawAddress(withdrawAddress_ common.Address) (*types.Transaction, error) {
	return _TokenWithdraw.Contract.SetWithdrawAddress(&_TokenWithdraw.TransactOpts, withdrawAddress_)
}

// SetWithdrawAmountLimit is a paid mutator transaction binding the contract method 0xe44c8d2a.
//
// Solidity: function setWithdrawAmountLimit(uint256 withdrawAmount) returns()
func (_TokenWithdraw *TokenWithdrawTransactor) SetWithdrawAmountLimit(opts *bind.TransactOpts, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _TokenWithdraw.contract.Transact(opts, "setWithdrawAmountLimit", withdrawAmount)
}

// SetWithdrawAmountLimit is a paid mutator transaction binding the contract method 0xe44c8d2a.
//
// Solidity: function setWithdrawAmountLimit(uint256 withdrawAmount) returns()
func (_TokenWithdraw *TokenWithdrawSession) SetWithdrawAmountLimit(withdrawAmount *big.Int) (*types.Transaction, error) {
	return _TokenWithdraw.Contract.SetWithdrawAmountLimit(&_TokenWithdraw.TransactOpts, withdrawAmount)
}

// SetWithdrawAmountLimit is a paid mutator transaction binding the contract method 0xe44c8d2a.
//
// Solidity: function setWithdrawAmountLimit(uint256 withdrawAmount) returns()
func (_TokenWithdraw *TokenWithdrawTransactorSession) SetWithdrawAmountLimit(withdrawAmount *big.Int) (*types.Transaction, error) {
	return _TokenWithdraw.Contract.SetWithdrawAmountLimit(&_TokenWithdraw.TransactOpts, withdrawAmount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenWithdraw *TokenWithdrawTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TokenWithdraw.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenWithdraw *TokenWithdrawSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenWithdraw.Contract.TransferOwnership(&_TokenWithdraw.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenWithdraw *TokenWithdrawTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenWithdraw.Contract.TransferOwnership(&_TokenWithdraw.TransactOpts, newOwner)
}

// WithdrawSx is a paid mutator transaction binding the contract method 0x15cd802a.
//
// Solidity: function withdrawSx() returns()
func (_TokenWithdraw *TokenWithdrawTransactor) WithdrawSx(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenWithdraw.contract.Transact(opts, "withdrawSx")
}

// WithdrawSx is a paid mutator transaction binding the contract method 0x15cd802a.
//
// Solidity: function withdrawSx() returns()
func (_TokenWithdraw *TokenWithdrawSession) WithdrawSx() (*types.Transaction, error) {
	return _TokenWithdraw.Contract.WithdrawSx(&_TokenWithdraw.TransactOpts)
}

// WithdrawSx is a paid mutator transaction binding the contract method 0x15cd802a.
//
// Solidity: function withdrawSx() returns()
func (_TokenWithdraw *TokenWithdrawTransactorSession) WithdrawSx() (*types.Transaction, error) {
	return _TokenWithdraw.Contract.WithdrawSx(&_TokenWithdraw.TransactOpts)
}

// TokenWithdrawOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TokenWithdraw contract.
type TokenWithdrawOwnershipTransferredIterator struct {
	Event *TokenWithdrawOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenWithdrawOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWithdrawOwnershipTransferred)
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
		it.Event = new(TokenWithdrawOwnershipTransferred)
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
func (it *TokenWithdrawOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWithdrawOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWithdrawOwnershipTransferred represents a OwnershipTransferred event raised by the TokenWithdraw contract.
type TokenWithdrawOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenWithdraw *TokenWithdrawFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenWithdrawOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenWithdraw.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenWithdrawOwnershipTransferredIterator{contract: _TokenWithdraw.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenWithdraw *TokenWithdrawFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenWithdrawOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenWithdraw.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWithdrawOwnershipTransferred)
				if err := _TokenWithdraw.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TokenWithdraw *TokenWithdrawFilterer) ParseOwnershipTransferred(log types.Log) (*TokenWithdrawOwnershipTransferred, error) {
	event := new(TokenWithdrawOwnershipTransferred)
	if err := _TokenWithdraw.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
