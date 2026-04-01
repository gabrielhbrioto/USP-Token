package transport

import (
	"net/http"
	"log"
	"os"
	
	"github.com/gabrielhbrioto/Usp-Token/packages/relayer-bundler/internal/service"
	"github.com/gabrielhbrioto/Usp-Token/packages/relayer-bundler/pkg/contracts"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

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
		var userOp contracts.UserOperation // Struct gerada pelo abigen

		// Bind JSON para Struct (Pode precisar de um DTO intermediário se os tipos não baterem exato)
		if err := c.ShouldBindJSON(&userOp); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
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
