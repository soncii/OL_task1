package transport

func (s Server) SetupRoutes() {
	v1 := s.App.Group("/api/v1")
	v1.POST("/user", s.handler.CreateUser)
}
