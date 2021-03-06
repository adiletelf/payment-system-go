package repository

import (
	"github.com/adiletelf/payment-system-go/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransactionRepoImpl struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) *TransactionRepoImpl {
	return &TransactionRepoImpl{
		db: db,
	}
}

func (tr *TransactionRepoImpl) Find(userId uint, email string) ([]model.Transaction, error) {
	var transactions []model.Transaction

	// query ignores parameter if zero field (0, "", false)
	if err := tr.db.Where(&model.Transaction{UserID: uint(userId), Email: email}).Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}

func (tr *TransactionRepoImpl) FindById(id uint) (*model.Transaction, error) {
	var t model.Transaction

	if err := tr.db.First(&t, id).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (tr *TransactionRepoImpl) UpdateStatusOrCreate(t *model.Transaction) error {
	err := tr.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"status"}),
	}).Create(t).Error

	if err != nil {
		return err
	}
	return nil
}
