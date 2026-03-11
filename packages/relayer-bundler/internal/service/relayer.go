package service

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"google.golang.org/api/idtoken"

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

type IdentityRegistryContract interface {
	AddStudent(opts *bind.TransactOpts, studentAddress common.Address, studentId string) (*types.Transaction, error)
}

// 2. A struct agora depende das interfaces, não das implementações concretas
type RelayerService struct {
	client     ChainClient
	cfg        *config.Config
	entryPoint EntryPointContract
	identityRegistry   IdentityRegistryContract
}

// 3. O construtor recebe as dependências prontas
func NewRelayerService(cfg *config.Config, client ChainClient, ep EntryPointContract, idReg IdentityRegistryContract) *RelayerService {
	return &RelayerService{
		client:     client,
		cfg:        cfg,
		entryPoint: ep,
		identityRegistry: idReg,
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

func (s *RelayerService) RegisterStudent(ctx context.Context, tokenString string, studentAddressHex string) (string, error) {
	// Valida o Token JWT com os servidores do Google
	payload, err := idtoken.Validate(ctx, tokenString, s.cfg.GoogleClientID)
	if err != nil {
		return "", fmt.Errorf("token do google invalido: %v", err)
	}

	// Extrai o email do payload
	emailItem, ok := payload.Claims["email"]
	if !ok {
		return "", fmt.Errorf("email nao encontrado no token")
	}
	email := emailItem.(string)

	// Verificação de domínio
	if !strings.HasSuffix(email, "@usp.br") {
		return "", fmt.Errorf("apenas emails da USP sao permitidos")
	}

	// Prepara a transação para a Blockchain
	privateKey, err := crypto.HexToECDSA(s.cfg.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("chave privada invalida: %v", err)
	}

	chainId, err := s.client.ChainID(ctx)
	if err != nil {
		return "", fmt.Errorf("falha ao obter chainID: %v", err)
	}

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	
	// Como o contrato já foi injetado, basta usá-lo!
	studentAddress := common.HexToAddress(studentAddressHex)
	
	tx, err := s.identityRegistry.AddStudent(auth, studentAddress, email)
	if err != nil { 
		return "", fmt.Errorf("falha ao registrar aluno na blockchain: %v", err) 
	}

	return tx.Hash().Hex(), nil
}