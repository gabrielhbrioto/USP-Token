// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

// test/USPToken.t.sol
import "forge-std/Test.sol";
import "../src/USPToken.sol";
import "../src/IdentityRegistry.sol";

contract USPTokenTest is Test {
    USPToken token;
    IdentityRegistry registry;
    address admin = address(1);
    address studentA = address(2);
    address studentB = address(3);

    function setUp() public {
        registry = new IdentityRegistry(admin); // [cite: 51]
        token = new USPToken(address(registry), admin); // [cite: 4]
        
        vm.prank(admin);
        registry.addStudent(studentA, "A1"); // 
    }

    function test_MintToActiveStudent() public {
        vm.prank(admin);
        token.mint(studentA, 1000); // [cite: 5]
        assertEq(token.balanceOf(studentA), 1000);
    }

    function test_FailTransferToInactiveStudent() public {
        vm.prank(admin);
        token.mint(studentA, 1000); // [cite: 5]

        vm.prank(studentA);
        vm.expectRevert("Destinatario inativo no registro");
        token.transfer(studentB, 500); // 
    }
}