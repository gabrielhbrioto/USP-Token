// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/StudentGovernance.sol";
import "../src/USPToken.sol";
import "../src/IdentityRegistry.sol";
import "../src/USPCertificate.sol";

contract MockPaymaster {
    uint256 public maxGasPerStudent;
    function setMaxGasPerStudent(uint256 v) external { maxGasPerStudent = v; }
}

contract StudentGovernanceTest is Test {
    StudentGovernance public governance;
    USPToken public token;
    IdentityRegistry public registry;
    USPCertificate public cert;

    address public admin = address(1);
    address public student1 = address(10);
    address public student2 = address(11);
    address public student3 = address(12);

    uint256 public constant INITIAL_STUDENT_BALANCE = 50 * 10**18;
    uint256 public constant INITIAL_VOTE_REWARD = 5 * 10**18;
    uint256 public constant PROPOSAL_COST = 50 * 10**18;

    // Configuração inicial antes de cada teste
    function setUp() public {
        // Inicia como admin para o deploy
        vm.startPrank(admin);

        // 1. Implanta Contratos
        registry = new IdentityRegistry(admin);
        token = new USPToken(address(registry), admin);
        governance = new StudentGovernance(address(registry), address(token), admin);
        cert = new USPCertificate(address(registry), address(token), admin);

        // 2. Configura Permissões (CRÍTICO)
        
        // A) USPToken: Governance precisa de MINTER_ROLE para o Vote-to-Earn
        token.grantRole(token.MINTER_ROLE(), address(governance));
        
        // B) USPCertificate: Governance precisa de GOVERNANCE_ROLE para mudar custos
        cert.grantRole(keccak256("GOVERNANCE_ROLE"), address(governance));

        // C) IdentityRegistry: Governance precisa de ADMIN_ROLE (Solução para o erro 0xa4...)
        // Isso permite que a Governance interaja livremente com o Registry se necessário
        registry.grantRole(registry.ADMIN_ROLE(), address(governance));

        // 3. Registrar contratos como 'estudantes ativos' para permitir transferências de caução
        // (Proposal Cost exige transferFrom)
        registry.addStudent(address(governance), "GOV");
        registry.addStudent(address(cert), "CERT");

        // 4. Cadastra Alunos de Teste (Bloco 1)
        vm.roll(1); 
        registry.addStudent(student1, "S1");
        registry.addStudent(student2, "S2");
        registry.addStudent(student3, "S3"); 

        // 5. Garante saldo inicial para testes que não mintam explicitamente
        token.mint(student1, INITIAL_STUDENT_BALANCE);

        vm.stopPrank();
    }

    // Testa o ciclo completo de Governança: Proposta -> Votação -> Execução
    function test_FullGovernanceCycle() public {

        // Garante bloco para snapshot
        vm.roll(10); 
        
        // Proposta
        bytes memory callData = abi.encodeWithSelector(
            USPCertificate.setCertificateCost.selector, 
            50 * 10**18
        );

        // Proposta feita pelo Aluno 1 (garantir approve da caução)
        uint256 pcost = governance.proposalCost();
        vm.prank(student1);
        token.approve(address(governance), pcost);
        vm.prank(student1);
        governance.propose(address(cert), callData, "Baixar preco certificado");
        uint256 proposalId = 0;
        
        // Avança 1 hora (seguro)
        vm.warp(block.timestamp + 1 hours); 

        // Voto Aluno 1
        vm.prank(student1);
        governance.castVote(proposalId, true);
        assertEq(token.balanceOf(student1), INITIAL_VOTE_REWARD + INITIAL_STUDENT_BALANCE - PROPOSAL_COST); // Verifica recompensa

        // Voto Aluno 2
        vm.prank(student2);
        governance.castVote(proposalId, true);
        assertEq(token.balanceOf(student2), INITIAL_VOTE_REWARD); // Verifica recompensa

        // Voto Aluno 3 - Necessário para atingir o mínimo de 3 votos da trava de segurança
        vm.prank(student3);
        governance.castVote(proposalId, true);
        assertEq(token.balanceOf(student3), INITIAL_VOTE_REWARD); // Verifica recompensa

        // Tentativa de Execução Precoce 
        vm.expectRevert("Votacao em andamento");
        governance.execute(proposalId);

        // Avanço no Tempo e Execução Final 
        vm.warp(block.timestamp + 2 days); 

        // Executa a proposta
        governance.execute(proposalId);
        
        // Validação Final: O custo no certificado mudou?
        assertEq(cert.certificateCost(), 50 * 10**18);
    }

    // Testa se a execução da proposta funciona corretamente
    function test_Governance_Execution_Success() public {
        
        uint256 NEW_COST = 50 * 10**18;
        uint256 pcost = governance.proposalCost();
        
        // Garantir que proposer tem saldo e allowance
        vm.prank(admin);
        token.mint(student1, pcost);
        vm.prank(student1);
        token.approve(address(governance), pcost);

        // Garantir que votantes estão ativos para o snapshot
        vm.roll(10);
        vm.prank(student1);
        governance.propose(address(cert), abi.encodeWithSelector(USPCertificate.setCertificateCost.selector, NEW_COST), "Mudar custo certificado");
        uint256 pid = 0;

        // Votos
        vm.prank(student1); governance.castVote(pid, true);
        vm.prank(student2); governance.castVote(pid, true);
        vm.prank(student3); governance.castVote(pid, true);

        // Execução após período
        vm.warp(block.timestamp + governance.votingPeriod() + 1);
        governance.execute(pid);

        // Verifica se o custo mudou no contrato Certificate
        assertEq(cert.certificateCost(), NEW_COST);
    }
 
    // Testa a alteração do custo do certificado via governança (sucesso)
    function test_CertificateCost_Change() public {
        uint256 NEW_COST = 123 * 10**18;
        uint256 pcost = governance.proposalCost();

        // Fornece saldo ao proposer e aprova a caução
        vm.prank(admin);
        token.mint(student1, pcost + 100 * 10**18);
        vm.prank(student1);
        token.approve(address(governance), pcost);

        // Define variáveis de estado
        uint256 initialBalance = token.balanceOf(student1);

        // Proposta
        vm.roll(10);
        vm.prank(student1);
        governance.propose(address(cert), abi.encodeWithSelector(USPCertificate.setCertificateCost.selector, NEW_COST), "Mudar custo certificado");
        uint256 pid = 0;

        // Votos
        vm.prank(student1); governance.castVote(pid, true);
        vm.prank(student2); governance.castVote(pid, true);
        vm.prank(student3); governance.castVote(pid, true);

        // Execução após período
        vm.warp(block.timestamp + governance.votingPeriod() + 1);
        governance.execute(pid);

        // Verificações
        assertEq(cert.certificateCost(), NEW_COST);

        // O proponente teve a caução devolvida e recebeu reward por votar
        uint256 balanceAfter = token.balanceOf(student1);
        assertEq(balanceAfter, initialBalance + governance.voteReward());
    }

    // Testa a falha na alteração do custo do certificado por falta de quorum
    function test_CertificateCost_FailsWithoutQuorum() public {
        uint256 NEW_COST = 321 * 10**18;
        uint256 pcost = governance.proposalCost();

        // Fornece saldo ao proposer e aprova a caução
        vm.prank(admin);
        token.mint(student1, pcost);
        vm.prank(student1);
        token.approve(address(governance), pcost);

        // Estado antes da proposta
        uint256 certCostBefore = cert.certificateCost();
        uint256 balanceBefore = token.balanceOf(student1);

        // Proposta
        vm.roll(11);
        vm.prank(student1);
        governance.propose(address(cert), abi.encodeWithSelector(USPCertificate.setCertificateCost.selector, NEW_COST), "Tentativa sem quorum");
        uint256 pid = 0;

        // Apenas 1 voto (insuficiente)
        vm.prank(student1); governance.castVote(pid, true);

        vm.warp(block.timestamp + governance.votingPeriod() + 1);
        governance.execute(pid);

        // custo do certificado nao mudou
        assertEq(cert.certificateCost(), certCostBefore);

        // caução nao foi devolvida (balance permanece menor)
        uint256 balanceAfter = token.balanceOf(student1);
        assertEq(balanceAfter, balanceBefore - pcost + governance.voteReward()); // votou e recebeu reward, mas caução consumida
    }

    // Testa a alteração da recompensa por voto via governança
    function test_VoteReward_Change() public {
        uint256 NEW_REWARD = 42 * 10**18;
        uint256 pcost = governance.proposalCost();

        // Fornece saldo ao proposer e aprova a caução
        vm.prank(admin);
        token.mint(student1, pcost + 100 * 10**18);
        vm.prank(student1);
        token.approve(address(governance), pcost);

        uint256 currentQuorum = governance.quorumPercentage();
        uint256 currentPeriod = governance.votingPeriod();

        // Proposta
        vm.roll(10);
        vm.prank(student1);
        governance.propose(address(governance), abi.encodeWithSelector(StudentGovernance.setGovernanceParams.selector, currentQuorum, currentPeriod, NEW_REWARD), "Mudar reward");
        uint256 pid = 0;

        // Votos
        vm.prank(student1); governance.castVote(pid, true);
        vm.prank(student2); governance.castVote(pid, true);
        vm.prank(student3); governance.castVote(pid, true);

        // Execução após período
        vm.warp(block.timestamp + governance.votingPeriod() + 1);
        governance.execute(pid);

        // Verificação da mudança
        assertEq(governance.voteReward(), NEW_REWARD);

        // Fornece saldo ao proposer e aprova a caução
        vm.prank(admin);
        token.mint(student1, pcost + 10 * 10**18);
        vm.prank(student1); 
        token.approve(address(governance), pcost);

        // Validar que novos votos mintam o novo reward
        vm.roll(20);
        vm.prank(student1);
        governance.propose(address(cert), abi.encodeWithSelector(USPCertificate.setCertificateCost.selector, 1 * 10**18), "p2");
        uint256 pid2 = 1;
        uint256 balanceBefore = token.balanceOf(student1);
        vm.prank(student1);
        governance.castVote(pid2, true);
        uint256 balanceAfter = token.balanceOf(student1);
        assertEq(balanceAfter - balanceBefore, NEW_REWARD);
    }

    // Testa a alteração da cota de gás por estudante (usa MockPaymaster)
    function test_GasQuotaPerStudent_Change() public {

        // Setup
        MockPaymaster pay = new MockPaymaster();
        // Registrar paymaster como ativo para que a caução possa ser transferida ao propor
        vm.prank(admin);
        registry.addStudent(address(pay), "PAY");
        uint256 NEW_GAS = 9999;
        uint256 pcost = governance.proposalCost();

        // Fornece saldo ao proposer e aprova a caução
        vm.prank(admin);
        token.mint(student1, pcost + 10 * 10**18);
        vm.prank(student1); 
        token.approve(address(governance), pcost);

        // Proposta 
        vm.roll(10);
        vm.prank(student1);
        governance.propose(address(pay), abi.encodeWithSelector(MockPaymaster.setMaxGasPerStudent.selector, NEW_GAS), "Ajustar gas");
        uint256 pid = 0;

        // Votos
        vm.prank(student1); governance.castVote(pid, true);
        vm.prank(student2); governance.castVote(pid, true);
        vm.prank(student3); governance.castVote(pid, true);

        // Execução após período
        vm.warp(block.timestamp + governance.votingPeriod() + 1);
        governance.execute(pid);

        // Verificação da mudança
        assertEq(pay.maxGasPerStudent(), NEW_GAS);
    }

    // Testa a alteração do quorumPercentage e validação funcional
    function test_QuorumPercentage_Change() public {

        vm.roll(10);

        // Cria mais estudantes para validação funcional mais eficiente
        address student4 = address(13);
        address student5 = address(14);
        address student6 = address(15);

        vm.startPrank(admin);
        registry.addStudent(student4, "S4");
        registry.addStudent(student5, "S5");
        registry.addStudent(student6, "S6");
        token.mint(student4, INITIAL_STUDENT_BALANCE);
        token.mint(student5, INITIAL_STUDENT_BALANCE);
        token.mint(student6, INITIAL_STUDENT_BALANCE);
        vm.stopPrank();

        // Setup
        uint256 NEW_QUORUM = 50;
        uint256 pcost = governance.proposalCost();

        // Fornece saldo ao proposer e aprova a caução
        vm.prank(admin);
        token.mint(student1, pcost + 10 * 10**18);
        vm.prank(student1); 
        token.approve(address(governance), pcost);

        uint256 currentPeriod = governance.votingPeriod();
        uint256 currentReward = governance.voteReward();

        // Proposta
        vm.roll(20);
        vm.prank(student1);
        governance.propose(address(governance), abi.encodeWithSelector(StudentGovernance.setGovernanceParams.selector, NEW_QUORUM, currentPeriod, currentReward), "Alterar quorum");
        uint256 pid = 0;

        // Votos
        vm.prank(student1); governance.castVote(pid, true);
        vm.prank(student2); governance.castVote(pid, true);
        vm.prank(student3); governance.castVote(pid, true);
        vm.prank(student4); governance.castVote(pid, true);
        vm.prank(student5); governance.castVote(pid, true);
        vm.prank(student6); governance.castVote(pid, true);

        // Execução após período
        vm.warp(block.timestamp + governance.votingPeriod() + 1);
        governance.execute(pid);

        // Verificação da mudança 
        assertEq(governance.quorumPercentage(), NEW_QUORUM);

        // Validação funcional: nova proposta exige mais votos
        uint256 total = registry.activeStudentCount();
        uint256 required = (total * NEW_QUORUM) / 100;
        if (required < 3) required = 3;

        // Fornece saldo ao proposer e aprova a caução
        vm.prank(admin);
        token.mint(student1, pcost + 10 * 10**18);
        vm.prank(student1); 
        token.approve(address(governance), pcost);

        // Cria proposta que deve ser rejeitada se votos < required
        vm.prank(student1);
        governance.propose(address(cert), abi.encodeWithSelector(USPCertificate.setCertificateCost.selector, 777 * 10**18), "p_quorum_test");
        uint256 pid2 = 1;

        // Vota com required-1 votos
        for (uint256 i = 0; i < required - 1; i++) {
            address voter = address(uint160(10 + i)); // student1, student2, ...
            vm.prank(voter);
            governance.castVote(pid2, true);
        }

        // Avança tempo e executa
        vm.warp(block.timestamp + governance.votingPeriod() + 1);
        governance.execute(pid2);

        // Verifica que o custo do certificado NÃO mudou (falta de quorum)
        assertTrue(cert.certificateCost() != 777 * 10**18);
    }

    // Testa a alteração do votingPeriod e validação funcional
    function test_VotingPeriod_Change() public {

        // Setup
        uint256 NEW_PERIOD = 1 days;
        uint256 pcost = governance.proposalCost();

        // Fornece saldo ao proposer e aprova a caução
        vm.prank(admin);
        token.mint(student1, pcost + 10 * 10**18);
        vm.prank(student1); 
        token.approve(address(governance), pcost);

        uint256 currentQuorum = governance.quorumPercentage();
        uint256 currentReward = governance.voteReward();

        vm.roll(10);
        vm.prank(student1);
        governance.propose(address(governance), abi.encodeWithSelector(StudentGovernance.setGovernanceParams.selector, currentQuorum, NEW_PERIOD, currentReward), "Alterar periodo");
        uint256 pid = 0;

        // Votos
        vm.prank(student1); governance.castVote(pid, true);
        vm.prank(student2); governance.castVote(pid, true);
        vm.prank(student3); governance.castVote(pid, true);

        // Execução após período
        vm.warp(block.timestamp + governance.votingPeriod() + 1);
        governance.execute(pid);

        // Verificação da mudança
        assertEq(governance.votingPeriod(), NEW_PERIOD);

        // Fornece saldo ao proposer e aprova a caução
        vm.prank(admin);
        token.mint(student1, pcost + 10 * 10**18);
        vm.prank(student1); 
        token.approve(address(governance), pcost);

        // Agora validar comportamento temporal com o novo periodo
        vm.prank(student1);
        governance.propose(address(cert), abi.encodeWithSelector(USPCertificate.setCertificateCost.selector, 999 * 10**18), "p_time");
        uint256 pid2 = 1;

        // Votos
        vm.prank(student1); governance.castVote(pid2, true);
        vm.prank(student2); governance.castVote(pid2, true);
        vm.prank(student3); governance.castVote(pid2, true);

        // tentativa precoce: antes do NEW_PERIOD
        vm.expectRevert("Votacao em andamento");
        governance.execute(pid2);

        // avanca o tempo e executa com sucesso
        vm.warp(block.timestamp + NEW_PERIOD + 1);
        governance.execute(pid2);

        // Verificação da mudança
        assertEq(cert.certificateCost(), 999 * 10**18);
    }

}