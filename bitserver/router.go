package bitserver

import "net/http"

// BitRouter defines the method set required for the BitServer router
type BitRouter interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	HandleFunc(pattern string, handler func(ResponseWriter http.ResponseWriter, Request *http.Request))
	Handle(pattern string, handler http.Handler)
}
