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
	userId := ctx.GetUint64(controller.UserID)
	return userId
}

func getUploadRpcService(ctx *gin.Context) storage.StorageService {
	rpcMap, _ := ctx.Get(RPCKEY)
	m, _ := rpcMap.(map[string]any)
	return m["storage"].(storage.StorageService)
}

func workspaceMiddleware(ctx *gin.Context) {
	userId := getCurrentUser(ctx)
	workspaceCode, _ := getWorkspaceCode(ctx)
	if service.ValidateUserWorkspace(ctx, userId, workspaceCode) {
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

	root.Use(mergeRpcMiddleware)
	registerCommonRoutes(root, loggedInMiddleware, workspaceMiddleware)
	registerDataSourceRoutes(root, loggedInMiddleware, workspaceMiddleware)
	registerProjectRoutes(root, loggedInMiddleware, workspaceMiddleware)
	registerGroupRoutes(root, loggedInMiddleware, workspaceMiddleware)
	registerAssetRoutes(root, loggedInMiddleware, workspaceMiddleware)
	registerThemeRoutes(root, loggedInMiddleware, workspaceMiddleware)

	return router
}
