package Storage

import (
	"gorm.io/gorm"
	"login/config"
	"login/entities"
)

type IUserRepository interface {
	Get()
	Create(user entities.User)
	Delete()
}
type Storage struct {
	Pg *gorm.DB
	IUserRepository
}

func NewStorage(cfg *config.Config) *Storage {
	return &Storage{IUserRepository: NewRepository(cfg)}
}
