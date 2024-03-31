package http

import (
	"github.com/cza14h/nino-work/apps/canvas-pro/consts"
	"github.com/cza14h/nino-work/apps/canvas-pro/service"
	"github.com/cza14h/nino-work/pkg/auth"
	"github.com/cza14h/nino-work/pkg/controller"
	"github.com/gin-gonic/gin"
)

type AssetGroupController struct {
	controller.BaseController
}

const assetGroupHandlerMessage = "[http] canvas asset group handler "

func (c *AssetGroupController) list(ctx *gin.Context) {
}

/*CRUD*/
type CreateAssetGroupReq struct {
	GroupName string `json:"groupName" binding:"required"`
	Workspace string `json:"workspace" binding:"required"`
}

func (c *AssetGroupController) create(ctx *gin.Context) {
	userId, _ := ctx.Get(auth.UserID)

	reqBody := CreateAssetGroupReq{}
	if err := ctx.BindJSON(&reqBody); err != nil {
		c.AbortClientError(ctx, assetGroupHandlerMessage+err.Error())
		return
	}

	service.GroupServiceImpl.Create(ctx, reqBody.GroupName, reqBody.Workspace, consts.GROUP)

}

type UpdateAssetGroupReq struct {
	CreateAssetGroupReq
	DeleteAssetGroupReq
}

// rename
func (c *AssetGroupController) update(ctx *gin.Context) {
	userId, _ := ctx.Get(auth.UserID)
}

type DeleteAssetGroupReq struct {
	GroupCode string `json:"groupCode" binding:"required"`
}

func (c *AssetGroupController) delete(ctx *gin.Context) {

}
