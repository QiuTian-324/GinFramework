package cmd

import (
	"fmt"
	"os"
	"yug_server/global"
	"yug_server/internal/config"
)

func StartServer() {
	global.Logger = config.LoggerInternal()
	gorminternal, err := config.GormInternal()
	if err != nil {
		fmt.Println("mysql数据库初始化失败" + err.Error())
		os.Exit(1)
	}
	global.DB = gorminternal
	redisClient, err := config.RedisInternal()
	if err != nil {
		fmt.Println("redis数据库初始化失败" + err.Error())
		os.Exit(1)
	}
	global.RedisClient = redisClient

	// 初始化雪花算法
	// internal.SnowFlakeinternal()
	// 初始化定时器
	// global.MyTicker = internal.Tickerinternal()
	// 开启定时器
	// go utils.TickerUse()
	// 初始化路由

	// 注册模块
	modules := []Module{
		&ChatModule{},
		&UserModule{},
	}

	// 初始化模块
	for _, module := range modules {
		if err := module.Initialize(global.DB, global.RedisClient, global.Logger); err != nil {
			fmt.Println("模块初始化失败: ", err)
			os.Exit(1)
		}
	}

	config.Routerinternal()

	config.GinInternal()
}
