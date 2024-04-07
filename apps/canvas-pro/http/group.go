package http

import (
	"github.com/cza14h/nino-work/apps/canvas-pro/consts"
	"github.com/cza14h/nino-work/apps/canvas-pro/service"
	"github.com/cza14h/nino-work/pkg/controller"
	"github.com/gin-gonic/gin"
)

const grouped_project_prefix = "group"

type GroupController struct {
	controller.BaseController
}

var groupController = &GroupController{
	controller.BaseController{
		ErrorPrefix: "[http] canvas asset group handler ",
	},
}

func (c *GroupController) list(ctx *gin.Context) {
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
	workspaceCode := getWorkspaceCode(ctx)
	reqBody := CreateAssetGroupReq{}
	if err := ctx.BindJSON(&reqBody); err != nil {
		c.AbortClientError(ctx, "create: "+err.Error())
		return
	}
	if err := service.GroupServiceImpl.Create(ctx, reqBody.GroupName, workspaceCode, typeTag); err != nil {
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
	workspaceCode := getWorkspaceCode(ctx)
	reqBody := &UpdateAssetGroupReq{}

	if err := ctx.BindJSON(reqBody); err != nil {
		c.AbortClientError(ctx, err.Error())
		return
	}

	if err := service.GroupServiceImpl.Rename(ctx, workspaceCode, reqBody.GroupCode, reqBody.GroupName, groupTypeTag); err != nil {
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
	if err := ctx.BindJSON(reqBody); err != nil {
		c.AbortClientError(ctx, err.Error())
		return
	}
	workspaceCode := getWorkspaceCode(ctx)
	if err := service.GroupServiceImpl.Delete(ctx, reqBody.GroupCode, workspaceCode, typeTag); err != nil {
		c.AbortClientError(ctx, err.Error())
		return
	}
	c.SuccessVoid(ctx)
}

type MoveAssetReq struct {
	DeleteAssetGroupReq
	Ids []string `json:"codes" binding:"required"`
}
type MoveProjectReq struct {
	DeleteAssetGroupReq
	Ids []string `json:"fileIds" binding:"required"`
}

func (c *GroupController) move(ctx *gin.Context) {

}

func (c *GroupController) moveProject(ctx *gin.Context) {
}

func (c *GroupController) moveAsset(ctx *gin.Context) {
}
