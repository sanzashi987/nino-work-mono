package http

import (
	"github.com/cza14h/nino-work/pkg/auth"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.Use(auth.ValidateMiddleware())

	controller := &ChatController{}

	v1 := router.Group("backend/v1")
	{
		v1.POST("chat", controller.Chat)

	}
	return router
}
