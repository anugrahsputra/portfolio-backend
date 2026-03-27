package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseURL string
	Env         string
	GmailUser   string
	GmailPass   string
}

func Load() *Config {
	env := getEnv("ENV", "development")

	_ = godotenv.Load(".env")

	return &Config{
		Port:        getEnv("PORT", "8082"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"),
		Env:         env,
		GmailUser:   getEnv("GMAIL_USER", ""),
		GmailPass:   getEnv("GMAIL_PASS", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
