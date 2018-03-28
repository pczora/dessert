package main

import (
	"github.com/gorilla/websocket"
)

type ClientConnections struct {
	clients   map[*WebsocketClient]bool
	register  chan *WebsocketClient
	sendToAll chan []byte
}

func NewConnections() *ClientConnections {
	return &ClientConnections{
		clients:   make(map[*WebsocketClient]bool),
		register:  make(chan *WebsocketClient),
		sendToAll: make(chan []byte, 2048),
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
}

func (c *ClientConnections) run() {
	for {
		select {
		case client := <-c.register:
			c.clients[client] = true
		case msg := <-c.sendToAll:
			for client := range c.clients {
				client.send <- msg
			}
		}
		// TODO: unregister clients/gracefully terminate connections
	}

}
