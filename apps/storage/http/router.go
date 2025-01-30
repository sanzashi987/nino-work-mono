package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/storage/consts"
	"github.com/sanzashi987/nino-work/pkg/controller"
)

func NewRouter(loginPageUrl, bucketPath, tmpPath string) *gin.Engine {
	apiEngine := gin.Default()

	applyPathInContext := func(ctx *gin.Context) {
		ctx.Set(consts.BucketPath, bucketPath)
		ctx.Set(consts.TmpPath, tmpPath)
	}

	authMiddleware := controller.ValidateMiddleware(loginPageUrl)
	v1 := apiEngine.Group("/backend/v1").Use(applyPathInContext)

	{
		authed := v1.Use(authMiddleware)
		// Bucket 管理
		authed.POST("/bucket/list", bucketController.ListBuckets)
		authed.POST("/bucket/create", bucketController.CreateBucket)
		authed.POST("/bucket/create-dir", bucketController.CreateDir)
		authed.GET("/bucket/info", bucketController.GetBucket)
		authed.GET("/bucket/list-dir", bucketController.ListBucketDir)

	}
	v1.POST("/asset/upload/:bucket_id", objectController.UploadFile)
	v1.GET("/asset/:bucket/:file_id", objectController.GetAsset)

	return apiEngine
}
