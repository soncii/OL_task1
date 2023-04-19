package http

import (
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "transactions/docs"
)

func (s *Server) SetupRoutes() {
	v1 := s.App.Group("/api/v1")

	v1.GET("/transaction/:tid", s.handler.GetTransaction)
	v1.POST("/transaction", s.handler.CreateTransaction)

	v1.GET("/swagger/*", echoSwagger.WrapHandler)
}
