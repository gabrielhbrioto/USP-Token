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
    address public studentB = address(3); // Inativo

    function setUp() public {
        vm.startPrank(admin);
        registry = new IdentityRegistry(admin);
        token = new USPToken(address(registry), admin);
        
        registry.addStudent(studentA, "A");
        registry.addStudent(studentB, "B");
        registry.setStudentStatus(studentB, false); // Desativa B
        
        token.mint(studentA, 1000);
        vm.stopPrank();
    }

    function test_Transfer_ActiveToActive_Success() public {
        vm.startPrank(admin);
        registry.setStudentStatus(studentB, true); // Ativa B
        vm.stopPrank();

        vm.prank(studentA);
        bool success = token.transfer(studentB, 100);
        assertTrue(success);
        assertEq(token.balanceOf(studentB), 100);
    }

    function test_Transfer_ToInactive_Revert() public {
        vm.prank(studentA);
        vm.expectRevert("Destinatario inativo no registro");
        token.transfer(studentB, 100);
    }

    function test_Mint_OnlyMinter() public {
        vm.prank(studentA);
        vm.expectRevert(); // AccessControl error
        token.mint(studentA, 100);
    }
}