package chat

import (
	"backend/internal/model"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Room string
}

type Hub struct {
	Rooms      map[string]map[*Client]bool
	Broadcast  chan model.Message
	Register   chan *Client
	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]map[*Client]bool),
		Broadcast:  make(chan model.Message),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	//! Infinite loops
	for {
		select {
		case client := <-h.Register:
			if h.Rooms[client.Room] == nil {
				h.Rooms[client.Room] = make(map[*Client]bool)
			}
			h.Rooms[client.Room][client] = true

		case client := <-h.Unregister:
			if _, ok := h.Rooms[client.Room][client]; ok {
				delete(h.Rooms[client.Room], client)
				client.Conn.Close()

				//! Cleaning empty Room
				if len(h.Rooms[client.Room]) == 0 {
					delete(h.Rooms, client.Room)
				}
			}
		case message := <-h.Broadcast:
			//! Only Loop Through Client in the specific room the message
			//! Make sure a model.Message has Room Field
			for client := range h.Rooms[message.Room] {
				err := client.Conn.WriteJSON(message)
				if err != nil {
					client.Conn.Close()
					delete(h.Rooms[message.Room], client)
				}
			}
		}
	}
}
