package model

import (
	"login/entities"
	"time"
)

type RecordGetCurrentUserReq struct {
	Email string
}

type RecordGetCurrentUserResp struct {
	Records []*entities.Record
}

type RecordCreateReq struct {
	UID string
	BID string
}
type RecordCreateResp struct {
	RID       uint
	CreatedAt time.Time
}
