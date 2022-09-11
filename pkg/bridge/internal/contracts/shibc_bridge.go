// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ShibcBridgeMetaData contains all meta data concerning the ShibcBridge contract.
var ShibcBridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CloseSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"OpenSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"APPROVER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tx_hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"close\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"history\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"open\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ShibcBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use ShibcBridgeMetaData.ABI instead.
var ShibcBridgeABI = ShibcBridgeMetaData.ABI

// ShibcBridge is an auto generated Go binding around an Ethereum contract.
type ShibcBridge struct {
	ShibcBridgeCaller     // Read-only binding to the contract
	ShibcBridgeTransactor // Write-only binding to the contract
	ShibcBridgeFilterer   // Log filterer for contract events
}

// ShibcBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ShibcBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShibcBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ShibcBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShibcBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ShibcBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShibcBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ShibcBridgeSession struct {
	Contract     *ShibcBridge      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ShibcBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ShibcBridgeCallerSession struct {
	Contract *ShibcBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ShibcBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ShibcBridgeTransactorSession struct {
	Contract     *ShibcBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ShibcBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ShibcBridgeRaw struct {
	Contract *ShibcBridge // Generic contract binding to access the raw methods on
}

// ShibcBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ShibcBridgeCallerRaw struct {
	Contract *ShibcBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// ShibcBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ShibcBridgeTransactorRaw struct {
	Contract *ShibcBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewShibcBridge creates a new instance of ShibcBridge, bound to a specific deployed contract.
func NewShibcBridge(address common.Address, backend bind.ContractBackend) (*ShibcBridge, error) {
	contract, err := bindShibcBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ShibcBridge{ShibcBridgeCaller: ShibcBridgeCaller{contract: contract}, ShibcBridgeTransactor: ShibcBridgeTransactor{contract: contract}, ShibcBridgeFilterer: ShibcBridgeFilterer{contract: contract}}, nil
}

// NewShibcBridgeCaller creates a new read-only instance of ShibcBridge, bound to a specific deployed contract.
func NewShibcBridgeCaller(address common.Address, caller bind.ContractCaller) (*ShibcBridgeCaller, error) {
	contract, err := bindShibcBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShibcBridgeCaller{contract: contract}, nil
}

// NewShibcBridgeTransactor creates a new write-only instance of ShibcBridge, bound to a specific deployed contract.
func NewShibcBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*ShibcBridgeTransactor, error) {
	contract, err := bindShibcBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShibcBridgeTransactor{contract: contract}, nil
}

// NewShibcBridgeFilterer creates a new log filterer instance of ShibcBridge, bound to a specific deployed contract.
func NewShibcBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*ShibcBridgeFilterer, error) {
	contract, err := bindShibcBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShibcBridgeFilterer{contract: contract}, nil
}

// bindShibcBridge binds a generic wrapper to an already deployed contract.
func bindShibcBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ShibcBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShibcBridge *ShibcBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShibcBridge.Contract.ShibcBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShibcBridge *ShibcBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShibcBridge.Contract.ShibcBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShibcBridge *ShibcBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShibcBridge.Contract.ShibcBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShibcBridge *ShibcBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShibcBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShibcBridge *ShibcBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShibcBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShibcBridge *ShibcBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShibcBridge.Contract.contract.Transact(opts, method, params...)
}

