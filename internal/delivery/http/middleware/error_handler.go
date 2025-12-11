package middleware

import (
	"log"

	"github.com/example/clean-arch-template/pkg/response"
	"github.com/gofiber/fiber/v2"
)

// ErrorHandler is a middleware to handle panics and errors
func ErrorHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic recovered: %v", r)
				response.InternalError(c, "Internal server error")
			}
		}()

		return c.Next()
	}
}
