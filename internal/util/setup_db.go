package util

import (
	"fmt"
	"os"

	"github.com/adiletelf/payment-system-go/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDB() (*gorm.DB, error) {
	filename := "test.db"
	deleteFileIfExists(filename)
	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Transaction{})

	populateDB(db)

	return db, nil
}

func populateDB(db *gorm.DB) {
	first := model.User{Email: "first@gmail.com"}
	second := model.User{Email: "second@gmail.com"}
	third := model.User{Email: "third@gmail.com"}
	db.Create([]*model.User{&first, &second, &third})

	transactions := []model.Transaction{
		model.NewTransaction(first, 100, model.Ruble),
		model.NewTransaction(first, 125, model.Ruble),
		model.NewTransaction(first, 126, model.Ruble),
		model.NewTransaction(second, 251, model.Dollar),
		model.NewTransaction(second, 252, model.Dollar),
		model.NewTransaction(third, 300, model.Euro),
	}

	db.Create(&transactions)
}

func deleteFileIfExists(filename string) {
	if _, err := os.Stat(filename); err == nil {
		os.Remove(filename)
	}
}
