package entities

import "github.com/google/uuid"

type User struct {
	UID      uuid.UUID
	Name     string
	Email    string
	Password string
}

func NewUser() *User {
	return &User{}
}
