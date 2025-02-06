package transactions

import (
	"context"

	"github.com/muhammad-asghar-ali/gox/fintech/durable/client"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/database"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/models"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/types"
)

func Transaction(userID string, req *types.TransactionReq) (*models.Account, error) {
	options := client.StartWorkflowOptions(client.DQueue)

	w := Workflow{}
	run, err := client.New().Client().
		ExecuteWorkflow(context.Background(), options, w.Transaction, req, userID)
	if err != nil {
		return nil, err
	}

	result := &models.Account{}
	err = run.Get(context.Background(), &result)
	if err != nil {
		return nil, err
	}

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
