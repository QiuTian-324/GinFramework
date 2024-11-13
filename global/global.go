package global

import (
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	// Logger 日志
	Logger *zap.Logger
	// DB mysql数据库
	DB *gorm.DB
	// RedisClient redis数据库
	RedisClient *redis.Client
	// Code 当前系统code 六位随机数
	Code int
	// MyTicker 全局定时器
	MyTicker *time.Ticker
)

