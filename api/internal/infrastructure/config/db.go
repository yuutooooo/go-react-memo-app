package config

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	if os.Getenv("ENV") != "docker" {
		dbHost = "localhost"
	}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSLMODE")
	if dbHost == "" || dbUser == "" || dbPassword == "" || dbName == "" || dbPort == "" || dbSSLMode == "" {
		return nil, errors.New("環境変数が取得できませんでした")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", dbHost, dbUser, dbPassword, dbName, dbPort, dbSSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// UUID拡張機能を有効化
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error; err != nil {
		return nil, fmt.Errorf("failed to create uuid-ossp extension: %w", err)
	}

	println("Connected to database")
	return db, nil
}
