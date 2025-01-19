package http

import (
	// "github.com/gin-contrib/static"

	// "fmt"
	// "path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/pkg/controller"
	// "github.com/sanzashi987/nino-work/pkg/utils"
)

func NewRouter(loginPageUrl string) *gin.Engine {
	apiEngine := gin.Default()

	// router.Use(static.Serve("/", static.LocalFile("./static/dist", true)))
	authMiddleware := controller.ValidateMiddleware(loginPageUrl)

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
		authed.POST("apps/create", appController.CreateApp)

		authed.GET("apps/list-permission", permissionController.ListPermissionsByApp)
	}

	return apiEngine
}
