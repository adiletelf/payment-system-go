package api

import (
	"fmt"
	"os"

	"github.com/adiletelf/payment-system-go/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDB() error {
	filename := "test.db"
	deleteFileIfExists(filename)
	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Transaction{})

	populateDB(db)

	DB = db
	return nil
}

func populateDB(db *gorm.DB) {
	first := models.User{Email: "first@gmail.com"}
	second := models.User{Email: "second@gmail.com"}
	third := models.User{Email: "third@gmail.com"}
	db.Create([]*models.User{&first, &second, &third})

	transactions := []models.Transaction{
		models.NewTransaction(first, 100, models.Ruble),
		models.NewTransaction(first, 125, models.Ruble),
		models.NewTransaction(first, 126, models.Ruble),
		models.NewTransaction(second, 251, models.Dollar),
		models.NewTransaction(second, 252, models.Dollar),
		models.NewTransaction(third, 300, models.Euro),
	}

	db.Create(&transactions)
}

func deleteFileIfExists(filename string) {
	if _, err := os.Stat(filename); err == nil {
		os.Remove(filename)
	}
}
