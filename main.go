package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/pczora/dessert/requestserver"
)

type LoggedRequest struct {
	Method string
	Path   string
}

var requests = make([]LoggedRequest, 0)
var clientConnections = NewConnections()

func main() {

	rServer := requestserver.NewRequestServer()
	monitoringServer := http.NewServeMux()
	monitoringServer.HandleFunc("/getRequests", getRequests)

	websocketServer := http.NewServeMux()
	websocketServer.HandleFunc("/ws", serveWs)

	requestServer := http.NewServeMux()
	requestServer.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logRequest(w, r)
		rServer.HandleRequest(w, r)
	})

	go func() {
		if err := http.ListenAndServe(":8080", requestServer); err != nil {
			panic(err)
		}
	}()

	go func() {
		if err := http.ListenAndServe(":8081", websocketServer); err != nil {
			panic(err)
		}
	}()

	if err := http.ListenAndServe(":9080", monitoringServer); err != nil {
		panic(err)
	}

}

func logRequest(w http.ResponseWriter, r *http.Request) {
	request := LoggedRequest{r.Method, r.URL.Path}
	jsonBytes, _ := json.Marshal(request)

	clientConnections.sendToAll <- jsonBytes

	requests = append(requests, request)
}

func getRequests(w http.ResponseWriter, _ *http.Request) {
	var buffer bytes.Buffer
	jsonBytes, _ := json.Marshal(requests)
	buffer.Write(jsonBytes)
	w.Write(buffer.Bytes())
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	go clientConnections.run()
	serveWebsocket(clientConnections, w, r)
}
