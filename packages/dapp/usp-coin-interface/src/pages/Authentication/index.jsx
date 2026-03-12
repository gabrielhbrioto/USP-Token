// src/pages/Authentication.jsx
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Box, Button, TextField, Typography, Paper, Container, MenuItem, CircularProgress } from '@mui/material';
import { ethers } from 'ethers';
import { useAuth } from '../../contexts/AuthContext';

// Lista de exemplos para o select
const INSTITUTOS = [
  'EESC - Escola de Engenharia de São Carlos',
  'ICMC - Instituto de Ciências Matemáticas e de Computação',
  'IFSC - Instituto de Física de São Carlos',
  'IQSC - Instituto de Química de São Carlos',
  'POLI - Escola Politécnica'
];

export function Authentication() {
  const { user, walletAddress } = useAuth();
  const navigate = useNavigate();
  
  const [uspNumber, setUspNumber] = useState('');
  const [instituto, setInstituto] = useState('');
  const [isSubmitting, setIsSubmitting] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setIsSubmitting(true);

    try {
      // 1. Preparação dos Dados
      // Estes são os dados que o contrato IdentityRegistry vai precisar
      const payload = {
        studentAddress: walletAddress,
        uspNumber: uspNumber,
        instituto: instituto,
        timestamp: Math.floor(Date.now() / 1000)
      };

      // 2. A Mágica da Embedded Wallet (Assinatura Local)
      // Como ainda não plugamos o Web3Auth/Privy, vamos simular a assinatura
      // gerando uma carteira temporária apenas para criar o hash criptográfico.
      // Na produção real, o SDK da carteira embutida fará isso em background.
      const dummyPrivateKey = "0x0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"; 
      const wallet = new ethers.Wallet(dummyPrivateKey);

      // Assinamos uma mensagem padronizada com os dados do aluno
      const messageToSign = JSON.stringify(payload);
      const signature = await wallet.signMessage(messageToSign);

      console.log("Assinatura gerada localmente:", signature);

      // 3. Envio para o Relayer em Go (Backend)
      // O frontend NUNCA envia a chave privada, apenas os dados públicos e a assinatura!
      const relayerUrl = `${import.meta.env.VITE_RELAYER_URL || 'http://localhost:8080'}/register`;
      
      // Descomente este bloco quando o backend em Go estiver rodando
      const response = await fetch(relayerUrl, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          data: payload,
          signature: signature 
        })
      });

      if (!response.ok) {
        throw new Error('Falha ao processar a transação no Relayer');
      }
      

      // Simulando o tempo de processamento do Relayer + Mineração no Anvil
      await new Promise(resolve => setTimeout(resolve, 2500));

      // 4. Sucesso e Redirecionamento
      // Aqui, o ideal é atualizar o `localStorage` e o estado do Context para isRegistered: true
      alert("Identidade registrada com sucesso na blockchain!");
      navigate('/wallet');

    } catch (error) {
      console.error("Erro durante o registro:", error);
      alert("Ocorreu um erro ao registrar sua identidade.");
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <Container maxWidth="sm">
      <Box sx={{ minHeight: '100vh', display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
        <Paper elevation={4} sx={{ p: 5, width: '100%', borderRadius: 4 }}>
          <Typography variant="h4" component="h1" fontWeight="bold" gutterBottom color="primary.main" textAlign="center">
            Complete seu Perfil
          </Typography>
          <Typography variant="body2" color="text.secondary" sx={{ mb: 4 }} textAlign="center">
            Para cunhar sua identidade no <strong>IdentityRegistry</strong>, precisamos de algumas informações acadêmicas.
          </Typography>

          <Box sx={{ mb: 3, p: 2, bgcolor: 'grey.100', borderRadius: 2 }}>
            <Typography variant="caption" color="text.secondary" display="block">
              Carteira Vinculada:
            </Typography>
            <Typography variant="body2" fontWeight="bold" sx={{ wordBreak: 'break-all' }}>
              {walletAddress || 'Carregando endereço...'}
            </Typography>
          </Box>

          <form onSubmit={handleSubmit}>
            <TextField
              fullWidth
              label="Número USP"
              variant="outlined"
              margin="normal"
              required
              type="number"
              value={uspNumber}
              onChange={(e) => setUspNumber(e.target.value)}
              disabled={isSubmitting}
            />

            <TextField
              fullWidth
              select
              label="Instituto"
              variant="outlined"
              margin="normal"
              required
              value={instituto}
              onChange={(e) => setInstituto(e.target.value)}
              disabled={isSubmitting}
            >
              {INSTITUTOS.map((inst) => (
                <MenuItem key={inst} value={inst}>
                  {inst}
                </MenuItem>
              ))}
            </TextField>

            <Button
              type="submit"
              variant="contained"
              size="large"
              fullWidth
              disabled={isSubmitting}
              sx={{ mt: 4, py: 1.5, fontSize: '1.1rem', borderRadius: 2 }}
            >
              {isSubmitting ? <CircularProgress size={24} color="inherit" /> : 'Cunhar Identidade na Rede'}
            </Button>
          </form>

          <Typography variant="caption" color="text.secondary" sx={{ display: 'block', mt: 3, textAlign: 'center' }}>
            Esta transação não possui custos (Gasless).
            <br />
            As taxas serão cobertas pelo USPPaymaster.
          </Typography>
        </Paper>
      </Box>
    </Container>
  );
}