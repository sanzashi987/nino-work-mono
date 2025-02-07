package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/pkg/controller"
)

func NewRouter(loginPageUrl string) *gin.Engine {
	apiEngine := gin.Default()

	authMiddleware := controller.ValidateMiddleware(loginPageUrl)

	v1 := apiEngine.Group("/backend/user/v1")
	{
		v1.POST("login", userController.UserLogin)
		v1.POST("register", userController.UserRegister)
		v1.GET("misc/importmap", miscController.GetImportMap)

		authed := v1.Use(authMiddleware)
		authed.GET("token", userController.TestToken)
		authed.GET("info", userController.UserInfo)

		authed.POST("apps/list", appController.ListApps)
		authed.POST("apps/create", appController.CreateApp)

		authed.GET("apps/list-permission", permissionController.ListPermissionsByApp)

		authed.POST("permission/create", permissionController.CreatePermission)
		authed.POST("permission/delete", permissionController.DeletePermission)
	}

	return apiEngine
}
