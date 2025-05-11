package usersroutes

import (
	"github.com/AdelGann/z0-backend-v1/internal/controllers/users_controller"
	"github.com/AdelGann/z0-backend-v1/internal/middlewares/auth_middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	api := app.Group("/api/v1/users")

	api.Get("/", authmiddleware.ValidateToken, usercontroller.GetUsers)
	api.Get("/:id", authmiddleware.ValidateToken, usercontroller.GetUserById)
	api.Post("/create", authmiddleware.ValidateToken, usercontroller.PostUser)
	api.Patch("/:id", func(c *fiber.Ctx) error {
		return c.SendString("/patch user")
	})
	api.Delete("/id", func(c *fiber.Ctx) error {
		return c.SendString("/delete user")
	})
}
