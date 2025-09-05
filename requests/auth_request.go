package requests

import "github.com/go-playground/validator/v10"

// RegisterRequest represents the request structure for user registration
type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// LoginRequest represents the request structure for user login
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// Validate validates the RegisterRequest using the validator
func (r *RegisterRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

// Validate validates the LoginRequest using the validator
func (r *LoginRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
