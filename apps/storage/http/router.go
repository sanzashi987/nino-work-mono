package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/pkg/auth"
)

func RegisterRoutes(r *gin.Engine) {
	bucketHandler := &BucketController{}
	
	storage := r.Group("/storage")
	storage.Use(auth.ValidateMiddleware("/login"))
	{
		// Bucket 管理
		storage.POST("/buckets", bucketHandler.CreateBucket)
		storage.GET("/buckets/:id", bucketHandler.GetBucket)
		
		// 文件上传（这部分应该通过RPC处理）
		storage.POST("/upload/:bucket", uploadHandler.Upload)
	}
} 