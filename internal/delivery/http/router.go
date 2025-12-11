package http

import (
	"github.com/example/clean-arch-template/internal/delivery/http/handler"
	"github.com/example/clean-arch-template/internal/delivery/http/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// SetupRouter configures all routes and middlewares
func SetupRouter(
	userHandler *handler.UserHandler,
	productHandler *handler.ProductHandler,
	orderHandler *handler.OrderHandler,
) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		},
	})

	// Middlewares
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(middleware.ErrorHandler())

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "healthy",
		})
	})

	// API v1 routes
	api := app.Group("/api/v1")

	// User routes
	users := api.Group("/users")
	users.Post("/register", userHandler.Register)
	users.Post("/login", userHandler.Login)
	users.Get("/:id", userHandler.GetProfile)

	// Product routes
	products := api.Group("/products")
	products.Post("/", productHandler.CreateProduct)
	products.Get("/:id", productHandler.GetProduct)
	products.Get("/", productHandler.ListProducts)
	products.Put("/:id", productHandler.UpdateProduct)
	products.Delete("/:id", productHandler.DeleteProduct)

	// Order routes
	orders := api.Group("/orders")
	orders.Post("/", orderHandler.CreateOrder)
	orders.Get("/:id", orderHandler.GetOrderDetail)
	orders.Get("/user/:user_id", orderHandler.ListUserOrders)

	return app
}
