package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"tg_app_backend/internal/storage/postgresql"
)

//// @BasePath /api/v1
//

type Server struct {
	db     *postgresql.Postgresql
	router *gin.Engine
}

func New(db *postgresql.Postgresql) *Server {
	return &Server{db: db, router: gin.New()}
}

func (s *Server) Start() error {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization", "ngrok-skip-browser-warning")
	s.router.Use(cors.New(corsConfig))
	s.routerSettings()
	err := s.router.Run(":4040")
	return err
}
