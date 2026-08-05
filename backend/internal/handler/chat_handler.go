package handler

import (
	"backend/internal/chat"
	"backend/internal/utils"
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

type ChatHandler struct {
	hub *chat.Hub
}

func NewChatHandler(hub *chat.Hub) *ChatHandler {
	return &ChatHandler{
		hub: hub,
	}
}

func (h *ChatHandler) ServeWS(c *gin.Context) {
	tokenString := c.Query("token")
	userID, err := utils.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}

	//* Send Connection to Hub
	h.hub.Register <- conn

	defer func() {
		h.hub.Unregister <- conn
	}()

	//* Looping for Read & Write
	for {
		//! Reading message
		var chatMsg chat.Message
		err := conn.ReadJSON(&chatMsg)
		if err != nil {
			break
		}
		chatMsg.SenderID = userID

		h.hub.Broadcast <- chatMsg
	}
}
