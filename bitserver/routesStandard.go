package bitserver

import (
	"net/http"
)

// GET /err?msg=
// shows the error message page
func (s *BitServer) HandleErr() HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		vals := request.URL.Query()
		_, err := session(writer, request)
		if err != nil {
			bitUtils.generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
		} else {
			bitUtils.generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
		}
	}
}

// GET
// responds with home page
func (s *BitServer) HandleIndex() HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		threads, err := data.Threads()
		if err != nil {
			error_message(writer, request, "Cannot get threads")
		} else {
			_, err := session(writer, request)
			if err != nil {
				bitUtils.generateHTML(writer, threads, "layout", "public.navbar", "index")
			} else {
				bitUtils.generateHTML(writer, threads, "layout", "private.navbar", "index")
			}
		}
	}
}

// GET
// static file server from /public
func (s *BitServer) HandleStatic() http.Handler {
	return http.StripPrefix(
		"/static/",
		http.FileServer(http.Dir("/public")),
	)
}
