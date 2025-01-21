package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/muhammad-asghar-ali/go/fintech/internal/routes"
)

func main() {
	r := routes.StartApi()

	fmt.Println("App is running on port :8888")
	log.Fatal(http.ListenAndServe(":8888", r))
}
