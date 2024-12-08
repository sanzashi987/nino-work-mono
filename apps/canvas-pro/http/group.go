package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/canvas-pro/consts"
	"github.com/sanzashi987/nino-work/apps/canvas-pro/service"
)

const grouped_project_prefix = "group"

type GroupController struct {
	CanvasController
}

var groupController = &GroupController{
	CanvasController: createCanvasController("[http] canvas asset group handler "),
}

func (c *GroupController) listAssetGroup(ctx *gin.Context) {
	c.listByType(ctx, consts.DESIGN)
}

func (c *GroupController) listProjectGroup(ctx *gin.Context) {
	c.listByType(ctx, consts.PROJECT)
}

type ListGroupReq struct {
	GroupName string `json:"groupName"`
}

func (c *GroupController) listByType(ctx *gin.Context, typeTag string) {

	reqBody := ListGroupReq{}

	if workspaceId, err := c.BindRequestJson(ctx, &reqBody, "list"); err != nil {
		return
	} else {
		output, err := service.GroupServiceImpl.ListGroups(ctx, workspaceId, reqBody.GroupName, typeTag)
		if err != nil {
			c.AbortServerError(ctx, "list: "+err.Error())
			return
		}
		c.ResponseJson(ctx, output)
	}

}

/*CRUD*/
type CreateAssetGroupReq struct {
	GroupName string `json:"groupName" binding:"required"`
	//TypeTag string `json:"type" binding:"required"`
}

func (c *GroupController) createProjectGroup(ctx *gin.Context) {
	c.create(ctx, consts.PROJECT)
}
func (c *GroupController) createDesginGroup(ctx *gin.Context) {
	c.create(ctx, consts.DESIGN)
}

func (c *GroupController) create(ctx *gin.Context, typeTag string) {
	reqBody := CreateAssetGroupReq{}
	if workspaceId, err := c.BindRequestJson(ctx, &reqBody, "create"); err != nil {
		return
	} else if _, err := service.GroupServiceImpl.Create(ctx, workspaceId, reqBody.GroupName, typeTag); err != nil {
		c.AbortClientError(ctx, "create: "+err.Error())
		return
	}
	c.SuccessVoid(ctx)
}

type UpdateAssetGroupReq struct {
	CreateAssetGroupReq
	DeleteAssetGroupReq
}

func (c *GroupController) projectRename(ctx *gin.Context) {
	c.rename(ctx, consts.PROJECT)
}

func (c *GroupController) assetRename(ctx *gin.Context) {
	c.rename(ctx, consts.DESIGN)
}

// rename
func (c *GroupController) rename(ctx *gin.Context, typeTag string) {
	groupTypeTag, _ := consts.GetGroupTypeTagFromBasic(typeTag)
	reqBody := &UpdateAssetGroupReq{}

	if workspaceId, err := c.BindRequestJson(ctx, &reqBody, "rename"); err != nil {
		return
	} else if err := service.GroupServiceImpl.Rename(ctx, workspaceId, reqBody.GroupCode, reqBody.GroupName, groupTypeTag); err != nil {
		c.AbortClientError(ctx, err.Error())
		return
	}

	c.SuccessVoid(ctx)
}

type DeleteAssetGroupReq struct {
	GroupCode string `json:"groupCode" binding:"required"`
}

func (c *GroupController) deleteProjectGroup(ctx *gin.Context) {
	c.delete(ctx, consts.PROJECT)
}
func (c *GroupController) deleteAssetGroup(ctx *gin.Context) {
	c.delete(ctx, consts.DESIGN)
}

func (c *GroupController) delete(ctx *gin.Context, typeTag string) {
	reqBody := &DeleteAssetGroupReq{}
	if workspaceId, err := c.BindRequestJson(ctx, &reqBody, "delete"); err != nil {
		return
	} else if err := service.GroupServiceImpl.Delete(ctx, workspaceId, reqBody.GroupCode, typeTag); err != nil {
		c.AbortClientError(ctx, err.Error())
		return
	}
	c.SuccessVoid(ctx)
}

type MoveProjectReq struct {
	DeleteAssetGroupReq
	Ids []string `json:"fileIds" binding:"required"`
}
