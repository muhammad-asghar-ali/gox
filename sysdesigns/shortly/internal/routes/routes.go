package routes

import (
	"github.com/julienschmidt/httprouter"

	"github.com/muhammad-asghar-ali/go/sysdesigns/shortly/internal/handlers"
)

func RegisterRoutes(router *httprouter.Router) {
	h := handlers.NewHealthHandler()
	us := handlers.NewUrlStoreHandler()

	router.GET("/api/v1/health", h.Health)
	router.POST("/api/v1/shorten", us.Shorten)
	router.GET("/api/v1/redirect/:code", us.Redirect)
}
