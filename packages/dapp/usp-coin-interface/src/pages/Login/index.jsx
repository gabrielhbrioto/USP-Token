import { Box, Typography, Paper, Container, CircularProgress } from '@mui/material';
import { useAuth } from '../../contexts/AuthContext';
import { GoogleLogin } from '@react-oauth/google';
export function Login() {
  const { handleLoginSuccess, isLoading } = useAuth(); 

  return (
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
  );
}