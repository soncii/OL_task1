package model

import (
	"github.com/google/uuid"
	"time"
)

type UserCreateReq struct {
	Name     string
	Email    string
	Password string
}

type CreateResp struct {
	UID       uuid.UUID
	CreatedAt time.Time
}
