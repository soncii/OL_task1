package model

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	UserID    uint
	BookID    uint
	Price     float64
}

type TransactionCreateReq struct {
	UID   uint
	BID   uint
	Price float64
}
type TransactionCreateResp struct {
	TID uint
}
