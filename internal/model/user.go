package model

type User struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Email string `json:"email"`
}
