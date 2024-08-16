package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// ViperInitialize  初始化viper配置
func ViperInitialize() {
	viper.SetConfigName("configs")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("配置文件读取失败" + err.Error())
		os.Exit(1)
	}
}
