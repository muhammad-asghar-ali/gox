package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/muhammad-asghar-ali/gox/fintech/internal/helpers"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/services/users"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/types"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body := helpers.ReadBody(r)

	req := &types.LoginRequest{}
	err := json.Unmarshal(body, req)
	helpers.HandleError(err)

	login, err := users.Login(req)
	helpers.HandleError(err)

	json.NewEncoder(w).Encode(login)
}

func Register(w http.ResponseWriter, r *http.Request) {
	body := helpers.ReadBody(r)

	req := &types.RegisterRequest{}
	err := json.Unmarshal(body, req)
	helpers.HandleError(err)

	register, err := users.Register(req.Username, req.Email, req.Password)
	helpers.HandleError(err)

	json.NewEncoder(w).Encode(register)
}
