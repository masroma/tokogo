package requests

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

// RegisterRequest represents the request structure for user registration
type RegisterRequest struct {
	Username        string `json:"username" validate:"required,min=3,max=50"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=6"`
}

// Validate validates the RegisterRequest using the validator
func (r *RegisterRequest) Validate() error {
	validate := validator.New()

	// Validasi struct fields
	if err := validate.Struct(r); err != nil {
		return err
	}

	// Validasi custom: password dan confirm_password harus sama
	if r.Password != r.ConfirmPassword {
		return errors.New("password and confirm_password must match")
	}

	return nil
}
