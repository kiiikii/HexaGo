package router

import (
	"backend/internal/chat"
	"backend/internal/handler"
	"backend/internal/logger"
	"backend/internal/middleware"
	"backend/internal/repository"
	"backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	logger.InitLogger()
	r := gin.New()

	r.Use(middleware.RequestLogger())
	r.Use(gin.Recovery())

	r.Use(middleware.CORSMiddleware())

	//! Repository Block
	userRepository := repository.NewUserRepository(db)
	messageRepository := repository.NewMessageRepository(db)

	//! Service Block
	userService := service.NewUserService(userRepository)
	messageService := service.NewMessageService(messageRepository)

	//! define the handler
	pingHandler := handler.NewPingHandler()
	authHandler := handler.NewAuthHandler(userService)
	messageHandler := handler.NewMessageHandler(messageService)

	hub := chat.NewHub()
	chatHandler := handler.NewChatHandler(hub, messageService)
	go hub.Run()

	v1 := r.Group("/api/v1")
	secure := v1.Group("/secure")
	secure.Use(middleware.RequireAuth())
	{
		secure.GET("/me", func(c *gin.Context) {
			userID, exists := c.Get("userID")
			if !exists {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to locate user context",
				})
				return
			}

			userIDStr, ok := userID.(string)
			if !ok {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Invalid user context type structure",
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "Hello, user",
				"id":      userIDStr,
			})
		})

		secure.GET("/messages", messageHandler.GetHistory)
		secure.DELETE("/messages/:id", messageHandler.DeleteMessage)
	}
	{
		v1.GET("/ping", pingHandler.Ping)
		v1.POST("/register", authHandler.Register)
		v1.POST("/login", authHandler.Login)
		v1.POST("/refresh", authHandler.RefreshToken)
		v1.GET("/ws", chatHandler.ServeWS)
	}

	return r
}
