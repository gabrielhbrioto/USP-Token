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

// PackedUserOperation is an auto generated low-level Go binding around an user-defined struct.
type PackedUserOperation struct {
	Sender             common.Address
	Nonce              *big.Int
	InitCode           []byte
	CallData           []byte
	AccountGasLimits   [32]byte
	PreVerificationGas *big.Int
	GasFees            [32]byte
	PaymasterAndData   []byte
	Signature          []byte
}

// USPPaymasterMetaData contains all meta data concerning the USPPaymaster contract.
var USPPaymasterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_entryPoint\",\"type\":\"address\",\"internalType\":\"contractIEntryPoint\"},{\"name\":\"_identityRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"entryPoint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEntryPoint\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gasSpentByStudent\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"identityRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIIdentityRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"makeDeposit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"maxGasPerStudent\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"postOp\",\"inputs\":[{\"name\":\"mode\",\"type\":\"uint8\",\"internalType\":\"enumIPaymaster.PostOpMode\"},{\"name\":\"context\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"actualGasCost\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxGasPerStudent\",\"inputs\":[{\"name\":\"_newLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTargetWhitelist\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"status\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validatePaymasterUserOp\",\"inputs\":[{\"name\":\"userOp\",\"type\":\"tuple\",\"internalType\":\"structPackedUserOperation\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"accountGasLimits\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"preVerificationGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasFees\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"paymasterAndData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"maxCost\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"context\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"validationData\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"whitelistedTargets\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawTo\",\"inputs\":[{\"name\":\"withdrawAddress\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"GasCharged\",\"inputs\":[{\"name\":\"student\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"actualGasCost\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"mode\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIPaymaster.PostOpMode\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GasLimitUpdated\",\"inputs\":[{\"name\":\"newLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// USPPaymasterABI is the input ABI used to generate the binding from.
// Deprecated: Use USPPaymasterMetaData.ABI instead.
var USPPaymasterABI = USPPaymasterMetaData.ABI

// USPPaymaster is an auto generated Go binding around an Ethereum contract.
type USPPaymaster struct {
	USPPaymasterCaller     // Read-only binding to the contract
	USPPaymasterTransactor // Write-only binding to the contract
	USPPaymasterFilterer   // Log filterer for contract events
}

// USPPaymasterCaller is an auto generated read-only Go binding around an Ethereum contract.
type USPPaymasterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USPPaymasterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type USPPaymasterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USPPaymasterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type USPPaymasterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USPPaymasterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type USPPaymasterSession struct {
	Contract     *USPPaymaster     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// USPPaymasterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type USPPaymasterCallerSession struct {
	Contract *USPPaymasterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// USPPaymasterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type USPPaymasterTransactorSession struct {
	Contract     *USPPaymasterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// USPPaymasterRaw is an auto generated low-level Go binding around an Ethereum contract.
type USPPaymasterRaw struct {
	Contract *USPPaymaster // Generic contract binding to access the raw methods on
}

// USPPaymasterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type USPPaymasterCallerRaw struct {
	Contract *USPPaymasterCaller // Generic read-only contract binding to access the raw methods on
}

// USPPaymasterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type USPPaymasterTransactorRaw struct {
	Contract *USPPaymasterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUSPPaymaster creates a new instance of USPPaymaster, bound to a specific deployed contract.
func NewUSPPaymaster(address common.Address, backend bind.ContractBackend) (*USPPaymaster, error) {
	contract, err := bindUSPPaymaster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &USPPaymaster{USPPaymasterCaller: USPPaymasterCaller{contract: contract}, USPPaymasterTransactor: USPPaymasterTransactor{contract: contract}, USPPaymasterFilterer: USPPaymasterFilterer{contract: contract}}, nil
}

// NewUSPPaymasterCaller creates a new read-only instance of USPPaymaster, bound to a specific deployed contract.
func NewUSPPaymasterCaller(address common.Address, caller bind.ContractCaller) (*USPPaymasterCaller, error) {
	contract, err := bindUSPPaymaster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &USPPaymasterCaller{contract: contract}, nil
}

