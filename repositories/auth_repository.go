package repositories

import (
	"errors"
	"tokogo/config"
	"tokogo/models"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

// NewAuthRepository membuat instance baru AuthRepository
func NewAuthRepository() *AuthRepository {
	return &AuthRepository{
		db: config.DB,
	}
}

// CreateUser membuat user baru
func (r *AuthRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

// GetUserByEmail mengambil user berdasarkan email (untuk cek duplikasi)
func (r *AuthRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
