package http

import (
	"math"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/storage/consts"
	"github.com/sanzashi987/nino-work/apps/storage/db/dao"
	"github.com/sanzashi987/nino-work/apps/storage/db/model"
	"github.com/sanzashi987/nino-work/pkg/controller"
	"github.com/sanzashi987/nino-work/pkg/db"
	"github.com/sanzashi987/nino-work/pkg/shared"
)

type BucketController struct {
	controller.BaseController
}

var bucketController BucketController = BucketController{
	controller.BaseController{
		ErrorPrefix: "[http]: bucket controller ",
	},
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

	bucketPath := ctx.GetString(consts.BucketPath)

	bucket, err := dao.NewBucketDao(ctx).CreateBucket(req.Name,bucketPath)
	if err != nil {
		c.AbortServerError(ctx, ""+err.Error())
		return
	}

	c.ResponseJson(ctx, bucket)
}

func (c *BucketController) GetBucket(ctx *gin.Context) {
	var uri struct {
		Id uint64 `uri:"id" binding:"required"`
	}

	if err := ctx.ShouldBindUri(&uri); err != nil {
		c.AbortClientError(ctx, "GetBucket payload error: "+err.Error())
		return
	}

	result, err := dao.NewBucketDao(ctx).GetBucket(uri.Id)
	if err != nil {
		c.AbortServerError(ctx, "GetBucket internal error: "+err.Error())
		return
	}

	c.ResponseJson(ctx, result)
}

func (c *BucketController) ListBuckets(ctx *gin.Context) {
	user := ctx.GetUint64(controller.UserID)

	pagination := shared.PaginationRequest{}
	if err := ctx.ShouldBindJSON(&pagination); err != nil {
		c.AbortClientError(ctx, "ListBuckets payload error: "+err.Error())
		return
	}

	offset := (pagination.Page - 1) * pagination.Size
	offset = int(math.Max(0, float64(offset)))

	paginationScope := db.Paginate(pagination.Page, pagination.Size)

	bucketDao := dao.NewBucketDao(ctx)
	u := model.User{
		UserId: user,
		Type:   model.USER,
	}
	err := bucketDao.GetOrm().Preload("Buckets").Scopes(paginationScope).Find(&u).Error
	if err != nil {
		c.AbortServerError(ctx, "ListBuckets internal error: "+err.Error())
		return
	}

	type Res struct {
		Code       string `json:"code"`
		UpdateTime int64  `json:"update_time"`
		CreateTime int64  `json:"create_time"`
	}

	res := make([]Res, len(u.Buckets))

	for i, bucket := range u.Buckets {
		res[i] = Res{
			Code:       bucket.Code,
			UpdateTime: bucket.UpdateTime.Unix(),
			CreateTime: bucket.CreateTime.Unix(),
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].UpdateTime > res[j].UpdateTime
	})
	c.ResponseJson(ctx, res)
}

func (c *BucketController) ListBucketDir(ctx *gin.Context) {

	var req struct {
		BucketID uint64 `uri:"id" binding:"required"`
	}

	if err:= ctx.ShouldBindUri(&req); err !=nil {
		
	}

}
