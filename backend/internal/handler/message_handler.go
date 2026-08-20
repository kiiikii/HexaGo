package handler

import (
	"backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	messageService service.MessageService
}

func NewMessageHandler(messageService service.MessageService) *MessageHandler {
	return &MessageHandler{messageService: messageService}
}

func (h *MessageHandler) GetMessage(c *gin.Context) {
	message, err := h.messageService.GetMessage(50)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": message})
}

func (h *MessageHandler) GetHistory(c *gin.Context) {
	//! get the string from url, default 50 and 0
	limitStr := c.DefaultQuery("limit", "50")
	offsetStr := c.DefaultQuery("offset", "0")
	room := c.DefaultQuery("room", "general")

	//! convert string to int
	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)

	//! Security check
	if limit > 100 {
		limit = 100
	}

	//! Calling the service
	messages, err := h.messageService.GetMessages(limit, offset, room)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//! Returning Data
	c.JSON(http.StatusOK, gin.H{"data": messages})
}

func (h *MessageHandler) DeleteMessage(c *gin.Context) {
	messageID := c.Param("id")
	userID, exist := c.Get("userID")

	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user context"})
		return
	}

	err := h.messageService.DeleteMessage(messageID, userIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": "Message deleted Successfully"})
}
