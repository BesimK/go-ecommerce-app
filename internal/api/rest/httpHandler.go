package rest

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/BesimK/go-ecommerce-app/config"
	"github.com/BesimK/go-ecommerce-app/internal/helper"
)

type RestHandler struct {
	App    *fiber.App
	DB     *gorm.DB
	Auth   helper.Auth
	Config config.AppConfig
}
