package http

import (
	"mime/multipart"

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
	FileId string `form:"fileId" binding:"required"`
}

/*CRUD*/
func (c *AssetController) read(ctx *gin.Context) {
	query := &ReadQuery{}
	if err := ctx.BindQuery(query); err != nil {
		c.AbortClientError(ctx, "read: "+err.Error())
		return
	}
}

type UpdateAssetQuery struct {
	GroupCode string `json:"groupCode"`
	GroupName string `json:"groupName"`
}

type UpdateAssetParam struct {
	FileId   string `json:"fileId"`
	FIleName string `json:"fileName"`
}

func (c *AssetController) update(ctx *gin.Context) {

	reqBody := UpdateAssetParam{}
	if err := ctx.BindJSON(&reqBody); err != nil {
		c.AbortClientError(ctx, "update: "+err.Error())
		return
	}

	_, workspaceId := getWorkspaceCode(ctx)

	if err := service.AssetServiceImpl.Update(ctx, workspaceId, reqBody.FIleName, reqBody.FileId); err != nil {
		c.AbortServerError(ctx, "update: "+err.Error())
		return
	}
	c.SuccessVoid(ctx)

}
func (c *AssetController) delete(ctx *gin.Context) {

}

type UploadAssetForm struct {
	GroupCode string                `form:"groupCode"`
	GroupName string                `form:"groupName"`
	Type      string                `form:"type"`
	File      *multipart.FileHeader `form:"file" binding:"required"`
}

func (c *AssetController) upload(ctx *gin.Context) {
	form := UploadAssetForm{}

	if err := ctx.Bind(&form); err != nil {
		c.AbortClientError(ctx, "upload: "+err.Error())
		return
	}

	uploadRpc := getUploadRpcService(ctx)

	_, workspaceId := getWorkspaceCode(ctx)
	res, err := service.AssetServiceImpl.UploadFile(ctx, uploadRpc, workspaceId, form.GroupCode, form.GroupName, form.File.Filename, form.Type, form.File)
	if err != nil {
		c.AbortServerError(ctx, "upload: "+err.Error())
		return
	}

	c.ResponseJson(ctx, res)

}
func (c *AssetController) replace(ctx *gin.Context) {
}

func (c *AssetController) download(ctx *gin.Context) {
}
func (c *AssetController) _import(ctx *gin.Context) {
}

type MoveGroupReq struct {
	FileIds []string `json:"fileIds" binding:"required"`
}

func (c *AssetController) moveGroup(ctx *gin.Context) {
	groupCode := ctx.Query("groupCode")
	groupName := ctx.Query("groupName")

	reqBody := MoveGroupReq{}
	if err := ctx.BindJSON(&reqBody); err != nil {
		c.AbortClientError(ctx, "move: "+err.Error())
		return
	}

	if groupCode == "" && groupName == "" {
		c.AbortClientError(ctx, "move: groupCode or groupName is required")
		return
	}

	_, workspaceId := getWorkspaceCode(ctx)

	if err := service.AssetServiceImpl.BatchMoveGroup(ctx, workspaceId, reqBody.FileIds, groupName, groupCode); err != nil {
		c.AbortServerError(ctx, "move: "+err.Error())
		return
	}

	c.SuccessVoid(ctx)

}
