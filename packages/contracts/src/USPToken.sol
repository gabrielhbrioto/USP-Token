// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

// @title USPToken
// @author Gabriel Henrique Brioto
// @notice Token ERC20 para alunos, integrando com o IdentityRegistry para validação de status
// @dev Este contrato implementa um token ERC20 para alunos, onde apenas estudantes ativos

// Importações necessárias
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "./IIdentityRegistry.sol";

contract USPToken is ERC20, ERC20Burnable, AccessControl {

    // Definição de papel para minters
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");
    IIdentityRegistry public identityRegistry;

    // Construtor que inicializa a referência ao IdentityRegistry e concede os papéis de administrador e minter
    constructor(address _identityRegistry, address admin) ERC20("USP Token", "U$PT") {
        identityRegistry = IIdentityRegistry(_identityRegistry);
        _grantRole(DEFAULT_ADMIN_ROLE, admin);
        _grantRole(MINTER_ROLE, admin);
    }

    /**
     * @dev Emite novos tokens para um endereço específico.
     * @param to Endereço do destinatário
     * @param amount Quantidade de tokens a serem emitidos
     */
    function mint(address to, uint256 amount) public onlyRole(MINTER_ROLE) {
        require(identityRegistry.isStudentActive(to), "Destinatario nao e um estudante ativo");
        _mint(to, amount);
    }

    /**
     * @dev Sobrescreve a função de transferência para incluir validação de status dos estudantes.
     * @param from Endereço do remetente
     * @param to Endereço do destinatário
     * @param value Quantidade de tokens a serem transferidos
     */
    function _update(address from, address to, uint256 value) internal override {
        // Ignora validacao se for emissao (mint) ou queima (burn)
        if (from != address(0) && to != address(0)) {
            require(identityRegistry.isStudentActive(from), "Remetente inativo no registro");
            require(identityRegistry.isStudentActive(to), "Destinatario inativo no registro");
        }
        super._update(from, to, value);
    }
}