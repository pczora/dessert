package main

import (
	"log"
	"net/http"

	"fmt"

	"github.com/gorilla/websocket"
)

type WebsocketClient struct {
	conn              *websocket.Conn
	send              chan []byte
	clientConnections *ClientConnections
}

func (c *WebsocketClient) run() {
	for {
		select {
		case msg := <-c.send:
			fmt.Println(string(msg))
			c.conn.WriteMessage(websocket.TextMessage, msg)
		}
	}
}

func serveWebsocket(c *ClientConnections, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &WebsocketClient{conn, make(chan []byte, 256), c}
	client.clientConnections.register <- client
	go client.run()
}
