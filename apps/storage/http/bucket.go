package http

import (
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
		Code        string `json:"code" binding:"required"`
		Description string `json:"description"`
	}

	userId := ctx.GetUint64(controller.UserID)
	// TODO validate the user has create authority

	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.AbortClientError(ctx, "CreateBucket "+err.Error())
		return
	}

	bucketPath := ctx.GetString(consts.BucketPath)

	tx := db.NewTx(ctx).Begin()

	bucket, err := dao.CreateBucket(tx, req.Code, bucketPath)
	if err != nil {
		tx.Rollback()
		c.AbortServerError(ctx, "CreateBucket "+err.Error())
		return
	}

	if err := dao.AddUsersToBucket(tx, bucket.Id, []uint64{userId}); err != nil {
		tx.Rollback()
		c.AbortServerError(ctx, "CreateBucket "+err.Error())
		return
	}

	tx.Commit()

	c.ResponseJson(ctx, bucket)
}

func (c *BucketController) AddUsersToBucket(ctx *gin.Context) {
	var req struct {
		BucketID uint64   `json:"bucket_id" binding:"required"`
		Users    []uint64 `json:"users" binding:"required"`
	}

	if err := ctx.ShouldBind(&req); err != nil {
		c.AbortClientError(ctx, "AddUsersToBucket params not passed: "+err.Error())
		return
	}

	tx := db.NewTx(ctx).Begin()
	if err := dao.AddUsersToBucket(tx, req.BucketID, req.Users); err != nil {
		tx.Rollback()
		c.AbortServerError(ctx, "AddUsersToBucket internal error: "+err.Error())
		return
	}
	tx.Commit()

	c.ResponseJson(ctx, nil)
}

