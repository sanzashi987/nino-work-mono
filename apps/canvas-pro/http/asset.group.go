package http

import (
	"github.com/cza14h/nino-work/pkg/controller"
	"github.com/gin-gonic/gin"
)

type AssetGroupController struct {
	controller.BaseController
}

func (c *AssetGroupController) list(ctx *gin.Context) {
}

/*CRUD*/
func (c *AssetGroupController) create(ctx *gin.Context) {

}

type UpdateAssetGroup struct {
	GroupCode string `json:"groupCode"`
	GroupName string `json:"groupName"`
}

func (c *AssetGroupController) update(ctx *gin.Context) {

}
func (c *AssetGroupController) delete(ctx *gin.Context) {

}
