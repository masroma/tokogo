package repositories

import (
	"tokogo/config"
	"tokogo/models"

	"gorm.io/gorm"
)

type UserManagementRepository struct {
	db *gorm.DB
}

// NewUserManagementRepository membuat instance baru UserManagementRepository
func NewUserManagementRepository() *UserManagementRepository {
	return &UserManagementRepository{
		db: config.DB,
	}
}

// CreateUser membuat user baru
func (r *UserManagementRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

// GetUserByID mengambil user berdasarkan ID
func (r *UserManagementRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByEmail mengambil user berdasarkan email
func (r *UserManagementRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAllUsers mengambil semua users dengan pagination
func (r *UserManagementRepository) GetAllUsers(page, limit int) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	// Hitung total
	if err := r.db.Model(&models.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Ambil data dengan pagination
	offset := (page - 1) * limit
	err := r.db.Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// UpdateUser mengupdate user
func (r *UserManagementRepository) UpdateUser(user *models.User) error {
	return r.db.Save(user).Error
}

// DeleteUser menghapus user (soft delete)
func (r *UserManagementRepository) DeleteUser(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

// UpdateUserRole mengupdate role user
func (r *UserManagementRepository) UpdateUserRole(id uint, role string) error {
	return r.db.Model(&models.User{}).Where("id = ?", id).Update("role", role).Error
}

// GetUsersByRole mengambil users berdasarkan role
func (r *UserManagementRepository) GetUsersByRole(role string, page, limit int) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	// Hitung total
	if err := r.db.Model(&models.User{}).Where("role = ?", role).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Ambil data dengan pagination
	offset := (page - 1) * limit
	err := r.db.Where("role = ?", role).Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
