// internal/services/chat_service.go
package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"yug_server/internal/dto"
	"yug_server/internal/libs"
	"yug_server/internal/repo"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type WsUseCase struct {
	repo   repo.ChatRepo
	rds    *redis.Client
	logger *zap.Logger
}

func NewWsUseCase(repo repo.ChatRepo, rds *redis.Client, logger *zap.Logger) *WsUseCase {
	return &WsUseCase{repo: repo, rds: rds, logger: logger}
}

// 发送消息
func (s *WsUseCase) SendMessage(conn *websocket.Conn, messageData []byte) error {
	var msg dto.Message
	err := json.Unmarshal(messageData, &msg)
	if err != nil {
		s.logger.Error("unmarshal message error", zap.Error(err))
		return err
	}

	switch msg.Type {
	case dto.TextMessageType:
		msgResponse := dto.NewMessage(
			msg.Type,
			msg.Content,
			msg.URL,
			msg.FileName,
			msg.SenderID,
			msg.ReceiverID,
			msg.GroupID,
		)
		err = conn.WriteJSON(msgResponse)
		if err != nil {
			s.logger.Error("消息发送失败", zap.Error(err))
			return err
		}

		// 判断接收者是否在线
		if !s.IsUserOnline(msg.ReceiverID) {
			s.logger.Info("接收者不在线，存储离线消息")
			// 存储离线消息到数据库
			err = s.StoreOfflineMessage(msg)
			if err != nil {
				s.logger.Error("存储离线消息失败", zap.Error(err))
			}
			return nil
		}

		// 查找接收者的连接并发送消息
		receiverConn := s.GetReceiverConnection(msg.ReceiverID)
		if receiverConn != nil {
			err = receiverConn.WriteJSON(msgResponse)
			if err != nil {
				s.logger.Error("消息发送给接收者失败", zap.Error(err))
				return err
			}
		}

	default:
		s.logger.Error("消息发送失败")
	}
	return nil
}

// 添加连接
func (s *WsUseCase) AddConnection(userID uint64, conn *websocket.Conn) {
	libs.Mu.Lock()
	libs.UserConnectionsMap[cast.ToString(userID)] = conn
	libs.Mu.Unlock()

	// 将用户在线状态存储到 Redis
	libs.RedisSet(context.Background(), fmt.Sprintf("user:%s:online", cast.ToString(userID)), "true", time.Hour*24)
}

// 移除连接
func (s *WsUseCase) RemoveConnection(userID uint64) {
	libs.Mu.Lock()
	delete(libs.UserConnectionsMap, cast.ToString(userID))
	libs.Mu.Unlock()

	// 从 Redis 中删除用户在线状态
	libs.RedisDelete(context.Background(), fmt.Sprintf("user:%s:online", cast.ToString(userID)))
}

// 获取接收者的连接
func (s *WsUseCase) GetReceiverConnection(receiverID uint64) *websocket.Conn {
	libs.Mu.Lock()
	defer libs.Mu.Unlock()
	return libs.UserConnectionsMap[cast.ToString(receiverID)]
}

// 心跳检测
func (s *WsUseCase) Heartbeat(conn *websocket.Conn, userID uint64) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for range ticker.C {
		err := conn.WriteMessage(websocket.PingMessage, nil)
		if err != nil {
			s.logger.Error("心跳检测失败", zap.Error(err))
			s.RemoveConnection(userID)
			return
		}
	}
}
func (s *WsUseCase) IsUserOnline(userID uint64) bool {
	onlineStatus, err := libs.RedisGet(context.Background(), fmt.Sprintf("user:%s:online", cast.ToString(userID)))
	if err != nil {
		s.logger.Error("获取用户在线状态失败", zap.Error(err))
		return false
	}
	return onlineStatus == "true"
}

func (s *WsUseCase) StoreOfflineMessage(msg dto.Message) error {

	// offlineMsg := chat.ChatMsg{
	// 	ContentType:    string(msg.Type),
	// 	Content:        msg.Content,
	// 	SenderID:       msg.SenderID,
	// 	ReceiverID:     msg.ReceiverID,
	// 	GroupID:        msg.GroupID,
	// 	Url:            msg.URL,
	// 	FileName:       msg.FileName,
	// 	OfflineMessage: true,
	// }

	// repo := chat.NewChatMsgRepository(s.db, s.rds, s.logger)

	// // 执行插入操作
	// err := repo.Insert(&offlineMsg)
	// if err != nil {
	// 	s.logger.Error("存储离线消息失败", zap.Error(err))
	// 	return err
	// }
	return nil
}
