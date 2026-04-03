import { useEffect, useState } from 'react';
import { Alert, Box, Typography, Paper, Container, CircularProgress, Snackbar } from '@mui/material';
import { useLocation, useNavigate } from 'react-router-dom';
import { useAuth } from '../../contexts/AuthContext';
import { GoogleLogin } from '@react-oauth/google';

export function Login() {
  const { handleLoginSuccess, isLoading } = useAuth();
  const location = useLocation();
  const navigate = useNavigate();
  const [logoutSnackbarOpen, setLogoutSnackbarOpen] = useState(false);

  useEffect(() => {
    if (location.state?.loggedOut) {
      setLogoutSnackbarOpen(true);
      navigate('/', { replace: true, state: null });
    }
  }, [location.state, navigate]);

  const handleCloseSnackbar = (_, reason) => {
    if (reason === 'clickaway') return;
    setLogoutSnackbarOpen(false);
  };

  return (
    <>
      <Container
        maxWidth="sm"
        sx={{
          minHeight: 'calc(100dvh - 2.5rem)',
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
        }}
      >
        <Paper elevation={4} sx={{ p: 5, width: '100%', textAlign: 'center', borderRadius: 4 }}>
          <Typography variant="h3" component="h1" fontWeight="800" gutterBottom color="primary.main">
            USPToken
          </Typography>
          <Typography variant="body1" color="text.secondary" sx={{ mb: 5 }}>
            A moeda social do ecossistema universitário.
          </Typography>

          <Box sx={{ display: 'flex', justifyContent: 'center' }}>
            {isLoading ? (
              <CircularProgress />
            ) : (
              /* Este componente substitui o seu Button antigo.
                 Ele gera o JWT que o seu handleLoginSuccess precisa.
              */
              <GoogleLogin
                onSuccess={handleLoginSuccess}
                onError={() => alert("Falha na autenticação com o Google")}
                useOneTap // Opcional: mostra um prompt flutuante se o usuário já estiver logado no Chrome
              />
            )}
          </Box>

          <Box sx={{ mt: 4, pt: 3, borderTop: 1, borderColor: 'divider' }}>
            <Typography variant="caption" color="text.secondary" sx={{ display: 'block' }}>
              Acesso restrito a contas institucionais.
            </Typography>
            <Typography variant="caption" color="text.secondary" sx={{ display: 'block', mt: 0.5 }}>
              Sua carteira digital será gerada automaticamente de forma segura.
            </Typography>
          </Box>
        </Paper>
      </Container>

      <Snackbar
        open={logoutSnackbarOpen}
        autoHideDuration={2600}
        onClose={handleCloseSnackbar}
        anchorOrigin={{ vertical: 'bottom', horizontal: 'center' }}
      >
        <Alert onClose={handleCloseSnackbar} severity="success" variant="filled" sx={{ width: '100%' }}>
          Sessao encerrada com sucesso.
        </Alert>
      </Snackbar>
    </>
  );
}