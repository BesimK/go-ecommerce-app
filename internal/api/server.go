package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/BesimK/go-ecommerce-app/config"
	"github.com/BesimK/go-ecommerce-app/internal/api/rest"
	"github.com/BesimK/go-ecommerce-app/internal/api/rest/handlers"
	"github.com/BesimK/go-ecommerce-app/internal/domain"
	"github.com/BesimK/go-ecommerce-app/internal/helper"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("database connection error %v\n", err)
	}

	db.AutoMigrate(&domain.User{})

	auth := helper.SetupAuth(config.AppSecret)

	rh := &rest.RestHandler{
		App:    app,
		DB:     db,
		Auth:   auth,
		Config: config,
	}

	handlers.SetupUserRoutes(rh)

	app.Listen(config.ServerPort)
}
