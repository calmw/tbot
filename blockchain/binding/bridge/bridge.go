// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridge

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

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ErrCallType\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"callType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"}],\"name\":\"CallLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"DepositData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dataLen\",\"type\":\"uint256\"}],\"name\":\"DepositDataLen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"callType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"}],\"name\":\"executeLog\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allAdmins\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"admins\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"amounts\",\"type\":\"address[]\"}],\"name\":\"batchAddAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeHandlerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"callType\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"callType\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"}],\"name\":\"depositCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"depositNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"someNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callType\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"callType\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getBytesData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"toxAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"usdtAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"miningAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridgeAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridgeHandlerAddress_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mining\",\"outputs\":[{\"internalType\":\"contractIMining\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"miningAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"topUpTest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tox\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdt\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeMetaData.ABI instead.
var BridgeABI = BridgeMetaData.ABI

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// AllAdmins is a free data retrieval call binding the contract method 0x40f32be6.
//
// Solidity: function allAdmins() view returns(address[] admins)
func (_Bridge *BridgeCaller) AllAdmins(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "allAdmins")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllAdmins is a free data retrieval call binding the contract method 0x40f32be6.
//
// Solidity: function allAdmins() view returns(address[] admins)
func (_Bridge *BridgeSession) AllAdmins() ([]common.Address, error) {
	return _Bridge.Contract.AllAdmins(&_Bridge.CallOpts)
}

