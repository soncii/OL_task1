package entities

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
}

type TransactionDetails struct {
	ID        uint           `gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	BookID    uint           `gorm:"primaryKey;autoIncrement:false"`
	Price     float64
	Amount    uint
}
