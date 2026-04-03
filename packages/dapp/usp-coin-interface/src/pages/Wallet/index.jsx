import React, { useState, useEffect } from 'react';
import { Alert, Snackbar } from '@mui/material';
import { useReadContracts, usePublicClient } from 'wagmi';
import { formatUnits, parseUnits, encodeFunctionData, keccak256, encodeAbiParameters, parseAbiParameters } from 'viem';

// Importação dos ABIs (assumindo que foram compilados para JSON)
import IdentityRegistryABI from '../../abi/IdentityRegistry.json';
import USPTokenABI from '../../abi/USPToken.json'; 
import StudentGovernanceABI from '../../abi/StudentGovernance.json';
import { useAccount, useSwitchChain } from 'wagmi';
import { sepolia } from 'wagmi/chains';
import { useAuth } from '../../contexts/AuthContext';
import { privateKeyToAccount } from 'viem/accounts';

// Variáveis de ambiente do Vite
const IDENTITY_REGISTRY_ADDR = import.meta.env.VITE_IDENTITY_REGISTRY_ADDRESS;
const USP_TOKEN_ADDR = import.meta.env.VITE_USP_TOKEN_ADDRESS;
const GOVERNANCE_ADDR = import.meta.env.VITE_GOVERNANCE_ADDRESS;

