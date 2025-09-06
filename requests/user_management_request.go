package requests

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

// CreateUserRequest represents the request structure for creating user
type CreateUserRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=255"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Role     string `json:"role" validate:"required,oneof=customer admin"`
}

// UpdateUserRequest represents the request structure for updating user
type UpdateUserRequest struct {
	Name  string `json:"name" validate:"omitempty,min=3,max=255"`
	Email string `json:"email" validate:"omitempty,email"`
	Role  string `json:"role" validate:"omitempty,oneof=customer admin"`
}

// ChangePasswordRequest represents the request structure for changing password
type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required,min=6"`
}

// Validate validates the CreateUserRequest using the validator
func (r *CreateUserRequest) Validate() error {
	validate := validator.New()

	// Validasi struct fields
	if err := validate.Struct(r); err != nil {
		return err
	}

	// Validasi custom: nama tidak boleh kosong setelah trim
	if strings.TrimSpace(r.Name) == "" {
		return errors.New("name cannot be empty")
	}

	// Validasi custom: email tidak boleh kosong setelah trim
	if strings.TrimSpace(r.Email) == "" {
		return errors.New("email cannot be empty")
	}

	// Validasi custom: password tidak boleh kosong setelah trim
	if strings.TrimSpace(r.Password) == "" {
		return errors.New("password cannot be empty")
	}

	return nil
}

// Validate validates the UpdateUserRequest using the validator
func (r *UpdateUserRequest) Validate() error {
	validate := validator.New()

	// Validasi struct fields
	if err := validate.Struct(r); err != nil {
		return err
	}

	// Validasi custom: jika name diisi, tidak boleh kosong setelah trim
	if r.Name != "" && strings.TrimSpace(r.Name) == "" {
		return errors.New("name cannot be empty")
	}

	// Validasi custom: jika email diisi, tidak boleh kosong setelah trim
	if r.Email != "" && strings.TrimSpace(r.Email) == "" {
		return errors.New("email cannot be empty")
	}

	return nil
}

// Validate validates the ChangePasswordRequest using the validator
func (r *ChangePasswordRequest) Validate() error {
	validate := validator.New()

	// Validasi struct fields
	if err := validate.Struct(r); err != nil {
		return err
	}

	// Validasi custom: current password tidak boleh kosong setelah trim
	if strings.TrimSpace(r.CurrentPassword) == "" {
		return errors.New("current password cannot be empty")
	}

	// Validasi custom: new password tidak boleh kosong setelah trim
	if strings.TrimSpace(r.NewPassword) == "" {
		return errors.New("new password cannot be empty")
	}

	return nil
}
