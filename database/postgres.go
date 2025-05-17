package database

import (
	"fmt"
	"log"

	"backend101/config"
	"backend101/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectPostgres() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Get("DB_HOST"),
		config.Get("DB_USER"),
		config.Get("DB_PASSWORD"),
		config.Get("DB_NAME"),
		config.Get("DB_PORT"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database: ", err)
	}

	log.Println("✅ Connected to PostgreSQL database!")

	err = DB.AutoMigrate(&models.User{}, &models.Transaction{})
	if err != nil {
		log.Fatal("❌ Failed to migrate models: ", err)
	}
	log.Println("📦 User table migrated!")
}
