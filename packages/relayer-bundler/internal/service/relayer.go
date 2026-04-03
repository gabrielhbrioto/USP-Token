package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"image/jpeg"
	"log"
	"math/big"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fogleman/gg"
	"github.com/gabrielhbrioto/Usp-Token/packages/relayer-bundler/internal/config"
	"github.com/gabrielhbrioto/Usp-Token/packages/relayer-bundler/pkg/contracts"
	"google.golang.org/api/idtoken"
)

// --- Structs do Padrão NFT e Pinata ---

type PinataResponse struct {
	IpfsHash string `json:"IpfsHash"`
}

type NFTAttribute struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

type NFTMetadata struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Image       string         `json:"image"`
	Attributes  []NFTAttribute `json:"attributes"`
}

type CertificateRequest struct {
	StudentAddress   string `json:"studentAddress"`
	StudentName      string `json:"studentName"`
	CourseName       string `json:"courseName"`
	Hours            string `json:"hours"`
	StartDate        string `json:"startDate"`
	EndDate          string `json:"endDate"`
	GradePercent     string `json:"gradePercent"`
	DirectorName     string `json:"directorName"`
	DirectorTitle    string `json:"directorTitle"`
	CoordinatorName  string `json:"coordinatorName"`
	CoordinatorTitle string `json:"coordinatorTitle"`
	EmissionDay      string `json:"emissionDay"`
	EmissionMonth    string `json:"emissionMonth"`
	EmissionYear     string `json:"emissionYear"`
	VerificationCode string `json:"verificationCode"`
}

type textConfig struct {
	boldPath    string
	regularPath string
	mainColor   string
	dataColor   string
}

// 1. Interfaces que representam as funções que precisamos da blockchain
type ChainClient interface {
	ChainID(ctx context.Context) (*big.Int, error)
}

type EntryPointContract interface {
	HandleOps(opts *bind.TransactOpts, ops []contracts.UserOperation, beneficiary common.Address) (*types.Transaction, error)
}

type IdentityRegistryContract interface {
	AddStudent(opts *bind.TransactOpts, studentAddress common.Address, studentId string) (*types.Transaction, error)
}

type CertificateContract interface {
	SystemMintCertificate(opts *bind.TransactOpts, student common.Address, metadataURI string) (*types.Transaction, error)
}

// 2. A struct agora depende das interfaces, não das implementações concretas
type RelayerService struct {
	client              ChainClient
	cfg                 *config.Config
	entryPoint          EntryPointContract
	identityRegistry    IdentityRegistryContract
	certificateContract CertificateContract
}

// 3. O construtor recebe as dependências prontas
func NewRelayerService(cfg *config.Config, client ChainClient, ep EntryPointContract, idReg IdentityRegistryContract, certContract CertificateContract) *RelayerService {
	return &RelayerService{
		client:              client,
		cfg:                 cfg,
		entryPoint:          ep,
		identityRegistry:    idReg,
		certificateContract: certContract,
	}
}

func bigIntString(v *big.Int) string {
	if v == nil {
		return "<nil>"
	}
	return v.String()
}

func logErrorChain(prefix string, err error) {
	if err == nil {
		return
	}

	log.Printf("%s: %v", prefix, err)
	for depth, cause := 1, errors.Unwrap(err); cause != nil; depth, cause = depth+1, errors.Unwrap(cause) {
		log.Printf("%s cause[%d]: %v", prefix, depth, cause)
	}
}

