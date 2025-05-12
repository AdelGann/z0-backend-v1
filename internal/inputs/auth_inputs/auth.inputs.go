package authinputs

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
type RegisterInput struct {
	FullName       string `json:"full_name" validate:"required"`
	UserName       string `json:"username" validate:"required"`
	Email          string `json:"email" validate:"required,email"`
	Password       string `json:"password" validate:"required"`
	RepeatPassword string `json:"repeat_password" validate:"required"`
}
