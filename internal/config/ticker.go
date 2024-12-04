package config

import (
	"time"

	"github.com/spf13/viper"
)

// Tickerinternal 初始化定时器
func Tickerinternal() *time.Ticker {
	return time.NewTicker(time.Duration(viper.GetInt("ticker.Second")) * time.Second)
}
