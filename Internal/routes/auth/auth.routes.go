package authroutes

import (
	"github.com/AdelGann/z0-backend-v1/Internal/controllers/AuthControllers"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	api := app.Group("/api/auth")

	api.Post("/login", authcontrollers.Login)
	api.Post("/register", func(c *fiber.Ctx) error {
		return c.SendString("/register user")
	})
}
