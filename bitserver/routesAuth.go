package bitserver

import (
	"net/http"
	"time"

	"github.com/OJOMB/bitChat/bituser"
	"github.com/OJOMB/bitChat/bitutils"
)

// HandleLogin Methods: [GET], Url: /login
// Show the login page
func (s *BitServer) HandleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := bitutils.ParseTemplateFiles("login.layout", "public.navbar", "login")
		t.Execute(w, nil)
	}
}

// HandleSignup Methods: [GET], Url: /signup
// Show the signup page
func (s *BitServer) HandleSignup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bitutils.GenerateHTML(w, nil, "login.layout", "public.navbar", "signup")
	}
}

// HandleSignupAccount Methods: [POST], Url: /signup
// Create the user account
func (s *BitServer) HandleSignupAccount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			s.logger.Error("failed to parse form with error: %s", err.Error())
		}
		user := bituser.User{
			Name:      r.PostFormValue("name"),
			Email:     r.PostFormValue("email"),
			Password:  r.PostFormValue("password"),
			Bio:       r.PostFormValue("bio"),
			CreatedAt: time.Now(),
		}
		if err := s.Repo.CreateUser(user.ToDocument()); err != nil {
			s.logger.Error(err, "Cannot create user")
		}
		http.Redirect(w, r, "/login", 302)
	}
}

// HandleAuthenticate Methods: [POST], Url: /authenticate
// Authenticate the user given the email and password
func (s *BitServer) HandleAuthenticate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			s.logger.Error(err, "Failed to parse data from Login form")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		user, err := s.Repo.GetUserByEmail(r.PostFormValue("email"))
		if err != nil {
			s.logger.Error("Failed to retrieve user with email: %s. Encountered the following error: %s", r.PostFormValue("email"), err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		if user.Password == bitutils.Encrypt(r.PostFormValue("password")) {
			sess, err := user.CreateSession()
			if err != nil {
				s.logger.Error("Failed to create session, encounter the following error: %s", err.Error())
			}
			cookie := http.Cookie{
				Name:     "_cookie",
				Value:    sess.Uuid,
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/", 302)
		} else {
			http.Redirect(w, r, "/login", 302)
		}
	}
}

// HandleLogout Methods: [GET], Url: /logout
// Logs the user out
func (s *BitServer) HandleLogout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("_cookie")
		if err != http.ErrNoCookie {
			s.Logger.Warning(err, "Failed to get cookie")
			session := data.Session{Uuid: cookie.Value}
			session.DeleteByUUID()
		}
		http.Redirect(w, r, "/", 302)
	}
}
