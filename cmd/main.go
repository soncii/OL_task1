package main

import (
	"context"
	"log"
	"login/config"
	"login/service"
	"login/storage"
	"login/transport/http"
)

func main() {
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
