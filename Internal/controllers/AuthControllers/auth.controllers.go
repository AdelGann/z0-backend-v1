package authcontrollers

import (
	"fmt"

	"github.com/AdelGann/z0-backend-v1/Internal/inputs/AuthInputs"
	"github.com/AdelGann/z0-backend-v1/Internal/services/AuthService"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	credentials := new(authinputs.LoginInput)
	if err := c.BodyParser(credentials); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": err.Error()})
	}
	response, err := authservices.Login(*credentials)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(response)
}
