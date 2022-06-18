package model

type Admin struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `gorm:"size:255;not null; unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

type AdminRepo interface {
	Save(*Admin) (*Admin, error)	
	LoginCheck(username, password string) (string, error)
}
