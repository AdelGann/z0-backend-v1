package main

import (
	"github.com/AdelGann/z0-backend-v1/config"
	"github.com/AdelGann/z0-backend-v1/internal/routes"
	"github.com/gofiber/fiber/v2"
	fiberlog "github.com/gofiber/fiber/v2/log"
	stdlog "log"
)

func main() {
	// connection to the database
	config.ConnectDB()

	// initialize app
	app := fiber.New()

	// middleware to prevent panic errors
	app.Use(func(c *fiber.Ctx) (err error) {
		defer func() {
			if r := recover(); r != nil {
				stdlog.Printf("Recovered from panic: %v", r)
				err = c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": r,
				})
			}
		}()
		return c.Next()
	})

	// user routes
	routes.MainRoutes(app)

	fiberlog.Fatal(app.Listen(":3000"))
}
