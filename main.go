package main

import (
	"fmt"

	"github.com/BesimK/go-ecommerce-app/config"
	"github.com/BesimK/go-ecommerce-app/internal/api"
)

func main() {
	fmt.Println("Fat Juicy & Wet ⚥")

	cfg := config.SetupEnv()

	api.StartServer(cfg)
}
