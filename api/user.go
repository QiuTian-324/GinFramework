package api

import (
	"yug_server/internal/handlers"
	"yug_server/internal/middleware"
	"yug_server/pkg"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes 注册用户相关的路由
func RegisterUserRoutes(router *gin.RouterGroup) {
	userHandler := pkg.GetHandler("userHandler").(*handlers.UserHandler)
	userGroup := router.Group("/user")
	userGroup.POST("/register", userHandler.Register)
	userGroup.POST("/login", userHandler.Login)
	// userGroup.POST("/github/authURL", handlers.GithubAuthURL)
	// userGroup.GET("/github/callback", handlers.GithubCallback)

	// 需要鉴权的接口
	userGroup.Use(middleware.AuthMiddleware())
	// userGroup.GET("/list", userHandler.GetUserSeesionList)
	// userGroup.POST("/add_friend", userHandler.AddFriend)
	// userGroup.GET("/friends", userHandler.GetFriends)
	userGroup.GET("/query", userHandler.QueryUser)
	userGroup.POST("/logout", userHandler.Logout)
}
