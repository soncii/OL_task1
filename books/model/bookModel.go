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
type BookRevenue struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Title     string
	BookID    uint
	Revenue   float64
}
type BookWithRevenue struct {
	ID      uint
	Title   string
	BookID  uint
	Revenue float64
}
type BookCreateReq struct {
	Title  string
	Author string
}
type BookGetByIDReq struct {
	BID uint
}
type BookGetByIDResp struct {
	BID    uint
	Title  string
	Author string
}
type BookCreateResp struct {
	BID       uint
	CreatedAt time.Time
}
