package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type loggedRequest struct {
	Method string
	Path   string
}

var requests = make([]loggedRequest, 0)
var clientConnections = NewConnections()

func main() {
	monitoringServer := http.NewServeMux()
	monitoringServer.HandleFunc("/getRequests", getRequests)

	websocketServer := http.NewServeMux()
	websocketServer.HandleFunc("/ws", serveWs)

	requestServer := http.NewServeMux()
	requestServer.HandleFunc("/", logRequest)

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
	request := loggedRequest{r.Method, r.URL.Path}
	jsonBytes, _ := json.Marshal(request)

	clientConnections.sendToAll <- jsonBytes

	requests = append(requests, request)
	w.WriteHeader(200)
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
