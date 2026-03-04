package service

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gabrielhbrioto/usp-project/relayer/internal/config"
	"github.com/gabrielhbrioto/usp-project/relayer/pkg/contracts"
)

type RelayerService struct {
	client     *ethclient.Client
	cfg        *config.Config
	entryPoint *contracts.EntryPoint
}

func NewRelayerService(cfg *config.Config) (*RelayerService, error) {
	client, err := ethclient.Dial(cfg.RpcUrl)
	if err != nil {
		return nil, err
	}

	address := common.HexToAddress(cfg.EntryPointAddr)
	ep, err := contracts.NewEntryPoint(address, client)
	if err != nil {
		return nil, err
	}

	return &RelayerService{client: client, cfg: cfg, entryPoint: ep}, nil
}

// SendBundle recebe a UserOperation (struct simplificada aqui) e envia
func (s *RelayerService) SendBundle(userOp contracts.UserOperation) (string, error) {
	// 1. Carregar chave privada do Bundler
	privateKey, _ := crypto.HexToECDSA(s.cfg.PrivateKey)

	// 2. Criar TransactOpts (autenticação do bundler)
	chainId, _ := s.client.ChainID(context.Background())
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainId)

	// 3. Empacotar a UserOp (Array de 1 item)
	ops := []contracts.UserOperation{userOp}

	// 4. Chamar handleOps no contrato EntryPoint
	// OBS: O 'beneficiary' é o endereço que recebe as taxas (o próprio bundler)
	bundlerAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	tx, err := s.entryPoint.HandleOps(auth, ops, bundlerAddress)
	if err != nil {
		return "", fmt.Errorf("falha ao enviar bundle: %v", err)
	}

	return tx.Hash().Hex(), nil
}
