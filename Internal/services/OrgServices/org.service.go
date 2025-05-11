package orgservices

import (
	"errors"

	"github.com/AdelGann/z0-backend-v1/Internal/inputs/OrgInputs"
	"github.com/AdelGann/z0-backend-v1/config"
	"github.com/AdelGann/z0-backend-v1/models"
	"github.com/google/uuid"
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
