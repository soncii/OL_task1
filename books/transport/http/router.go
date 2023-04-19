package http

import (
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "login/docs"
)

func (s *Server) SetupRoutes() {
	v1 := s.App.Group("/api/v1")
	v1.GET("/swagger/*", echoSwagger.WrapHandler)
	jwtGroup := v1.Group("", s.handler.jwtMiddleware)

	//ADDING TO DB
	jwtGroup.POST("/book", s.handler.CreateBook)
	jwtGroup.PUT("/password", s.handler.ChangePassword)
	jwtGroup.POST("/borrow", s.handler.CreateRecord)

	jwtGroup.GET("/books/borrowed", s.handler.GetBorrowedBooks)
	jwtGroup.GET("/records", s.handler.GetRecords)
	jwtGroup.GET("/records/month", s.handler.GetUsersLastMonth)
	jwtGroup.GET("/users", s.handler.GetUsersWithBooks)

	v1.POST("/user", s.handler.CreateUser)
	v1.POST("/login", s.handler.Validate)
}
