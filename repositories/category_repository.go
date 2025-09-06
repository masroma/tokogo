package repositories

import (
	"errors"
	"tokogo/config"
	"tokogo/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

// NewCategoryRepository membuat instance baru CategoryRepository
func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{
		db: config.DB,
	}
}

// CreateCategory menyimpan category baru ke database
func (r *CategoryRepository) CreateCategory(category *models.Category) error {
	return r.db.Create(category).Error
}

// GetCategoryByID mengambil category berdasarkan ID
func (r *CategoryRepository) GetCategoryByID(id uint) (*models.Category, error) {
	var category models.Category
	err := r.db.Where("id = ?", id).First(&category).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}

	return &category, nil
}

// GetCategoryBySlug mengambil category berdasarkan slug
func (r *CategoryRepository) GetCategoryBySlug(slug string) (*models.Category, error) {
	var category models.Category
	err := r.db.Where("slug = ?", slug).First(&category).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}

	return &category, nil
}

// GetAllCategories mengambil semua categories dengan pagination
func (r *CategoryRepository) GetAllCategories(page, limit int) ([]models.Category, int64, error) {
	var categories []models.Category
	var total int64

	// Hitung total records
	if err := r.db.Model(&models.Category{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Ambil data dengan pagination
	offset := (page - 1) * limit
	err := r.db.Offset(offset).Limit(limit).Order("created_at DESC").Find(&categories).Error

	return categories, total, err
}

// UpdateCategory mengupdate category berdasarkan ID
func (r *CategoryRepository) UpdateCategory(id uint, category *models.Category) error {
	return r.db.Where("id = ?", id).Updates(category).Error
}

// DeleteCategory menghapus category berdasarkan ID (soft delete)
func (r *CategoryRepository) DeleteCategory(id uint) error {
	return r.db.Where("id = ?", id).Delete(&models.Category{}).Error
}

// CheckCategoryExists mengecek apakah category dengan nama tertentu sudah ada
func (r *CategoryRepository) CheckCategoryExists(name string, excludeID uint) (bool, error) {
	var count int64
	query := r.db.Model(&models.Category{}).Where("name = ?", name)

	if excludeID > 0 {
		query = query.Where("id != ?", excludeID)
	}

	err := query.Count(&count).Error
	return count > 0, err
}
