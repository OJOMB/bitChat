package bitserver

import (
	"errors"
	"net/http"

	"github.com/OJOMB/bitChat/bitconfig"
	"github.com/OJOMB/bitChat/bitrepo"
	"github.com/OJOMB/bitChat/bituser"
)

// BitServer models an instance of the bitChat Server
type BitServer struct {
	Config *bitconfig.Config
	Router BitRouter
	Repo   bitrepo.BitRepo
	logger bitlogger.Logger
}

// NewBitServer returns a BitServer instance
func NewBitServer(config *bitconfig.Config, router Router, repo bitRepo.Repo, logger bitLog.Logger) *bitServer {
	var s *BitServer = &Server{
		config,
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

// ListenAndServe Listens and serves on the address specified in the BitServer Config
func (s *BitServer) ListenAndServe() {
	addr := s.Config["Address"] + ":" + s.Config["Port"]
	s.Logger.Fatal(http.ListenAndServe(addr, s))
}

func (s *BitServer) setupRoutes() {
	// register handlers against patterns in the DefaultServeMux
	files := http.FileServer(http.Dir("public"))
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

// CheckSession checks the request for an active session cookie
func (s *BitServer) CheckSession(r *http.Request) (*bituser.Session, error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess, err := s.Repo.GetSession(cookie.value)
		if err != nil {
			s.Logger.Warning(
				"Found session cookie but encountered the following error whilst attempting to retrieve matching session from DB",
			)
		}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}
