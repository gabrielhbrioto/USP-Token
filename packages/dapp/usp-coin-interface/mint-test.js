import { 
  encodeFunctionData, 
  parseUnits, 
  createPublicClient, 
  http, 
  encodeAbiParameters, 
  keccak256, 
  concat 
} from 'viem';
import { mnemonicToAccount } from 'viem/accounts';
import { sepolia } from 'viem/chains';
import { decodeErrorResult } from 'viem';

// 1. CONFIGURAÇÕES
const RPC_URL = 'https://sepolia.infura.io/v3/bd7646012a6547b4b8f5aed540422dd1';
const MNEMONIC = 'monkey fat require arctic grit blouse steak planet hurry betray mechanic tool';
const RELAYER_URL = 'http://localhost:8080/relay';

// Endereços do seu deploy.log
const IDENTITY_REGISTRY_ADDRESS = '0x2986C5c7FceE1ecA6c6146031ef65e797C965CD3';
const USP_TOKEN_ADDRESS = '0x56412963Afd3db0C27929A71Ef9cbbfbE3D2cf63';
const ENTRY_POINT_ADDRESS = '0x0000000071727De22E5E9d8BAf0edAc6f37da032';
const USPPAYMASTER_ADDRESS = '0xc1E42816B1CDF6977692C6eb4dF18f440a57D77F';

const targetStudentAddress = '0x6A19f2C4ab35827F38b8eE796Bb18dC0DDba6240';
const senderSmartAccount = '0x9e26481b51e382B8563a888B525b060B1e9C07A2'; 

// Utilitário para converter Hex para Base64 (Requisito do seu Relayer Go)
const hexToBase64 = (hex) => {
  const cleanHex = hex.startsWith('0x') ? hex.slice(2) : hex;
  return Buffer.from(cleanHex, 'hex').toString('base64');
};

