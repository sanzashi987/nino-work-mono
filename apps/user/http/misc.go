package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/pkg/controller"
)

type MiscController struct {
	controller.BaseController
}

var miscController = MiscController{}

func (c *MiscController) GetImportMap(ctx *gin.Context) {
	data := gin.H{
		"message": "test token",
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		c.AbortServerError(ctx, "Failed to create JSON")
		return
	}

	ctx.Writer.WriteHeader(http.StatusOK)
	// Set the appropriate headers for the file response
	ctx.Header("Content-Disposition", "attachment; filename=importmap.json")
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Content-Length", fmt.Sprintf("%d", len(jsonData)))

	// Write the JSON content to the response body
	ctx.Writer.Write(jsonData)
}
