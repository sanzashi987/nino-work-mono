package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/apps/canvix/service"
	"github.com/sanzashi987/nino-work/pkg/shared"
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

type ListGroupReq struct {
	GroupCode string `json:"code"`
	GroupName string `json:"name"`
	TypeTag   string `json:"type" binding:"required"`
	shared.PaginationRequest
}

func (c *GroupController) list(ctx *gin.Context) {

	req := ListGroupReq{}
	workspaceId, err := c.BindRequestJson(ctx, &req, "list")
	if err != nil {
		return
	}

	output, err := service.GroupServiceImpl.ListGroups(ctx, workspaceId, req.GroupName, typeTag)
	if err != nil {
		c.AbortServerError(ctx, "[http] list group error "+err.Error())
		return
	}
	c.ResponseJson(ctx, output)

}

/*CRUD*/
type CreateAssetGroupReq struct {
	GroupName string `json:"name" binding:"required"`
	TypeTag   string `json:"type" binding:"required"`
}

func (c *GroupController) create(ctx *gin.Context) {
	req := CreateAssetGroupReq{}
	workspaceId, err := c.BindRequestJson(ctx, &req, "create")
	if err != nil {
		return
	}

	if _, err := service.GroupServiceImpl.Create(ctx, workspaceId, req.GroupName, typeTag); err != nil {
		c.AbortClientError(ctx, "[http] create group error "+err.Error())
		return
	}
	c.SuccessVoid(ctx)
}

type UpdateAssetGroupReq struct {
	CreateAssetGroupReq
	DeleteAssetGroupReq
}

// rename
func (c *GroupController) rename(ctx *gin.Context) {
	groupTypeTag, _ := consts.GetGroupTypeTagFromBasic(typeTag)
	req := &UpdateAssetGroupReq{}
	workspaceId, err := c.BindRequestJson(ctx, &req, "rename")
	if err != nil {
		return
	}
	if err := service.GroupServiceImpl.Rename(ctx, workspaceId, req.GroupCode, req.GroupName, groupTypeTag); err != nil {
		c.AbortClientError(ctx, err.Error())
		return
	}

	c.SuccessVoid(ctx)
}

type DeleteAssetGroupReq struct {
	GroupCode string `json:"code" binding:"required"`
	TypeTag   string `json:"type" binding:"required"`
}

func (c *GroupController) delete(ctx *gin.Context) {
	req := &DeleteAssetGroupReq{}
	workspaceId, err := c.BindRequestJson(ctx, &req, "delete")
	if err != nil {
		return
	}

	if err := service.GroupServiceImpl.Delete(ctx, workspaceId, req.GroupCode, typeTag); err != nil {
		c.AbortServerError(ctx, err.Error())
		return
	}
	c.SuccessVoid(ctx)
}

type MoveProjectReq struct {
	DeleteAssetGroupReq
	Ids []string `json:"fileIds" binding:"required"`
}
