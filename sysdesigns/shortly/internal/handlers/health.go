package handlers

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type (
	HealthHanlder struct{}
)

func NewHealthHandler() *HealthHanlder {
	return &HealthHanlder{}
}

func (hh *HealthHanlder) Health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"alive": true}`)
}
