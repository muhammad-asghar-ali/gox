package models

import (
	"authis/internal/database"
	"errors"

	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		Username string `gorm:"unique"`
		Password string
	}
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}

func (u *User) Create() error {
	return database.GetDB().Create(u).Error
}

func (u *User) FindByUsername(username string) error {
	database.GetDB().Where("username = ?", username).First(&u)

	if u.ID == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (u *User) GetUser(user_id uint) {
	database.GetDB().First(u, user_id)
}
