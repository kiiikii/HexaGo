package router

import (
	"backend/internal/handler"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	pingHandler := handler.NewPingHandler()

	v1 := r.Group("api/v1")
	{
		v1.GET("/ping", pingHandler.Ping)
	}

	return r
}
