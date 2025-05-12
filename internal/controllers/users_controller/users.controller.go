package usercontroller

import (
	"fmt"

	"github.com/AdelGann/z0-backend-v1/internal/inputs/users_inputs"
	"github.com/AdelGann/z0-backend-v1/internal/services/users_service"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	users, err := userservices.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch users"})
	}
	return c.JSON(users)
}
func GetUserById(c *fiber.Ctx) error {
	user, err := userservices.GetUserById(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch users"})
	}
	return c.JSON(user)
}
func PostUser(c *fiber.Ctx) error {
	user := new(userinputs.CreateUserInput)

	if err := c.BodyParser(user); err != nil {
		fmt.Println(err)
		return c.Status(400).JSON(fiber.Map{"error": "Failed while parsing the body"})
	}
	response, err := userservices.SaveUser(*user)

	if err != nil {
		fmt.Print(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(response)

}
