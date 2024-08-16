package main

import (
	"fmt"
	"gin_template/global"
	"gin_template/initialize"
	"gin_template/utils"
	"os"
)

func init() {
	initialize.ViperInitialize()

}
func main() {
	initializes()
}
func initializes() {
	global.Logger = initialize.LoggerInitialize()
	gormInitialize, err := initialize.GormInitialize()
	if err != nil {
		fmt.Println("mysql数据库初始化失败" + err.Error())
		os.Exit(1)
	}
	global.DB = gormInitialize
	redisClient, err := initialize.RedisInitialize()
	if err != nil {
		fmt.Println("redis数据库初始化失败" + err.Error())
		os.Exit(1)
	}
	global.RedisClient = redisClient
	// 初始化雪花算法
	initialize.SnowFlakeInitialize()
	// 初始化定时器
	global.MyTicker = initialize.TickerInitialize()
	// 开启定时器
	go utils.TickerUse()
	// 初始化路由要放到最后面不然会阻止其他的函数初始化
	// 初始化路由
	initialize.RouterInitialize()
}
