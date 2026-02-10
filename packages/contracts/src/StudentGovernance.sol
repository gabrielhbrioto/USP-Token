// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

interface IIdentityRegistryGov {
    function isStudentValidForSnapshot(address student, uint256 snapshotBlock) external view returns (bool);
    function activeStudentCount() external view returns (uint256);
}

interface IGovernanceToken is IERC20 {
    function mint(address to, uint256 amount) external;
    function burn(uint256 amount) external;
}

contract StudentGovernance is AccessControl, ReentrancyGuard {
    IIdentityRegistryGov public identityRegistry;
    IGovernanceToken public token;

    // --- Parâmetros de Governança (Variáveis) ---
    uint256 public quorumPercentage = 15; // Ex: 15%
    uint256 public votingPeriod = 2 days; // Duração padrão
    uint256 public voteReward = 5 * 10**18; // 5 Tokens por voto
    uint256 public proposalCost = 50 * 10**18;
    
    struct Proposal {
        address target;         // Contrato a ser alterado (Token, Certificado, Paymaster, ou Governance)
        bytes callData;         // A função e argumentos a executar (ex: setCertificateCost(200))
        string description;
        uint256 snapshotBlock;  // Bloco onde a proposta foi criada
        uint256 votingDeadline;
        uint256 yesVotes;
        uint256 noVotes;
        uint256 totalVotes;
        bool executed;
        address proposer;      
        bool depositReturned;  
        mapping(address => bool) hasVoted;
    }

    mapping(uint256 => Proposal) public proposals;
    uint256 public nextProposalId;

    event ProposalCreated(uint256 indexed id, string description, uint256 deadline);
    event VoteCast(address indexed voter, uint256 indexed proposalId, bool support);
    event ProposalExecuted(uint256 indexed id);
    event GovernanceParamsUpdated(uint256 newQuorum, uint256 newPeriod, uint256 newReward);
    event ProposalRejected(uint256 indexed id, bool quorumMet, bool approved);

    constructor(address _identityRegistry, address _token, address _admin) {
        identityRegistry = IIdentityRegistryGov(_identityRegistry);
        token = IGovernanceToken(_token);
        _grantRole(DEFAULT_ADMIN_ROLE, _admin);
    }

    function propose(address _target, bytes memory _callData, string memory _description) external {
        require(identityRegistry.isStudentValidForSnapshot(msg.sender, block.number), "Apenas alunos ativos");

        // COBRANÇA DA CAUÇÃO
        // O aluno precisa ter dado 'approve' no token para o contrato de governança antes
        require(token.transferFrom(msg.sender, address(this), proposalCost), "Falha no deposito da caucao");

        Proposal storage p = proposals[nextProposalId];
        p.target = _target;
        p.callData = _callData;
        p.description = _description;
        p.snapshotBlock = block.number;
        p.votingDeadline = block.timestamp + votingPeriod;
        p.proposer = msg.sender; // Salva o dono da proposta

        emit ProposalCreated(nextProposalId, _description, p.votingDeadline);
        nextProposalId++;
    }

    function castVote(uint256 proposalId, bool support) external nonReentrant {
        Proposal storage p = proposals[proposalId];
        
        require(block.timestamp < p.votingDeadline, "Votacao encerrada");
        require(!p.hasVoted[msg.sender], "Ja votou");
        
        // --- LÓGICA DE SNAPSHOT ---
        // Verifica se o aluno já existia quando a proposta foi criada
        require(
            identityRegistry.isStudentValidForSnapshot(msg.sender, p.snapshotBlock), 
            "Elegibilidade invalida para este snapshot"
        );

        p.hasVoted[msg.sender] = true;
        p.totalVotes++;

        if (support) {
            p.yesVotes++;
        } else {
            p.noVotes++;
        }

        // --- RECOMPENSA DE VOTO (Vote-to-Earn) ---
        // Governance deve ter MINTER_ROLE no USPToken
        token.mint(msg.sender, voteReward);

        emit VoteCast(msg.sender, proposalId, support);
    }

    function execute(uint256 proposalId) external nonReentrant {
        Proposal storage p = proposals[proposalId];
        
        // 1. Validações de Tempo e Estado (Essas AINDA devem reverter se falharem)
        require(block.timestamp >= p.votingDeadline, "Votacao em andamento");
        require(!p.executed, "Ja processada");

        p.executed = true; // Marca como processada para evitar reentrância/duplicação

        // 2. Cálculos de Aprovação
        uint256 totalStudents = identityRegistry.activeStudentCount();
        uint256 requiredVotes = (totalStudents * quorumPercentage) / 100;
        if (requiredVotes < 3) requiredVotes = 3;

        bool quorumMet = p.totalVotes >= requiredVotes;
        bool approved = p.yesVotes > p.noVotes;

        // 3. O Grande IF/ELSE (A Mágica acontece aqui)
        if (quorumMet && approved) {
            // --- CENÁRIO DE SUCESSO ---
            
            // Devolve a caução
            if (!p.depositReturned) {
                p.depositReturned = true;
                token.transfer(p.proposer, proposalCost); // Devolve
            }

            // Executa a ação
            (bool success, ) = p.target.call(p.callData);
            require(success, "Falha na chamada externa"); // Este revert é ok, pois é falha técnica
            
            emit ProposalExecuted(proposalId);

        } else {
            // --- CENÁRIO DE FRACASSO (Spam, rejeição ou falta de quórum) ---
            
            // A caução é queimada (ou enviada para o tesouro)
            p.depositReturned = true; // Marca como "lidada"
            
            // Tenta queimar (usando try/catch caso o token não suporte burn ou falhe)
            try token.burn(proposalCost) {
                // Queimou com sucesso
            } catch {
                // Se falhar o burn, envia para endereço morto para garantir punição
                token.transfer(address(0xdead), proposalCost);
            }
            
            // Não emite ProposalExecuted, pois falhou.
            // Pode emitir um evento de falha se quiser:
            emit ProposalRejected(proposalId, quorumMet, approved);
        }
    }

    // Função para alterar os parâmetros da PRÓPRIA governança via proposta
    // Só pode ser chamada pelo próprio contrato (via execute)
    function setGovernanceParams(uint256 _quorum, uint256 _period, uint256 _reward) external {
        require(msg.sender == address(this), "Apenas via Governanca");
        quorumPercentage = _quorum;
        votingPeriod = _period;
        voteReward = _reward;
        emit GovernanceParamsUpdated(_quorum, _period, _reward);
    }
}