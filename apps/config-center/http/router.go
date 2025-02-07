package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/pkg/controller"
)

func NewRouter(loginPageUrl string) *gin.Engine {
	apiEngine := gin.Default()

	authMiddleware := controller.ValidateMiddleware(loginPageUrl)

	v1 := apiEngine.Group("/backend/config-center/v1")
	{
		v1.Use(authMiddleware)
	}

	return apiEngine
}
