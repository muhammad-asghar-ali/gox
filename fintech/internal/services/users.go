package services

import (
	"github.com/muhammad-asghar-ali/go/fintech/internal/helpers"
	"github.com/muhammad-asghar-ali/go/fintech/internal/models"
)

// TODO: need a proper way to handle response
func Login(username string, pass string) map[string]any {
	user := &models.User{}
	if err := user.CheckUser(username); err != nil {
		return map[string]any{"message": err.Error(), "success": false}
	}

	if err := helpers.ComparePassword(pass, user.Password); err != nil {
		return map[string]any{"message": "Wrong password", "success": false}
	}

	account := models.ResponseAccount{}
	accounts := account.UserAccounts(user.ID)

	res := &models.ResponseUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Accounts: accounts,
	}

	token, err := helpers.GenerateToken(user.ID)
	helpers.HandleError(err)

	var response = map[string]any{"message": "all is fine", "success": true}
	response["jwt"] = token
	response["data"] = res

	return response
}
