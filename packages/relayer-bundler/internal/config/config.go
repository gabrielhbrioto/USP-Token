package config

import (
	"os"
	"fmt"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	RpcUrl        string
	PrivateKey    string
	EntryPointAddr string
	Port          string
	GoogleClientID       string
	IdentityRegistryAddr string
}

func Load() (*Config, error) {
	godotenv.Load() // Ignora erro se .env não existir (prod)
	
	requiredVars := []string{
		"SEPOLIA_RPC_URL",
		"BUNDLER_PRIVATE_KEY",
		"ENTRYPOINT_ADDRESS",
		"GOOGLE_CLIENT_ID",
		"IDENTITY_REGISTRY_ADDRESS",
	}

	var missing []string
	for _, req := range requiredVars {
		if os.Getenv(req) == "" {
			missing = append(missing, req)
		}
	}

	if len(missing) > 0 {
		return nil, fmt.Errorf("variáveis de ambiente obrigatórias ausentes: %s", strings.Join(missing, ", "))
	}

	return &Config{
		RpcUrl:        os.Getenv("SEPOLIA_RPC_URL"),
		PrivateKey:    os.Getenv("BUNDLER_PRIVATE_KEY"),
		EntryPointAddr: os.Getenv("ENTRYPOINT_ADDRESS"),
		Port:          getEnv("PORT", "8080"),
		GoogleClientID:       os.Getenv("GOOGLE_CLIENT_ID"),      
		IdentityRegistryAddr: os.Getenv("IDENTITY_REGISTRY_ADDRESS"),
	}, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}