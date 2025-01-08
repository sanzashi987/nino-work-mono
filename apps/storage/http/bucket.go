package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/storage/db/dao"
	"github.com/sanzashi987/nino-work/pkg/controller"
)

type BucketController struct {
	controller.BaseController
}

func (c *BucketController) CreateBucket(ctx *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
		Code string `json:"code" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.AbortClientError(ctx, ""+err.Error())
		return
	}

	bucket, err := dao.NewBucketDao(ctx).CreateBucket(req.Name)
	if err != nil {
		c.AbortServerError(ctx, ""+err.Error())
		return
	}

	c.ResponseJson(ctx, bucket)
}

func (c *BucketController) GetBucket(ctx *gin.Context) {
	id := ctx.Param("id")
	bucketid, err := strconv.ParseUint(id, 10, 64)
	if err != nil || bucketid == 0 {
		c.AbortClientError(ctx, "[http]: get bucket params error id is not allowed")
		return
	}

	result, err := dao.NewBucketDao(ctx).GetBucket(uint(bucketid))
	if err != nil {
		c.AbortServerError(ctx, "[http]: get bucket service error: "+ err.Error())
		return
	}

	c.ResponseJson(ctx, result)
}
