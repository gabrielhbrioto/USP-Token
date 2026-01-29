// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "./IIdentityRegistry.sol";

contract USPToken is ERC20, ERC20Burnable, AccessControl {
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");
    IIdentityRegistry public identityRegistry;

    constructor(address _identityRegistry, address admin) ERC20("USP Token", "U$PT") {
        identityRegistry = IIdentityRegistry(_identityRegistry);
        _grantRole(DEFAULT_ADMIN_ROLE, admin);
        _grantRole(MINTER_ROLE, admin);
    }

    // Função para emitir tokens como recompensa
    function mint(address to, uint256 amount) public onlyRole(MINTER_ROLE) {
        require(identityRegistry.isStudentActive(to), "Destinatario nao e um estudante ativo");
        _mint(to, amount);
    }

    // Sobrescrita da funcao de transferencia para validar o registro de identidade
    function _update(address from, address to, uint256 value) internal override {
        // Ignora validacao se for emissao (mint) ou queima (burn)
        if (from != address(0) && to != address(0)) {
            require(identityRegistry.isStudentActive(from), "Remetente inativo no registro");
            require(identityRegistry.isStudentActive(to), "Destinatario inativo no registro");
        }
        super._update(from, to, value);
    }
}