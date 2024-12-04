package chat

import (
	"context"
	"yug_server/internal/repo"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type chatRepo struct {
	db     *gorm.DB
	rds    *redis.Client
	logger *zap.Logger
}

func NewChatRepo(db *gorm.DB, rds *redis.Client, logger *zap.Logger) repo.ChatRepo {
	return &chatRepo{db: db, rds: rds, logger: logger}
}

func (r *chatRepo) SendMessage(ctx context.Context, conn *websocket.Conn, messageData []byte) error {
	return nil
}

// func (r *chatRepo) Insert(msg *model.ChatMsg) error {
// 	return r.DB.Create(msg).Error
// }

// func (r *chatRepo) GetById(id uint64) (*model.ChatMsg, error) {
// 	var msg model.ChatMsg
// 	if err := r.DB.Where("id = ?", id).First(&msg).Error; err != nil {
// 		return nil, err
// 	}
// 	return &msg, nil
// }
