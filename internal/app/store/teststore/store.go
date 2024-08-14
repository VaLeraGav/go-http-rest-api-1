package teststore

import (
	"github.com/VaLeraGav/go-http-rest-api-1/internal/app/model"
	"github.com/VaLeraGav/go-http-rest-api-1/internal/app/store"
)

type Store struct {
	userRepository *UserRepository
}

// New …
func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.userRepository
}
