package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/storage/db/dao"
)

type BucketHandler struct{}

func (h *BucketHandler) CreateBucket(c *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bucket, err := dao.NewBucketDao(c).CreateBucket(req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bucket)
}

func (h *BucketHandler) GetBucket(c *gin.Context) {
	id := c.Param("id")
	var bucket uint
	if _, err := fmt.Sscanf(id, "%d", &bucket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid bucket id"})
		return
	}

	result, err := dao.NewBucketDao(c).GetBucket(bucket)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "bucket not found"})
		return
	}

	c.JSON(http.StatusOK, result)
} 