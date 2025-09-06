package requests

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

// UpdateTransactionStatusRequest represents the request structure for updating transaction status
type UpdateTransactionStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=pending paid failed expired"`
}

// GetTransactionsRequest represents the request structure for getting transactions with filters
type GetTransactionsRequest struct {
	Status string `json:"status" validate:"omitempty,oneof=pending paid failed expired"`
	Page   int    `json:"page" validate:"omitempty,min=1"`
	Limit  int    `json:"limit" validate:"omitempty,min=1,max=100"`
}

// Validate validates the UpdateTransactionStatusRequest using the validator
func (r *UpdateTransactionStatusRequest) Validate() error {
	validate := validator.New()

	// Validasi struct fields
	if err := validate.Struct(r); err != nil {
		return err
	}

	// Validasi custom: status tidak boleh kosong setelah trim
	if strings.TrimSpace(r.Status) == "" {
		return errors.New("status cannot be empty")
	}

	return nil
}

// Validate validates the GetTransactionsRequest using the validator
func (r *GetTransactionsRequest) Validate() error {
	validate := validator.New()

	// Validasi struct fields
	if err := validate.Struct(r); err != nil {
		return err
	}

	// Set default values
	if r.Page <= 0 {
		r.Page = 1
	}
	if r.Limit <= 0 {
		r.Limit = 10
	}

	return nil
}