func (s *RelayerService) SendBundle(userOp contracts.UserOperation) (string, error) {
	privateKey, err := crypto.HexToECDSA(s.cfg.PrivateKey)
	if err != nil {
		wrappedErr := fmt.Errorf("chave privada invalida: %w", err)
		logErrorChain("SendBundle: falha ao carregar chave privada", wrappedErr)
		return "", wrappedErr
	}

	chainId, err := s.client.ChainID(context.Background())
	if err != nil {
		wrappedErr := fmt.Errorf("falha ao obter chainID: %w", err)
		logErrorChain("SendBundle: falha ao obter chainID", wrappedErr)
		return "", wrappedErr
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		wrappedErr := fmt.Errorf("falha ao criar transactor com chainID %s: %w", chainId.String(), err)
		logErrorChain("SendBundle: falha ao criar transactor", wrappedErr)
		return "", wrappedErr
	}

	ops := []contracts.UserOperation{userOp}
	bundlerAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	tx, err := s.entryPoint.HandleOps(auth, ops, bundlerAddress)
	if err != nil {
		wrappedErr := fmt.Errorf("falha ao enviar bundle: %w", err)
		logErrorChain(
			fmt.Sprintf(
				"SendBundle: HandleOps falhou sender=%s nonce=%s chainID=%s callDataLen=%d paymasterDataLen=%d bundler=%s",
				userOp.Sender.Hex(),
				bigIntString(userOp.Nonce),
				chainId.String(),
				len(userOp.CallData),
				len(userOp.PaymasterAndData),
				bundlerAddress.Hex(),
			),
			wrappedErr,
		)
		return "", wrappedErr
	}

	return tx.Hash().Hex(), nil
}

func (s *RelayerService) RegisterStudent(ctx context.Context, tokenString string, studentAddressHex string) (string, error) {
	// Valida o Token JWT com os servidores do Google
	payload, err := idtoken.Validate(ctx, tokenString, s.cfg.GoogleClientID)
	if err != nil {
		wrappedErr := fmt.Errorf("token do google invalido: %w", err)
		logErrorChain("RegisterStudent: validacao de token falhou", wrappedErr)
		return "", wrappedErr
	}

	// Extrai o email do payload
	emailItem, ok := payload.Claims["email"]
	if !ok {
		return "", fmt.Errorf("email nao encontrado no token")
	}
	email := emailItem.(string)

	// Verificação de domínio
	// if !strings.HasSuffix(email, "@usp.br") {
	if !strings.HasSuffix(email, "@gmail.com") {
		return "", fmt.Errorf("apenas emails da USP sao permitidos")
	}

	// Prepara a transação para a Blockchain
	privateKey, err := crypto.HexToECDSA(s.cfg.PrivateKey)
	if err != nil {
		wrappedErr := fmt.Errorf("chave privada invalida: %w", err)
		logErrorChain("RegisterStudent: falha ao carregar chave privada", wrappedErr)
		return "", wrappedErr
	}

	chainId, err := s.client.ChainID(ctx)
	if err != nil {
		wrappedErr := fmt.Errorf("falha ao obter chainID: %w", err)
		logErrorChain("RegisterStudent: falha ao obter chainID", wrappedErr)
		return "", wrappedErr
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		wrappedErr := fmt.Errorf("falha ao criar transactor com chainID %s: %w", chainId.String(), err)
		logErrorChain("RegisterStudent: falha ao criar transactor", wrappedErr)
		return "", wrappedErr
	}

	// Como o contrato já foi injetado, basta usá-lo!
	studentAddress := common.HexToAddress(studentAddressHex)

	tx, err := s.identityRegistry.AddStudent(auth, studentAddress, email)
	if err != nil {
		wrappedErr := fmt.Errorf("falha ao registrar aluno na blockchain: %w", err)
		logErrorChain(
			fmt.Sprintf("RegisterStudent: AddStudent falhou student=%s email=%s chainID=%s", studentAddress.Hex(), email, chainId.String()),
			wrappedErr,
		)
		return "", wrappedErr
	}

	return tx.Hash().Hex(), nil
}

func getAssetPath(filename string) string {
	pathRoot := "assets/" + filename
	if _, err := os.Stat(pathRoot); err == nil {
		return pathRoot
	}
	pathTest := "../../assets/" + filename
	if _, err := os.Stat(pathTest); err == nil {
		return pathTest
	}
	return filename // fallback
}

