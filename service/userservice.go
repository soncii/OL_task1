package service

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"login/config"
	"login/entities"
	"login/model"
	"login/storage"
	"time"
)

type UserService struct {
	r        *storage.Storage
	hashCost int
}

func NewUserService(r *storage.Storage, cfg *config.Config) *UserService {
	return &UserService{r: r, hashCost: cfg.HashCost}
}
func (s *UserService) Get() {
}

func (s *UserService) Create(req model.UserCreateReq) model.UserCreateResp {
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), s.hashCost)
	if err != nil {
		print(err)
	}
	u := entities.User{Email: req.Email, Name: req.Name, Password: hashed}
	s.r.UserRepo.CreateUser(&u)
	return model.UserCreateResp{UID: u.ID, CreatedAt: u.CreatedAt}
}
func (s *UserService) UpdatePassword(req model.UserEmailPassReq) (resp model.UserChangePassResp, err1 error) {
	u := s.r.UserRepo.GetUserByEmail(req.Email)
	u.Password, _ = bcrypt.GenerateFromPassword([]byte(req.Password), s.hashCost)
	err := s.r.UserRepo.Update(u)
	if err != nil {
		return model.UserChangePassResp{}, err
	}
	return model.UserChangePassResp{UpdatedAt: u.UpdatedAt}, nil
}
func (s *UserService) Delete() {

}

func (s *UserService) Validate(req model.UserCreateReq) bool {
	u := s.r.UserRepo.GetUserByEmail(req.Email)
	err := bcrypt.CompareHashAndPassword(u.Password, []byte(req.Password))
	if err != nil {
		return false
	}
	return true
}
func (s *UserService) GetUserRecords(email string) model.RecordGetCurrentUserResp {
	records := s.r.UserRepo.GetUserRecords(email)
	return model.RecordGetCurrentUserResp{Records: records}
}

func (s *UserService) GetUsers() model.GetUsersResp {
	Users := s.r.UserRepo.GetUsers()
	res := make([]*entities.UserWithBooks, 0)
	for _, user := range Users {
		res = append(res, user.ConvertToUserWithBooks())
	}
	return model.GetUsersResp{Users: res}

}
func (s *UserService) GetUsersLastMonth() model.GetUsersWithRecordResp {
	monthAgo := time.Now().AddDate(0, -1, 0).Format("2006-01-02 15:04:05.999999 -07:00")
	Users := s.r.UserRepo.GetUsersLastMonth(monthAgo)
	UsersWithRecord := make([]model.UserWithRecord, 0)
	fmt.Println(len(Users))
	for _, u := range Users {
		fmt.Println(u)
		UsersWithRecord = append(UsersWithRecord, model.UserWithRecord{ID: u.ID, Name: u.Name, Email: u.Email, Count: len(u.Records)})
	}
	fmt.Println(UsersWithRecord)
	return model.GetUsersWithRecordResp{Users: UsersWithRecord}
}
