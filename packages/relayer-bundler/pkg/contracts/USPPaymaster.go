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
	Bin: "0x60a060405266b1a2bc2ec5000060035534801561001a575f5ffd5b50604051611728380380611728833981810160405281019061003c9190610293565b805f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036100ad575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016100a491906102f2565b60405180910390fd5b6100bc8161013960201b60201c565b508273ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508160015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505061030b565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610227826101fe565b9050919050565b5f6102388261021d565b9050919050565b6102488161022e565b8114610252575f5ffd5b50565b5f815190506102638161023f565b92915050565b6102728161021d565b811461027c575f5ffd5b50565b5f8151905061028d81610269565b92915050565b5f5f5f606084860312156102aa576102a96101fa565b5b5f6102b786828701610255565b93505060206102c88682870161027f565b92505060406102d98682870161027f565b9150509250925092565b6102ec8161021d565b82525050565b5f6020820190506103055f8301846102e3565b92915050565b6080516113f06103385f395f81816103ac0152818161043801528181610688015261087f01526113f05ff3fe6080604052600436106100dc575f3560e01c806388155ee81161007e578063bfee769f11610058578063bfee769f1461025d578063c849329014610285578063f2fde38b146102c1578063fb9ed257146102e9576100dc565b806388155ee8146101df5780638da5cb5b14610209578063b0d691fe14610233576100dc565b806340732c89116100ba57806340732c891461015a57806352b7512c14610164578063715018a6146101a15780637c627b21146101b7576100dc565b8063108bdc14146100e0578063134e18f414610108578063205c287814610132575b5f5ffd5b3480156100eb575f5ffd5b5061010660048036038101906101019190610bc5565b610325565b005b348015610113575f5ffd5b5061011c61037d565b6040516101299190610c5e565b60405180910390f35b34801561013d575f5ffd5b5061015860048036038101906101539190610ce5565b6103a2565b005b610162610436565b005b34801561016f575f5ffd5b5061018a60048036038101906101859190610d79565b6104c0565b604051610198929190610e64565b60405180910390f35b3480156101ac575f5ffd5b506101b5610780565b005b3480156101c2575f5ffd5b506101dd60048036038101906101d89190610f16565b610793565b005b3480156101ea575f5ffd5b506101f3610850565b6040516102009190610f9a565b60405180910390f35b348015610214575f5ffd5b5061021d610856565b60405161022a9190610fc2565b60405180910390f35b34801561023e575f5ffd5b5061024761087d565b6040516102549190610ffb565b60405180910390f35b348015610268575f5ffd5b50610283600480360381019061027e9190611014565b6108a1565b005b348015610290575f5ffd5b506102ab60048036038101906102a6919061103f565b6108ea565b6040516102b89190610f9a565b60405180910390f35b3480156102cc575f5ffd5b506102e760048036038101906102e2919061103f565b6108ff565b005b3480156102f4575f5ffd5b5061030f600480360381019061030a919061103f565b610983565b60405161031c9190611079565b60405180910390f35b8060045f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6103aa6109a0565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663205c287883836040518363ffffffff1660e01b81526004016104059291906110a1565b5f604051808303815f87803b15801561041c575f5ffd5b505af115801561042e573d5f5f3e3d5ffd5b505050505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663b760faf934306040518363ffffffff1660e01b81526004016104909190610fc2565b5f604051808303818588803b1580156104a7575f5ffd5b505af11580156104b9573d5f5f3e3d5ffd5b5050505050565b60605f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663095f7b89865f016020810190610512919061103f565b6040518263ffffffff1660e01b815260040161052e9190610fc2565b602060405180830381865afa158015610549573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061056d91906110dc565b61058c57600160405180602001604052805f8152509091509150610778565b5f6105a58680606001906105a09190611113565b610a27565b905060045f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1661061157600160405180602001604052805f815250909250925050610778565b60035460025f885f016020810190610629919061103f565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054111561068557600160405180602001604052805f815250909250925050610778565b837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016106df9190610fc2565b602060405180830381865afa1580156106fa573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061071e9190611189565b101561074057600160405180602001604052805f815250909250925050610778565b855f016020810190610752919061103f565b6040516020016107629190610fc2565b6040516020818303038152906040525f92509250505b935093915050565b6107886109a0565b6107915f610a66565b565b5f84848101906107a391906111b4565b90508260025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546107f1919061120c565b925050819055508073ffffffffffffffffffffffffffffffffffffffff167f1d7e0371f837c884650a0c4575efcbdeff16ce3ca7dfeb357092a19bbb31241684886040516108409291906112b2565b60405180910390a2505050505050565b60035481565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b7f000000000000000000000000000000000000000000000000000000000000000081565b6108a96109a0565b806003819055507f3d1394ba0f6fca9c1e344f10a3efe1bfca63bc591232bb0d76755690f409450c816040516108df9190610f9a565b60405180910390a150565b6002602052805f5260405f205f915090505481565b6109076109a0565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610977575f6040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040161096e9190610fc2565b60405180910390fd5b61098081610a66565b50565b6004602052805f5260405f205f915054906101000a900460ff1681565b6109a8610b27565b73ffffffffffffffffffffffffffffffffffffffff166109c6610856565b73ffffffffffffffffffffffffffffffffffffffff1614610a25576109e9610b27565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401610a1c9190610fc2565b60405180910390fd5b565b5f6014838390501015610a3c575f9050610a60565b82825f90601492610a4f939291906112e1565b90610a5a919061135c565b60601c90505b92915050565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f33905090565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610b5f82610b36565b9050919050565b610b6f81610b55565b8114610b79575f5ffd5b50565b5f81359050610b8a81610b66565b92915050565b5f8115159050919050565b610ba481610b90565b8114610bae575f5ffd5b50565b5f81359050610bbf81610b9b565b92915050565b5f5f60408385031215610bdb57610bda610b2e565b5b5f610be885828601610b7c565b9250506020610bf985828601610bb1565b9150509250929050565b5f819050919050565b5f610c26610c21610c1c84610b36565b610c03565b610b36565b9050919050565b5f610c3782610c0c565b9050919050565b5f610c4882610c2d565b9050919050565b610c5881610c3e565b82525050565b5f602082019050610c715f830184610c4f565b92915050565b5f610c8182610b36565b9050919050565b610c9181610c77565b8114610c9b575f5ffd5b50565b5f81359050610cac81610c88565b92915050565b5f819050919050565b610cc481610cb2565b8114610cce575f5ffd5b50565b5f81359050610cdf81610cbb565b92915050565b5f5f60408385031215610cfb57610cfa610b2e565b5b5f610d0885828601610c9e565b9250506020610d1985828601610cd1565b9150509250929050565b5f5ffd5b5f6101208284031215610d3d57610d3c610d23565b5b81905092915050565b5f819050919050565b610d5881610d46565b8114610d62575f5ffd5b50565b5f81359050610d7381610d4f565b92915050565b5f5f5f60608486031215610d9057610d8f610b2e565b5b5f84013567ffffffffffffffff811115610dad57610dac610b32565b5b610db986828701610d27565b9350506020610dca86828701610d65565b9250506040610ddb86828701610cd1565b9150509250925092565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610e2782610de5565b610e318185610def565b9350610e41818560208601610dff565b610e4a81610e0d565b840191505092915050565b610e5e81610cb2565b82525050565b5f6040820190508181035f830152610e7c8185610e1d565b9050610e8b6020830184610e55565b9392505050565b60038110610e9e575f5ffd5b50565b5f81359050610eaf81610e92565b92915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f840112610ed657610ed5610eb5565b5b8235905067ffffffffffffffff811115610ef357610ef2610eb9565b5b602083019150836001820283011115610f0f57610f0e610ebd565b5b9250929050565b5f5f5f5f5f60808688031215610f2f57610f2e610b2e565b5b5f610f3c88828901610ea1565b955050602086013567ffffffffffffffff811115610f5d57610f5c610b32565b5b610f6988828901610ec1565b94509450506040610f7c88828901610cd1565b9250506060610f8d88828901610cd1565b9150509295509295909350565b5f602082019050610fad5f830184610e55565b92915050565b610fbc81610b55565b82525050565b5f602082019050610fd55f830184610fb3565b92915050565b5f610fe582610c2d565b9050919050565b610ff581610fdb565b82525050565b5f60208201905061100e5f830184610fec565b92915050565b5f6020828403121561102957611028610b2e565b5b5f61103684828501610cd1565b91505092915050565b5f6020828403121561105457611053610b2e565b5b5f61106184828501610b7c565b91505092915050565b61107381610b90565b82525050565b5f60208201905061108c5f83018461106a565b92915050565b61109b81610c77565b82525050565b5f6040820190506110b45f830185611092565b6110c16020830184610e55565b9392505050565b5f815190506110d681610b9b565b92915050565b5f602082840312156110f1576110f0610b2e565b5b5f6110fe848285016110c8565b91505092915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f8335600160200384360303811261112f5761112e611107565b5b80840192508235915067ffffffffffffffff8211156111515761115061110b565b5b60208301925060018202360383131561116d5761116c61110f565b5b509250929050565b5f8151905061118381610cbb565b92915050565b5f6020828403121561119e5761119d610b2e565b5b5f6111ab84828501611175565b91505092915050565b5f602082840312156111c9576111c8610b2e565b5b5f6111d684828501610c9e565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61121682610cb2565b915061122183610cb2565b9250828201905080821115611239576112386111df565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b6003811061127d5761127c61123f565b5b50565b5f81905061128d8261126c565b919050565b5f61129c82611280565b9050919050565b6112ac81611292565b82525050565b5f6040820190506112c55f830185610e55565b6112d260208301846112a3565b9392505050565b5f5ffd5b5f5ffd5b5f5f858511156112f4576112f36112d9565b5b83861115611305576113046112dd565b5b6001850283019150848603905094509492505050565b5f82905092915050565b5f7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000082169050919050565b5f82821b905092915050565b5f611367838361131b565b826113728135611325565b925060148210156113b2576113ad7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000083601403600802611350565b831692505b50509291505056fea26469706673582212208d1f56844ec26d4c53d1252358075dde560203957b5911e80007aef8dc6a6cb564736f6c63430008210033",
}

// USPPaymasterABI is the input ABI used to generate the binding from.
// Deprecated: Use USPPaymasterMetaData.ABI instead.
var USPPaymasterABI = USPPaymasterMetaData.ABI

// USPPaymasterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use USPPaymasterMetaData.Bin instead.
var USPPaymasterBin = USPPaymasterMetaData.Bin

// DeployUSPPaymaster deploys a new Ethereum contract, binding an instance of USPPaymaster to it.
func DeployUSPPaymaster(auth *bind.TransactOpts, backend bind.ContractBackend, _entryPoint common.Address, _identityRegistry common.Address, _admin common.Address) (common.Address, *types.Transaction, *USPPaymaster, error) {
	parsed, err := USPPaymasterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(USPPaymasterBin), backend, _entryPoint, _identityRegistry, _admin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &USPPaymaster{USPPaymasterCaller: USPPaymasterCaller{contract: contract}, USPPaymasterTransactor: USPPaymasterTransactor{contract: contract}, USPPaymasterFilterer: USPPaymasterFilterer{contract: contract}}, nil
}

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
