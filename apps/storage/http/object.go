package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/storage/service"
	"github.com/sanzashi987/nino-work/pkg/controller"
)

type ObjectController struct {
	controller.BaseController
}

var objectController ObjectController = ObjectController{
	controller.BaseController{
		ErrorPrefix: "[http] object controller ",
	},
}

func (c *ObjectController) UploadFile(ctx *gin.Context) {
	user := ctx.GetUint64(controller.UserID)

	req := service.UploadFilePayload{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.AbortClientError(ctx, "upload payload error: "+err.Error())
		return
	}

	service.UploadServiceWebImpl.UploadFile(ctx, user, &req)

}

func (c *ObjectController) UploadLargeFile(ctx *gin.Context) {

}

func (c *ObjectController) CheckFileExists(ctx *gin.Context) {

	hash := ctx.Param("hash")
	if hash == "" {
		c.AbortClientError(ctx, "[http]: check file exist client error, hash is not provided ")
	}

}
