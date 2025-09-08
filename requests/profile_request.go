package requests

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

// UpdateProfileRequest represents the request structure for updating profile
type UpdateProfileRequest struct {
	Name  string `json:"name" validate:"required,min=3,max=255"`
	Email string `json:"email" validate:"required,email,max=255"`
}

// ChangePasswordRequest represents the request structure for changing password
type ChangeUserPasswordRequest struct {
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required,min=6,max=255"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

// Validate validates the UpdateProfileRequest using the validator
func (r *UpdateProfileRequest) Validate() error {
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

	return nil
}

// Validate validates the ChangePasswordRequest using the validator
func (r *ChangeUserPasswordRequest) Validate() error {
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

	// Validasi custom: confirm password tidak boleh kosong setelah trim
	if strings.TrimSpace(r.ConfirmPassword) == "" {
		return errors.New("confirm password cannot be empty")
	}

	// Validasi password match
	if r.NewPassword != r.ConfirmPassword {
		return errors.New("new password and confirm password must match")
	}

	// Validasi password tidak sama dengan current
	if r.CurrentPassword == r.NewPassword {
		return errors.New("new password cannot be the same as current password")
	}

	return nil
}
