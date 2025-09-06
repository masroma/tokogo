package requests

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

// CreateProductRequest represents the request structure for creating product
type CreateProductRequest struct {
	Name          string  `json:"name" validate:"required,min=3,max=255"`
	Description   string  `json:"description"`
	PurchasePrice float64 `json:"purchase_price" validate:"required,min=0"`
	SellingPrice  float64 `json:"selling_price" validate:"required,min=0"`
	Stock         int     `json:"stock" validate:"min=0"`
	CategoryID    uint    `json:"category_id" validate:"required"`
}

// UpdateProductRequest represents the request structure for updating product
type UpdateProductRequest struct {
	Name          string  `json:"name" validate:"required,min=3,max=255"`
	Description   string  `json:"description"`
	PurchasePrice float64 `json:"purchase_price" validate:"required,min=0"`
	SellingPrice  float64 `json:"selling_price" validate:"required,min=0"`
	Stock         int     `json:"stock" validate:"min=0"`
	CategoryID    uint    `json:"category_id" validate:"required"`
}

// Validate validates the CreateProductRequest using the validator
func (r *CreateProductRequest) Validate() error {
	validate := validator.New()

	// Validasi struct fields
	if err := validate.Struct(r); err != nil {
		return err
	}

	// Validasi custom: nama tidak boleh kosong setelah trim
	if strings.TrimSpace(r.Name) == "" {
		return errors.New("name cannot be empty")
	}

	// Validasi custom: selling price harus lebih besar dari purchase price
	if r.SellingPrice <= r.PurchasePrice {
		return errors.New("selling price must be greater than purchase price")
	}

	return nil
}

// Validate validates the UpdateProductRequest using the validator
func (r *UpdateProductRequest) Validate() error {
	validate := validator.New()

	// Validasi struct fields
	if err := validate.Struct(r); err != nil {
		return err
	}

	// Validasi custom: nama tidak boleh kosong setelah trim
	if strings.TrimSpace(r.Name) == "" {
		return errors.New("name cannot be empty")
	}

	// Validasi custom: selling price harus lebih besar dari purchase price
	if r.SellingPrice <= r.PurchasePrice {
		return errors.New("selling price must be greater than purchase price")
	}

	return nil
}
