package http

import (
	"net/http"

	"github.com/cza14h/nino-work/apps/canvas-pro/http/middleware"
	"github.com/cza14h/nino-work/apps/canvas-pro/service"
	"github.com/cza14h/nino-work/pkg/auth"
	"github.com/gin-gonic/gin"
)

func getWorkspaceCode(ctx *gin.Context) string {
	return ctx.GetHeader("Projectcode")
}

func getCurrentUser(ctx *gin.Context) uint64 {
	userId, _ := ctx.Get(auth.UserID)
	return userId.(uint64)
}

func UserWorkspace(ctx *gin.Context) {
	userId, workspaceCode := getCurrentUser(ctx), getWorkspaceCode(ctx)

	if service.UserServiceImpl.ValidateUserWorkspace(ctx, userId, workspaceCode) {
		ctx.Next()
	} else {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "User does not has the access to the given workspace",
			"data": nil,
		})
	}
}

func NewRouter(loginPageUrl string) *gin.Engine {
	router := gin.Default()

	root := router.Group("enc-oss-canvas/V1")
	canvasAuthMiddleWare := []gin.HandlerFunc{
		middleware.CanvasUserLoggedIn(loginPageUrl),
		UserWorkspace,
	}

	{
		loginGroup := root.Group(login_prefix)
		loginGroup.POST("login", loginController.login)
		loginGroup.GET("login-verification/get-uuidkey", loginController.getUuid)
		loginGroup.GET("login-verification/get-verification-code", loginController.getVerifyCode)
		loginGroup.GET("logout", loginController.logout)
	}

	{
		commonRoutes := root.Group(common_prefix).Use(canvasAuthMiddleWare...)
		commonRoutes.POST("search", commonController.searchComponents)
		commonRoutes.GET("userInfo", commonController.getUserInfo)
	}

	{
		dataSourceRoutes := root.Group(data_source_prefix).Use(canvasAuthMiddleWare...)

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
		projectScreenRoutes := root.Group(project_prefix).Use(canvasAuthMiddleWare...)

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

		groupedProjectRoutes := root.Group(grouped_project_prefix).Use(canvasAuthMiddleWare...)

		groupedProjectRoutes.POST("list", groupController.list)
		groupedProjectRoutes.POST("create", groupController.createProjectGroup)
		groupedProjectRoutes.POST("update", groupController.rename)
		groupedProjectRoutes.DELETE("delete", groupController.delete)
		// for adapt
		projectScreenRoutes.POST("move", groupController.move)

	}

	{
		assetRoutes := root.Group(asset_prefix).Use(canvasAuthMiddleWare)
		assetRoutes.POST("selectMyAssets", assetController.list)
		assetRoutes.POST("updateMyAssetsName", assetController.update)
		assetRoutes.POST("updateAssetsGroup", assetController.update)
		assetRoutes.DELETE("deleteAssets", assetController.delete)
		assetRoutes.POST("upload", assetController.upload)
		assetRoutes.POST("detail", assetController.read)
		assetRoutes.POST("replace", assetController.replace)
		assetRoutes.POST("loadAsset", assetController.download)
		assetRoutes.POST("importAsset", assetController._import)

		assetRoutes.POST("addGroup", groupController.createDesginGroup)
		assetRoutes.GET("deleteGroup", groupController.delete)
		assetRoutes.POST("updateGroupsName", groupController.rename)
		assetRoutes.POST("selectGroup", groupController.list)
	}

	{
		themeRoutes := root.Group(theme_prefix).Use(canvasAuthMiddleWare)

		themeRoutes.POST("list", themeController.list)
		themeRoutes.POST("create", themeController.create)
		themeRoutes.POST("update", themeController.update)
		themeRoutes.DELETE("delete", themeController.delete)

	}

	return router
}
