package model

import (
	"login/entities"
	"time"
)

type TransactionCreateReq struct {
	UID   uint
	Books []*entities.TransactionDetails
}
type TransactionCreateResp struct {
	TID       uint
	CreatedAt time.Time
}
