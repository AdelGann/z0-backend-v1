package orgservices

import (
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/AdelGann/z0-backend-v1/config"
	"github.com/AdelGann/z0-backend-v1/internal/inputs/org_inputs"
	"github.com/AdelGann/z0-backend-v1/models"
	"github.com/AdelGann/z0-backend-v1/pkg/utils/mail"
	"github.com/google/uuid"
	"math/big"
)

func SaveOrg(org orginputs.CreateOrgInput, founderID uuid.UUID) (models.Org, error) {
	if len(founderID) == 0 {
		return models.Org{}, errors.New("FounderID is required")
	}
	if len(org.Name) == 0 {
		return models.Org{}, errors.New("OrgName is required")
	}
	newOrg := models.Org{
		ID:        uuid.New(),
		FounderID: founderID,
		Name:      org.Name,
	}
	if err := config.DB.Create(&newOrg).Error; err != nil {
		return models.Org{}, err
	}
	return models.Org{}, nil
}
func GetOrgById(id uuid.UUID) (models.Org, error) {
	var org models.Org
	result := config.DB.Where("id = ?", id).Find(&org)
	return org, result.Error
}
func JoinOrg() {

}
func SendInvitation(OrgID uuid.UUID, UserEmail string, founderID uuid.UUID) (string, error) {
	mail.LoadEnv()
	mail.Builder()

	org, err := GetOrgById(OrgID)

	if err != nil {
		return "", errors.New("org not founded")
	}

	if org.FounderID != founderID {
		return "", errors.New("only founders have the privilege to invite others to join their organizations")
	}

	user := models.User{}

	res := config.DB.Where("email = ?", UserEmail).Find(&user)

	if res.Error != nil {
		return "", errors.New(res.Error.Error())
	}

	length := 6
	code := make([]byte, length)
	for i := range code {
		var num *big.Int
		for {
			num, _ = rand.Int(rand.Reader, big.NewInt(127))
			if (num.Int64() >= 48 && num.Int64() <= 57) ||
				(num.Int64() >= 65 && num.Int64() <= 90) {
				break
			}
		}
		code[i] = byte(num.Int64())
	}

	sendInv := models.OrgInvitation{
		ID:     uuid.New(),
		UserID: user.ID,
		OrgID:  OrgID,
		Code:   string(code),
	}

	if err := config.DB.Create(&sendInv).Error; err != nil {
		return "", err
	}

	msg := []byte(
		"From: z0 Team <no-reply@z0-team.com>\r\n" +
			"To: " + UserEmail + "\r\n" +
			"Subject: Team Code Invitation\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/plain; charset=\"utf-8\"\r\n" +
			"\r\n" +
			"You have been invited to join " + org.Name + "\r\n\r\n" +
			"Your invitation code is: " + string(code) + "\r\n",
	)

	var addresses []string
	addresses = append(addresses, UserEmail)
	fmt.Print(addresses)

	mail.SendEmailSSL(msg, addresses)

	return string(code), nil
}
