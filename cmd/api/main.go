package main

import (
	"log"

	"github.com/example/clean-arch-template/config"
	"github.com/example/clean-arch-template/internal/delivery/http"
	"github.com/example/clean-arch-template/internal/delivery/http/handler"
	"github.com/example/clean-arch-template/internal/infrastructure/database"
	"github.com/example/clean-arch-template/internal/infrastructure/persistence"
	"github.com/example/clean-arch-template/internal/usecase"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file (optional)
	_ = godotenv.Load()

	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database connection
	db, err := database.NewPostgresConnection(&cfg.Database)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Run auto migration
	if err := database.AutoMigrate(db); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize Repositories
	// Repositories are initialized with normal DB connection
	// They will be re-created with transaction (tx) when needed in use cases
	userRepo := persistence.NewUserRepository(db)
	productRepo := persistence.NewProductRepository(db)

	// Initialize Use Cases
	userUseCase := usecase.NewUserUseCase(userRepo)
	productUseCase := usecase.NewProductUseCase(productRepo)

	// OrderUseCase receives the DB instance directly for transaction management
	orderUseCase := usecase.NewOrderUseCase(db)

	// Initialize Handlers
	userHandler := handler.NewUserHandler(userUseCase)
	productHandler := handler.NewProductHandler(productUseCase)
	orderHandler := handler.NewOrderHandler(orderUseCase)

	// Setup Router
	app := http.SetupRouter(userHandler, productHandler, orderHandler)

	// Start server
	log.Printf("Server starting on port %s...", cfg.Server.Port)
	if err := app.Listen(":" + cfg.Server.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