// APPROVER is a free data retrieval call binding the contract method 0x3d111c7e.
//
// Solidity: function APPROVER() view returns(bytes32)
func (_ShibcBridge *ShibcBridgeCaller) APPROVER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ShibcBridge.contract.Call(opts, &out, "APPROVER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// APPROVER is a free data retrieval call binding the contract method 0x3d111c7e.
//
// Solidity: function APPROVER() view returns(bytes32)
func (_ShibcBridge *ShibcBridgeSession) APPROVER() ([32]byte, error) {
	return _ShibcBridge.Contract.APPROVER(&_ShibcBridge.CallOpts)
}

// APPROVER is a free data retrieval call binding the contract method 0x3d111c7e.
//
// Solidity: function APPROVER() view returns(bytes32)
func (_ShibcBridge *ShibcBridgeCallerSession) APPROVER() ([32]byte, error) {
	return _ShibcBridge.Contract.APPROVER(&_ShibcBridge.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ShibcBridge *ShibcBridgeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ShibcBridge.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ShibcBridge *ShibcBridgeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ShibcBridge.Contract.DEFAULTADMINROLE(&_ShibcBridge.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ShibcBridge *ShibcBridgeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ShibcBridge.Contract.DEFAULTADMINROLE(&_ShibcBridge.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ShibcBridge *ShibcBridgeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ShibcBridge.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ShibcBridge *ShibcBridgeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ShibcBridge.Contract.GetRoleAdmin(&_ShibcBridge.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ShibcBridge *ShibcBridgeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ShibcBridge.Contract.GetRoleAdmin(&_ShibcBridge.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ShibcBridge *ShibcBridgeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ShibcBridge.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ShibcBridge *ShibcBridgeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ShibcBridge.Contract.HasRole(&_ShibcBridge.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ShibcBridge *ShibcBridgeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ShibcBridge.Contract.HasRole(&_ShibcBridge.CallOpts, role, account)
}

// History is a free data retrieval call binding the contract method 0xed26bceb.
//
// Solidity: function history(bytes32 ) view returns(address)
func (_ShibcBridge *ShibcBridgeCaller) History(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ShibcBridge.contract.Call(opts, &out, "history", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// History is a free data retrieval call binding the contract method 0xed26bceb.
//
// Solidity: function history(bytes32 ) view returns(address)
func (_ShibcBridge *ShibcBridgeSession) History(arg0 [32]byte) (common.Address, error) {
	return _ShibcBridge.Contract.History(&_ShibcBridge.CallOpts, arg0)
}

// History is a free data retrieval call binding the contract method 0xed26bceb.
//
// Solidity: function history(bytes32 ) view returns(address)
func (_ShibcBridge *ShibcBridgeCallerSession) History(arg0 [32]byte) (common.Address, error) {
	return _ShibcBridge.Contract.History(&_ShibcBridge.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ShibcBridge *ShibcBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShibcBridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ShibcBridge *ShibcBridgeSession) Owner() (common.Address, error) {
	return _ShibcBridge.Contract.Owner(&_ShibcBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ShibcBridge *ShibcBridgeCallerSession) Owner() (common.Address, error) {
	return _ShibcBridge.Contract.Owner(&_ShibcBridge.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ShibcBridge *ShibcBridgeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ShibcBridge.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ShibcBridge *ShibcBridgeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ShibcBridge.Contract.SupportsInterface(&_ShibcBridge.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ShibcBridge *ShibcBridgeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ShibcBridge.Contract.SupportsInterface(&_ShibcBridge.CallOpts, interfaceId)
}

// Close is a paid mutator transaction binding the contract method 0xfe3cce58.
//
// Solidity: function close(bytes32 tx_hash, address recipient, uint256 amount) returns()
func (_ShibcBridge *ShibcBridgeTransactor) Close(opts *bind.TransactOpts, tx_hash [32]byte, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ShibcBridge.contract.Transact(opts, "close", tx_hash, recipient, amount)
}

// Close is a paid mutator transaction binding the contract method 0xfe3cce58.
//
// Solidity: function close(bytes32 tx_hash, address recipient, uint256 amount) returns()
func (_ShibcBridge *ShibcBridgeSession) Close(tx_hash [32]byte, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ShibcBridge.Contract.Close(&_ShibcBridge.TransactOpts, tx_hash, recipient, amount)
}

// Close is a paid mutator transaction binding the contract method 0xfe3cce58.
//
// Solidity: function close(bytes32 tx_hash, address recipient, uint256 amount) returns()
func (_ShibcBridge *ShibcBridgeTransactorSession) Close(tx_hash [32]byte, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ShibcBridge.Contract.Close(&_ShibcBridge.TransactOpts, tx_hash, recipient, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ShibcBridge *ShibcBridgeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ShibcBridge.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ShibcBridge *ShibcBridgeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ShibcBridge.Contract.GrantRole(&_ShibcBridge.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ShibcBridge *ShibcBridgeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ShibcBridge.Contract.GrantRole(&_ShibcBridge.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ShibcBridge *ShibcBridgeTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShibcBridge.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ShibcBridge *ShibcBridgeSession) Initialize() (*types.Transaction, error) {
	return _ShibcBridge.Contract.Initialize(&_ShibcBridge.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ShibcBridge *ShibcBridgeTransactorSession) Initialize() (*types.Transaction, error) {
	return _ShibcBridge.Contract.Initialize(&_ShibcBridge.TransactOpts)
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() payable returns()
func (_ShibcBridge *ShibcBridgeTransactor) Open(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShibcBridge.contract.Transact(opts, "open")
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() payable returns()
func (_ShibcBridge *ShibcBridgeSession) Open() (*types.Transaction, error) {
	return _ShibcBridge.Contract.Open(&_ShibcBridge.TransactOpts)
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() payable returns()
func (_ShibcBridge *ShibcBridgeTransactorSession) Open() (*types.Transaction, error) {
	return _ShibcBridge.Contract.Open(&_ShibcBridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ShibcBridge *ShibcBridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShibcBridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ShibcBridge *ShibcBridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _ShibcBridge.Contract.RenounceOwnership(&_ShibcBridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ShibcBridge *ShibcBridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ShibcBridge.Contract.RenounceOwnership(&_ShibcBridge.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ShibcBridge *ShibcBridgeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ShibcBridge.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ShibcBridge *ShibcBridgeSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ShibcBridge.Contract.RenounceRole(&_ShibcBridge.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ShibcBridge *ShibcBridgeTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ShibcBridge.Contract.RenounceRole(&_ShibcBridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ShibcBridge *ShibcBridgeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ShibcBridge.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ShibcBridge *ShibcBridgeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ShibcBridge.Contract.RevokeRole(&_ShibcBridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ShibcBridge *ShibcBridgeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ShibcBridge.Contract.RevokeRole(&_ShibcBridge.TransactOpts, role, account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ShibcBridge *ShibcBridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ShibcBridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ShibcBridge *ShibcBridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ShibcBridge.Contract.TransferOwnership(&_ShibcBridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ShibcBridge *ShibcBridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ShibcBridge.Contract.TransferOwnership(&_ShibcBridge.TransactOpts, newOwner)
}

// ShibcBridgeCloseSuccessIterator is returned from FilterCloseSuccess and is used to iterate over the raw logs and unpacked data for CloseSuccess events raised by the ShibcBridge contract.
type ShibcBridgeCloseSuccessIterator struct {
	Event *ShibcBridgeCloseSuccess // Event containing the contract specifics and raw log

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
func (it *ShibcBridgeCloseSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShibcBridgeCloseSuccess)
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
		it.Event = new(ShibcBridgeCloseSuccess)
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
func (it *ShibcBridgeCloseSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShibcBridgeCloseSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShibcBridgeCloseSuccess represents a CloseSuccess event raised by the ShibcBridge contract.
type ShibcBridgeCloseSuccess struct {
	Sender common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCloseSuccess is a free log retrieval operation binding the contract event 0x0fa9114275ab33b627cc27197ac4c1e022ab9f6ec20d305ce7537c030369daa9.
//
// Solidity: event CloseSuccess(address sender, address to, uint256 amount)
func (_ShibcBridge *ShibcBridgeFilterer) FilterCloseSuccess(opts *bind.FilterOpts) (*ShibcBridgeCloseSuccessIterator, error) {

	logs, sub, err := _ShibcBridge.contract.FilterLogs(opts, "CloseSuccess")
	if err != nil {
		return nil, err
	}
	return &ShibcBridgeCloseSuccessIterator{contract: _ShibcBridge.contract, event: "CloseSuccess", logs: logs, sub: sub}, nil
}

// WatchCloseSuccess is a free log subscription operation binding the contract event 0x0fa9114275ab33b627cc27197ac4c1e022ab9f6ec20d305ce7537c030369daa9.
//
// Solidity: event CloseSuccess(address sender, address to, uint256 amount)
func (_ShibcBridge *ShibcBridgeFilterer) WatchCloseSuccess(opts *bind.WatchOpts, sink chan<- *ShibcBridgeCloseSuccess) (event.Subscription, error) {

	logs, sub, err := _ShibcBridge.contract.WatchLogs(opts, "CloseSuccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShibcBridgeCloseSuccess)
				if err := _ShibcBridge.contract.UnpackLog(event, "CloseSuccess", log); err != nil {
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

// ParseCloseSuccess is a log parse operation binding the contract event 0x0fa9114275ab33b627cc27197ac4c1e022ab9f6ec20d305ce7537c030369daa9.
//
// Solidity: event CloseSuccess(address sender, address to, uint256 amount)
func (_ShibcBridge *ShibcBridgeFilterer) ParseCloseSuccess(log types.Log) (*ShibcBridgeCloseSuccess, error) {
	event := new(ShibcBridgeCloseSuccess)
	if err := _ShibcBridge.contract.UnpackLog(event, "CloseSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShibcBridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ShibcBridge contract.
type ShibcBridgeInitializedIterator struct {
	Event *ShibcBridgeInitialized // Event containing the contract specifics and raw log

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
func (it *ShibcBridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShibcBridgeInitialized)
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
		it.Event = new(ShibcBridgeInitialized)
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
func (it *ShibcBridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShibcBridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShibcBridgeInitialized represents a Initialized event raised by the ShibcBridge contract.
type ShibcBridgeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ShibcBridge *ShibcBridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*ShibcBridgeInitializedIterator, error) {

	logs, sub, err := _ShibcBridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ShibcBridgeInitializedIterator{contract: _ShibcBridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ShibcBridge *ShibcBridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ShibcBridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _ShibcBridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShibcBridgeInitialized)
				if err := _ShibcBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ShibcBridge *ShibcBridgeFilterer) ParseInitialized(log types.Log) (*ShibcBridgeInitialized, error) {
	event := new(ShibcBridgeInitialized)
	if err := _ShibcBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShibcBridgeOpenSuccessIterator is returned from FilterOpenSuccess and is used to iterate over the raw logs and unpacked data for OpenSuccess events raised by the ShibcBridge contract.
type ShibcBridgeOpenSuccessIterator struct {
	Event *ShibcBridgeOpenSuccess // Event containing the contract specifics and raw log

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
func (it *ShibcBridgeOpenSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShibcBridgeOpenSuccess)
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
		it.Event = new(ShibcBridgeOpenSuccess)
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
func (it *ShibcBridgeOpenSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShibcBridgeOpenSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShibcBridgeOpenSuccess represents a OpenSuccess event raised by the ShibcBridge contract.
type ShibcBridgeOpenSuccess struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOpenSuccess is a free log retrieval operation binding the contract event 0x3c8b0033b2ea5ab33dfdfc8a7f4a1db9393c3594eeb58c1a2f8757a609edc734.
//
// Solidity: event OpenSuccess(address sender, uint256 amount)
func (_ShibcBridge *ShibcBridgeFilterer) FilterOpenSuccess(opts *bind.FilterOpts) (*ShibcBridgeOpenSuccessIterator, error) {

	logs, sub, err := _ShibcBridge.contract.FilterLogs(opts, "OpenSuccess")
	if err != nil {
		return nil, err
	}
	return &ShibcBridgeOpenSuccessIterator{contract: _ShibcBridge.contract, event: "OpenSuccess", logs: logs, sub: sub}, nil
}

// WatchOpenSuccess is a free log subscription operation binding the contract event 0x3c8b0033b2ea5ab33dfdfc8a7f4a1db9393c3594eeb58c1a2f8757a609edc734.
//
// Solidity: event OpenSuccess(address sender, uint256 amount)
func (_ShibcBridge *ShibcBridgeFilterer) WatchOpenSuccess(opts *bind.WatchOpts, sink chan<- *ShibcBridgeOpenSuccess) (event.Subscription, error) {

	logs, sub, err := _ShibcBridge.contract.WatchLogs(opts, "OpenSuccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShibcBridgeOpenSuccess)
				if err := _ShibcBridge.contract.UnpackLog(event, "OpenSuccess", log); err != nil {
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

// ParseOpenSuccess is a log parse operation binding the contract event 0x3c8b0033b2ea5ab33dfdfc8a7f4a1db9393c3594eeb58c1a2f8757a609edc734.
//
// Solidity: event OpenSuccess(address sender, uint256 amount)
func (_ShibcBridge *ShibcBridgeFilterer) ParseOpenSuccess(log types.Log) (*ShibcBridgeOpenSuccess, error) {
	event := new(ShibcBridgeOpenSuccess)
	if err := _ShibcBridge.contract.UnpackLog(event, "OpenSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShibcBridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ShibcBridge contract.
type ShibcBridgeOwnershipTransferredIterator struct {
	Event *ShibcBridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ShibcBridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShibcBridgeOwnershipTransferred)
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
		it.Event = new(ShibcBridgeOwnershipTransferred)
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
func (it *ShibcBridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShibcBridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShibcBridgeOwnershipTransferred represents a OwnershipTransferred event raised by the ShibcBridge contract.
type ShibcBridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ShibcBridge *ShibcBridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ShibcBridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ShibcBridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ShibcBridgeOwnershipTransferredIterator{contract: _ShibcBridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ShibcBridge *ShibcBridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ShibcBridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ShibcBridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShibcBridgeOwnershipTransferred)
				if err := _ShibcBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ShibcBridge *ShibcBridgeFilterer) ParseOwnershipTransferred(log types.Log) (*ShibcBridgeOwnershipTransferred, error) {
	event := new(ShibcBridgeOwnershipTransferred)
	if err := _ShibcBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShibcBridgeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ShibcBridge contract.
type ShibcBridgeRoleAdminChangedIterator struct {
	Event *ShibcBridgeRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ShibcBridgeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShibcBridgeRoleAdminChanged)
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
		it.Event = new(ShibcBridgeRoleAdminChanged)
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
func (it *ShibcBridgeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShibcBridgeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShibcBridgeRoleAdminChanged represents a RoleAdminChanged event raised by the ShibcBridge contract.
type ShibcBridgeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ShibcBridge *ShibcBridgeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ShibcBridgeRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ShibcBridge.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ShibcBridgeRoleAdminChangedIterator{contract: _ShibcBridge.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ShibcBridge *ShibcBridgeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ShibcBridgeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ShibcBridge.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShibcBridgeRoleAdminChanged)
				if err := _ShibcBridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ShibcBridge *ShibcBridgeFilterer) ParseRoleAdminChanged(log types.Log) (*ShibcBridgeRoleAdminChanged, error) {
	event := new(ShibcBridgeRoleAdminChanged)
	if err := _ShibcBridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShibcBridgeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ShibcBridge contract.
type ShibcBridgeRoleGrantedIterator struct {
	Event *ShibcBridgeRoleGranted // Event containing the contract specifics and raw log

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
func (it *ShibcBridgeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShibcBridgeRoleGranted)
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
		it.Event = new(ShibcBridgeRoleGranted)
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
func (it *ShibcBridgeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShibcBridgeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShibcBridgeRoleGranted represents a RoleGranted event raised by the ShibcBridge contract.
type ShibcBridgeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ShibcBridge *ShibcBridgeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ShibcBridgeRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ShibcBridge.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ShibcBridgeRoleGrantedIterator{contract: _ShibcBridge.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ShibcBridge *ShibcBridgeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ShibcBridgeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ShibcBridge.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShibcBridgeRoleGranted)
				if err := _ShibcBridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ShibcBridge *ShibcBridgeFilterer) ParseRoleGranted(log types.Log) (*ShibcBridgeRoleGranted, error) {
	event := new(ShibcBridgeRoleGranted)
	if err := _ShibcBridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShibcBridgeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ShibcBridge contract.
type ShibcBridgeRoleRevokedIterator struct {
	Event *ShibcBridgeRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ShibcBridgeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShibcBridgeRoleRevoked)
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
		it.Event = new(ShibcBridgeRoleRevoked)
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
func (it *ShibcBridgeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShibcBridgeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShibcBridgeRoleRevoked represents a RoleRevoked event raised by the ShibcBridge contract.
type ShibcBridgeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ShibcBridge *ShibcBridgeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ShibcBridgeRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ShibcBridge.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ShibcBridgeRoleRevokedIterator{contract: _ShibcBridge.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ShibcBridge *ShibcBridgeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ShibcBridgeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ShibcBridge.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShibcBridgeRoleRevoked)
				if err := _ShibcBridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ShibcBridge *ShibcBridgeFilterer) ParseRoleRevoked(log types.Log) (*ShibcBridgeRoleRevoked, error) {
	event := new(ShibcBridgeRoleRevoked)
	if err := _ShibcBridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
