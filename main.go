package main

import (
	"github.com/BesimK/go-ecommerce-app/config"
	"github.com/BesimK/go-ecommerce-app/internal/api"
)

func main() {
	cfg := config.SetupEnv()

	api.StartServer(cfg)
}
