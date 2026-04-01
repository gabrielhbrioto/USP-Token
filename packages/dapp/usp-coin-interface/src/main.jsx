import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.jsx'
import { GoogleOAuthProvider } from '@react-oauth/google';
import { createConfig, http, WagmiProvider } from 'wagmi';
import { sepolia } from 'wagmi/chains';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

const clientId = import.meta.env.VITE_GOOGLE_CLIENT_ID;
const RPC_URL = import.meta.env.VITE_INFURA_RPC_URL || 'https://rpc.sepolia.org';


export const wagmiConfig = createConfig({
  // Trava o dApp na Sepolia
  chains: [sepolia], 
  transports: {
    [sepolia.id]: http(RPC_URL),
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
