package main

import (
	"context"
	"transactions/config"
	"transactions/service"
	"transactions/storage"
	"transactions/transport/http"
)

// @title Transaction API
// @version 1.0
// @description This is a microservice that manages financial transactions of book borrowing
// @termsOfService http://swagger.io/terms/

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
	transactionService := service.NewTransactionService(st)
	h := http.NewHandler(transactionService, cfg)
	server := http.NewServer(cfg, h)
	err = server.StartHTTPServer(ctx)
	if err != nil {
		return
	}

}
