package config

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"
	"yug_server/pkg"

	"github.com/spf13/viper"
)

func GinInternal() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	portUse := viper.GetString("service.port")
	if portUse == "" {
		portUse = "8080"
	}
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", portUse),
		Handler: RouterGin,
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
