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

func main() {
	monitoringServer := http.NewServeMux()
	monitoringServer.HandleFunc("/getRequests", getRequests)

	requestServer := http.NewServeMux()
	requestServer.HandleFunc("/", logRequest)

	go func() {
		if err := http.ListenAndServe("localhost:8080", requestServer); err != nil {
			panic(err)
		}
	}()

	if err := http.ListenAndServe("localhost:9080", monitoringServer); err != nil {
		panic(err)
	}

}

func logRequest(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	url := r.URL.Path
	requests = append(requests, loggedRequest{method, url})
	w.WriteHeader(200)
}

func getRequests(w http.ResponseWriter, r *http.Request) {
	var buffer bytes.Buffer
	for _, r := range requests {
		jsonString, _ := json.Marshal(r)
		buffer.Write(jsonString)
		buffer.WriteString("\n")
	}
	w.Write(buffer.Bytes())
}
