package config

import (
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	RpcUrl        string
	PrivateKey    string
	EntryPointAddr string
	Port          string
}

func Load() (*Config, error) {
	godotenv.Load() // Ignora erro se .env não existir (prod)
	return &Config{
		RpcUrl:        os.Getenv("SEPOLIA_RPC_URL"),
		PrivateKey:    os.Getenv("BUNDLER_PRIVATE_KEY"),
		EntryPointAddr: os.Getenv("ENTRYPOINT_ADDRESS"),
		Port:          getEnv("PORT", "8080"),
	}, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}