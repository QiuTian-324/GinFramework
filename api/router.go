package api

import (
	"gin_template/internal/libs"

	"github.com/gin-gonic/gin"
)

// CollectRoutes 注册所有路由
func CollectRoutes(router *gin.Engine) {
	api := router.Group("/api")

	// 示例：返回一个 ping 响应
	api.GET("/ping", func(ctx *gin.Context) {
		libs.SuccessResponse(ctx, "pong", nil)
	})

	// 注册用户路由
	RegisterUserRoutes(api)
}
