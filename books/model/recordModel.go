package model

import (
	"gorm.io/gorm"
	"time"
)

type Record struct {
	ID            uint           `gorm:"primaryKey"`
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	UserID        uint
	BookID        uint
	TransactionID uint
	Book          Book `gorm:"foreignKey:BookID"`
	TakenAt       time.Time
	ReturnedAt    time.Time
	Borrowed      bool
}

type RecordGetCurrentUserReq struct {
	Email string
}

type RecordGetCurrentUserResp struct {
	Records []*Record
}

type RecordCreateReq struct {
	UID   uint
	BID   uint
	Price float64
}
type TransactionCreateReq struct {
	UID       uint
	BID       uint
	Price     float64
	CreatedAt time.Time
}
type TransactionCreateResp struct {
	TID uint
}
type RecordCreateResp struct {
	RID       uint
	CreatedAt time.Time
}
