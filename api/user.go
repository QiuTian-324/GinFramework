package api

import (
	"gin_template/internal/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes 注册用户相关的路由
func RegisterUserRoutes(router *gin.RouterGroup) {
	userGroup := router.Group("/user")

	userGroup.POST("/register", handlers.Register)
	userGroup.POST("/login", handlers.Login)
}
