package models

import (
	"errors"

	"gorm.io/gorm"

	"github.com/pvfm/enube/api/services"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique" json:"email"`
	Name     string `json:"name"`
	Password string `gorm:"type:varchar(255)" json:"password"`
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if tx.Statement.Changed("Password") {
		hashPassword, err := services.HashPassword(u.Password)

		if err != nil {
			return errors.New("Error Transaction")
		}

		u.Password = hashPassword
	}

	return
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashPassword, err := services.HashPassword(u.Password)

	if err != nil {
		return errors.New("Error Transaction")
	}

	u.Password = hashPassword

	return
}
