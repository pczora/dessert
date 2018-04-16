package requestserver

import (
	"log"
	"net/http"
)

type Request struct {
	method string
	URL    string
}

type Response struct {
	statusCode int
	body       string
}

type RequestServer struct {
	endpoints map[Request]Response
}

func NewRequestServer() RequestServer {
	endpoints := make(map[Request]Response)
	// TODO: Read config from file/db/cli params/whatever
	endpoints[Request{http.MethodGet, "/test"}] = Response{201, "alright!"}
	return RequestServer{endpoints}
}

func (rs RequestServer) HandleRequest(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	response := rs.endpoints[Request{r.Method, r.URL.String()}]
	w.WriteHeader(response.statusCode)
	w.Write([]byte(response.body))
}
