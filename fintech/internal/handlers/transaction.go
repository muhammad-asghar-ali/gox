package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/muhammad-asghar-ali/gox/fintech/internal/helpers"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/services"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/types"
)

func Transaction(w http.ResponseWriter, r *http.Request) {
	body := helpers.ReadBody(r)

	req := &types.TransactionReq{}
	err := json.Unmarshal(body, req)
	helpers.HandleError(err)

	userID, err := helpers.GetUserIdFromCtx(r)
	helpers.HandleError(err)

	res, err := services.Transaction(userID, req)
	helpers.HandleError(err)

	json.NewEncoder(w).Encode(res)
}