// AllAdmins is a free data retrieval call binding the contract method 0x40f32be6.
//
// Solidity: function allAdmins() view returns(address[] admins)
func (_Bridge *BridgeCallerSession) AllAdmins() ([]common.Address, error) {
	return _Bridge.Contract.AllAdmins(&_Bridge.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Bridge *BridgeCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Bridge *BridgeSession) BridgeAddress() (common.Address, error) {
	return _Bridge.Contract.BridgeAddress(&_Bridge.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Bridge *BridgeCallerSession) BridgeAddress() (common.Address, error) {
	return _Bridge.Contract.BridgeAddress(&_Bridge.CallOpts)
}

// BridgeHandlerAddress is a free data retrieval call binding the contract method 0x8d7b0d4e.
//
// Solidity: function bridgeHandlerAddress() view returns(address)
func (_Bridge *BridgeCaller) BridgeHandlerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "bridgeHandlerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeHandlerAddress is a free data retrieval call binding the contract method 0x8d7b0d4e.
//
// Solidity: function bridgeHandlerAddress() view returns(address)
func (_Bridge *BridgeSession) BridgeHandlerAddress() (common.Address, error) {
	return _Bridge.Contract.BridgeHandlerAddress(&_Bridge.CallOpts)
}

// BridgeHandlerAddress is a free data retrieval call binding the contract method 0x8d7b0d4e.
//
// Solidity: function bridgeHandlerAddress() view returns(address)
func (_Bridge *BridgeCallerSession) BridgeHandlerAddress() (common.Address, error) {
	return _Bridge.Contract.BridgeHandlerAddress(&_Bridge.CallOpts)
}

// DepositNonce is a free data retrieval call binding the contract method 0x87e306f5.
//
// Solidity: function depositNonce(address ) view returns(uint64)
func (_Bridge *BridgeCaller) DepositNonce(opts *bind.CallOpts, arg0 common.Address) (uint64, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "depositNonce", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DepositNonce is a free data retrieval call binding the contract method 0x87e306f5.
//
// Solidity: function depositNonce(address ) view returns(uint64)
func (_Bridge *BridgeSession) DepositNonce(arg0 common.Address) (uint64, error) {
	return _Bridge.Contract.DepositNonce(&_Bridge.CallOpts, arg0)
}

// DepositNonce is a free data retrieval call binding the contract method 0x87e306f5.
//
// Solidity: function depositNonce(address ) view returns(uint64)
func (_Bridge *BridgeCallerSession) DepositNonce(arg0 common.Address) (uint64, error) {
	return _Bridge.Contract.DepositNonce(&_Bridge.CallOpts, arg0)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_Bridge *BridgeCaller) IsAdmin(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "isAdmin", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_Bridge *BridgeSession) IsAdmin(account common.Address) (bool, error) {
	return _Bridge.Contract.IsAdmin(&_Bridge.CallOpts, account)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_Bridge *BridgeCallerSession) IsAdmin(account common.Address) (bool, error) {
	return _Bridge.Contract.IsAdmin(&_Bridge.CallOpts, account)
}

// Mining is a free data retrieval call binding the contract method 0x662fac39.
//
// Solidity: function mining() view returns(address)
func (_Bridge *BridgeCaller) Mining(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "mining")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mining is a free data retrieval call binding the contract method 0x662fac39.
//
// Solidity: function mining() view returns(address)
func (_Bridge *BridgeSession) Mining() (common.Address, error) {
	return _Bridge.Contract.Mining(&_Bridge.CallOpts)
}

// Mining is a free data retrieval call binding the contract method 0x662fac39.
//
// Solidity: function mining() view returns(address)
func (_Bridge *BridgeCallerSession) Mining() (common.Address, error) {
	return _Bridge.Contract.Mining(&_Bridge.CallOpts)
}

// MiningAddress is a free data retrieval call binding the contract method 0x88ece43f.
//
// Solidity: function miningAddress() view returns(address)
func (_Bridge *BridgeCaller) MiningAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "miningAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MiningAddress is a free data retrieval call binding the contract method 0x88ece43f.
//
// Solidity: function miningAddress() view returns(address)
func (_Bridge *BridgeSession) MiningAddress() (common.Address, error) {
	return _Bridge.Contract.MiningAddress(&_Bridge.CallOpts)
}

// MiningAddress is a free data retrieval call binding the contract method 0x88ece43f.
//
// Solidity: function miningAddress() view returns(address)
func (_Bridge *BridgeCallerSession) MiningAddress() (common.Address, error) {
	return _Bridge.Contract.MiningAddress(&_Bridge.CallOpts)
}

// Tox is a free data retrieval call binding the contract method 0x46c811fd.
//
// Solidity: function tox() view returns(address)
func (_Bridge *BridgeCaller) Tox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "tox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tox is a free data retrieval call binding the contract method 0x46c811fd.
//
// Solidity: function tox() view returns(address)
func (_Bridge *BridgeSession) Tox() (common.Address, error) {
	return _Bridge.Contract.Tox(&_Bridge.CallOpts)
}

// Tox is a free data retrieval call binding the contract method 0x46c811fd.
//
// Solidity: function tox() view returns(address)
func (_Bridge *BridgeCallerSession) Tox() (common.Address, error) {
	return _Bridge.Contract.Tox(&_Bridge.CallOpts)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_Bridge *BridgeCaller) Usdt(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "usdt")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_Bridge *BridgeSession) Usdt() (common.Address, error) {
	return _Bridge.Contract.Usdt(&_Bridge.CallOpts)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_Bridge *BridgeCallerSession) Usdt() (common.Address, error) {
	return _Bridge.Contract.Usdt(&_Bridge.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_Bridge *BridgeTransactor) AddAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "addAdmin", account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_Bridge *BridgeSession) AddAdmin(account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AddAdmin(&_Bridge.TransactOpts, account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_Bridge *BridgeTransactorSession) AddAdmin(account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AddAdmin(&_Bridge.TransactOpts, account)
}

// BatchAddAdmin is a paid mutator transaction binding the contract method 0x2c9ab42b.
//
// Solidity: function batchAddAdmin(address[] amounts) returns()
func (_Bridge *BridgeTransactor) BatchAddAdmin(opts *bind.TransactOpts, amounts []common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "batchAddAdmin", amounts)
}

// BatchAddAdmin is a paid mutator transaction binding the contract method 0x2c9ab42b.
//
// Solidity: function batchAddAdmin(address[] amounts) returns()
func (_Bridge *BridgeSession) BatchAddAdmin(amounts []common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.BatchAddAdmin(&_Bridge.TransactOpts, amounts)
}

// BatchAddAdmin is a paid mutator transaction binding the contract method 0x2c9ab42b.
//
// Solidity: function batchAddAdmin(address[] amounts) returns()
func (_Bridge *BridgeTransactorSession) BatchAddAdmin(amounts []common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.BatchAddAdmin(&_Bridge.TransactOpts, amounts)
}

// Deposit is a paid mutator transaction binding the contract method 0x528e1bc2.
//
// Solidity: function deposit(uint8 destinationChainID, bytes32 resourceID, uint256 callType, address recipient, uint256 amount) payable returns()
func (_Bridge *BridgeTransactor) Deposit(opts *bind.TransactOpts, destinationChainID uint8, resourceID [32]byte, callType *big.Int, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "deposit", destinationChainID, resourceID, callType, recipient, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x528e1bc2.
//
// Solidity: function deposit(uint8 destinationChainID, bytes32 resourceID, uint256 callType, address recipient, uint256 amount) payable returns()
func (_Bridge *BridgeSession) Deposit(destinationChainID uint8, resourceID [32]byte, callType *big.Int, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Deposit(&_Bridge.TransactOpts, destinationChainID, resourceID, callType, recipient, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x528e1bc2.
//
// Solidity: function deposit(uint8 destinationChainID, bytes32 resourceID, uint256 callType, address recipient, uint256 amount) payable returns()
func (_Bridge *BridgeTransactorSession) Deposit(destinationChainID uint8, resourceID [32]byte, callType *big.Int, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Deposit(&_Bridge.TransactOpts, destinationChainID, resourceID, callType, recipient, amount)
}

// DepositCallback is a paid mutator transaction binding the contract method 0x9e165780.
//
// Solidity: function depositCallback(uint256 callType, address caller, address recipient, uint256 amount, uint64 depositNonce) returns()
func (_Bridge *BridgeTransactor) DepositCallback(opts *bind.TransactOpts, callType *big.Int, caller common.Address, recipient common.Address, amount *big.Int, depositNonce uint64) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "depositCallback", callType, caller, recipient, amount, depositNonce)
}

// DepositCallback is a paid mutator transaction binding the contract method 0x9e165780.
//
// Solidity: function depositCallback(uint256 callType, address caller, address recipient, uint256 amount, uint64 depositNonce) returns()
func (_Bridge *BridgeSession) DepositCallback(callType *big.Int, caller common.Address, recipient common.Address, amount *big.Int, depositNonce uint64) (*types.Transaction, error) {
	return _Bridge.Contract.DepositCallback(&_Bridge.TransactOpts, callType, caller, recipient, amount, depositNonce)
}

// DepositCallback is a paid mutator transaction binding the contract method 0x9e165780.
//
// Solidity: function depositCallback(uint256 callType, address caller, address recipient, uint256 amount, uint64 depositNonce) returns()
func (_Bridge *BridgeTransactorSession) DepositCallback(callType *big.Int, caller common.Address, recipient common.Address, amount *big.Int, depositNonce uint64) (*types.Transaction, error) {
	return _Bridge.Contract.DepositCallback(&_Bridge.TransactOpts, callType, caller, recipient, amount, depositNonce)
}

// Execute is a paid mutator transaction binding the contract method 0xbe9e0cd7.
//
// Solidity: function execute(uint256 someNumber, uint256 callType, address recipient, uint256 amount, uint64 depositNonce) returns()
func (_Bridge *BridgeTransactor) Execute(opts *bind.TransactOpts, someNumber *big.Int, callType *big.Int, recipient common.Address, amount *big.Int, depositNonce uint64) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "execute", someNumber, callType, recipient, amount, depositNonce)
}

// Execute is a paid mutator transaction binding the contract method 0xbe9e0cd7.
//
// Solidity: function execute(uint256 someNumber, uint256 callType, address recipient, uint256 amount, uint64 depositNonce) returns()
func (_Bridge *BridgeSession) Execute(someNumber *big.Int, callType *big.Int, recipient common.Address, amount *big.Int, depositNonce uint64) (*types.Transaction, error) {
	return _Bridge.Contract.Execute(&_Bridge.TransactOpts, someNumber, callType, recipient, amount, depositNonce)
}

// Execute is a paid mutator transaction binding the contract method 0xbe9e0cd7.
//
// Solidity: function execute(uint256 someNumber, uint256 callType, address recipient, uint256 amount, uint64 depositNonce) returns()
func (_Bridge *BridgeTransactorSession) Execute(someNumber *big.Int, callType *big.Int, recipient common.Address, amount *big.Int, depositNonce uint64) (*types.Transaction, error) {
	return _Bridge.Contract.Execute(&_Bridge.TransactOpts, someNumber, callType, recipient, amount, depositNonce)
}

// GetBytesData is a paid mutator transaction binding the contract method 0xcdf6dd35.
//
// Solidity: function getBytesData(uint256 callType, address recipient, uint256 amount) returns(bytes)
func (_Bridge *BridgeTransactor) GetBytesData(opts *bind.TransactOpts, callType *big.Int, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "getBytesData", callType, recipient, amount)
}

