package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/muhammad-asghar-ali/go/fintech/internal/helpers"
	"github.com/muhammad-asghar-ali/go/fintech/internal/services"
)

type (
	LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	ErrResponse struct {
		Message string `json:"message"`
	}
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	helpers.HandleError(err)

	req := &LoginRequest{}
	err = json.Unmarshal(body, req)
	helpers.HandleError(err)

	login, err := services.Login(req.Username, req.Password)
	helpers.HandleError(err)

	json.NewEncoder(w).Encode(login)
}
