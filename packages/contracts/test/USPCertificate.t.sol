// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/USPCertificate.sol";
import "../src/USPToken.sol";
import "../src/IdentityRegistry.sol";

contract USPCertificateTest is Test {
    USPCertificate public cert;
    USPToken public token;
    IdentityRegistry public registry;
    address public admin = address(1);
    address public governance = address(4); // Simula contrato de Gov
    address public student = address(2);

    function setUp() public {
        vm.startPrank(admin);
        registry = new IdentityRegistry(admin);
        token = new USPToken(address(registry), admin);
        cert = new USPCertificate(address(registry), address(token), admin);
        
        // Configura permissões
        token.grantRole(keccak256("MINTER_ROLE"), admin);
        cert.grantRole(keccak256("GOVERNANCE_ROLE"), governance);
        
        registry.addStudent(student, "123");
        token.mint(student, 500 * 10**18);
        vm.stopPrank();
    }

    function test_RedeemCertificate_BurnsTokens() public {
        vm.startPrank(student);
        token.approve(address(cert), 100 * 10**18);
        
        cert.redeemCertificate("ipfs://metadata");
        
        // Verifica saldo (tinha 500, custou 100, sobrou 400)
        assertEq(token.balanceOf(student), 400 * 10**18);
        // Verifica posse
        assertEq(cert.ownerOf(0), student);
        // Verifica lock (SBT)
        assertTrue(cert.locked(0));
        vm.stopPrank();
    }

    function test_SetCost_OnlyGovernance() public {
        vm.startPrank(governance);
        cert.setCertificateCost(200 * 10**18);
        assertEq(cert.certificateCost(), 200 * 10**18);
        vm.stopPrank();

        // Tenta com admin normal (deve falhar pois admin perdeu papel ou nunca teve essa role especifica se configurado assim, mas aqui admin tem DEFAULT_ADMIN, entao cuidado com hierarquia)
        // No setup eu dei GOVERNANCE_ROLE para governance. Admin tem admin role.
        
        vm.prank(student);
        vm.expectRevert();
        cert.setCertificateCost(0);
    }
}