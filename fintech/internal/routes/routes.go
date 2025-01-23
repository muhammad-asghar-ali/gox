package routes

import (
	"github.com/gorilla/mux"

	"github.com/muhammad-asghar-ali/go/fintech/internal/handlers"
)

func StartApi() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/login", handlers.Login).Methods("POST")
	router.HandleFunc("/register", handlers.Register).Methods("POST")

	return router
}
