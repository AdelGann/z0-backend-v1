package authroutes

import (
	"github.com/AdelGann/z0-backend-v1/internal/controllers/auth_controller"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	api := app.Group("/api/auth")

	api.Post("/login", authcontrollers.Login)
	api.Post("/register", authcontrollers.Register)
}
