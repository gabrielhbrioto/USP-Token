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

// StudentGovernanceMetaData contains all meta data concerning the StudentGovernance contract.
var StudentGovernanceMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_identityRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"castVote\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"support\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"identityRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIIdentityRegistryGov\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextProposalId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalCost\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposals\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"snapshotBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"votingDeadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"yesVotes\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"noVotes\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalVotes\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"proposer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"depositReturned\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"propose\",\"inputs\":[{\"name\":\"_target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_description\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"quorumPercentage\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGovernanceParams\",\"inputs\":[{\"name\":\"_quorum\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_period\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_reward\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIGovernanceToken\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"voteReward\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"votingPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"GovernanceParamsUpdated\",\"inputs\":[{\"name\":\"newQuorum\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newPeriod\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newReward\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProposalCreated\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"description\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProposalExecuted\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProposalRejected\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"quorumMet\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VoteCast\",\"inputs\":[{\"name\":\"voter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"proposalId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"support\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
}

// StudentGovernanceABI is the input ABI used to generate the binding from.
// Deprecated: Use StudentGovernanceMetaData.ABI instead.
var StudentGovernanceABI = StudentGovernanceMetaData.ABI

// StudentGovernance is an auto generated Go binding around an Ethereum contract.
type StudentGovernance struct {
	StudentGovernanceCaller     // Read-only binding to the contract
	StudentGovernanceTransactor // Write-only binding to the contract
	StudentGovernanceFilterer   // Log filterer for contract events
}

// StudentGovernanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type StudentGovernanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StudentGovernanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StudentGovernanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StudentGovernanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StudentGovernanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StudentGovernanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StudentGovernanceSession struct {
	Contract     *StudentGovernance // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StudentGovernanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StudentGovernanceCallerSession struct {
	Contract *StudentGovernanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// StudentGovernanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StudentGovernanceTransactorSession struct {
	Contract     *StudentGovernanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// StudentGovernanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type StudentGovernanceRaw struct {
	Contract *StudentGovernance // Generic contract binding to access the raw methods on
}

// StudentGovernanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StudentGovernanceCallerRaw struct {
	Contract *StudentGovernanceCaller // Generic read-only contract binding to access the raw methods on
}

// StudentGovernanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StudentGovernanceTransactorRaw struct {
	Contract *StudentGovernanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStudentGovernance creates a new instance of StudentGovernance, bound to a specific deployed contract.
func NewStudentGovernance(address common.Address, backend bind.ContractBackend) (*StudentGovernance, error) {
	contract, err := bindStudentGovernance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StudentGovernance{StudentGovernanceCaller: StudentGovernanceCaller{contract: contract}, StudentGovernanceTransactor: StudentGovernanceTransactor{contract: contract}, StudentGovernanceFilterer: StudentGovernanceFilterer{contract: contract}}, nil
}

// NewStudentGovernanceCaller creates a new read-only instance of StudentGovernance, bound to a specific deployed contract.
func NewStudentGovernanceCaller(address common.Address, caller bind.ContractCaller) (*StudentGovernanceCaller, error) {
	contract, err := bindStudentGovernance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StudentGovernanceCaller{contract: contract}, nil
}

// NewStudentGovernanceTransactor creates a new write-only instance of StudentGovernance, bound to a specific deployed contract.
func NewStudentGovernanceTransactor(address common.Address, transactor bind.ContractTransactor) (*StudentGovernanceTransactor, error) {
	contract, err := bindStudentGovernance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StudentGovernanceTransactor{contract: contract}, nil
}

// NewStudentGovernanceFilterer creates a new log filterer instance of StudentGovernance, bound to a specific deployed contract.
func NewStudentGovernanceFilterer(address common.Address, filterer bind.ContractFilterer) (*StudentGovernanceFilterer, error) {
	contract, err := bindStudentGovernance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StudentGovernanceFilterer{contract: contract}, nil
}

