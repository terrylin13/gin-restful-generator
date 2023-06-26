package ws

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var connManager = &ConnectionManager{
	connections: make(map[string]*Client),
}

type ConnectionManager struct {
	sync.RWMutex
	connections map[string]*Client
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
