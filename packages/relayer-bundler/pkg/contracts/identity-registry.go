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
	_ = abi.ConvertType
)

// IdentityRegistryMetaData contains all meta data concerning the IdentityRegistry contract.
var IdentityRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activeStudentCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addStudent\",\"inputs\":[{\"name\":\"studentAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"studentId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isStudentActive\",\"inputs\":[{\"name\":\"studentAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isStudentValidForSnapshot\",\"inputs\":[{\"name\":\"student\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"snapshotBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStudentStatus\",\"inputs\":[{\"name\":\"studentAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"status\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"students\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"studentId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"registrationBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StudentAdded\",\"inputs\":[{\"name\":\"studentAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"studentId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StudentStatusChanged\",\"inputs\":[{\"name\":\"studentAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"isActive\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// IdentityRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use IdentityRegistryMetaData.ABI instead.
var IdentityRegistryABI = IdentityRegistryMetaData.ABI

// IdentityRegistry is an auto generated Go binding around an Ethereum contract.
type IdentityRegistry struct {
	IdentityRegistryCaller     // Read-only binding to the contract
	IdentityRegistryTransactor // Write-only binding to the contract
	IdentityRegistryFilterer   // Log filterer for contract events
}

// IdentityRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentityRegistrySession struct {
	Contract     *IdentityRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdentityRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityRegistryCallerSession struct {
	Contract *IdentityRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IdentityRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityRegistryTransactorSession struct {
	Contract     *IdentityRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IdentityRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityRegistryRaw struct {
	Contract *IdentityRegistry // Generic contract binding to access the raw methods on
}

// IdentityRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityRegistryCallerRaw struct {
	Contract *IdentityRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityRegistryTransactorRaw struct {
	Contract *IdentityRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentityRegistry creates a new instance of IdentityRegistry, bound to a specific deployed contract.
func NewIdentityRegistry(address common.Address, backend bind.ContractBackend) (*IdentityRegistry, error) {
	contract, err := bindIdentityRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistry{IdentityRegistryCaller: IdentityRegistryCaller{contract: contract}, IdentityRegistryTransactor: IdentityRegistryTransactor{contract: contract}, IdentityRegistryFilterer: IdentityRegistryFilterer{contract: contract}}, nil
}

// NewIdentityRegistryCaller creates a new read-only instance of IdentityRegistry, bound to a specific deployed contract.
func NewIdentityRegistryCaller(address common.Address, caller bind.ContractCaller) (*IdentityRegistryCaller, error) {
	contract, err := bindIdentityRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryCaller{contract: contract}, nil
}

// NewIdentityRegistryTransactor creates a new write-only instance of IdentityRegistry, bound to a specific deployed contract.
func NewIdentityRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityRegistryTransactor, error) {
	contract, err := bindIdentityRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryTransactor{contract: contract}, nil
}

// NewIdentityRegistryFilterer creates a new log filterer instance of IdentityRegistry, bound to a specific deployed contract.
func NewIdentityRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityRegistryFilterer, error) {
	contract, err := bindIdentityRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryFilterer{contract: contract}, nil
}

// bindIdentityRegistry binds a generic wrapper to an already deployed contract.
func bindIdentityRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IdentityRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityRegistry *IdentityRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdentityRegistry.Contract.IdentityRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityRegistry *IdentityRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.IdentityRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityRegistry *IdentityRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.IdentityRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityRegistry *IdentityRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdentityRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityRegistry *IdentityRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityRegistry *IdentityRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_IdentityRegistry *IdentityRegistryCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_IdentityRegistry *IdentityRegistrySession) ADMINROLE() ([32]byte, error) {
	return _IdentityRegistry.Contract.ADMINROLE(&_IdentityRegistry.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_IdentityRegistry *IdentityRegistryCallerSession) ADMINROLE() ([32]byte, error) {
	return _IdentityRegistry.Contract.ADMINROLE(&_IdentityRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_IdentityRegistry *IdentityRegistryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_IdentityRegistry *IdentityRegistrySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _IdentityRegistry.Contract.DEFAULTADMINROLE(&_IdentityRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_IdentityRegistry *IdentityRegistryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _IdentityRegistry.Contract.DEFAULTADMINROLE(&_IdentityRegistry.CallOpts)
}

// ActiveStudentCount is a free data retrieval call binding the contract method 0xa2c6d8e0.
//
// Solidity: function activeStudentCount() view returns(uint256)
func (_IdentityRegistry *IdentityRegistryCaller) ActiveStudentCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "activeStudentCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveStudentCount is a free data retrieval call binding the contract method 0xa2c6d8e0.
//
// Solidity: function activeStudentCount() view returns(uint256)
func (_IdentityRegistry *IdentityRegistrySession) ActiveStudentCount() (*big.Int, error) {
	return _IdentityRegistry.Contract.ActiveStudentCount(&_IdentityRegistry.CallOpts)
}

// ActiveStudentCount is a free data retrieval call binding the contract method 0xa2c6d8e0.
//
// Solidity: function activeStudentCount() view returns(uint256)
func (_IdentityRegistry *IdentityRegistryCallerSession) ActiveStudentCount() (*big.Int, error) {
	return _IdentityRegistry.Contract.ActiveStudentCount(&_IdentityRegistry.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IdentityRegistry *IdentityRegistryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IdentityRegistry *IdentityRegistrySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IdentityRegistry.Contract.GetRoleAdmin(&_IdentityRegistry.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IdentityRegistry *IdentityRegistryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IdentityRegistry.Contract.GetRoleAdmin(&_IdentityRegistry.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IdentityRegistry *IdentityRegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IdentityRegistry *IdentityRegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IdentityRegistry.Contract.HasRole(&_IdentityRegistry.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IdentityRegistry *IdentityRegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IdentityRegistry.Contract.HasRole(&_IdentityRegistry.CallOpts, role, account)
}

// IsStudentActive is a free data retrieval call binding the contract method 0x095f7b89.
//
// Solidity: function isStudentActive(address studentAddress) view returns(bool)
func (_IdentityRegistry *IdentityRegistryCaller) IsStudentActive(opts *bind.CallOpts, studentAddress common.Address) (bool, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "isStudentActive", studentAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStudentActive is a free data retrieval call binding the contract method 0x095f7b89.
//
// Solidity: function isStudentActive(address studentAddress) view returns(bool)
func (_IdentityRegistry *IdentityRegistrySession) IsStudentActive(studentAddress common.Address) (bool, error) {
	return _IdentityRegistry.Contract.IsStudentActive(&_IdentityRegistry.CallOpts, studentAddress)
}

// IsStudentActive is a free data retrieval call binding the contract method 0x095f7b89.
//
// Solidity: function isStudentActive(address studentAddress) view returns(bool)
func (_IdentityRegistry *IdentityRegistryCallerSession) IsStudentActive(studentAddress common.Address) (bool, error) {
	return _IdentityRegistry.Contract.IsStudentActive(&_IdentityRegistry.CallOpts, studentAddress)
}

// IsStudentValidForSnapshot is a free data retrieval call binding the contract method 0x3acbef55.
//
// Solidity: function isStudentValidForSnapshot(address student, uint256 snapshotBlock) view returns(bool)
func (_IdentityRegistry *IdentityRegistryCaller) IsStudentValidForSnapshot(opts *bind.CallOpts, student common.Address, snapshotBlock *big.Int) (bool, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "isStudentValidForSnapshot", student, snapshotBlock)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStudentValidForSnapshot is a free data retrieval call binding the contract method 0x3acbef55.
//
// Solidity: function isStudentValidForSnapshot(address student, uint256 snapshotBlock) view returns(bool)
func (_IdentityRegistry *IdentityRegistrySession) IsStudentValidForSnapshot(student common.Address, snapshotBlock *big.Int) (bool, error) {
	return _IdentityRegistry.Contract.IsStudentValidForSnapshot(&_IdentityRegistry.CallOpts, student, snapshotBlock)
}

// IsStudentValidForSnapshot is a free data retrieval call binding the contract method 0x3acbef55.
//
// Solidity: function isStudentValidForSnapshot(address student, uint256 snapshotBlock) view returns(bool)
func (_IdentityRegistry *IdentityRegistryCallerSession) IsStudentValidForSnapshot(student common.Address, snapshotBlock *big.Int) (bool, error) {
	return _IdentityRegistry.Contract.IsStudentValidForSnapshot(&_IdentityRegistry.CallOpts, student, snapshotBlock)
}

// Students is a free data retrieval call binding the contract method 0xa6c807a9.
//
// Solidity: function students(address ) view returns(bool isActive, string studentId, uint256 registrationBlock)
func (_IdentityRegistry *IdentityRegistryCaller) Students(opts *bind.CallOpts, arg0 common.Address) (struct {
	IsActive          bool
	StudentId         string
	RegistrationBlock *big.Int
}, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "students", arg0)

	outstruct := new(struct {
		IsActive          bool
		StudentId         string
		RegistrationBlock *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsActive = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.StudentId = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.RegistrationBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Students is a free data retrieval call binding the contract method 0xa6c807a9.
//
// Solidity: function students(address ) view returns(bool isActive, string studentId, uint256 registrationBlock)
func (_IdentityRegistry *IdentityRegistrySession) Students(arg0 common.Address) (struct {
	IsActive          bool
	StudentId         string
	RegistrationBlock *big.Int
}, error) {
	return _IdentityRegistry.Contract.Students(&_IdentityRegistry.CallOpts, arg0)
}

// Students is a free data retrieval call binding the contract method 0xa6c807a9.
//
// Solidity: function students(address ) view returns(bool isActive, string studentId, uint256 registrationBlock)
func (_IdentityRegistry *IdentityRegistryCallerSession) Students(arg0 common.Address) (struct {
	IsActive          bool
	StudentId         string
	RegistrationBlock *big.Int
}, error) {
	return _IdentityRegistry.Contract.Students(&_IdentityRegistry.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IdentityRegistry *IdentityRegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IdentityRegistry *IdentityRegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IdentityRegistry.Contract.SupportsInterface(&_IdentityRegistry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IdentityRegistry *IdentityRegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IdentityRegistry.Contract.SupportsInterface(&_IdentityRegistry.CallOpts, interfaceId)
}

// AddStudent is a paid mutator transaction binding the contract method 0x5e8b8091.
//
// Solidity: function addStudent(address studentAddress, string studentId) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) AddStudent(opts *bind.TransactOpts, studentAddress common.Address, studentId string) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "addStudent", studentAddress, studentId)
}

// AddStudent is a paid mutator transaction binding the contract method 0x5e8b8091.
//
// Solidity: function addStudent(address studentAddress, string studentId) returns()
func (_IdentityRegistry *IdentityRegistrySession) AddStudent(studentAddress common.Address, studentId string) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.AddStudent(&_IdentityRegistry.TransactOpts, studentAddress, studentId)
}

// AddStudent is a paid mutator transaction binding the contract method 0x5e8b8091.
//
// Solidity: function addStudent(address studentAddress, string studentId) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) AddStudent(studentAddress common.Address, studentId string) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.AddStudent(&_IdentityRegistry.TransactOpts, studentAddress, studentId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IdentityRegistry *IdentityRegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.GrantRole(&_IdentityRegistry.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.GrantRole(&_IdentityRegistry.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_IdentityRegistry *IdentityRegistrySession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.RenounceRole(&_IdentityRegistry.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.RenounceRole(&_IdentityRegistry.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IdentityRegistry *IdentityRegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.RevokeRole(&_IdentityRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.RevokeRole(&_IdentityRegistry.TransactOpts, role, account)
}

// SetStudentStatus is a paid mutator transaction binding the contract method 0xc95c2fc7.
//
// Solidity: function setStudentStatus(address studentAddress, bool status) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) SetStudentStatus(opts *bind.TransactOpts, studentAddress common.Address, status bool) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "setStudentStatus", studentAddress, status)
}

// SetStudentStatus is a paid mutator transaction binding the contract method 0xc95c2fc7.
//
// Solidity: function setStudentStatus(address studentAddress, bool status) returns()
func (_IdentityRegistry *IdentityRegistrySession) SetStudentStatus(studentAddress common.Address, status bool) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.SetStudentStatus(&_IdentityRegistry.TransactOpts, studentAddress, status)
}

// SetStudentStatus is a paid mutator transaction binding the contract method 0xc95c2fc7.
//
// Solidity: function setStudentStatus(address studentAddress, bool status) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) SetStudentStatus(studentAddress common.Address, status bool) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.SetStudentStatus(&_IdentityRegistry.TransactOpts, studentAddress, status)
}

// IdentityRegistryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the IdentityRegistry contract.
type IdentityRegistryRoleAdminChangedIterator struct {
	Event *IdentityRegistryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryRoleAdminChanged)
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
		it.Event = new(IdentityRegistryRoleAdminChanged)
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
func (it *IdentityRegistryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryRoleAdminChanged represents a RoleAdminChanged event raised by the IdentityRegistry contract.
type IdentityRegistryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IdentityRegistry *IdentityRegistryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*IdentityRegistryRoleAdminChangedIterator, error) {

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

	logs, sub, err := _IdentityRegistry.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryRoleAdminChangedIterator{contract: _IdentityRegistry.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IdentityRegistry *IdentityRegistryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *IdentityRegistryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _IdentityRegistry.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryRoleAdminChanged)
				if err := _IdentityRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_IdentityRegistry *IdentityRegistryFilterer) ParseRoleAdminChanged(log types.Log) (*IdentityRegistryRoleAdminChanged, error) {
	event := new(IdentityRegistryRoleAdminChanged)
	if err := _IdentityRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityRegistryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the IdentityRegistry contract.
type IdentityRegistryRoleGrantedIterator struct {
	Event *IdentityRegistryRoleGranted // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryRoleGranted)
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
		it.Event = new(IdentityRegistryRoleGranted)
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
func (it *IdentityRegistryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryRoleGranted represents a RoleGranted event raised by the IdentityRegistry contract.
type IdentityRegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IdentityRegistry *IdentityRegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IdentityRegistryRoleGrantedIterator, error) {

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

	logs, sub, err := _IdentityRegistry.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryRoleGrantedIterator{contract: _IdentityRegistry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IdentityRegistry *IdentityRegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *IdentityRegistryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _IdentityRegistry.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryRoleGranted)
				if err := _IdentityRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_IdentityRegistry *IdentityRegistryFilterer) ParseRoleGranted(log types.Log) (*IdentityRegistryRoleGranted, error) {
	event := new(IdentityRegistryRoleGranted)
	if err := _IdentityRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityRegistryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the IdentityRegistry contract.
type IdentityRegistryRoleRevokedIterator struct {
	Event *IdentityRegistryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryRoleRevoked)
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
		it.Event = new(IdentityRegistryRoleRevoked)
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
func (it *IdentityRegistryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryRoleRevoked represents a RoleRevoked event raised by the IdentityRegistry contract.
type IdentityRegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IdentityRegistry *IdentityRegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IdentityRegistryRoleRevokedIterator, error) {

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

	logs, sub, err := _IdentityRegistry.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryRoleRevokedIterator{contract: _IdentityRegistry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IdentityRegistry *IdentityRegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *IdentityRegistryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _IdentityRegistry.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryRoleRevoked)
				if err := _IdentityRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_IdentityRegistry *IdentityRegistryFilterer) ParseRoleRevoked(log types.Log) (*IdentityRegistryRoleRevoked, error) {
	event := new(IdentityRegistryRoleRevoked)
	if err := _IdentityRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityRegistryStudentAddedIterator is returned from FilterStudentAdded and is used to iterate over the raw logs and unpacked data for StudentAdded events raised by the IdentityRegistry contract.
type IdentityRegistryStudentAddedIterator struct {
	Event *IdentityRegistryStudentAdded // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryStudentAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryStudentAdded)
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
		it.Event = new(IdentityRegistryStudentAdded)
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
func (it *IdentityRegistryStudentAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryStudentAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryStudentAdded represents a StudentAdded event raised by the IdentityRegistry contract.
type IdentityRegistryStudentAdded struct {
	StudentAddress common.Address
	StudentId      string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterStudentAdded is a free log retrieval operation binding the contract event 0x546ed786c8923815c5efd04ecbc2b720211ff8b4998a8971c38897c3e75d1996.
//
// Solidity: event StudentAdded(address indexed studentAddress, string studentId)
func (_IdentityRegistry *IdentityRegistryFilterer) FilterStudentAdded(opts *bind.FilterOpts, studentAddress []common.Address) (*IdentityRegistryStudentAddedIterator, error) {

	var studentAddressRule []interface{}
	for _, studentAddressItem := range studentAddress {
		studentAddressRule = append(studentAddressRule, studentAddressItem)
	}

	logs, sub, err := _IdentityRegistry.contract.FilterLogs(opts, "StudentAdded", studentAddressRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryStudentAddedIterator{contract: _IdentityRegistry.contract, event: "StudentAdded", logs: logs, sub: sub}, nil
}

// WatchStudentAdded is a free log subscription operation binding the contract event 0x546ed786c8923815c5efd04ecbc2b720211ff8b4998a8971c38897c3e75d1996.
//
// Solidity: event StudentAdded(address indexed studentAddress, string studentId)
func (_IdentityRegistry *IdentityRegistryFilterer) WatchStudentAdded(opts *bind.WatchOpts, sink chan<- *IdentityRegistryStudentAdded, studentAddress []common.Address) (event.Subscription, error) {

	var studentAddressRule []interface{}
	for _, studentAddressItem := range studentAddress {
		studentAddressRule = append(studentAddressRule, studentAddressItem)
	}

	logs, sub, err := _IdentityRegistry.contract.WatchLogs(opts, "StudentAdded", studentAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryStudentAdded)
				if err := _IdentityRegistry.contract.UnpackLog(event, "StudentAdded", log); err != nil {
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

// ParseStudentAdded is a log parse operation binding the contract event 0x546ed786c8923815c5efd04ecbc2b720211ff8b4998a8971c38897c3e75d1996.
//
// Solidity: event StudentAdded(address indexed studentAddress, string studentId)
func (_IdentityRegistry *IdentityRegistryFilterer) ParseStudentAdded(log types.Log) (*IdentityRegistryStudentAdded, error) {
	event := new(IdentityRegistryStudentAdded)
	if err := _IdentityRegistry.contract.UnpackLog(event, "StudentAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityRegistryStudentStatusChangedIterator is returned from FilterStudentStatusChanged and is used to iterate over the raw logs and unpacked data for StudentStatusChanged events raised by the IdentityRegistry contract.
type IdentityRegistryStudentStatusChangedIterator struct {
	Event *IdentityRegistryStudentStatusChanged // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryStudentStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryStudentStatusChanged)
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
		it.Event = new(IdentityRegistryStudentStatusChanged)
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
func (it *IdentityRegistryStudentStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryStudentStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryStudentStatusChanged represents a StudentStatusChanged event raised by the IdentityRegistry contract.
type IdentityRegistryStudentStatusChanged struct {
	StudentAddress common.Address
	IsActive       bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterStudentStatusChanged is a free log retrieval operation binding the contract event 0x2c7c6b09ce2a51196bb33d1fc78177ab366b43fd142006433161c89ad982d1a5.
//
// Solidity: event StudentStatusChanged(address indexed studentAddress, bool isActive)
func (_IdentityRegistry *IdentityRegistryFilterer) FilterStudentStatusChanged(opts *bind.FilterOpts, studentAddress []common.Address) (*IdentityRegistryStudentStatusChangedIterator, error) {

	var studentAddressRule []interface{}
	for _, studentAddressItem := range studentAddress {
		studentAddressRule = append(studentAddressRule, studentAddressItem)
	}

	logs, sub, err := _IdentityRegistry.contract.FilterLogs(opts, "StudentStatusChanged", studentAddressRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryStudentStatusChangedIterator{contract: _IdentityRegistry.contract, event: "StudentStatusChanged", logs: logs, sub: sub}, nil
}

// WatchStudentStatusChanged is a free log subscription operation binding the contract event 0x2c7c6b09ce2a51196bb33d1fc78177ab366b43fd142006433161c89ad982d1a5.
//
// Solidity: event StudentStatusChanged(address indexed studentAddress, bool isActive)
func (_IdentityRegistry *IdentityRegistryFilterer) WatchStudentStatusChanged(opts *bind.WatchOpts, sink chan<- *IdentityRegistryStudentStatusChanged, studentAddress []common.Address) (event.Subscription, error) {

	var studentAddressRule []interface{}
	for _, studentAddressItem := range studentAddress {
		studentAddressRule = append(studentAddressRule, studentAddressItem)
	}

	logs, sub, err := _IdentityRegistry.contract.WatchLogs(opts, "StudentStatusChanged", studentAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryStudentStatusChanged)
				if err := _IdentityRegistry.contract.UnpackLog(event, "StudentStatusChanged", log); err != nil {
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

// ParseStudentStatusChanged is a log parse operation binding the contract event 0x2c7c6b09ce2a51196bb33d1fc78177ab366b43fd142006433161c89ad982d1a5.
//
// Solidity: event StudentStatusChanged(address indexed studentAddress, bool isActive)
func (_IdentityRegistry *IdentityRegistryFilterer) ParseStudentStatusChanged(log types.Log) (*IdentityRegistryStudentStatusChanged, error) {
	event := new(IdentityRegistryStudentStatusChanged)
	if err := _IdentityRegistry.contract.UnpackLog(event, "StudentStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
