// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

// @title USPPaymaster
// @author Gabriel Henrique Brioto
// @notice Paymaster para alunos, integrando com o IdentityRegistry para validação e controle de gastos
// @dev Este contrato implementa um Paymaster para alunos, onde cada aluno tem

// Importações necessárias
import "@openzeppelin/contracts/access/Ownable.sol";
import "@account-abstraction/contracts/interfaces/IPaymaster.sol";
import "@account-abstraction/contracts/interfaces/IEntryPoint.sol";
import "./IIdentityRegistry.sol";

contract USPPaymaster is IPaymaster, Ownable {

    // Referências para o contrato de identidade e o EntryPoint
    IIdentityRegistry public identityRegistry;
    IEntryPoint public immutable entryPoint;
    
    // Variáveis para controle de gastos
    mapping(address => uint256) public gasSpentByStudent;
    uint256 public maxGasPerStudent = 0.05 ether;
    mapping(address => bool) public whitelistedTargets;

    // Eventos
    event GasCharged(address indexed student, uint256 actualGasCost, PostOpMode mode);
    event GasLimitUpdated(uint256 newLimit);

    // Construtor que inicializa as referências aos contratos e define o owner
    constructor(IEntryPoint _entryPoint, address _identityRegistry, address _admin) Ownable(_admin) {
        entryPoint = _entryPoint;
        identityRegistry = IIdentityRegistry(_identityRegistry);
    }

    /**
     * @dev Define o limite máximo de gás que um aluno pode gastar.
     * @param _newLimit Novo limite de gás
     */
    function setMaxGasPerStudent(uint256 _newLimit) external onlyOwner {
        maxGasPerStudent = _newLimit;
        emit GasLimitUpdated(_newLimit);
    }

    /**
     * @dev Adiciona ou remove um endereço da lista de whitelist.
     * @param target Endereço a ser adicionado ou removido
     * @param status Status (true para adicionar, false para remover)
     */
    function setTargetWhitelist(address target, bool status) external onlyOwner {
        whitelistedTargets[target] = status;
    }

    /**
     * @dev Valida uma operação de usuário para o Paymaster.
     * @param userOp Dados da operação do usuário
     * @param maxCost Custo máximo permitido
     */
    function validatePaymasterUserOp(
        PackedUserOperation calldata userOp,
        bytes32,
        uint256 maxCost
    ) external view override returns (bytes memory context, uint256 validationData) {

        // Verifica se o aluno está ativo
        if (!identityRegistry.isStudentActive(userOp.sender)) return ("", 1);

        // Verifica se o alvo está na whitelist
        address target = _getTargetFromCallData(userOp.callData);
        if (!whitelistedTargets[target]) return ("", 1);

        // Verifica contra a variável maxGasPerStudent
        if (gasSpentByStudent[userOp.sender] > maxGasPerStudent) return ("", 1); 

        // Verifica se o Paymaster tem fundos suficientes para cobrir o custo máximo
        if (entryPoint.balanceOf(address(this)) < maxCost) return ("", 1);

        // Se todas as verificações passarem, retorna o contexto e dados de validação
        return (abi.encode(userOp.sender), 0);
    }

    /**
     * @dev Função de pós-operação para cobrar o gás gasto.
     * @param mode Modo da operação (opReverted, opSucceeded, etc.)
     * @param context Contexto codificado (contém o endereço do aluno)
     * @param actualGasCost Custo real do gás gasto
     */
    function postOp(PostOpMode mode, bytes calldata context, uint256 actualGasCost, uint256) external override {
        address student = abi.decode(context, (address));
        gasSpentByStudent[student] += actualGasCost;
        emit GasCharged(student, actualGasCost, mode);
    }

    /**
     * @dev Faz um depósito no Paymaster.
     */
    function makeDeposit() external payable {
        entryPoint.depositTo{value: msg.value}(address(this));
    }

    /**
     * @dev Retira fundos do Paymaster para um endereço específico.
     * @param withdrawAddress Endereço para o qual os fundos serão retirados
     * @param amount Quantidade a ser retirada
     */
    function withdrawTo(address payable withdrawAddress, uint256 amount) external onlyOwner {
        entryPoint.withdrawTo(withdrawAddress, amount);
    }

    /** @dev Função interna para extrair o endereço do alvo a partir dos dados da chamada.
     * @param data Dados da chamada
     * @return target Endereço do alvo extraído
     */
    function _getTargetFromCallData(bytes calldata data) internal pure returns (address target) {
        if (data.length < 20) return address(0);
        return address(bytes20(data[0:20])); 
    }
}