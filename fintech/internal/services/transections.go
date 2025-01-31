package services

import (
	"errors"

	"github.com/muhammad-asghar-ali/gox/fintech/internal/helpers"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/models"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/types"
)

func Transaction(userId string, req *types.TransactionReq) (*models.Account, error) {
	db := helpers.GetDatabase()
	user := &models.User{}
	if err := user.GetUserByID(db, userId); err != nil {
		return nil, err
	}

	fromAccount := &models.Account{}
	toAccount := &models.Account{}
	a := &models.Account{}

	if err := fromAccount.GetAccount(db, req.From); err != nil {
		return nil, err
	}
	if err := toAccount.GetAccount(db, req.To); err != nil {
		return nil, err
	}

	if fromAccount.UserID != user.ID {
		return nil, errors.New("you are not owner of the account")
	} else if int(fromAccount.Balance) < req.Amount {
		return nil, errors.New("account balance is too small")
	}

	if err := a.UpdateAccount(db, req.From, int(fromAccount.Balance)-req.Amount); err != nil {
		return nil, err
	}

	if err := a.UpdateAccount(db, req.To, int(toAccount.Balance)+req.Amount); err != nil {
		return nil, err
	}

	transaction := &models.Transaction{
		From: req.From, To: req.To, Amount: req.Amount,
	}

	if err := transaction.CreateTransaction(db); err != nil {
		return nil, err
	}

	return a, nil
}
