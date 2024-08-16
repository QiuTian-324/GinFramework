package utils

import (
	"fmt"
	"gin_template/global"
)

func TickerUse() {
	for range global.MyTicker.C {
		// 到了设置的时间就会执行当前函数
		fmt.Println("定时器执行")
	}
	defer global.MyTicker.Stop()
}
