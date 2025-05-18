package usercontroller

import (
	"fmt"

	userinputs "github.com/AdelGann/z0-backend-v1/internal/inputs/users_inputs"
	userservices "github.com/AdelGann/z0-backend-v1/internal/services/users_service"
	"github.com/gofiber/fiber/v2"
)

// GetUsers godoc
// @Summary Obtener todos los usuarios
// @Description Retorna una lista con todos los usuarios
// @Tags usuarios
// @Produce json
// @Success 200 {array} User
// @Failure 500 {object} fiber.Map{"error":string}
// @Router /users [get]
func GetUsers(c *fiber.Ctx) error {
	users, err := userservices.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch users"})
	}
	return c.JSON(users)
}

// GetUserById godoc
// @Summary Obtener usuario por ID
// @Description Retorna un usuario dado su ID
// @Tags usuarios
// @Produce json
// @Param id path string true "ID del usuario"
// @Success 200 {object} User
// @Failure 500 {object} fiber.Map{"error":string}
// @Router /users/{id} [get]
func GetUserById(c *fiber.Ctx) error {
	user, err := userservices.GetUserById(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch user"})
	}
	return c.JSON(user)
}

// PostUser godoc
// @Summary Crear un nuevo usuario
// @Description Crea un usuario con la información proporcionada
// @Tags usuarios
// @Accept json
// @Produce json
// @Param user body userinputs.CreateUserInput true "Datos del usuario"
// @Success 201 {object} User
// @Failure 400 {object} fiber.Map{"error":string}
// @Router /users [post]
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
