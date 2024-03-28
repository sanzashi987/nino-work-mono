package http

import (
	"github.com/cza14h/nino-work/pkg/auth"
	"github.com/cza14h/nino-work/pkg/controller"
	"github.com/gin-gonic/gin"
)

type AssetGroupController struct {
	controller.BaseController
}

func (c *AssetGroupController) list(ctx *gin.Context) {
}

/*CRUD*/
type CreateAssetGroupReq struct {
	GroupName string `json:"groupName" binding:"required"`
}

func (c *AssetGroupController) create(ctx *gin.Context) {

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
