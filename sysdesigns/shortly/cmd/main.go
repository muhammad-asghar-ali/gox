package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"

	"github.com/muhammad-asghar-ali/go/sysdesigns/shortly/internal/config"
	"github.com/muhammad-asghar-ali/go/sysdesigns/shortly/internal/routes"
)

func main() {
	config.NewLogger()
	port := config.GetConfig().GetPort()

	r := httprouter.New()
	routes.RegisterRoutes(r)

	log.Printf("Starting server on port %v\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
