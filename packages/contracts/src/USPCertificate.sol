// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "./IIdentityRegistry.sol";

// Interface para interagir com a Moeda Social
interface ISocialCurrency {
    function burnFrom(address account, uint256 amount) external;
}

contract USPCertificate is ERC721, ERC721URIStorage, AccessControl {
    bytes32 public constant ISSUER_ROLE = keccak256("ISSUER_ROLE");
    
    ISocialCurrency public socialCurrency;
    IIdentityRegistry public identityRegistry;
    
    uint256 private _nextTokenId;
    uint256 public constant CERTIFICATE_COST = 100 * 10**18; // Exemplo: 100 moedas

    // Evento para conformidade com ERC-5192 (Soulbound)
    event Locked(uint256 tokenId);

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
     * Ela queima as moedas sociais e emite o NFT.
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
        
        // 3. Emite evento de que o token está "preso" (Soulbound)
        emit Locked(tokenId);
    }

    /**
     * @dev Bloqueia qualquer tentativa de transferência (Soulbound Logic).
     */
    function _update(address to, uint256 tokenId, address auth) internal override returns (address) {
        address from = _ownerOf(tokenId);
        
        // Permite apenas o MINT (from == address(0)) e o BURN (to == address(0))
        if (from != address(0) && to != address(0)) {
            revert("Tokens Soulbound nao podem ser transferidos");
        }
        
        return super._update(to, tokenId, auth);
    }

    // Funções de suporte necessárias para o ERC721URIStorage e AccessControl
    function supportsInterface(bytes4 interfaceId) public view override(ERC721, ERC721URIStorage, AccessControl) returns (bool) {
        return super.supportsInterface(interfaceId);
    }

    function tokenURI(uint256 tokenId) public view override(ERC721, ERC721URIStorage) returns (string memory) {
        return super.tokenURI(tokenId);
    }
}