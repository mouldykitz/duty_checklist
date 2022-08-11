package store

import (
	"database/sql"

	_ "github.com/lib/pq" // драйвер для postgres
)

type Store struct {
	config         *Config
	db             *sql.DB
	userRepository *UserRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Инициализация хранилища и попытка подключения к БД
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL) //найти кака правильно указать драйвер для mysql
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

// Отключение от БД
func (s *Store) Close() {
	s.db.Close()
}

// Не хотим, чтобы репозиторием пользовались в обход хранилища, поэтому создаем этот метод
func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository

	// store.User().Create()
}
