package main

import (
	"gin_template/cmd"
	"gin_template/internal/config"
)

func main() {
	// 启动服务
	cmd.StartServer()
}

// 初始化配置
func init() {
	config.Viperinternal()
}
