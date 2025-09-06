package models

import "time"

type Product struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	Name          string    `json:"name" gorm:"not null"`
	Description   string    `json:"description"`
	PurchasePrice float64   `json:"purchase_price" gorm:"not null;type:decimal(10,2)"`
	SellingPrice  float64   `json:"selling_price" gorm:"not null;type:decimal(10,2)"`
	Stock         int       `json:"stock" gorm:"not null;default:0"`
	CategoryID    uint      `json:"category_id" gorm:"not null"`
	Category      Category  `json:"category" gorm:"foreignKey:CategoryID"`
	ImageURL      string    `json:"image_url"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (Product) TableName() string {
	return "products"
}
