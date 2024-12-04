package config

import (
	router "yug_server/api"
	"yug_server/internal/libs"
	"yug_server/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var RouterGin *gin.Engine

func Routerinternal() {

	// 默认是线上模式
	ginMode := gin.ReleaseMode
	if viper.GetBool("development.develop") {
		// 当开发者模式是true时，使用debug模式
		ginMode = gin.DebugMode
	}
	gin.SetMode(ginMode)
	RouterGin = gin.New()

	RouterGin.Use(gin.Logger(), gin.Recovery()) // 日志
	// RouterGin.Use(middleware.InjectDB())        // 注入数据库
	RouterGin.Use(middleware.Cors()) // 直接放行全部跨域请求
	router.CollectRoutes(RouterGin)

	RouterGin.NoRoute(func(c *gin.Context) {
		libs.NotFoundResponse(c, "抱歉，您请求的接口不存在")
	})

}
