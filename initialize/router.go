package initialize

import (
	"context"
	"fmt"
	"gin_template/global"
	"gin_template/middleware"
	"gin_template/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func RouterInitialize() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	// 默认是线上模式
	ginMode := gin.ReleaseMode
	if viper.GetBool("development.develop") {
		// 当开发者模式是true时，使用debug模式
		ginMode = gin.DebugMode
	}
	gin.SetMode(ginMode)
	routerGin := gin.New()
	routerGin.Use(gin.Logger(), gin.Recovery())
	// 跨域，如需跨域可以打开下面的注释
	routerGin.Use(middleware.Cors()) // 直接放行全部跨域请求
	router.CollectRoutes(routerGin)

	routerGin.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "接口不存在"})
	})
	portUse := viper.GetString("service.port")
	if portUse == "" {
		portUse = "8080"
	}
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", portUse),
		Handler: routerGin,
	}
	fmt.Printf("服务开始运行，当前服务运行在 %s 端口", viper.GetString("service.port"))
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Errorf("服务启动失败: %s", err.Error())
			return
		}
	}()
	<-ctx.Done()
	stop()
	global.Logger.Info("shutting down server...")
	// 创建一个5秒超时的context
	ctx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()
	if err := server.Shutdown(ctx); err != nil {
		global.Logger.Error("服务关闭失败:" + err.Error())
		return
	}
	global.Logger.Debug("服务关闭成功")
}
