package http

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/storage/db/model"
	"github.com/sanzashi987/nino-work/apps/storage/service"
	"github.com/sanzashi987/nino-work/pkg/controller"
	"github.com/sanzashi987/nino-work/pkg/db"
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

	form, err := ctx.MultipartForm()

	if err != nil {
		c.AbortClientError(ctx, "upload payload error: "+err.Error())
		return
	}

	bucket_id, path_id := form.Value["bucket_id"], form.Value["path_id"]
	files := form.File["file[]"]

	if len(bucket_id) == 0 || len(path_id) == 0 || len(files) == 0 {
		c.AbortClientError(ctx, "upload payload error: bucket_id or path_id is not provided")
		return
	}

	bucketID, err := strconv.ParseUint(bucket_id[0], 10, 64)
	pathId, err2 := strconv.ParseUint(path_id[0], 10, 64)

	if err != nil || err2 != nil {
		c.AbortClientError(ctx, "upload payload error: bucket_id or path_id is not a number")
		return
	}

	payload := service.UploadFilePayload{
		BucketID: bucketID,
		PathId:   pathId,
		Files:    files,
	}

	uid, err := service.UploadServiceWebImpl.UploadFile(ctx, user, &payload)
	if err != nil {
		c.AbortServerError(ctx, "upload file error: "+err.Error())
		return
	}

	c.ResponseJson(ctx, gin.H{"file_ids": uid})

}

func (c *ObjectController) GetAsset(ctx *gin.Context) {
	// user := ctx.GetUint64(controller.UserID)
	// // if the api is called not authed
	// authed := user == 0

	var req struct {
		BucketID uint64 `uri:"bucket" binding:"required"`
		FileId   string `uri:"file_id" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&req); err != nil {
		c.AbortClientError(ctx, "GetAsset payload error: "+err.Error())
		return
	}

	tx := db.NewTx(ctx)
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
