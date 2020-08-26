package bitServer

import (
	"net/http"
)

// GET /login
// Show the login page
func (s *BitServer) HandleLogin() HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		t := bitUtils.parseTemplateFiles("login.layout", "public.navbar", "login")
		t.Execute(w, nil)
	}
}

// GET /signup
// Show the signup page
func (s *BitServer) HandleSignup() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bitUtils.generateHTML(w, nil, "login.layout", "public.navbar", "signup")
	}
}

// POST /signup
// Create the user account
func (s *BitServer) HandleSignupAccount() HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		err := request.ParseForm()
		if err != nil {
			s.Logger.Error(err, "Cannot parse form")
		}
		user := data.User{
			Name:     request.PostFormValue("name"),
			Email:    request.PostFormValue("email"),
			Password: request.PostFormValue("password"),
		}
		if err := user.Create(); err != nil {
			s.Logger.Error(err, "Cannot create user")
		}
		http.Redirect(w, request, "/login", 302)
	}
}

// POST /authenticate
// Authenticate the user given the email and password
func (s *BitServer) HandleAuthenticate() HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		err := request.ParseForm()
		user, err := data.UserByEmail(request.PostFormValue("email"))
		if err != nil {
			s.Logger.Error(err, "Cannot find user")
		}
		if user.Password == data.Encrypt(request.PostFormValue("password")) {
			session, err := user.CreateSession()
			if err != nil {
				s.Logger.Error(err, "Cannot create session")
			}
			cookie := http.Cookie{
				Name:     "_cookie",
				Value:    session.Uuid,
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)
			http.Redirect(w, request, "/", 302)
		} else {
			http.Redirect(w, request, "/login", 302)
		}
	}
}

// GET /logout
// Logs the user out
func (s *BitServer) HandleLogout() HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		cookie, err := request.Cookie("_cookie")
		if err != http.ErrNoCookie {
			s.Logger.Warning(err, "Failed to get cookie")
			session := data.Session{Uuid: cookie.Value}
			session.DeleteByUUID()
		}
		http.Redirect(w, request, "/", 302)
	}
}
