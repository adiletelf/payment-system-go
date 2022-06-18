package repository

import (
	"html"
	"strings"

	"github.com/adiletelf/payment-system-go/internal/model"
	"github.com/adiletelf/payment-system-go/internal/token"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AdminRepoImpl struct {
	db *gorm.DB
}

func NewAdminRepo(db *gorm.DB) *AdminRepoImpl {
	return &AdminRepoImpl{
		db: db,
	}
}

func (ar *AdminRepoImpl) Save(a *model.Admin) (*model.Admin, error) {
	ar.BeforeSave(a)
	if err := ar.db.Create(a).Error; err != nil {
		return &model.Admin{}, err
	}

	return a, nil
}

func (ar *AdminRepoImpl) BeforeSave(a *model.Admin) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	a.Password = string(hashedPassword)
	a.Username = html.EscapeString(strings.TrimSpace(a.Username))

	return nil
}

func (ar *AdminRepoImpl) LoginCheck(username, password string) (string, error) {
	var err error
	a := model.Admin{}

	err = ar.db.Model(model.Admin{}).Where("username = ?", username).Take(&a).Error
	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, a.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(a.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
