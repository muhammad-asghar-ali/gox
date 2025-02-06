package main

import (
	"log"
	"net/http"

	"github.com/muhammad-asghar-ali/gox/fintech/durable/worker"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/routes"
)

func main() {
	go func() {
		if err := worker.Run(); err != nil {
			log.Fatalln("Unable to start Temporal worker:", err)
		}
	}()

	r := routes.StartApi()

	log.Println("App is running on port :8888")
	log.Fatal(http.ListenAndServe(":8888", r))
}