async function main() {
  const account = mnemonicToAccount(MNEMONIC);
  const publicClient = createPublicClient({ chain: sepolia, transport: http(RPC_URL) });

  console.log(`Assinando com a conta: ${account.address}`);

  // 2. BUSCA DE DADOS ON-CHAIN (Nonce)
  const nonce = await publicClient.readContract({
    address: ENTRY_POINT_ADDRESS,
    abi: [{ "inputs": [{ "name": "sender", "type": "address" }, { "name": "key", "type": "uint192" }], "name": "getNonce", "outputs": [{ "name": "nonce", "type": "uint256" }], "stateMutability": "view", "type": "function" }],
    functionName: 'getNonce',
    args: [senderSmartAccount, 0n],
  });

  const mintCallData = encodeFunctionData({
    abi: [{ "inputs": [{ "name": "to", "type": "address" }, { "name": "amount", "type": "uint256" }], "name": "mint", "outputs": [], "stateMutability": "nonpayable", "type": "function" }],
    functionName: 'mint',
    args: [targetStudentAddress, parseUnits('1000', 18)]
  });

  // 3. ESTRUTURA DA USER OPERATION (Valores numéricos para o JSON)
  const userOp = {
    sender: senderSmartAccount,
    nonce: Number(nonce),
    initCode: "0x",
    callData: mintCallData,
    callGasLimit: parseInt("0x30D40", 16),
    verificationGasLimit: parseInt("0x493E0", 16), // Aumentado para 300k p/ validação
    preVerificationGas: parseInt("0xC350", 16),
    maxFeePerGas: parseInt("0x77359400", 16),
    maxPriorityFeePerGas: parseInt("0x3B9ACA00", 16),
    paymasterAndData: USPPAYMASTER_ADDRESS,
    signature: "0x" 
  };

  // 4. GERAÇÃO DA ASSINATURA (PADRÃO ERC-4337)
  // O hash da UserOp é o que o EntryPoint valida
  const userOpHash = keccak256(encodeAbiParameters(
    [
      { type: 'address' }, { type: 'uint256' }, { type: 'bytes32' }, { type: 'bytes32' },
      { type: 'uint256' }, { type: 'uint256' }, { type: 'uint256' }, { type: 'uint256' },
      { type: 'uint256' }, { type: 'bytes32' }
    ],
    [
      userOp.sender, BigInt(userOp.nonce), keccak256(userOp.initCode), keccak256(userOp.callData),
      BigInt(userOp.callGasLimit), BigInt(userOp.verificationGasLimit), BigInt(userOp.preVerificationGas),
      BigInt(userOp.maxFeePerGas), BigInt(userOp.maxPriorityFeePerGas), keccak256(userOp.paymasterAndData)
    ]
  ));

  // O hash final inclui o EntryPoint e ChainID para evitar Replay em outras redes
  const finalHash = keccak256(encodeAbiParameters(
    [{ type: 'bytes32' }, { type: 'address' }, { type: 'uint256' }],
    [userOpHash, ENTRY_POINT_ADDRESS, BigInt(sepolia.id)]
  ));

  // Assinando o hash final com a sua frase mnemônica
  userOp.signature = await account.signMessage({ message: { raw: finalHash } });

  console.log("UserOp assinada com sucesso!");

  // 5. ENVIO FORMATADO EM BASE64 PARA O RELAYER GO
  const payload = {
    ...userOp,
    initCode: hexToBase64(userOp.initCode),
    callData: hexToBase64(userOp.callData),
    paymasterAndData: hexToBase64(userOp.paymasterAndData),
    signature: hexToBase64(userOp.signature)
  };

  // Substitua o seu bloco de simulação no mint-test.js por este:

  try {
    console.log("Iniciando simulação local...");
    await publicClient.simulateContract({
      address: ENTRY_POINT_ADDRESS,
      abi: [
        {
          "inputs": [
            { 
              "components": [
                { "name": "sender", "type": "address" },
                { "name": "nonce", "type": "uint256" },
                { "name": "initCode", "type": "bytes" },
                { "name": "callData", "type": "bytes" },
                { "name": "callGasLimit", "type": "uint256" },
                { "name": "verificationGasLimit", "type": "uint256" },
                { "name": "preVerificationGas", "type": "uint256" },
                { "name": "maxFeePerGas", "type": "uint256" },
                { "name": "maxPriorityFeePerGas", "type": "uint256" },
                { "name": "paymasterAndData", "type": "bytes" },
                { "name": "signature", "type": "bytes" }
              ], 
              "name": "ops", 
              "type": "tuple[]" 
            },
            { "name": "beneficiary", "type": "address" }
          ],
          "name": "handleOps",
          "outputs": [],
          "stateMutability": "nonpayable",
          "type": "function"
        },
        // 🚨 DEFINIÇÃO DO ERRO ADICIONADA AQUI 🚨
        {
          "inputs": [
            { "name": "opIndex", "type": "uint256" },
            { "name": "reason", "type": "string" }
          ],
          "name": "FailedOp",
          "type": "error"
        }
      ],
      functionName: 'handleOps',
      args: [[userOp], account.address],
    });
    console.log("✅ Simulação passou com sucesso!");
  } catch (error) {
    // O Viem aninha os erros de contrato em 'cause' ou 'metaMessages'
    console.error("❌ Falha na simulação!");
    // Vasculha os erros aninhados buscando dados brutos (hex)
    const revertData = error.walk?.()?.data || error?.cause?.data;

    if (revertData) {
      try {
        const decodedError = decodeErrorResult({
          abi: [{
            "inputs": [
              { "name": "opIndex", "type": "uint256" },
              { "name": "reason", "type": "string" }
            ],
            "name": "FailedOp",
            "type": "error"
          }],
          data: revertData
        });
        console.error(`⚠️ Erro formatado do EntryPoint: [Index: ${decodedError.args[0]}] Motivo: ${decodedError.args[1]}`);
      } catch (decodeErr) {
        console.error("⚠️ Dados do Revert (não foi possível decodificar como FailedOp):", revertData);
      }
    } else {
      console.error(error.shortMessage || error.message);
    }
    return; 
  }

  try {
    const response = await fetch(RELAYER_URL, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload),
    });
    const data = await response.json();
    console.log(response.ok ? '✅ Transação aceita pelo Relayer!' : '❌ Erro:', data);
  } catch (e) {
    console.error('❌ Erro na conexão:', e.message);
  }
}

main();