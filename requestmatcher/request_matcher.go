package requestmatcher

import (
	"io/ioutil"
	"net/http"
)

// RequestMatcher is used to determine, whether an HTTP request matches specific criteria
type RequestMatcher interface {
	Match(r *http.Request) bool
}

// DefaultRequestMatcher matches only on HTTP method, request path and payload, if any
type DefaultRequestMatcher struct {
	Method string
	Path   string
	Body   string
}

// Match returns always true
func (d DefaultRequestMatcher) Match(r *http.Request) bool {
	return d.matchMethod(r) && d.matchPath(r) && d.matchBody(r)
}

func (d DefaultRequestMatcher) matchMethod(r *http.Request) bool {
	return d.Method == r.Method || d.Method == ""
}

func (d DefaultRequestMatcher) matchPath(r *http.Request) bool {
	return d.Path == r.URL.RequestURI() || d.Path == ""
}

func (d DefaultRequestMatcher) matchBody(r *http.Request) bool {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return false
	}

	return d.Body == string(body) || d.Body == ""
}
