package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/canvix/service/group"
)

type GroupController struct {
	CanvixController
}

func registerGroupRoutes(router *gin.RouterGroup, loggedMiddleware, workspaceMiddleware gin.HandlerFunc) {
	groupController := GroupController{}

	groupRoutes := router.Group("group").Use(loggedMiddleware, workspaceMiddleware)

	groupRoutes.POST("list", groupController.list)
	groupRoutes.POST("create", groupController.create)
	groupRoutes.POST("update", groupController.rename)
	groupRoutes.DELETE("delete", groupController.delete)
	groupRoutes.POST("move", groupController.move)
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
	req := group.UpdateAssetGroupReq{}
	workspaceId, err := c.BindRequestJson(ctx, &req, "rename")
	if err != nil {
		return
	}
	if err := group.Rename(ctx, workspaceId, &req); err != nil {
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

func (c *GroupController) move(ctx *gin.Context) {
	req := group.BatchMoveProjectGroupReq{}

	workspaceId, err := c.BindRequestJson(ctx, &req, "move group")

	if err != nil {
		return
	}

	if req.GroupCode == "" && req.GroupName == "" {
		c.AbortClientError(ctx, "move: groupCode or groupName is required")
		return
	}

	if err := group.BatchMoveGroup(ctx, workspaceId, &req); err != nil {
		c.AbortServerError(ctx, "move: "+err.Error())
		return
	}

	c.SuccessVoid(ctx)
}
