// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Test, console} from "forge-std/Test.sol";
import {USPCertificate} from "../src/USPCertificate.sol";
import {USPToken} from "../src/USPToken.sol";
import {IdentityRegistry} from "../src/IdentityRegistry.sol";

contract USPCertificateTest is Test {
    USPCertificate public certificate;
    USPToken public token;
    IdentityRegistry public registry;

    address public admin = address(1);
    address public student = address(2);
    address public otherPerson = address(3);

    // Evento do EIP-5192 para verificação
    event Locked(uint256 tokenId);

    function setUp() public {
        vm.startPrank(admin);

        // 1. Deploy da infraestrutura base
        registry = new IdentityRegistry(admin);
        token = new USPToken(address(registry), admin);

        // 2. Deploy do Certificado
        certificate = new USPCertificate(
            address(registry),
            address(token),
            admin
        );

        // 3. Configuração inicial:
        // - Adiciona aluno ao registro
        registry.addStudent(student, "123456");
        
        // - Dá permissão para o Certificado queimar tokens (MINTER_ROLE não é necessário para burnFrom, 
        // mas o contrato precisa ser confiável se houvesse lógica de mint reverso. 
        // Aqui o foco é o aluno ter saldo).
        
        // - Emite tokens iniciais para o aluno pagar pelo certificado
        token.mint(student, 500 * 10**18); // Aluno recebe 500 tokens

        vm.stopPrank();
    }

    function test_RedeemCertificate_HappyPath() public {
        vm.startPrank(student);

        // Aluno aprova o contrato de certificado a gastar 100 tokens
        token.approve(address(certificate), 100 * 10**18);

        // Verifica a expectativa do Evento Locked (EIP-5192)
        // O primeiro token ID será 0
        vm.expectEmit(true, true, true, true);
        emit Locked(0);

        // Executa o resgate
        certificate.redeemCertificate("ipfs://exemplo-metadata");

        vm.stopPrank();

        // VALIDAÇÕES:
        // 1. Saldo de tokens do aluno diminuiu em 100
        assertEq(token.balanceOf(student), 400 * 10**18);

        // 2. Aluno é dono do NFT ID 0
        assertEq(certificate.ownerOf(0), student);

        // 3. URI está correta
        assertEq(certificate.tokenURI(0), "ipfs://exemplo-metadata");
    }

    function test_EIP5192_Compliance() public {
        vm.startPrank(student);
        token.approve(address(certificate), 100 * 10**18);
        certificate.redeemCertificate("ipfs://metadata");
        vm.stopPrank();

        uint256 tokenId = 0;

        // VALIDAÇÃO 1: Função locked() retorna true
        assertTrue(certificate.locked(tokenId));

        // VALIDAÇÃO 2: Suporte à Interface (ERC-165)
        // O ID da interface EIP-5192 é 0xb45a3c0e
        bytes4 interfaceId5192 = 0xb45a3c0e;
        assertTrue(certificate.supportsInterface(interfaceId5192));
    }

    function test_Fail_Transfer_Soulbound() public {
        // Preparação: Aluno resgata o certificado
        vm.startPrank(student);
        token.approve(address(certificate), 100 * 10**18);
        certificate.redeemCertificate("ipfs://metadata");

        // Tentativa de Transferência para outra pessoa
        vm.expectRevert("Tokens Soulbound nao podem ser transferidos");
        certificate.transferFrom(student, otherPerson, 0);
        
        vm.stopPrank();
    }

    function test_Fail_Redeem_InactiveStudent() public {
        // ETAPA 1: Preparação (Como Admin)
        vm.startPrank(admin); 
        
        // Cadastra o "otherPerson" temporariamente
        registry.addStudent(otherPerson, "999");
        
        // Dá tokens para ele (Só o Admin com MINTER_ROLE pode fazer isso)
        // O erro original acontecia aqui porque o 'prank' estava errado
        token.mint(otherPerson, 200 * 10**18);
        
        // Agora desativa o aluno para testar o erro
        registry.setStudentStatus(otherPerson, false); 
        
        vm.stopPrank(); // Encerra o modo Admin

        // ETAPA 2: Ação de Falha (Como Aluno Inativo)
        vm.startPrank(otherPerson);
        
        // Aprova o gasto (isso funciona mesmo inativo, pois é ERC20 padrão)
        token.approve(address(certificate), 100 * 10**18);

        // AQUI ESTÁ O SEGREDO:
        // Dizemos ao Foundry: "A próxima linha TEM que falhar com essa mensagem exata"
        vm.expectRevert("Apenas alunos ativos podem resgatar");
        
        // Chamamos a função que vai falhar
        certificate.redeemCertificate("ipfs://fail");
        
        vm.stopPrank();
    }

    function test_Fail_Redeem_InsufficientBalance() public {
        vm.startPrank(student);
        
        // Gasta todos os tokens ou transfere para outro (se pudesse)
        // Vamos apenas tentar resgatar SEM ter aprovado o suficiente ou sem saldo
        // Simulando saldo baixo: Queima os tokens do aluno via admin (se tivesse burn access)
        // Ou mais simples: Aluno tenta resgatar 6 certificados com saldo de 500
        
        token.approve(address(certificate), 1000 * 10**18); // Aprova muito
        
        // Resgata 5 vezes (Custo total 500) - OK
        for(uint i=0; i<5; i++){
            certificate.redeemCertificate("ipfs://ok");
        }

        // Tenta o 6º (Saldo insuficiente)
        // O erro virá do ERC20: "ERC20: burn amount exceeds balance" ou similar
        // Dependendo da implementação do OpenZeppelin, pode ser Panic(0x11) ou mensagem customizada
        vm.expectRevert(); 
        certificate.redeemCertificate("ipfs://fail");

        vm.stopPrank();
    }
}