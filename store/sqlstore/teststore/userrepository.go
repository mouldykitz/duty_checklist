package teststore

import (
	"github.com/mouldykitz/duty_checklist/model"
	"github.com/mouldykitz/duty_checklist/store"
)

type UserRepository struct {
	store *Store
	users map[int]*model.User
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	//r.users[u.Email] = u
	//u.ID = len(r.users)

	u.ID = len(r.users) + 1
	r.users[u.ID] = u

	return nil
}

// Find by id
func (r *UserRepository) Find(id int) (*model.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}

	return nil, store.ErrRecordNotFound
	// u, ok := r.users[email]
	// if !ok {
	// 	return nil, store.ErrRecordNotFound
	// }

	// return u, nil
}
