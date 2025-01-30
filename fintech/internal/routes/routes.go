package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/muhammad-asghar-ali/go/fintech/internal/handlers"
	"github.com/muhammad-asghar-ali/go/fintech/internal/middlewares"
)

func StartApi() *mux.Router {
	router := mux.NewRouter()
	router.Use(middlewares.PanicHandler)

	router.HandleFunc("/login", handlers.Login).Methods("POST")
	router.HandleFunc("/register", handlers.Register).Methods("POST")

	router.Handle("/me", middlewares.VerifyAuthorization(http.HandlerFunc(handlers.GetUser))).Methods("GET")
	return router
}
