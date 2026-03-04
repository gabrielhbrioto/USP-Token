package service

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gabrielhbrioto/Usp-Token/packages/relayer-bundler/internal/config"
	"github.com/gabrielhbrioto/Usp-Token/packages/relayer-bundler/pkg/contracts"
)

// 1. Interfaces que representam as funções que precisamos da blockchain
type ChainClient interface {
	ChainID(ctx context.Context) (*big.Int, error)
}

type EntryPointContract interface {
	HandleOps(opts *bind.TransactOpts, ops []contracts.UserOperation, beneficiary common.Address) (*types.Transaction, error)
}

// 2. A struct agora depende das interfaces, não das implementações concretas
type RelayerService struct {
	client     ChainClient
	cfg        *config.Config
	entryPoint EntryPointContract
}

// 3. O construtor recebe as dependências prontas
func NewRelayerService(cfg *config.Config, client ChainClient, ep EntryPointContract) *RelayerService {
	return &RelayerService{
		client:     client,
		cfg:        cfg,
		entryPoint: ep,
	}
}

func (s *RelayerService) SendBundle(userOp contracts.UserOperation) (string, error) {
	privateKey, err := crypto.HexToECDSA(s.cfg.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("chave privada invalida: %v", err)
	}

	chainId, err := s.client.ChainID(context.Background())
	if err != nil {
		return "", fmt.Errorf("falha ao obter chainID: %v", err)
	}

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	ops := []contracts.UserOperation{userOp}
	bundlerAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	tx, err := s.entryPoint.HandleOps(auth, ops, bundlerAddress)
	if err != nil {
		return "", fmt.Errorf("falha ao enviar bundle: %v", err)
	}

	return tx.Hash().Hex(), nil
}