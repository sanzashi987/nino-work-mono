package http

import (
	"github.com/cza14h/nino-work/pkg/controller"
	"github.com/gin-gonic/gin"
)

const asset_prefix = "assets"

type AssetController struct {
	controller.BaseController
}

func (c *AssetController) list(ctx *gin.Context) {
}

/*CRUD*/
func (c *AssetController) create(ctx *gin.Context) {

}
func (c *AssetController) read(ctx *gin.Context) {

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
