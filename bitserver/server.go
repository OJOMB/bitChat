package bitserver

import (
	"net/http"

	"github.com/OJOMB/bitChat/bitLog"
)

type BitServer struct {
	address string
	Router  bitRouter
	Repo    bitRepo
	logger  bitLogger
}

func NewBitServer(address string, router Router, repo bitRepo.Repo, logger bitLog.Logger) *bitServer {
	var s *BitServer = &Server{
		address
		router,
		repo,
		logger,
	}
	s.Logger.Info("Init server: Establishing routes")
	s.setupRoutes()
	return s
}

func (s *BitServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func (s *BitServer) ListenAndServe() error {
	s.Logger.Fatal(http.ListenAndServe(s.address, s))
}

func (s *BitServer) setupRoutes() {
	// register handlers against patterns in the DefaultServeMux
	s.Router.Handle("/static/", http.StripPrefix("/static/", files))
	s.Router.HandleFunc("/", s.HandleIndex())
	s.Router.HandleFunc("/err", s.HandleErr())
	s.Router.HandleFunc("/login", s.HandleLogin())
	s.Router.HandleFunc("/logout", s.HandleLogout())
	s.Router.HandleFunc("/signup", s.HandleSignup())
	s.Router.HandleFunc("/signup_account", s.HandleSignupAccount())
	s.Router.HandleFunc("/authenticate", s.HandleAuthenticate())
	s.Router.HandleFunc("/thread/new", s.HandleNewThread())
	s.Router.HandleFunc("/thread/create", s.HandleCreateThread())
	s.Router.HandleFunc("/thread/post", s.HandlePostThread())
	s.Router.HandleFunc("/thread/read", s.HandleReadThread())

}
