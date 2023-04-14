package http

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"login/config"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	cfg     *config.Config
	handler *Handler
	App     *echo.Echo
}

func NewServer(cfg *config.Config, handler *Handler) *Server {
	s := Server{cfg: cfg, handler: handler}
	return &s
}

func (s *Server) StartHTTPServer(ctx context.Context) error {
	s.App = echo.New()

	s.SetupRoutes()
	if err := s.App.Start(s.cfg.Port); err != http.ErrServerClosed {
		fmt.Printf("%v", err)
	}
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Print("gracefully shutting down the server")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.App.Shutdown(ctx); err != nil {
		s.App.Logger.Fatal(err)
	}
	return nil
}
