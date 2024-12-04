package config

import (
	"context"
	"fmt"
	"yug_server/pkg"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

// RedisInternal  初始化redis
func RedisInternal() (*redis.Client, error) {
	reClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", viper.GetString("redis.host"), viper.GetString("redis.port")),
		Password: "",
		DB:       viper.GetInt("redis.db"),
	})
	_, err := reClient.Ping(context.Background()).Result()
	if err != nil {
		pkg.Error("Redis连接失败", err)
		return nil, err
	}
	return reClient, nil

}
