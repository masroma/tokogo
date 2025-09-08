package repositories

import (
	"tokogo/models"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{
		db: db,
	}
}

func (r *CartRepository) Create(cart *models.Cart) error {
	return r.db.Create(cart).Error
}

func (r *CartRepository) GetByUserID(userID uint) ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Preload("Product").Preload("User").Where("user_id = ?", userID).Find(&carts).Error
	return carts, err
}

func (r *CartRepository) GetByUserIDAndProductID(userID, productID uint) (*models.Cart, error) {
	var cart models.Cart
	err := r.db.Preload("Product").Preload("User").Where("user_id = ? AND product_id = ?", userID, productID).First(&cart).Error
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func (r *CartRepository) Update(cart *models.Cart) error {
	return r.db.Save(cart).Error
}

func (r *CartRepository) Delete(cartID uint) error {
	return r.db.Delete(&models.Cart{}, cartID).Error
}

func (r *CartRepository) DeleteByUserIDAndProductID(userID, productID uint) error {
	return r.db.Where("user_id = ? AND product_id = ?", userID, productID).Delete(&models.Cart{}).Error
}

func (r *CartRepository) ClearCart(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&models.Cart{}).Error
}

func (r *CartRepository) GetCartItemCount(userID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Cart{}).Where("user_id = ?", userID).Count(&count).Error
	return count, err
}
