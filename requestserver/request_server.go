package requestserver

import (
	"log"
	"net/http"

	"github.com/pczora/dessert/requestmatcher"
)

type Response struct {
	StatusCode int
	Body       string
}
type RequestServer struct {
	endpoints map[requestmatcher.RequestMatcher]Response
}

func NewRequestServer() RequestServer {
	endpoints := make(map[requestmatcher.RequestMatcher]Response)
	// TODO: Read config from file/db/cli params/whatever

	endpoints[requestmatcher.DefaultRequestMatcher{Path: "/test", Method: http.MethodGet}] = Response{http.StatusOK, "ok!"}
	return RequestServer{endpoints}
}

func (rs RequestServer) HandleRequest(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)

	for matcher, response := range rs.endpoints {
		if matcher.Match(r) {
			w.WriteHeader(response.StatusCode)
			w.Write([]byte(response.Body))
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)

}
