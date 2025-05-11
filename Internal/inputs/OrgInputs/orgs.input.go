package orginputs

type CreateOrgInput struct {
	Name string `json:"org_name" validate:"required"`
}
