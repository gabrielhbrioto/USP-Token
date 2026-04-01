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

// USPTokenMetaData contains all meta data concerning the USPToken contract.
var USPTokenMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_identityRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MINTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"identityRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIIdentityRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561000f575f5ffd5b5060405161233638038061233683398181016040528101906100319190610309565b6040518060400160405280600981526020017f55535020546f6b656e00000000000000000000000000000000000000000000008152506040518060400160405280600481526020017f552450540000000000000000000000000000000000000000000000000000000081525081600390816100ac9190610595565b5080600490816100bc9190610595565b5050508160065f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506101115f5f1b8261014a60201b60201c565b506101427f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a68261014a60201b60201c565b505050610664565b5f61015b838361024060201b60201c565b61023657600160055f8581526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055506101d36102a460201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001905061023a565b5f90505b92915050565b5f60055f8481526020019081526020015f205f015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b5f33905090565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6102d8826102af565b9050919050565b6102e8816102ce565b81146102f2575f5ffd5b50565b5f81519050610303816102df565b92915050565b5f5f6040838503121561031f5761031e6102ab565b5b5f61032c858286016102f5565b925050602061033d858286016102f5565b9150509250929050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806103c257607f821691505b6020821081036103d5576103d461037e565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026104377fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826103fc565b61044186836103fc565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61048561048061047b84610459565b610462565b610459565b9050919050565b5f819050919050565b61049e8361046b565b6104b26104aa8261048c565b848454610408565b825550505050565b5f5f905090565b6104c96104ba565b6104d4818484610495565b505050565b5f5b828110156104fa576104ef5f8284016104c1565b6001810190506104db565b505050565b601f82111561054d578282111561054c57610519816103db565b610522836103ed565b61052b856103ed565b6020861015610538575f90505b808301610547828403826104d9565b505050505b5b505050565b5f82821c905092915050565b5f61056d5f1984600802610552565b1980831691505092915050565b5f610585838361055e565b9150826002028217905092915050565b61059e82610347565b67ffffffffffffffff8111156105b7576105b6610351565b5b6105c182546103ab565b6105cc8282856104ff565b5f60209050601f8311600181146105fd575f84156105eb578287015190505b6105f5858261057a565b86555061065c565b601f19841661060b866103db565b5f5b828110156106325784890151825560018201915060208501945060208101905061060d565b8683101561064f578489015161064b601f89168261055e565b8355505b6001600288020188555050505b505050505050565b611cc5806106715f395ff3fe608060405234801561000f575f5ffd5b5060043610610135575f3560e01c806340c10f19116100b657806395d89b411161007a57806395d89b411461035d578063a217fddf1461037b578063a9059cbb14610399578063d5391393146103c9578063d547741f146103e7578063dd62ed3e1461040357610135565b806340c10f19146102a957806342966c68146102c557806370a08231146102e157806379cc67901461031157806391d148541461032d57610135565b806323b872dd116100fd57806323b872dd146101f3578063248a9ca3146102235780632f2ff15d14610253578063313ce5671461026f57806336568abe1461028d57610135565b806301ffc9a71461013957806306fdde0314610169578063095ea7b314610187578063134e18f4146101b757806318160ddd146101d5575b5f5ffd5b610153600480360381019061014e919061156a565b610433565b60405161016091906115af565b60405180910390f35b6101716104ac565b60405161017e9190611638565b60405180910390f35b6101a1600480360381019061019c91906116e5565b61053c565b6040516101ae91906115af565b60405180910390f35b6101bf61055e565b6040516101cc919061177e565b60405180910390f35b6101dd610583565b6040516101ea91906117a6565b60405180910390f35b61020d600480360381019061020891906117bf565b61058c565b60405161021a91906115af565b60405180910390f35b61023d60048036038101906102389190611842565b6105ba565b60405161024a919061187c565b60405180910390f35b61026d60048036038101906102689190611895565b6105d7565b005b6102776105f9565b60405161028491906118ee565b60405180910390f35b6102a760048036038101906102a29190611895565b610601565b005b6102c360048036038101906102be91906116e5565b61067c565b005b6102df60048036038101906102da9190611907565b610762565b005b6102fb60048036038101906102f69190611932565b610776565b60405161030891906117a6565b60405180910390f35b61032b600480360381019061032691906116e5565b6107bb565b005b61034760048036038101906103429190611895565b6107db565b60405161035491906115af565b60405180910390f35b61036561083f565b6040516103729190611638565b60405180910390f35b6103836108cf565b604051610390919061187c565b60405180910390f35b6103b360048036038101906103ae91906116e5565b6108d5565b6040516103c091906115af565b60405180910390f35b6103d16108f7565b6040516103de919061187c565b60405180910390f35b61040160048036038101906103fc9190611895565b61091b565b005b61041d6004803603810190610418919061195d565b61093d565b60405161042a91906117a6565b60405180910390f35b5f7f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806104a557506104a4826109bf565b5b9050919050565b6060600380546104bb906119c8565b80601f01602080910402602001604051908101604052809291908181526020018280546104e7906119c8565b80156105325780601f1061050957610100808354040283529160200191610532565b820191905f5260205f20905b81548152906001019060200180831161051557829003601f168201915b5050505050905090565b5f5f610546610a28565b9050610553818585610a2f565b600191505092915050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f600254905090565b5f5f610596610a28565b90506105a3858285610a41565b6105ae858585610ad4565b60019150509392505050565b5f60055f8381526020019081526020015f20600101549050919050565b6105e0826105ba565b6105e981610bc4565b6105f38383610bd8565b50505050565b5f6012905090565b610609610a28565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461066d576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106778282610cc2565b505050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663095f7b89836040518263ffffffff1660e01b81526004016106d69190611a07565b602060405180830381865afa1580156106f1573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107159190611a4a565b610754576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161074b90611ae5565b60405180910390fd5b61075e8282610dac565b5050565b61077361076d610a28565b82610e2b565b50565b5f5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b6107cd826107c7610a28565b83610a41565b6107d78282610e2b565b5050565b5f60055f8481526020019081526020015f205f015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b60606004805461084e906119c8565b80601f016020809104026020016040519081016040528092919081815260200182805461087a906119c8565b80156108c55780601f1061089c576101008083540402835291602001916108c5565b820191905f5260205f20905b8154815290600101906020018083116108a857829003601f168201915b5050505050905090565b5f5f1b81565b5f5f6108df610a28565b90506108ec818585610ad4565b600191505092915050565b7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a681565b610924826105ba565b61092d81610bc4565b6109378383610cc2565b50505050565b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b5f33905090565b610a3c8383836001610eaa565b505050565b5f610a4c848461093d565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811015610ace5781811015610abf578281836040517ffb8f41b2000000000000000000000000000000000000000000000000000000008152600401610ab693929190611b03565b60405180910390fd5b610acd84848484035f610eaa565b5b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610b44575f6040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600401610b3b9190611a07565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610bb4575f6040517fec442f05000000000000000000000000000000000000000000000000000000008152600401610bab9190611a07565b60405180910390fd5b610bbf838383611079565b505050565b610bd581610bd0610a28565b6112a7565b50565b5f610be383836107db565b610cb857600160055f8581526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550610c55610a28565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019050610cbc565b5f90505b92915050565b5f610ccd83836107db565b15610da2575f60055f8581526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550610d3f610a28565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a460019050610da6565b5f90505b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610e1c575f6040517fec442f05000000000000000000000000000000000000000000000000000000008152600401610e139190611a07565b60405180910390fd5b610e275f8383611079565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610e9b575f6040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600401610e929190611a07565b60405180910390fd5b610ea6825f83611079565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610f1a575f6040517fe602df05000000000000000000000000000000000000000000000000000000008152600401610f119190611a07565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610f8a575f6040517f94280d62000000000000000000000000000000000000000000000000000000008152600401610f819190611a07565b60405180910390fd5b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508015611073578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161106a91906117a6565b60405180910390a35b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141580156110e157505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b156112975760065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663095f7b89846040518263ffffffff1660e01b81526004016111409190611a07565b602060405180830381865afa15801561115b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061117f9190611a4a565b6111be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111b590611b82565b60405180910390fd5b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663095f7b89836040518263ffffffff1660e01b81526004016112189190611a07565b602060405180830381865afa158015611233573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112579190611a4a565b611296576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161128d90611bea565b60405180910390fd5b5b6112a28383836112f8565b505050565b6112b182826107db565b6112f45780826040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526004016112eb929190611c08565b60405180910390fd5b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603611348578060025f82825461133c9190611c5c565b92505081905550611416565b5f5f5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050818110156113d1578381836040517fe450d38c0000000000000000000000000000000000000000000000000000000081526004016113c893929190611b03565b60405180910390fd5b8181035f5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361145d578060025f82825403925050819055506114a7565b805f5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161150491906117a6565b60405180910390a3505050565b5f5ffd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b61154981611515565b8114611553575f5ffd5b50565b5f8135905061156481611540565b92915050565b5f6020828403121561157f5761157e611511565b5b5f61158c84828501611556565b91505092915050565b5f8115159050919050565b6115a981611595565b82525050565b5f6020820190506115c25f8301846115a0565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61160a826115c8565b61161481856115d2565b93506116248185602086016115e2565b61162d816115f0565b840191505092915050565b5f6020820190508181035f8301526116508184611600565b905092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61168182611658565b9050919050565b61169181611677565b811461169b575f5ffd5b50565b5f813590506116ac81611688565b92915050565b5f819050919050565b6116c4816116b2565b81146116ce575f5ffd5b50565b5f813590506116df816116bb565b92915050565b5f5f604083850312156116fb576116fa611511565b5b5f6117088582860161169e565b9250506020611719858286016116d1565b9150509250929050565b5f819050919050565b5f61174661174161173c84611658565b611723565b611658565b9050919050565b5f6117578261172c565b9050919050565b5f6117688261174d565b9050919050565b6117788161175e565b82525050565b5f6020820190506117915f83018461176f565b92915050565b6117a0816116b2565b82525050565b5f6020820190506117b95f830184611797565b92915050565b5f5f5f606084860312156117d6576117d5611511565b5b5f6117e38682870161169e565b93505060206117f48682870161169e565b9250506040611805868287016116d1565b9150509250925092565b5f819050919050565b6118218161180f565b811461182b575f5ffd5b50565b5f8135905061183c81611818565b92915050565b5f6020828403121561185757611856611511565b5b5f6118648482850161182e565b91505092915050565b6118768161180f565b82525050565b5f60208201905061188f5f83018461186d565b92915050565b5f5f604083850312156118ab576118aa611511565b5b5f6118b88582860161182e565b92505060206118c98582860161169e565b9150509250929050565b5f60ff82169050919050565b6118e8816118d3565b82525050565b5f6020820190506119015f8301846118df565b92915050565b5f6020828403121561191c5761191b611511565b5b5f611929848285016116d1565b91505092915050565b5f6020828403121561194757611946611511565b5b5f6119548482850161169e565b91505092915050565b5f5f6040838503121561197357611972611511565b5b5f6119808582860161169e565b92505060206119918582860161169e565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806119df57607f821691505b6020821081036119f2576119f161199b565b5b50919050565b611a0181611677565b82525050565b5f602082019050611a1a5f8301846119f8565b92915050565b611a2981611595565b8114611a33575f5ffd5b50565b5f81519050611a4481611a20565b92915050565b5f60208284031215611a5f57611a5e611511565b5b5f611a6c84828501611a36565b91505092915050565b7f44657374696e61746172696f206e616f206520756d206573747564616e7465205f8201527f617469766f000000000000000000000000000000000000000000000000000000602082015250565b5f611acf6025836115d2565b9150611ada82611a75565b604082019050919050565b5f6020820190508181035f830152611afc81611ac3565b9050919050565b5f606082019050611b165f8301866119f8565b611b236020830185611797565b611b306040830184611797565b949350505050565b7f52656d6574656e746520696e617469766f206e6f20726567697374726f0000005f82015250565b5f611b6c601d836115d2565b9150611b7782611b38565b602082019050919050565b5f6020820190508181035f830152611b9981611b60565b9050919050565b7f44657374696e61746172696f20696e617469766f206e6f20726567697374726f5f82015250565b5f611bd46020836115d2565b9150611bdf82611ba0565b602082019050919050565b5f6020820190508181035f830152611c0181611bc8565b9050919050565b5f604082019050611c1b5f8301856119f8565b611c28602083018461186d565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611c66826116b2565b9150611c71836116b2565b9250828201905080821115611c8957611c88611c2f565b5b9291505056fea2646970667358221220c111127c9783b09d5c378e2d71e62d590d0e61b6d08f1c3a0ed386b120b42c6864736f6c63430008210033",
}

// USPTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use USPTokenMetaData.ABI instead.
var USPTokenABI = USPTokenMetaData.ABI

// USPTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use USPTokenMetaData.Bin instead.
var USPTokenBin = USPTokenMetaData.Bin

// DeployUSPToken deploys a new Ethereum contract, binding an instance of USPToken to it.
func DeployUSPToken(auth *bind.TransactOpts, backend bind.ContractBackend, _identityRegistry common.Address, admin common.Address) (common.Address, *types.Transaction, *USPToken, error) {
	parsed, err := USPTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(USPTokenBin), backend, _identityRegistry, admin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &USPToken{USPTokenCaller: USPTokenCaller{contract: contract}, USPTokenTransactor: USPTokenTransactor{contract: contract}, USPTokenFilterer: USPTokenFilterer{contract: contract}}, nil
}

// USPToken is an auto generated Go binding around an Ethereum contract.
type USPToken struct {
	USPTokenCaller     // Read-only binding to the contract
	USPTokenTransactor // Write-only binding to the contract
	USPTokenFilterer   // Log filterer for contract events
}

// USPTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type USPTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USPTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type USPTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USPTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type USPTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USPTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type USPTokenSession struct {
	Contract     *USPToken         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// USPTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type USPTokenCallerSession struct {
	Contract *USPTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// USPTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type USPTokenTransactorSession struct {
	Contract     *USPTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// USPTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type USPTokenRaw struct {
	Contract *USPToken // Generic contract binding to access the raw methods on
}

// USPTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type USPTokenCallerRaw struct {
	Contract *USPTokenCaller // Generic read-only contract binding to access the raw methods on
}

// USPTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type USPTokenTransactorRaw struct {
	Contract *USPTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUSPToken creates a new instance of USPToken, bound to a specific deployed contract.
func NewUSPToken(address common.Address, backend bind.ContractBackend) (*USPToken, error) {
	contract, err := bindUSPToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &USPToken{USPTokenCaller: USPTokenCaller{contract: contract}, USPTokenTransactor: USPTokenTransactor{contract: contract}, USPTokenFilterer: USPTokenFilterer{contract: contract}}, nil
}

// NewUSPTokenCaller creates a new read-only instance of USPToken, bound to a specific deployed contract.
func NewUSPTokenCaller(address common.Address, caller bind.ContractCaller) (*USPTokenCaller, error) {
	contract, err := bindUSPToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &USPTokenCaller{contract: contract}, nil
}

// NewUSPTokenTransactor creates a new write-only instance of USPToken, bound to a specific deployed contract.
func NewUSPTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*USPTokenTransactor, error) {
	contract, err := bindUSPToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &USPTokenTransactor{contract: contract}, nil
}

// NewUSPTokenFilterer creates a new log filterer instance of USPToken, bound to a specific deployed contract.
func NewUSPTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*USPTokenFilterer, error) {
	contract, err := bindUSPToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &USPTokenFilterer{contract: contract}, nil
}

