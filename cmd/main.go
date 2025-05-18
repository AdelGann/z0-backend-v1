package main

import (
	"fmt"

	_ "github.com/AdelGann/z0-backend-v1/cmd/docs"
	"github.com/AdelGann/z0-backend-v1/config"
	"github.com/AdelGann/z0-backend-v1/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

// @title z0-api
// @version 1.0
// @description Documentación de ejemplo con Fiber y swaggo
// @host localhost:3000
// @BasePath /
func main() {
	// connection to the database
	config.ConnectDB()

	// initialize app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	app.Use(cors.New())
	// cors.Config{
	// 	AllowOrigins:     "*",
	// 	AllowMethods:     "GET,POST,PUT,DELETE,PATCH,OPTIONS",
	// 	AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
	// 	ExposeHeaders:    "Content-Length",
	// 	AllowCredentials: true,
	// }
	//TODO: implement cors settings when the frontend will integrate it

	app.Use(func(c *fiber.Ctx) (err error) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered from panic: %v\n", r)
				err = c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": fmt.Sprintf("%v", r),
				})
			} //? DO I NEED THIS HERE?
		}()
		return c.Next()
	})

	// swagger configurated
	app.Get("/swagger/*", swagger.HandlerDefault)

	// user routes
	routes.MainRoutes(app)

	log.Fatal(app.Listen("0.0.0.0:3000"))
}
