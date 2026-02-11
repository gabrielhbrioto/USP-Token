// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

// @title DeployUSP
// @author Gabriel Henrique Brioto
// @notice Este script é responsável por realizar o deploy de todos os contratos do projeto USP
// @dev Ele deve ser executado usando Foundry, e espera que a variável de ambiente PRIVATE_KEY esteja configurada com a chave privada do deployer.

// Importação dos contratos a serem implantados
import "forge-std/Script.sol";
import "../src/USPToken.sol";
import "../src/USPPaymaster.sol";
import "../src/USPCertificate.sol";
import "../src/StudentGovernance.sol";
import "../src/IdentityRegistry.sol";

contract DeployUSP is Script {
    // Endereço oficial do EntryPoint na Sepolia (v0.6)
    // Verifique se seus contratos usam v0.6 ou v0.7. O endereço abaixo é o padrão para v0.6.
    // Se for v0.7, use 0x0000000071727De22E5E9d8BAf0edAc6f37da032
    address constant ENTRYPOINT_ADDR = 0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789;

    /**
     * @dev Função principal que executa o script de deploy.
     */
    function run() external {
        // 1. Configuração Inicial
        
        // CORREÇÃO: Lê a variável como String (Frase) em vez de Uint (Número)
        string memory mnemonic = vm.envString("PRIVATE_KEY");
        
        // Deriva a chave privada a partir da frase mnemônica (índice 0)
        uint256 deployerPrivateKey = vm.deriveKey(mnemonic, 0);
        
        address deployerAddr = vm.addr(deployerPrivateKey);
        
        vm.startBroadcast(deployerPrivateKey);

        console.log("Iniciando Deploy com a conta:", deployerAddr);

        // ------------------------------
        // ETAPA 1: Identity Registry
        // ------------------------------
        IdentityRegistry registry = new IdentityRegistry(deployerAddr);
        console.log("IdentityRegistry deployado em:", address(registry));

        // ------------------------------
        // ETAPA 2: USPToken
        // ------------------------------
        USPToken token = new USPToken(address(registry), deployerAddr);
        console.log("USPToken deployado em:", address(token));

        // ------------------------------
        // ETAPA 3: USPCertificate
        // ------------------------------
        USPCertificate cert = new USPCertificate(
            address(registry), 
            address(token), 
            deployerAddr
        );
        console.log("USPCertificate deployado em:", address(cert));

        // ------------------------------
        // ETAPA 4: StudentGovernance
        // ------------------------------
        StudentGovernance governance = new StudentGovernance(
            address(registry),
            address(token),
            deployerAddr
        );
        console.log("StudentGovernance deployado em:", address(governance));

        // ------------------------------
        // ETAPA 5: USPPaymaster
        // ------------------------------
        USPPaymaster paymaster = new USPPaymaster(
            IEntryPoint(ENTRYPOINT_ADDR),
            address(registry),
            deployerAddr
        );
        console.log("USPPaymaster deployado em:", address(paymaster));

        // ------------------------------
        // ETAPA 6: Configuração de Permissões (Orquestração)
        // ------------------------------
        console.log("Configurando permissoes...");

        // A) Dar permissão para a Governança emitir tokens (Vote-to-Earn)
        token.grantRole(token.MINTER_ROLE(), address(governance));
        console.log("- Governance agora tem MINTER_ROLE no Token");

        // B) Dar permissão para a Governança alterar o custo do certificado
        cert.grantRole(cert.GOVERNANCE_ROLE(), address(governance));
        console.log("- Governance agora tem GOVERNANCE_ROLE no Certificado");

        // C) Transferir propriedade do Paymaster para a Governança
        // (Isso permite que a DAO altere limites de gás e whitelist)
        paymaster.transferOwnership(address(governance));
        console.log("- Governance agora e Owner do Paymaster");

        // ------------------------------
        // Opcional: Depositar ETH no Paymaster para ele começar funcionando
        // ------------------------------
        // paymaster.makeDeposit{value: 0.05 ether}();
        // console.log("- Deposito inicial feito no Paymaster");

        vm.stopBroadcast();
    }
}