// bindUSPToken binds a generic wrapper to an already deployed contract.
func bindUSPToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := USPTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_USPToken *USPTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USPToken.Contract.USPTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_USPToken *USPTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USPToken.Contract.USPTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_USPToken *USPTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USPToken.Contract.USPTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_USPToken *USPTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USPToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_USPToken *USPTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USPToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_USPToken *USPTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USPToken.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_USPToken *USPTokenCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _USPToken.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_USPToken *USPTokenSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _USPToken.Contract.DEFAULTADMINROLE(&_USPToken.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_USPToken *USPTokenCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _USPToken.Contract.DEFAULTADMINROLE(&_USPToken.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_USPToken *USPTokenCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _USPToken.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_USPToken *USPTokenSession) MINTERROLE() ([32]byte, error) {
	return _USPToken.Contract.MINTERROLE(&_USPToken.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_USPToken *USPTokenCallerSession) MINTERROLE() ([32]byte, error) {
	return _USPToken.Contract.MINTERROLE(&_USPToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_USPToken *USPTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _USPToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_USPToken *USPTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _USPToken.Contract.Allowance(&_USPToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_USPToken *USPTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _USPToken.Contract.Allowance(&_USPToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_USPToken *USPTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _USPToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_USPToken *USPTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _USPToken.Contract.BalanceOf(&_USPToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_USPToken *USPTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _USPToken.Contract.BalanceOf(&_USPToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_USPToken *USPTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _USPToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_USPToken *USPTokenSession) Decimals() (uint8, error) {
	return _USPToken.Contract.Decimals(&_USPToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_USPToken *USPTokenCallerSession) Decimals() (uint8, error) {
	return _USPToken.Contract.Decimals(&_USPToken.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_USPToken *USPTokenCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _USPToken.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_USPToken *USPTokenSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _USPToken.Contract.GetRoleAdmin(&_USPToken.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_USPToken *USPTokenCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _USPToken.Contract.GetRoleAdmin(&_USPToken.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_USPToken *USPTokenCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _USPToken.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_USPToken *USPTokenSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _USPToken.Contract.HasRole(&_USPToken.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_USPToken *USPTokenCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _USPToken.Contract.HasRole(&_USPToken.CallOpts, role, account)
}

// IdentityRegistry is a free data retrieval call binding the contract method 0x134e18f4.
//
// Solidity: function identityRegistry() view returns(address)
func (_USPToken *USPTokenCaller) IdentityRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USPToken.contract.Call(opts, &out, "identityRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IdentityRegistry is a free data retrieval call binding the contract method 0x134e18f4.
//
// Solidity: function identityRegistry() view returns(address)
func (_USPToken *USPTokenSession) IdentityRegistry() (common.Address, error) {
	return _USPToken.Contract.IdentityRegistry(&_USPToken.CallOpts)
}

// IdentityRegistry is a free data retrieval call binding the contract method 0x134e18f4.
//
// Solidity: function identityRegistry() view returns(address)
func (_USPToken *USPTokenCallerSession) IdentityRegistry() (common.Address, error) {
	return _USPToken.Contract.IdentityRegistry(&_USPToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_USPToken *USPTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _USPToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_USPToken *USPTokenSession) Name() (string, error) {
	return _USPToken.Contract.Name(&_USPToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_USPToken *USPTokenCallerSession) Name() (string, error) {
	return _USPToken.Contract.Name(&_USPToken.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_USPToken *USPTokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _USPToken.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_USPToken *USPTokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _USPToken.Contract.SupportsInterface(&_USPToken.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_USPToken *USPTokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _USPToken.Contract.SupportsInterface(&_USPToken.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_USPToken *USPTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _USPToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_USPToken *USPTokenSession) Symbol() (string, error) {
	return _USPToken.Contract.Symbol(&_USPToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_USPToken *USPTokenCallerSession) Symbol() (string, error) {
	return _USPToken.Contract.Symbol(&_USPToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_USPToken *USPTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _USPToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_USPToken *USPTokenSession) TotalSupply() (*big.Int, error) {
	return _USPToken.Contract.TotalSupply(&_USPToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_USPToken *USPTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _USPToken.Contract.TotalSupply(&_USPToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_USPToken *USPTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _USPToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_USPToken *USPTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _USPToken.Contract.Approve(&_USPToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_USPToken *USPTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _USPToken.Contract.Approve(&_USPToken.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_USPToken *USPTokenTransactor) Burn(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _USPToken.contract.Transact(opts, "burn", value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_USPToken *USPTokenSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _USPToken.Contract.Burn(&_USPToken.TransactOpts, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_USPToken *USPTokenTransactorSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _USPToken.Contract.Burn(&_USPToken.TransactOpts, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_USPToken *USPTokenTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, value *big.Int) (*types.Transaction, error) {
	return _USPToken.contract.Transact(opts, "burnFrom", account, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_USPToken *USPTokenSession) BurnFrom(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _USPToken.Contract.BurnFrom(&_USPToken.TransactOpts, account, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_USPToken *USPTokenTransactorSession) BurnFrom(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _USPToken.Contract.BurnFrom(&_USPToken.TransactOpts, account, value)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_USPToken *USPTokenTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _USPToken.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_USPToken *USPTokenSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _USPToken.Contract.GrantRole(&_USPToken.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_USPToken *USPTokenTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _USPToken.Contract.GrantRole(&_USPToken.TransactOpts, role, account)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_USPToken *USPTokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USPToken.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_USPToken *USPTokenSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USPToken.Contract.Mint(&_USPToken.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_USPToken *USPTokenTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USPToken.Contract.Mint(&_USPToken.TransactOpts, to, amount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_USPToken *USPTokenTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _USPToken.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_USPToken *USPTokenSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _USPToken.Contract.RenounceRole(&_USPToken.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_USPToken *USPTokenTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _USPToken.Contract.RenounceRole(&_USPToken.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_USPToken *USPTokenTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _USPToken.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_USPToken *USPTokenSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _USPToken.Contract.RevokeRole(&_USPToken.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_USPToken *USPTokenTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _USPToken.Contract.RevokeRole(&_USPToken.TransactOpts, role, account)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_USPToken *USPTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _USPToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_USPToken *USPTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _USPToken.Contract.Transfer(&_USPToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_USPToken *USPTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _USPToken.Contract.Transfer(&_USPToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_USPToken *USPTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _USPToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_USPToken *USPTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _USPToken.Contract.TransferFrom(&_USPToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_USPToken *USPTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _USPToken.Contract.TransferFrom(&_USPToken.TransactOpts, from, to, value)
}

// USPTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the USPToken contract.
type USPTokenApprovalIterator struct {
	Event *USPTokenApproval // Event containing the contract specifics and raw log

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
func (it *USPTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USPTokenApproval)
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
		it.Event = new(USPTokenApproval)
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
func (it *USPTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USPTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USPTokenApproval represents a Approval event raised by the USPToken contract.
type USPTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_USPToken *USPTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*USPTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _USPToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &USPTokenApprovalIterator{contract: _USPToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_USPToken *USPTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *USPTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _USPToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USPTokenApproval)
				if err := _USPToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_USPToken *USPTokenFilterer) ParseApproval(log types.Log) (*USPTokenApproval, error) {
	event := new(USPTokenApproval)
	if err := _USPToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// USPTokenRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the USPToken contract.
type USPTokenRoleAdminChangedIterator struct {
	Event *USPTokenRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *USPTokenRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USPTokenRoleAdminChanged)
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
		it.Event = new(USPTokenRoleAdminChanged)
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
func (it *USPTokenRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USPTokenRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USPTokenRoleAdminChanged represents a RoleAdminChanged event raised by the USPToken contract.
type USPTokenRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_USPToken *USPTokenFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*USPTokenRoleAdminChangedIterator, error) {

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

	logs, sub, err := _USPToken.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &USPTokenRoleAdminChangedIterator{contract: _USPToken.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_USPToken *USPTokenFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *USPTokenRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _USPToken.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USPTokenRoleAdminChanged)
				if err := _USPToken.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_USPToken *USPTokenFilterer) ParseRoleAdminChanged(log types.Log) (*USPTokenRoleAdminChanged, error) {
	event := new(USPTokenRoleAdminChanged)
	if err := _USPToken.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// USPTokenRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the USPToken contract.
type USPTokenRoleGrantedIterator struct {
	Event *USPTokenRoleGranted // Event containing the contract specifics and raw log

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
func (it *USPTokenRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USPTokenRoleGranted)
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
		it.Event = new(USPTokenRoleGranted)
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
func (it *USPTokenRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USPTokenRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USPTokenRoleGranted represents a RoleGranted event raised by the USPToken contract.
type USPTokenRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_USPToken *USPTokenFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*USPTokenRoleGrantedIterator, error) {

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

	logs, sub, err := _USPToken.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &USPTokenRoleGrantedIterator{contract: _USPToken.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_USPToken *USPTokenFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *USPTokenRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _USPToken.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USPTokenRoleGranted)
				if err := _USPToken.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_USPToken *USPTokenFilterer) ParseRoleGranted(log types.Log) (*USPTokenRoleGranted, error) {
	event := new(USPTokenRoleGranted)
	if err := _USPToken.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// USPTokenRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the USPToken contract.
type USPTokenRoleRevokedIterator struct {
	Event *USPTokenRoleRevoked // Event containing the contract specifics and raw log

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
func (it *USPTokenRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USPTokenRoleRevoked)
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
		it.Event = new(USPTokenRoleRevoked)
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
func (it *USPTokenRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USPTokenRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USPTokenRoleRevoked represents a RoleRevoked event raised by the USPToken contract.
type USPTokenRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_USPToken *USPTokenFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*USPTokenRoleRevokedIterator, error) {

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

	logs, sub, err := _USPToken.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &USPTokenRoleRevokedIterator{contract: _USPToken.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_USPToken *USPTokenFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *USPTokenRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _USPToken.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USPTokenRoleRevoked)
				if err := _USPToken.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_USPToken *USPTokenFilterer) ParseRoleRevoked(log types.Log) (*USPTokenRoleRevoked, error) {
	event := new(USPTokenRoleRevoked)
	if err := _USPToken.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// USPTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the USPToken contract.
type USPTokenTransferIterator struct {
	Event *USPTokenTransfer // Event containing the contract specifics and raw log

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
func (it *USPTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USPTokenTransfer)
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
		it.Event = new(USPTokenTransfer)
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
func (it *USPTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USPTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USPTokenTransfer represents a Transfer event raised by the USPToken contract.
type USPTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_USPToken *USPTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*USPTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USPToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &USPTokenTransferIterator{contract: _USPToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_USPToken *USPTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *USPTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USPToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USPTokenTransfer)
				if err := _USPToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_USPToken *USPTokenFilterer) ParseTransfer(log types.Log) (*USPTokenTransfer, error) {
	event := new(USPTokenTransfer)
	if err := _USPToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
