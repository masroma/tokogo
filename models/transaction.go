package models

import "time"

// Transaction represents the transaction model
type Transaction struct {
	ID                 uint                `json:"id" gorm:"primaryKey"`
	UserID             uint                `json:"user_id" gorm:"not null"`
	User               User                `json:"user" gorm:"foreignKey:UserID"`
	Status             string              `json:"status" gorm:"type:enum('pending','paid','failed','expired');default:'pending'"`
	TotalAmount        float64             `json:"total_amount" gorm:"type:decimal(15,2);not null"`
	ShippingAddress    string              `json:"shipping_address" gorm:"type:text;not null"`
	PaymentMethod      string              `json:"payment_method" gorm:"type:varchar(50);not null"`
	PaymentURL         string              `json:"payment_url" gorm:"type:varchar(500)"`
	PaymentProof       string              `json:"payment_proof" gorm:"type:varchar(500)"`
	Notes              string              `json:"notes" gorm:"type:text"`
	TransactionDetails []TransactionDetail `json:"transaction_details" gorm:"foreignKey:TransactionID"`
	CreatedAt          time.Time           `json:"created_at"`
	UpdatedAt          time.Time           `json:"updated_at"`
}

// TransactionDetail represents the transaction detail model
type TransactionDetail struct {
	ID            uint        `json:"id" gorm:"primaryKey"`
	TransactionID uint        `json:"transaction_id" gorm:"not null"`
	Transaction   Transaction `json:"transaction" gorm:"foreignKey:TransactionID"`
	ProductID     uint        `json:"product_id" gorm:"not null"`
	Product       Product     `json:"product" gorm:"foreignKey:ProductID"`
	Quantity      int         `json:"quantity" gorm:"not null"`
	Price         float64     `json:"price" gorm:"type:decimal(15,2);not null"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}

// TableName returns the table name for Transaction
func (Transaction) TableName() string {
	return "transactions"
}

// TableName returns the table name for TransactionDetail
func (TransactionDetail) TableName() string {
	return "transaction_details"
}
