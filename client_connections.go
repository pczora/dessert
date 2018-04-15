package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
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
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (c *ClientConnections) run() {
	for {
		select {
		case client := <-c.register:
			log.Println("Client registered")
			c.clients[client] = true
		case msg := <-c.sendToAll:
			log.Println("Sending message to clients...")
			for client := range c.clients {
				client.send <- msg
				log.Println("sent.")
			}
		}
		// TODO: unregister clients/gracefully terminate connections
	}

}
