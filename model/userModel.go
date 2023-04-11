package model

import (
	"login/entities"
	"time"
)

type UserCreateReq struct {
	Name     string
	Email    string
	Password string
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
	Users []*entities.UserWithBooks
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
