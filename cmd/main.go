package main

import (
	"context"
	"login/config"
	"login/service"
	"login/storage"
	"login/transport"
)

func main() {
	run()

}
func run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cfg, err := config.NewConfig()
	if err != nil {
		return
	}
	st := storage.NewStorage(cfg)
	userService := service.NewUserService(st, cfg)
	bookService := service.NewBookService(st)
	recordService := service.NewRecordService(st)
	transactionService := service.NewTransactionService(st)
	h := transport.NewHandler(userService, bookService, recordService, transactionService, cfg)
	server := transport.NewServer(cfg, h)
	err = server.StartHTTPServer(ctx)
	if err != nil {
		return
	}

}
