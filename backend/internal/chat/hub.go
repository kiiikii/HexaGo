package chat

import "github.com/gorilla/websocket"

type Hub struct {
	Clients    map[*websocket.Conn]bool
	Broadcast  chan Message
	Register   chan *websocket.Conn
	Unregister chan *websocket.Conn
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*websocket.Conn]bool),
		Broadcast:  make(chan Message),
		Register:   make(chan *websocket.Conn),
		Unregister: make(chan *websocket.Conn),
	}
}

func (h *Hub) Run() {
	//! Infinite loops
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				client.Close()
			}
		case message := <-h.Broadcast:
			//! Looping to send the text message
			for client := range h.Clients {
				err := client.WriteJSON(message)
				if err != nil {
					client.Close()
					delete(h.Clients, client)
				}
			}
		}
	}
}
