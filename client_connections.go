package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type ClientConnections struct {
	conns []*WebsocketClient
}

func NewConnections() *ClientConnections {
	return &ClientConnections{}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (c *ClientConnections) sendToAll(msg string) {
	for _, client := range c.conns {
		client.send([]byte(msg))
	}
}

func serveWebsocket(c *ClientConnections, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &WebsocketClient{conn}
	c.conns = append(c.conns, client)
}
