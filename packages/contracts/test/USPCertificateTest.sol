// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

// test/USPCertificate.t.sol
import "forge-std/Test.sol";
import "../src/USPCertificate.sol";
import "../src/IdentityRegistry.sol";
import "../src/USPToken.sol";

contract USPCertificateTest is Test {
    USPCertificate sbt;
    USPToken token;
    IdentityRegistry registry;
    address admin = address(1);
    address student = address(2);

    function setUp() public {
        registry = new IdentityRegistry(admin); // [cite: 51]
        token = new USPToken(address(registry), admin); // [cite: 4]
        sbt = new USPCertificate(address(registry), address(token), admin); // [cite: 33]

        vm.startPrank(admin);
        registry.addStudent(student, "S1"); // 
        token.mint(student, 200 * 10**18); // [cite: 5]
        vm.stopPrank();
    }

    function test_RedeemCertificate() public {
        vm.startPrank(student);
        token.approve(address(sbt), 100 * 10**18); // 
        sbt.redeemCertificate("ipfs://metadata"); // [cite: 36]
        vm.stopPrank();

        assertEq(sbt.balanceOf(student), 1); // [cite: 38]
        assertEq(token.balanceOf(student), 100 * 10**18); // 
    }

    function test_FailTransferSBT() public {
        vm.startPrank(student);
        token.approve(address(sbt), 100 * 10**18); // 
        sbt.redeemCertificate("ipfs://metadata"); // [cite: 36]

        vm.expectRevert("Tokens Soulbound nao podem ser transferidos");
        sbt.transferFrom(student, address(4), 0); // 
        vm.stopPrank();
    }
}