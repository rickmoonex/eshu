package middleware

import (
	"github.com/gofiber/fiber/v2"
	basicMiddleware "github.com/gofiber/fiber/v2/middleware/basicauth"
)

func BasicAuthProtected() func(*fiber.Ctx) error {
	return basicMiddleware.New(basicMiddleware.Config{
		Users: map[string]string{
			"john": "doe",
		},
	})
}
