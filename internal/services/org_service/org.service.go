package orgservices

import (
	"errors"
	"github.com/AdelGann/z0-backend-v1/config"
	"github.com/AdelGann/z0-backend-v1/internal/inputs/org_inputs"
	"github.com/AdelGann/z0-backend-v1/models"
	"github.com/AdelGann/z0-backend-v1/pkg/utils/helpers/gen"
	"github.com/AdelGann/z0-backend-v1/pkg/utils/mail"
	"github.com/google/uuid"
	"time"
)

func SaveOrg(org orginputs.CreateOrgInput, founderID uuid.UUID) (models.Org, error) {
	if founderID == uuid.Nil {
		return models.Org{}, errors.New("founderID is required")
	}
	if org.Name == "" {
		return models.Org{}, errors.New("organization name is required")
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

func JoinOrg(code string, userID uuid.UUID) (models.Employee, error) {
	// getting user invitations
	var invitation models.OrgInvitation
	res := config.DB.Where("user_id = ?", userID).Where("state = ?", models.PENDING).First(&invitation)
	if res.Error != nil {
		return models.Employee{}, res.Error
	}
	if invitation.ID != uuid.Nil {
		// Check if the invitation is older than 30 minutes
		if time.Since(invitation.CreatedAt) >= 30*time.Minute {
			// if it is, cancel the invitation
			err := config.DB.Model(&models.OrgInvitation{}).
				Where("id = ?", invitation.Code).
				Update("state", models.CANCELED).Error
			if err != nil {
				return models.Employee{}, err
			}
			// return an error if invitation is older than 30 minutes
			return models.Employee{}, errors.New("invitation expired")
		}
	}
	if invitation.Code != code {
		// Check if the invitation code is correct
		return models.Employee{}, errors.New("the code provided is not correct")
	}
	// generate a new employee
	newEmployee := models.Employee{
		ID:     uuid.New(),
		OrgID:  invitation.OrgID,
		UserID: userID,
		Role:   models.USER,
	}
	// if the employee already exists, return an error
	if config.DB.Where("user_id = ? AND org_id = ?", userID, invitation.OrgID).First(&models.Employee{}).RowsAffected > 0 {
		return models.Employee{}, errors.New("user already exists in the organization")
	}
	if err := config.DB.Create(&newEmployee).Error; err != nil {
		// if there is an error creating the employee, return an error
		return models.Employee{}, err
	}
	// if there is no error, update the invitation state to accepted
	// if there is an error updating the invitation, return an error
	if err := config.DB.Model(&models.OrgInvitation{}).
		Where("id = ?", invitation.Code).
		Update("state", models.ACCEPTED).Error; err != nil {
		return models.Employee{}, err
	}
	// if there is no error, update the invitation state to accepted
	return newEmployee, nil
}

func SendInvitation(OrgID uuid.UUID, UserEmail string, founderID uuid.UUID) (string, error) {
	// loading mail settings
	mail.Builder()
	// getting organization
	org, err := GetOrgById(OrgID)
	// if the organization is not founded, return an error
	if err != nil {
		return "", errors.New("org not founded")
	}
	// if the organization is not founded, return an error
	if org.FounderID != founderID {
		return "", errors.New("only founders have the privilege to invite others to join their organizations")
	}
	// searching user by email
	user := models.User{}
	res := config.DB.Where("email = ?", UserEmail).Find(&user)

	if res.Error != nil {
		return "", errors.New(res.Error.Error())
	}
	// searching if user already exists in Employee table
	// if the user already exists, return an error
	var existingEmployee models.Employee
	result := config.DB.Where("user_id = ?", user.ID).First(&existingEmployee)
	if result.Error == nil {
		return "", errors.New("user already exists")
	}
	// searching if user already exists in OrgInvitation table
	invitations := models.OrgInvitation{}
	err = config.DB.Where("user_id = ?", user.ID).Where("org_id", OrgID).First(&invitations).Error
	if err != nil {
		return "", errors.New(err.Error())
	}
	// if user already have an invitation, check if it is older than 30 minutes
	if invitations.ID != uuid.Nil {
		if time.Since(invitations.CreatedAt) >= 30*time.Minute {
			err := config.DB.Model(&models.OrgInvitation{}).
				Where("id = ?", invitations.ID).
				Update("state", models.CANCELED).Error
			if err != nil {
				return "", err
			}
		} else {
			// if the invitation is not older than 30 minutes, return an error
			return "", errors.New("an invitation already exists for user " + UserEmail)
		}
	}
	// generating a new invitation code
	code, err := gen.GenerateCode(6)
	if err != nil {
		return "", errors.New(err.Error())
	}
	// generating invitation model for db
	sendInv := models.OrgInvitation{
		ID:     uuid.New(),
		UserID: user.ID,
		OrgID:  OrgID,
		Code:   string(code),
		State:  models.PENDING,
	}
	// creating the invitation in the db
	// if there is an error creating the invitation, return an error
	if err := config.DB.Create(&sendInv).Error; err != nil {
		return "", err
	}
	// sending the invitation email
	// generating email message
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
	// setting up the email addresses
	var addresses []string
	addresses = append(addresses, UserEmail)

	// sending the email
	err = mail.SendEmailSSL(msg, addresses)
	if err != nil {
		// if there is an error sending the email, return an error
		return "", err
	}
	// if there is no error, return the code
	return string(code), nil
}
