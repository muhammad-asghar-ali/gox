package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/muhammad-asghar-ali/gox/fintech/internal/database"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/helpers"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/models"
)

func createAccounts() {
	db := database.GetDatabase()

	users := [2]models.User{
		{Username: "John", Email: "john@john.com"},
		{Username: "Michael", Email: "michael@michael.com"},
	}

	for i := 0; i < len(users); i++ {
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := &models.User{
			Username: users[i].Username,
			Email:    users[i].Email,
			Password: generatedPassword,
		}
		db.Create(user)

		account := &models.Account{
			Type:    "Daily Account",
			Name:    string(users[i].Username + "'s" + " account"),
			Balance: uint(10000 * int(i+1)),
			UserID:  user.ID,
		}
		db.Create(account)
	}

	defer db.Close()
}

func main() {
	db := database.GetDatabase()
	db.AutoMigrate(&models.User{}, &models.Account{})
	db.AutoMigrate(&models.Transaction{})
	defer db.Close()

	createAccounts()
}
