// go:build wireinject
//go:build wireinject
// +build wireinject

package cmd

import (
	"yug_server/internal/handlers"
	"yug_server/internal/server"

	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func InitializeChatHandler(db *gorm.DB, rds *redis.Client, logger *zap.Logger) *handlers.ChatHandler {
	wire.Build(server.ChatHandlerSet)
	return &handlers.ChatHandler{}
}

func InitializeUserHandler(db *gorm.DB, rds *redis.Client, logger *zap.Logger) *handlers.UserHandler {
	wire.Build(server.UserHandlerSet)
	return &handlers.UserHandler{}
}
