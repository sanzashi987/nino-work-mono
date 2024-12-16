package http

import (
	// "github.com/gin-contrib/static"

	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/pkg/auth"
	"github.com/sanzashi987/nino-work/pkg/utils"
)

func NewRouter(loginPageUrl string) *gin.Engine {
	router := gin.Default()

	userController := UserController{}

	// router.Use(static.Serve("/", static.LocalFile("./static/dist", true)))
	authMiddleware := auth.ValidateMiddleware(loginPageUrl)

	v1 := router.Group("backend/v1")
	{
		v1.POST("login", userController.UserLogin)
		v1.POST("register", userController.UserRegister)
		v1.GET("dashboard", userController.Dashboard).Use(authMiddleware)
	}
	appRoot := utils.GetAppRoot()
	feRoot := filepath.Join(appRoot, "./static/dist/")
	router.Static("/", feRoot)
	router.NoRoute(func(ctx *gin.Context) {
		ctx.File(filepath.Join(feRoot, "index.html"))
	})

	return router
}
