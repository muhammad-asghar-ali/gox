package transactions

import (
	"context"
	"errors"

	"github.com/muhammad-asghar-ali/gox/fintech/internal/database"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/models"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/types"
)

type (
	Activities struct{}
)

func (a *Activities) GetAccount(ctx context.Context, id uint) (*models.Account, error) {
	db := database.GetDatabase()
	account := &models.Account{}

	if err := account.GetAccount(db, id); err != nil {
		return nil, err
	}

	return account, nil
}

func (a *Activities) MatchAccount(ctx context.Context, from_user_id, user_id uint) error {
	if from_user_id != user_id {
		return errors.New("you are not owner of the account")
	}

	return nil
}

func (a *Activities) CheckBalance(ctx context.Context, curr, want_to int) error {
	if curr < want_to {
		return errors.New("account balance is too small")
	}

	return nil
}

func (a *Activities) Tranfer(ctx context.Context, fromAccount, toAccount *models.Account, req *types.TransactionReq) error {
	db := database.GetDatabase()

	if err := fromAccount.UpdateAccount(db, req.From, int(fromAccount.Balance)-req.Amount); err != nil {
		return err
	}

	if err := toAccount.UpdateAccount(db, req.To, int(toAccount.Balance)+req.Amount); err != nil {
		return err
	}

	return nil
}

func (a *Activities) CreateTransaction(ctx context.Context, req *types.TransactionReq) error {
	db := database.GetDatabase()

	transaction := &models.Transaction{
		From: req.From, To: req.To, Amount: req.Amount,
	}

	if err := transaction.CreateTransaction(db); err != nil {
		return err
	}

	return nil
}