func (s *RelayerService) GenerateCertificateImage(req CertificateRequest) ([]byte, error) {
	txtCfg := textConfig{
		boldPath: getAssetPath("Roboto-Bold.ttf"),
		// regularPath: getAssetPath("Roboto-Regular.ttf"),
		mainColor: "#2C3E50",
		dataColor: "#333333",
	}

	img, err := gg.LoadImage(getAssetPath("template-certificado.png")) // Certifique-se de que este é o template LIMPO
	if err != nil {
		log.Printf("ERRO INTERNO [GenerateCertificate]: %v\n", err)
		return nil, fmt.Errorf("falha ao carregar o template: %v", err)
	}

	bounds := img.Bounds()
	dc := gg.NewContext(bounds.Dx(), bounds.Dy())
	dc.DrawImage(img, 0, 0)

	// Transformamos em float para usar cálculo percentual
	W := float64(bounds.Dx())
	H := float64(bounds.Dy())

	// Tamanhos de fonte dinâmicos baseados na altura da imagem
	titleFontSize := H * 0.045
	courseFontSize := H * 0.035
	normalFontSize := H * 0.025
	signatureFontSize := H * 0.030
	signatureTitleFontSize := H * 0.020

	// --- 1. Títulos Principais (Bold) ---
	dc.SetHexColor(txtCfg.mainColor)

	// Nome do Aluno
	dc.LoadFontFace(txtCfg.boldPath, titleFontSize)
	dc.DrawStringAnchored(req.StudentName, W*0.50, H*0.275, 0.5, 0.5)

	// Nome do Curso
	dc.LoadFontFace(txtCfg.boldPath, courseFontSize)
	dc.DrawStringAnchored(req.CourseName, W*0.50, H*0.385, 0.5, 0.5)

	// --- 2. Dados Detalhados (Regular) ---
	dc.SetHexColor(txtCfg.dataColor)
	dc.LoadFontFace(txtCfg.regularPath, normalFontSize)

	// Linha de Caixas (Carga Horária, Período, Nível)
	yRow1 := H * 0.505
	dc.DrawStringAnchored(req.Hours, W*0.215, yRow1, 0.5, 0.5)
	dc.DrawStringAnchored(req.StartDate+" a "+req.EndDate, W*0.50, yRow1, 0.5, 0.5)
	dc.DrawStringAnchored(req.GradePercent, W*0.785, yRow1, 0.5, 0.5)

	// Data de Emissão
	yEmission := H * 0.675
	dc.DrawStringAnchored(req.EmissionDay, W*0.46, yEmission, 0.5, 0.5)
	dc.DrawStringAnchored(req.EmissionMonth, W*0.54, yEmission, 0.5, 0.5)
	dc.DrawStringAnchored(req.EmissionYear, W*0.63, yEmission, 0.5, 0.5)

	// Código de Verificação
	yVerify := H * 0.895
	dc.DrawStringAnchored(req.VerificationCode, W*0.60, yVerify, 0.5, 0.5)

	// --- 3. Assinaturas ---
	ySignName := H * 0.78 // Logo acima da linha
	// ySignTitle := H * 0.87 // Abaixo do texto estático "Diretor/Coordenador"

	// Nomes (Bold)
	dc.SetHexColor(txtCfg.mainColor)
	dc.LoadFontFace(txtCfg.boldPath, signatureFontSize)
	dc.DrawStringAnchored(req.DirectorName, W*0.25, ySignName, 0.5, 0.5)
	dc.DrawStringAnchored(req.CoordinatorName, W*0.75, ySignName, 0.5, 0.5)

	// Cargos (Regular)
	dc.SetHexColor(txtCfg.dataColor)
	dc.LoadFontFace(txtCfg.regularPath, signatureTitleFontSize)
	// dc.DrawStringAnchored(req.DirectorTitle, W*0.25, ySignTitle, 0.5, 0.5)
	// dc.DrawStringAnchored(req.CoordinatorTitle, W*0.75, ySignTitle, 0.5, 0.5)

	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, dc.Image(), &jpeg.Options{Quality: 95}); err != nil {
		return nil, fmt.Errorf("falha ao encodar a imagem gerada: %v", err)
	}

	return buf.Bytes(), nil
}

// --- Funções de Upload para o IPFS ---

// uploadFileToPinata envia o buffer JPEG gerado para o IPFS
func (s *RelayerService) uploadFileToPinata(fileBytes []byte, filename string) (string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return "", err
	}
	part.Write(fileBytes)
	writer.Close()

	req, err := http.NewRequest("POST", "https://api.pinata.cloud/pinning/pinFileToIPFS", body)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+s.cfg.PinataJWT)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("erro do pinata file: %s", resp.Status)
	}

	var pinataResp PinataResponse
	json.NewDecoder(resp.Body).Decode(&pinataResp)

	return "ipfs://" + pinataResp.IpfsHash, nil
}

