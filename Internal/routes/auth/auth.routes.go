package authroutes

import (
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	api := app.Group("/api/auth")

	api.Post("/login", func(c *fiber.Ctx) error {
		return c.SendString("/login user")
	})
	api.Post("/register", func(c *fiber.Ctx) error {
		return c.SendString("/register user")
	})
}
