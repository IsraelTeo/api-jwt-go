package model

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID       uint64 `json:"id" gorm:"primary_key;autoIncrement"`
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"size:100;unique;not null"`
	Status   bool   `json:"status" gorm:"defaut:true"`
	Password string `json:"-" gorm:"defaut:true"`
	RoleID   uint64 `json:"role_id"`
	Role     Role   `json:"role"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(passwordHashed string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHashed), []byte(password))
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	passwordHashed, err := Hash(u.Password)
	if err != nil {
		return err
	}

	u.Password = string(passwordHashed)
	return nil
}
func (u *User) Prepare() {
	u.ID = 0
	u.Name = html.EscapeString(strings.ToUpper(strings.TrimSpace(u.Name)))
	u.Name = html.EscapeString(strings.TrimSpace(u.Email))
}
