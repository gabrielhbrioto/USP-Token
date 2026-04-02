package transport

import (
	"net/http"
	"log"
	"os"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gabrielhbrioto/Usp-Token/packages/relayer-bundler/internal/service"
	"github.com/gabrielhbrioto/Usp-Token/packages/relayer-bundler/pkg/contracts"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// DTO para receber os dados puros do Frontend em texto e número
type UserOperationDTO struct {
	Sender               string `json:"sender"`
	Nonce                int64  `json:"nonce"`
	InitCode             string `json:"initCode"`
	CallData             string `json:"callData"`
	CallGasLimit         int64  `json:"callGasLimit"`
	VerificationGasLimit int64  `json:"verificationGasLimit"`
	PreVerificationGas   int64  `json:"preVerificationGas"`
	MaxFeePerGas         int64  `json:"maxFeePerGas"`
	MaxPriorityFeePerGas int64  `json:"maxPriorityFeePerGas"`
	PaymasterAndData     string `json:"paymasterAndData"`
	Signature            string `json:"signature"`
}

type RegisterRequest struct {
	IDToken        string `json:"idToken"`
	StudentAddress string `json:"studentAddress"`
}

func SetupRouter(svc *service.RelayerService) *gin.Engine {
	r := gin.Default()

	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:8081"
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{frontendURL} 
	corsConfig.AllowMethods = []string{"POST", "GET", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	
	r.Use(cors.New(corsConfig))

	r.POST("/relay", func(c *gin.Context) {
		var dto UserOperationDTO

		// 1. Recebe os dados no formato seguro (String/Int)
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "JSON invalido: " + err.Error()})
			return
		}

		// 2. Converte as strings hexadecimais (0x...) para array de bytes ([]byte)
		initCode, _ := hexutil.Decode(dto.InitCode)
		callData, _ := hexutil.Decode(dto.CallData)
		paymasterAndData, _ := hexutil.Decode(dto.PaymasterAndData)
		signature, _ := hexutil.Decode(dto.Signature)

		// 3. Monta a struct oficial gerada pelo Abigen
		userOp := contracts.UserOperation{
			Sender:               common.HexToAddress(dto.Sender),
			Nonce:                big.NewInt(dto.Nonce),
			InitCode:             initCode,
			CallData:             callData,
			CallGasLimit:         big.NewInt(dto.CallGasLimit),
			VerificationGasLimit: big.NewInt(dto.VerificationGasLimit),
			PreVerificationGas:   big.NewInt(dto.PreVerificationGas),
			MaxFeePerGas:         big.NewInt(dto.MaxFeePerGas),
			MaxPriorityFeePerGas: big.NewInt(dto.MaxPriorityFeePerGas),
			PaymasterAndData:     paymasterAndData,
			Signature:            signature,
		}

		txHash, err := svc.SendBundle(userOp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"txHash": txHash})
	})

	r.POST("/register", func(c *gin.Context) {
		var req RegisterRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "JSON invalido"})
			return
		}

		// Chama a nova função do serviço, passando o Context do Gin
		txHash, err := svc.RegisterStudent(c.Request.Context(), req.IDToken, req.StudentAddress)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Aluno matriculado com sucesso",
			"txHash":  txHash,
		})
	})

	// Nova rota para gerar o certificado em memória
	r.POST("/certificate/generate", func(c *gin.Context) {
		var req service.CertificateRequest

		// Bind JSON para a struct do serviço
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "JSON invalido"})
			return
		}

		// Chama o serviço para gerar os bytes da imagem
		imgBytes, err := svc.GenerateCertificateImage(req)
		if err != nil {
			log.Printf("ERRO INTERNO [GenerateCertificate]: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// No Gin, retornamos bytes brutos (como imagens) usando c.Data
		c.Data(http.StatusOK, "image/jpeg", imgBytes)
	})

	// Rota que prepara o certificado (Gera Imagem -> IPFS -> Retorna URI)
	r.POST("/certificate/prepare", func(c *gin.Context) {
		var req service.CertificateRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "JSON invalido"})
			return
		}

		// Chama a nova função orquestradora!
		metadataURI, err := svc.GenerateAndPinCertificate(req)
		if err != nil {
			log.Printf("ERRO INTERNO [GenerateAndPinCertificate]: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Devolve a URL final para o frontend React
		c.JSON(http.StatusOK, gin.H{
			"metadataURI": metadataURI,
		})
	})

	// Nova rota completa: Emite e Registra na Blockchain
	r.POST("/certificate/issue", func(c *gin.Context) {
		var req service.CertificateRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "JSON invalido"})
			return
		}

		// 1. Gera imagem e sobe pro IPFS
		metadataURI, err := svc.GenerateAndPinCertificate(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha no IPFS: " + err.Error()})
			return
		}

		// 2. Envia a transação na blockchain (O Relayer paga o gás!)
		txHash, err := svc.IssueCertificateOnChain(c.Request.Context(), req.StudentAddress, metadataURI)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha na Blockchain: " + err.Error()})
			return
		}

		// 3. Sucesso total!
		c.JSON(http.StatusOK, gin.H{
			"message": "Certificado NFT emitido com sucesso!",
			"metadataURI": metadataURI,
			"txHash": txHash,
		})
	})

	return r
}
