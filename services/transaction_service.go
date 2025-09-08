package services

import (
	"tokogo/config"
	"tokogo/repositories"
	"tokogo/requests"
	"tokogo/responses"
)

type TransactionService struct {
	transactionRepo *repositories.TransactionRepository
}

// NewTransactionService membuat instance baru TransactionService
func NewTransactionService() *TransactionService {
	return &TransactionService{
		transactionRepo: repositories.NewTransactionRepository(config.DB),
	}
}

// GetAllTransactions mengambil semua transaksi dengan pagination dan filter
func (s *TransactionService) GetAllTransactions(page, limit int, status string) (*responses.TransactionListResponse, error) {
	// Set default values
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	// Get transactions from repository
	transactions, total, err := s.transactionRepo.GetAllTransactions(page, limit, status)
	if err != nil {
		return nil, err
	}

	// Convert to response format
	var transactionResponses []responses.TransactionResponse
	for _, transaction := range transactions {
		transactionResponse := responses.TransactionResponse{
			ID:          int64(transaction.ID),
			UserID:      int64(transaction.UserID),
			UserName:    transaction.User.Name,
			UserEmail:   transaction.User.Email,
			Status:      transaction.Status,
			TotalAmount: transaction.TotalAmount,
			PaymentURL:  transaction.PaymentURL,
			CreatedAt:   transaction.CreatedAt,
			UpdatedAt:   transaction.UpdatedAt,
		}
		transactionResponses = append(transactionResponses, transactionResponse)
	}

	return &responses.TransactionListResponse{
		Transactions: transactionResponses,
		Total:        int64(total),
		Page:         page,
		Limit:        limit,
	}, nil
}

// GetTransactionByID mengambil transaksi berdasarkan ID dengan detail
func (s *TransactionService) GetTransactionByID(id uint) (*responses.TransactionResponse, error) {
	// Get transaction from repository
	transaction, err := s.transactionRepo.GetTransactionByID(id)
	if err != nil {
		return nil, err
	}

	// Convert details to response format
	var detailResponses []responses.TransactionDetailResponse
	for _, detail := range transaction.TransactionDetails {
		detailResponse := responses.TransactionDetailResponse{
			ID:            int64(detail.ID),
			TransactionID: int64(detail.TransactionID),
			ProductID:     int64(detail.ProductID),
			ProductName:   detail.Product.Name,
			ProductImage:  detail.Product.ImageURL,
			Quantity:      detail.Quantity,
			Price:         detail.Price,
			Subtotal:      float64(detail.Quantity) * detail.Price,
		}
		detailResponses = append(detailResponses, detailResponse)
	}

	// Convert to response format
	transactionResponse := &responses.TransactionResponse{
		ID:          int64(transaction.ID),
		UserID:      int64(transaction.UserID),
		UserName:    transaction.User.Name,
		UserEmail:   transaction.User.Email,
		Status:      transaction.Status,
		TotalAmount: transaction.TotalAmount,
		PaymentURL:  transaction.PaymentURL,
		CreatedAt:   transaction.CreatedAt,
		UpdatedAt:   transaction.UpdatedAt,
		Details:     detailResponses,
	}

	return transactionResponse, nil
}

// UpdateTransactionStatus mengupdate status transaksi
func (s *TransactionService) UpdateTransactionStatus(id uint, req requests.UpdateTransactionStatusRequest) (*responses.TransactionResponse, error) {
	// Validasi request
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// Cek apakah transaction ada
	_, err := s.transactionRepo.GetTransactionByID(id)
	if err != nil {
		return nil, err
	}

	// Update transaction status
	err = s.transactionRepo.UpdateTransactionStatus(id, req.Status)
	if err != nil {
		return nil, err
	}

	// Get updated transaction
	transactionResponse, err := s.GetTransactionByID(id)
	if err != nil {
		return nil, err
	}

	return transactionResponse, nil
}
