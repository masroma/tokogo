package repositories

import (
	"tokogo/config"
	"tokogo/models"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

// NewTransactionRepository membuat instance baru TransactionRepository
func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{
		db: config.DB,
	}
}

// GetAllTransactions mengambil semua transaksi dengan pagination dan filter
func (r *TransactionRepository) GetAllTransactions(page, limit int, status string) ([]models.Transaction, int64, error) {
	var transactions []models.Transaction
	var total int64

	// Build query dengan GORM
	query := r.db.Preload("User").Model(&models.Transaction{})

	// Apply filter status jika ada
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// Count total records
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination dan ambil data
	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&transactions).Error; err != nil {
		return nil, 0, err
	}

	return transactions, total, nil
}

// GetTransactionByID mengambil transaksi berdasarkan ID dengan detail
func (r *TransactionRepository) GetTransactionByID(id uint) (*models.Transaction, error) {
	var transaction models.Transaction

	// Get transaction dengan preload User dan TransactionDetails
	err := r.db.Preload("User").Preload("TransactionDetails.Product").First(&transaction, id).Error
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

// UpdateTransactionStatus mengupdate status transaksi
func (r *TransactionRepository) UpdateTransactionStatus(id uint, status string) error {
	err := r.db.Model(&models.Transaction{}).Where("id = ?", id).Update("status", status).Error
	if err != nil {
		return err
	}
	return nil
}
