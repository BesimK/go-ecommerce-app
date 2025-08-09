package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
	Dsn        string
}

func SetupEnv() (cfg AppConfig) {
	godotenv.Load()
	
	var	name:string = "Bes"
	port := os.Getenv("PORT")

	if len(port) < 1 {
		log.Fatalln("config file is not loaded properly")
	}

	Dsn := os.Getenv("DSN")

	if len(Dsn) < 1 {
		log.Fatalln("config file is not loaded properly")
	}

	return AppConfig{ServerPort: port, Dsn: Dsn}
}
