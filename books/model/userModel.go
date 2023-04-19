package model

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Name      string
	Email     string `gorm:"uniqueIndex"`
	Password  []byte `json:"-"`
	Records   []*Record
}

type UserWithBooks struct {
	ID    uint
	Name  string
	Email string
	Books []*Book
}

func (u *User) ConvertToUserWithBooks() *UserWithBooks {
	Books := make([]*Book, 0)
	for _, r := range u.Records {
		Books = append(Books, &Book{ID: r.BookID, Title: r.Book.Title, Author: r.Book.Author})
		fmt.Println(r.Book)
	}

	return &UserWithBooks{ID: u.ID, Name: u.Name, Email: u.Email, Books: Books}
}

type UserCreateReq struct {
	Name     string
	Email    string
	Password string
}
type UserDeleteReq struct {
	UID  uint
	Hard bool
}
type UserCreateResp struct {
	UID       uint
	CreatedAt time.Time
}
type UserEmailPassReq struct {
	Email    string
	Password string
}
type UserChangePassResp struct {
	UpdatedAt time.Time
}
type GetUsersResp struct {
	Users []*UserWithBooks
}
type UserWithRecord struct {
	ID    uint
	Name  string
	Email string
	Count int
}
type GetUsersWithRecordResp struct {
	Users []UserWithRecord
}
