package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/muhammad-asghar-ali/go/sysdesigns/shortly/internal/config"
	"github.com/muhammad-asghar-ali/go/sysdesigns/shortly/internal/routes"
)

func main() {
	port := config.GetConfig().GetPort()

	fmt.Printf("Starting server on port %v\n", port)

	r := httprouter.New()
	routes.RegisterRoutes(r)

	http.ListenAndServe(port, r)
}
