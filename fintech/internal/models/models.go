package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/types"
)

type (
	User struct {
		gorm.Model
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	Account struct {
		gorm.Model
		Type    string `json:"type"`
		Name    string `json:"name"`
		Balance uint   `json:"balance"`
		UserID  uint   `json:"user_id"`
	}

	Transaction struct {
		gorm.Model
		From   uint `json:"from"`
		To     uint `json:"to"`
		Amount int  `json:"amount"`
	}
)

func (u *User) CheckUser(db *gorm.DB, username string) error {
	if db.Where("username = ? ", username).First(&u).RecordNotFound() {
		return errors.New("user not found")
	}

	return nil
}

func (u *User) Create(db *gorm.DB) error {
	return db.Create(&u).Error
}

func (a *Account) UserAccounts(db *gorm.DB, userID uint) []*types.AccountResponse {
	accounts := make([]*Account, 0)
	ra := make([]*types.AccountResponse, 0)

	db.Table("accounts").Where("user_id = ?", userID).Find(&accounts)

	if len(accounts) > 0 {
		for _, account := range accounts {
			ra = append(ra, &types.AccountResponse{
				Type:    account.Type,
				Name:    account.Name,
				Balance: account.Balance,
			})
		}
	}

	return ra
}

func (u *User) GetUserByID(db *gorm.DB, id string) error {
	if db.Where("id = ? ", id).First(&u).RecordNotFound() {
		return errors.New("user not found")
	}

	return nil
}

func (a *Account) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a *Account) UpdateAccount(db *gorm.DB, id uint, amount int) error {
	a.Balance = uint(amount)
	db.Save(&a)

	return nil
}

func (a *Account) GetAccount(db *gorm.DB, id uint) error {
	if db.Where("id = ? ", id).First(&a).RecordNotFound() {
		return errors.New("account not found")
	}

	return nil
}

func (t *Transaction) CreateTransaction(db *gorm.DB) error {
	return db.Create(&t).Error
}