export function Wallet() {
  const { walletAddress: studentAddress, privateKey } = useAuth(); 
  const publicClient = usePublicClient();
  const [receiverAddress, setReceiverAddress] = useState('');
  const [transferAmount, setTransferAmount] = useState('');
  const [transferStatus, setTransferStatus] = useState('');
  const [readyProposals, setReadyProposals] = useState([]);
  const { chainId, isConnected } = useAccount();
  const { switchChain } = useSwitchChain();
  const isWrongNetwork = isConnected && chainId !== sepolia.id;
  const [transferTxHash, setTransferTxHash] = useState('');
  const [copySnackbar, setCopySnackbar] = useState({
    open: false,
    message: '',
    severity: 'success',
  });

  const ENTRY_POINT_ADDRESS = '0x0000000071727De22E5E9d8BAf0edAc6f37da032'; // Padrão ERC-4337 v0.6
  const CHAIN_ID = 11155111; // ID da Sepolia

  const getUserOpHash = (userOp) => {
  // 1. Hashea os campos dinâmicos
  const hashedInitCode = keccak256(userOp.initCode || '0x');
  const hashedCallData = keccak256(userOp.callData);
  const hashedPaymasterAndData = keccak256(userOp.paymasterAndData || '0x');

  // 2. Empacota todos os dados exatamente como a EVM faz (ABI Encode)
  const packedUserOp = encodeAbiParameters(
    parseAbiParameters('address, uint256, bytes32, bytes32, uint256, uint256, uint256, uint256, uint256, bytes32'),
    [
      userOp.sender,
      BigInt(userOp.nonce),
      hashedInitCode,
      hashedCallData,
      BigInt(userOp.callGasLimit),
      BigInt(userOp.verificationGasLimit),
      BigInt(userOp.preVerificationGas),
      BigInt(userOp.maxFeePerGas),
      BigInt(userOp.maxPriorityFeePerGas),
      hashedPaymasterAndData
    ]
  );

  const encodedUserOpHash = keccak256(packedUserOp);

  // 3. Adiciona o EntryPoint e o ChainID para evitar replay attacks
  const finalHash = keccak256(
    encodeAbiParameters(
      parseAbiParameters('bytes32, address, uint256'),
      [encodedUserOpHash, ENTRY_POINT_ADDRESS, BigInt(CHAIN_ID)]
    )
  );

  return finalHash; // Esse é o Hash oficial que a carteira precisa assinar!
};

  // 1. Leitura em Lote (Status do Aluno, Saldo e Dados Globais da Governança)
  const { data: onChainData, isLoading, error: globalError } = useReadContracts({
    contracts: [
      {
        address: IDENTITY_REGISTRY_ADDR,
        abi: IdentityRegistryABI,
        functionName: 'isStudentActive',
        args: [studentAddress],
      },
      {
        address: USP_TOKEN_ADDR,
        abi: USPTokenABI,
        functionName: 'balanceOf',
        args: [studentAddress],
      },
      {
        address: GOVERNANCE_ADDR,
        abi: StudentGovernanceABI,
        functionName: 'quorumPercentage',
      },
      {
        address: IDENTITY_REGISTRY_ADDR,
        abi: IdentityRegistryABI,
        functionName: 'activeStudentCount',
      }
    ]
  });

  useEffect(() => {
    console.log("1. Endereços carregados do .env:", { 
      registry: IDENTITY_REGISTRY_ADDR, 
      token: USP_TOKEN_ADDR 
    });
    console.log("2. Status das chamadas do Wagmi:", onChainData);
    if (globalError) console.error("3. Erro Global do Wagmi:", globalError);
  }, [onChainData, globalError]);

  // O wagmi retorna os dados na mesma ordem do array de contratos
  const isActive = onChainData?.[0]?.result; // Retorna true se o estudante estiver ativo no IdentityRegistry [cite: 73]
  const balance = onChainData?.[1]?.result ? formatUnits(onChainData[1].result, 18) : '0';
  const quorumPercentage = onChainData?.[2]?.result || 30n; // Padrão definido no contrato [cite: 15]
  const activeStudentCount = onChainData?.[3]?.result || 0n; // Utilizado para calcular o quórum dinâmico [cite: 58, 61]

  // Função auxiliar para construir a UserOperation (Padrão ERC-4337)
  const buildUserOp = async (senderAddress, targetContract, encodedData) => {
    // NOTA ARQUITETURAL: Em um cenário real de Account Abstraction (AA),
    // o 'callData' da UserOp NÃO é direto o callData do Token.
    // Ele deve ser a chamada da função 'execute(target, value, data)' da sua Smart Account.
    // Mas vamos manter a estrutura para o seu backend conseguir parsear:

    return {
      sender: senderAddress,
      nonce: 0, // ⚠️ Atenção: Em produção, você deve ler o nonce atual no EntryPoint
      initCode: "0x", // Vazio se a Smart Account do aluno já estiver "deployed"
      callData: encodedData,
      callGasLimit: 100000, // Limites simulados para o Relayer pagar
      verificationGasLimit: 100000,
      preVerificationGas: 50000,
      maxFeePerGas: 0, // 0 pois o Relayer (USPPaymaster) que vai arcar com o gás
      maxPriorityFeePerGas: 0,
      paymasterAndData: "0x", // O seu Relayer em Go vai preencher isso antes de enviar pra rede!
      signature: "0x" // ⚠️ Atenção: O aluno precisará assinar o Hash dessa UserOp.
    };
  };

  const handleCopyAddress = async () => {
    if (!studentAddress) return;

    try {
      await navigator.clipboard.writeText(studentAddress);
      setCopySnackbar({
        open: true,
        message: 'Endereço copiado para a area de transferencia.',
        severity: 'success',
      });
    } catch {
      setCopySnackbar({
        open: true,
        message: 'Nao foi possivel copiar o endereco automaticamente.',
        severity: 'error',
      });
    }
  };

  const handleCloseSnackbar = (_, reason) => {
    if (reason === 'clickaway') return;
    setCopySnackbar((prev) => ({ ...prev, open: false }));
  };

  // 2. Validação e Envio da Transferência via Account Abstraction
 const handleTransfer = async (e) => {
    e.preventDefault();
    setTransferStatus('Validando destinatário...');
    setTransferTxHash(''); // Limpa o hash de transferências anteriores
    console.log('Validando destinatário...');

    try {
      // 1. Regra de Negócio: O destinatário DEVE ser um aluno ativo no IdentityRegistry
      const isReceiverActive = await publicClient.readContract({
        address: IDENTITY_REGISTRY_ADDR,
        abi: IdentityRegistryABI,
        functionName: 'isStudentActive',
        args: [receiverAddress]
      });

      if (!isReceiverActive) {
        setTransferStatus('Erro: O endereço de destino não é um aluno ativo no registro.');
        console.error('Destinatário não é um aluno ativo.');
        return;
      }

      setTransferStatus('Montando UserOperation...');
      console.log('Montando UserOperation...');

      // 2. Gera o callData da transferência
      const callData = encodeFunctionData({ 
        abi: USPTokenABI, 
        functionName: 'transfer', 
        args: [receiverAddress, parseUnits(transferAmount, 18)] 
      });
      
      // 3. Constrói a UserOp base (ainda com signature vazia "0x")
      const userOp = await buildUserOp(studentAddress, USP_TOKEN_ADDR, callData);

      // 4. Calcula o Hash criptográfico da UserOp simulando o EntryPoint
      const userOpHash = getUserOpHash(userOp);
      
      setTransferStatus('Assinando transação invisível...');
      
      // 5. Assinatura Invisível (Local)
      // Instancia a carteira local a partir da chave privada do aluno em memória
      const account = privateKeyToAccount(privateKey);
      
      // Assina o Hash instantaneamente sem nenhum pop-up
      const signature = await account.signMessage({
        message: { raw: userOpHash },
      });

      // 6. Acopla a assinatura verdadeira na nossa operação
      userOp.signature = signature;

      setTransferStatus('Enviando transação para o Relayer...');
      
      // 7. Envia para o servidor em Go
      const response = await fetch(import.meta.env.VITE_RELAYER_URL, { 
        method: 'POST', 
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(userOp) 
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || 'Falha ao processar pelo Relayer');
      }

      const data = await response.json();
      
      // 8. Sucesso! Atualiza a tela e exibe o link do Etherscan
      setTransferStatus('Transferência realizada com sucesso!');
      setTransferTxHash(data.txHash);
      
    } catch (error) {
      // Captura tanto rejeição de assinatura quanto erros do servidor Go
      setTransferStatus(`Erro na transação: ${error.message}`);
      console.error('Erro ao enviar UserOperation:', error);
    }
  };

  // 3. Simulação de checagem de propostas (Idealmente via Indexador/Backend)
  // Checa se as propostas do aluno atendem a todos os critérios do StudentGovernance.sol
  useEffect(() => {
    const checkProposals = async () => {
      // Exemplo: O backend informou que as propostas ID 1 e 2 são deste usuário
      const userProposalIds = [1, 2]; 
      const readyIds = [];

      for (const id of userProposalIds) {
        try {
          const proposal = await publicClient.readContract({
            address: GOVERNANCE_ADDR,
            abi: StudentGovernanceABI,
            functionName: 'proposals',
            args: [BigInt(id)]
          });

          // Desestruturação baseada na struct Proposal do contrato [cite: 20, 21]
          const [
            target, callData, description, snapshotBlock, votingDeadline, 
            yesVotes, noVotes, totalVotes, executed, proposer, depositReturned
          ] = proposal;

          // Regras de Execução do Smart Contract:
          const blockTimestamp = BigInt(Math.floor(Date.now() / 1000));
          const isDeadlinePassed = blockTimestamp >= votingDeadline; // Requer que o período de votação tenha encerrado [cite: 39]
          
          let requiredVotes = (activeStudentCount * quorumPercentage) / 100n; // Cálculo dinâmico do quórum [cite: 41, 42]
          if (requiredVotes < 3n) requiredVotes = 3n; // Quórum mínimo de 3 votos [cite: 42]
          
          const quorumMet = totalVotes >= requiredVotes; // Verifica se o quórum foi atingido [cite: 43]
          const isApproved = yesVotes > noVotes; // Verifica se há mais votos a favor do que contra [cite: 43]

          if (isDeadlinePassed && !executed && quorumMet && isApproved && proposer === studentAddress) {
            readyIds.push(id);
          }
        } catch (error) {
          console.error(`Erro ao checar proposta ${id}:`, error);
        }
      }
      setReadyProposals(readyIds);
    };

    if (activeStudentCount > 0n) checkProposals();
  }, [publicClient, activeStudentCount, quorumPercentage, studentAddress]);

  if (isWrongNetwork) {
    return (
      <div className="max-w-4xl mx-auto p-6 bg-red-50 border border-red-200 text-center rounded-lg mt-8">
        <h2 className="text-xl font-bold text-red-800 mb-2">Rede Incorreta detectada</h2>
        <p className="text-red-600 mb-4">
          O dApp do USPToken funciona exclusivamente na rede Sepolia. 
          Por favor, altere a rede na sua carteira.
        </p>
        <button 
          onClick={() => switchChain({ chainId: sepolia.id })}
          className="bg-red-600 hover:bg-red-700 text-white font-bold py-2 px-4 rounded"
        >
          Trocar para Sepolia
        </button>
      </div>
    );
  }

  if (isLoading) return <div className="p-8 text-center text-gray-600">Carregando carteira USP...</div>;

  return (
    <>
      <div className="max-w-4xl mx-auto p-6 bg-white/95 shadow-xl rounded-2xl mt-8 border border-slate-100">
      
      {/* Alertas de Governança */}
      {readyProposals.map(proposalId => (
        <div key={proposalId} className="bg-yellow-100 border-l-4 border-yellow-500 p-4 mb-6 rounded">
          <p className="font-bold text-yellow-800">Proposta Aprovada!</p>
          <p className="text-yellow-700 text-sm mt-1">
            Sua proposta #{proposalId} atingiu o quórum e encerrou o prazo de votação. Ela já pode ser executada através de uma chamada externa.
          </p>
          <button 
            className="mt-3 bg-yellow-600 text-white px-4 py-2 rounded shadow hover:bg-yellow-700 transition"
            onClick={() => {/* Montar UserOp chamando execute(proposalId) */}}
          >
            Executar Proposta #{proposalId}
          </button>
        </div>
      ))}

      {/* Header / Status do Identity Registry */}
      <div className="flex flex-col md:flex-row justify-between items-start md:items-center border-b pb-4 mb-6">
        <div>
          <h1 className="text-3xl font-extrabold text-blue-700 tracking-tight">Dashboard do Aluno</h1>
          <div className="mt-2 flex items-center gap-2 flex-wrap">
            <p className="text-sm text-gray-600 font-mono bg-slate-100 px-3 py-1.5 rounded-md border border-slate-200 break-all">
              {studentAddress}
            </p>
            <button
              type="button"
              onClick={handleCopyAddress}
              className="px-3 py-1.5 text-xs font-semibold rounded-md bg-blue-600 text-white hover:bg-blue-700 transition"
            >
              Copiar endereco
            </button>
          </div>
        </div>
        <div className={`mt-2 md:mt-0 px-4 py-1 rounded-full text-sm font-semibold border ${isActive ? 'bg-green-100 text-green-800 border-green-200' : 'bg-red-100 text-red-800 border-red-200'}`}>
          {isActive ? 'Ativo no Registro' : 'Inativo no Registro'}
        </div>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
        {/* Painel de Saldo */}
        <div className="p-6 rounded-xl border border-blue-200 bg-gradient-to-br from-blue-50 via-indigo-50 to-cyan-50 shadow-sm flex flex-col justify-center">
          <h2 className="text-xs font-bold text-blue-800 uppercase tracking-[0.18em] mb-3">Saldo em Carteira</h2>
          <p className="text-5xl font-black text-blue-900 leading-tight tracking-tight tabular-nums">
            {balance}
            <span className="text-2xl font-semibold text-blue-700 ml-2">U$PT</span>
          </p>
          <p className="mt-3 text-sm text-blue-700/80">Disponivel para transferencias gasless na Sepolia.</p>
        </div>

        {/* Formulário de Transferência */}
        <div className="bg-gradient-to-b from-white to-slate-50 p-6 rounded-xl border border-slate-200 shadow-sm">
          <h2 className="text-lg font-semibold text-slate-800 mb-4">Enviar U$PT</h2>
          <form onSubmit={handleTransfer} className="space-y-4">
            <div>
              <label className="block text-xs font-semibold uppercase tracking-wider text-slate-600 mb-1.5">Destinatario</label>
              <input 
                type="text" 
                placeholder="0x..."
                value={receiverAddress}
                onChange={(e) => setReceiverAddress(e.target.value)}
                className="block w-full rounded-lg border-slate-300 bg-white text-slate-800 shadow-sm p-2.5 border focus:ring-2 focus:ring-blue-500 focus:border-blue-500 font-mono text-sm outline-none"
                required
              />
            </div>
            <div>
              <label className="block text-xs font-semibold uppercase tracking-wider text-slate-600 mb-1.5">Valor</label>
              <input 
                type="number" 
                step="0.01"
                min="0"
                placeholder="0.00"
                value={transferAmount}
                onChange={(e) => setTransferAmount(e.target.value)}
                className="block w-full rounded-lg border-slate-300 bg-white text-slate-800 shadow-sm p-2.5 border focus:ring-2 focus:ring-blue-500 focus:border-blue-500 outline-none"
                required
              />
            </div>
            <button 
              type="submit" 
              disabled={!isActive}
              className="w-full bg-gradient-to-r from-slate-900 to-blue-800 text-white font-semibold py-2.5 px-4 rounded-lg hover:from-slate-800 hover:to-blue-700 focus:ring-4 focus:ring-blue-200 disabled:opacity-50 disabled:cursor-not-allowed transition"
            >
              Assinar Transação (Gasless)
            </button>
            {transferStatus && (
              <p className={`text-sm mt-2 font-medium ${transferStatus.includes('Erro') ? 'text-red-600' : 'text-blue-600'}`}>
                {transferStatus}
              </p>
            )}
            {/* Bloco de Feedback e Link do Etherscan */}
            {(transferStatus || transferTxHash) && (
              <div className="mt-4 p-3 bg-gray-50 border border-gray-200 rounded-md">
                <p className={`text-sm font-medium ${transferStatus.includes('Erro') ? 'text-red-600' : 'text-green-600'}`}>
                  {transferStatus}
                </p>
                
                {transferTxHash && (
                  <a 
                    href={`https://sepolia.etherscan.io/tx/${transferTxHash}`} 
                    target="_blank" 
                    rel="noopener noreferrer"
                    className="inline-flex items-center mt-2 text-sm text-blue-600 hover:text-blue-800 hover:underline"
                  >
                    <svg className="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
                    </svg>
                    Ver transação no Etherscan
                  </a>
                )}
              </div>
            )}
          </form>
        </div>
      </div>
      </div>

      <Snackbar
        open={copySnackbar.open}
        autoHideDuration={2600}
        onClose={handleCloseSnackbar}
        anchorOrigin={{ vertical: 'bottom', horizontal: 'center' }}
      >
        <Alert onClose={handleCloseSnackbar} severity={copySnackbar.severity} variant="filled" sx={{ width: '100%' }}>
          {copySnackbar.message}
        </Alert>
      </Snackbar>
    </>
  );
}