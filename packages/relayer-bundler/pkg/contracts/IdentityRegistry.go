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
	Bin: "0x608060405234801561000f575f5ffd5b50604051611a02380380611a0283398181016040528101906100319190610238565b6100435f5f1b8261007b60201b60201c565b506100747fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c217758261007b60201b60201c565b5050610263565b5f61008c838361017060201b60201c565b6101665760015f5f8581526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055506101036101d360201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001905061016a565b5f90505b92915050565b5f5f5f8481526020019081526020015f205f015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b5f33905090565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610207826101de565b9050919050565b610217816101fd565b8114610221575f5ffd5b50565b5f815190506102328161020e565b92915050565b5f6020828403121561024d5761024c6101da565b5b5f61025a84828501610224565b91505092915050565b611792806102705f395ff3fe608060405234801561000f575f5ffd5b50600436106100e8575f3560e01c806375b238fc1161008a578063a2c6d8e011610064578063a2c6d8e01461026c578063a6c807a91461028a578063c95c2fc7146102bc578063d547741f146102d8576100e8565b806375b238fc1461020057806391d148541461021e578063a217fddf1461024e576100e8565b80632f2ff15d116100c65780632f2ff15d1461017c57806336568abe146101985780633acbef55146101b45780635e8b8091146101e4576100e8565b806301ffc9a7146100ec578063095f7b891461011c578063248a9ca31461014c575b5f5ffd5b61010660048036038101906101019190610de4565b6102f4565b6040516101139190610e29565b60405180910390f35b61013660048036038101906101319190610e9c565b61036d565b6040516101439190610e29565b60405180910390f35b61016660048036038101906101619190610efa565b6103c1565b6040516101739190610f34565b60405180910390f35b61019660048036038101906101919190610f4d565b6103dd565b005b6101b260048036038101906101ad9190610f4d565b6103ff565b005b6101ce60048036038101906101c99190610fbe565b61047a565b6040516101db9190610e29565b60405180910390f35b6101fe60048036038101906101f99190611138565b6105ab565b005b610208610768565b6040516102159190610f34565b60405180910390f35b61023860048036038101906102339190610f4d565b61078c565b6040516102459190610e29565b60405180910390f35b6102566107ef565b6040516102639190610f34565b60405180910390f35b6102746107f5565b60405161028191906111a1565b60405180910390f35b6102a4600480360381019061029f9190610e9c565b6107fb565b6040516102b39392919061121a565b60405180910390f35b6102d660048036038101906102d19190611280565b6108b3565b005b6102f260048036038101906102ed9190610f4d565b610ab5565b005b5f7f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610366575061036582610ad7565b5b9050919050565b5f60015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f9054906101000a900460ff169050919050565b5f5f5f8381526020019081526020015f20600101549050919050565b6103e6826103c1565b6103ef81610b40565b6103f98383610b54565b50505050565b610407610c3d565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461046b576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6104758282610c44565b505050565b5f5f60015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206040518060600160405290815f82015f9054906101000a900460ff161515151581526020016001820180546104ed906112eb565b80601f0160208091040260200160405190810160405280929190818152602001828054610519906112eb565b80156105645780601f1061053b57610100808354040283529160200191610564565b820191905f5260205f20905b81548152906001019060200180831161054757829003601f168201915b5050505050815260200160028201548152505090505f81604001510361058d575f9150506105a5565b805f015180156105a1575082816040015111155b9150505b92915050565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c217756105d581610b40565b5f60015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001018054610621906112eb565b905014610663576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065a90611365565b60405180910390fd5b60405180606001604052806001151581526020018381526020014381525060015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f820151815f015f6101000a81548160ff02191690831515021790555060208201518160010190816106f09190611534565b506040820151816002015590505060025f81548092919061071090611630565b91905055508273ffffffffffffffffffffffffffffffffffffffff167f546ed786c8923815c5efd04ecbc2b720211ff8b4998a8971c38897c3e75d19968360405161075b9190611677565b60405180910390a2505050565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177581565b5f5f5f8481526020019081526020015f205f015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b5f5f1b81565b60025481565b6001602052805f5260405f205f91509050805f015f9054906101000a900460ff169080600101805461082c906112eb565b80601f0160208091040260200160405190810160405280929190818152602001828054610858906112eb565b80156108a35780601f1061087a576101008083540402835291602001916108a3565b820191905f5260205f20905b81548152906001019060200180831161088657829003601f168201915b5050505050908060020154905083565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c217756108dd81610b40565b5f60015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001018054610929906112eb565b90501161096b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610962906116e1565b60405180910390fd5b5f60015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f9054906101000a900460ff16905082151581151514610aaf578260015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f6101000a81548160ff0219169083151502179055508215610a3d5760025f815480929190610a3390611630565b9190505550610a60565b5f6002541115610a5f5760025f815480929190610a59906116ff565b91905055505b5b8373ffffffffffffffffffffffffffffffffffffffff167f2c7c6b09ce2a51196bb33d1fc78177ab366b43fd142006433161c89ad982d1a584604051610aa69190610e29565b60405180910390a25b50505050565b610abe826103c1565b610ac781610b40565b610ad18383610c44565b50505050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b610b5181610b4c610c3d565b610d2d565b50565b5f610b5f838361078c565b610c335760015f5f8581526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550610bd0610c3d565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019050610c37565b5f90505b92915050565b5f33905090565b5f610c4f838361078c565b15610d23575f5f5f8581526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550610cc0610c3d565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a460019050610d27565b5f90505b92915050565b610d37828261078c565b610d7a5780826040517fe2517d3f000000000000000000000000000000000000000000000000000000008152600401610d71929190611735565b60405180910390fd5b5050565b5f604051905090565b5f5ffd5b5f5ffd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b610dc381610d8f565b8114610dcd575f5ffd5b50565b5f81359050610dde81610dba565b92915050565b5f60208284031215610df957610df8610d87565b5b5f610e0684828501610dd0565b91505092915050565b5f8115159050919050565b610e2381610e0f565b82525050565b5f602082019050610e3c5f830184610e1a565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610e6b82610e42565b9050919050565b610e7b81610e61565b8114610e85575f5ffd5b50565b5f81359050610e9681610e72565b92915050565b5f60208284031215610eb157610eb0610d87565b5b5f610ebe84828501610e88565b91505092915050565b5f819050919050565b610ed981610ec7565b8114610ee3575f5ffd5b50565b5f81359050610ef481610ed0565b92915050565b5f60208284031215610f0f57610f0e610d87565b5b5f610f1c84828501610ee6565b91505092915050565b610f2e81610ec7565b82525050565b5f602082019050610f475f830184610f25565b92915050565b5f5f60408385031215610f6357610f62610d87565b5b5f610f7085828601610ee6565b9250506020610f8185828601610e88565b9150509250929050565b5f819050919050565b610f9d81610f8b565b8114610fa7575f5ffd5b50565b5f81359050610fb881610f94565b92915050565b5f5f60408385031215610fd457610fd3610d87565b5b5f610fe185828601610e88565b9250506020610ff285828601610faa565b9150509250929050565b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61104a82611004565b810181811067ffffffffffffffff8211171561106957611068611014565b5b80604052505050565b5f61107b610d7e565b90506110878282611041565b919050565b5f67ffffffffffffffff8211156110a6576110a5611014565b5b6110af82611004565b9050602081019050919050565b828183375f83830152505050565b5f6110dc6110d78461108c565b611072565b9050828152602081018484840111156110f8576110f7611000565b5b6111038482856110bc565b509392505050565b5f82601f83011261111f5761111e610ffc565b5b813561112f8482602086016110ca565b91505092915050565b5f5f6040838503121561114e5761114d610d87565b5b5f61115b85828601610e88565b925050602083013567ffffffffffffffff81111561117c5761117b610d8b565b5b6111888582860161110b565b9150509250929050565b61119b81610f8b565b82525050565b5f6020820190506111b45f830184611192565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f6111ec826111ba565b6111f681856111c4565b93506112068185602086016111d4565b61120f81611004565b840191505092915050565b5f60608201905061122d5f830186610e1a565b818103602083015261123f81856111e2565b905061124e6040830184611192565b949350505050565b61125f81610e0f565b8114611269575f5ffd5b50565b5f8135905061127a81611256565b92915050565b5f5f6040838503121561129657611295610d87565b5b5f6112a385828601610e88565b92505060206112b48582860161126c565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061130257607f821691505b602082108103611315576113146112be565b5b50919050565b7f4a61207265676973747261646f000000000000000000000000000000000000005f82015250565b5f61134f600d836111c4565b915061135a8261131b565b602082019050919050565b5f6020820190508181035f83015261137c81611343565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026113df7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826113a4565b6113e986836113a4565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61142461141f61141a84610f8b565b611401565b610f8b565b9050919050565b5f819050919050565b61143d8361140a565b6114516114498261142b565b8484546113b0565b825550505050565b5f5f905090565b611468611459565b611473818484611434565b505050565b5f5b828110156114995761148e5f828401611460565b60018101905061147a565b505050565b601f8211156114ec57828211156114eb576114b881611383565b6114c183611395565b6114ca85611395565b60208610156114d7575f90505b8083016114e682840382611478565b505050505b5b505050565b5f82821c905092915050565b5f61150c5f19846008026114f1565b1980831691505092915050565b5f61152483836114fd565b9150826002028217905092915050565b61153d826111ba565b67ffffffffffffffff81111561155657611555611014565b5b61156082546112eb565b61156b82828561149e565b5f60209050601f83116001811461159c575f841561158a578287015190505b6115948582611519565b8655506115fb565b601f1984166115aa86611383565b5f5b828110156115d1578489015182556001820191506020850194506020810190506115ac565b868310156115ee57848901516115ea601f8916826114fd565b8355505b6001600288020188555050505b505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61163a82610f8b565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361166c5761166b611603565b5b600182019050919050565b5f6020820190508181035f83015261168f81846111e2565b905092915050565b7f4e616f207265676973747261646f0000000000000000000000000000000000005f82015250565b5f6116cb600e836111c4565b91506116d682611697565b602082019050919050565b5f6020820190508181035f8301526116f8816116bf565b9050919050565b5f61170982610f8b565b91505f820361171b5761171a611603565b5b600182039050919050565b61172f81610e61565b82525050565b5f6040820190506117485f830185611726565b6117556020830184610f25565b939250505056fea2646970667358221220a94f651d1fbf30e635d5738113f59f9c560647cb1a7450225a7a54b0f513604164736f6c63430008210033",
}

// IdentityRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use IdentityRegistryMetaData.ABI instead.
var IdentityRegistryABI = IdentityRegistryMetaData.ABI

// IdentityRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use IdentityRegistryMetaData.Bin instead.
var IdentityRegistryBin = IdentityRegistryMetaData.Bin

// DeployIdentityRegistry deploys a new Ethereum contract, binding an instance of IdentityRegistry to it.
func DeployIdentityRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, admin common.Address) (common.Address, *types.Transaction, *IdentityRegistry, error) {
	parsed, err := IdentityRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(IdentityRegistryBin), backend, admin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IdentityRegistry{IdentityRegistryCaller: IdentityRegistryCaller{contract: contract}, IdentityRegistryTransactor: IdentityRegistryTransactor{contract: contract}, IdentityRegistryFilterer: IdentityRegistryFilterer{contract: contract}}, nil
}

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
