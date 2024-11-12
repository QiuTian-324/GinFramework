package cmd

import (
	"fmt"
	"gin_template/global"
	"gin_template/internal/config"
	"gin_template/internal/middleware"
	"os"
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

	// 注入db
	middleware.InjectDB()

	// 初始化雪花算法
	// internal.SnowFlakeinternal()
	// 初始化定时器
	// global.MyTicker = internal.Tickerinternal()
	// 开启定时器
	// go utils.TickerUse()
	// 初始化路由要放到最后面不然会阻止其他的函数初始化
	// 初始化路由
	config.Routerinternal()
}
