package config

import (
	"fmt"
	"time"

	"tokogo/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB() {
	// Ambil konfigurasi database dari environment variables
	dbUser := GetEnv("DB_USER", "root")
	dbPassword := GetEnv("DB_PASSWORD", "")
	dbHost := GetEnv("DB_HOST", "localhost")
	dbPort := GetEnv("DB_PORT", "3306")
	dbName := GetEnv("DB_NAME", "tokogo")

	// Buat DSN (Data Source Name) untuk koneksi MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	var err error

	// Buka koneksi database menggunakan GORM
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // Gunakan nama table singular
		},
	})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	// Konfigurasi connection pool
	sqlDB, err := DB.DB()
	if err != nil {
		panic("Failed to get database instance!")
	}

	// Set connection pool settings
	sqlDB.SetMaxOpenConns(10)           // Maksimal 10 koneksi terbuka
	sqlDB.SetMaxIdleConns(5)            // Maksimal 5 koneksi idle
	sqlDB.SetConnMaxLifetime(time.Hour) // Maksimal 1 jam lifetime

	// Auto migrate model User (akan kita buat di bab selanjutnya)
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		panic(fmt.Sprintf("AutoMigrate failed: %v", err))
	}

}
