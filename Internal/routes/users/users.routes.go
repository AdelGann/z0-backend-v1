package usersroutes

import (
	"github.com/AdelGann/z0-backend-v1/Internal/controllers/UsersControllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	api := app.Group("/api/v1/users")

	api.Get("/", usercontroller.GetUsers)
	api.Get("/:id", usercontroller.GetUserById)
	api.Post("/create", usercontroller.PostUser)
	api.Patch("/:id", func(c *fiber.Ctx) error {
		return c.SendString("/patch user")
	})
	api.Delete("/id", func(c *fiber.Ctx) error {
		return c.SendString("/delete user")
	})
}
