package transactions

import (
	"context"
	"fmt"
	"time"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"

	"github.com/muhammad-asghar-ali/gox/fintech/internal/database"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/durable"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/models"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/types"
)

func Transaction(userID string, req *types.TransactionReq) (*models.Account, error) {
	db := database.GetDatabase()
	user := &models.User{}
	if err := user.GetUserByID(db, userID); err != nil {
		return nil, err
	}

	options := client.StartWorkflowOptions{
		ID:        "transaction-workflow-001",
		TaskQueue: durable.MoneyTransferTaskQueueName,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second * 2,
			BackoffCoefficient: 2.0,
			MaximumAttempts:    5,
		},
	}

	w := Workflow{}
	run, err := durable.Set().Client().
		ExecuteWorkflow(context.Background(), options, w.Transaction, req, user.ID)
	if err != nil {
		return nil, err
	}

	fmt.Println("Started Workflow:", "ID:", run.GetID(), "RunID:", run.GetRunID())

	result := &models.Account{}
	err = run.Get(context.Background(), &result)
	if err != nil {
		return nil, err
	}

	fmt.Println("result", result)

	return result, nil
}

func GetMyTransactions(userID string) []types.TransactionResponse {
	db := database.GetDatabase()
	a := models.Account{}

	accounts, _ := a.GetAccountsByUserID(db, userID)

	transactions := []types.TransactionResponse{}
	t := models.Transaction{}
	for i := 0; i < len(accounts); i++ {
		accTransactions := t.GetTransactionsByAccount(db, accounts[i].ID)
		transactions = append(transactions, accTransactions...)
	}

	return transactions
}
