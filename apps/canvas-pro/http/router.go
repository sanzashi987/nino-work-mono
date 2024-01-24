package http

import (
	"github.com/cza14h/nino-work/apps/canvas-pro/http/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(loginPageUrl string) *gin.Engine {
	router := gin.Default()

	root := router.Group("enc-oss-canvas/V1")
	canvasAuthMiddleWare := middleware.CanvasMiddleware(loginPageUrl)
	{
		loginGroup := root.Group(login_prefix)
		loginController := &LoginController{}

		loginGroup.POST("login", loginController.login)
		loginGroup.GET("login-verification/get-uuidkey", loginController.getUuid)
		loginGroup.GET("login-verification/get-verification-code", loginController.getVerifyCode)
		loginGroup.GET("logout", loginController.logout)
	}

	{
		commonRoutes := root.Group(common_prefix).Use(canvasAuthMiddleWare)
		commonController := &CommonController{}

		commonRoutes.POST("search", commonController.searchComponents)
		commonRoutes.GET("userInfo", commonController.getUserInfo)
	}

	{
		dataSourceRoutes := root.Group(data_source_prefix).Use(canvasAuthMiddleWare)
		dataSourceController := &DataSourceController{}

		dataSourceRoutes.POST("create", dataSourceController.create)
		dataSourceRoutes.GET("info/:sourceId", dataSourceController.read)
		dataSourceRoutes.POST("update", dataSourceController.update)
		dataSourceRoutes.DELETE("delete", dataSourceController.delete)
		dataSourceRoutes.POST("replaceIp", dataSourceController.replaceIp)
		dataSourceRoutes.POST("list-all", dataSourceController.list)
		dataSourceRoutes.POST("list-page", dataSourceController.list)
		dataSourceRoutes.POST("searchByIp", dataSourceController.list)
	}

	{
		projectScreenRoutes := root.Group(project_prefix).Use(canvasAuthMiddleWare)
		projectController := &ProjectController{}

		projectScreenRoutes.POST("create", projectController.create)
		projectScreenRoutes.POST("addByTemplate", projectController.create)
		projectScreenRoutes.GET("info/:id", projectController.read)
		projectScreenRoutes.POST("update", projectController.update)
		projectScreenRoutes.DELETE("delete", projectController.delete)
		projectScreenRoutes.POST("list", projectController.list)
		projectScreenRoutes.GET("copy/:id", projectController.duplicate)
		projectScreenRoutes.POST("publish", projectController.publish)
		projectScreenRoutes.POST("downloadScreen", projectController.export)
		projectScreenRoutes.POST("importScreen", projectController._import)
		projectScreenRoutes.POST("checkRef", projectController.getInteraction)

		groupedProjectRoutes := root.Group(grouped_project_prefix).Use(canvasAuthMiddleWare)
		groupedProjectController := &ProjectGroupController{}

		groupedProjectRoutes.POST("list", groupedProjectController.list)
		groupedProjectRoutes.POST("create", groupedProjectController.create)
		groupedProjectRoutes.POST("update", groupedProjectController.update)
		groupedProjectRoutes.DELETE("delete", projectController.delete)
		// for adapt
		projectScreenRoutes.POST("move", groupedProjectController.move)

	}

	return router
}
