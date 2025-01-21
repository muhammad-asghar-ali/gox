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
		Username string
		Password string
	}

	ErrResponse struct {
		Message string
	}
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	helpers.HandleError(err)

	req := &LoginRequest{}
	err = json.Unmarshal(body, req)
	helpers.HandleError(err)

	login := services.Login(req.Username, req.Password)

	if login["success"] == true {
		resp := login
		json.NewEncoder(w).Encode(resp)
	} else {

	}
}
