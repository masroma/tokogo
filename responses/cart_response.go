package responses

import "tokogo/models"

type CartItemResponse struct {
	ID           uint    `json:"id"`
	UserID       uint    `json:"user_id"`
	ProductID    uint    `json:"product_id"`
	ProductName  string  `json:"product_name"`
	ProductPrice float64 `json:"product_price"`
	ProductImage string  `json:"product_image"`
	Quantity     int     `json:"quantity"`
	Subtotal     float64 `json:"subtotal"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

type CartResponse struct {
	Items      []CartItemResponse `json:"items"`
	TotalItems int                `json:"total_items"`
	TotalPrice float64            `json:"total_price"`
}

func ConvertCartToResponse(cart models.Cart) CartItemResponse {
	subtotal := float64(cart.Quantity) * cart.Product.SellingPrice

	return CartItemResponse{
		ID:           cart.ID,
		UserID:       cart.UserID,
		ProductID:    cart.ProductID,
		ProductName:  cart.Product.Name,
		ProductPrice: cart.Product.SellingPrice,
		ProductImage: cart.Product.ImageURL,
		Quantity:     cart.Quantity,
		Subtotal:     subtotal,
		CreatedAt:    cart.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    cart.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ConvertCartsToResponse(carts []models.Cart) []CartItemResponse {
	var responses []CartItemResponse
	var totalPrice float64

	for _, cart := range carts {
		response := ConvertCartToResponse(cart)
		responses = append(responses, response)
		totalPrice += response.Subtotal
	}

	return responses
}

func CreateCartResponse(carts []models.Cart) CartResponse {
	items := ConvertCartsToResponse(carts)

	var totalItems int
	var totalPrice float64

	for _, item := range items {
		totalItems += item.Quantity
		totalPrice += item.Subtotal
	}

	return CartResponse{
		Items:      items,
		TotalItems: totalItems,
		TotalPrice: totalPrice,
	}
}
