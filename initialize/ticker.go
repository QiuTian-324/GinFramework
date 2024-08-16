package initialize

import (
	"github.com/spf13/viper"
	"time"
)

// TickerInitialize 初始化定时器
func TickerInitialize() *time.Ticker {
	return time.NewTicker(time.Duration(viper.GetInt("ticker.Second")) * time.Second)
}
