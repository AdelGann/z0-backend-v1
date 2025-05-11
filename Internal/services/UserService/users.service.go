package userservices

import (
	"errors"
	"fmt"

	"github.com/AdelGann/z0-backend-v1/Internal/inputs/OrgInputs"
	"github.com/AdelGann/z0-backend-v1/Internal/inputs/UsersInput"
	"github.com/AdelGann/z0-backend-v1/Internal/services/OrgServices"
	"github.com/AdelGann/z0-backend-v1/config"
	"github.com/AdelGann/z0-backend-v1/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func UserToDto() {}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := config.DB.Find(&users)
	return users, result.Error
}
func GetUserById(id string) (models.User, error) {
	var user models.User
	result := config.DB.Where("id = ?", id).Find(&user)
	return user, result.Error
}
func GetUserByEmail(email *string) (models.User, error) {
	var user models.User
	result := config.DB.Where("email = ?", *email).Find(&user)
	return user, result.Error
}
func GetUserByUserName(username *string) (models.User, error) {
	var user models.User
	userName := "@" + *username
	result := config.DB.Where("user_name = ?", userName).Find(&user)
	return user, result.Error
}

func SaveUser(user userinputs.CreateUserInput) (models.User, error) {

	if user.FullName == "" || user.UserName == "" || user.Email == "" || user.Password == "" {
		return models.User{}, errors.New("all fields are required")
	}

	ExistingEmail, _ := GetUserByEmail(&user.Email)
	ExistingUserName, _ := GetUserByUserName(&user.UserName)

	if ExistingUserName.ID == uuid.Nil {
		return models.User{}, errors.New("UserName already registered: " + ExistingUserName.UserName)
	}

	if ExistingEmail.ID == uuid.Nil {
		return models.User{}, errors.New("Email already registered: " + ExistingEmail.Email)
	}

	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	// Crear usuario
	newUser := models.User{
		ID:       uuid.New(),
		FullName: user.FullName,
		UserName: "@" + user.UserName,
		Email:    user.Email,
		Password: string(passwordHashed),
	}

	if err := config.DB.Create(&newUser).Error; err != nil {
		return models.User{}, err
	}

	// Crear organización
	newOrg := orginputs.CreateOrgInput{
		Name: fmt.Sprintf("%s's Organization", newUser.UserName),
	}

	_, err = orgservices.SaveOrg(newOrg, newUser.ID)

	if err != nil {
		return models.User{}, err
	}

	return newUser, nil
}
