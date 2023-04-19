package service

import (
	"login/config"
	"login/storage"
)

type Manager struct {
	UserService   IUserService
	BookService   IBookService
	RecordService IRecordService
}

func NewManager(st *storage.Storage, cfg *config.Config) *Manager {
	userService := NewUserService(st, cfg)
	bookService := NewBookService(st)
	recordService := NewRecordService(st)
	m := &Manager{UserService: userService, BookService: bookService, RecordService: recordService}
	return m
}
