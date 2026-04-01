import { createWalletClient, http } from 'viem';
import { mnemonicToAccount } from 'viem/accounts';
import { sepolia } from 'viem/chains';

const RPC_URL = 'https://sepolia.infura.io/v3/bd7646012a6547b4b8f5aed540422dd1';
const MNEMONIC = 'monkey fat require arctic grit blouse steak planet hurry betray mechanic tool';
const USP_TOKEN_ADDRESS = '0x56412963Afd3db0C27929A71Ef9cbbfbE3D2cf63';
const targetStudentAddress = '0x6A19f2C4ab35827F38b8eE796Bb18dC0DDba6240';

async function main() {
  const account = mnemonicToAccount(MNEMONIC);
  const client = createWalletClient({ account, chain: sepolia, transport: http(RPC_URL) });

  console.log(`Tentando mint direto com admin: ${account.address}`);

  try {
    const hash = await client.writeContract({
      address: USP_TOKEN_ADDRESS,
      abi: [{ "inputs": [{ "name": "to", "type": "address" }, { "name": "amount", "type": "uint256" }], "name": "mint", "outputs": [], "stateMutability": "nonpayable", "type": "function" }],
      functionName: 'mint',
      args: [targetStudentAddress, 10000n * 10n**18n] // 1000 tokens
    });
    console.log(`✅ Sucesso! Transação enviada: ${hash}`);
  } catch (error) {
    console.error(`❌ Falha no mint direto: ${error.shortMessage || error.message}`);
  }
}

main();