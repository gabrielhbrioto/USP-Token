// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
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

    // Configuração inicial antes de cada teste
    function setUp() public {
        vm.startPrank(admin);

        // Implanta contratos
        registry = new IdentityRegistry(admin);
        token = new USPToken(address(registry), admin);
        cert = new USPCertificate(address(registry), address(token), admin);
        
        // Configura permissões
        token.grantRole(keccak256("MINTER_ROLE"), admin);
        token.grantRole(token.DEFAULT_ADMIN_ROLE(), address(cert));
        cert.grantRole(keccak256("GOVERNANCE_ROLE"), governance);
        
        // Adiciona aluno ativo e fornece tokens
        registry.addStudent(student, "123");
        token.mint(student, 500 * 10**18);

        vm.stopPrank();
    }

    // Testa resgate de certificado
    function test_RedeemCertificate_BurnsTokens() public {
        vm.startPrank(student);

        // Resgata certificado
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

    // Testa alteração de custo do certificado
    function test_SetCost_OnlyGovernance() public {
        vm.startPrank(governance);

        // Altera custo com papel correto
        cert.setCertificateCost(200 * 10**18);
        assertEq(cert.certificateCost(), 200 * 10**18);
        vm.stopPrank();

        // Tenta com admin normal (deve falhar pois admin perdeu papel ou nunca teve essa role especifica se configurado assim, mas aqui admin tem DEFAULT_ADMIN, entao cuidado com hierarquia)        
        vm.prank(student);
        vm.expectRevert();
        cert.setCertificateCost(0);
    }

    // Testa resgate por aluno inativo
    function test_RedeemCertificate_FailsIfNotActiveStudent() public {
        // Torna aluno inativo
        vm.startPrank(admin);
        registry.setStudentStatus(student, false);
        vm.stopPrank();

        // Tenta resgatar certificado
        vm.startPrank(student);
        token.approve(address(cert), 100 * 10**18);

        // Use expectRevert() genérico se não tiver certeza da string de revert
        vm.expectRevert("Apenas alunos ativos podem resgatar");
        cert.redeemCertificate("ipfs://metadata");
        vm.stopPrank();

        // Verifica saldo de tokens não consumido
        assertEq(token.balanceOf(student), 500 * 10**18);

        // Garante que nenhum certificado foi emitido
        assertEq(cert.balanceOf(student), 0);
    }

    function test_SystemMintCertificate_BurnsTokensAndSetsURI() public {
        vm.prank(admin);
        cert.systemMintCertificate(student, "ipfs://system-metadata");

        assertEq(token.balanceOf(student), 400 * 10**18);
        assertEq(cert.ownerOf(0), student);
        assertEq(cert.tokenURI(0), "ipfs://system-metadata");
        assertTrue(cert.supportsInterface(type(IERC5192).interfaceId));
        assertTrue(cert.supportsInterface(type(IERC721).interfaceId));
    }

    function test_SystemMintCertificate_FailsForInactiveStudent() public {
        vm.startPrank(admin);
        registry.setStudentStatus(student, false);
        vm.expectRevert("Apenas alunos ativos podem receber");
        cert.systemMintCertificate(student, "ipfs://system-metadata");
        vm.stopPrank();
    }

    function test_Locked_RevertsForMissingToken() public {
        vm.expectRevert("Token nao existe");
        cert.locked(999);
    }

    function test_Transfer_RevertsWhenTokenIsSoulbound() public {
        vm.prank(student);
        token.approve(address(cert), 100 * 10**18);
        vm.prank(student);
        cert.redeemCertificate("ipfs://metadata");

        vm.prank(student);
        vm.expectRevert("Tokens Soulbound nao podem ser transferidos");
        cert.transferFrom(student, address(3), 0);
    }
}