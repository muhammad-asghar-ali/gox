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

func (u *User) CheckUser(username string) error {
	db := helpers.ConnectDB()
	defer db.Close()

	if db.Where("username = ? ", username).First(u).RecordNotFound() {
		return errors.New("user not found")
	}

	return nil
}

func (ra *ResponseAccount) UserAccounts(userID uint) []*ResponseAccount {
	db := helpers.ConnectDB()
	defer db.Close()

	accounts := make([]*ResponseAccount, 0)
	db.Table("accounts").Select("id", "name", "balance").Where("user_id = ?", userID).Scan(accounts)
	return accounts
}
