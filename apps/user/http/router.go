package http

import (
	// "github.com/gin-contrib/static"

	// "fmt"
	// "path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/pkg/auth"
	// "github.com/sanzashi987/nino-work/pkg/utils"
)

func NewRouter(loginPageUrl string) *gin.Engine {
	apiEngine := gin.Default()

	userController := UserController{}
	appController := AppController{}

	// router.Use(static.Serve("/", static.LocalFile("./static/dist", true)))
	authMiddleware := auth.ValidateMiddleware(loginPageUrl)

	// appRoot := utils.GetAppRoot()
	// feRoot := filepath.Join(appRoot, "./apps/user/static/dist/")
	// fmt.Printf("static root: %s", feRoot)

	// apiEngine.Use(static.Serve("/", static.LocalFile(feRoot, true)))

	v1 := apiEngine.Group("/backend/v1")
	{
		v1.POST("login", userController.UserLogin)
		v1.POST("register", userController.UserRegister)
		authed := v1.Use(authMiddleware)
		authed.GET("token", userController.TestToken)
		authed.GET("info", userController.UserInfo)

		authed.GET("apps/list", appController.ListApps)
		authed.POST("apps/create", appController.ListApps)
	}

	return apiEngine
}
