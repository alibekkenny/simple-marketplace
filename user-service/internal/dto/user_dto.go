package dto

type RegisterInput struct {
	Email    string `validate:"required,email"`
	Username string `validate:"required,min=3"`
	Password string `validate:"required,min=3"`
	Role     string `validate:"required"`
}

type LoginInput struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=3"`
}
