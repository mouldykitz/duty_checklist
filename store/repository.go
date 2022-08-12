package store

import "github.com/mouldykitz/duty_checklist/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
