package model

import "time"

type BookCreateReq struct {
	Title  string
	Author string
}

type BookCreateResp struct {
	BID       uint
	CreatedAt time.Time
}
