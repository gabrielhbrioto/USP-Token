// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/AccessControl.sol";

contract IdentityRegistry is AccessControl {
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

    struct Student {
        bool isActive;
        string studentId;
        uint256 registrationBlock; // Para Snapshot
    }

    mapping(address => Student) public students;
    uint256 public activeStudentCount; // Para cálculo de quórum dinâmico

    event StudentAdded(address indexed studentAddress, string studentId);
    event StudentStatusChanged(address indexed studentAddress, bool isActive);

    constructor(address admin) {
        _grantRole(DEFAULT_ADMIN_ROLE, admin);
        _grantRole(ADMIN_ROLE, admin);
    }

    function addStudent(address studentAddress, string memory studentId) public onlyRole(ADMIN_ROLE) {
        require(bytes(students[studentAddress].studentId).length == 0, "Ja registrado");
        
        students[studentAddress] = Student({
            isActive: true,
            studentId: studentId,
            registrationBlock: block.number // Registra o bloco atual
        });
        
        activeStudentCount++; // Incrementa contador
        emit StudentAdded(studentAddress, studentId);
    }

    function setStudentStatus(address studentAddress, bool status) public onlyRole(ADMIN_ROLE) {
        require(bytes(students[studentAddress].studentId).length > 0, "Nao registrado");
        
        bool currentStatus = students[studentAddress].isActive;
        if (currentStatus != status) {
            students[studentAddress].isActive = status;
            if (status) {
                activeStudentCount++;
            } else {
                // Evita underflow por segurança
                if (activeStudentCount > 0) activeStudentCount--;
            }
            emit StudentStatusChanged(studentAddress, status);
        }
    }

    function isStudentActive(address studentAddress) public view returns (bool) {
        return students[studentAddress].isActive;
    }

    // Novo: Verifica se era ativo E se já existia no bloco do snapshot
    function isStudentValidForSnapshot(address student, uint256 snapshotBlock) external view returns (bool) {
        Student memory s = students[student];
        if (s.registrationBlock == 0) return false;
        return s.isActive && s.registrationBlock <= snapshotBlock;
    }
}