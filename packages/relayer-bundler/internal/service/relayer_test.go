package service

import (
	"bytes"
	"context"
	"fmt"
	"image/jpeg"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gabrielhbrioto/Usp-Token/packages/relayer-bundler/internal/config"
	"github.com/gabrielhbrioto/Usp-Token/packages/relayer-bundler/pkg/contracts"
)

// --- MOCKS ---

// Simulando o Client da Blockchain
type MockChainClient struct{}

func (m *MockChainClient) ChainID(ctx context.Context) (*big.Int, error) {
	return big.NewInt(11155111), nil // Simula o ID da Sepolia
}

// Simulando o Contrato EntryPoint
type MockEntryPoint struct{}

func (m *MockEntryPoint) HandleOps(opts *bind.TransactOpts, ops []contracts.UserOperation, beneficiary common.Address) (*types.Transaction, error) {
	// Cria uma transação falsa genérica apenas para gerar um Hash no teste
	fakeTx := types.NewTx(&types.LegacyTx{
		Nonce:    1,
		GasPrice: big.NewInt(1000000000),
		Gas:      21000,
		To:       &beneficiary,
		Value:    big.NewInt(0),
	})
	return fakeTx, nil
}

// --- TESTES DE BLOCKCHAIN ---

func TestSendBundle_Success(t *testing.T) {
	// 1. Prepara a configuração com uma chave privada válida (aleatória para teste)
	cfg := &config.Config{
		PrivateKey: "fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19",
	}

	// 2. Injeta os Mocks no serviço
	svc := NewRelayerService(cfg, &MockChainClient{}, &MockEntryPoint{}, nil)

	// 3. Cria uma UserOperation vazia (irrelevante para este teste de infra)
	mockOp := contracts.UserOperation{}

	// 4. Executa a função
	hash, err := svc.SendBundle(mockOp)

	// 5. Valida os resultados
	if err != nil {
		t.Fatalf("Não esperava erro, mas obteve: %v", err)
	}

	if hash == "" {
		t.Errorf("Esperava um hash de transação, obteve string vazia")
	}

	t.Logf("Teste passou! Hash simulado: %s", hash)
}

// --- MOCK DE FALHA ---

// Simulando o Contrato EntryPoint retornando um erro
type MockEntryPointError struct{}

func (m *MockEntryPointError) HandleOps(opts *bind.TransactOpts, ops []contracts.UserOperation, beneficiary common.Address) (*types.Transaction, error) {
	// Retornamos nil para a transação e um erro simulado do contrato
	return nil, fmt.Errorf("execution reverted: AA21 didn't pay prefund")
}

func TestSendBundle_Failure_ContractRevert(t *testing.T) {
	// 1. Prepara a configuração com uma chave válida para não falhar na etapa de criptografia
	cfg := &config.Config{
		PrivateKey: "fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19", 
	}

	// 2. Injeta o Mock de SUCESSO para o Client, mas o Mock de ERRO para o EntryPoint
	svc := NewRelayerService(cfg, &MockChainClient{}, &MockEntryPointError{}, nil)

	// 3. Cria a UserOperation
	mockOp := contracts.UserOperation{}

	// 4. Executa a função
	hash, err := svc.SendBundle(mockOp)

	// 5. Validações (Asserts) rigorosas para garantir que o erro foi tratado

	// O erro NÃO pode ser nulo
	if err == nil {
		t.Fatalf("Esperava que a função retornasse um erro, mas retornou nil")
	}

	// A hash da transação DEVE ser vazia, pois ela nunca foi enviada
	if hash != "" {
		t.Errorf("Esperava uma hash vazia em caso de falha, mas obteve: %s", hash)
	}

	// Verifica se a mensagem de erro foi encapsulada corretamente pelo nosso serviço
	expectedErrorSubstring := "falha ao enviar bundle: execution reverted: AA21 didn't pay prefund"
	if err.Error() != expectedErrorSubstring {
		t.Errorf("Mensagem de erro inesperada.\nEsperava: %s\nObteve: %v", expectedErrorSubstring, err.Error())
	}

	t.Logf("Teste de falha passou! Erro capturado com sucesso: %v", err)
}

// --- TESTES DE IMAGEM ---

func TestGenerateCertificateImage(t *testing.T) {
	// 1. Setup
	svc := &RelayerService{}

	// Preenche uma requisição completa com dados de teste
	req := CertificateRequest{
		StudentName:        "GABRIEL HENRIQUE BRIOTO", // Caps para bold ficar mais visível
		CourseName:         "ENGENHARIA DE COMPUTAÇÃO WEB3", // Caps para bold
		Hours:              "40",
		StartDate:          "17/03/2026",
		EndDate:            "17/04/2026",
		GradePercent:       "95",
		DirectorName:       "Prof. Dr. Fulano de Tal",
		DirectorTitle:      "Diretor da Poli-USP",
		CoordinatorName:    "Ciclana de Souza",
		CoordinatorTitle:   "Coordenadora do Curso",
		EmissionDay:        "17",
		EmissionMonth:      "03",
		EmissionYear:       "2026",
		VerificationCode:   "USP-12345678",
	}

	// 2. Execução
	imgBytes, err := svc.GenerateCertificateImage(req)

	// 3. Asserções (Verificações)
	if err != nil {
		t.Fatalf("Esperava erro nil, mas obteve: %v\nVerifique se os arquivos de template e fonte existem na pasta assets/.", err)
	}

	if len(imgBytes) == 0 {
		t.Fatalf("Esperava um slice de bytes preenchido, mas obteve 0 bytes")
	}

	// 4. Validação de formato
	// Tenta decodificar o slice de bytes para garantir que é um JPEG estruturalmente válido
	_, err = jpeg.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		t.Errorf("Os bytes retornados não formam um JPEG válido: %v", err)
	}

	// 5. Salvar a imagem no disco para visualização física (opcional)
	err = os.WriteFile("resultado_teste_certificado_ajustado.jpg", imgBytes, 0644)
	if err != nil {
		t.Fatalf("Falha ao salvar a imagem de teste no disco: %v", err)
	}

	t.Log("Sucesso! Imagem salva como 'resultado_teste_certificado_ajustado.jpg'")
}