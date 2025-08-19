package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/BesimK/go-ecommerce-app/internal/api/rest"
	"github.com/BesimK/go-ecommerce-app/internal/repository"
	"github.com/BesimK/go-ecommerce-app/internal/service"
)

type CatalogHandler struct {
	svc service.CatalogService
}

func SetupCatalogRoutes(rh *rest.RestHandler) {
	app := rh.App

	svc := service.CatalogService{
		Repo:   repository.NewCatalogRepository(rh.DB),
		Auth:   rh.Auth,
		Config: rh.Config, // düzeltildi (Confing → Config)
	}

	handler := CatalogHandler{
		svc: svc,
	}

	// Public routes (no authentication)
	app.Get("/products", handler.GetProducts)
	app.Get("/products/:id", handler.GetProduct)
	app.Get("/categories", handler.GetCategories)
	app.Get("/categories/:id", handler.GetCategoryById)

	// Private routes (seller authentication)
	selRoutes := app.Group("/seller", rh.Auth.AuthorizeSeller)

	// Categories
	selRoutes.Post("/categories", handler.CreateCategories)
	selRoutes.Patch("/categories/:id", handler.EditCategory)
	selRoutes.Delete("/categories/:id", handler.DeleteCategory)

	// Products
	selRoutes.Post("/products", handler.CreateProducts)
	selRoutes.Get("/products", handler.GetSellerProducts)
	selRoutes.Get("/products/:id", handler.GetProduct)
	selRoutes.Put("/products/:id", handler.EditProduct)
	selRoutes.Patch("/products/:id/stock", handler.UpdateStock)
	selRoutes.Delete("/products/:id", handler.DeleteProduct)
}

// ---------- Categories ----------

func (h CatalogHandler) CreateCategories(ctx *fiber.Ctx) error {
	// Placeholder: parse input and create category
	return rest.SuccessMessage(ctx, "Category created successfully", nil)
}

func (h CatalogHandler) EditCategory(ctx *fiber.Ctx) error {
	// Placeholder: update category details
	return rest.SuccessMessage(ctx, "Category updated successfully", nil)
}

func (h CatalogHandler) DeleteCategory(ctx *fiber.Ctx) error {
	// Placeholder: delete category by ID
	return rest.SuccessMessage(ctx, "Category deleted successfully", nil)
}

func (h CatalogHandler) GetCategories(ctx *fiber.Ctx) error {
	// Placeholder: fetch list of categories
	return rest.SuccessMessage(ctx, "Fetched categories", []string{})
}

func (h CatalogHandler) GetCategoryById(ctx *fiber.Ctx) error {
	// Placeholder: fetch category by ID
	return rest.SuccessMessage(ctx, "Fetched category details", nil)
}

// ---------- Products ----------

func (h CatalogHandler) CreateProducts(ctx *fiber.Ctx) error {
	// Placeholder: create new product
	return rest.SuccessMessage(ctx, "Product created successfully", nil)
}

func (h CatalogHandler) GetProducts(ctx *fiber.Ctx) error {
	// Placeholder: fetch all products
	return rest.SuccessMessage(ctx, "Fetched products", []string{})
}

func (h CatalogHandler) GetSellerProducts(ctx *fiber.Ctx) error {
	// Placeholder: fetch seller's products
	return rest.SuccessMessage(ctx, "Fetched seller products", []string{})
}

func (h CatalogHandler) GetProduct(ctx *fiber.Ctx) error {
	// Placeholder: fetch product by ID
	return rest.SuccessMessage(ctx, "Fetched product details", nil)
}

func (h CatalogHandler) EditProduct(ctx *fiber.Ctx) error {
	// Placeholder: update product details
	return rest.SuccessMessage(ctx, "Product updated successfully", nil)
}

func (h CatalogHandler) UpdateStock(ctx *fiber.Ctx) error {
	// Placeholder: update stock quantity
	return rest.SuccessMessage(ctx, "Stock updated successfully", nil)
}

func (h CatalogHandler) DeleteProduct(ctx *fiber.Ctx) error {
	// Placeholder: delete product by ID
	return rest.SuccessMessage(ctx, "Product deleted successfully", nil)
}
