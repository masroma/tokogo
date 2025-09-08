package repositories

import (
	"errors"
	"tokogo/config"
	"tokogo/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type ProfileRepository struct {
	db *gorm.DB
}

// NewProfileRepository membuat instance baru ProfileRepository
func NewProfileRepository() *ProfileRepository {
	return &ProfileRepository{
		db: config.DB,
	}
}

// GetProfileByID mengambil profile user berdasarkan ID
func (r *ProfileRepository) GetProfileByID(userID uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateProfile mengupdate profile user
func (r *ProfileRepository) UpdateProfile(userID uint, name, email string) (*models.User, error) {
	var user models.User

	// Cek apakah user ada
	if err := r.db.First(&user, userID).Error; err != nil {
		return nil, errors.New("user not found")
	}

	// Cek apakah email sudah digunakan oleh user lain
	var existingUser models.User
	if err := r.db.Where("email = ? AND id != ?", email, userID).First(&existingUser).Error; err == nil {
		return nil, errors.New("email already exists")
	}

	// Update profile
	user.Name = name
	user.Email = email

	if err := r.db.Save(&user).Error; err != nil {
		return nil, errors.New("failed to update profile")
	}

	return &user, nil
}

// ChangePassword mengubah password user
func (r *ProfileRepository) ChangeUserPassword(userID uint, currentPassword, newPassword string) error {
	var user models.User

	// Cek apakah user ada
	if err := r.db.First(&user, userID).Error; err != nil {
		return errors.New("user not found")
	}

	// Verifikasi current password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(currentPassword)); err != nil {
		return errors.New("current password is incorrect")
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash new password")
	}

	// Update password
	user.Password = string(hashedPassword)

	if err := r.db.Save(&user).Error; err != nil {
		return errors.New("failed to change password")
	}

	return nil
}
