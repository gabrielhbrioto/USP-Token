// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/USPPaymaster.sol";
import "../src/IdentityRegistry.sol";
import "@account-abstraction/contracts/interfaces/PackedUserOperation.sol";

// 1. Criamos um Mock simples do EntryPoint
contract MockEntryPoint {
    mapping(address => uint256) public balances;

    function balanceOf(address account) public view returns (uint256) {
        return balances[account];
    }

    function depositTo(address to) public payable {
        balances[to] += msg.value;
    }
    
    // Funções extras que podem ser chamadas mas não afetam este teste
    function withdrawTo(address payable, uint256) public {}
}

contract USPPaymasterTest is Test {
    USPPaymaster public paymaster;
    IdentityRegistry public registry;
    MockEntryPoint public mockEntryPoint; // Variável para o Mock
    
    address public admin = address(1);
    address public student = address(2);
    address public targetContract = address(3);

    function setUp() public {
        vm.startPrank(admin);

        // 2. Implantamos o Mock e o Registry
        mockEntryPoint = new MockEntryPoint();
        registry = new IdentityRegistry(admin);

        // Passamos o endereço do Mock para o Paymaster
        paymaster = new USPPaymaster(IEntryPoint(address(mockEntryPoint)), address(registry), admin);

        // Configurações iniciais
        registry.addStudent(student, "123");
        paymaster.setTargetWhitelist(targetContract, true);
        
        // 3. O Pulo do Gato: Não basta dar deal no Paymaster, 
        // tem que depositar no EntryPoint para passar na validação de saldo!
        
        // Dá ETH para o Paymaster
        vm.deal(admin, 10 ether);
        
        // Força o Paymaster a depositar no EntryPoint (Mock)
        // O Paymaster tem a função makeDeposit() que chama entryPoint.depositTo
        paymaster.makeDeposit{value: 5 ether}();

        vm.stopPrank();
    }

    function test_ValidateUserOp_Success() public {
        PackedUserOperation memory op = _createOp(student, targetContract);
        
        // Simula chamada vindo do EntryPoint (Mock)
        vm.prank(address(mockEntryPoint));
        
        (bytes memory context, uint256 validationData) = paymaster.validatePaymasterUserOp(op, bytes32(0), 1000);
        
        assertEq(validationData, 0); // Sucesso
        
        address decodedStudent = abi.decode(context, (address));
        assertEq(decodedStudent, student);
    }

    function test_ValidateUserOp_Fail_NotWhitelisted() public {
        PackedUserOperation memory op = _createOp(student, address(999)); 
        
        vm.prank(address(mockEntryPoint));
        (, uint256 validationData) = paymaster.validatePaymasterUserOp(op, bytes32(0), 1000);
        
        assertEq(validationData, 1); // Falha
    }

    function test_PostOp_ChargesGas() public {
        bytes memory context = abi.encode(student);
        
        vm.prank(address(mockEntryPoint));
        paymaster.postOp(IPaymaster.PostOpMode.opSucceeded, context, 0.01 ether, 0);
        
        assertEq(paymaster.gasSpentByStudent(student), 0.01 ether);
    }

    function _createOp(address sender, address target) internal pure returns (PackedUserOperation memory) {
        PackedUserOperation memory op;
        op.sender = sender;
        op.callData = abi.encodePacked(target, hex"1234");
        return op;
    }
}