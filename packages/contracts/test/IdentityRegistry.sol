// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

// test/IdentityRegistry.t.sol
import "forge-std/Test.sol";
import "../src/IdentityRegistry.sol";

contract IdentityRegistryTest is Test {
    IdentityRegistry registry;
    address admin = address(1);
    address student = address(2);

    function setUp() public {
        registry = new IdentityRegistry(admin); // [cite: 51]
    }

    function test_AddStudent() public {
        vm.prank(admin);
        registry.addStudent(student, "2024123"); // 
        assertTrue(registry.isStudentActive(student)); // [cite: 55]
    }

    function test_FailAddStudentNonAdmin() public {
        vm.prank(student);
        vm.expectRevert(); 
        registry.addStudent(student, "2024123"); // 
    }
}