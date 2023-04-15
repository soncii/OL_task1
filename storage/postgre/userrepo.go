package postgre

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"login/model"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	u := &model.User{}
	return u, r.db.Where("email=?", email).First(u).Error

}
func (r *UserRepository) GetUserByID(ctx context.Context, UID uint) (*model.User, error) {
	u := &model.User{}
	return u, r.db.Where("id=?", UID).First(u).Error

}

func (r *UserRepository) CreateUser(ctx context.Context, user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) Delete(ctx context.Context, uid uint, hard bool) error {
	if hard {
		return r.db.Unscoped().Delete(&model.User{}, uid).Error
	}
	return r.db.Delete(&model.User{}, uid).Error
}

func (r *UserRepository) Update(ctx context.Context, user *model.User) error {
	err := r.db.First(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserRecords(ctx context.Context, email string) ([]*model.Record, error) {
	var user model.User
	err := r.db.Model(&model.User{}).Preload("Records").Where("email=?", email).First(&user).Error
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	return user.Records, nil
}

func (r *UserRepository) GetUsers(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	err := r.db.Model(&model.User{}).Preload("Records", "borrowed=true").
		Preload("Records.Book").Find(&users).Error
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	return users, nil
}
func (r *UserRepository) GetUsersLastMonth(ctx context.Context, time string) ([]*model.User, error) {
	var users []*model.User
	err := r.db.Model(&model.User{}).
		Preload("Records", fmt.Sprintf("taken_at>'%v'", time)).
		Find(&users).Error
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	return users, err
}
