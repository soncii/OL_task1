package Storage

import (
	"login/config"
	"login/entities"
)

type UserRepository struct {
	db []entities.User
}

func NewRepository(config *config.Config) *UserRepository {
	return &UserRepository{db: make([]entities.User, 0, config.Capacity)}
}

func (r *UserRepository) Get() {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepository) Create(user entities.User) {
	r.db = append(r.db, user)
}

func (r *UserRepository) Delete() {
	//TODO implement me
	panic("implement me")
}
