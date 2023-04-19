package main

import (
	"context"
	"fmt"
	"log"
	"login/config"
	"login/service"
	"login/storage"
	"login/transport/http"
	"time"
)

// @title Library Management System API
// @description This API manages a library system, including users, books, and borrowing records.
// @version 1.0
// @contact Damir Gimaletdinov
// @host localhost:8080
// @BasePath /api/v1
func main() {
	fmt.Println(time.Now())
	log.Fatal(run())

}
func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}
	st, err := storage.NewStorage(cfg)
	if err != nil {
		return err
	}
	m := service.NewManager(st, cfg)
	h := http.NewHandler(m, cfg)
	server := http.NewServer(cfg, h)
	err = server.StartHTTPServer(ctx)
	if err != nil {
		return err
	}
	return nil
}
