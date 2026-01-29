// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/AccessControl.sol";

contract IdentityRegistry is AccessControl {
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

    struct Student {
        bool isActive;
        string studentId; // Número USP
    }

    mapping(address => Student) public students;

    event StudentAdded(address indexed studentAddress, string studentId);
    event StudentStatusChanged(address indexed studentAddress, bool isActive);

    constructor(address admin) {
        _grantRole(DEFAULT_ADMIN_ROLE, admin);
        _grantRole(ADMIN_ROLE, admin);
    }

    function addStudent(address studentAddress, string memory studentId) public onlyRole(ADMIN_ROLE) {
        students[studentAddress] = Student(true, studentId);
        emit StudentAdded(studentAddress, studentId);
    }

    function setStudentStatus(address studentAddress, bool status) public onlyRole(ADMIN_ROLE) {
        require(bytes(students[studentAddress].studentId).length > 0, "Student not registered");
        students[studentAddress].isActive = status;
        emit StudentStatusChanged(studentAddress, status);
    }

    function isStudentActive(address studentAddress) public view returns (bool) {
        return students[studentAddress].isActive;
    }
}