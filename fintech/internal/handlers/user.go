package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/muhammad-asghar-ali/gox/fintech/internal/helpers"
	services "github.com/muhammad-asghar-ali/gox/fintech/internal/services/users"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	userID, err := helpers.GetUserIdFromCtx(r)
	helpers.HandleError(err)

	res, err := services.GetUser(userID)
	helpers.HandleError(err)

	json.NewEncoder(w).Encode(res)
}
