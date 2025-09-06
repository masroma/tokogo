package models

import (
    "time"
    "strings"
    "gorm.io/gorm"
)

type Category struct {
    ID        uint           `gorm:"primaryKey;column:id;type:BIGINT UNSIGNED AUTO_INCREMENT" json:"id"`
    Name      string         `gorm:"column:name;type:VARCHAR(255);not null" json:"name"`
    Slug      string         `gorm:"column:slug;type:VARCHAR(255);uniqueIndex;not null" json:"slug"`
    CreatedAt time.Time      `gorm:"column:created_at;type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt time.Time      `gorm:"column:updated_at;type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"-"`
}

// TableName mengembalikan nama tabel untuk model Category
func (Category) TableName() string {
    return "categories"
}

// GenerateSlug menghasilkan slug dari nama kategori
func (c *Category) GenerateSlug() {
    slug := strings.ToLower(c.Name)
    slug = strings.ReplaceAll(slug, " ", "-")
    slug = strings.ReplaceAll(slug, "_", "-")
    c.Slug = slug
}

// BeforeCreate hook untuk generate slug sebelum create
func (c *Category) BeforeCreate(tx *gorm.DB) error {
    c.GenerateSlug()
    return nil
}

// BeforeUpdate hook untuk generate slug sebelum update
func (c *Category) BeforeUpdate(tx *gorm.DB) error {
    c.GenerateSlug()
    return nil
}