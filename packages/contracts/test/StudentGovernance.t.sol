// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/StudentGovernance.sol";
import "../src/USPToken.sol";
import "../src/IdentityRegistry.sol";
import "../src/USPCertificate.sol";

contract StudentGovernanceTest is Test {
    StudentGovernance public governance;
    USPToken public token;
    IdentityRegistry public registry;
    USPCertificate public cert;

    address public admin = address(1);
    address public student1 = address(10);
    address public student2 = address(11);
    address public student3 = address(12);

    function setUp() public {
        vm.startPrank(admin);
        
        // 1. Deploy Base
        registry = new IdentityRegistry(admin);
        token = new USPToken(address(registry), admin);
        governance = new StudentGovernance(address(registry), address(token), admin);
        cert = new USPCertificate(address(registry), address(token), admin);

        // 2. Configura Permissões
        // Token precisa dar MINTER_ROLE para Governance (para Vote-to-Earn)
        token.grantRole(keccak256("MINTER_ROLE"), address(governance));
        
        // Certificado precisa dar GOVERNANCE_ROLE para Governance (para alterar custo)
        cert.grantRole(keccak256("GOVERNANCE_ROLE"), address(governance));

        // 3. Cadastra Alunos (Bloco 1)
        vm.roll(1); 
        registry.addStudent(student1, "S1");
        registry.addStudent(student2, "S2");
        registry.addStudent(student3, "S3"); // 3 alunos ativos

        vm.stopPrank();
    }

    function test_FullGovernanceCycle() public {
        // --- ETAPA 1: Criar Proposta ---
        vm.roll(10); // Garante bloco para snapshot
        
        bytes memory callData = abi.encodeWithSelector(
            USPCertificate.setCertificateCost.selector, 
            50 * 10**18
        );

        vm.prank(student1);
        governance.propose(address(cert), callData, "Baixar preco certificado");
        uint256 proposalId = 0;
        
        // --- ETAPA 2: Votação (DENTRO DO PRAZO) ---
        vm.warp(block.timestamp + 1 hours); // Avança 1 hora (seguro)

        // Voto Aluno 1
        vm.prank(student1);
        governance.castVote(proposalId, true);
        assertEq(token.balanceOf(student1), 5 * 10**18); // Verifica recompensa

        // Voto Aluno 2
        vm.prank(student2);
        governance.castVote(proposalId, true);

        // Voto Aluno 3 (CRÍTICO: Deve votar AGORA, antes do prazo expirar)
        // Necessário para atingir o mínimo de 3 votos da trava de segurança
        vm.prank(student3);
        governance.castVote(proposalId, true);

        // --- ETAPA 3: Tentativa de Execução Precoce ---
        vm.expectRevert("Votacao em andamento");
        governance.execute(proposalId);

        // --- ETAPA 4: Avanço no Tempo e Execução Final ---
        // Agora sim avançamos para depois do prazo (2 dias + 1 segundo do início)
        // Como já avançamos 1 hora antes, basta somar o restante ou definir absoluto
        vm.warp(block.timestamp + 2 days); 

        // Executa a proposta
        governance.execute(proposalId);
        
        // Validação Final: O custo no certificado mudou?
        assertEq(cert.certificateCost(), 50 * 10**18);
    }

    function test_Governance_Execution_Success() public {
         // Setup
        vm.roll(10);
        bytes memory callData = abi.encodeWithSelector(USPCertificate.setCertificateCost.selector, 50 * 10**18);
        vm.prank(student1);
        governance.propose(address(cert), callData, "Desc");

        // Votação (3 votos a favor para garantir quorum minimo de 3)
        vm.prank(student1); governance.castVote(0, true);
        vm.prank(student2); governance.castVote(0, true);
        vm.prank(student3); governance.castVote(0, true);

        // Avança tempo
        vm.warp(block.timestamp + 3 days);

        // Executa
        governance.execute(0);

        // Verifica se o custo mudou no contrato Certificate
        assertEq(cert.certificateCost(), 50 * 10**18);
    }
}