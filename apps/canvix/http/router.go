package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/apps/canvix/http/middleware"
	"github.com/sanzashi987/nino-work/apps/canvix/service"
	"github.com/sanzashi987/nino-work/pkg/controller"
	"github.com/sanzashi987/nino-work/proto/storage"
)

const RPCKEY = "RPCCLIENTS"

func getWorkspaceCode(ctx *gin.Context) (string, uint64) {
	workspaceCode := ctx.GetHeader("Projectcode")
	workspaceId, _, _ := consts.GetIdFromCode(workspaceCode)
	return workspaceCode, workspaceId
}

func getCurrentUser(ctx *gin.Context) uint64 {
	userId, _ := ctx.Get(controller.UserID)
	return userId.(uint64)
}

func getUploadRpcService(ctx *gin.Context) storage.StorageService {
	rpcMap, _ := ctx.Get(RPCKEY)
	m, _ := rpcMap.(map[string]any)
	return m["storage"].(storage.StorageService)
}

func workspaceMiddleware(ctx *gin.Context) {
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

	root := router.Group("/backend/canvix/v1")

	mergeRpcMiddleware := func(ctx *gin.Context) {
		ctx.Set(RPCKEY, rpcServices)
	}

	loggedInMiddleware := middleware.CanvasUserLoggedIn(loginPageUrl)

	canvasAuthMiddleWare := []gin.HandlerFunc{
		loggedInMiddleware,
		workspaceMiddleware,
	}

	root.Use(mergeRpcMiddleware)

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
		dataSourceRoutes.POST("searchByIp", dataSourceController.search)
	}

	registerProjectRoutes(root, loggedInMiddleware, workspaceMiddleware)
	{
		groupedProjectRoutes := root.Group(grouped_project_prefix).Use(canvasAuthMiddleWare...)
		groupedProjectRoutes.POST("list", groupController.listProjectGroup)
		groupedProjectRoutes.POST("create", groupController.createProjectGroup)
		groupedProjectRoutes.POST("update", groupController.projectRename)
		groupedProjectRoutes.DELETE("delete", groupController.deleteProjectGroup)

	}

	registerAssetRoutes(root, loggedInMiddleware, workspaceMiddleware)

	{
		assetRoutes.POST("addGroup", groupController.createDesginGroup)
		assetRoutes.GET("deleteGroup", groupController.deleteAssetGroup)
		assetRoutes.POST("updateGroupsName", groupController.assetRename)
		assetRoutes.POST("selectGroup", groupController.listAssetGroup)
	}

	registerThemeRoutes(root, loggedInMiddleware, workspaceMiddleware)

	return router
}