// bindStudentGovernance binds a generic wrapper to an already deployed contract.
func bindStudentGovernance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StudentGovernanceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StudentGovernance *StudentGovernanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StudentGovernance.Contract.StudentGovernanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StudentGovernance *StudentGovernanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StudentGovernance.Contract.StudentGovernanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StudentGovernance *StudentGovernanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StudentGovernance.Contract.StudentGovernanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StudentGovernance *StudentGovernanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StudentGovernance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StudentGovernance *StudentGovernanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StudentGovernance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StudentGovernance *StudentGovernanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StudentGovernance.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StudentGovernance *StudentGovernanceCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StudentGovernance.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StudentGovernance *StudentGovernanceSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StudentGovernance.Contract.DEFAULTADMINROLE(&_StudentGovernance.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StudentGovernance *StudentGovernanceCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StudentGovernance.Contract.DEFAULTADMINROLE(&_StudentGovernance.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StudentGovernance *StudentGovernanceCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _StudentGovernance.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StudentGovernance *StudentGovernanceSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StudentGovernance.Contract.GetRoleAdmin(&_StudentGovernance.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StudentGovernance *StudentGovernanceCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StudentGovernance.Contract.GetRoleAdmin(&_StudentGovernance.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StudentGovernance *StudentGovernanceCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _StudentGovernance.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StudentGovernance *StudentGovernanceSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StudentGovernance.Contract.HasRole(&_StudentGovernance.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StudentGovernance *StudentGovernanceCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StudentGovernance.Contract.HasRole(&_StudentGovernance.CallOpts, role, account)
}

// IdentityRegistry is a free data retrieval call binding the contract method 0x134e18f4.
//
// Solidity: function identityRegistry() view returns(address)
func (_StudentGovernance *StudentGovernanceCaller) IdentityRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StudentGovernance.contract.Call(opts, &out, "identityRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IdentityRegistry is a free data retrieval call binding the contract method 0x134e18f4.
//
// Solidity: function identityRegistry() view returns(address)
func (_StudentGovernance *StudentGovernanceSession) IdentityRegistry() (common.Address, error) {
	return _StudentGovernance.Contract.IdentityRegistry(&_StudentGovernance.CallOpts)
}

// IdentityRegistry is a free data retrieval call binding the contract method 0x134e18f4.
//
// Solidity: function identityRegistry() view returns(address)
func (_StudentGovernance *StudentGovernanceCallerSession) IdentityRegistry() (common.Address, error) {
	return _StudentGovernance.Contract.IdentityRegistry(&_StudentGovernance.CallOpts)
}

// NextProposalId is a free data retrieval call binding the contract method 0x2ab09d14.
//
// Solidity: function nextProposalId() view returns(uint256)
func (_StudentGovernance *StudentGovernanceCaller) NextProposalId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StudentGovernance.contract.Call(opts, &out, "nextProposalId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextProposalId is a free data retrieval call binding the contract method 0x2ab09d14.
//
// Solidity: function nextProposalId() view returns(uint256)
func (_StudentGovernance *StudentGovernanceSession) NextProposalId() (*big.Int, error) {
	return _StudentGovernance.Contract.NextProposalId(&_StudentGovernance.CallOpts)
}

// NextProposalId is a free data retrieval call binding the contract method 0x2ab09d14.
//
// Solidity: function nextProposalId() view returns(uint256)
func (_StudentGovernance *StudentGovernanceCallerSession) NextProposalId() (*big.Int, error) {
	return _StudentGovernance.Contract.NextProposalId(&_StudentGovernance.CallOpts)
}

// ProposalCost is a free data retrieval call binding the contract method 0xe664f3b2.
//
// Solidity: function proposalCost() view returns(uint256)
func (_StudentGovernance *StudentGovernanceCaller) ProposalCost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StudentGovernance.contract.Call(opts, &out, "proposalCost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalCost is a free data retrieval call binding the contract method 0xe664f3b2.
//
// Solidity: function proposalCost() view returns(uint256)
func (_StudentGovernance *StudentGovernanceSession) ProposalCost() (*big.Int, error) {
	return _StudentGovernance.Contract.ProposalCost(&_StudentGovernance.CallOpts)
}

// ProposalCost is a free data retrieval call binding the contract method 0xe664f3b2.
//
// Solidity: function proposalCost() view returns(uint256)
func (_StudentGovernance *StudentGovernanceCallerSession) ProposalCost() (*big.Int, error) {
	return _StudentGovernance.Contract.ProposalCost(&_StudentGovernance.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(address target, bytes callData, string description, uint256 snapshotBlock, uint256 votingDeadline, uint256 yesVotes, uint256 noVotes, uint256 totalVotes, bool executed, address proposer, bool depositReturned)
func (_StudentGovernance *StudentGovernanceCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Target          common.Address
	CallData        []byte
	Description     string
	SnapshotBlock   *big.Int
	VotingDeadline  *big.Int
	YesVotes        *big.Int
	NoVotes         *big.Int
	TotalVotes      *big.Int
	Executed        bool
	Proposer        common.Address
	DepositReturned bool
}, error) {
	var out []interface{}
	err := _StudentGovernance.contract.Call(opts, &out, "proposals", arg0)

	outstruct := new(struct {
		Target          common.Address
		CallData        []byte
		Description     string
		SnapshotBlock   *big.Int
		VotingDeadline  *big.Int
		YesVotes        *big.Int
		NoVotes         *big.Int
		TotalVotes      *big.Int
		Executed        bool
		Proposer        common.Address
		DepositReturned bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Target = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.CallData = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.Description = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.SnapshotBlock = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VotingDeadline = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.YesVotes = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.NoVotes = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.TotalVotes = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Executed = *abi.ConvertType(out[8], new(bool)).(*bool)
	outstruct.Proposer = *abi.ConvertType(out[9], new(common.Address)).(*common.Address)
	outstruct.DepositReturned = *abi.ConvertType(out[10], new(bool)).(*bool)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(address target, bytes callData, string description, uint256 snapshotBlock, uint256 votingDeadline, uint256 yesVotes, uint256 noVotes, uint256 totalVotes, bool executed, address proposer, bool depositReturned)
func (_StudentGovernance *StudentGovernanceSession) Proposals(arg0 *big.Int) (struct {
	Target          common.Address
	CallData        []byte
	Description     string
	SnapshotBlock   *big.Int
	VotingDeadline  *big.Int
	YesVotes        *big.Int
	NoVotes         *big.Int
	TotalVotes      *big.Int
	Executed        bool
	Proposer        common.Address
	DepositReturned bool
}, error) {
	return _StudentGovernance.Contract.Proposals(&_StudentGovernance.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(address target, bytes callData, string description, uint256 snapshotBlock, uint256 votingDeadline, uint256 yesVotes, uint256 noVotes, uint256 totalVotes, bool executed, address proposer, bool depositReturned)
func (_StudentGovernance *StudentGovernanceCallerSession) Proposals(arg0 *big.Int) (struct {
	Target          common.Address
	CallData        []byte
	Description     string
	SnapshotBlock   *big.Int
	VotingDeadline  *big.Int
	YesVotes        *big.Int
	NoVotes         *big.Int
	TotalVotes      *big.Int
	Executed        bool
	Proposer        common.Address
	DepositReturned bool
}, error) {
	return _StudentGovernance.Contract.Proposals(&_StudentGovernance.CallOpts, arg0)
}

// QuorumPercentage is a free data retrieval call binding the contract method 0x4fa76ec9.
//
// Solidity: function quorumPercentage() view returns(uint256)
func (_StudentGovernance *StudentGovernanceCaller) QuorumPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StudentGovernance.contract.Call(opts, &out, "quorumPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumPercentage is a free data retrieval call binding the contract method 0x4fa76ec9.
//
// Solidity: function quorumPercentage() view returns(uint256)
func (_StudentGovernance *StudentGovernanceSession) QuorumPercentage() (*big.Int, error) {
	return _StudentGovernance.Contract.QuorumPercentage(&_StudentGovernance.CallOpts)
}

// QuorumPercentage is a free data retrieval call binding the contract method 0x4fa76ec9.
//
// Solidity: function quorumPercentage() view returns(uint256)
func (_StudentGovernance *StudentGovernanceCallerSession) QuorumPercentage() (*big.Int, error) {
	return _StudentGovernance.Contract.QuorumPercentage(&_StudentGovernance.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StudentGovernance *StudentGovernanceCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _StudentGovernance.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StudentGovernance *StudentGovernanceSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StudentGovernance.Contract.SupportsInterface(&_StudentGovernance.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StudentGovernance *StudentGovernanceCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StudentGovernance.Contract.SupportsInterface(&_StudentGovernance.CallOpts, interfaceId)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_StudentGovernance *StudentGovernanceCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StudentGovernance.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_StudentGovernance *StudentGovernanceSession) Token() (common.Address, error) {
	return _StudentGovernance.Contract.Token(&_StudentGovernance.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_StudentGovernance *StudentGovernanceCallerSession) Token() (common.Address, error) {
	return _StudentGovernance.Contract.Token(&_StudentGovernance.CallOpts)
}

// VoteReward is a free data retrieval call binding the contract method 0x542b1c5d.
//
// Solidity: function voteReward() view returns(uint256)
func (_StudentGovernance *StudentGovernanceCaller) VoteReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StudentGovernance.contract.Call(opts, &out, "voteReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoteReward is a free data retrieval call binding the contract method 0x542b1c5d.
//
// Solidity: function voteReward() view returns(uint256)
func (_StudentGovernance *StudentGovernanceSession) VoteReward() (*big.Int, error) {
	return _StudentGovernance.Contract.VoteReward(&_StudentGovernance.CallOpts)
}

// VoteReward is a free data retrieval call binding the contract method 0x542b1c5d.
//
// Solidity: function voteReward() view returns(uint256)
func (_StudentGovernance *StudentGovernanceCallerSession) VoteReward() (*big.Int, error) {
	return _StudentGovernance.Contract.VoteReward(&_StudentGovernance.CallOpts)
}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() view returns(uint256)
func (_StudentGovernance *StudentGovernanceCaller) VotingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StudentGovernance.contract.Call(opts, &out, "votingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() view returns(uint256)
func (_StudentGovernance *StudentGovernanceSession) VotingPeriod() (*big.Int, error) {
	return _StudentGovernance.Contract.VotingPeriod(&_StudentGovernance.CallOpts)
}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() view returns(uint256)
func (_StudentGovernance *StudentGovernanceCallerSession) VotingPeriod() (*big.Int, error) {
	return _StudentGovernance.Contract.VotingPeriod(&_StudentGovernance.CallOpts)
}

// CastVote is a paid mutator transaction binding the contract method 0x15373e3d.
//
// Solidity: function castVote(uint256 proposalId, bool support) returns()
func (_StudentGovernance *StudentGovernanceTransactor) CastVote(opts *bind.TransactOpts, proposalId *big.Int, support bool) (*types.Transaction, error) {
	return _StudentGovernance.contract.Transact(opts, "castVote", proposalId, support)
}

// CastVote is a paid mutator transaction binding the contract method 0x15373e3d.
//
// Solidity: function castVote(uint256 proposalId, bool support) returns()
func (_StudentGovernance *StudentGovernanceSession) CastVote(proposalId *big.Int, support bool) (*types.Transaction, error) {
	return _StudentGovernance.Contract.CastVote(&_StudentGovernance.TransactOpts, proposalId, support)
}

// CastVote is a paid mutator transaction binding the contract method 0x15373e3d.
//
// Solidity: function castVote(uint256 proposalId, bool support) returns()
func (_StudentGovernance *StudentGovernanceTransactorSession) CastVote(proposalId *big.Int, support bool) (*types.Transaction, error) {
	return _StudentGovernance.Contract.CastVote(&_StudentGovernance.TransactOpts, proposalId, support)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 proposalId) returns()
func (_StudentGovernance *StudentGovernanceTransactor) Execute(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _StudentGovernance.contract.Transact(opts, "execute", proposalId)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 proposalId) returns()
func (_StudentGovernance *StudentGovernanceSession) Execute(proposalId *big.Int) (*types.Transaction, error) {
	return _StudentGovernance.Contract.Execute(&_StudentGovernance.TransactOpts, proposalId)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 proposalId) returns()
func (_StudentGovernance *StudentGovernanceTransactorSession) Execute(proposalId *big.Int) (*types.Transaction, error) {
	return _StudentGovernance.Contract.Execute(&_StudentGovernance.TransactOpts, proposalId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StudentGovernance *StudentGovernanceTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StudentGovernance.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StudentGovernance *StudentGovernanceSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StudentGovernance.Contract.GrantRole(&_StudentGovernance.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StudentGovernance *StudentGovernanceTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StudentGovernance.Contract.GrantRole(&_StudentGovernance.TransactOpts, role, account)
}

// Propose is a paid mutator transaction binding the contract method 0x3153fedb.
//
// Solidity: function propose(address _target, bytes _callData, string _description) returns()
func (_StudentGovernance *StudentGovernanceTransactor) Propose(opts *bind.TransactOpts, _target common.Address, _callData []byte, _description string) (*types.Transaction, error) {
	return _StudentGovernance.contract.Transact(opts, "propose", _target, _callData, _description)
}

// Propose is a paid mutator transaction binding the contract method 0x3153fedb.
//
// Solidity: function propose(address _target, bytes _callData, string _description) returns()
func (_StudentGovernance *StudentGovernanceSession) Propose(_target common.Address, _callData []byte, _description string) (*types.Transaction, error) {
	return _StudentGovernance.Contract.Propose(&_StudentGovernance.TransactOpts, _target, _callData, _description)
}

// Propose is a paid mutator transaction binding the contract method 0x3153fedb.
//
// Solidity: function propose(address _target, bytes _callData, string _description) returns()
func (_StudentGovernance *StudentGovernanceTransactorSession) Propose(_target common.Address, _callData []byte, _description string) (*types.Transaction, error) {
	return _StudentGovernance.Contract.Propose(&_StudentGovernance.TransactOpts, _target, _callData, _description)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_StudentGovernance *StudentGovernanceTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _StudentGovernance.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_StudentGovernance *StudentGovernanceSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _StudentGovernance.Contract.RenounceRole(&_StudentGovernance.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_StudentGovernance *StudentGovernanceTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _StudentGovernance.Contract.RenounceRole(&_StudentGovernance.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StudentGovernance *StudentGovernanceTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StudentGovernance.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StudentGovernance *StudentGovernanceSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StudentGovernance.Contract.RevokeRole(&_StudentGovernance.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StudentGovernance *StudentGovernanceTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StudentGovernance.Contract.RevokeRole(&_StudentGovernance.TransactOpts, role, account)
}

// SetGovernanceParams is a paid mutator transaction binding the contract method 0x639ab0ae.
//
// Solidity: function setGovernanceParams(uint256 _quorum, uint256 _period, uint256 _reward) returns()
func (_StudentGovernance *StudentGovernanceTransactor) SetGovernanceParams(opts *bind.TransactOpts, _quorum *big.Int, _period *big.Int, _reward *big.Int) (*types.Transaction, error) {
	return _StudentGovernance.contract.Transact(opts, "setGovernanceParams", _quorum, _period, _reward)
}

// SetGovernanceParams is a paid mutator transaction binding the contract method 0x639ab0ae.
//
// Solidity: function setGovernanceParams(uint256 _quorum, uint256 _period, uint256 _reward) returns()
func (_StudentGovernance *StudentGovernanceSession) SetGovernanceParams(_quorum *big.Int, _period *big.Int, _reward *big.Int) (*types.Transaction, error) {
	return _StudentGovernance.Contract.SetGovernanceParams(&_StudentGovernance.TransactOpts, _quorum, _period, _reward)
}

// SetGovernanceParams is a paid mutator transaction binding the contract method 0x639ab0ae.
//
// Solidity: function setGovernanceParams(uint256 _quorum, uint256 _period, uint256 _reward) returns()
func (_StudentGovernance *StudentGovernanceTransactorSession) SetGovernanceParams(_quorum *big.Int, _period *big.Int, _reward *big.Int) (*types.Transaction, error) {
	return _StudentGovernance.Contract.SetGovernanceParams(&_StudentGovernance.TransactOpts, _quorum, _period, _reward)
}

// StudentGovernanceGovernanceParamsUpdatedIterator is returned from FilterGovernanceParamsUpdated and is used to iterate over the raw logs and unpacked data for GovernanceParamsUpdated events raised by the StudentGovernance contract.
type StudentGovernanceGovernanceParamsUpdatedIterator struct {
	Event *StudentGovernanceGovernanceParamsUpdated // Event containing the contract specifics and raw log

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
func (it *StudentGovernanceGovernanceParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StudentGovernanceGovernanceParamsUpdated)
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
		it.Event = new(StudentGovernanceGovernanceParamsUpdated)
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
func (it *StudentGovernanceGovernanceParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StudentGovernanceGovernanceParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StudentGovernanceGovernanceParamsUpdated represents a GovernanceParamsUpdated event raised by the StudentGovernance contract.
type StudentGovernanceGovernanceParamsUpdated struct {
	NewQuorum *big.Int
	NewPeriod *big.Int
	NewReward *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGovernanceParamsUpdated is a free log retrieval operation binding the contract event 0x19692f5ce33d869604f6435a807e9b9cb563a045c2fe38511e0daf7bc8095cd6.
//
// Solidity: event GovernanceParamsUpdated(uint256 newQuorum, uint256 newPeriod, uint256 newReward)
func (_StudentGovernance *StudentGovernanceFilterer) FilterGovernanceParamsUpdated(opts *bind.FilterOpts) (*StudentGovernanceGovernanceParamsUpdatedIterator, error) {

	logs, sub, err := _StudentGovernance.contract.FilterLogs(opts, "GovernanceParamsUpdated")
	if err != nil {
		return nil, err
	}
	return &StudentGovernanceGovernanceParamsUpdatedIterator{contract: _StudentGovernance.contract, event: "GovernanceParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchGovernanceParamsUpdated is a free log subscription operation binding the contract event 0x19692f5ce33d869604f6435a807e9b9cb563a045c2fe38511e0daf7bc8095cd6.
//
// Solidity: event GovernanceParamsUpdated(uint256 newQuorum, uint256 newPeriod, uint256 newReward)
func (_StudentGovernance *StudentGovernanceFilterer) WatchGovernanceParamsUpdated(opts *bind.WatchOpts, sink chan<- *StudentGovernanceGovernanceParamsUpdated) (event.Subscription, error) {

	logs, sub, err := _StudentGovernance.contract.WatchLogs(opts, "GovernanceParamsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StudentGovernanceGovernanceParamsUpdated)
				if err := _StudentGovernance.contract.UnpackLog(event, "GovernanceParamsUpdated", log); err != nil {
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

// ParseGovernanceParamsUpdated is a log parse operation binding the contract event 0x19692f5ce33d869604f6435a807e9b9cb563a045c2fe38511e0daf7bc8095cd6.
//
// Solidity: event GovernanceParamsUpdated(uint256 newQuorum, uint256 newPeriod, uint256 newReward)
func (_StudentGovernance *StudentGovernanceFilterer) ParseGovernanceParamsUpdated(log types.Log) (*StudentGovernanceGovernanceParamsUpdated, error) {
	event := new(StudentGovernanceGovernanceParamsUpdated)
	if err := _StudentGovernance.contract.UnpackLog(event, "GovernanceParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StudentGovernanceProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the StudentGovernance contract.
type StudentGovernanceProposalCreatedIterator struct {
	Event *StudentGovernanceProposalCreated // Event containing the contract specifics and raw log

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
func (it *StudentGovernanceProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StudentGovernanceProposalCreated)
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
		it.Event = new(StudentGovernanceProposalCreated)
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
func (it *StudentGovernanceProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StudentGovernanceProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StudentGovernanceProposalCreated represents a ProposalCreated event raised by the StudentGovernance contract.
type StudentGovernanceProposalCreated struct {
	Id          *big.Int
	Description string
	Deadline    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0xf5dc0da9afa6d080780a49fd8c680f03d2ddf3c067cd6d0c85360c096644d2d3.
//
// Solidity: event ProposalCreated(uint256 indexed id, string description, uint256 deadline)
func (_StudentGovernance *StudentGovernanceFilterer) FilterProposalCreated(opts *bind.FilterOpts, id []*big.Int) (*StudentGovernanceProposalCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _StudentGovernance.contract.FilterLogs(opts, "ProposalCreated", idRule)
	if err != nil {
		return nil, err
	}
	return &StudentGovernanceProposalCreatedIterator{contract: _StudentGovernance.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0xf5dc0da9afa6d080780a49fd8c680f03d2ddf3c067cd6d0c85360c096644d2d3.
//
// Solidity: event ProposalCreated(uint256 indexed id, string description, uint256 deadline)
func (_StudentGovernance *StudentGovernanceFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *StudentGovernanceProposalCreated, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _StudentGovernance.contract.WatchLogs(opts, "ProposalCreated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StudentGovernanceProposalCreated)
				if err := _StudentGovernance.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
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

// ParseProposalCreated is a log parse operation binding the contract event 0xf5dc0da9afa6d080780a49fd8c680f03d2ddf3c067cd6d0c85360c096644d2d3.
//
// Solidity: event ProposalCreated(uint256 indexed id, string description, uint256 deadline)
func (_StudentGovernance *StudentGovernanceFilterer) ParseProposalCreated(log types.Log) (*StudentGovernanceProposalCreated, error) {
	event := new(StudentGovernanceProposalCreated)
	if err := _StudentGovernance.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StudentGovernanceProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the StudentGovernance contract.
type StudentGovernanceProposalExecutedIterator struct {
	Event *StudentGovernanceProposalExecuted // Event containing the contract specifics and raw log

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
func (it *StudentGovernanceProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StudentGovernanceProposalExecuted)
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
		it.Event = new(StudentGovernanceProposalExecuted)
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
func (it *StudentGovernanceProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StudentGovernanceProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StudentGovernanceProposalExecuted represents a ProposalExecuted event raised by the StudentGovernance contract.
type StudentGovernanceProposalExecuted struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 indexed id)
func (_StudentGovernance *StudentGovernanceFilterer) FilterProposalExecuted(opts *bind.FilterOpts, id []*big.Int) (*StudentGovernanceProposalExecutedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _StudentGovernance.contract.FilterLogs(opts, "ProposalExecuted", idRule)
	if err != nil {
		return nil, err
	}
	return &StudentGovernanceProposalExecutedIterator{contract: _StudentGovernance.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 indexed id)
func (_StudentGovernance *StudentGovernanceFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *StudentGovernanceProposalExecuted, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _StudentGovernance.contract.WatchLogs(opts, "ProposalExecuted", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StudentGovernanceProposalExecuted)
				if err := _StudentGovernance.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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

// ParseProposalExecuted is a log parse operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 indexed id)
func (_StudentGovernance *StudentGovernanceFilterer) ParseProposalExecuted(log types.Log) (*StudentGovernanceProposalExecuted, error) {
	event := new(StudentGovernanceProposalExecuted)
	if err := _StudentGovernance.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StudentGovernanceProposalRejectedIterator is returned from FilterProposalRejected and is used to iterate over the raw logs and unpacked data for ProposalRejected events raised by the StudentGovernance contract.
type StudentGovernanceProposalRejectedIterator struct {
	Event *StudentGovernanceProposalRejected // Event containing the contract specifics and raw log

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
func (it *StudentGovernanceProposalRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StudentGovernanceProposalRejected)
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
		it.Event = new(StudentGovernanceProposalRejected)
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
func (it *StudentGovernanceProposalRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StudentGovernanceProposalRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StudentGovernanceProposalRejected represents a ProposalRejected event raised by the StudentGovernance contract.
type StudentGovernanceProposalRejected struct {
	Id        *big.Int
	QuorumMet bool
	Approved  bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProposalRejected is a free log retrieval operation binding the contract event 0xd973f1640c6d7b706edfc94170f9f4fd263bedf67a06968fde698fbb5846e2ce.
//
// Solidity: event ProposalRejected(uint256 indexed id, bool quorumMet, bool approved)
func (_StudentGovernance *StudentGovernanceFilterer) FilterProposalRejected(opts *bind.FilterOpts, id []*big.Int) (*StudentGovernanceProposalRejectedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _StudentGovernance.contract.FilterLogs(opts, "ProposalRejected", idRule)
	if err != nil {
		return nil, err
	}
	return &StudentGovernanceProposalRejectedIterator{contract: _StudentGovernance.contract, event: "ProposalRejected", logs: logs, sub: sub}, nil
}

// WatchProposalRejected is a free log subscription operation binding the contract event 0xd973f1640c6d7b706edfc94170f9f4fd263bedf67a06968fde698fbb5846e2ce.
//
// Solidity: event ProposalRejected(uint256 indexed id, bool quorumMet, bool approved)
func (_StudentGovernance *StudentGovernanceFilterer) WatchProposalRejected(opts *bind.WatchOpts, sink chan<- *StudentGovernanceProposalRejected, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _StudentGovernance.contract.WatchLogs(opts, "ProposalRejected", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StudentGovernanceProposalRejected)
				if err := _StudentGovernance.contract.UnpackLog(event, "ProposalRejected", log); err != nil {
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

// ParseProposalRejected is a log parse operation binding the contract event 0xd973f1640c6d7b706edfc94170f9f4fd263bedf67a06968fde698fbb5846e2ce.
//
// Solidity: event ProposalRejected(uint256 indexed id, bool quorumMet, bool approved)
func (_StudentGovernance *StudentGovernanceFilterer) ParseProposalRejected(log types.Log) (*StudentGovernanceProposalRejected, error) {
	event := new(StudentGovernanceProposalRejected)
	if err := _StudentGovernance.contract.UnpackLog(event, "ProposalRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StudentGovernanceRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the StudentGovernance contract.
type StudentGovernanceRoleAdminChangedIterator struct {
	Event *StudentGovernanceRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *StudentGovernanceRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StudentGovernanceRoleAdminChanged)
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
		it.Event = new(StudentGovernanceRoleAdminChanged)
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
func (it *StudentGovernanceRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StudentGovernanceRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StudentGovernanceRoleAdminChanged represents a RoleAdminChanged event raised by the StudentGovernance contract.
type StudentGovernanceRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StudentGovernance *StudentGovernanceFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StudentGovernanceRoleAdminChangedIterator, error) {

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

	logs, sub, err := _StudentGovernance.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StudentGovernanceRoleAdminChangedIterator{contract: _StudentGovernance.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StudentGovernance *StudentGovernanceFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StudentGovernanceRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _StudentGovernance.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StudentGovernanceRoleAdminChanged)
				if err := _StudentGovernance.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_StudentGovernance *StudentGovernanceFilterer) ParseRoleAdminChanged(log types.Log) (*StudentGovernanceRoleAdminChanged, error) {
	event := new(StudentGovernanceRoleAdminChanged)
	if err := _StudentGovernance.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StudentGovernanceRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the StudentGovernance contract.
type StudentGovernanceRoleGrantedIterator struct {
	Event *StudentGovernanceRoleGranted // Event containing the contract specifics and raw log

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
func (it *StudentGovernanceRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StudentGovernanceRoleGranted)
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
		it.Event = new(StudentGovernanceRoleGranted)
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
func (it *StudentGovernanceRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StudentGovernanceRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StudentGovernanceRoleGranted represents a RoleGranted event raised by the StudentGovernance contract.
type StudentGovernanceRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StudentGovernance *StudentGovernanceFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StudentGovernanceRoleGrantedIterator, error) {

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

	logs, sub, err := _StudentGovernance.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StudentGovernanceRoleGrantedIterator{contract: _StudentGovernance.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StudentGovernance *StudentGovernanceFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StudentGovernanceRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StudentGovernance.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StudentGovernanceRoleGranted)
				if err := _StudentGovernance.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_StudentGovernance *StudentGovernanceFilterer) ParseRoleGranted(log types.Log) (*StudentGovernanceRoleGranted, error) {
	event := new(StudentGovernanceRoleGranted)
	if err := _StudentGovernance.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StudentGovernanceRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the StudentGovernance contract.
type StudentGovernanceRoleRevokedIterator struct {
	Event *StudentGovernanceRoleRevoked // Event containing the contract specifics and raw log

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
func (it *StudentGovernanceRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StudentGovernanceRoleRevoked)
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
		it.Event = new(StudentGovernanceRoleRevoked)
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
func (it *StudentGovernanceRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StudentGovernanceRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StudentGovernanceRoleRevoked represents a RoleRevoked event raised by the StudentGovernance contract.
type StudentGovernanceRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StudentGovernance *StudentGovernanceFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StudentGovernanceRoleRevokedIterator, error) {

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

	logs, sub, err := _StudentGovernance.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StudentGovernanceRoleRevokedIterator{contract: _StudentGovernance.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StudentGovernance *StudentGovernanceFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StudentGovernanceRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StudentGovernance.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StudentGovernanceRoleRevoked)
				if err := _StudentGovernance.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_StudentGovernance *StudentGovernanceFilterer) ParseRoleRevoked(log types.Log) (*StudentGovernanceRoleRevoked, error) {
	event := new(StudentGovernanceRoleRevoked)
	if err := _StudentGovernance.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StudentGovernanceVoteCastIterator is returned from FilterVoteCast and is used to iterate over the raw logs and unpacked data for VoteCast events raised by the StudentGovernance contract.
type StudentGovernanceVoteCastIterator struct {
	Event *StudentGovernanceVoteCast // Event containing the contract specifics and raw log

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
func (it *StudentGovernanceVoteCastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StudentGovernanceVoteCast)
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
		it.Event = new(StudentGovernanceVoteCast)
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
func (it *StudentGovernanceVoteCastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StudentGovernanceVoteCastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StudentGovernanceVoteCast represents a VoteCast event raised by the StudentGovernance contract.
type StudentGovernanceVoteCast struct {
	Voter      common.Address
	ProposalId *big.Int
	Support    bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteCast is a free log retrieval operation binding the contract event 0xd356173ae8eeea8691aee4c1be712c314a975a3d43ebc48b08ca54d0dac91228.
//
// Solidity: event VoteCast(address indexed voter, uint256 indexed proposalId, bool support)
func (_StudentGovernance *StudentGovernanceFilterer) FilterVoteCast(opts *bind.FilterOpts, voter []common.Address, proposalId []*big.Int) (*StudentGovernanceVoteCastIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _StudentGovernance.contract.FilterLogs(opts, "VoteCast", voterRule, proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &StudentGovernanceVoteCastIterator{contract: _StudentGovernance.contract, event: "VoteCast", logs: logs, sub: sub}, nil
}

// WatchVoteCast is a free log subscription operation binding the contract event 0xd356173ae8eeea8691aee4c1be712c314a975a3d43ebc48b08ca54d0dac91228.
//
// Solidity: event VoteCast(address indexed voter, uint256 indexed proposalId, bool support)
func (_StudentGovernance *StudentGovernanceFilterer) WatchVoteCast(opts *bind.WatchOpts, sink chan<- *StudentGovernanceVoteCast, voter []common.Address, proposalId []*big.Int) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _StudentGovernance.contract.WatchLogs(opts, "VoteCast", voterRule, proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StudentGovernanceVoteCast)
				if err := _StudentGovernance.contract.UnpackLog(event, "VoteCast", log); err != nil {
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

// ParseVoteCast is a log parse operation binding the contract event 0xd356173ae8eeea8691aee4c1be712c314a975a3d43ebc48b08ca54d0dac91228.
//
// Solidity: event VoteCast(address indexed voter, uint256 indexed proposalId, bool support)
func (_StudentGovernance *StudentGovernanceFilterer) ParseVoteCast(log types.Log) (*StudentGovernanceVoteCast, error) {
	event := new(StudentGovernanceVoteCast)
	if err := _StudentGovernance.contract.UnpackLog(event, "VoteCast", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
