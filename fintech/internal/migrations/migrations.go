package migrations

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/muhammad-asghar-ali/go/fintech/internal/helpers"
	"github.com/muhammad-asghar-ali/go/fintech/internal/models"
)

// TODO - to singleton
func connect() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=user dbname=dbname password=password sslmode=disable")
	helpers.HandleError(err)

	return db
}

// test account
// TODO - figure out some way
func createAccounts() {
	db := connect()

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

func Migrate() {
	db := connect()
	db.AutoMigrate(&models.User{}, &models.Account{})
	defer db.Close()

	createAccounts()
}
