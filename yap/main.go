package main

import (
	"fmt"
	"net/http"

	"github.com/muhammad-asghar-ali/go/yap/ws"
)

func main() {
	fmt.Println("Starting application...")

	manager := &ws.ClientManager{
		Broadcast:  make(chan []byte),
		Register:   make(chan *ws.Client),
		Unregister: make(chan *ws.Client),
		Clients:    make(map[*ws.Client]bool),
	}

	go manager.Start()
	http.HandleFunc("/ws", ws.Ws)
	http.ListenAndServe(":12345", nil)
}