// GetBytesData is a paid mutator transaction binding the contract method 0xcdf6dd35.
//
// Solidity: function getBytesData(uint256 callType, address recipient, uint256 amount) returns(bytes)
func (_Bridge *BridgeSession) GetBytesData(callType *big.Int, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.GetBytesData(&_Bridge.TransactOpts, callType, recipient, amount)
}

// GetBytesData is a paid mutator transaction binding the contract method 0xcdf6dd35.
//
// Solidity: function getBytesData(uint256 callType, address recipient, uint256 amount) returns(bytes)
func (_Bridge *BridgeTransactorSession) GetBytesData(callType *big.Int, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.GetBytesData(&_Bridge.TransactOpts, callType, recipient, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address toxAddress_, address usdtAddress_, address miningAddress_, address bridgeAddress_, address bridgeHandlerAddress_) returns()
func (_Bridge *BridgeTransactor) Initialize(opts *bind.TransactOpts, toxAddress_ common.Address, usdtAddress_ common.Address, miningAddress_ common.Address, bridgeAddress_ common.Address, bridgeHandlerAddress_ common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "initialize", toxAddress_, usdtAddress_, miningAddress_, bridgeAddress_, bridgeHandlerAddress_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address toxAddress_, address usdtAddress_, address miningAddress_, address bridgeAddress_, address bridgeHandlerAddress_) returns()
func (_Bridge *BridgeSession) Initialize(toxAddress_ common.Address, usdtAddress_ common.Address, miningAddress_ common.Address, bridgeAddress_ common.Address, bridgeHandlerAddress_ common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.Initialize(&_Bridge.TransactOpts, toxAddress_, usdtAddress_, miningAddress_, bridgeAddress_, bridgeHandlerAddress_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address toxAddress_, address usdtAddress_, address miningAddress_, address bridgeAddress_, address bridgeHandlerAddress_) returns()
func (_Bridge *BridgeTransactorSession) Initialize(toxAddress_ common.Address, usdtAddress_ common.Address, miningAddress_ common.Address, bridgeAddress_ common.Address, bridgeHandlerAddress_ common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.Initialize(&_Bridge.TransactOpts, toxAddress_, usdtAddress_, miningAddress_, bridgeAddress_, bridgeHandlerAddress_)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address account) returns()
func (_Bridge *BridgeTransactor) RemoveAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "removeAdmin", account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address account) returns()
func (_Bridge *BridgeSession) RemoveAdmin(account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RemoveAdmin(&_Bridge.TransactOpts, account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address account) returns()
func (_Bridge *BridgeTransactorSession) RemoveAdmin(account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RemoveAdmin(&_Bridge.TransactOpts, account)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_Bridge *BridgeTransactor) RenounceAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "renounceAdmin")
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_Bridge *BridgeSession) RenounceAdmin() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceAdmin(&_Bridge.TransactOpts)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_Bridge *BridgeTransactorSession) RenounceAdmin() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceAdmin(&_Bridge.TransactOpts)
}

