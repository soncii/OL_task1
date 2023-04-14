package model

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Title     string
	Author    string
}
type BookCreateReq struct {
	Title  string
	Author string
}

type BookCreateResp struct {
	BID       uint
	CreatedAt time.Time
}
