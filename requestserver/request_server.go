package requestserver

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pczora/dessert/requestmatcher"
)

// TODO: Make configurable
const basePath = "./base/"

type Response struct {
	StatusCode int
	Body       string
}
type RequestServer struct {
	endpoints map[requestmatcher.RequestMatcher]Response
}

func NewRequestServer() RequestServer {
	endpoints := parseFiles()
	return RequestServer{endpoints}
}

func parseFiles() map[requestmatcher.RequestMatcher]Response {

	endpoints := make(map[requestmatcher.RequestMatcher]Response)
	files, err := ioutil.ReadDir(basePath)

	if err != nil {
		log.Println("Error reading directory")

	}

	for _, file := range files {
		fileContent, err := ioutil.ReadFile(basePath + file.Name())
		if err != nil {
			log.Printf("Could not read file: %v", err)
		}
		// TODO: ATM it's GET only. What about the other methods?
		endpoints[requestmatcher.DefaultRequestMatcher{Path: "/" + file.Name(), Method: http.MethodGet}] = Response{http.StatusOK, string(fileContent)}
		log.Printf("Found file: %s", file.Name())
	}

	return endpoints
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
