// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "./IIdentityRegistry.sol";

// ... (Interfaces IERC5192 e ISocialCurrency mantidas iguais) ...
interface IERC5192 { function locked(uint256 tokenId) external view returns (bool); }
interface ISocialCurrency { function burnFrom(address account, uint256 amount) external; }

contract USPCertificate is ERC721, ERC721URIStorage, AccessControl, IERC5192 {
    bytes32 public constant GOVERNANCE_ROLE = keccak256("GOVERNANCE_ROLE"); // Role para alterar preço
    
    ISocialCurrency public socialCurrency;
    IIdentityRegistry public identityRegistry;
    
    uint256 private _nextTokenId;
    
    // Agora é uma variável, não constante
    uint256 public certificateCost = 100 * 10**18; 

    event CostUpdated(uint256 newCost);
    event Locked(uint256 tokenId); // EIP-5192

    constructor(address _identityRegistry, address _socialCurrency, address admin) 
        ERC721("USP Certificate", "USPC") 
    {
        identityRegistry = IIdentityRegistry(_identityRegistry);
        socialCurrency = ISocialCurrency(_socialCurrency);
        
        _grantRole(DEFAULT_ADMIN_ROLE, admin);
        _grantRole(GOVERNANCE_ROLE, admin); // Admin inicial, depois transferir para o contrato Governance
    }

    // Função Setter controlada pela Governança
    function setCertificateCost(uint256 newCost) external onlyRole(GOVERNANCE_ROLE) {
        certificateCost = newCost;
        emit CostUpdated(newCost);
    }

    function redeemCertificate(string memory metadataURI) public {
        address student = msg.sender;
        require(identityRegistry.isStudentActive(student), "Apenas alunos ativos podem resgatar");
        
        // Usa a variável certificateCost
        socialCurrency.burnFrom(student, certificateCost);
        
        uint256 tokenId = _nextTokenId++;
        _safeMint(student, tokenId);
        _setTokenURI(tokenId, metadataURI);
        emit Locked(tokenId);
    }

    // ... (Funções locked, supportsInterface, _update, tokenURI mantidas iguais) ...
    function locked(uint256 tokenId) external view override returns (bool) {
        require(_ownerOf(tokenId) != address(0), "Token nao existe");
        return true; 
    }

    function _update(address to, uint256 tokenId, address auth) internal override returns (address) {
        address from = _ownerOf(tokenId);
        if (from != address(0) && to != address(0)) {
            revert("Tokens Soulbound nao podem ser transferidos");
        }
        return super._update(to, tokenId, auth);
    }

    function supportsInterface(bytes4 interfaceId) public view override(ERC721, ERC721URIStorage, AccessControl) returns (bool) {
        return interfaceId == type(IERC5192).interfaceId || super.supportsInterface(interfaceId);
    }

    function tokenURI(uint256 tokenId) public view override(ERC721, ERC721URIStorage) returns (string memory) {
        return super.tokenURI(tokenId);
    }
}