package http

import (
	// "github.com/gin-contrib/static"

	"path/filepath"

	"github.com/cza14h/nino-work/pkg/utils"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	userController := UserController{}

	// router.Use(static.Serve("/", static.LocalFile("./static/dist", true)))
	v1 := router.Group("backend/v1")
	{
		v1.POST("login", userController.UserLogin)
		v1.POST("register", userController.UserRegister)
	}
	appRoot := utils.GetAppRoot()
	feRoot := filepath.Join(appRoot, "./static/dist/")
	router.Static("/", feRoot)
	router.NoRoute(func(ctx *gin.Context) {
		ctx.File(filepath.Join(feRoot, "index.html"))
	})

	return router
}
