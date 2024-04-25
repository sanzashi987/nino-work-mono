package http

import (
	"net/http"

	"github.com/cza14h/nino-work/apps/canvas-pro/consts"
	"github.com/cza14h/nino-work/apps/canvas-pro/http/middleware"
	"github.com/cza14h/nino-work/apps/canvas-pro/service"
	"github.com/cza14h/nino-work/pkg/auth"
	"github.com/cza14h/nino-work/proto/upload"
	"github.com/gin-gonic/gin"
)

const RPCKEY = "RPCCLIENTS"

func getWorkspaceCode(ctx *gin.Context) (string, uint64) {
	workspaceCode := ctx.GetHeader("Projectcode")
	workspaceId, _, _ := consts.GetIdFromCode(workspaceCode)
	return workspaceCode, workspaceId
}

func getCurrentUser(ctx *gin.Context) uint64 {
	userId, _ := ctx.Get(auth.UserID)
	return userId.(uint64)
}

func getUploadRpcService(ctx *gin.Context) upload.FileUploadService {
	rpcMap, _ := ctx.Get(RPCKEY)
	m, _ := rpcMap.(map[string]any)
	return m["upload"].(upload.FileUploadService)
}

func UserWorkspace(ctx *gin.Context) {
	userId := getCurrentUser(ctx)
	workspaceCode, _ := getWorkspaceCode(ctx)
	if service.UserServiceImpl.ValidateUserWorkspace(ctx, userId, workspaceCode) {
		ctx.Next()
	} else {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "Current user does not have the access right to the given workspace",
			"data": nil,
		})
	}
}

func NewRouter(loginPageUrl string, rpcServices map[string]any) *gin.Engine {
	router := gin.Default()

	root := router.Group("enc-oss-canvas/V1")

	mergeRpcMiddleware := func(ctx *gin.Context) {
		ctx.Set(RPCKEY, rpcServices)
	}

	canvasAuthMiddleWare := []gin.HandlerFunc{
		middleware.CanvasUserLoggedIn(loginPageUrl),
		UserWorkspace,
	}

	root.Use(mergeRpcMiddleware)

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
		projectScreenRoutes.POST("move", projectController.moveGroup)

		groupedProjectRoutes := root.Group(grouped_project_prefix).Use(canvasAuthMiddleWare...)

		groupedProjectRoutes.POST("list", groupController.listProjectGroup)
		groupedProjectRoutes.POST("create", groupController.createProjectGroup)
		groupedProjectRoutes.POST("update", groupController.projectRename)
		groupedProjectRoutes.DELETE("delete", groupController.deleteProjectGroup)

	}

	{
		assetRoutes := root.Group(asset_prefix).Use(canvasAuthMiddleWare...)
		assetRoutes.POST("selectMyAssets", assetController.list)
		assetRoutes.POST("updateMyAssetsName", assetController.update)
		assetRoutes.POST("updateAssetsGroup", assetController.moveGroup)
		assetRoutes.DELETE("deleteAssets", assetController.delete)
		assetRoutes.POST("upload", assetController.upload)
		assetRoutes.POST("detail", assetController.read)
		assetRoutes.POST("replace", assetController.replace)
		assetRoutes.POST("loadAsset", assetController.download)
		assetRoutes.POST("importAsset", assetController._import)

		assetRoutes.POST("addGroup", groupController.createDesginGroup)
		assetRoutes.GET("deleteGroup", groupController.deleteAssetGroup)
		assetRoutes.POST("updateGroupsName", groupController.assetRename)
		assetRoutes.POST("selectGroup", groupController.listAssetGroup)
	}

	{
		themeRoutes := root.Group(theme_prefix).Use(canvasAuthMiddleWare...)

		themeRoutes.POST("list", themeController.list)
		themeRoutes.POST("create", themeController.create)
		themeRoutes.POST("update", themeController.update)
		themeRoutes.DELETE("delete", themeController.delete)

	}

	return router
}