// NewUSPPaymasterTransactor creates a new write-only instance of USPPaymaster, bound to a specific deployed contract.
func NewUSPPaymasterTransactor(address common.Address, transactor bind.ContractTransactor) (*USPPaymasterTransactor, error) {
	contract, err := bindUSPPaymaster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &USPPaymasterTransactor{contract: contract}, nil
}

// NewUSPPaymasterFilterer creates a new log filterer instance of USPPaymaster, bound to a specific deployed contract.
func NewUSPPaymasterFilterer(address common.Address, filterer bind.ContractFilterer) (*USPPaymasterFilterer, error) {
	contract, err := bindUSPPaymaster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &USPPaymasterFilterer{contract: contract}, nil
}

// bindUSPPaymaster binds a generic wrapper to an already deployed contract.
func bindUSPPaymaster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := USPPaymasterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_USPPaymaster *USPPaymasterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USPPaymaster.Contract.USPPaymasterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_USPPaymaster *USPPaymasterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USPPaymaster.Contract.USPPaymasterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_USPPaymaster *USPPaymasterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USPPaymaster.Contract.USPPaymasterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_USPPaymaster *USPPaymasterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USPPaymaster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_USPPaymaster *USPPaymasterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USPPaymaster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_USPPaymaster *USPPaymasterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USPPaymaster.Contract.contract.Transact(opts, method, params...)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_USPPaymaster *USPPaymasterCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USPPaymaster.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_USPPaymaster *USPPaymasterSession) EntryPoint() (common.Address, error) {
	return _USPPaymaster.Contract.EntryPoint(&_USPPaymaster.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_USPPaymaster *USPPaymasterCallerSession) EntryPoint() (common.Address, error) {
	return _USPPaymaster.Contract.EntryPoint(&_USPPaymaster.CallOpts)
}

// GasSpentByStudent is a free data retrieval call binding the contract method 0xc8493290.
//
// Solidity: function gasSpentByStudent(address ) view returns(uint256)
func (_USPPaymaster *USPPaymasterCaller) GasSpentByStudent(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _USPPaymaster.contract.Call(opts, &out, "gasSpentByStudent", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasSpentByStudent is a free data retrieval call binding the contract method 0xc8493290.
//
// Solidity: function gasSpentByStudent(address ) view returns(uint256)
func (_USPPaymaster *USPPaymasterSession) GasSpentByStudent(arg0 common.Address) (*big.Int, error) {
	return _USPPaymaster.Contract.GasSpentByStudent(&_USPPaymaster.CallOpts, arg0)
}

// GasSpentByStudent is a free data retrieval call binding the contract method 0xc8493290.
//
// Solidity: function gasSpentByStudent(address ) view returns(uint256)
func (_USPPaymaster *USPPaymasterCallerSession) GasSpentByStudent(arg0 common.Address) (*big.Int, error) {
	return _USPPaymaster.Contract.GasSpentByStudent(&_USPPaymaster.CallOpts, arg0)
}

// IdentityRegistry is a free data retrieval call binding the contract method 0x134e18f4.
//
// Solidity: function identityRegistry() view returns(address)
func (_USPPaymaster *USPPaymasterCaller) IdentityRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USPPaymaster.contract.Call(opts, &out, "identityRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IdentityRegistry is a free data retrieval call binding the contract method 0x134e18f4.
//
// Solidity: function identityRegistry() view returns(address)
func (_USPPaymaster *USPPaymasterSession) IdentityRegistry() (common.Address, error) {
	return _USPPaymaster.Contract.IdentityRegistry(&_USPPaymaster.CallOpts)
}

// IdentityRegistry is a free data retrieval call binding the contract method 0x134e18f4.
//
// Solidity: function identityRegistry() view returns(address)
func (_USPPaymaster *USPPaymasterCallerSession) IdentityRegistry() (common.Address, error) {
	return _USPPaymaster.Contract.IdentityRegistry(&_USPPaymaster.CallOpts)
}

// MaxGasPerStudent is a free data retrieval call binding the contract method 0x88155ee8.
//
// Solidity: function maxGasPerStudent() view returns(uint256)
func (_USPPaymaster *USPPaymasterCaller) MaxGasPerStudent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _USPPaymaster.contract.Call(opts, &out, "maxGasPerStudent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGasPerStudent is a free data retrieval call binding the contract method 0x88155ee8.
//
// Solidity: function maxGasPerStudent() view returns(uint256)
func (_USPPaymaster *USPPaymasterSession) MaxGasPerStudent() (*big.Int, error) {
	return _USPPaymaster.Contract.MaxGasPerStudent(&_USPPaymaster.CallOpts)
}

// MaxGasPerStudent is a free data retrieval call binding the contract method 0x88155ee8.
//
// Solidity: function maxGasPerStudent() view returns(uint256)
func (_USPPaymaster *USPPaymasterCallerSession) MaxGasPerStudent() (*big.Int, error) {
	return _USPPaymaster.Contract.MaxGasPerStudent(&_USPPaymaster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_USPPaymaster *USPPaymasterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USPPaymaster.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_USPPaymaster *USPPaymasterSession) Owner() (common.Address, error) {
	return _USPPaymaster.Contract.Owner(&_USPPaymaster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_USPPaymaster *USPPaymasterCallerSession) Owner() (common.Address, error) {
	return _USPPaymaster.Contract.Owner(&_USPPaymaster.CallOpts)
}

// ValidatePaymasterUserOp is a free data retrieval call binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 , uint256 maxCost) view returns(bytes context, uint256 validationData)
func (_USPPaymaster *USPPaymasterCaller) ValidatePaymasterUserOp(opts *bind.CallOpts, userOp PackedUserOperation, arg1 [32]byte, maxCost *big.Int) (struct {
	Context        []byte
	ValidationData *big.Int
}, error) {
	var out []interface{}
	err := _USPPaymaster.contract.Call(opts, &out, "validatePaymasterUserOp", userOp, arg1, maxCost)

	outstruct := new(struct {
		Context        []byte
		ValidationData *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Context = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.ValidationData = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ValidatePaymasterUserOp is a free data retrieval call binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 , uint256 maxCost) view returns(bytes context, uint256 validationData)
func (_USPPaymaster *USPPaymasterSession) ValidatePaymasterUserOp(userOp PackedUserOperation, arg1 [32]byte, maxCost *big.Int) (struct {
	Context        []byte
	ValidationData *big.Int
}, error) {
	return _USPPaymaster.Contract.ValidatePaymasterUserOp(&_USPPaymaster.CallOpts, userOp, arg1, maxCost)
}

// ValidatePaymasterUserOp is a free data retrieval call binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 , uint256 maxCost) view returns(bytes context, uint256 validationData)
func (_USPPaymaster *USPPaymasterCallerSession) ValidatePaymasterUserOp(userOp PackedUserOperation, arg1 [32]byte, maxCost *big.Int) (struct {
	Context        []byte
	ValidationData *big.Int
}, error) {
	return _USPPaymaster.Contract.ValidatePaymasterUserOp(&_USPPaymaster.CallOpts, userOp, arg1, maxCost)
}

// WhitelistedTargets is a free data retrieval call binding the contract method 0xfb9ed257.
//
// Solidity: function whitelistedTargets(address ) view returns(bool)
func (_USPPaymaster *USPPaymasterCaller) WhitelistedTargets(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _USPPaymaster.contract.Call(opts, &out, "whitelistedTargets", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistedTargets is a free data retrieval call binding the contract method 0xfb9ed257.
//
// Solidity: function whitelistedTargets(address ) view returns(bool)
func (_USPPaymaster *USPPaymasterSession) WhitelistedTargets(arg0 common.Address) (bool, error) {
	return _USPPaymaster.Contract.WhitelistedTargets(&_USPPaymaster.CallOpts, arg0)
}

// WhitelistedTargets is a free data retrieval call binding the contract method 0xfb9ed257.
//
// Solidity: function whitelistedTargets(address ) view returns(bool)
func (_USPPaymaster *USPPaymasterCallerSession) WhitelistedTargets(arg0 common.Address) (bool, error) {
	return _USPPaymaster.Contract.WhitelistedTargets(&_USPPaymaster.CallOpts, arg0)
}

// MakeDeposit is a paid mutator transaction binding the contract method 0x40732c89.
//
// Solidity: function makeDeposit() payable returns()
func (_USPPaymaster *USPPaymasterTransactor) MakeDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USPPaymaster.contract.Transact(opts, "makeDeposit")
}

// MakeDeposit is a paid mutator transaction binding the contract method 0x40732c89.
//
// Solidity: function makeDeposit() payable returns()
func (_USPPaymaster *USPPaymasterSession) MakeDeposit() (*types.Transaction, error) {
	return _USPPaymaster.Contract.MakeDeposit(&_USPPaymaster.TransactOpts)
}

// MakeDeposit is a paid mutator transaction binding the contract method 0x40732c89.
//
// Solidity: function makeDeposit() payable returns()
func (_USPPaymaster *USPPaymasterTransactorSession) MakeDeposit() (*types.Transaction, error) {
	return _USPPaymaster.Contract.MakeDeposit(&_USPPaymaster.TransactOpts)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 ) returns()
func (_USPPaymaster *USPPaymasterTransactor) PostOp(opts *bind.TransactOpts, mode uint8, context []byte, actualGasCost *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _USPPaymaster.contract.Transact(opts, "postOp", mode, context, actualGasCost, arg3)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 ) returns()
func (_USPPaymaster *USPPaymasterSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _USPPaymaster.Contract.PostOp(&_USPPaymaster.TransactOpts, mode, context, actualGasCost, arg3)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 ) returns()
func (_USPPaymaster *USPPaymasterTransactorSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _USPPaymaster.Contract.PostOp(&_USPPaymaster.TransactOpts, mode, context, actualGasCost, arg3)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_USPPaymaster *USPPaymasterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USPPaymaster.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_USPPaymaster *USPPaymasterSession) RenounceOwnership() (*types.Transaction, error) {
	return _USPPaymaster.Contract.RenounceOwnership(&_USPPaymaster.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_USPPaymaster *USPPaymasterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _USPPaymaster.Contract.RenounceOwnership(&_USPPaymaster.TransactOpts)
}

// SetMaxGasPerStudent is a paid mutator transaction binding the contract method 0xbfee769f.
//
// Solidity: function setMaxGasPerStudent(uint256 _newLimit) returns()
func (_USPPaymaster *USPPaymasterTransactor) SetMaxGasPerStudent(opts *bind.TransactOpts, _newLimit *big.Int) (*types.Transaction, error) {
	return _USPPaymaster.contract.Transact(opts, "setMaxGasPerStudent", _newLimit)
}

// SetMaxGasPerStudent is a paid mutator transaction binding the contract method 0xbfee769f.
//
// Solidity: function setMaxGasPerStudent(uint256 _newLimit) returns()
func (_USPPaymaster *USPPaymasterSession) SetMaxGasPerStudent(_newLimit *big.Int) (*types.Transaction, error) {
	return _USPPaymaster.Contract.SetMaxGasPerStudent(&_USPPaymaster.TransactOpts, _newLimit)
}

// SetMaxGasPerStudent is a paid mutator transaction binding the contract method 0xbfee769f.
//
// Solidity: function setMaxGasPerStudent(uint256 _newLimit) returns()
func (_USPPaymaster *USPPaymasterTransactorSession) SetMaxGasPerStudent(_newLimit *big.Int) (*types.Transaction, error) {
	return _USPPaymaster.Contract.SetMaxGasPerStudent(&_USPPaymaster.TransactOpts, _newLimit)
}

// SetTargetWhitelist is a paid mutator transaction binding the contract method 0x108bdc14.
//
// Solidity: function setTargetWhitelist(address target, bool status) returns()
func (_USPPaymaster *USPPaymasterTransactor) SetTargetWhitelist(opts *bind.TransactOpts, target common.Address, status bool) (*types.Transaction, error) {
	return _USPPaymaster.contract.Transact(opts, "setTargetWhitelist", target, status)
}

// SetTargetWhitelist is a paid mutator transaction binding the contract method 0x108bdc14.
//
// Solidity: function setTargetWhitelist(address target, bool status) returns()
func (_USPPaymaster *USPPaymasterSession) SetTargetWhitelist(target common.Address, status bool) (*types.Transaction, error) {
	return _USPPaymaster.Contract.SetTargetWhitelist(&_USPPaymaster.TransactOpts, target, status)
}

// SetTargetWhitelist is a paid mutator transaction binding the contract method 0x108bdc14.
//
// Solidity: function setTargetWhitelist(address target, bool status) returns()
func (_USPPaymaster *USPPaymasterTransactorSession) SetTargetWhitelist(target common.Address, status bool) (*types.Transaction, error) {
	return _USPPaymaster.Contract.SetTargetWhitelist(&_USPPaymaster.TransactOpts, target, status)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_USPPaymaster *USPPaymasterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _USPPaymaster.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_USPPaymaster *USPPaymasterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _USPPaymaster.Contract.TransferOwnership(&_USPPaymaster.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_USPPaymaster *USPPaymasterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _USPPaymaster.Contract.TransferOwnership(&_USPPaymaster.TransactOpts, newOwner)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_USPPaymaster *USPPaymasterTransactor) WithdrawTo(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USPPaymaster.contract.Transact(opts, "withdrawTo", withdrawAddress, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_USPPaymaster *USPPaymasterSession) WithdrawTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USPPaymaster.Contract.WithdrawTo(&_USPPaymaster.TransactOpts, withdrawAddress, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_USPPaymaster *USPPaymasterTransactorSession) WithdrawTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USPPaymaster.Contract.WithdrawTo(&_USPPaymaster.TransactOpts, withdrawAddress, amount)
}

// USPPaymasterGasChargedIterator is returned from FilterGasCharged and is used to iterate over the raw logs and unpacked data for GasCharged events raised by the USPPaymaster contract.
type USPPaymasterGasChargedIterator struct {
	Event *USPPaymasterGasCharged // Event containing the contract specifics and raw log

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
func (it *USPPaymasterGasChargedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USPPaymasterGasCharged)
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
		it.Event = new(USPPaymasterGasCharged)
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
func (it *USPPaymasterGasChargedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USPPaymasterGasChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USPPaymasterGasCharged represents a GasCharged event raised by the USPPaymaster contract.
type USPPaymasterGasCharged struct {
	Student       common.Address
	ActualGasCost *big.Int
	Mode          uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterGasCharged is a free log retrieval operation binding the contract event 0x1d7e0371f837c884650a0c4575efcbdeff16ce3ca7dfeb357092a19bbb312416.
//
// Solidity: event GasCharged(address indexed student, uint256 actualGasCost, uint8 mode)
func (_USPPaymaster *USPPaymasterFilterer) FilterGasCharged(opts *bind.FilterOpts, student []common.Address) (*USPPaymasterGasChargedIterator, error) {

	var studentRule []interface{}
	for _, studentItem := range student {
		studentRule = append(studentRule, studentItem)
	}

	logs, sub, err := _USPPaymaster.contract.FilterLogs(opts, "GasCharged", studentRule)
	if err != nil {
		return nil, err
	}
	return &USPPaymasterGasChargedIterator{contract: _USPPaymaster.contract, event: "GasCharged", logs: logs, sub: sub}, nil
}

// WatchGasCharged is a free log subscription operation binding the contract event 0x1d7e0371f837c884650a0c4575efcbdeff16ce3ca7dfeb357092a19bbb312416.
//
// Solidity: event GasCharged(address indexed student, uint256 actualGasCost, uint8 mode)
func (_USPPaymaster *USPPaymasterFilterer) WatchGasCharged(opts *bind.WatchOpts, sink chan<- *USPPaymasterGasCharged, student []common.Address) (event.Subscription, error) {

	var studentRule []interface{}
	for _, studentItem := range student {
		studentRule = append(studentRule, studentItem)
	}

	logs, sub, err := _USPPaymaster.contract.WatchLogs(opts, "GasCharged", studentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USPPaymasterGasCharged)
				if err := _USPPaymaster.contract.UnpackLog(event, "GasCharged", log); err != nil {
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

// ParseGasCharged is a log parse operation binding the contract event 0x1d7e0371f837c884650a0c4575efcbdeff16ce3ca7dfeb357092a19bbb312416.
//
// Solidity: event GasCharged(address indexed student, uint256 actualGasCost, uint8 mode)
func (_USPPaymaster *USPPaymasterFilterer) ParseGasCharged(log types.Log) (*USPPaymasterGasCharged, error) {
	event := new(USPPaymasterGasCharged)
	if err := _USPPaymaster.contract.UnpackLog(event, "GasCharged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// USPPaymasterGasLimitUpdatedIterator is returned from FilterGasLimitUpdated and is used to iterate over the raw logs and unpacked data for GasLimitUpdated events raised by the USPPaymaster contract.
type USPPaymasterGasLimitUpdatedIterator struct {
	Event *USPPaymasterGasLimitUpdated // Event containing the contract specifics and raw log

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
func (it *USPPaymasterGasLimitUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USPPaymasterGasLimitUpdated)
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
		it.Event = new(USPPaymasterGasLimitUpdated)
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
func (it *USPPaymasterGasLimitUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USPPaymasterGasLimitUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USPPaymasterGasLimitUpdated represents a GasLimitUpdated event raised by the USPPaymaster contract.
type USPPaymasterGasLimitUpdated struct {
	NewLimit *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGasLimitUpdated is a free log retrieval operation binding the contract event 0x3d1394ba0f6fca9c1e344f10a3efe1bfca63bc591232bb0d76755690f409450c.
//
// Solidity: event GasLimitUpdated(uint256 newLimit)
func (_USPPaymaster *USPPaymasterFilterer) FilterGasLimitUpdated(opts *bind.FilterOpts) (*USPPaymasterGasLimitUpdatedIterator, error) {

	logs, sub, err := _USPPaymaster.contract.FilterLogs(opts, "GasLimitUpdated")
	if err != nil {
		return nil, err
	}
	return &USPPaymasterGasLimitUpdatedIterator{contract: _USPPaymaster.contract, event: "GasLimitUpdated", logs: logs, sub: sub}, nil
}

// WatchGasLimitUpdated is a free log subscription operation binding the contract event 0x3d1394ba0f6fca9c1e344f10a3efe1bfca63bc591232bb0d76755690f409450c.
//
// Solidity: event GasLimitUpdated(uint256 newLimit)
func (_USPPaymaster *USPPaymasterFilterer) WatchGasLimitUpdated(opts *bind.WatchOpts, sink chan<- *USPPaymasterGasLimitUpdated) (event.Subscription, error) {

	logs, sub, err := _USPPaymaster.contract.WatchLogs(opts, "GasLimitUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USPPaymasterGasLimitUpdated)
				if err := _USPPaymaster.contract.UnpackLog(event, "GasLimitUpdated", log); err != nil {
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

// ParseGasLimitUpdated is a log parse operation binding the contract event 0x3d1394ba0f6fca9c1e344f10a3efe1bfca63bc591232bb0d76755690f409450c.
//
// Solidity: event GasLimitUpdated(uint256 newLimit)
func (_USPPaymaster *USPPaymasterFilterer) ParseGasLimitUpdated(log types.Log) (*USPPaymasterGasLimitUpdated, error) {
	event := new(USPPaymasterGasLimitUpdated)
	if err := _USPPaymaster.contract.UnpackLog(event, "GasLimitUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// USPPaymasterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the USPPaymaster contract.
type USPPaymasterOwnershipTransferredIterator struct {
	Event *USPPaymasterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *USPPaymasterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USPPaymasterOwnershipTransferred)
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
		it.Event = new(USPPaymasterOwnershipTransferred)
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
func (it *USPPaymasterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USPPaymasterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USPPaymasterOwnershipTransferred represents a OwnershipTransferred event raised by the USPPaymaster contract.
type USPPaymasterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_USPPaymaster *USPPaymasterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*USPPaymasterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _USPPaymaster.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &USPPaymasterOwnershipTransferredIterator{contract: _USPPaymaster.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_USPPaymaster *USPPaymasterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *USPPaymasterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _USPPaymaster.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USPPaymasterOwnershipTransferred)
				if err := _USPPaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_USPPaymaster *USPPaymasterFilterer) ParseOwnershipTransferred(log types.Log) (*USPPaymasterOwnershipTransferred, error) {
	event := new(USPPaymasterOwnershipTransferred)
	if err := _USPPaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
