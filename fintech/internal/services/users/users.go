package users

import (
	"context"

	"github.com/muhammad-asghar-ali/gox/fintech/durable/client"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/database"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/helpers"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/models"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/types"
)

func Login(req *types.LoginRequest) (*types.LoginResponse, error) {
	if err := helpers.Validation([]helpers.Validate{
		{Value: req.Username, Valid: "username"},
		{Value: req.Password, Valid: "password"},
	}); err != nil {
		return nil, err
	}

	options := client.StartWorkflowOptions(client.DQueue)

	w := &Workflow{}
	run, err := client.New().Client().
		ExecuteWorkflow(context.Background(), options, w.Login, req)
	if err != nil {
		return nil, err
	}

	var token string
	err = run.Get(context.Background(), &token)
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		Message: "Ok",
		Token:   token,
	}, nil
}

func Register(req *types.RegisterRequest) (*types.RegisterResponse, error) {
	if err := helpers.Validation([]helpers.Validate{
		{Value: req.Username, Valid: "username"},
		{Value: req.Password, Valid: "password"},
		{Value: req.Email, Valid: "email"},
	}); err != nil {
		return nil, err
	}

	options := client.StartWorkflowOptions(client.DQueue)

	w := &Workflow{}
	res := &types.RegisterResponse{}
	run, err := client.New().Client().
		ExecuteWorkflow(context.Background(), options, w.Register, req)
	if err != nil {
		return nil, err
	}

	err = run.Get(context.Background(), &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetUser(id string) (*types.UserResponse, error) {
	db := database.GetDatabase()

	user := &models.User{}
	if err := user.GetUserByID(db, id); err != nil {
		return nil, err
	}

	a := models.Account{}
	accounts := a.UserAccounts(db, user.ID)

	res := &types.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Accounts: accounts,
	}

	return res, nil
}
