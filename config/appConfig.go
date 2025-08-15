package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	envPort      = "PORT"
	envDSN       = "DSN"
	envAPPSECRET = "APPSECRET"
)

type AppConfig struct {
	ServerPort string
	Dsn        string
	AppSecret  string
}

func getEnvOrFail(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("%s could not be found in environment variables", key)
	}
	return val
}

func SetupEnv() (cfg AppConfig) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return AppConfig{
		ServerPort: getEnvOrFail(envPort),
		Dsn:        getEnvOrFail(envDSN),
		AppSecret:  getEnvOrFail(envAPPSECRET),
	}
}
