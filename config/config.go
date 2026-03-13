package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseURL string
	Env         string
}

func Load() *Config {
	env := getEnv("GO_ENV", "development")

	// Attempt to load .env file based on environment
	envFile := ".env." + env
	if err := godotenv.Load(envFile); err != nil {
		log.Printf("Warning: No %s file found, using system environment variables", envFile)
	}

	return &Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"),
		Env:         env,
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
