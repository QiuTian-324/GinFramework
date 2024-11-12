package config

import (
	"context"
	"fmt"
	router "gin_template/api"
	"gin_template/internal/libs"
	"gin_template/internal/middleware"
	"gin_template/pkg"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Routerinternal() {
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
	routerGin.Use(middleware.InjectDB())
	// 跨域，如需跨域可以打开下面的注释
	routerGin.Use(middleware.Cors()) // 直接放行全部跨域请求
	router.CollectRoutes(routerGin)

	routerGin.NoRoute(func(c *gin.Context) {
		libs.NotFoundResponse(c, "抱歉，您请求的接口不存在")
	})
	portUse := viper.GetString("service.port")
	if portUse == "" {
		portUse = "8080"
	}
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", portUse),
		Handler: routerGin,
	}
	pkg.Info("Fox IM 服务运行在：" + portUse + "端口上")
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			pkg.Error("服务启动失败", err)
			return
		}
	}()
	<-ctx.Done()
	stop()
	pkg.Info("服务关闭")
	// 创建一个5秒超时的context
	ctx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()
	if err := server.Shutdown(ctx); err != nil {
		pkg.Error("服务关闭失败", err)
		return
	}
	pkg.Info("服务关闭成功")
}
