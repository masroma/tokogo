package requests

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

// CreateCategoryRequest represents the request structure for creating category
type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required,min=2,max=100"`
}

// UpdateCategoryRequest represents the request structure for updating category
type UpdateCategoryRequest struct {
	Name string `json:"name" validate:"required,min=2,max=100"`
}

// Validate validates the CreateCategoryRequest using the validator
func (r *CreateCategoryRequest) Validate() error {
	validate := validator.New()

	// Validasi struct fields
	if err := validate.Struct(r); err != nil {
		return err
	}

	// Validasi custom: nama tidak boleh kosong setelah trim
	if strings.TrimSpace(r.Name) == "" {
		return errors.New("name cannot be empty")
	}

	return nil
}

// Validate validates the UpdateCategoryRequest using the validator
func (r *UpdateCategoryRequest) Validate() error {
	validate := validator.New()

	// Validasi struct fields
	if err := validate.Struct(r); err != nil {
		return err
	}

	// Validasi custom: nama tidak boleh kosong setelah trim
	if strings.TrimSpace(r.Name) == "" {
		return errors.New("name cannot be empty")
	}

	return nil
}
