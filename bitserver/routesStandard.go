package bitserver

import (
	"net/http"
)

// HandleErr Methods: [GET], Url: /err?msg=
// shows the error message page
func (s *BitServer) HandleErr() http.HandlerFunc {
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

// HandleIndex Methods: [GET], Url: /
// responds with home page
func (s *BitServer) HandleIndex() http.HandlerFunc {
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

// HandleStatic Methods: [GET], Url: /public
// static file server
func (s *BitServer) HandleStatic() http.Handler {
	return http.StripPrefix(
		"/static/",
		http.FileServer(http.Dir("/public")),
	)
}
