// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/USPToken.sol";
import "../src/IdentityRegistry.sol";

contract USPTokenTest is Test {
    USPToken public token;
    IdentityRegistry public registry;
    address public admin = address(1);
    address public studentA = address(2);
    address public studentB = address(3); 

    /// @dev Configuração inicial antes de cada teste
    function setUp() public {
        vm.startPrank(admin);

        registry = new IdentityRegistry(admin);
        token = new USPToken(address(registry), admin);
        
        // Adiciona estudantes
        registry.addStudent(studentA, "A");
        registry.addStudent(studentB, "B");

        // Desativa B
        registry.setStudentStatus(studentB, false); 
        
        // Emite tokens para A
        token.mint(studentA, 1000);

        vm.stopPrank();
    }

    // Testa transferência entre estudantes ativos
    function test_Transfer_ActiveToActive_Success() public {
        vm.startPrank(admin);

        // Ativa B
        registry.setStudentStatus(studentB, true); 
        vm.stopPrank();

        // Transferência de A para B
        vm.prank(studentA);
        bool success = token.transfer(studentB, 100);
        assertTrue(success);
        assertEq(token.balanceOf(studentB), 100);
    }

    // Testa transferência de estudante ativo para inativo
    function test_Transfer_ToInactive_Revert() public {
        vm.prank(studentA);
        vm.expectRevert("Destinatario inativo no registro");
        token.transfer(studentB, 100);
    }

    // Testa transferência de estudante inativo 
    function test_Transfer_FromInactive_Revert() public {
        vm.prank(studentB);
        vm.expectRevert("Remetente inativo no registro");
        token.transfer(studentA, 100);
    }

    // Testa restrição de mintagem de tokens apenas por minter autorizado
    function test_Mint_OnlyMinter() public {
        vm.prank(studentA);
        vm.expectRevert(); // AccessControl error
        token.mint(studentA, 100);
    }
}