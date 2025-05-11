package userinputs

type CreateUserInput struct {
	FullName string `json:"full_name" validate:"required"`
	UserName string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
