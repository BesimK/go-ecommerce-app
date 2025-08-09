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
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("database connection error %v\n", err)
	}

	log.Println("database connected!")
	db.AutoMigrate(&domain.User{})
	rh := &rest.RestHandler{
		App: app,
		DB:  db,
	}

	setupRoutes(rh)
	app.Listen(config.ServerPort)
}

func setupRoutes(rh *rest.RestHandler) {
	handlers.SetupUserRoutes(rh)
}
