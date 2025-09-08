package requests

import "errors"

type AddToCartRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,min=1"`
}

func (r *AddToCartRequest) Validate() error {
	if r.ProductID == 0 {
		return errors.New("product_id is required")
	}
	if r.Quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}
	return nil
}

type UpdateCartItemRequest struct {
	Quantity int `json:"quantity" binding:"required,min=1"`
}

func (r *UpdateCartItemRequest) Validate() error {
	if r.Quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}
	return nil
}

type RemoveFromCartRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
}

func (r *RemoveFromCartRequest) Validate() error {
	if r.ProductID == 0 {
		return errors.New("product_id is required")
	}
	return nil
}
