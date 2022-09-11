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

// ETHBridgeMetaData contains all meta data concerning the ETHBridge contract.
var ETHBridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CloseSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"OpenSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"APPROVER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHIB\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tx_hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"close\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"history\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_shibToken\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"open\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ETHBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use ETHBridgeMetaData.ABI instead.
var ETHBridgeABI = ETHBridgeMetaData.ABI

// ETHBridge is an auto generated Go binding around an Ethereum contract.
type ETHBridge struct {
	ETHBridgeCaller     // Read-only binding to the contract
	ETHBridgeTransactor // Write-only binding to the contract
	ETHBridgeFilterer   // Log filterer for contract events
}

// ETHBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ETHBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ETHBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ETHBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ETHBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ETHBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ETHBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ETHBridgeSession struct {
	Contract     *ETHBridge        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ETHBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ETHBridgeCallerSession struct {
	Contract *ETHBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ETHBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ETHBridgeTransactorSession struct {
	Contract     *ETHBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ETHBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ETHBridgeRaw struct {
	Contract *ETHBridge // Generic contract binding to access the raw methods on
}

// ETHBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ETHBridgeCallerRaw struct {
	Contract *ETHBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// ETHBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ETHBridgeTransactorRaw struct {
	Contract *ETHBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewETHBridge creates a new instance of ETHBridge, bound to a specific deployed contract.
func NewETHBridge(address common.Address, backend bind.ContractBackend) (*ETHBridge, error) {
	contract, err := bindETHBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ETHBridge{ETHBridgeCaller: ETHBridgeCaller{contract: contract}, ETHBridgeTransactor: ETHBridgeTransactor{contract: contract}, ETHBridgeFilterer: ETHBridgeFilterer{contract: contract}}, nil
}

// NewETHBridgeCaller creates a new read-only instance of ETHBridge, bound to a specific deployed contract.
func NewETHBridgeCaller(address common.Address, caller bind.ContractCaller) (*ETHBridgeCaller, error) {
	contract, err := bindETHBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ETHBridgeCaller{contract: contract}, nil
}

// NewETHBridgeTransactor creates a new write-only instance of ETHBridge, bound to a specific deployed contract.
func NewETHBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*ETHBridgeTransactor, error) {
	contract, err := bindETHBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ETHBridgeTransactor{contract: contract}, nil
}

// NewETHBridgeFilterer creates a new log filterer instance of ETHBridge, bound to a specific deployed contract.
func NewETHBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*ETHBridgeFilterer, error) {
	contract, err := bindETHBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ETHBridgeFilterer{contract: contract}, nil
}

// bindETHBridge binds a generic wrapper to an already deployed contract.
func bindETHBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ETHBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ETHBridge *ETHBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ETHBridge.Contract.ETHBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ETHBridge *ETHBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHBridge.Contract.ETHBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ETHBridge *ETHBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ETHBridge.Contract.ETHBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ETHBridge *ETHBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ETHBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ETHBridge *ETHBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ETHBridge *ETHBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ETHBridge.Contract.contract.Transact(opts, method, params...)
}

