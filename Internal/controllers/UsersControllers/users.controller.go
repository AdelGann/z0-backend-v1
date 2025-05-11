package usercontroller

import (
	"github.com/AdelGann/z0-backend-v1/Internal/inputs/UsersInput"
	"github.com/AdelGann/z0-backend-v1/Internal/services/UserService"
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
		return c.Status(400).JSON(fiber.Map{"error": "Failed while parsing the body"})
	}
	response, err := userservices.SaveUser(*user)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	return c.Status(fiber.StatusCreated).JSON(response)

}
