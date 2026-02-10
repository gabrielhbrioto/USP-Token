// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

// @title IdentityRegistry
// @author Gabriel Henrique Brioto
// @notice Gerencia o registro de estudantes, permitindo adição, ativação/desativação, consulta de status e possui integração com Snapshot para governança
// @dev Utiliza AccessControl para gerenciamento de permissões, garantindo que apenas administradores possam modificar o registro de estudantes. O contrato mantém um contador de estudantes ativos para facilitar o cálculo de quórum dinâmico na governança

import "@openzeppelin/contracts/access/AccessControl.sol";

contract IdentityRegistry is AccessControl {

    // Definição de papel para administradores
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

    // Estrutura para armazenar informações do estudante, incluindo o bloco de registro para integração com Snapshot
    struct Student {
        bool isActive; 
        string studentId;
        uint256 registrationBlock; 
    }

    mapping(address => Student) public students;
    uint256 public activeStudentCount; 

    // Eventos
    event StudentAdded(address indexed studentAddress, string studentId);
    event StudentStatusChanged(address indexed studentAddress, bool isActive);

    // Construtor que concede os papéis de administrador ao endereço fornecido
    constructor(address admin) {
        _grantRole(DEFAULT_ADMIN_ROLE, admin);
        _grantRole(ADMIN_ROLE, admin);
    }

    /**
     * @dev Adiciona um estudante ao registro.
     * @param studentAddress Endereço do estudante a ser adicionado
     * @param studentId Identificador do estudante (pode ser um número de matrícula ou outro ID)
     */
    function addStudent(address studentAddress, string memory studentId) public onlyRole(ADMIN_ROLE) {
        require(bytes(students[studentAddress].studentId).length == 0, "Ja registrado");
        
        students[studentAddress] = Student({
            isActive: true,
            studentId: studentId,
            registrationBlock: block.number
        });
        
        activeStudentCount++;
        emit StudentAdded(studentAddress, studentId);
    }

    /**
     * @dev Altera o status de um estudante (ativo/inativo).
     * @param studentAddress Endereço do estudante a ser alterado
     * @param status Novo status do estudante
     */
    function setStudentStatus(address studentAddress, bool status) public onlyRole(ADMIN_ROLE) {
        require(bytes(students[studentAddress].studentId).length > 0, "Nao registrado");
        
        bool currentStatus = students[studentAddress].isActive;
        if (currentStatus != status) {
            students[studentAddress].isActive = status;
            if (status) {
                activeStudentCount++;
            } else {
                if (activeStudentCount > 0) activeStudentCount--;
            }
            emit StudentStatusChanged(studentAddress, status);
        }
    }

    /**
     * @param studentAddress Endereço do estudante a ser verificado
     * @return bool Retorna true se o estudante estiver ativo, false caso contrário
     */
    function isStudentActive(address studentAddress) public view returns (bool) {
        return students[studentAddress].isActive;
    }

    /**
     * @param student Endereço do estudante a ser verificado
     * @param snapshotBlock Bloco do snapshot a ser verificado
     */
    function isStudentValidForSnapshot(address student, uint256 snapshotBlock) external view returns (bool) {
        Student memory s = students[student];
        if (s.registrationBlock == 0) return false;
        return s.isActive && s.registrationBlock <= snapshotBlock;
    }
}