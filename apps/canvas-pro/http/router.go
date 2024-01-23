package http

import (
	"github.com/cza14h/nino-work/apps/canvas-pro/http/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(loginPageUrl string) *gin.Engine {
	router := gin.Default()

	root := router.Group("enc-oss-canvas/V1")
	canvasAuthMiddleWare := middleware.CanvasMiddleware(loginPageUrl)
	loginGroup := root.Group(login_group)
	{
		loginController := &LoginController{}
		loginGroup.POST("login", loginController.login)
		loginGroup.GET("login-verification/get-uuidkey", loginController.getUuid)
		loginGroup.GET("login-verification/get-verification-code", loginController.getVerifyCode)
		loginGroup.GET("logout", loginController.logout)
	}

	commonGroup := root.Group(common_group).Use(canvasAuthMiddleWare)
	{
		commonController := &CommonController{}
		commonGroup.POST("search", commonController.searchComponents)
		commonGroup.GET("userInfo", commonController.getUserInfo)
	}

	dataSourceGroup := root.Group(data_source_group).Use(canvasAuthMiddleWare)
	{
		dataSourceController := &DataSourceController{}

		dataSourceGroup.POST("create", dataSourceController.createDataSource)
		dataSourceGroup.GET("info/:sourceId", dataSourceController.readDataSource)
		dataSourceGroup.DELETE("delete", dataSourceController.deleteDataSource)
		dataSourceGroup.POST("update", dataSourceController.updateDataSource)
		dataSourceGroup.POST("replaceIp", dataSourceController.replaceIp)
		dataSourceGroup.POST("list-all", dataSourceController.queryDataSourceList)
		dataSourceGroup.POST("list-page", dataSourceController.queryDataSourceList)
		dataSourceGroup.POST("searchByIp", dataSourceController.queryDataSourceList)

	}

	return router
}
