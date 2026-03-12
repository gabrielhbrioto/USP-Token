import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.jsx'
import { GoogleOAuthProvider } from '@react-oauth/google';

const clientId = import.meta.env.VITE_GOOGLE_CLIENT_ID;

import { WagmiProvider, createConfig, http } from 'wagmi';
import { localhost, sepolia } from 'wagmi/chains'; // Importe a rede que você está usando
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

// 2. Criação da configuração do Wagmi
const wagmiConfig = createConfig({
  chains: [localhost], // Coloque aqui a rede que você está usando (ex: localhost para anvil/hardhat, ou sepolia)
  transports: {
    // O http() pega a URL padrão da rede, mas você pode passar a sua RPC customizada: http('http://127.0.0.1:8545')
    [localhost.id]: http(), 
  },
});

// 3. Instância do QueryClient
const queryClient = new QueryClient();

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <WagmiProvider config={wagmiConfig}>
      <QueryClientProvider client={queryClient}>
        <GoogleOAuthProvider clientId={clientId}>
          <App />
        </GoogleOAuthProvider>
      </QueryClientProvider>
    </WagmiProvider>
  </StrictMode>,
)
