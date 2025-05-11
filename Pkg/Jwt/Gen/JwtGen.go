package jwtgen

import (
	"log"
	"os"
	"time"

	"github.com/AdelGann/z0-backend-v1/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func GenerateJWT(email string, id uuid.UUID, role models.Roles) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error charching file .env")
	}

	secret := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"sub":   id,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secret))
}
