package repositories

import (
	"tokogo/models"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

// NewTransactionRepository membuat instance baru TransactionRepository
func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
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

// Create membuat transaksi baru
func (r *TransactionRepository) Create(transaction *models.Transaction) error {
	return r.db.Create(transaction).Error
}

// Update mengupdate transaksi
func (r *TransactionRepository) Update(transaction *models.Transaction) error {
	return r.db.Save(transaction).Error
}

// GetByID mengambil transaksi berdasarkan ID
func (r *TransactionRepository) GetByID(id uint) (*models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").Preload("TransactionDetails.Product").First(&transaction, id).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

// GetByUserID mengambil transaksi berdasarkan User ID
func (r *TransactionRepository) GetByUserID(userID uint) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("User").Preload("TransactionDetails.Product").Where("user_id = ?", userID).Order("created_at DESC").Find(&transactions).Error
	return transactions, err
}

// CreateTransactionDetail membuat detail transaksi baru
func (r *TransactionRepository) CreateTransactionDetail(detail *models.TransactionDetail) error {
	return r.db.Create(detail).Error
}
