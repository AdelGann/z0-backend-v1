package authservices

import (
	"errors"
	"fmt"

	"github.com/AdelGann/z0-backend-v1/internal/dtos/auth_dto"
	"github.com/AdelGann/z0-backend-v1/internal/inputs/auth_inputs"
	"github.com/AdelGann/z0-backend-v1/internal/inputs/users_inputs"
	"github.com/AdelGann/z0-backend-v1/internal/services/users_service"
	"github.com/AdelGann/z0-backend-v1/pkg/jwt/gen"
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
func Register(credentials authinputs.RegisterInput) (authdto.JwtResponse, error) {
	if credentials.Email == "" ||
		credentials.FullName == "" ||
		credentials.UserName == "" ||
		credentials.Password == "" ||
		credentials.RepeatPassword == "" {
		return authdto.JwtResponse{}, errors.New("all credentials are required")
	}

	if credentials.Password != credentials.RepeatPassword {
		return authdto.JwtResponse{}, errors.New("password aren't equal")
	}

	newUser := userinputs.CreateUserInput{
		FullName: credentials.FullName,
		UserName: credentials.UserName,
		Email:    credentials.Email,
		Password: credentials.RepeatPassword,
	}

	User, err := userservices.SaveUser(newUser)
	if err != nil {
		fmt.Println(err)
		return authdto.JwtResponse{}, errors.New(err.Error())
	}

	token, err := jwtgen.GenerateJWT(User.Email, User.ID, User.Role)
	if err != nil {
		fmt.Println(err)
		return authdto.JwtResponse{}, errors.New("error generating token")
	}

	return authdto.JwtResponse{Token: token}, nil
}
