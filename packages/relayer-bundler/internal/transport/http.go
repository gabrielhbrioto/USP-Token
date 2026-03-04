package transport

import (
	"net/http"

	"github.com/gabrielhbrioto/usp-project/relayer/internal/service"
	"github.com/gabrielhbrioto/usp-project/relayer/pkg/contracts"
	"github.com/gin-gonic/gin"
)

func SetupRouter(svc *service.RelayerService) *gin.Engine {
	r := gin.Default()

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

	return r
}