// TopUpTest is a paid mutator transaction binding the contract method 0x6d2df8a4.
//
// Solidity: function topUpTest(address recipient, uint256 amount) returns()
func (_Bridge *BridgeTransactor) TopUpTest(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "topUpTest", recipient, amount)
}

// TopUpTest is a paid mutator transaction binding the contract method 0x6d2df8a4.
//
// Solidity: function topUpTest(address recipient, uint256 amount) returns()
func (_Bridge *BridgeSession) TopUpTest(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.TopUpTest(&_Bridge.TransactOpts, recipient, amount)
}

// TopUpTest is a paid mutator transaction binding the contract method 0x6d2df8a4.
//
// Solidity: function topUpTest(address recipient, uint256 amount) returns()
func (_Bridge *BridgeTransactorSession) TopUpTest(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.TopUpTest(&_Bridge.TransactOpts, recipient, amount)
}

// BridgeAdminAddedIterator is returned from FilterAdminAdded and is used to iterate over the raw logs and unpacked data for AdminAdded events raised by the Bridge contract.
type BridgeAdminAddedIterator struct {
	Event *BridgeAdminAdded // Event containing the contract specifics and raw log

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
func (it *BridgeAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeAdminAdded)
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
		it.Event = new(BridgeAdminAdded)
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
func (it *BridgeAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeAdminAdded represents a AdminAdded event raised by the Bridge contract.
type BridgeAdminAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminAdded is a free log retrieval operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address indexed account)
func (_Bridge *BridgeFilterer) FilterAdminAdded(opts *bind.FilterOpts, account []common.Address) (*BridgeAdminAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "AdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &BridgeAdminAddedIterator{contract: _Bridge.contract, event: "AdminAdded", logs: logs, sub: sub}, nil
}

// WatchAdminAdded is a free log subscription operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address indexed account)
func (_Bridge *BridgeFilterer) WatchAdminAdded(opts *bind.WatchOpts, sink chan<- *BridgeAdminAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "AdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeAdminAdded)
				if err := _Bridge.contract.UnpackLog(event, "AdminAdded", log); err != nil {
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

// ParseAdminAdded is a log parse operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address indexed account)
func (_Bridge *BridgeFilterer) ParseAdminAdded(log types.Log) (*BridgeAdminAdded, error) {
	event := new(BridgeAdminAdded)
	if err := _Bridge.contract.UnpackLog(event, "AdminAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the Bridge contract.
type BridgeAdminRemovedIterator struct {
	Event *BridgeAdminRemoved // Event containing the contract specifics and raw log

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
func (it *BridgeAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeAdminRemoved)
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
		it.Event = new(BridgeAdminRemoved)
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
func (it *BridgeAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeAdminRemoved represents a AdminRemoved event raised by the Bridge contract.
type BridgeAdminRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed account)
func (_Bridge *BridgeFilterer) FilterAdminRemoved(opts *bind.FilterOpts, account []common.Address) (*BridgeAdminRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "AdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &BridgeAdminRemovedIterator{contract: _Bridge.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed account)
func (_Bridge *BridgeFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *BridgeAdminRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "AdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeAdminRemoved)
				if err := _Bridge.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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

// ParseAdminRemoved is a log parse operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed account)
func (_Bridge *BridgeFilterer) ParseAdminRemoved(log types.Log) (*BridgeAdminRemoved, error) {
	event := new(BridgeAdminRemoved)
	if err := _Bridge.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeCallLogIterator is returned from FilterCallLog and is used to iterate over the raw logs and unpacked data for CallLog events raised by the Bridge contract.
type BridgeCallLogIterator struct {
	Event *BridgeCallLog // Event containing the contract specifics and raw log

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
func (it *BridgeCallLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeCallLog)
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
		it.Event = new(BridgeCallLog)
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
func (it *BridgeCallLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeCallLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeCallLog represents a CallLog event raised by the Bridge contract.
type BridgeCallLog struct {
	CallType     *big.Int
	Caller       common.Address
	Recipient    common.Address
	Amount       *big.Int
	DepositNonce uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCallLog is a free log retrieval operation binding the contract event 0x7f81ab8ec2567cfeb7dba357c192c8275473902b84ef5ad9319b4aff3cc3b2d0.
//
// Solidity: event CallLog(uint256 callType, address caller, address recipient, uint256 amount, uint64 depositNonce)
func (_Bridge *BridgeFilterer) FilterCallLog(opts *bind.FilterOpts) (*BridgeCallLogIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "CallLog")
	if err != nil {
		return nil, err
	}
	return &BridgeCallLogIterator{contract: _Bridge.contract, event: "CallLog", logs: logs, sub: sub}, nil
}

// WatchCallLog is a free log subscription operation binding the contract event 0x7f81ab8ec2567cfeb7dba357c192c8275473902b84ef5ad9319b4aff3cc3b2d0.
//
// Solidity: event CallLog(uint256 callType, address caller, address recipient, uint256 amount, uint64 depositNonce)
func (_Bridge *BridgeFilterer) WatchCallLog(opts *bind.WatchOpts, sink chan<- *BridgeCallLog) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "CallLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeCallLog)
				if err := _Bridge.contract.UnpackLog(event, "CallLog", log); err != nil {
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

// ParseCallLog is a log parse operation binding the contract event 0x7f81ab8ec2567cfeb7dba357c192c8275473902b84ef5ad9319b4aff3cc3b2d0.
//
// Solidity: event CallLog(uint256 callType, address caller, address recipient, uint256 amount, uint64 depositNonce)
func (_Bridge *BridgeFilterer) ParseCallLog(log types.Log) (*BridgeCallLog, error) {
	event := new(BridgeCallLog)
	if err := _Bridge.contract.UnpackLog(event, "CallLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeDepositDataIterator is returned from FilterDepositData and is used to iterate over the raw logs and unpacked data for DepositData events raised by the Bridge contract.
type BridgeDepositDataIterator struct {
	Event *BridgeDepositData // Event containing the contract specifics and raw log

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
func (it *BridgeDepositDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDepositData)
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
		it.Event = new(BridgeDepositData)
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
func (it *BridgeDepositDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDepositDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDepositData represents a DepositData event raised by the Bridge contract.
type BridgeDepositData struct {
	Data []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDepositData is a free log retrieval operation binding the contract event 0x6e28c8e27992732aa04aa17ea84c29f96a423c949ee6e8b6d7b447b40f78fe64.
//
// Solidity: event DepositData(bytes data)
func (_Bridge *BridgeFilterer) FilterDepositData(opts *bind.FilterOpts) (*BridgeDepositDataIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DepositData")
	if err != nil {
		return nil, err
	}
	return &BridgeDepositDataIterator{contract: _Bridge.contract, event: "DepositData", logs: logs, sub: sub}, nil
}

// WatchDepositData is a free log subscription operation binding the contract event 0x6e28c8e27992732aa04aa17ea84c29f96a423c949ee6e8b6d7b447b40f78fe64.
//
// Solidity: event DepositData(bytes data)
func (_Bridge *BridgeFilterer) WatchDepositData(opts *bind.WatchOpts, sink chan<- *BridgeDepositData) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DepositData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDepositData)
				if err := _Bridge.contract.UnpackLog(event, "DepositData", log); err != nil {
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

// ParseDepositData is a log parse operation binding the contract event 0x6e28c8e27992732aa04aa17ea84c29f96a423c949ee6e8b6d7b447b40f78fe64.
//
// Solidity: event DepositData(bytes data)
func (_Bridge *BridgeFilterer) ParseDepositData(log types.Log) (*BridgeDepositData, error) {
	event := new(BridgeDepositData)
	if err := _Bridge.contract.UnpackLog(event, "DepositData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeDepositDataLenIterator is returned from FilterDepositDataLen and is used to iterate over the raw logs and unpacked data for DepositDataLen events raised by the Bridge contract.
type BridgeDepositDataLenIterator struct {
	Event *BridgeDepositDataLen // Event containing the contract specifics and raw log

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
func (it *BridgeDepositDataLenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDepositDataLen)
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
		it.Event = new(BridgeDepositDataLen)
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
func (it *BridgeDepositDataLenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDepositDataLenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDepositDataLen represents a DepositDataLen event raised by the Bridge contract.
type BridgeDepositDataLen struct {
	DataLen *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositDataLen is a free log retrieval operation binding the contract event 0xcb18bfcf532f628f2ff0432f09761ad4f7ac185bef650cce1608c73dcb31a2ad.
//
// Solidity: event DepositDataLen(uint256 dataLen)
func (_Bridge *BridgeFilterer) FilterDepositDataLen(opts *bind.FilterOpts) (*BridgeDepositDataLenIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DepositDataLen")
	if err != nil {
		return nil, err
	}
	return &BridgeDepositDataLenIterator{contract: _Bridge.contract, event: "DepositDataLen", logs: logs, sub: sub}, nil
}

// WatchDepositDataLen is a free log subscription operation binding the contract event 0xcb18bfcf532f628f2ff0432f09761ad4f7ac185bef650cce1608c73dcb31a2ad.
//
// Solidity: event DepositDataLen(uint256 dataLen)
func (_Bridge *BridgeFilterer) WatchDepositDataLen(opts *bind.WatchOpts, sink chan<- *BridgeDepositDataLen) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DepositDataLen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDepositDataLen)
				if err := _Bridge.contract.UnpackLog(event, "DepositDataLen", log); err != nil {
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

// ParseDepositDataLen is a log parse operation binding the contract event 0xcb18bfcf532f628f2ff0432f09761ad4f7ac185bef650cce1608c73dcb31a2ad.
//
// Solidity: event DepositDataLen(uint256 dataLen)
func (_Bridge *BridgeFilterer) ParseDepositDataLen(log types.Log) (*BridgeDepositDataLen, error) {
	event := new(BridgeDepositDataLen)
	if err := _Bridge.contract.UnpackLog(event, "DepositDataLen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Bridge contract.
type BridgeInitializedIterator struct {
	Event *BridgeInitialized // Event containing the contract specifics and raw log

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
func (it *BridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeInitialized)
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
		it.Event = new(BridgeInitialized)
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
func (it *BridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeInitialized represents a Initialized event raised by the Bridge contract.
type BridgeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bridge *BridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*BridgeInitializedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BridgeInitializedIterator{contract: _Bridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bridge *BridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeInitialized)
				if err := _Bridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bridge *BridgeFilterer) ParseInitialized(log types.Log) (*BridgeInitialized, error) {
	event := new(BridgeInitialized)
	if err := _Bridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the Bridge contract.
type BridgeReceivedIterator struct {
	Event *BridgeReceived // Event containing the contract specifics and raw log

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
func (it *BridgeReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeReceived)
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
		it.Event = new(BridgeReceived)
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
func (it *BridgeReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeReceived represents a Received event raised by the Bridge contract.
type BridgeReceived struct {
	Caller common.Address
	Fee    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address caller, uint256 fee)
func (_Bridge *BridgeFilterer) FilterReceived(opts *bind.FilterOpts) (*BridgeReceivedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return &BridgeReceivedIterator{contract: _Bridge.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address caller, uint256 fee)
func (_Bridge *BridgeFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *BridgeReceived) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeReceived)
				if err := _Bridge.contract.UnpackLog(event, "Received", log); err != nil {
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
// Solidity: event Received(address caller, uint256 fee)
func (_Bridge *BridgeFilterer) ParseReceived(log types.Log) (*BridgeReceived, error) {
	event := new(BridgeReceived)
	if err := _Bridge.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeExecuteLogIterator is returned from FilterExecuteLog and is used to iterate over the raw logs and unpacked data for ExecuteLog events raised by the Bridge contract.
type BridgeExecuteLogIterator struct {
	Event *BridgeExecuteLog // Event containing the contract specifics and raw log

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
func (it *BridgeExecuteLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeExecuteLog)
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
		it.Event = new(BridgeExecuteLog)
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
func (it *BridgeExecuteLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeExecuteLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeExecuteLog represents a ExecuteLog event raised by the Bridge contract.
type BridgeExecuteLog struct {
	CallType     *big.Int
	Recipient    common.Address
	Amount       *big.Int
	DepositNonce uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterExecuteLog is a free log retrieval operation binding the contract event 0x026ce1aa14e0922ccf07c4bd7bcdbfb06b67dbf3a5dd824f37f7de947b73b67d.
//
// Solidity: event executeLog(uint256 callType, address recipient, uint256 amount, uint64 depositNonce)
func (_Bridge *BridgeFilterer) FilterExecuteLog(opts *bind.FilterOpts) (*BridgeExecuteLogIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "executeLog")
	if err != nil {
		return nil, err
	}
	return &BridgeExecuteLogIterator{contract: _Bridge.contract, event: "executeLog", logs: logs, sub: sub}, nil
}

// WatchExecuteLog is a free log subscription operation binding the contract event 0x026ce1aa14e0922ccf07c4bd7bcdbfb06b67dbf3a5dd824f37f7de947b73b67d.
//
// Solidity: event executeLog(uint256 callType, address recipient, uint256 amount, uint64 depositNonce)
func (_Bridge *BridgeFilterer) WatchExecuteLog(opts *bind.WatchOpts, sink chan<- *BridgeExecuteLog) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "executeLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeExecuteLog)
				if err := _Bridge.contract.UnpackLog(event, "executeLog", log); err != nil {
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

// ParseExecuteLog is a log parse operation binding the contract event 0x026ce1aa14e0922ccf07c4bd7bcdbfb06b67dbf3a5dd824f37f7de947b73b67d.
//
// Solidity: event executeLog(uint256 callType, address recipient, uint256 amount, uint64 depositNonce)
func (_Bridge *BridgeFilterer) ParseExecuteLog(log types.Log) (*BridgeExecuteLog, error) {
	event := new(BridgeExecuteLog)
	if err := _Bridge.contract.UnpackLog(event, "executeLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
