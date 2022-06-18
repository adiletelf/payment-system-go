package util

import (
	"fmt"
	"log"
	"os"

	"github.com/adiletelf/payment-system-go/internal/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDB() (*gorm.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	filename := os.Getenv("DB_NAME")
	deleteFileIfExists(filename)
	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Transaction{})
	db.AutoMigrate(&model.Admin{})

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
