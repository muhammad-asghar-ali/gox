package users

import (
	"context"

	"github.com/muhammad-asghar-ali/gox/fintech/internal/database"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/helpers"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/models"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/types"
)

type (
	Activities struct{}
)

var db = database.GetDatabase()

func (a *Activities) CheckUser(ctx context.Context, req *types.LoginRequest) (*models.User, error) {
	user := &models.User{}

	if err := user.CheckUser(db, req.Username); err != nil {
		return nil, err
	}

	if err := helpers.ComparePassword(req.Password, user.Password); err != nil {
		return nil, err
	}

	return user, nil
}

func (a *Activities) GenerateToken(ctx context.Context, uid uint) (string, error) {
	token, err := helpers.GenerateToken(uid)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a *Activities) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	user := &models.User{}
	if err := user.GetUserByID(db, id); err != nil {
		return nil, err
	}

	return user, nil
}

func (a *Activities) CreateUser(ctx context.Context, req *types.RegisterRequest) (*models.User, error) {
	generatedPassword := helpers.HashAndSalt([]byte(req.Password))

	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: generatedPassword,
	}
	if err := user.Create(db); err != nil {
		return nil, err
	}

	return user, nil
}

func (a *Activities) CreateUserAccount(ctx context.Context, uid uint, username string) (*types.AccountResponse, error) {
	account := &models.Account{
		Type:    "Daily Account",
		Name:    string(username + "'s" + " account"),
		Balance: 0,
		UserID:  uid,
	}
	if err := account.Create(db); err != nil {
		return nil, err
	}

	ra := &types.AccountResponse{
		Type:    account.Type,
		Name:    account.Name,
		Balance: account.Balance,
	}

	return ra, nil
}
