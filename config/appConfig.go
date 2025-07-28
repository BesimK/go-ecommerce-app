package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
}

func SetupEnv() (cfg AppConfig) {
	godotenv.Load()

	port := os.Getenv("PORT")

	if len(port) < 1 {
		log.Fatalln("config file is not loaded properly")
	}

	return AppConfig{ServerPort: port}
}
