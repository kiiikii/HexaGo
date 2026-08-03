package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	//* Allow all origin testing
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type ChatHandler struct{}

func NewChatHandler() *ChatHandler {
	return &ChatHandler{}
}

func (h *ChatHandler) ServeWS(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Client connection successfuly")

	//* Looping for Read & Write
	for {
		//! Reading message
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Read error:", err)
			break
		}

		//! Print Recieved Message
		fmt.Printf("Recieved: %s\n", msg)

		//! Echo Exact Message back to client
		err = conn.WriteMessage(msgType, msg)
		if err != nil {
			fmt.Println("Write error:", err)
			break
		}
	}
}
