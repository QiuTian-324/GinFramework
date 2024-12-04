package main

import (
	"yug_server/cmd"
	"yug_server/internal/config"
)

func main() {
	// 启动服务
	cmd.StartServer()

}

// 初始化配置
func init() {
	config.Viperinternal()
}
