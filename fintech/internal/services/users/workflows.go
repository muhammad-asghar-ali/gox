package users

import (
	"time"

	"github.com/muhammad-asghar-ali/gox/fintech/internal/models"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/types"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

type (
	Workflow struct{}
)

func (w *Workflow) Login(ctx workflow.Context, req *types.LoginRequest) (string, error) {
	retrypolicy := &temporal.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2.0,
		MaximumInterval:    100 * time.Second,
		MaximumAttempts:    500,
	}

	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
		RetryPolicy:         retrypolicy,
	}

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
