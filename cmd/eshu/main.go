package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rickmoonex/eshu/pkg/routes"
)

func main() {
	app := fiber.New(fiber.Config{})

	routes.PrivateRoutes(app)
	routes.NotFoundRoute(app)

	app.Listen(":3003")
}
