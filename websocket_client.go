package main

import (
	"github.com/gorilla/websocket"
)

type WebsocketClient struct {
	conn *websocket.Conn
}

func (c *WebsocketClient) send(payload []byte) {
	c.conn.WriteMessage(websocket.TextMessage, payload)
}
