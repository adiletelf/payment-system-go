package model

type Admin struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `gorm:"size:255;not null; unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

func (a *Admin) PrepareGive() {
	a.Password = ""
}

type AdminRepo interface {
	Save(*Admin) (*Admin, error)	
	LoginCheck(username, password string) (string, error)
	GetAdminById(uid uint) (Admin, error)
}