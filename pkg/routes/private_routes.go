package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rickmoonex/eshu/app/controllers"
	"github.com/rickmoonex/eshu/pkg/middleware"
)

func PrivateRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Post("/new", middleware.BasicAuthProtected(), controllers.NewMemo)
}
