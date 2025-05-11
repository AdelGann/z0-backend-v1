package authservices

import (
	"errors"
	"fmt"

	authdto "github.com/AdelGann/z0-backend-v1/Internal/dtos/AuthDto"
	"github.com/AdelGann/z0-backend-v1/Internal/inputs/AuthInputs"
	userservices "github.com/AdelGann/z0-backend-v1/Internal/services/UserService"
	jwtgen "github.com/AdelGann/z0-backend-v1/Pkg/Jwt/Gen"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Login(credentials authinputs.LoginInput) (authdto.JwtResponse, error) {

	if credentials.Email == "" || credentials.Password == "" {
		return authdto.JwtResponse{}, errors.New("email and password cannot be empty")
	}

	User, _ := userservices.GetUserByEmail(&credentials.Email)

	if User.ID == uuid.Nil {
		return authdto.JwtResponse{}, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(credentials.Password)); err != nil {
		fmt.Println(err)
		return authdto.JwtResponse{}, errors.New("invalid credentials")
	}

	token, err := jwtgen.GenerateJWT(User.Email, User.ID, User.Role)
	if err != nil {
		fmt.Println(err)
		return authdto.JwtResponse{}, errors.New("error generating token")
	}

	return authdto.JwtResponse{Token: token}, nil
}
