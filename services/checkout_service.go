package services

import (
	"errors"
	"fmt"
	"tokogo/config"
	"tokogo/models"
	"tokogo/repositories"
	"tokogo/requests"
	"tokogo/responses"
)

type CheckoutService struct {
	cartRepo        *repositories.CartRepository
	productRepo     *repositories.ProductRepository
	transactionRepo *repositories.TransactionRepository
}

func NewCheckoutService() *CheckoutService {
	return &CheckoutService{
		cartRepo:        repositories.NewCartRepository(config.DB),
		productRepo:     repositories.NewProductRepository(config.DB),
		transactionRepo: repositories.NewTransactionRepository(config.DB),
	}
}

func (s *CheckoutService) GetCheckoutSummary(userID uint, req requests.CheckoutRequest) (*responses.CheckoutSummaryResponse, error) {
	// Get user's cart
	carts, err := s.cartRepo.GetByUserID(userID)
	if err != nil {
		return nil, errors.New("failed to get cart")
	}

	if len(carts) == 0 {
		return nil, errors.New("cart is empty")
	}

	// Calculate shipping cost (simple logic - can be enhanced)
	shippingCost := s.calculateShippingCost(carts)

	// Create checkout summary
	summary := responses.CreateCheckoutSummaryResponse(carts, shippingCost, req.PaymentMethod, req.ShippingAddress)

	return &summary, nil
}

func (s *CheckoutService) ProcessCheckout(userID uint, req requests.CheckoutRequest) (*responses.CheckoutResponse, error) {
	// Get user's cart
	carts, err := s.cartRepo.GetByUserID(userID)
	if err != nil {
		return nil, errors.New("failed to get cart")
	}

	if len(carts) == 0 {
		return nil, errors.New("cart is empty")
	}

	// Validate stock for all items
	for _, cart := range carts {
		product, err := s.productRepo.GetByID(cart.ProductID)
		if err != nil {
			return nil, fmt.Errorf("product with ID %d not found", cart.ProductID)
		}

		if product.Stock < cart.Quantity {
			return nil, fmt.Errorf("insufficient stock for product %s (available: %d, requested: %d)",
				product.Name, product.Stock, cart.Quantity)
		}
	}

	// Calculate total amount
	var totalAmount float64
	for _, cart := range carts {
		totalAmount += float64(cart.Quantity) * cart.Product.SellingPrice
	}

	// Add shipping cost
	shippingCost := s.calculateShippingCost(carts)
	totalAmount += shippingCost

	// Create transaction
	transaction := &models.Transaction{
		UserID:          userID,
		Status:          "pending",
		TotalAmount:     totalAmount,
		ShippingAddress: req.ShippingAddress,
		PaymentMethod:   req.PaymentMethod,
		Notes:           req.Notes,
	}

	// Save transaction
	if err := s.transactionRepo.Create(transaction); err != nil {
		return nil, errors.New("failed to create transaction")
	}

	// Create transaction details
	for _, cart := range carts {
		detail := &models.TransactionDetail{
			TransactionID: transaction.ID,
			ProductID:     cart.ProductID,
			Quantity:      cart.Quantity,
			Price:         cart.Product.SellingPrice,
		}

		if err := s.transactionRepo.CreateTransactionDetail(detail); err != nil {
			return nil, errors.New("failed to create transaction detail")
		}

		// Update product stock
		product, _ := s.productRepo.GetByID(cart.ProductID)
		product.Stock -= cart.Quantity
		if err := s.productRepo.Update(product); err != nil {
			return nil, errors.New("failed to update product stock")
		}
	}

	// Generate payment URL (mock implementation)
	paymentURL := s.generatePaymentURL(transaction.ID, req.PaymentMethod)
	transaction.PaymentURL = paymentURL

	if err := s.transactionRepo.Update(transaction); err != nil {
		return nil, errors.New("failed to update payment URL")
	}

	// Clear user's cart
	if err := s.cartRepo.ClearCart(userID); err != nil {
		// Log error but don't fail the checkout
		fmt.Printf("Warning: Failed to clear cart for user %d: %v\n", userID, err)
	}

	// Get transaction with details for response
	createdTransaction, err := s.transactionRepo.GetByID(transaction.ID)
	if err != nil {
		return nil, errors.New("failed to retrieve transaction")
	}

	response := responses.ConvertTransactionToCheckoutResponse(*createdTransaction)
	return &response, nil
}

func (s *CheckoutService) ConfirmPayment(userID uint, transactionID uint, req requests.ConfirmPaymentRequest) (*responses.CheckoutResponse, error) {
	// Get transaction
	transaction, err := s.transactionRepo.GetByID(transactionID)
	if err != nil {
		return nil, errors.New("transaction not found")
	}

	// Check if transaction belongs to user
	if transaction.UserID != userID {
		return nil, errors.New("unauthorized access to transaction")
	}

	// Check if transaction is in pending status
	if transaction.Status != "pending" {
		return nil, errors.New("transaction is not in pending status")
	}

	// Update transaction status to paid
	transaction.Status = "paid"
	transaction.PaymentProof = req.PaymentProof
	transaction.Notes = req.Notes

	if err := s.transactionRepo.Update(transaction); err != nil {
		return nil, errors.New("failed to update transaction status")
	}

	response := responses.ConvertTransactionToCheckoutResponse(*transaction)
	return &response, nil
}

func (s *CheckoutService) GetUserTransactions(userID uint) ([]responses.CheckoutResponse, error) {
	transactions, err := s.transactionRepo.GetByUserID(userID)
	if err != nil {
		return nil, errors.New("failed to get transactions")
	}

	var responseList []responses.CheckoutResponse
	for _, transaction := range transactions {
		response := responses.ConvertTransactionToCheckoutResponse(transaction)
		responseList = append(responseList, response)
	}

	return responseList, nil
}

func (s *CheckoutService) GetTransactionByID(userID uint, transactionID uint) (*responses.CheckoutResponse, error) {
	transaction, err := s.transactionRepo.GetByID(transactionID)
	if err != nil {
		return nil, errors.New("transaction not found")
	}

	// Check if transaction belongs to user
	if transaction.UserID != userID {
		return nil, errors.New("unauthorized access to transaction")
	}

	response := responses.ConvertTransactionToCheckoutResponse(*transaction)
	return &response, nil
}

// Helper methods
func (s *CheckoutService) calculateShippingCost(carts []models.Cart) float64 {
	// Simple shipping calculation - can be enhanced with more complex logic
	var totalWeight float64
	for _, cart := range carts {
		// Assume each product has weight of 1kg (can be added to product model)
		totalWeight += float64(cart.Quantity)
	}

	// Shipping cost: 5000 per kg, minimum 10000
	shippingCost := totalWeight * 5000
	if shippingCost < 10000 {
		shippingCost = 10000
	}

	return shippingCost
}

func (s *CheckoutService) generatePaymentURL(transactionID uint, paymentMethod string) string {
	// Mock payment URL generation - in real implementation, integrate with payment gateway
	baseURL := "https://payment.example.com"

	switch paymentMethod {
	case "bank_transfer":
		return fmt.Sprintf("%s/bank-transfer?transaction_id=%d", baseURL, transactionID)
	case "credit_card":
		return fmt.Sprintf("%s/credit-card?transaction_id=%d", baseURL, transactionID)
	case "e_wallet":
		return fmt.Sprintf("%s/e-wallet?transaction_id=%d", baseURL, transactionID)
	case "cod":
		return "" // COD doesn't need payment URL
	default:
		return fmt.Sprintf("%s/payment?transaction_id=%d", baseURL, transactionID)
	}
}
