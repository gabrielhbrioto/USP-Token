package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Setup: Definimos variáveis de ambiente simuladas para o teste
	os.Setenv("SEPOLIA_RPC_URL", "https://fake-rpc.com")
	os.Setenv("BUNDLER_PRIVATE_KEY", "1234567890abcdef")
	os.Setenv("ENTRYPOINT_ADDRESS", "0xFakeAddress")
	os.Setenv("PORT", "9090")

	// Clean up: Limpamos as variáveis após o teste para não afetar outros testes
	defer func() {
		os.Unsetenv("SEPOLIA_RPC_URL")
		os.Unsetenv("BUNDLER_PRIVATE_KEY")
		os.Unsetenv("ENTRYPOINT_ADDRESS")
		os.Unsetenv("PORT")
	}()

	// Execução: Chamamos a função Load
	cfg, err := Load()

	// Validação (Asserts)
	if err != nil {
		t.Fatalf("Esperava erro nil, obteve %v", err)
	}

	if cfg.RpcUrl != "https://fake-rpc.com" {
		t.Errorf("Esperava RPC URL 'https://fake-rpc.com', obteve '%s'", cfg.RpcUrl)
	}

	if cfg.Port != "9090" {
		t.Errorf("Esperava Port '9090', obteve '%s'", cfg.Port)
	}
}

func TestGetEnvFallback(t *testing.T) {
	os.Unsetenv("PORT") // Garantimos que a variável não existe

	// Execução: Testando a função getEnv com fallback
	port := getEnv("PORT", "8080")

	if port != "8080" {
		t.Errorf("Esperava fallback '8080', obteve '%s'", port)
	}
}
