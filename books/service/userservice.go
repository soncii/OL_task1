package service

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"login/config"
	"login/model"
	"login/storage"
	"time"
)

type UserService struct {
	r        *storage.Storage
	hashCost int
}
type IUserService interface {
	Create(ctx context.Context, req model.UserCreateReq) (model.UserCreateResp, error)
	UpdatePassword(ctx context.Context, req model.UserEmailPassReq) (resp model.UserChangePassResp, err1 error)
	Validate(ctx context.Context, req model.UserCreateReq) (bool, error)
	GetUserRecords(ctx context.Context, email string) (model.RecordGetCurrentUserResp, error)
	GetUsers(ctx context.Context) (model.GetUsersResp, error)
	GetUsersLastMonth(ctx context.Context) (model.GetUsersWithRecordResp, error)
	Delete(ctx context.Context, req model.UserDeleteReq) error
}

func NewUserService(r *storage.Storage, cfg *config.Config) *UserService {
	return &UserService{r: r, hashCost: cfg.HashCost}
}

func (s *UserService) Create(ctx context.Context, req model.UserCreateReq) (model.UserCreateResp, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), s.hashCost)
	if err != nil {
		print(err)
	}
	u := model.User{Email: req.Email, Name: req.Name, Password: hashed}
	err = s.r.UserRepo.CreateUser(ctx, &u)
	if err != nil {
		return model.UserCreateResp{}, err
	}
	return model.UserCreateResp{UID: u.ID, CreatedAt: u.CreatedAt}, nil
}
func (s *UserService) UpdatePassword(ctx context.Context, req model.UserEmailPassReq) (resp model.UserChangePassResp, err1 error) {
	u, err := s.r.UserRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return model.UserChangePassResp{}, err
	}
	u.Password, err = bcrypt.GenerateFromPassword([]byte(req.Password), s.hashCost)
	if err != nil {
		return model.UserChangePassResp{}, err
	}
	err = s.r.UserRepo.Update(ctx, u)
	if err != nil {
		return model.UserChangePassResp{}, err
	}
	return model.UserChangePassResp{UpdatedAt: u.UpdatedAt}, nil
}

func (s *UserService) Validate(ctx context.Context, req model.UserCreateReq) (bool, error) {
	u, err := s.r.UserRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return false, err
	}
	err = bcrypt.CompareHashAndPassword(u.Password, []byte(req.Password))
	if err != nil {
		return false, err
	}
	return true, nil
}
func (s *UserService) GetUserRecords(ctx context.Context, email string) (model.RecordGetCurrentUserResp, error) {
	records, err := s.r.UserRepo.GetUserRecords(ctx, email)
	if err != nil {
		return model.RecordGetCurrentUserResp{}, err
	}
	return model.RecordGetCurrentUserResp{Records: records}, nil
}

func (s *UserService) GetUsers(ctx context.Context) (model.GetUsersResp, error) {
	Users, err := s.r.UserRepo.GetUsers(ctx)
	if err != nil {
		return model.GetUsersResp{}, err
	}
	res := make([]*model.UserWithBooks, 0)
	for _, user := range Users {
		res = append(res, user.ConvertToUserWithBooks())
	}
	return model.GetUsersResp{Users: res}, nil

}
func (s *UserService) GetUsersLastMonth(ctx context.Context) (model.GetUsersWithRecordResp, error) {
	monthAgo := time.Now().AddDate(0, -1, 0).Format("2006-01-02 15:04:05.999999 -07:00")
	Users, err := s.r.UserRepo.GetUsersLastMonth(ctx, monthAgo)
	if err != nil {
		return model.GetUsersWithRecordResp{}, err
	}
	UsersWithRecord := make([]model.UserWithRecord, 0)
	fmt.Println(len(Users))
	for _, u := range Users {
		fmt.Println(u)
		UsersWithRecord = append(UsersWithRecord, model.UserWithRecord{ID: u.ID, Name: u.Name, Email: u.Email, Count: len(u.Records)})
	}
	return model.GetUsersWithRecordResp{Users: UsersWithRecord}, nil
}
func (s *UserService) Delete(ctx context.Context, req model.UserDeleteReq) error {
	return s.r.UserRepo.Delete(ctx, req.UID, req.Hard)
}
