package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/pkg/auth"
)

func RegisterRoutes(loginPageUrl string) {
	apiEngine := gin.Default()

	bucketHandler := &BucketController{}

	authMiddleware := auth.ValidateMiddleware(loginPageUrl)
	v1 := apiEngine.Group("/backend/v1")

	{
		authed := v1.Use(authMiddleware)
		// Bucket 管理
		authed.POST("/bucket/list", bucketHandler.ListBuckets)
		authed.GET("/bucket/:id", bucketHandler.GetBucket)

		// 文件上传（这部分应该通过RPC处理）
		authed.POST("/upload/:bucket", bucketHandler.Upload)
	}
}
