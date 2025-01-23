package services

import (
	"github.com/muhammad-asghar-ali/go/fintech/internal/helpers"
	"github.com/muhammad-asghar-ali/go/fintech/internal/models"
	"github.com/muhammad-asghar-ali/go/fintech/internal/types"
)

func Login(username string, pass string) (*types.LoginResponse, error) {
	if err := helpers.Validation([]helpers.Validate{
		{Value: username, Valid: "username"},
		{Value: pass, Valid: "password"},
	}); err != nil {
		return nil, err
	}

	db := helpers.GetDatabase()
	user := &models.User{}
	if err := user.CheckUser(db, username); err != nil {
		return &types.LoginResponse{
			Message: err.Error(),
		}, err
	}

	if err := helpers.ComparePassword(pass, user.Password); err != nil {
		return &types.LoginResponse{
			Message: err.Error(),
		}, err
	}

	a := models.Account{}
	accounts := a.UserAccounts(db, user.ID)

	res := &types.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Accounts: accounts,
	}

	token, err := helpers.GenerateToken(user.ID)
	helpers.HandleError(err)

	return &types.LoginResponse{
		Message: "Ok",
		Token:   token,
		Data:    res,
	}, nil
}

func Register(username, email, pass string) (*types.RegisterResponse, error) {
	if err := helpers.Validation([]helpers.Validate{
		{Value: username, Valid: "username"},
		{Value: pass, Valid: "password"},
		{Value: email, Valid: "email"},
	}); err != nil {
		return nil, err
	}

	db := helpers.GetDatabase()
	generatedPassword := helpers.HashAndSalt([]byte(pass))

	// TODO - apply transections
	user := &models.User{
		Username: username,
		Email:    email,
		Password: generatedPassword,
	}
	if err := user.Create(db); err != nil {
		return nil, err
	}

	account := &models.Account{
		Type:    "Daily Account",
		Name:    string(username + "'s" + " account"),
		Balance: 0,
		UserID:  user.ID,
	}
	if err := account.Create(db); err != nil {
		return nil, err
	}

	ra := &types.AccountResponse{
		Type:    account.Type,
		Name:    account.Name,
		Balance: account.Balance,
	}

	res := &types.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Accounts: []*types.AccountResponse{ra},
	}

	return &types.RegisterResponse{
		Message: "Ok",
		Data:    res,
	}, nil

}
