package responses

import "tokogo/models"

type CheckoutResponse struct {
	TransactionID   uint                   `json:"transaction_id"`
	UserID          uint                   `json:"user_id"`
	Status          string                 `json:"status"`
	TotalAmount     float64                `json:"total_amount"`
	ShippingAddress string                 `json:"shipping_address"`
	PaymentMethod   string                 `json:"payment_method"`
	PaymentURL      string                 `json:"payment_url,omitempty"`
	Items           []CheckoutItemResponse `json:"items"`
	CreatedAt       string                 `json:"created_at"`
	UpdatedAt       string                 `json:"updated_at"`
}

type CheckoutItemResponse struct {
	ProductID    uint    `json:"product_id"`
	ProductName  string  `json:"product_name"`
	ProductPrice float64 `json:"product_price"`
	Quantity     int     `json:"quantity"`
	Subtotal     float64 `json:"subtotal"`
}

type CheckoutSummaryResponse struct {
	TotalItems      int     `json:"total_items"`
	TotalAmount     float64 `json:"total_amount"`
	ShippingCost    float64 `json:"shipping_cost"`
	GrandTotal      float64 `json:"grand_total"`
	PaymentMethod   string  `json:"payment_method"`
	ShippingAddress string  `json:"shipping_address"`
}

func ConvertTransactionToCheckoutResponse(transaction models.Transaction) CheckoutResponse {
	var items []CheckoutItemResponse
	var totalAmount float64

	for _, detail := range transaction.TransactionDetails {
		item := CheckoutItemResponse{
			ProductID:    detail.ProductID,
			ProductName:  detail.Product.Name,
			ProductPrice: detail.Price,
			Quantity:     detail.Quantity,
			Subtotal:     float64(detail.Quantity) * detail.Price,
		}
		items = append(items, item)
		totalAmount += item.Subtotal
	}

	return CheckoutResponse{
		TransactionID:   transaction.ID,
		UserID:          transaction.UserID,
		Status:          transaction.Status,
		TotalAmount:     transaction.TotalAmount,
		ShippingAddress: transaction.ShippingAddress,
		PaymentMethod:   transaction.PaymentMethod,
		PaymentURL:      transaction.PaymentURL,
		Items:           items,
		CreatedAt:       transaction.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:       transaction.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func CreateCheckoutSummaryResponse(carts []models.Cart, shippingCost float64, paymentMethod, shippingAddress string) CheckoutSummaryResponse {
	var totalItems int
	var totalAmount float64

	for _, cart := range carts {
		totalItems += cart.Quantity
		totalAmount += float64(cart.Quantity) * cart.Product.SellingPrice
	}

	return CheckoutSummaryResponse{
		TotalItems:      totalItems,
		TotalAmount:     totalAmount,
		ShippingCost:    shippingCost,
		GrandTotal:      totalAmount + shippingCost,
		PaymentMethod:   paymentMethod,
		ShippingAddress: shippingAddress,
	}
}
