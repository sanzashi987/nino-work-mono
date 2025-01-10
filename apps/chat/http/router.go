package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/pkg/controller"
)

func NewRouter(loginPageUrl string) *gin.Engine {
	router := gin.Default()

	router.Use(controller.ValidateMiddleware(loginPageUrl))

	controller := &ChatController{}

	v1 := router.Group("backend/v1")
	{
		v1.POST("chat", controller.Chat)
	}
	return router
}
