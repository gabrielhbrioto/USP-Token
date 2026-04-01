import React, { useState } from 'react';
import { useAuth } from '../../contexts/AuthContext';

export function Certificates() {
  // Estado para armazenar os dados do formulário
  // Preenchido com valores padrão para facilitar os testes
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
  // Função disparada ao submeter o formulário
  const handleSubmit = async (e) => {
    e.preventDefault();
    setIsLoading(true);
    setError('');
    setTxHash(''); // Limpa o hash anterior

    try {
      // 1. Dispara a emissão oficial na Blockchain (IPFS + Mint via Relayer)
      const issueResponse = await fetch('http://localhost:8080/certificate/issue', {
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
      const generateResponse = await fetch('http://localhost:8080/certificate/generate', {
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
    <div className="min-h-screen bg-gray-50 p-8 font-sans">
      <div className="max-w-6xl mx-auto">
        <h1 className="text-3xl font-bold text-gray-800 mb-8">Emissão de Certificado</h1>

        <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
          
          {/* LADO ESQUERDO: Formulário */}
          <div className="bg-white p-6 rounded-xl shadow-sm border border-gray-200">
            <h2 className="text-xl font-semibold mb-4 text-gray-700">Dados do Aluno</h2>
            <form onSubmit={handleSubmit} className="space-y-4">
              
              <div>
                <label className="block text-sm font-medium text-gray-700">Nome do Aluno</label>
                <input type="text" name="studentName" value={formData.studentName} onChange={handleChange} required
                  className="mt-1 w-full rounded-md border-gray-300 shadow-sm p-2 border focus:ring-blue-500 focus:border-blue-500" />
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700">Nome do Curso</label>
                <input type="text" name="courseName" value={formData.courseName} onChange={handleChange} required
                  className="mt-1 w-full rounded-md border-gray-300 shadow-sm p-2 border focus:ring-blue-500 focus:border-blue-500" />
              </div>

              <div className="grid grid-cols-3 gap-4">
                <div>
                  <label className="block text-sm font-medium text-gray-700">Carga Horária</label>
                  <input type="number" name="hours" value={formData.hours} onChange={handleChange} required
                    className="mt-1 w-full rounded-md border-gray-300 shadow-sm p-2 border" />
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700">Data Início</label>
                  <input type="text" name="startDate" value={formData.startDate} onChange={handleChange} required
                    className="mt-1 w-full rounded-md border-gray-300 shadow-sm p-2 border placeholder-gray-400" placeholder="DD/MM/AAAA" />
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700">Data Fim</label>
                  <input type="text" name="endDate" value={formData.endDate} onChange={handleChange} required
                    className="mt-1 w-full rounded-md border-gray-300 shadow-sm p-2 border placeholder-gray-400" placeholder="DD/MM/AAAA" />
                </div>
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700">Nota / Aproveitamento (%)</label>
                <input type="number" name="gradePercent" value={formData.gradePercent} onChange={handleChange} required
                  className="mt-1 w-full rounded-md border-gray-300 shadow-sm p-2 border" />
              </div>

              {/* Botão de Submit */}
              <button
                type="submit"
                disabled={isLoading}
                className={`w-full mt-6 text-white font-bold py-3 px-4 rounded-lg transition-colors ${
                  isLoading ? 'bg-blue-400 cursor-not-allowed' : 'bg-blue-600 hover:bg-blue-700'
                }`}
              >
                {isLoading ? 'Gerando Certificado...' : 'Gerar Certificado'}
              </button>

              {error && <p className="text-red-500 text-sm mt-2">{error}</p>}
            </form>
          </div>

          {/* LADO DIREITO: Preview da Imagem */}
          <div className="bg-white p-6 rounded-xl shadow-sm border border-gray-200 flex flex-col items-center justify-center min-h-[400px]">
            {txHash && (
              <div className="w-full mb-4 p-4 bg-green-50 border border-green-200 rounded-lg text-center">
                <p className="text-green-800 font-semibold mb-1">✅ NFT Emitido com Sucesso!</p>
                <a 
                  href={`https://sepolia.etherscan.io/tx/${txHash}`} 
                  target="_blank" 
                  rel="noopener noreferrer"
                  className="text-sm text-blue-600 hover:underline"
                >
                  Ver transação no Etherscan
                </a>
              </div>
            )}
            {generatedImage ? (
              <div className="w-full flex flex-col items-center">
                <h2 className="text-xl font-semibold mb-4 text-gray-700 self-start">Visualização</h2>
                <img 
                  src={generatedImage} 
                  alt="Certificado Gerado" 
                  className="w-full h-auto rounded-md shadow-md border border-gray-300"
                />
                <a 
                  href={generatedImage} 
                  download={`Certificado_${formData.studentName.replace(/\s+/g, '_')}.jpg`}
                  className="mt-6 bg-green-600 hover:bg-green-700 text-white font-bold py-2 px-6 rounded-lg transition-colors"
                >
                  Baixar Imagem
                </a>
              </div>
            ) : (
              <div className="text-center text-gray-400">
                <svg className="mx-auto h-12 w-12 text-gray-300 mb-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
                  <path vectorEffect="non-scaling-stroke" strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 13h6m-3-3v6m-9 1V7a2 2 0 012-2h6l2 2h6a2 2 0 012 2v8a2 2 0 01-2 2H5a2 2 0 01-2-2z" />
                </svg>
                <p>Preencha os dados e clique em "Visualizar" para gerar a imagem.</p>
              </div>
            )}
          </div>

        </div>
      </div>
    </div>
  );
}