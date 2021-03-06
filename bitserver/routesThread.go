package bitserver

import (
	"fmt"
	"net/http"
)

// HandleNewThread  Methods: [GET], Url: /threads/new
// Show the new thread form page
func (s *BitServer) HandleNewThread() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := session(w, r)
		if err != nil {
			http.Redirect(w, r, "/login", 302)
		} else {
			bitUtils.generateHTML(w, nil, "layout", "private.navbar", "new.thread")
		}
	}
}

// HandleCreateThread Methods:[POST], Url: /signup
// Create the user account
func (s *BitServer) HandleCreateThread() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sess, err := session(w, r)
		if err != nil {
			http.Redirect(w, r, "/login", 302)
		} else {
			err = r.ParseForm()
			if err != nil {
				danger(err, "Cannot parse form")
			}
			user, err := sess.User()
			if err != nil {
				danger(err, "Cannot get user from session")
			}
			topic := r.PostFormValue("topic")
			if _, err := user.CreateThread(topic); err != nil {
				danger(err, "Cannot create thread")
			}
			http.Redirect(w, r, "/", 302)
		}
	}
}

// HandleReadThread Methods: [GET], Url:/thread/read
// Show the details of the thread, including the posts and the form to write a post
func (s *BitServer) HandleReadThread() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vals := r.URL.Query()
		uuid := vals.Get("id")
		thread, err := data.ThreadByUUID(uuid)
		if err != nil {
			error_message(w, r, "Cannot read thread")
		} else {
			_, err := session(w, r)
			if err != nil {
				generateHTML(w, &thread, "layout", "public.navbar", "public.thread")
			} else {
				generateHTML(w, &thread, "layout", "private.navbar", "private.thread")
			}
		}
	}
}

// HandlePostThread Methods: [POST], Url:/thread/post
// Create the post
func (s *BitServer) HandlePostThread() {
	return func(w http.ResponseWriter, r *http.Request) {
		sess, err := session(w, r)
		if err != nil {
			http.Redirect(w, r, "/login", 302)
		} else {
			err = r.ParseForm()
			if err != nil {
				danger(err, "Cannot parse form")
			}
			user, err := sess.User()
			if err != nil {
				danger(err, "Cannot get user from session")
			}
			body := r.PostFormValue("body")
			uuid := r.PostFormValue("uuid")
			thread, err := data.ThreadByUUID(uuid)
			if err != nil {
				error_message(w, r, "Cannot read thread")
			}
			if _, err := user.CreatePost(thread, body); err != nil {
				danger(err, "Cannot create post")
			}
			url := fmt.Sprint("/thread/read?id=", uuid)
			http.Redirect(w, r, url, 302)
		}
	}
}
