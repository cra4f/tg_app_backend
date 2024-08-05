package server

func (s *Server) routerSettings() {
	s.personagesRoutes()
}

func (s *Server) personagesRoutes() {
	scoreboardRoutes := s.router.Group("personages")
	{
		scoreboardRoutes.GET("/personages", s.getRatings)
	}
}
