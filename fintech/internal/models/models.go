package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/muhammad-asghar-ali/go/fintech/internal/helpers"
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

	ResponseUser struct {
		ID       uint
		Username string
		Email    string
		Accounts []*Account
	}
)

func (u *User) CheckUser(username string) error {
	db := helpers.GetDatabase()
	// defer db.Close()

	if db.Where("username = ? ", username).First(&u).RecordNotFound() {
		return errors.New("user not found")
	}

	return nil
}

func (a *Account) UserAccounts(userID uint) []*Account {
	db := helpers.GetDatabase()
	// defer db.Close()

	accounts := make([]*Account, 0)
	db.Table("accounts").Where("user_id = ?", userID).Find(&accounts)

	return accounts
}
