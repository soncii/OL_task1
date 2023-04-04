package service

import (
	"github.com/google/uuid"
	"login/Storage"
	"login/entities"
	"login/model"
	"time"
)

type UserService struct {
	r *Storage.Storage
}

func NewService(r *Storage.Storage) *UserService {
	return &UserService{r: r}
}
func (s *UserService) Get() {
}

func (s *UserService) Create(req model.UserCreateReq) model.CreateResp {
	uid := uuid.New()
	u := entities.User{UID: uid, Email: req.Email, Name: req.Name, Password: req.Password}
	s.r.Create(u)
	return model.CreateResp{UID: uid, CreatedAt: time.Now()}
}

func (*UserService) Delete() {

}
