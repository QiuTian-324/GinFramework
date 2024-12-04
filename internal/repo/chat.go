package repo

import (
	"context"

	"github.com/gorilla/websocket"
)

type ChatRepo interface {
	SendMessage(ctx context.Context, conn *websocket.Conn, messageData []byte) error
}
