// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@account-abstraction/contracts/interfaces/IPaymaster.sol";
import "@account-abstraction/contracts/interfaces/IEntryPoint.sol";

import "./IIdentityRegistry.sol";

contract USPPaymaster is IPaymaster, Ownable {
    IIdentityRegistry public identityRegistry;
    IEntryPoint public immutable entryPoint;
    
    mapping(address => uint256) public gasSpentByStudent;
    uint256 public constant MAX_GAS_PER_STUDENT = 0.05 ether;

    mapping(address => bool) public whitelistedTargets;

    event GasCharged(address indexed student, uint256 actualGasCost, PostOpMode mode);

    constructor(IEntryPoint _entryPoint, address _identityRegistry, address _admin) Ownable(_admin) {
        entryPoint = _entryPoint;
        identityRegistry = IIdentityRegistry(_identityRegistry);
    }

    /**
     * @dev Valida se o Paymaster aceita pagar pela operação.
     * Removido o 'view' para bater com a interface IPaymaster.
     */
    function validatePaymasterUserOp(
        PackedUserOperation calldata userOp,
        bytes32 /*userOpHash*/,
        uint256 maxCost
    ) external view override returns (bytes memory context, uint256 validationData) {
        
        // 1. Verificações de segurança
        if (!identityRegistry.isStudentActive(userOp.sender)) {
            return ("", 1); 
        }

        address target = _getTargetFromCallData(userOp.callData);
        if (!whitelistedTargets[target]) {
            return ("", 1);
        }

        if (entryPoint.balanceOf(address(this)) < maxCost) {
            return ("", 1);
        }

        // Retornamos o endereço do aluno para ser usado na postOp
        return (abi.encode(userOp.sender), 0); 
    }

    /**
     * @dev Função corrigida com os 4 parâmetros exigidos pela interface v0.7
     */
    function postOp(
        PostOpMode mode,
        bytes calldata context,
        uint256 actualGasCost,
        uint256 /* actualUserOpFeePerGas */ 
    ) external override {
        // Decodifica o endereço do aluno enviado pelo validatePaymasterUserOp
        address student = abi.decode(context, (address));

        // Registra o gasto real (em unidades de gás multiplicado pela taxa)
        // Você pode registrar apenas o actualGasCost ou o custo total em ETH
        gasSpentByStudent[student] += actualGasCost;

        emit GasCharged(student, actualGasCost, mode);
    }

    // --- Funções Administrativas ---

    function setTargetWhitelist(address target, bool status) external onlyOwner {
        whitelistedTargets[target] = status;
    }

    function make_deposit() external payable {
        entryPoint.depositTo{value: msg.value}(address(this));
    }

    function withdrawTo(address payable withdrawAddress, uint256 amount) external onlyOwner {
        entryPoint.withdrawTo(withdrawAddress, amount);
    }

    function _getTargetFromCallData(bytes calldata data) internal pure returns (address target) {
        // Para a maioria das Smart Accounts, o target está nos primeiros 20 bytes do callData
        if (data.length < 20) return address(0);
        return address(bytes20(data[0:20])); 
    }
}