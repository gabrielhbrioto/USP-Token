package main

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gabrielhbrioto/Usp-Token/packages/relayer-bundler/internal/config"
	"github.com/gabrielhbrioto/Usp-Token/packages/relayer-bundler/internal/service"
	"github.com/gabrielhbrioto/Usp-Token/packages/relayer-bundler/internal/transport"
	"github.com/gabrielhbrioto/Usp-Token/packages/relayer-bundler/pkg/contracts"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Erro ao carregar config:", err)
	}

	// Instanciamos as conexões reais aqui!
	client, err := ethclient.Dial(cfg.RpcUrl)
	if err != nil {
		log.Fatal("Erro ao conectar no Ethereum:", err)
	}

	address := common.HexToAddress(cfg.EntryPointAddr)
	ep, err := contracts.NewEntryPoint(address, client)
	if err != nil {
		log.Fatal("Erro ao instanciar EntryPoint:", err)
	}

	// --- NOVO: Instanciamos o IdentityRegistry real ---
	registryAddr := common.HexToAddress(cfg.IdentityRegistryAddr)
	idReg, err := contracts.NewIdentityRegistry(registryAddr, client)
	if err != nil {
		log.Fatal("Erro ao instanciar IdentityRegistry:", err)
	}

	// Passamos o 'idReg' como o 4º parâmetro
	svc := service.NewRelayerService(cfg, client, ep, idReg)
	r := transport.SetupRouter(svc)

	log.Printf("Relayer rodando na porta %s", cfg.Port)
	r.Run(":" + cfg.Port)
}