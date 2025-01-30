package ws

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

type (
	ClientManager struct {
		Clients    map[*Client]bool
		Broadcast  chan []byte
		Register   chan *Client
		Unregister chan *Client
	}

	Client struct {
		ID     string
		Socket *websocket.Conn
		Send   chan []byte
	}

	Message struct {
		Sender    string `json:"sender,omitempty"`
		Recipient string `json:"recipient,omitempty"`
		Content   string `json:"content,omitempty"`
	}
)

func (manager *ClientManager) send(message []byte, ignore *Client) {
	for conn := range manager.Clients {
		if conn != ignore {
			conn.Send <- message
		}
	}
}

func (manager *ClientManager) Start() {
	for {
		select {
		case conn := <-manager.Register:
			manager.Clients[conn] = true
			jsonMessage, _ := json.Marshal(&Message{Content: "/A new socket has connected."})
			manager.send(jsonMessage, conn)

		case conn := <-manager.Unregister:
			if _, ok := manager.Clients[conn]; ok {
				close(conn.Send)
				delete(manager.Clients, conn)
				jsonMessage, _ := json.Marshal(&Message{Content: "/A socket has disconnected."})
				manager.send(jsonMessage, conn)
			}

		case message := <-manager.Broadcast:
			for conn := range manager.Clients {
				select {
				case conn.Send <- message:
				default:
					close(conn.Send)
					delete(manager.Clients, conn)
				}
			}
		}
	}
}

func (c *Client) read(manager *ClientManager) {
	defer func() {
		manager.Unregister <- c
		c.Socket.Close()
	}()

	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			manager.Unregister <- c
			c.Socket.Close()
			break
		}
		jsonMessage, _ := json.Marshal(&Message{Sender: c.ID, Content: string(message)})
		manager.Broadcast <- jsonMessage
	}
}

func (c *Client) write() {
	defer func() {
		c.Socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.Socket.WriteMessage(websocket.TextMessage, message)
		default:
			fmt.Println("default case")
		}
	}
}

func Ws(res http.ResponseWriter, req *http.Request) {
	conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if error != nil {
		http.NotFound(res, req)
		return
	}

	id, _ := uuid.NewV4()
	client := &Client{ID: id.String(), Socket: conn, Send: make(chan []byte)}

	manager := &ClientManager{}
	manager.Register <- client

	go client.read(manager)
	go client.write()
}
