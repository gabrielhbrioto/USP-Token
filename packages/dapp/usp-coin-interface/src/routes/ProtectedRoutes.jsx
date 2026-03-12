import { Navigate } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';
// import { CircularProgress, Box } from '@mui/material'; // Descomente ao usar MUI

export function ProtectedRoute({ children }) {
  const { user, isLoading } = useAuth();

  // Enquanto verifica o estado (evita redirecionar prematuramente)
  if (isLoading) {
    return (
      <div style={{ display: 'flex', justifyContent: 'center', marginTop: '20vh' }}>
        {/* <CircularProgress /> */}
        <p>Carregando dados da blockchain...</p>
      </div>
    );
  }

  // Se não tem usuário, manda de volta pra tela de login
  if (!user) {
    return <Navigate to="/" replace />;
  }

  // Se o usuário está logado mas não está registrado no IdentityRegistry, trava na tela de cadastro
  // (Isso impede que ele force a URL /wallet pelo navegador sem ter a identidade cunhada)
  if (user && !user.isRegistered && window.location.pathname !== '/auth') {
    return <Navigate to="/auth" replace />;
  }

  // Se passou por tudo, renderiza a tela (ex: Dashboard)
  return children;
}