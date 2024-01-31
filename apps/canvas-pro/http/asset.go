package http

import (
	"net/http"

	"github.com/cza14h/nino-work/pkg/controller"
	"github.com/gin-gonic/gin"
)

const asset_prefix = "assets"

type AssetController struct {
	controller.BaseController
}

func (c *AssetController) list(ctx *gin.Context) {
}

type ReadQuery struct {
	FileId string `json:"fileId"`
}

/*CRUD*/
func (c *AssetController) read(ctx *gin.Context) {
	query := &ReadQuery{}
	if err := ctx.BindQuery(query); err != nil {
		c.AbortJson(ctx, http.StatusBadRequest, "FileId should be provided in query", nil)
		return
	}
}

type UpdateAssetParam struct {
	FileId   string `json:"fileId"`
	FIleName string `json:"fileName"`
}

type UpdateAssetQuery struct {
	GroupCode string `json:"groupCode"`
	GroupName string `json:"groupName"`
}

func (c *AssetController) update(ctx *gin.Context) {

}
func (c *AssetController) delete(ctx *gin.Context) {

}

func (c *AssetController) upload(ctx *gin.Context) {
}
func (c *AssetController) replace(ctx *gin.Context) {
}

func (c *AssetController) download(ctx *gin.Context) {
}
func (c *AssetController) _import(ctx *gin.Context) {
}
