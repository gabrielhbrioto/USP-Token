package main

import (
	"log"

	"github.com/gabrielhbrioto/usp-project/relayer/internal/config"
	"github.com/gabrielhbrioto/usp-project/relayer/internal/service"
	"github.com/gabrielhbrioto/usp-project/relayer/internal/transport"
)

func main() {
	// 1. Config
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Erro ao carregar config:", err)
	}

	// 2. Service
	svc, err := service.NewRelayerService(cfg)
	if err != nil {
		log.Fatal("Erro ao conectar no Ethereum:", err)
	}

	// 3. Server
	r := transport.SetupRouter(svc)

	log.Printf("Relayer rodando na porta %s", cfg.Port)
	r.Run(":" + cfg.Port)
}
