package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func New(router *gin.Engine) *Server {
	return &Server{
		router: router,
	}
}

func (s *Server) Start() {
	s.routerSettings()
}
