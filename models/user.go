package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey;column:id;type:BIGINT UNSIGNED AUTO_INCREMENT" json:"id"`
	Name      string         `gorm:"column:name;type:VARCHAR(255);not null" json:"name"`
	Email     string         `gorm:"column:email;type:VARCHAR(255);uniqueIndex;not null" json:"email"`
	Password  string         `gorm:"column:password;type:VARCHAR(255);not null" json:"-"` // Hidden dari JSON response
	Role      string         `gorm:"column:role;type:ENUM('customer','admin');default:'customer'" json:"role"`
	CreatedAt time.Time      `gorm:"column:created_at;type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"-"`
}

// TableName mengembalikan nama tabel untuk model User
func (User) TableName() string {
	return "user"
}
