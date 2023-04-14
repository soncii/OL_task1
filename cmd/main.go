package main

import (
	"context"
	"login/config"
	"login/service"
	"login/storage"
	"login/transport/http"
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
	m := service.NewManager(st, cfg)
	h := http.NewHandler(m, cfg)
	server := http.NewServer(cfg, h)
	err = server.StartHTTPServer(ctx)
	if err != nil {
		return
	}

}
