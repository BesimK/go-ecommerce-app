// Package service
package service

import (
	"github.com/BesimK/go-ecommerce-app/config"
	"github.com/BesimK/go-ecommerce-app/internal/helper"
	"github.com/BesimK/go-ecommerce-app/internal/repository"
)

type CatalogService struct {
	Repo   repository.CatalogRepository
	Auth   helper.Auth
	Config config.AppConfig
}
