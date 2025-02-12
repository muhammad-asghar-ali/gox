package users

import (
	"github.com/muhammad-asghar-ali/gox/fintech/durable/client"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/models"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/types"
	"go.temporal.io/sdk/workflow"
)

type (
	Workflow struct{}
)

func (w *Workflow) Login(ctx workflow.Context, req *types.LoginRequest) (string, error) {
	options := client.ActivityOptions()

	ctx = workflow.WithActivityOptions(ctx, options)
	act := &Activities{}
	user := &models.User{}
	var token string

	err := workflow.ExecuteActivity(ctx, act.CheckUser, req).Get(ctx, user)
	if err != nil {
		return "", err
	}

	err = workflow.ExecuteActivity(ctx, act.GenerateToken, user.ID).Get(ctx, &token)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (w *Workflow) Register(ctx workflow.Context, req *types.RegisterRequest) (*types.RegisterResponse, error) {
	options := client.ActivityOptions()

	ctx = workflow.WithActivityOptions(ctx, options)
	act := &Activities{}
	user := &models.User{}
	ra := &types.AccountResponse{}

	err := workflow.ExecuteActivity(ctx, act.CreateUser, req).Get(ctx, user)
	if err != nil {
		return nil, err
	}

	err = workflow.ExecuteActivity(ctx, act.CreateUserAccount, user.ID, user.Username).Get(ctx, ra)
	if err != nil {
		return nil, err
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
