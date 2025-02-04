package main

import (
	"log"
	"net/http"

	"github.com/muhammad-asghar-ali/gox/fintech/internal/durable"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/routes"
)

func main() {
	go func() {
		if err := durable.Run(); err != nil {
			log.Fatalln("Unable to start Temporal worker:", err)
		}
	}()

	r := routes.StartApi()

	log.Println("App is running on port :8888")
	log.Fatal(http.ListenAndServe(":8888", r))
}
