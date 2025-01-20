package models

import (
	"github.com/jinzhu/gorm"
)

type (
	User struct {
		gorm.Model
		Username string
		Email    string
		Password string
	}

	Account struct {
		gorm.Model
		Type    string
		Name    string
		Balance uint
		UserID  uint
	}

	ResponseAccount struct {
		ID      uint
		Name    string
		Balance int
	}

	ResponseUser struct {
		ID       uint
		Username string
		Email    string
		Accounts []*ResponseAccount
	}
)
