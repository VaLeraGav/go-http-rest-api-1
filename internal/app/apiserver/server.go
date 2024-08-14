package apiserver

import (
	"net/http"

	"github.com/VaLeraGav/go-http-rest-api-1/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}

	s.configureRouting()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouting() {
	s.router.HandleFunc("/users", s.handelUserCreate()).Methods("POST")
}

func (s *server) handelUserCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
