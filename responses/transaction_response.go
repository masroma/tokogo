package responses

import "time"

// TransactionResponse represents the response structure for transaction
type TransactionResponse struct {
	ID          int64                       `json:"id"`
	UserID      int64                       `json:"user_id"`
	UserName    string                      `json:"user_name"`
	UserEmail   string                      `json:"user_email"`
	Status      string                      `json:"status"`
	TotalAmount float64                     `json:"total_amount"`
	PaymentURL  string                      `json:"payment_url,omitempty"`
	CreatedAt   time.Time                   `json:"created_at"`
	UpdatedAt   time.Time                   `json:"updated_at"`
	Details     []TransactionDetailResponse `json:"details,omitempty"`
}

// TransactionDetailResponse represents the response structure for transaction detail
type TransactionDetailResponse struct {
	ID            int64   `json:"id"`
	TransactionID int64   `json:"transaction_id"`
	ProductID     int64   `json:"product_id"`
	ProductName   string  `json:"product_name"`
	ProductImage  string  `json:"product_image,omitempty"`
	Quantity      int     `json:"quantity"`
	Price         float64 `json:"price"`
	Subtotal      float64 `json:"subtotal"`
}

// TransactionListResponse represents the response structure for transaction list
type TransactionListResponse struct {
	Transactions []TransactionResponse `json:"transactions"`
	Total        int64                 `json:"total"`
	Page         int                   `json:"page"`
	Limit        int                   `json:"limit"`
	TotalPages   int                   `json:"total_pages"`
}

// TransactionStatusResponse represents the response structure for transaction status update
type TransactionStatusResponse struct {
	Message string              `json:"message"`
	Data    TransactionResponse `json:"data"`
}
