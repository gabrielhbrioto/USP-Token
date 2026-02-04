// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "./IIdentityRegistry.sol";

// Interface Oficial do Padrão EIP-5192
interface IERC5192 {
    /// @notice Emitted when the locking status is changed to locked.
    /// @dev If a token is minted and the status is locked, this event should be emitted.
    event Locked(uint256 tokenId);

    /// @notice Emitted when the locking status is changed to unlocked.
    /// @dev If a token is minted and the status is unlocked, this event should be emitted.
    event Unlocked(uint256 tokenId);

    /// @notice Returns the locking status of an Soulbound Token
    /// @dev SBTs assigned to zero address are considered invalid, and queries
    /// about them do throw.
    /// @param tokenId The identifier for an SBT.
    function locked(uint256 tokenId) external view returns (bool);
}

// Interface para interagir com a Moeda Social
interface ISocialCurrency {
    function burnFrom(address account, uint256 amount) external;
}

contract USPCertificate is ERC721, ERC721URIStorage, AccessControl, IERC5192 {
    bytes32 public constant ISSUER_ROLE = keccak256("ISSUER_ROLE");
    
    ISocialCurrency public socialCurrency;
    IIdentityRegistry public identityRegistry;
    
    uint256 private _nextTokenId;
    uint256 public constant CERTIFICATE_COST = 100 * 10**18; // Custo fixo: 100 moedas

    constructor(
        address _identityRegistry, 
        address _socialCurrency, 
        address admin
    ) ERC721("USP Certificate", "USPC") {
        identityRegistry = IIdentityRegistry(_identityRegistry);
        socialCurrency = ISocialCurrency(_socialCurrency);
        
        _grantRole(DEFAULT_ADMIN_ROLE, admin);
        _grantRole(ISSUER_ROLE, admin);
    }

    /**
     * @dev Função para o aluno resgatar um certificado.
     * Ela queima as moedas sociais e emite o NFT já bloqueado.
     */
    function redeemCertificate(string memory metadataURI) public {
        address student = msg.sender;
        
        require(identityRegistry.isStudentActive(student), "Apenas alunos ativos podem resgatar");
        
        // 1. Queima os tokens de moeda social (Exige 'allowance' prévio)
        socialCurrency.burnFrom(student, CERTIFICATE_COST);
        
        // 2. Minta o certificado
        uint256 tokenId = _nextTokenId++;
        _safeMint(student, tokenId);
        _setTokenURI(tokenId, metadataURI);
        
        // 3. Emite evento obrigatório do EIP-5192
        emit Locked(tokenId);
    }

    /**
     * @notice Função obrigatória do EIP-5192.
     * @dev Retorna se o token está bloqueado (Soulbound).
     */
    function locked(uint256 tokenId) external view override returns (bool) {
        // Verifica se o token existe (proprietário não é endereço zero) [cite: 53]
        require(_ownerOf(tokenId) != address(0), "Token nao existe");
        return true; 
    }

    /**
     * @dev Bloqueia qualquer tentativa de transferência.
     */
    function _update(address to, uint256 tokenId, address auth) internal override returns (address) {
        address from = _ownerOf(tokenId);
        
        // Lógica Soulbound: Permite apenas MINT (from == 0) e BURN (to == 0) [cite: 49]
        if (from != address(0) && to != address(0)) {
            revert("Tokens Soulbound nao podem ser transferidos");
        }
        
        return super._update(to, tokenId, auth);
    }

    /**
     * @dev Atualizado para anunciar suporte ao EIP-5192 (0xb45a3c0e).
     */
    function supportsInterface(bytes4 interfaceId) public view override(ERC721, ERC721URIStorage, AccessControl) returns (bool) {
        // 0xb45a3c0e é o ID da interface EIP-5192
        return interfaceId == type(IERC5192).interfaceId || super.supportsInterface(interfaceId);
    }

    function tokenURI(uint256 tokenId) public view override(ERC721, ERC721URIStorage) returns (string memory) {
        return super.tokenURI(tokenId);
    }
}