package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/pkg/auth"
)

func NewRouter(loginPageUrl string) *gin.Engine {
	router := gin.Default()

	router.Use(auth.ValidateMiddleware(loginPageUrl))

	controller := &ChatController{}

	v1 := router.Group("backend/v1")
	{
		v1.POST("chat", controller.Chat)
	}
	return router
}
