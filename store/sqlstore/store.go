package sqlstore

import (
	"database/sql"

	_ "github.com/lib/pq" // драйвер для postgres
	"github.com/mouldykitz/duty_checklist/store"
)

type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository

	// store.User().Create()
}
