package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/muhammad-asghar-ali/gox/fintech/internal/helpers"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/services/transactions"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/types"
)

func Transaction(w http.ResponseWriter, r *http.Request) {
	body := helpers.ReadBody(r)

	req := &types.TransactionReq{}
	err := json.Unmarshal(body, req)
	helpers.HandleError(err)

	userID, err := helpers.GetUserIdFromCtx(r)
	helpers.HandleError(err)

	res, err := transactions.Transaction(userID, req)
	helpers.HandleError(err)

	json.NewEncoder(w).Encode(res)
}

func GetMyTransactions(w http.ResponseWriter, r *http.Request) {
	userID, err := helpers.GetUserIdFromCtx(r)
	helpers.HandleError(err)

	transactions := transactions.GetMyTransactions(userID)

	json.NewEncoder(w).Encode(transactions)
}
