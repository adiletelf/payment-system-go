package api

import (
	"fmt"

	"github.com/adiletelf/payment-system-go/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Transaction{})

	populateDB(db)

	return db, nil
}

func populateDB(db *gorm.DB) {
	first := models.User{Email: "first@gmail.com"}
	second := models.User{Email: "second@gmail.com"}
	db.Create(&[]models.User{first, second})

	transactions := []models.Transaction{
		models.NewTransaction(first, 100, models.Ruble),
		models.NewTransaction(first, 125, models.Ruble),
		models.NewTransaction(second, 150, models.Ruble),
	}

	db.Create(&transactions)
}