// uploadJSONToPinata envia os metadados do certificado para o IPFS
func (s *RelayerService) uploadJSONToPinata(metadata NFTMetadata) (string, error) {
	jsonData, err := json.Marshal(metadata)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.pinata.cloud/pinning/pinJSONToIPFS", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+s.cfg.PinataJWT)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("erro do pinata json: %s", resp.Status)
	}

	var pinataResp PinataResponse
	json.NewDecoder(resp.Body).Decode(&pinataResp)

	return "ipfs://" + pinataResp.IpfsHash, nil
}

// --- Orquestrador Final ---

// GenerateAndPinCertificate gera a imagem, sobe pro IPFS, cria o JSON, sobe pro IPFS e retorna o URI final
func (s *RelayerService) GenerateAndPinCertificate(req CertificateRequest) (string, error) {
	// 1. Gera a imagem em memória (função que já criamos e validamos)
	imgBytes, err := s.GenerateCertificateImage(req)
	if err != nil {
		return "", fmt.Errorf("falha ao gerar imagem: %v", err)
	}

	// 2. Faz o upload da imagem para o IPFS
	imageURI, err := s.uploadFileToPinata(imgBytes, "certificado_"+strings.ReplaceAll(req.StudentName, " ", "_")+".jpg")
	if err != nil {
		return "", fmt.Errorf("falha no upload da imagem: %v", err)
	}

	// 3. Monta o JSON padrão ERC-721
	metadata := NFTMetadata{
		Name:        "Certificado USP - " + req.StudentName,
		Description: "Certificado Soulbound de conclusão do curso de " + req.CourseName + " no ecossistema USP.",
		Image:       imageURI, // <-- Aqui vai o link ipfs://Qm... da imagem gerada!
		Attributes: []NFTAttribute{
			{TraitType: "Curso", Value: req.CourseName},
			{TraitType: "Carga Horária", Value: req.Hours + "h"},
			{TraitType: "Nota", Value: req.GradePercent},
			{TraitType: "Data de Emissão", Value: fmt.Sprintf("%s/%s/%s", req.EmissionDay, req.EmissionMonth, req.EmissionYear)},
		},
	}

	// 4. Faz o upload do JSON para o IPFS
	metadataURI, err := s.uploadJSONToPinata(metadata)
	if err != nil {
		return "", fmt.Errorf("falha no upload dos metadados: %v", err)
	}

	// 5. Retorna o metadataURI final!
	return metadataURI, nil
}

// IssueCertificateOnChain assina e envia a transação de Mint do NFT
func (s *RelayerService) IssueCertificateOnChain(ctx context.Context, studentAddressHex string, metadataURI string) (string, error) {
	privateKey, err := crypto.HexToECDSA(s.cfg.PrivateKey)
	if err != nil {
		wrappedErr := fmt.Errorf("chave privada invalida: %w", err)
		logErrorChain("IssueCertificateOnChain: falha ao carregar chave privada", wrappedErr)
		return "", wrappedErr
	}

	chainId, err := s.client.ChainID(ctx)
	if err != nil {
		wrappedErr := fmt.Errorf("falha ao obter chainID: %w", err)
		logErrorChain("IssueCertificateOnChain: falha ao obter chainID", wrappedErr)
		return "", wrappedErr
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		wrappedErr := fmt.Errorf("falha ao criar transactor com chainID %s: %w", chainId.String(), err)
		logErrorChain("IssueCertificateOnChain: falha ao criar transactor", wrappedErr)
		return "", wrappedErr
	}

	studentAddress := common.HexToAddress(studentAddressHex)

	// AQUI: Assumindo que você adicionou o certificado no RelayerService (como fez com IdentityRegistry)
	// Chamamos a nova função systemMintCertificate do Smart Contract
	tx, err := s.certificateContract.SystemMintCertificate(auth, studentAddress, metadataURI)
	if err != nil {
		wrappedErr := fmt.Errorf("falha ao emitir NFT na blockchain: %w", err)
		logErrorChain(
			fmt.Sprintf("IssueCertificateOnChain: mint falhou student=%s chainID=%s metadataURI=%s", studentAddress.Hex(), chainId.String(), metadataURI),
			wrappedErr,
		)
		return "", wrappedErr
	}

	return tx.Hash().Hex(), nil
}
