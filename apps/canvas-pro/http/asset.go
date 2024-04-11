package http

import (
	"encoding/json"

	"github.com/cza14h/nino-work/apps/canvas-pro/http/request"
	"github.com/cza14h/nino-work/apps/canvas-pro/service"
	"github.com/cza14h/nino-work/pkg/controller"
	"github.com/gin-gonic/gin"
)

const asset_prefix = "assets"

type AssetController struct {
	controller.BaseController
}

var assetController = &AssetController{
	controller.BaseController{
		ErrorPrefix: "[http] canvas asset handler ",
	},
}

type ListAssetReq struct {
	GroupCode *string `json:"groupCode"`
	MimeType  *string `json:"filter"`
	FileName  *string `json:"fileName"`
	FileType  *string `json:"fileType"`
	Sort      *int    `json:"sort"`
	request.PaginationRequest
}

func (c *AssetController) list(ctx *gin.Context) {
}

type ReadQuery struct {
	FileId string `json:"fileId" binding:"required"`
}

/*CRUD*/
func (c *AssetController) read(ctx *gin.Context) {
	query := &ReadQuery{}
	if err := ctx.ShouldBindQuery(query); err != nil {
		c.AbortClientError(ctx, "[http] canvas asset read: fail to get required field in query, "+err.Error())
		return
	}
}

type UpdateAssetParam struct {
	FileId   string `json:"fileId"`
	FIleName string `json:"fileName"`
}

type UpdateAssetQuery struct {
	GroupCode string `json:"groupCode"`
	GroupName string `json:"groupName"`
}

func (c *AssetController) update(ctx *gin.Context) {

}
func (c *AssetController) delete(ctx *gin.Context) {

}

type UploadAssetReq struct {
	GroupCode string `json:"groupCode"`
	GroupName string `json:"groupName"`
	Type      string `json:"type"`
}

type UploadAssetRes struct {
	FileId   string `json:"fileId"`
	MimeType string `json:"mimeType"`
	Name     string `json:"name"`
	Size     int64  `json:"size"`
	Suffix   string `json:"suffix"`
}

func (c *AssetController) upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		c.AbortClientError(ctx, "upload: "+err.Error())
	}

}
func (c *AssetController) replace(ctx *gin.Context) {
}

func (c *AssetController) download(ctx *gin.Context) {
}
func (c *AssetController) _import(ctx *gin.Context) {
}

func (c *AssetController) moveGroup(ctx *gin.Context) {
	groupCode := ctx.Query("groupCode")
	assetCodesString := ctx.Query("fileIds")

	if groupCode == "" || assetCodesString == "" {
		c.AbortClientError(ctx, "move: "+"groupCode or fileIds is not provide")
		return
	}

	assetCodes := []string{}

	if err := json.Unmarshal([]byte(assetCodesString), &assetCodes); err != nil {
		c.AbortClientError(ctx, "move: "+err.Error())
		return
	}

	if err := service.AssetServiceImpl.BatchMoveGroup(ctx, assetCodes, groupCode, getWorkspaceCode(ctx)); err != nil {
		c.AbortClientError(ctx, "move: "+err.Error())
		return
	}

	c.SuccessVoid(ctx)

}
