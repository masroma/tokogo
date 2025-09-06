package repositories

import (
	"tokogo/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepository) GetAll(page, limit int) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64

	offset := (page - 1) * limit

	// Count total records
	if err := r.db.Model(&models.Product{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get products with pagination
	err := r.db.Preload("Category").
		Offset(offset).
		Limit(limit).
		Find(&products).Error

	return products, total, err
}

func (r *ProductRepository) GetByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.Preload("Category").First(&product, id).Error
	return &product, err
}

func (r *ProductRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *ProductRepository) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}

func (r *ProductRepository) GetByCategoryID(categoryID uint, page, limit int) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64

	offset := (page - 1) * limit

	// Count total records
	if err := r.db.Model(&models.Product{}).Where("category_id = ?", categoryID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get products with pagination
	err := r.db.Preload("Category").
		Where("category_id = ?", categoryID).
		Offset(offset).
		Limit(limit).
		Find(&products).Error

	return products, total, err
}
