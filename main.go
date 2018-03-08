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
		if err := http.ListenAndServe(":8080", requestServer); err != nil {
			panic(err)
		}
	}()

	if err := http.ListenAndServe(":9080", monitoringServer); err != nil {
		panic(err)
	}

}

func logRequest(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	url := r.URL.Path
	requests = append(requests, loggedRequest{method, url})
	w.WriteHeader(200)
}

func getRequests(w http.ResponseWriter, _ *http.Request) {
	var buffer bytes.Buffer
	jsonString, _ := json.Marshal(requests)
	buffer.Write(jsonString)
	w.Write(buffer.Bytes())
}
