// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

// @title IIdentityRegistry
// @author Gabriel Henrique Brioto
// @notice Interface para o contrato IdentityRegistry, definindo as funções essenciais para consulta de status
// @dev Esta interface é utilizada para garantir que outros contratos possam interagir com o IdentityRegistry de forma consistente, permitindo a verificação do status dos estudantes

interface IIdentityRegistry {
    function isStudentActive(address studentAddress) external view returns (bool);
}
