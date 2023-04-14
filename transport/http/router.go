package http

func (s *Server) SetupRoutes() {
	v1 := s.App.Group("/api/v1")

	jwtGroup := v1.Group("", s.handler.jwtMiddleware)

	//ADDING TO DB
	jwtGroup.POST("/book", s.handler.CreateBook)
	jwtGroup.POST("/password", s.handler.ChangePassword)
	jwtGroup.POST("/record", s.handler.CreateRecord)

	jwtGroup.GET("/records", s.handler.GetRecords)
	jwtGroup.GET("/records/month", s.handler.GetUsersLastMonth)
	jwtGroup.GET("/users", s.handler.GetUsersWithBooks)

	v1.POST("/user", s.handler.CreateUser)
	v1.POST("/login", s.handler.Validate)
}
