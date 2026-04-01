import React, { useState, useEffect } from 'react';
import { useReadContracts, usePublicClient } from 'wagmi';
import { formatUnits, parseUnits } from 'viem';

// Importação dos ABIs (assumindo que foram compilados para JSON)
import IdentityRegistryABI from '../../abi/IdentityRegistry.json';
import USPTokenABI from '../../abi/USPToken.json'; 
import StudentGovernanceABI from '../../abi/StudentGovernance.json';
import { useAccount, useSwitchChain } from 'wagmi';
import { sepolia } from 'wagmi/chains';
import { useAuth } from '../../contexts/AuthContext';

// Variáveis de ambiente do Vite
const IDENTITY_REGISTRY_ADDR = import.meta.env.VITE_IDENTITY_REGISTRY_ADDRESS;
const USP_TOKEN_ADDR = import.meta.env.VITE_USP_TOKEN_ADDRESS;
const GOVERNANCE_ADDR = import.meta.env.VITE_GOVERNANCE_ADDRESS;

export function Wallet() {
  const { walletAddress: studentAddress } = useAuth(); 
  const publicClient = usePublicClient();
  const [receiverAddress, setReceiverAddress] = useState('');
  const [transferAmount, setTransferAmount] = useState('');
  const [transferStatus, setTransferStatus] = useState('');
  const [readyProposals, setReadyProposals] = useState([]);
  const { chainId, isConnected } = useAccount();
  const { switchChain } = useSwitchChain();
  const isWrongNetwork = isConnected && chainId !== sepolia.id;

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

  // 2. Validação e Envio da Transferência via Account Abstraction
  const handleTransfer = async (e) => {
    e.preventDefault();
    setTransferStatus('Validando destinatário...');

    try {
      // Regra de Negócio: O destinatário DEVE ser um aluno ativo no IdentityRegistry [cite: 7, 8]
      const isReceiverActive = await publicClient.readContract({
        address: IDENTITY_REGISTRY_ADDR,
        abi: IdentityRegistryABI,
        functionName: 'isStudentActive',
        args: [receiverAddress]
      });

      if (!isReceiverActive) {
        setTransferStatus('Erro: O endereço de destino não é um aluno ativo no registro.');
        return;
      }

      setTransferStatus('Montando UserOperation...');
      
      // Lógica de integração com o Relayer em Go
      const callData = encodeFunctionData({ abi: USPTokenABI, functionName: 'transfer', args: [receiverAddress, parseUnits(transferAmount, 18)] });
      const userOp = await buildUserOp(studentAddress, USP_TOKEN_ADDR, callData);
      await fetch(import.meta.env.VITE_RELAYER_URL, { method: 'POST', body: JSON.stringify(userOp) });
      
      setTransferStatus('UserOperation enviada ao relayer com sucesso!');
    } catch (error) {
      setTransferStatus(`Erro na transação: ${error.message}`);
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
    <div className="max-w-4xl mx-auto p-6 bg-white shadow-md rounded-lg mt-8">
      
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
          <h1 className="text-2xl font-bold text-gray-800">Dashboard do Aluno</h1>
          <p className="text-sm text-gray-500 font-mono mt-1">{studentAddress}</p>
        </div>
        <div className={`mt-2 md:mt-0 px-4 py-1 rounded-full text-sm font-semibold border ${isActive ? 'bg-green-100 text-green-800 border-green-200' : 'bg-red-100 text-red-800 border-red-200'}`}>
          {isActive ? 'Ativo no Registro' : 'Inativo no Registro'}
        </div>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
        {/* Painel de Saldo */}
        <div className="bg-blue-50 p-6 rounded-lg border border-blue-100 flex flex-col justify-center">
          <h2 className="text-sm font-semibold text-blue-800 uppercase tracking-wider mb-2">Saldo em Carteira</h2>
          <p className="text-4xl font-extrabold text-blue-900">{balance} <span className="text-xl font-medium">U$PT</span></p>
        </div>

        {/* Formulário de Transferência */}
        <div className="bg-gray-50 p-6 rounded-lg border border-gray-200">
          <h2 className="text-lg font-semibold text-gray-800 mb-4">Enviar U$PT</h2>
          <form onSubmit={handleTransfer} className="space-y-4">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">Destinatário</label>
              <input 
                type="text" 
                placeholder="0x..."
                value={receiverAddress}
                onChange={(e) => setReceiverAddress(e.target.value)}
                className="block w-full rounded-md border-gray-300 shadow-sm p-2.5 border focus:ring-blue-500 focus:border-blue-500 font-mono text-sm"
                required
              />
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">Valor</label>
              <input 
                type="number" 
                step="0.01"
                min="0"
                placeholder="0.00"
                value={transferAmount}
                onChange={(e) => setTransferAmount(e.target.value)}
                className="block w-full rounded-md border-gray-300 shadow-sm p-2.5 border focus:ring-blue-500 focus:border-blue-500"
                required
              />
            </div>
            <button 
              type="submit" 
              disabled={!isActive}
              className="w-full bg-gray-900 text-white font-medium py-2.5 px-4 rounded-md hover:bg-gray-800 focus:ring-4 focus:ring-gray-300 disabled:opacity-50 disabled:cursor-not-allowed transition"
            >
              Assinar Transação (Gasless)
            </button>
            {transferStatus && (
              <p className={`text-sm mt-2 font-medium ${transferStatus.includes('Erro') ? 'text-red-600' : 'text-blue-600'}`}>
                {transferStatus}
              </p>
            )}
          </form>
        </div>
      </div>
    </div>
  );
}