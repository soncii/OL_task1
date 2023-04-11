package entities

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

func NewBook() *Book {
	return &Book{}
}
