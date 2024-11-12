package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Viperinternal  初始化viper配置
func Viperinternal() {
	viper.SetConfigName("configs")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("配置文件读取失败" + err.Error())
		os.Exit(1)
	}
}
