package main

import (
	"github.com/AdelGann/z0-backend-v1/config"
	"github.com/AdelGann/z0-backend-v1/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	// connection to the database
	config.ConnectDB()

	// initialize app
	app := fiber.New()

	// user routes
	routes.UserRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
