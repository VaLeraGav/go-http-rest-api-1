package apiserver

import (
	"net/http"

	"github.com/VaLeraGav/go-http-rest-api-1/internal/app/handlers"
	"github.com/VaLeraGav/go-http-rest-api-1/internal/app/middleware"
	"github.com/VaLeraGav/go-http-rest-api-1/internal/app/store"
	gorilla_handle "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
)

const (
	sessionName      = "auth_seesion_name"
	ctxKeyUser  int8 = iota
	ctxKeyRequestID
)

type server struct {
	router       *mux.Router
	logger       *logrus.Logger
	store        store.Store
	sessionStore sessions.Store
}

func newServer(store store.Store, sessionStore sessions.Store) *server {
	s := &server{
		router:       mux.NewRouter(),
		logger:       logrus.New(),
		store:        store,
		sessionStore: sessionStore,
	}

	s.configureRouting()
	// s.configureLogger()

	return s
}

// TODO: нужно сделать установку логирования
// func (s *server) configureLogger() error {
// 	level, err := logrus.ParseLevel(s.config.LogLevel)
// 	if err != nil {
// 		return err
// 	}

// 	s.logger.SetLevel(level)
// 	return nil
// }

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouting() {
	s.router.Use(middleware.SetRequestID(s.logger, ctxKeyRequestID))
	s.router.Use(middleware.LogRequest(s.logger, ctxKeyRequestID))

	s.router.Use(gorilla_handle.CORS(gorilla_handle.AllowedOrigins([]string{"*"})))

	s.router.HandleFunc("/users/create", handlers.HandelUserCreate(s.store)).Methods("POST")
	s.router.HandleFunc("/auth", handlers.SessionsCreate(s.store, s.sessionStore, sessionName)).Methods("POST")

	private := s.router.PathPrefix("/private").Subrouter()
	private.Use(middleware.AuthenticateUser(s.store, s.sessionStore, sessionName, ctxKeyUser))
	private.HandleFunc("/whoami", handlers.Whoami(ctxKeyUser))
}
