package http

import (
	"github.com/cza14h/nino-work/apps/canvas-pro/consts"
	"github.com/cza14h/nino-work/apps/canvas-pro/service"
	"github.com/cza14h/nino-work/pkg/controller"
	"github.com/gin-gonic/gin"
)

const assetGroupHandlerMessage = "[http] canvas asset group handler "
const grouped_project_prefix = "group"

type GroupController struct {
	controller.BaseController
}

var groupController = &GroupController{}

func (c *GroupController) list(ctx *gin.Context) {
}

/*CRUD*/
type CreateAssetGroupReq struct {
	GroupName string `json:"groupName" binding:"required"`
	//TypeTag string `json:"type" binding:"required"`
}

// func (c *GroupController) create(ctx *gin.Context) {

// }

func (c *GroupController) createProjectGroup(ctx *gin.Context) {
	c.create(ctx, consts.PROJECT)
}
func (c *GroupController) createDesginGroup(ctx *gin.Context) {
	c.create(ctx, consts.DESIGN)
}

func (c *GroupController) create(ctx *gin.Context, typeTag string) {
	userId, workspaceCode := getCurrentUser(ctx), getWorkspaceCode(ctx)
	reqBody := CreateAssetGroupReq{}
	if err := ctx.BindJSON(&reqBody); err != nil {
		c.AbortClientError(ctx, assetGroupHandlerMessage+err.Error())
		return
	}

	service.GroupServiceImpl.Create(ctx, reqBody.GroupName, workspaceCode, typeTag)

}

type UpdateAssetGroupReq struct {
	CreateAssetGroupReq
	DeleteAssetGroupReq
}

func (c *GroupController) projectRename(ctx *gin.Context) {

}

// rename
func (c *GroupController) rename(ctx *gin.Context) {
	userId, workspaceCode := getCurrentUser(ctx), getWorkspaceCode(ctx)
	reqBody := &UpdateAssetGroupReq{}

	if err := ctx.BindJSON(reqBody); err != nil {
		c.AbortClientError(ctx, assetGroupHandlerMessage+err.Error())
		return
	}

	if err := service.GroupServiceImpl.Rename(ctx, userId, workspaceCode, reqBody.GroupCode, reqBody.GroupName); err != nil {
		c.AbortClientError(ctx, assetGroupHandlerMessage+err.Error())
		return
	}

	c.ResponseJson(ctx)
}

type DeleteAssetGroupReq struct {
	GroupCode string `json:"groupCode" binding:"required"`
}

func (c *GroupController) delete(ctx *gin.Context) {

}

func (c *GroupController) move(ctx *gin.Context) {

}
