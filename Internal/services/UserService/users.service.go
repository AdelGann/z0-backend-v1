package userservices

import (
	"errors"
	"github.com/AdelGann/z0-backend-v1/Internal/inputs/OrgInputs"
	"github.com/AdelGann/z0-backend-v1/Internal/inputs/UsersInput"
	"github.com/AdelGann/z0-backend-v1/Internal/services/OrgServices"
	"github.com/AdelGann/z0-backend-v1/config"
	"github.com/AdelGann/z0-backend-v1/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"reflect"
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
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return models.User{}, err
	}

	Values := reflect.ValueOf(user)

	ExistingEmail, _ := GetUserByEmail(&user.Email)          // for validate if email exists before the db throws an error
	ExistingUserName, _ := GetUserByUserName(&user.UserName) // for validate if username exists before the db throws an error

	if len(ExistingUserName.ID) > 0 {
		return models.User{}, errors.New("UserName already registered: " + ExistingUserName.UserName)
	}

	if len(ExistingEmail.ID) > 0 {
		return models.User{}, errors.New("Email already registered: " + ExistingEmail.Email)
	}

	for i := 0; i < Values.NumField(); i++ {
		fieldValue := Values.Field(i).String()

		if len(fieldValue) == 0 {
			return models.User{}, errors.New("Field " + Values.Type().Field(i).Name + " cannot be null")
		}

	}
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

	newOrg := orginputs.CreateOrgInput{
		Name: newUser.UserName + "'s Organization",
	}

	_, err = orgservices.SaveOrg(newOrg, newUser.ID)

	if err != nil {
		return models.User{}, err
	}

	return newUser, nil
}
