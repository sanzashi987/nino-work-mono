package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/canvix/service"
)

type CommonController struct {
	CanvixController
}

func registerCommonRoutes(router *gin.RouterGroup, loggedMiddleware, workspaceMiddleware gin.HandlerFunc) {

	commonController := CommonController{}
	nonAuthed := router.Group("common")
	authed := nonAuthed.Use(loggedMiddleware)

	authed.GET("user", commonController.getUserInfo)
	authed.GET("console", commonController.GetConsoleInfo)

}

// func (c *CommonController) searchComponents(ctx *gin.Context) {

// }
// TODO, call rpc from sso, to fetch the canvix permissions only
func (c *CommonController) getUserInfo(ctx *gin.Context) {

}

func (c *CommonController) GetConsoleInfo(ctx *gin.Context) {
	req := service.GetConsolenfoReq{}

	if err := ctx.ShouldBindUri(&req); err != nil {
		c.AbortClientError(ctx, "get workspace info error: "+err.Error())
		return
	}

	res, err := service.GetConsoleInfo(ctx, &req)
	if err != nil {
		c.AbortServerError(ctx, "get workspace info error: "+err.Error())
		return
	}
	c.ResponseJson(ctx, res)
}