type BucketInfo struct {
	Id          uint64 `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
	UpdateTime  int64  `json:"update_time"`
	CreateTime  int64  `json:"create_time"`
}

func (c *BucketController) GetBucket(ctx *gin.Context) {
	var uri struct {
		Id uint64 `form:"bucket_id" binding:"required"`
	}
	tx := db.NewTx(ctx)
	if err := ctx.ShouldBind(&uri); err != nil {
		c.AbortClientError(ctx, "GetBucket payload error: "+err.Error())
		return
	}

	result, err := dao.GetBucket(tx, uri.Id)
	if err != nil {
		c.AbortServerError(ctx, "GetBucket internal error: "+err.Error())
		return
	}

	rootDir := model.Object{}

	if err := tx.Where("bucket_id = ? AND dir = ? AND parent_id = ?", result.Id, true, 0).Find(&rootDir).Error; err != nil { 
		c.AbortServerError(ctx, "GetBucket internal error: "+err.Error())
		return
	}

	data, err := dao.ListObjectsByDir(tx, result.Id, rootDir.Id)
	if err != nil {
		c.AbortServerError(ctx, "ListBucketDir query files error: "+err.Error())
		return
	}
	files, dirs := ClusterObjects(data)

	type InfoWithRootDir struct {
		BucketInfo
		DirContents DirResponse `json:"dir_contents"`
		RootPathId  uint64      `json:"root_path_id"`
	}

	res := InfoWithRootDir{}
	res.Id, res.Code, res.UpdateTime, res.CreateTime = result.Id, result.Code, result.UpdateTime.Unix(), result.CreateTime.Unix()
	res.Description = result.Description
	res.DirContents.File, res.DirContents.Directory, res.RootPathId = files, dirs, rootDir.Id

	c.ResponseJson(ctx, res)
}

func (c *BucketController) ListBuckets(ctx *gin.Context) {
	user := ctx.GetUint64(controller.UserID)

	pagination := shared.PaginationRequest{}
	if err := ctx.ShouldBindJSON(&pagination); err != nil {
		c.AbortClientError(ctx, "ListBuckets payload error: "+err.Error())
		return
	}

	var total int64
	tx := db.NewTx(ctx)
	// TODO many2many delete may cause problem
	if err := tx.Table("bucket_user").Where("user_id = ?", user).Count(&total).Error; err != nil {
		c.AbortServerError(ctx, "ListBuckets internal error: "+err.Error())
		return
	}

	if total < int64(pagination.Page*pagination.Size) {
		pagination.Page = 1
		pagination.Size = int(total)
	}

	paginationScope := db.Paginate(pagination.Page, pagination.Size)

	buckets := []*model.Bucket{}

	if err := tx.Table("buckets").
		Joins("INNER JOIN bucket_user ON bucket_user.bucket_id = buckets.id").
		Where("user_id = ? ", user).
		Scopes(paginationScope).Order("update_time DESC").Find(&buckets).Error; err != nil {
		c.AbortServerError(ctx, "ListBuckets internal error: "+err.Error())
		return
	}

	res := make([]*BucketInfo, len(buckets))

	for i, bucket := range buckets {
		res[i] = &BucketInfo{
			Id:         bucket.Id,
			Code:       bucket.Code,
			UpdateTime: bucket.UpdateTime.Unix(),
			CreateTime: bucket.CreateTime.Unix(),
		}
	}

	paginationResponse := shared.PaginationResponse{
		PageIndex:   pagination.Page,
		PageSize:    pagination.Size,
		RecordTotal: int(total),
	}

	var r struct {
		Data []*BucketInfo `json:"data"`
		*shared.PaginationResponse
	}
	r.PaginationResponse = &paginationResponse
	r.Data = res

	c.ResponseJson(ctx, r)

}

type FileInfo struct {
	FileId     string `json:"file_id"`
	Name       string `json:"name"`
	URI        string `json:"uri"`
	UpdateTime int64  `json:"update_time"`
	CreateTime int64  `json:"create_time"`
}

type DirInfo struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}
type DirResponse struct {
	File      []*FileInfo `json:"files"`
	Directory []*DirInfo  `json:"dirs"`
}

func ClusterObjects(data []*model.Object) ([]*FileInfo, []*DirInfo) {
	files := []*FileInfo{}
	dirs := []*DirInfo{}
	for _, d := range data {
		if d.Dir {
			dirs = append(dirs, &DirInfo{
				Id:   d.Id,
				Name: d.Name,
			})
		} else {

			files = append(files, &FileInfo{
				FileId:     d.FileId,
				Name:       d.Name,
				URI:        d.URI,
				UpdateTime: d.UpdateTime.Unix(),
				CreateTime: d.CreateTime.Unix(),
			})
		}
	}

	return files, dirs
}

func (c *BucketController) ListBucketDir(ctx *gin.Context) {

	var req struct {
		BucketID uint64 `form:"bucket_id" binding:"required"`
		PathId   uint64 `form:"path_id" binding:"required"`
	}

	if err := ctx.ShouldBind(&req); err != nil {
		c.AbortClientError(ctx, "ListBucketDir params not passed: "+err.Error())
		return
	}

	tx := db.NewTx(ctx)
	data, err := dao.ListObjectsByDir(tx, req.BucketID, req.PathId)
	if err != nil {
		c.AbortServerError(ctx, "ListBucketDir query files error: "+err.Error())
		return
	}
	files, dirs := ClusterObjects(data)
	type Response struct {
		File      []*FileInfo `json:"files"`
		Directory []*DirInfo  `json:"dirs"`
	}

	res := Response{
		File:      files,
		Directory: dirs,
	}

	c.ResponseJson(ctx, res)
}

func (c BucketController) CreateDir(ctx *gin.Context) {
	var req struct {
		BucketID uint64 `json:"bucket_id" binding:"required"`
		ParentID uint64 `json:"parent_id" binding:"required"`
		Name     string `json:"name" binding:"required"`
	}

	if err := ctx.ShouldBind(&req); err != nil {
		c.AbortClientError(ctx, "CreateDir params not passed: "+err.Error())
		return
	}

	tx := db.NewTx(ctx)

	// add deduplicate validate
	var existing model.Object
	if err := tx.Where("bucket_id = ? AND parent_id = ? AND name = ?", req.BucketID, req.ParentID, req.Name).First(&existing).Error; err == nil {
		c.AbortClientError(ctx, "CreateDir directory already exists")
		return
	}

	toCreate := model.Object{
		BucketID: req.BucketID,
		ParentId: req.ParentID,
		Name:     req.Name,
		Dir:      true,
	}

	if err := tx.Create(&toCreate).Error; err != nil {
		c.AbortServerError(ctx, "CreateDir internal error: "+err.Error())
		return
	}

	c.ResponseJson(ctx, nil)
}
