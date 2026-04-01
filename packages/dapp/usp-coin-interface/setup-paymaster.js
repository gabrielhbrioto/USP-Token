import { createWalletClient, http, parseEther } from 'viem';
import { mnemonicToAccount } from 'viem/accounts';
import { sepolia } from 'viem/chains';

// 1. CONFIGURAÇÕES
const RPC_URL = 'https://sepolia.infura.io/v3/bd7646012a6547b4b8f5aed540422dd1';
const MNEMONIC = 'monkey fat require arctic grit blouse steak planet hurry betray mechanic tool';
const PAYMASTER_ADDRESS = '0xc1E42816B1CDF6977692C6eb4dF18f440a57D77F';
const USP_TOKEN_ADDRESS = '0x56412963Afd3db0C27929A71Ef9cbbfbE3D2cf63';

const PaymasterABI = [
  { "inputs": [], "name": "makeDeposit", "outputs": [], "stateMutability": "payable", "type": "function" },
  { "inputs": [{ "name": "target", "type": "address" }, { "name": "status", "type": "bool" }], "name": "setTargetWhitelist", "outputs": [], "stateMutability": "nonpayable", "type": "function" }
];

async function main() {
  const account = mnemonicToAccount(MNEMONIC);
  const client = createWalletClient({ account, chain: sepolia, transport: http(RPC_URL) });

  console.log(`Configurando Paymaster como: ${account.address}`);

  try {
    // PASSO 1: Depositar ETH no EntryPoint via Paymaster
    console.log("Enviando 0.2 ETH para o depósito do EntryPoint...");
    const hashDep = await client.writeContract({
      address: PAYMASTER_ADDRESS,
      abi: PaymasterABI,
      functionName: 'makeDeposit',
      value: parseEther('0.2'), // Quantidade de Sepolia ETH para cobrir várias transações
    });
    console.log(`✔ Depósito realizado! Hash: ${hashDep}`);

    // PASSO 2: Adicionar o USPToken na Whitelist
    console.log("Adicionando USPToken na Whitelist do Paymaster...");
    const hashWhite = await client.writeContract({
      address: PAYMASTER_ADDRESS,
      abi: PaymasterABI,
      functionName: 'setTargetWhitelist',
      args: [USP_TOKEN_ADDRESS, true],
    });
    console.log(`✔ Whitelist configurada! Hash: ${hashWhite}`);

    console.log("✅ Paymaster pronto para uso!");
  } catch (error) {
    console.error("❌ Erro:", error.message);
  }
}

main();