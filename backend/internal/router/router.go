package router

import (
	"backend/internal/handler"
	"backend/internal/middleware"
	"backend/internal/repository"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	//! Repository Block
	userRepository := repository.NewUserRepository()

	//! Service Block
	userService := service.NewUserService(userRepository)

	//! define the handler
	pingHandler := handler.NewPingHandler()
	authHandler := handler.NewAuthHandler(userService)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", pingHandler.Ping)
		v1.POST("/register", authHandler.Register)
	}

	return r
}
