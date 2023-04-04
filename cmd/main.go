package main

import (
	"context"
	"login/Storage"
	"login/config"
	"login/service"
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
	st := Storage.NewStorage(cfg)
	s := service.NewService(st)
	h := transport.NewHandler(s)
	server := transport.NewServer(cfg, h)
	server.StartHTTPServer(ctx)

}
