package model

import (
	"time"

	"gorm.io/gorm"
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
type RecordWithTransaction struct {
	ID            uint
	UserID        uint
	BookID        uint
	TransactionID uint
	TakenAt       time.Time
	ReturnedAt    time.Time
	Borrowed      bool
	Price         float64
}

func (r *Record) MapWithPrice(price float64) *RecordWithTransaction {
	return &RecordWithTransaction{
		ID:            r.ID,
		UserID:        r.UserID,
		BookID:        r.BookID,
		TransactionID: r.TransactionID,
		TakenAt:       r.TakenAt,
		ReturnedAt:    r.ReturnedAt,
		Borrowed:      r.Borrowed,
		Price:         price,
	}
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
type TransactionGetResp struct {
	ID     uint
	BookID uint
	Price  float64
	UserID uint
}
type RecordCreateResp struct {
	RID       uint
	CreatedAt time.Time
}
