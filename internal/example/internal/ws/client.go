package ws

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Msg  chan []byte
}

func (c *Client) Send(msg []byte) error {
	return c.Conn.WriteMessage(websocket.TextMessage, []byte("ping"))

}
