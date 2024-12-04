package api

import (
	"yug_server/internal/handlers"
	"yug_server/internal/middleware"
	"yug_server/pkg"

	"github.com/gin-gonic/gin"
)

func RegisterChatRoutes(router *gin.RouterGroup) {
	chatHandler := pkg.GetHandler("chatHandler").(*handlers.ChatHandler)

	chatGroup := router.Group("/chat", middleware.AuthMiddleware())
	// websocket连接
	chatGroup.GET("/ws", chatHandler.Ws)
}
