package transactions

import (
	"go.temporal.io/sdk/workflow"

	"github.com/muhammad-asghar-ali/gox/fintech/durable/client"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/models"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/services/users"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/types"
)

type (
	Workflow struct{}
)

func (w *Workflow) Transaction(ctx workflow.Context, req *types.TransactionReq, user_id string) (*models.Account, error) {
	options := client.ActivityOptions()

	ctx = workflow.WithActivityOptions(ctx, options)

	act := &Activities{}
	toAccount := &models.Account{}
	fromAccount := &models.Account{}

	uact := users.Activities{}
	user := &models.User{}
	err := workflow.ExecuteActivity(ctx, uact.GetUserByID, user_id).Get(ctx, user)
	if err != nil {
		return nil, err
	}

	err = workflow.ExecuteActivity(ctx, act.GetAccount, req.From).Get(ctx, fromAccount)
	if err != nil {
		return nil, err
	}

	err = workflow.ExecuteActivity(ctx, act.GetAccount, req.To).Get(ctx, toAccount)
	if err != nil {
		return nil, err
	}

	err = workflow.ExecuteActivity(ctx, act.MatchAccount, fromAccount.UserID, user_id).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	err = workflow.ExecuteActivity(ctx, act.CheckBalance, int(fromAccount.Balance), req.Amount).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	err = workflow.ExecuteActivity(ctx, act.Tranfer, fromAccount, toAccount, req).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	err = workflow.ExecuteActivity(ctx, act.CreateTransaction, req).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	return fromAccount, nil
}
