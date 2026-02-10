// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

// @title StudentGovernance
// @author Gabriel Henrique Brioto
// @notice Contrato de governança para alunos, permitindo a criação de propostas, votação e
// @dev Este contrato implementa um sistema de governança para alunos, onde eles podem criar propostas para alterar parâmetros do sistema (como custos de certificados, regras de paymaster, ou até mesmo parâmetros da própria governança). O contrato utiliza um sistema de caução para evitar spam, um mecanismo de votação baseado em snapshots para garantir que apenas alunos ativos possam votar, e um sistema de recompensa para incentivar a participação. A execução das propostas é feita através de chamadas externas, permitindo uma grande flexibilidade nas mudanças que podem ser propostas.

// importações necessárias
import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

/** @title IIdentityRegistryGov
 *  @author Gabriel Henrique Brioto
 *  @notice Interface para o contrato IdentityRegistryGov, definindo as funções essenciais para consulta de status
 *  @dev Esta interface é utilizada para garantir que outros contratos possam interagir com o IdentityRegistryGov de forma consistente, permitindo a verificação do status dos estudantes
 */
interface IIdentityRegistryGov {
    function isStudentValidForSnapshot(address student, uint256 snapshotBlock) external view returns (bool);
    function activeStudentCount() external view returns (uint256);
}

/**
 * @title IGovernanceToken
 * @author Gabriel Henrique Brioto
 * @notice Interface para o contrato de token de governança, definindo as funções essenciais para mint e burn
 * @dev Esta interface é utilizada para garantir que o contrato de governança possa interagir com o token de governança de forma consistente, permitindo a mintagem de tokens como recompensa por voto e a queima de tokens como penalidade por propostas rejeitadas ou spam
 */
interface IGovernanceToken is IERC20 {
    function mint(address to, uint256 amount) external;
    function burn(uint256 amount) external;
}

contract StudentGovernance is AccessControl, ReentrancyGuard {

    IIdentityRegistryGov public identityRegistry;
    IGovernanceToken public token;

    // --- Parâmetros de Governança (Variáveis) ---
    uint256 public quorumPercentage = 30; // Valor em porcentagem (%)
    uint256 public votingPeriod = 2 days; // Duração do período de votação em dias
    uint256 public voteReward = 5 * 10**18; // Recompensa por voto em tokens (considerando 18 decimais)
    uint256 public proposalCost = 50 * 10**18; // Caução para criar uma proposta em tokens (considerando 18 decimais)
    
    // Estrutura para armazenar informações de cada proposta
    struct Proposal {
        address target;         
        bytes callData;         
        string description;
        uint256 snapshotBlock;  
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

    // Eventos
    event ProposalCreated(uint256 indexed id, string description, uint256 deadline);
    event VoteCast(address indexed voter, uint256 indexed proposalId, bool support);
    event ProposalExecuted(uint256 indexed id);
    event GovernanceParamsUpdated(uint256 newQuorum, uint256 newPeriod, uint256 newReward);
    event ProposalRejected(uint256 indexed id, bool quorumMet, bool approved);

    // Construtor que inicializa o contrato com os endereços do IdentityRegistry e do token de governança, e concede o papel de administrador ao endereço fornecido
    constructor(address _identityRegistry, address _token, address _admin) {
        identityRegistry = IIdentityRegistryGov(_identityRegistry);
        token = IGovernanceToken(_token);
        _grantRole(DEFAULT_ADMIN_ROLE, _admin);
    }

    /**
     * @dev Cria uma nova proposta.
     * @param _target Endereço do contrato a ser alterado
     * @param _callData Dados da chamada a ser executada
     * @param _description Descrição da proposta
     */
    function propose(address _target, bytes memory _callData, string memory _description) external {

        // Verifica se o aluno é válido para criar uma proposta, garantindo que ele esteja registrado e ativo no momento da criação da proposta (usando o bloco atual para o snapshot)
        require(identityRegistry.isStudentValidForSnapshot(msg.sender, block.number), "Apenas alunos ativos");

        // Cobrança do caução: o aluno precisa ter dado 'approve' no token para o contrato de governança antes
        require(token.transferFrom(msg.sender, address(this), proposalCost), "Falha no deposito da caucao");

        // Criação da proposta
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

    /**
     * @dev Registra um voto em uma proposta.
     * @param proposalId ID da proposta a ser votada
     * @param support Indica se o voto é a favor ou contra
     */
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

    /**
     * @dev Executa uma proposta aprovada.
     * @param proposalId ID da proposta a ser executada
     */
    function execute(uint256 proposalId) external nonReentrant {
        Proposal storage p = proposals[proposalId];
        
        // Verificações básicas
        require(block.timestamp >= p.votingDeadline, "Votacao em andamento");
        require(!p.executed, "Ja processada");

        // Marca como processada para evitar reentrância/duplicação
        p.executed = true; 

        // --- LÓGICA DE QUÓRUM E APROVAÇÃO ---
        uint256 totalStudents = identityRegistry.activeStudentCount();
        uint256 requiredVotes = (totalStudents * quorumPercentage) / 100;
        if (requiredVotes < 3) requiredVotes = 3;

        bool quorumMet = p.totalVotes >= requiredVotes;
        bool approved = p.yesVotes > p.noVotes;

        // --- CENÁRIO DE SUCESSO ---
        if (quorumMet && approved) {
            
            // Devolve o caução
            if (!p.depositReturned) {
                p.depositReturned = true;
                token.transfer(p.proposer, proposalCost); // Devolve
            }

            // Executa a ação
            (bool success, ) = p.target.call(p.callData);
            require(success, "Falha na chamada externa"); // Este revert é ok, pois é falha técnica
            
            emit ProposalExecuted(proposalId);

        // --- CENÁRIO DE FRACASSO (Spam, rejeição ou falta de quórum) ---
        } else {
            
            // O caução é queimado
            p.depositReturned = true; // Marca como "lidada"
            
            // Tenta queimar (usando try/catch caso o token não suporte burn ou falhe)
            try token.burn(proposalCost) {
                // Queimou com sucesso
            } catch {
                // Se falhar o burn, envia para endereço morto para garantir punição
                token.transfer(address(0xdead), proposalCost);
            }
            
            // Não emite ProposalExecuted, pois falhou.
            // Emitir um evento de falha
            emit ProposalRejected(proposalId, quorumMet, approved);
        }
    }

    /**
     * @dev Altera os parâmetros de governança.
     * @param _quorum Novo percentual de quórum
     * @param _period Novo período de votação
     * @param _reward Nova recompensa por voto
     */
    function setGovernanceParams(uint256 _quorum, uint256 _period, uint256 _reward) external {
        require(msg.sender == address(this), "Apenas via Governanca");
        quorumPercentage = _quorum;
        votingPeriod = _period;
        voteReward = _reward;
        emit GovernanceParamsUpdated(_quorum, _period, _reward);
    }
}