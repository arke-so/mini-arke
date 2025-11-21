package main

import (
	"log"
	"os"

	"github.com/arke-so/mini-arke/internal/api"
	"github.com/arke-so/mini-arke/internal/database"
	"github.com/arke-so/mini-arke/internal/repository"
	"github.com/arke-so/mini-arke/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	oapimiddleware "github.com/oapi-codegen/echo-middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Load OpenAPI spec for validation
	swagger, err := api.GetSwagger()
	if err != nil {
		log.Fatal("Failed to load OpenAPI spec:", err)
	}

	// Skip server validation (we're using localhost)
	swagger.Servers = nil

	// Add OpenAPI validation middleware - validates all requests automatically
	e.Use(oapimiddleware.OapiRequestValidator(swagger))

	// Setup database
	db, err := database.NewPostgresDBFromURL(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := database.RunMigrations(db); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Setup dependencies
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	handler := api.NewHandler(productService)

	// Register handlers
	api.RegisterHandlers(e, handler)

	port := getEnv("PORT", "8080")
	log.Printf("Server starting on port %s", port)
	e.Logger.Fatal(e.Start(":" + port))
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
