package global

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
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

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 允许所有来源的连接
	CheckOrigin: func(r *http.Request) bool { return true },
}
