package global

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var (
	// Logger 日志
	Logger *zap.SugaredLogger
	// DB mysql数据库
	DB *gorm.DB
	// RedisClient redis数据库
	RedisClient *redis.Client
	// Code 当前系统code 六位随机数
	Code int
	// MyTicker 全局定时器
	MyTicker *time.Ticker
)

// RandomKey 随机数相关的种子
var RandomKey string = "Key12368975490255"

// RedisKey redis存储随机前缀
var RedisKey string = "redis_"
