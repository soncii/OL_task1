package postgre

import (
	"fmt"
	"gorm.io/gorm"
	"login/entities"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByEmail(email string) *entities.User {
	u := &entities.User{}
	r.db.Where("email=?", email).First(u)
	return u
}
func (r *UserRepository) GetUserByID(UID uint) *entities.User {
	u := &entities.User{}
	r.db.Where("id=?", UID).First(u)
	return u
}

func (r *UserRepository) CreateUser(user *entities.User) {
	r.db.Create(user)
}

func (r *UserRepository) Delete() {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepository) Update(user *entities.User) error {
	err := r.db.First(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserRecords(email string) []*entities.Record {
	var user entities.User
	err := r.db.Model(&entities.User{}).Preload("Records").Where("email=?", email).First(&user).Error
	if err != nil {
		fmt.Print(err)
	}
	return user.Records
}

func (r *UserRepository) GetUsers() []*entities.User {
	var users []*entities.User
	err := r.db.Model(&entities.User{}).Preload("Records", "borrowed=true").
		Preload("Records.Book").Find(&users).Error
	if err != nil {
		fmt.Print(err)
	}
	return users
}
func (r *UserRepository) GetUsersLastMonth(time string) []*entities.User {
	var users []*entities.User
	err := r.db.Model(&entities.User{}).
		Preload("Records", fmt.Sprintf("taken_at>'%v'", time)).
		Find(&users).Error
	if err != nil {
		fmt.Print(err)
	}
	return users
}
