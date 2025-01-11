package http

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/storage/db/dao"
	"github.com/sanzashi987/nino-work/apps/storage/db/model"
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

func (c *ObjectController) GetAsset(ctx *gin.Context) {
	user := ctx.GetUint64(controller.UserID)
	authed := user == 0

	var req struct {
		BucketID uint64 `uri:"bucket" binding:"required"`
		FileId   string `uri:"file_id" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&req); err != nil {
		c.AbortClientError(ctx, "GetAsset payload error: "+err.Error())
		return
	}

	tx := dao.NewObjectDao(ctx).GetOrm()
	object := model.Object{}
	if err := tx.Where("bucket_id = ? AND file_id = ?", req.BucketID, req.FileId).Find(&object).Error; err != nil {
		c.AbortServerError(ctx, "GetAsset internal error "+err.Error())
		return
	}
	reader, err := os.Open(object.URI)
	if err != nil {
		c.AbortServerError(ctx, "GetAsset read file error "+err.Error())
		return
	}
	ctx.DataFromReader(http.StatusOK, object.Size, object.MimeType, reader, map[string]string{})
}

func (c *ObjectController) UploadLargeFile(ctx *gin.Context) {

}

func (c *ObjectController) CheckFileExists(ctx *gin.Context) {

	hash := ctx.Param("hash")
	if hash == "" {
		c.AbortClientError(ctx, "[http]: check file exist client error, hash is not provided ")
	}

}
