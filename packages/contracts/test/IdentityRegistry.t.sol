// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/IdentityRegistry.sol";

contract IdentityRegistryTest is Test {
    IdentityRegistry public registry;
    address public admin = address(1);
    address public student = address(2);

    function setUp() public {
        vm.startPrank(admin);
        registry = new IdentityRegistry(admin);
        vm.stopPrank();
    }

    function test_AddStudent_IncrementsCountAndSetsBlock() public {
        vm.startPrank(admin);
        
        uint256 blockBefore = block.number;
        registry.addStudent(student, "12345");
        
        // Verifica contagem
        assertEq(registry.activeStudentCount(), 1);
        
        // Verifica dados da struct
        (bool isActive, string memory id, uint256 regBlock) = registry.students(student);
        assertTrue(isActive);
        assertEq(id, "12345");
        assertEq(regBlock, blockBefore); // Deve ser o bloco atual
        vm.stopPrank();
    }

    function test_SnapshotLogic() public {
        vm.startPrank(admin);
        
        // Bloco 10: Adiciona aluno
        vm.roll(10);
        registry.addStudent(student, "12345");
        
        // Bloco 20: Verifica snapshot
        vm.roll(20);
        
        // O aluno deve ser válido para um snapshot no bloco 10 ou maior
        assertTrue(registry.isStudentValidForSnapshot(student, 10));
        assertTrue(registry.isStudentValidForSnapshot(student, 15));
        
        // O aluno NÃO deve ser válido para um snapshot no bloco 5 (antes de entrar)
        assertFalse(registry.isStudentValidForSnapshot(student, 5));
        
        vm.stopPrank();
    }

    function test_SetStudentStatus_UpdatesCount() public {
        vm.startPrank(admin);
        registry.addStudent(student, "12345");
        assertEq(registry.activeStudentCount(), 1);

        registry.setStudentStatus(student, false);
        assertEq(registry.activeStudentCount(), 0);
        
        registry.setStudentStatus(student, true);
        assertEq(registry.activeStudentCount(), 1);
        vm.stopPrank();
    }
}