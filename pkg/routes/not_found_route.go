package routes

import (
	"github.com/gofiber/fiber/v2"
)

// NotFoundRoute func for handling a 404 Error response.
 func NotFoundRoute(a *fiber.App) {
	// Register new special route.
	a.Use(
		func(c *fiber.Ctx) error {
			// Return HTTP 404 status with JSON response.
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg": "sorry, endpoint not found",
			})
		},
		)
}
