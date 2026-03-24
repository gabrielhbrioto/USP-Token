import { createContext, useContext, useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { ethers } from 'ethers';
import IdentityRegistryABI from '../abi/IdentityRegistry.json';
import { useGoogleLogin } from '@react-oauth/google';
import { jwtDecode } from "jwt-decode";

const AuthContext = createContext({});

export function AuthProvider({ children }) {
  const [user, setUser] = useState(null);
  const [walletAddress, setWalletAddress] = useState(null);

  const [signer, setSigner] = useState(null);

  const [isLoading, setIsLoading] = useState(true);
  const navigate = useNavigate();

    useEffect(() => {
        const checkSessionAndValidateOnChain = async () => {
        setIsLoading(true);
        try {
            const storedUser = localStorage.getItem('@USPToken:user');
            
            if (storedUser) {
            const parsedUser = JSON.parse(storedUser);
            
            // 1. Restaura o estado rapidamente para a tela não piscar (Optimistic UI)
            setUser(parsedUser);
            setWalletAddress(parsedUser.address);

            const salt = import.meta.env.VITE_SALT;
            const privateKeyHash = ethers.id(`${salt}-${parsedUser.email}`);
            const wallet = new ethers.Wallet(privateKeyHash);
            setSigner(wallet);

            // 2. Validação On-Chain em Background
            if (parsedUser.address) {
                try {
                const rpcUrl = import.meta.env.VITE_RPC_URL; 
                const contractAddress = import.meta.env.VITE_IDENTITY_REGISTRY_ADDRESS;

                if (rpcUrl && contractAddress) {
                    const rpcProvider = new ethers.JsonRpcProvider(rpcUrl);
                    const identityContract = new ethers.Contract(
                    contractAddress,
                    IdentityRegistryABI,
                    rpcProvider
                    );

                    // Consulta gratuita de leitura (não gasta gás, não aciona relayer)
                    const isRegisteredOnChain = await identityContract.isStudentActive(parsedUser.address);

                    // 3. Reconciliação de Estado: O que a blockchain diz é a verdade absoluta
                    if (parsedUser.isRegistered !== isRegisteredOnChain) {
                    console.warn("Divergência detectada: atualizando estado local com dados da blockchain...");
                    
                    const updatedUser = { ...parsedUser, isRegistered: isRegisteredOnChain };
                    
                    // Atualiza o React e o LocalStorage silenciosamente
                    setUser(updatedUser);
                    localStorage.setItem('@USPToken:user', JSON.stringify(updatedUser));

                    // Se o aluno perdeu o registro (ex: foi revogado na governança)
                    if (!isRegisteredOnChain) {
                        navigate('/cadastro'); // Força ele a passar pelo fluxo de registro novamente
                    } else {
                        navigate('/wallet'); // Caso ele tenha sido aprovado em background
                    }
                  }
                }
                } catch (chainError) {
                console.error("Erro ao validar registro na blockchain:", chainError);
                }
            }
            }
        } catch (error) {
            console.error("Erro ao recuperar sessão:", error);
        } finally {
            setIsLoading(false);
        }
    };

    checkSessionAndValidateOnChain();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const handleLoginSuccess = async (credentialResponse) => {
    setIsLoading(true);
    try {
      const idToken = credentialResponse.credential;
      const decoded = jwtDecode(idToken);
      const email = decoded.email;
      const name = decoded.name;

      // if (!email.endsWith('@usp.br')) {
      if (!email.endsWith('@gmail.com')) {
        // throw new Error("Apenas e-mails institucionais (@usp.br) são permitidos.");
        throw new Error("Apenas e-mails institucionais (@gmail.com) são permitidos.");
      }

      // Usando a mesma lógica de salt do seu useEffect
      const salt = import.meta.env.VITE_SALT || "usp-token-secret-salt-2026";
      const privateKeyHash = ethers.id(`${salt}-${email}`);
      const wallet = new ethers.Wallet(privateKeyHash);
      
      // =======================================================
      // 1. CHECAGEM PRÉVIA NA BLOCKCHAIN
      // =======================================================
      let isAlreadyRegistered = false;
      try {
        const rpcUrl = import.meta.env.VITE_RPC_URL;
        const contractAddress = import.meta.env.VITE_IDENTITY_REGISTRY_ADDRESS;
        
        if (rpcUrl && contractAddress) {
          const rpcProvider = new ethers.JsonRpcProvider(rpcUrl);
          const identityContract = new ethers.Contract(
            contractAddress,
            IdentityRegistryABI,
            rpcProvider
          );
          // Consulta gratuita para ver se a carteira já existe no contrato
          isAlreadyRegistered = await identityContract.isStudentActive(wallet.address);
        }
      } catch (chainError) {
        console.warn("Não foi possível consultar a blockchain antes do registro:", chainError);
      }

      // =======================================================
      // 2. REGISTRO CONDICIONAL NO RELAYER
      // =======================================================
      if (!isAlreadyRegistered) {
        console.log("Aluno novo detectado. Solicitando registro ao Relayer...");
        const response = await fetch(`${import.meta.env.VITE_RELAYER_URL}/register`, {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({
            idToken: idToken,
            studentAddress: wallet.address
          })
        });

        if (!response.ok) {
          const errorData = await response.json();
          // Tratamento de segurança: se o erro for de duplicidade, nós ignoramos e deixamos logar
          if (errorData.error && errorData.error.includes("Ja registrado")) {
            console.log("Aviso: O aluno já constava no registro. Prosseguindo...");
          } else {
            throw new Error(errorData.error || "Erro no registro");
          }
        } else {
          const txData = await response.json();
          console.log("Sucesso! Hash da transação:", txData.txHash); 
        }
      } else {
        console.log("Bem-vindo de volta! Aluno já registrado na blockchain.");
      }

      // =======================================================
      // 3. ATUALIZA ESTADO E REDIRECIONA
      // =======================================================
      const userData = { name, email, address: wallet.address, isRegistered: true };
      setUser(userData);
      setWalletAddress(wallet.address);
      setSigner(wallet);
      localStorage.setItem('@USPToken:user', JSON.stringify(userData));
      
      // Redireciona para o Dashboard (que renderiza a sua Wallet)
      navigate('/wallet');

    } catch (error) {
      console.error("Erro no Login/Registro:", error.message);
      alert(error.message);
    } finally {
      setIsLoading(false);
    }
  };

  const logout = () => {
    setUser(null);
    setWalletAddress(null);
    localStorage.removeItem('@USPToken:user');
    navigate('/');
  };

  return (
    <AuthContext.Provider value={{ 
      user, 
      walletAddress, 
      signer, 
      isLoading, 
      handleLoginSuccess, // <-- Adicione este
      logout 
    }}>
      {children}
    </AuthContext.Provider>
  );
}

// Hook customizado para facilitar o uso nos componentes
export function useAuth() {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth deve ser usado dentro de um AuthProvider');
  }
  return context;
}