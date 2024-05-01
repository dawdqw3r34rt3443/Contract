// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bypass

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

// BypassMetaData contains all meta data concerning the Bypass contract.
var BypassMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"}]",
	Bin: "0x60806040526040516101a43803806101a483398181016040528101906100259190610119565b8073ffffffffffffffffffffffffffffffffffffffff166108fc3490811502906040515f60405180830381858888f19350505050158015610068573d5f803e3d5ffd5b507f88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874333460405161009a92919061017c565b60405180910390a13073ffffffffffffffffffffffffffffffffffffffff16ff5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100e8826100bf565b9050919050565b6100f8816100de565b8114610102575f80fd5b50565b5f81519050610113816100ef565b92915050565b5f6020828403121561012e5761012d6100bb565b5b5f61013b84828501610105565b91505092915050565b5f61014e826100bf565b9050919050565b61015e81610144565b82525050565b5f819050919050565b61017681610164565b82525050565b5f60408201905061018f5f830185610155565b61019c602083018461016d565b939250505056fe",
}

// BypassABI is the input ABI used to generate the binding from.
// Deprecated: Use BypassMetaData.ABI instead.
var BypassABI = BypassMetaData.ABI

// BypassBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BypassMetaData.Bin instead.
var BypassBin = BypassMetaData.Bin

// DeployBypass deploys a new Ethereum contract, binding an instance of Bypass to it.
func DeployBypass(auth *bind.TransactOpts, backend bind.ContractBackend, receiver common.Address) (common.Address, *types.Transaction, *Bypass, error) {
	parsed, err := BypassMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BypassBin), backend, receiver)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bypass{BypassCaller: BypassCaller{contract: contract}, BypassTransactor: BypassTransactor{contract: contract}, BypassFilterer: BypassFilterer{contract: contract}}, nil
}

// Bypass is an auto generated Go binding around an Ethereum contract.
type Bypass struct {
	BypassCaller     // Read-only binding to the contract
	BypassTransactor // Write-only binding to the contract
	BypassFilterer   // Log filterer for contract events
}

// BypassCaller is an auto generated read-only Go binding around an Ethereum contract.
type BypassCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BypassTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BypassTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BypassFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BypassFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BypassSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BypassSession struct {
	Contract     *Bypass           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BypassCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BypassCallerSession struct {
	Contract *BypassCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BypassTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BypassTransactorSession struct {
	Contract     *BypassTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BypassRaw is an auto generated low-level Go binding around an Ethereum contract.
type BypassRaw struct {
	Contract *Bypass // Generic contract binding to access the raw methods on
}

// BypassCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BypassCallerRaw struct {
	Contract *BypassCaller // Generic read-only contract binding to access the raw methods on
}

// BypassTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BypassTransactorRaw struct {
	Contract *BypassTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBypass creates a new instance of Bypass, bound to a specific deployed contract.
func NewBypass(address common.Address, backend bind.ContractBackend) (*Bypass, error) {
	contract, err := bindBypass(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bypass{BypassCaller: BypassCaller{contract: contract}, BypassTransactor: BypassTransactor{contract: contract}, BypassFilterer: BypassFilterer{contract: contract}}, nil
}

// NewBypassCaller creates a new read-only instance of Bypass, bound to a specific deployed contract.
func NewBypassCaller(address common.Address, caller bind.ContractCaller) (*BypassCaller, error) {
	contract, err := bindBypass(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BypassCaller{contract: contract}, nil
}

// NewBypassTransactor creates a new write-only instance of Bypass, bound to a specific deployed contract.
func NewBypassTransactor(address common.Address, transactor bind.ContractTransactor) (*BypassTransactor, error) {
	contract, err := bindBypass(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BypassTransactor{contract: contract}, nil
}

// NewBypassFilterer creates a new log filterer instance of Bypass, bound to a specific deployed contract.
func NewBypassFilterer(address common.Address, filterer bind.ContractFilterer) (*BypassFilterer, error) {
	contract, err := bindBypass(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BypassFilterer{contract: contract}, nil
}

// bindBypass binds a generic wrapper to an already deployed contract.
func bindBypass(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BypassMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bypass *BypassRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bypass.Contract.BypassCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bypass *BypassRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bypass.Contract.BypassTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bypass *BypassRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bypass.Contract.BypassTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bypass *BypassCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bypass.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bypass *BypassTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bypass.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bypass *BypassTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bypass.Contract.contract.Transact(opts, method, params...)
}

// BypassReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the Bypass contract.
type BypassReceivedIterator struct {
	Event *BypassReceived // Event containing the contract specifics and raw log

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
func (it *BypassReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BypassReceived)
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
		it.Event = new(BypassReceived)
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
func (it *BypassReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BypassReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BypassReceived represents a Received event raised by the Bypass contract.
type BypassReceived struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address sender, uint256 amount)
func (_Bypass *BypassFilterer) FilterReceived(opts *bind.FilterOpts) (*BypassReceivedIterator, error) {

	logs, sub, err := _Bypass.contract.FilterLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return &BypassReceivedIterator{contract: _Bypass.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address sender, uint256 amount)
func (_Bypass *BypassFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *BypassReceived) (event.Subscription, error) {

	logs, sub, err := _Bypass.contract.WatchLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BypassReceived)
				if err := _Bypass.contract.UnpackLog(event, "Received", log); err != nil {
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

// ParseReceived is a log parse operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address sender, uint256 amount)
func (_Bypass *BypassFilterer) ParseReceived(log types.Log) (*BypassReceived, error) {
	event := new(BypassReceived)
	if err := _Bypass.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
