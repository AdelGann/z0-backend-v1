package routes

import (
	"github.com/AdelGann/z0-backend-v1/internal/routes/auth"
	"github.com/AdelGann/z0-backend-v1/internal/routes/users"
	"github.com/gofiber/fiber/v2"
)

func MainRoutes(c *fiber.App) {
	usersroutes.UserRoutes(c)
	authroutes.AuthRoutes(c)
}
