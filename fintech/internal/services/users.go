package services

import (
	"github.com/muhammad-asghar-ali/go/fintech/internal/helpers"
	"github.com/muhammad-asghar-ali/go/fintech/internal/models"
)

type (
	LoginResponse struct {
		Message string               `json:"message"`
		Token   *string              `json:"token"`
		Data    *models.ResponseUser `json:"data"`
	}
)

func Login(username string, pass string) (*LoginResponse, error) {
	user := &models.User{}
	if err := user.CheckUser(username); err != nil {
		return &LoginResponse{
			Message: err.Error(),
		}, err
	}

	if err := helpers.ComparePassword(pass, user.Password); err != nil {
		return &LoginResponse{
			Message: err.Error(),
		}, err
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

	return &LoginResponse{
		Message: "Ok",
		Token:   token,
		Data:    res,
	}, nil
}
