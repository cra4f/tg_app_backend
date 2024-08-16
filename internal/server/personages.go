package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary get personages
// @Schemes
// @Description Получить список всех персонажей игры
// @Tags personages
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /personages [get]
func (s *Server) getPersonages(c *gin.Context) {
	initData, ok := ctxInitData(c.Request.Context())
	fmt.Println("getPersonages", initData)
	if !ok {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "Init data not found",
		})
		return
	}
	if personages, err := s.db.GetPersonages(); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"personages": personages,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}
