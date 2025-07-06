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

const WORKSPACE_ID = "workspaceId"
const WORKSPACE_CODE = "workspaceCode"

func getWorkspaceCode(ctx *gin.Context) (string, uint64, error) {
	workspaceCode := ctx.GetHeader("workspace")
	workspaceId, _, err := consts.GetIdFromCode(workspaceCode)
	return workspaceCode, workspaceId, err
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
	workspaceCode, workspaceId, err := getWorkspaceCode(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "Fail to parse workspace code",
			"data": nil,
		})
		return
	}
	if service.ValidateUserWorkspace(ctx, userId, workspaceCode) {
		ctx.Set(WORKSPACE_CODE, workspaceCode)
		ctx.Set(WORKSPACE_ID, workspaceId)
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

	loggedInMiddleware := middleware.CanvixUserLoggedIn(loginPageUrl)

	root.Use(mergeRpcMiddleware)
	registerCommonRoutes(root, loggedInMiddleware, workspaceMiddleware)
	registerDataSourceRoutes(root, loggedInMiddleware, workspaceMiddleware)
	registerProjectRoutes(root, loggedInMiddleware, workspaceMiddleware)
	registerGroupRoutes(root, loggedInMiddleware, workspaceMiddleware)
	registerAssetRoutes(root, loggedInMiddleware, workspaceMiddleware)
	registerThemeRoutes(root, loggedInMiddleware, workspaceMiddleware)

	return router
}
