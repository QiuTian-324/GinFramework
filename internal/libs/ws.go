package libs

import (
	"sync"

	"github.com/gorilla/websocket"
)

var (
	UserConnectionsMap = make(map[string]*websocket.Conn)
	Mu                 sync.Mutex
)
