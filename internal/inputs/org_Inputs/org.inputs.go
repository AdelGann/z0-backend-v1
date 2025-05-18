package orginputs

import "github.com/google/uuid"

type CreateOrgInput struct {
	Name string `json:"org_name" validate:"required"`
}
type InviteOrgInput struct {
	OrgID uuid.UUID `json:"org_id" validate:"required"`
	Email string    `json:"user_email" validate:"required"`
}
