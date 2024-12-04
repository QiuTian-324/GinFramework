package utils

import (
	"fmt"
	"yug_server/global"
)

func TickerUse() {
	for range global.MyTicker.C {
		// 到了设置的时间就会执行当前函数
		fmt.Println("定时器执行")
	}
	defer global.MyTicker.Stop()
}
