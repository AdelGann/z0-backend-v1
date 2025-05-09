package routes

import (
	"github.com/AdelGann/z0-backend-v1/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	api := app.Group("/api/v1/users")

	api.Get("/", controllers.GetUsers)
}
