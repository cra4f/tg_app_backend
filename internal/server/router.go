package server

import (
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"tg_app_backend/cmd/docs"
)

//func main() {
//	// Your secret bot token.
//	token := "1234567890:ABC"
//
//	r := gin.New()
//
//	r.Use(authMiddleware(token))
//	r.GET("/", showInitDataMiddleware)
//
//	if err := r.Run(":3000"); err != nil {
//		panic(err)
//	}
//}

func (s *Server) routerSettings() {
	docs.SwaggerInfo.BasePath = "/api/v1"

	// токен бота
	const token = "7402987260:AAH4Ps89Hsx9fjdPfkGFQE-LxPus7AG5vZQ"
	s.router.Use(s.authMiddleware(token))
	v1 := s.router.Group("/api/v1")
	{
		//v1.POST("/auth", s.)
		personageRoutes := v1.Group("/personages")
		{
			personageRoutes.GET("/", s.getPersonages)
		}

		//	exchangeGoldToCoin()
		//
	}
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
