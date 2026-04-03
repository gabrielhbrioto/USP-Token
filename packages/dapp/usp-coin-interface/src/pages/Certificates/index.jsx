import React, { useState } from 'react';
import { useAuth } from '../../contexts/AuthContext';

export function Certificates() {
  const { walletAddress: studentAddress } = useAuth(); 
  
  const [formData, setFormData] = useState({
    studentAddress: studentAddress,
    studentName: 'GABRIEL HENRIQUE BRIOTO',
    courseName: 'ENGENHARIA DE COMPUTAÇÃO WEB3',
    hours: '40',
    startDate: '17/03/2026',
    endDate: '17/04/2026',
    gradePercent: '95',
    directorName: 'Prof. Dr. Fulano de Tal',
    directorTitle: 'Diretor da Poli-USP',
    coordinatorName: 'Ciclana de Souza',
    coordinatorTitle: 'Coordenadora do Curso',
    emissionDay: new Date().getDate().toString().padStart(2, '0'),
    emissionMonth: (new Date().getMonth() + 1).toString().padStart(2, '0'),
    emissionYear: new Date().getFullYear().toString(),
    verificationCode: 'USP-' + Math.floor(Math.random() * 10000000),
  });

  // Estados de controle da interface
  const [isLoading, setIsLoading] = useState(false);
  const [generatedImage, setGeneratedImage] = useState(null);
  const [error, setError] = useState('');
  const [txHash, setTxHash] = useState('');

  // Atualiza o estado conforme o usuário digita
  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  // Função disparada ao submeter o formulário
  const handleSubmit = async (e) => {
    e.preventDefault();
    setIsLoading(true);
    setError('');
    setTxHash(''); // Limpa o hash anterior

    try {
      // 1. Dispara a emissão oficial na Blockchain (IPFS + Mint via Relayer)
      const issueResponse = await fetch(`${import.meta.env.VITE_RELAYER_URL}/certificate/issue`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });

      if (!issueResponse.ok) {
        const errData = await issueResponse.json();
        throw new Error(errData.error || 'Falha ao emitir o certificado na blockchain.');
      }

      // Extrai o JSON com os dados do Mint
      const issueData = await issueResponse.json();
      setTxHash(issueData.txHash); // Salva o hash para mostrar na tela

      // 2. Busca a imagem gerada em memória para exibição/download no navegador
      const generateResponse = await fetch(`${import.meta.env.VITE_RELAYER_URL}/certificate/generate`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });

      if (!generateResponse.ok) {
        throw new Error('Certificado emitido, mas falha ao carregar a visualização.');
      }

      // Como a rota /generate retorna a imagem diretamente, pegamos o Blob
      const imageBlob = await generateResponse.blob();
      
      // Cria uma URL temporária no navegador para exibir a imagem gerada
      const imageUrl = URL.createObjectURL(imageBlob);
      setGeneratedImage(imageUrl);

    } catch (err) {
      setError(err.message);
      console.error(err);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-neutral-900 p-8 font-sans text-gray-200">
      <div className="max-w-7xl mx-auto">
        <h1 className="text-4xl font-bold text-white mb-8">Emissão de Certificado</h1>

        <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
          
          {/* LADO ESQUERDO: Formulário */}
          <div className="bg-neutral-800 p-6 rounded-xl shadow-lg border border-neutral-700">
            <h2 className="text-xl font-semibold mb-6 text-gray-300">Dados do Aluno</h2>
            <form onSubmit={handleSubmit} className="space-y-5">
              
              <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div>
                  <label className="block text-sm font-medium text-gray-400 mb-1">Nome do Aluno</label>
                  <div className="relative">
                    <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-gray-500">
                      <svg width="16" height="16" className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path></svg>
                    </div>
                    <input type="text" name="studentName" value={formData.studentName} onChange={handleChange} required
                      className="w-full bg-transparent border border-neutral-600 rounded-md py-2 pl-10 pr-3 text-white placeholder-gray-500 focus:outline-none focus:ring-1 focus:ring-yellow-500 focus:border-yellow-500" />
                  </div>
                </div>

                <div>
                  <label className="block text-sm font-medium text-gray-400 mb-1">Nome do Curso</label>
                  <div className="relative">
                    <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-gray-500">
                      <svg width="16" height="16" className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path d="M12 14l9-5-9-5-9 5 9 5z"></path><path d="M12 14l6.16-3.422a12.083 12.083 0 01.665 6.479A11.952 11.952 0 0012 20.055a11.952 11.952 0 00-6.824-2.998 12.078 12.078 0 01.665-6.479L12 14z"></path><path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M12 14l9-5-9-5-9 5 9 5zm0 0l6.16-3.422a12.083 12.083 0 01.665 6.479A11.952 11.952 0 0012 20.055a11.952 11.952 0 00-6.824-2.998 12.078 12.078 0 01.665-6.479L12 14zm-4 6v-7.5l4-2.222"></path></svg>
                    </div>
                    <input type="text" name="courseName" value={formData.courseName} onChange={handleChange} required
                      className="w-full bg-transparent border border-neutral-600 rounded-md py-2 pl-10 pr-3 text-white placeholder-gray-500 focus:outline-none focus:ring-1 focus:ring-yellow-500 focus:border-yellow-500" />
                  </div>
                </div>
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-400 mb-1">Carga Horária (horas)</label>
                <div className="relative">
                  <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-gray-500">
                    <svg width="16" height="16" className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
                  </div>
                  <input type="number" name="hours" value={formData.hours} onChange={handleChange} required
                    className="w-full bg-transparent border border-neutral-600 rounded-md py-2 pl-10 pr-3 text-white placeholder-gray-500 focus:outline-none focus:ring-1 focus:ring-yellow-500 focus:border-yellow-500" />
                </div>
              </div>

              <div className="grid grid-cols-2 gap-4">
                <div>
                  <label className="block text-sm font-medium text-gray-400 mb-1">Data Início</label>
                  <div className="relative">
                    <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-gray-500">
                      <svg width="16" height="16" className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path></svg>
                    </div>
                    <input type="text" name="startDate" value={formData.startDate} onChange={handleChange} required
                      className="w-full bg-transparent border border-neutral-600 rounded-md py-2 pl-10 pr-3 text-white placeholder-gray-500 focus:outline-none focus:ring-1 focus:ring-yellow-500 focus:border-yellow-500" placeholder="DD/MM/AAAA" />
                  </div>
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-400 mb-1">Data Fim</label>
                  <div className="relative">
                    <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-gray-500">
                      <svg width="16" height="16" className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path></svg>
                    </div>
                    <input type="text" name="endDate" value={formData.endDate} onChange={handleChange} required
                      className="w-full bg-transparent border border-neutral-600 rounded-md py-2 pl-10 pr-3 text-white placeholder-gray-500 focus:outline-none focus:ring-1 focus:ring-yellow-500 focus:border-yellow-500" placeholder="DD/MM/AAAA" />
                  </div>
                </div>
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-400 mb-1">Nota / Aproveitamento (%)</label>
                <div className="relative">
                  <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-gray-500 font-bold">
                    %
                  </div>
                  <input type="number" name="gradePercent" value={formData.gradePercent} onChange={handleChange} required
                    className="w-full bg-transparent border border-neutral-600 rounded-md py-2 pl-10 pr-3 text-white placeholder-gray-500 focus:outline-none focus:ring-1 focus:ring-yellow-500 focus:border-yellow-500" />
                </div>
              </div>

              {/* Botão de Submit */}
              <button
                type="submit"
                disabled={isLoading}
                className={`w-full mt-6 text-neutral-900 font-semibold py-3 px-4 rounded-md transition-colors ${
                  isLoading ? 'bg-yellow-600/50 cursor-not-allowed' : 'bg-yellow-500 hover:bg-yellow-600'
                }`}
              >
                {isLoading ? 'Gerando Certificado...' : 'Gerar Certificado'}
              </button>

              {error && <p className="text-red-400 text-sm mt-2">{error}</p>}
            </form>
          </div>

          {/* LADO DIREITO: Preview da Imagem */}
          <div className="bg-neutral-800 p-6 rounded-xl shadow-lg border border-neutral-700 flex flex-col items-center justify-center min-h-[400px]">
            {txHash && (
              <div className="w-full mb-4 p-4 bg-green-900/30 border border-green-800 rounded-md text-left">
                <p className="text-green-400 font-medium mb-1">✅ NFT Emitido com Sucesso!</p>
                <p className="text-sm text-gray-300">
                  Transação no Etherscan: <a 
                    href={`https://sepolia.etherscan.io/tx/${txHash}`} 
                    target="_blank" 
                    rel="noopener noreferrer"
                    className="text-green-400 hover:underline"
                  >
                    {txHash.substring(0, 6)}...{txHash.substring(txHash.length - 4)}
                  </a>
                </p>
              </div>
            )}
            
            {generatedImage ? (
              <div className="w-full flex flex-col items-center">
                <div className="w-full bg-neutral-900 border border-neutral-700 rounded-md p-4 mb-4">
                  <img 
                    src={generatedImage} 
                    alt="Certificado Gerado" 
                    className="w-full h-auto object-contain border border-neutral-800"
                  />
                </div>
                <a 
                  href={generatedImage} 
                  download={`Certificado_${formData.studentName.replace(/\s+/g, '_')}.jpg`}
                  className="w-full text-center bg-green-600 hover:bg-green-700 text-white font-semibold py-3 px-6 rounded-md transition-colors"
                >
                  Baixar Imagem
                </a>
              </div>
            ) : (
              <div className="text-center text-gray-500">
                <svg width="40" height="40" className="mx-auto h-10 w-10 text-gray-600 mb-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
                  <path vectorEffect="non-scaling-stroke" strokeLinecap="round" strokeLinejoin="round" strokeWidth={1.5} d="M9 13h6m-3-3v6m-9 1V7a2 2 0 012-2h6l2 2h6a2 2 0 012 2v8a2 2 0 01-2 2H5a2 2 0 01-2-2z" />
                </svg>
                <p>Preencha os dados e clique em "Gerar Certificado" para visualizar.</p>
              </div>
            )}
          </div>

        </div>
      </div>
    </div>
  );
}