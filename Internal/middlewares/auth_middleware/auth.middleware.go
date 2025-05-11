package authmiddleware

import (
	"errors"
	"log"
	"os"
	"strings"
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

func ExtractClaims(c *fiber.Ctx) (map[string]interface{}, error) {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return nil, errors.New(fiber.ErrUnauthorized.Message)
	}

	parts := strings.Split(tokenString, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, errors.New(fiber.ErrUnauthorized.Message)
	}

	token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New(fiber.ErrUnauthorized.Message)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, errors.New(fiber.ErrUnauthorized.Message)
}

func ValidateRole(c *fiber.Ctx, allowedRoles []string) bool {
	claims, err := ExtractClaims(c)
	if err != nil {
		return false
	}

	role, ok := claims["role"].(string)
	if !ok {
		return false
	}

	for _, allowedRole := range allowedRoles {
		if role == allowedRole {
			return true
		}
	}

	return false
}
