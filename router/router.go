package router

import (
	v1 "gin_template/api/v1"
	"gin_template/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	publicGroup := api.Group("")
	{
		publicGroup.POST("login", v1.LoginApi)
	}
	privateGroup := api.Group("")
	privateGroup.Use(middleware.AuthToken)
	{

	}
}
