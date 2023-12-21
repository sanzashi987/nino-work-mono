package http

import (
	"github.com/gin-gonic/gin"
)


func NewRouter() *gin.Engine {
	router := gin.Default()

	var userController = UserController{}
	v1 := router.Group("backend/v1")
	{
		v1.POST("login", userController.UserLogin)
		v1.POST("register", userController.UserRegister)
	}

	return router
}

