package teststore

import (
	"github.com/mouldykitz/duty_checklist/model"
	"github.com/mouldykitz/duty_checklist/store"
)

type Store struct {
	userRepository *UserRepository
}

func New() *Store {
	return &Store{}
}

// Не хотим, чтобы репозиторием пользовались в обход хранилища, поэтому создаем этот метод
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}

	return s.userRepository

	// store.User().Create()
}