// APPROVER is a free data retrieval call binding the contract method 0x3d111c7e.
//
// Solidity: function APPROVER() view returns(bytes32)
func (_ETHBridge *ETHBridgeCaller) APPROVER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ETHBridge.contract.Call(opts, &out, "APPROVER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// APPROVER is a free data retrieval call binding the contract method 0x3d111c7e.
//
// Solidity: function APPROVER() view returns(bytes32)
func (_ETHBridge *ETHBridgeSession) APPROVER() ([32]byte, error) {
	return _ETHBridge.Contract.APPROVER(&_ETHBridge.CallOpts)
}

// APPROVER is a free data retrieval call binding the contract method 0x3d111c7e.
//
// Solidity: function APPROVER() view returns(bytes32)
func (_ETHBridge *ETHBridgeCallerSession) APPROVER() ([32]byte, error) {
	return _ETHBridge.Contract.APPROVER(&_ETHBridge.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ETHBridge *ETHBridgeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ETHBridge.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ETHBridge *ETHBridgeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ETHBridge.Contract.DEFAULTADMINROLE(&_ETHBridge.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ETHBridge *ETHBridgeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ETHBridge.Contract.DEFAULTADMINROLE(&_ETHBridge.CallOpts)
}

// SHIB is a free data retrieval call binding the contract method 0xe24b85e7.
//
// Solidity: function SHIB() view returns(address)
func (_ETHBridge *ETHBridgeCaller) SHIB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ETHBridge.contract.Call(opts, &out, "SHIB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SHIB is a free data retrieval call binding the contract method 0xe24b85e7.
//
// Solidity: function SHIB() view returns(address)
func (_ETHBridge *ETHBridgeSession) SHIB() (common.Address, error) {
	return _ETHBridge.Contract.SHIB(&_ETHBridge.CallOpts)
}

// SHIB is a free data retrieval call binding the contract method 0xe24b85e7.
//
// Solidity: function SHIB() view returns(address)
func (_ETHBridge *ETHBridgeCallerSession) SHIB() (common.Address, error) {
	return _ETHBridge.Contract.SHIB(&_ETHBridge.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ETHBridge *ETHBridgeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ETHBridge.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ETHBridge *ETHBridgeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ETHBridge.Contract.GetRoleAdmin(&_ETHBridge.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ETHBridge *ETHBridgeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ETHBridge.Contract.GetRoleAdmin(&_ETHBridge.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ETHBridge *ETHBridgeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ETHBridge.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ETHBridge *ETHBridgeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ETHBridge.Contract.HasRole(&_ETHBridge.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ETHBridge *ETHBridgeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ETHBridge.Contract.HasRole(&_ETHBridge.CallOpts, role, account)
}

// History is a free data retrieval call binding the contract method 0xed26bceb.
//
// Solidity: function history(bytes32 ) view returns(address)
func (_ETHBridge *ETHBridgeCaller) History(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ETHBridge.contract.Call(opts, &out, "history", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// History is a free data retrieval call binding the contract method 0xed26bceb.
//
// Solidity: function history(bytes32 ) view returns(address)
func (_ETHBridge *ETHBridgeSession) History(arg0 [32]byte) (common.Address, error) {
	return _ETHBridge.Contract.History(&_ETHBridge.CallOpts, arg0)
}

// History is a free data retrieval call binding the contract method 0xed26bceb.
//
// Solidity: function history(bytes32 ) view returns(address)
func (_ETHBridge *ETHBridgeCallerSession) History(arg0 [32]byte) (common.Address, error) {
	return _ETHBridge.Contract.History(&_ETHBridge.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ETHBridge *ETHBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ETHBridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ETHBridge *ETHBridgeSession) Owner() (common.Address, error) {
	return _ETHBridge.Contract.Owner(&_ETHBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ETHBridge *ETHBridgeCallerSession) Owner() (common.Address, error) {
	return _ETHBridge.Contract.Owner(&_ETHBridge.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ETHBridge *ETHBridgeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ETHBridge.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ETHBridge *ETHBridgeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ETHBridge.Contract.SupportsInterface(&_ETHBridge.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ETHBridge *ETHBridgeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ETHBridge.Contract.SupportsInterface(&_ETHBridge.CallOpts, interfaceId)
}

// Close is a paid mutator transaction binding the contract method 0xfe3cce58.
//
// Solidity: function close(bytes32 tx_hash, address to, uint256 amount) returns()
func (_ETHBridge *ETHBridgeTransactor) Close(opts *bind.TransactOpts, tx_hash [32]byte, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ETHBridge.contract.Transact(opts, "close", tx_hash, to, amount)
}

// Close is a paid mutator transaction binding the contract method 0xfe3cce58.
//
// Solidity: function close(bytes32 tx_hash, address to, uint256 amount) returns()
func (_ETHBridge *ETHBridgeSession) Close(tx_hash [32]byte, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ETHBridge.Contract.Close(&_ETHBridge.TransactOpts, tx_hash, to, amount)
}

// Close is a paid mutator transaction binding the contract method 0xfe3cce58.
//
// Solidity: function close(bytes32 tx_hash, address to, uint256 amount) returns()
func (_ETHBridge *ETHBridgeTransactorSession) Close(tx_hash [32]byte, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ETHBridge.Contract.Close(&_ETHBridge.TransactOpts, tx_hash, to, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ETHBridge *ETHBridgeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ETHBridge.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ETHBridge *ETHBridgeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ETHBridge.Contract.GrantRole(&_ETHBridge.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ETHBridge *ETHBridgeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ETHBridge.Contract.GrantRole(&_ETHBridge.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _shibToken) returns()
func (_ETHBridge *ETHBridgeTransactor) Initialize(opts *bind.TransactOpts, _shibToken common.Address) (*types.Transaction, error) {
	return _ETHBridge.contract.Transact(opts, "initialize", _shibToken)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _shibToken) returns()
func (_ETHBridge *ETHBridgeSession) Initialize(_shibToken common.Address) (*types.Transaction, error) {
	return _ETHBridge.Contract.Initialize(&_ETHBridge.TransactOpts, _shibToken)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _shibToken) returns()
func (_ETHBridge *ETHBridgeTransactorSession) Initialize(_shibToken common.Address) (*types.Transaction, error) {
	return _ETHBridge.Contract.Initialize(&_ETHBridge.TransactOpts, _shibToken)
}

// Open is a paid mutator transaction binding the contract method 0x690e7c09.
//
// Solidity: function open(uint256 amount) returns()
func (_ETHBridge *ETHBridgeTransactor) Open(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ETHBridge.contract.Transact(opts, "open", amount)
}

// Open is a paid mutator transaction binding the contract method 0x690e7c09.
//
// Solidity: function open(uint256 amount) returns()
func (_ETHBridge *ETHBridgeSession) Open(amount *big.Int) (*types.Transaction, error) {
	return _ETHBridge.Contract.Open(&_ETHBridge.TransactOpts, amount)
}

// Open is a paid mutator transaction binding the contract method 0x690e7c09.
//
// Solidity: function open(uint256 amount) returns()
func (_ETHBridge *ETHBridgeTransactorSession) Open(amount *big.Int) (*types.Transaction, error) {
	return _ETHBridge.Contract.Open(&_ETHBridge.TransactOpts, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ETHBridge *ETHBridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHBridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ETHBridge *ETHBridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _ETHBridge.Contract.RenounceOwnership(&_ETHBridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ETHBridge *ETHBridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ETHBridge.Contract.RenounceOwnership(&_ETHBridge.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ETHBridge *ETHBridgeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ETHBridge.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ETHBridge *ETHBridgeSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ETHBridge.Contract.RenounceRole(&_ETHBridge.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ETHBridge *ETHBridgeTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ETHBridge.Contract.RenounceRole(&_ETHBridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ETHBridge *ETHBridgeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ETHBridge.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ETHBridge *ETHBridgeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ETHBridge.Contract.RevokeRole(&_ETHBridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ETHBridge *ETHBridgeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ETHBridge.Contract.RevokeRole(&_ETHBridge.TransactOpts, role, account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ETHBridge *ETHBridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ETHBridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ETHBridge *ETHBridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ETHBridge.Contract.TransferOwnership(&_ETHBridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ETHBridge *ETHBridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ETHBridge.Contract.TransferOwnership(&_ETHBridge.TransactOpts, newOwner)
}

// ETHBridgeCloseSuccessIterator is returned from FilterCloseSuccess and is used to iterate over the raw logs and unpacked data for CloseSuccess events raised by the ETHBridge contract.
type ETHBridgeCloseSuccessIterator struct {
	Event *ETHBridgeCloseSuccess // Event containing the contract specifics and raw log

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
func (it *ETHBridgeCloseSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHBridgeCloseSuccess)
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
		it.Event = new(ETHBridgeCloseSuccess)
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
func (it *ETHBridgeCloseSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHBridgeCloseSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHBridgeCloseSuccess represents a CloseSuccess event raised by the ETHBridge contract.
type ETHBridgeCloseSuccess struct {
	Sender common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCloseSuccess is a free log retrieval operation binding the contract event 0x0fa9114275ab33b627cc27197ac4c1e022ab9f6ec20d305ce7537c030369daa9.
//
// Solidity: event CloseSuccess(address sender, address to, uint256 amount)
func (_ETHBridge *ETHBridgeFilterer) FilterCloseSuccess(opts *bind.FilterOpts) (*ETHBridgeCloseSuccessIterator, error) {

	logs, sub, err := _ETHBridge.contract.FilterLogs(opts, "CloseSuccess")
	if err != nil {
		return nil, err
	}
	return &ETHBridgeCloseSuccessIterator{contract: _ETHBridge.contract, event: "CloseSuccess", logs: logs, sub: sub}, nil
}

// WatchCloseSuccess is a free log subscription operation binding the contract event 0x0fa9114275ab33b627cc27197ac4c1e022ab9f6ec20d305ce7537c030369daa9.
//
// Solidity: event CloseSuccess(address sender, address to, uint256 amount)
func (_ETHBridge *ETHBridgeFilterer) WatchCloseSuccess(opts *bind.WatchOpts, sink chan<- *ETHBridgeCloseSuccess) (event.Subscription, error) {

	logs, sub, err := _ETHBridge.contract.WatchLogs(opts, "CloseSuccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHBridgeCloseSuccess)
				if err := _ETHBridge.contract.UnpackLog(event, "CloseSuccess", log); err != nil {
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
func (_ETHBridge *ETHBridgeFilterer) ParseCloseSuccess(log types.Log) (*ETHBridgeCloseSuccess, error) {
	event := new(ETHBridgeCloseSuccess)
	if err := _ETHBridge.contract.UnpackLog(event, "CloseSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHBridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ETHBridge contract.
type ETHBridgeInitializedIterator struct {
	Event *ETHBridgeInitialized // Event containing the contract specifics and raw log

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
func (it *ETHBridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHBridgeInitialized)
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
		it.Event = new(ETHBridgeInitialized)
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
func (it *ETHBridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHBridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHBridgeInitialized represents a Initialized event raised by the ETHBridge contract.
type ETHBridgeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ETHBridge *ETHBridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*ETHBridgeInitializedIterator, error) {

	logs, sub, err := _ETHBridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ETHBridgeInitializedIterator{contract: _ETHBridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ETHBridge *ETHBridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ETHBridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _ETHBridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHBridgeInitialized)
				if err := _ETHBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ETHBridge *ETHBridgeFilterer) ParseInitialized(log types.Log) (*ETHBridgeInitialized, error) {
	event := new(ETHBridgeInitialized)
	if err := _ETHBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHBridgeOpenSuccessIterator is returned from FilterOpenSuccess and is used to iterate over the raw logs and unpacked data for OpenSuccess events raised by the ETHBridge contract.
type ETHBridgeOpenSuccessIterator struct {
	Event *ETHBridgeOpenSuccess // Event containing the contract specifics and raw log

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
func (it *ETHBridgeOpenSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHBridgeOpenSuccess)
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
		it.Event = new(ETHBridgeOpenSuccess)
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
func (it *ETHBridgeOpenSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHBridgeOpenSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHBridgeOpenSuccess represents a OpenSuccess event raised by the ETHBridge contract.
type ETHBridgeOpenSuccess struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOpenSuccess is a free log retrieval operation binding the contract event 0x3c8b0033b2ea5ab33dfdfc8a7f4a1db9393c3594eeb58c1a2f8757a609edc734.
//
// Solidity: event OpenSuccess(address sender, uint256 amount)
func (_ETHBridge *ETHBridgeFilterer) FilterOpenSuccess(opts *bind.FilterOpts) (*ETHBridgeOpenSuccessIterator, error) {

	logs, sub, err := _ETHBridge.contract.FilterLogs(opts, "OpenSuccess")
	if err != nil {
		return nil, err
	}
	return &ETHBridgeOpenSuccessIterator{contract: _ETHBridge.contract, event: "OpenSuccess", logs: logs, sub: sub}, nil
}

// WatchOpenSuccess is a free log subscription operation binding the contract event 0x3c8b0033b2ea5ab33dfdfc8a7f4a1db9393c3594eeb58c1a2f8757a609edc734.
//
// Solidity: event OpenSuccess(address sender, uint256 amount)
func (_ETHBridge *ETHBridgeFilterer) WatchOpenSuccess(opts *bind.WatchOpts, sink chan<- *ETHBridgeOpenSuccess) (event.Subscription, error) {

	logs, sub, err := _ETHBridge.contract.WatchLogs(opts, "OpenSuccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHBridgeOpenSuccess)
				if err := _ETHBridge.contract.UnpackLog(event, "OpenSuccess", log); err != nil {
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
func (_ETHBridge *ETHBridgeFilterer) ParseOpenSuccess(log types.Log) (*ETHBridgeOpenSuccess, error) {
	event := new(ETHBridgeOpenSuccess)
	if err := _ETHBridge.contract.UnpackLog(event, "OpenSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHBridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ETHBridge contract.
type ETHBridgeOwnershipTransferredIterator struct {
	Event *ETHBridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ETHBridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHBridgeOwnershipTransferred)
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
		it.Event = new(ETHBridgeOwnershipTransferred)
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
func (it *ETHBridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHBridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHBridgeOwnershipTransferred represents a OwnershipTransferred event raised by the ETHBridge contract.
type ETHBridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ETHBridge *ETHBridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ETHBridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ETHBridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ETHBridgeOwnershipTransferredIterator{contract: _ETHBridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ETHBridge *ETHBridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ETHBridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ETHBridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHBridgeOwnershipTransferred)
				if err := _ETHBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ETHBridge *ETHBridgeFilterer) ParseOwnershipTransferred(log types.Log) (*ETHBridgeOwnershipTransferred, error) {
	event := new(ETHBridgeOwnershipTransferred)
	if err := _ETHBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHBridgeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ETHBridge contract.
type ETHBridgeRoleAdminChangedIterator struct {
	Event *ETHBridgeRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ETHBridgeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHBridgeRoleAdminChanged)
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
		it.Event = new(ETHBridgeRoleAdminChanged)
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
func (it *ETHBridgeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHBridgeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHBridgeRoleAdminChanged represents a RoleAdminChanged event raised by the ETHBridge contract.
type ETHBridgeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ETHBridge *ETHBridgeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ETHBridgeRoleAdminChangedIterator, error) {

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

	logs, sub, err := _ETHBridge.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ETHBridgeRoleAdminChangedIterator{contract: _ETHBridge.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ETHBridge *ETHBridgeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ETHBridgeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ETHBridge.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHBridgeRoleAdminChanged)
				if err := _ETHBridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_ETHBridge *ETHBridgeFilterer) ParseRoleAdminChanged(log types.Log) (*ETHBridgeRoleAdminChanged, error) {
	event := new(ETHBridgeRoleAdminChanged)
	if err := _ETHBridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHBridgeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ETHBridge contract.
type ETHBridgeRoleGrantedIterator struct {
	Event *ETHBridgeRoleGranted // Event containing the contract specifics and raw log

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
func (it *ETHBridgeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHBridgeRoleGranted)
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
		it.Event = new(ETHBridgeRoleGranted)
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
func (it *ETHBridgeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHBridgeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHBridgeRoleGranted represents a RoleGranted event raised by the ETHBridge contract.
type ETHBridgeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ETHBridge *ETHBridgeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ETHBridgeRoleGrantedIterator, error) {

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

	logs, sub, err := _ETHBridge.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ETHBridgeRoleGrantedIterator{contract: _ETHBridge.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ETHBridge *ETHBridgeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ETHBridgeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ETHBridge.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHBridgeRoleGranted)
				if err := _ETHBridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ETHBridge *ETHBridgeFilterer) ParseRoleGranted(log types.Log) (*ETHBridgeRoleGranted, error) {
	event := new(ETHBridgeRoleGranted)
	if err := _ETHBridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHBridgeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ETHBridge contract.
type ETHBridgeRoleRevokedIterator struct {
	Event *ETHBridgeRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ETHBridgeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHBridgeRoleRevoked)
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
		it.Event = new(ETHBridgeRoleRevoked)
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
func (it *ETHBridgeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHBridgeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHBridgeRoleRevoked represents a RoleRevoked event raised by the ETHBridge contract.
type ETHBridgeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ETHBridge *ETHBridgeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ETHBridgeRoleRevokedIterator, error) {

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

	logs, sub, err := _ETHBridge.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ETHBridgeRoleRevokedIterator{contract: _ETHBridge.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ETHBridge *ETHBridgeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ETHBridgeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ETHBridge.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHBridgeRoleRevoked)
				if err := _ETHBridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ETHBridge *ETHBridgeFilterer) ParseRoleRevoked(log types.Log) (*ETHBridgeRoleRevoked, error) {
	event := new(ETHBridgeRoleRevoked)
	if err := _ETHBridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
