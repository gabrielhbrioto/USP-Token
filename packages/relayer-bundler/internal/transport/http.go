package transport

import (
	"net/http"
	"log"

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

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8081"} 
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

	return r
}
