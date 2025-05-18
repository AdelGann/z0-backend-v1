package authmiddleware

import (
	"log"
	"os"
	"strings"

	"github.com/AdelGann/z0-backend-v1/pkg/utils/helpers/validations"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var SecretKey string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env")
	}
	SecretKey = os.Getenv("JWT_SECRET")
}

func ValidateToken(c *fiber.Ctx) error {
	Authorization := c.Get("Authorization")
	if Authorization == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.ErrUnauthorized)
	}

	splited := strings.Split(Authorization, " ")
	if len(splited) != 2 || splited[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.ErrUnauthorized)
	}

	token, err := jwt.Parse(splited[1], func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.ErrUnauthorized)
	}

	return c.Next()
}

func RoleMiddleware(allowedRoles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if !validations.ValidateRole(c, allowedRoles) {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Access denied",
			})
		}
		return c.Next()
	}
}
