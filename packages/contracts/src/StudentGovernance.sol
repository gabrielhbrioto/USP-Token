// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

interface IIdentityRegistryGov {
    function isStudentValidForSnapshot(address student, uint256 snapshotBlock) external view returns (bool);
    function activeStudentCount() external view returns (uint256);
}

interface IMintableToken {
    function mint(address to, uint256 amount) external;
}

contract StudentGovernance is AccessControl, ReentrancyGuard {
    IIdentityRegistryGov public identityRegistry;
    IMintableToken public token;

    // --- Parâmetros de Governança (Variáveis) ---
    uint256 public quorumPercentage = 15; // Ex: 15%
    uint256 public votingPeriod = 2 days; // Duração padrão
    uint256 public voteReward = 5 * 10**18; // 5 Tokens por voto

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
        mapping(address => bool) hasVoted;
    }

    mapping(uint256 => Proposal) public proposals;
    uint256 public nextProposalId;

    event ProposalCreated(uint256 indexed id, string description, uint256 deadline);
    event VoteCast(address indexed voter, uint256 indexed proposalId, bool support);
    event ProposalExecuted(uint256 indexed id);
    event GovernanceParamsUpdated(uint256 newQuorum, uint256 newPeriod, uint256 newReward);

    constructor(address _identityRegistry, address _token, address _admin) {
        identityRegistry = IIdentityRegistryGov(_identityRegistry);
        token = IMintableToken(_token);
        _grantRole(DEFAULT_ADMIN_ROLE, _admin);
    }

    // Cria uma proposta para alterar QUALQUER coisa no sistema
    function propose(
        address _target, 
        bytes memory _callData, 
        string memory _description
    ) external {
        // Valida snapshot simples (deve ser ativo agora para propor)
        require(identityRegistry.isStudentValidForSnapshot(msg.sender, block.number), "Apenas alunos ativos");

        Proposal storage p = proposals[nextProposalId];
        p.target = _target;
        p.callData = _callData;
        p.description = _description;
        p.snapshotBlock = block.number;
        p.votingDeadline = block.timestamp + votingPeriod;

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
        require(block.timestamp >= p.votingDeadline, "Votacao em andamento");
        require(!p.executed, "Ja executada");

        // --- CÁLCULO DE QUORUM ---
        uint256 totalStudents = identityRegistry.activeStudentCount();
        uint256 requiredVotes = (totalStudents * quorumPercentage) / 100;
        
        // Trava de segurança mínima (ex: min 3 votos sempre)
        if (requiredVotes < 3) requiredVotes = 3;

        require(p.totalVotes >= requiredVotes, "Quorum nao atingido");

        p.executed = true;

        if (p.yesVotes > p.noVotes) {
            // Executa a chamada no contrato alvo
            (bool success, ) = p.target.call(p.callData);
            require(success, "Falha na execucao da proposta");
            emit ProposalExecuted(proposalId);
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