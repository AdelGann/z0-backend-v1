package usersroutes

import (
	"github.com/AdelGann/z0-backend-v1/internal/controllers/users_controller"
	"github.com/AdelGann/z0-backend-v1/internal/middlewares/auth_middleware"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	api := app.Group("/api/v1/users", authmiddleware.ValidateToken)

	api.Get("/", authmiddleware.RoleMiddleware([]string{"ADMIN"}), usercontroller.GetUsers)
	api.Get("/:id", authmiddleware.RoleMiddleware([]string{"ADMIN"}), usercontroller.GetUserById)
	api.Post("/create", authmiddleware.RoleMiddleware([]string{"ADMIN"}), usercontroller.PostUser)
	api.Patch("/:id", func(c *fiber.Ctx) error {
		return c.SendString("/patch user")
	})
	api.Delete("/:id", func(c *fiber.Ctx) error {
		return c.SendString("/delete user")
	})
}
