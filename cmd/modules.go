package cmd

import (
	"yug_server/pkg"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Module interface {
	Initialize(db *gorm.DB, rds *redis.Client, logger *zap.Logger) error
}

type UserModule struct{}

func (m *UserModule) Initialize(db *gorm.DB, rds *redis.Client, logger *zap.Logger) error {
	userHandler := InitializeUserHandler(db, rds, logger)
	pkg.RegisterHandler("userHandler", userHandler)
	return nil
}

type ChatModule struct{}

func (m *ChatModule) Initialize(db *gorm.DB, rds *redis.Client, logger *zap.Logger) error {
	chatHandler := InitializeChatHandler(db, rds, logger)
	pkg.RegisterHandler("chatHandler", chatHandler)
	return nil
}
