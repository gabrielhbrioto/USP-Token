// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

// @title USPCertificate
// @author Gabriel Henrique Brioto
// @notice Contrato de certificado NFT para alunos, integrando com o IdentityRegistry para valida
// @dev Este contrato implementa um sistema de certificados NFT para alunos, onde cada certificado é soulbound (não transferível) e possui um custo em SocialCurrency para resgate. O contrato integra com o IdentityRegistry para garantir que apenas alunos ativos possam resgatar certificados, e utiliza AccessControl para permitir que a governança possa alterar o custo dos certificados. Além disso, o contrato implementa a interface IERC5192 para indicar que os tokens são soulbound.

// Importações necessárias
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "./IIdentityRegistry.sol";

// Interfaces para integração com outros contratos
interface IERC5192 { function locked(uint256 tokenId) external view returns (bool); }
interface ISocialCurrency { function burnFrom(address account, uint256 amount) external; }

contract USPCertificate is ERC721, ERC721URIStorage, AccessControl, IERC5192 {

    // Definição de papel para governança
    bytes32 public constant GOVERNANCE_ROLE = keccak256("GOVERNANCE_ROLE"); 
    
    // Referências para os contratos de identidade e moeda social
    ISocialCurrency public socialCurrency;
    IIdentityRegistry public identityRegistry;
    
    // Variável para controle do próximo ID de token e custo do certificado
    uint256 private _nextTokenId;
    uint256 public certificateCost = 100 * 10**18; 

    // Eventos
    event CostUpdated(uint256 newCost);
    event Locked(uint256 tokenId); // EIP-5192

    // Construtor que inicializa as referências aos contratos e concede os papéis de governança
    constructor(address _identityRegistry, address _socialCurrency, address admin) 
        ERC721("USP Certificate", "USPC") 
    {
        identityRegistry = IIdentityRegistry(_identityRegistry);
        socialCurrency = ISocialCurrency(_socialCurrency);
        
        _grantRole(DEFAULT_ADMIN_ROLE, admin);
        _grantRole(GOVERNANCE_ROLE, admin);
    }

    /**
     * @dev Altera o custo do certificado.
     * @param newCost Novo custo do certificado em SocialCurrency
     */
    function setCertificateCost(uint256 newCost) external onlyRole(GOVERNANCE_ROLE) {
        certificateCost = newCost;
        emit CostUpdated(newCost);
    }

    /**
     * @dev Resgata um certificado NFT.
     * @param metadataURI URI dos metadados do certificado
     */
    function redeemCertificate(string memory metadataURI) public {

        // Verifica se o remetente é um aluno ativo no IdentityRegistry
        address student = msg.sender;
        require(identityRegistry.isStudentActive(student), "Apenas alunos ativos podem resgatar");
        
        // Usa a variável certificateCost
        socialCurrency.burnFrom(student, certificateCost);
        
        // Mint do certificado NFT para o aluno
        uint256 tokenId = _nextTokenId++;
        _safeMint(student, tokenId);
        _setTokenURI(tokenId, metadataURI);
        emit Locked(tokenId);
    }

    /**
     * @dev Verifica se um certificado está bloqueado (soulbound).
     * @param tokenId ID do token a ser verificado
     */
    function locked(uint256 tokenId) external view override returns (bool) {
        require(_ownerOf(tokenId) != address(0), "Token nao existe");
        return true; 
    }

    /**
     * @dev Atualiza o proprietário de um token (internamente utilizado).
     * @param to Novo proprietário do token
     * @param tokenId ID do token a ser atualizado
     * @param auth Endereço do contrato que está autorizando a atualização
     */
    function _update(address to, uint256 tokenId, address auth) internal override returns (address) {
        address from = _ownerOf(tokenId);
        if (from != address(0) && to != address(0)) {
            revert("Tokens Soulbound nao podem ser transferidos");
        }
        return super._update(to, tokenId, auth);
    }

    /**
     * @dev Verifica se um contrato suporta uma determinada interface.
     * @param interfaceId ID da interface a ser verificada
     */
    function supportsInterface(bytes4 interfaceId) public view override(ERC721, ERC721URIStorage, AccessControl) returns (bool) {
        return interfaceId == type(IERC5192).interfaceId || super.supportsInterface(interfaceId);
    }

    /**
     * @dev Retorna a URI do token.
     * @param tokenId ID do token a ser consultado
     */
    function tokenURI(uint256 tokenId) public view override(ERC721, ERC721URIStorage) returns (string memory) {
        return super.tokenURI(tokenId);
    }
}