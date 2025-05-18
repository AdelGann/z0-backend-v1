package main

import (
	"github.com/AdelGann/z0-backend-v1/config"
	"github.com/AdelGann/z0-backend-v1/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// connection to the database
	config.ConnectDB()

	// initialize app
	app := fiber.New()

	app.Use(recover.New())

	// user routes
	routes.MainRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
