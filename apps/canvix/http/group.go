package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/apps/canvix/service/group"
)

type GroupController struct {
	CanvixController
}

func registerGroupRoutes(router *gin.RouterGroup, loggedMiddleware, workspaceMiddleware gin.HandlerFunc) {
	groupController := GroupController{}

	groupRoutes := router.Group("group")
	groupRoutes.Use(loggedMiddleware, workspaceMiddleware)
	{
		groupRoutes.POST("list", groupController.list)
		groupRoutes.POST("create", groupController.create)
		groupRoutes.POST("update", groupController.rename)
		groupRoutes.DELETE("delete", groupController.delete)
	}
}

func (c *GroupController) list(ctx *gin.Context) {

	req := group.ListGroupReq{}
	workspaceId, err := c.BindRequestJson(ctx, &req, "list")
	if err != nil {
		return
	}

	output, err := group.List(ctx, workspaceId, &req)
	if err != nil {
		c.AbortServerError(ctx, "[http] list group error "+err.Error())
		return
	}
	c.ResponseJson(ctx, output)

}

/*CRUD*/

func (c *GroupController) create(ctx *gin.Context) {
	req := group.CreateAssetGroupReq{}
	workspaceId, err := c.BindRequestJson(ctx, &req, "create")
	if err != nil {
		return
	}

	if _, err := group.Create(ctx, workspaceId, &req); err != nil {
		c.AbortClientError(ctx, "[http] create group error "+err.Error())
		return
	}
	c.SuccessVoid(ctx)
}

// rename
func (c *GroupController) rename(ctx *gin.Context) {
	req := &group.UpdateAssetGroupReq{}
	workspaceId, err := c.BindRequestJson(ctx, &req, "rename")
	groupTypeTag, _ := consts.GetGroupTypeTagFromBasic(req.TypeTag)
	if err != nil {
		return
	}
	if err := group.Rename(ctx, workspaceId, req.GroupCode, req.GroupName, groupTypeTag); err != nil {
		c.AbortClientError(ctx, err.Error())
		return
	}

	c.SuccessVoid(ctx)
}

func (c *GroupController) delete(ctx *gin.Context) {
	req := &group.DeleteAssetGroupReq{}
	workspaceId, err := c.BindRequestJson(ctx, &req, "delete")
	if err != nil {
		return
	}

	if err := group.Delete(ctx, workspaceId, req); err != nil {
		c.AbortServerError(ctx, err.Error())
		return
	}
	c.SuccessVoid(ctx)
}

type MoveProjectReq struct {
	group.DeleteAssetGroupReq
	Ids []string `json:"fileIds" binding:"required"`
}
