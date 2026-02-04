// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

// test/Integration.t.sol
import "forge-std/Test.sol";
import "../src/USPPaymaster.sol";
// ... (outros imports)

contract IntegrationTest is Test {
    // Setup de todos os contratos...
    
    function test_FullStudentJourney() public {
        // 1. Admin registra aluno e dá moedas iniciais [cite: 5, 52]
        // 2. Paymaster valida que o aluno é ativo [cite: 15]
        // 3. Aluno solicita resgate do certificado [cite: 36]
        // 4. Moedas são queimadas e certificado emitido [cite: 37, 38]
        // 5. Paymaster registra o gasto de gás do aluno [cite: 21, 22]
    }
}