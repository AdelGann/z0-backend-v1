package orgroutes

import (
	orgcontrollers "github.com/AdelGann/z0-backend-v1/internal/controllers/org_controller"
	"github.com/AdelGann/z0-backend-v1/internal/middlewares/auth_middleware"
	"github.com/gofiber/fiber/v2"
)

func OrgRoutes(app *fiber.App) {
	api := app.Group("/api/v1/organization", authmiddleware.ValidateToken)

	api.Post("/send-invitation", orgcontrollers.Invite)
}
