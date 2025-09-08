package requests

import "errors"

type CheckoutRequest struct {
	ShippingAddress string `json:"shipping_address" binding:"required"`
	PaymentMethod   string `json:"payment_method" binding:"required"`
	Notes           string `json:"notes"`
}

func (r *CheckoutRequest) Validate() error {
	if r.ShippingAddress == "" {
		return errors.New("shipping_address is required")
	}
	if r.PaymentMethod == "" {
		return errors.New("payment_method is required")
	}

	// Validate payment method
	validPaymentMethods := []string{"bank_transfer", "credit_card", "e_wallet", "cod"}
	isValid := false
	for _, method := range validPaymentMethods {
		if r.PaymentMethod == method {
			isValid = true
			break
		}
	}
	if !isValid {
		return errors.New("invalid payment method")
	}

	return nil
}

type ConfirmPaymentRequest struct {
	PaymentProof string `json:"payment_proof" binding:"required"`
	Notes        string `json:"notes"`
}

func (r *ConfirmPaymentRequest) Validate() error {
	if r.PaymentProof == "" {
		return errors.New("payment_proof is required")
	}
	return nil
}
