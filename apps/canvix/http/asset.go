package http

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/apps/canvix/service"
	"github.com/sanzashi987/nino-work/pkg/shared"
)

const asset_prefix = "assets"

type AssetController struct {
	CanvixController
}

var assetController = &AssetController{
	CanvixController: createCanvixController("[http] canvas asset handler "),
}

type ListAssetResponse struct {
	Data []*service.ListAssetRes `json:"data"`
	shared.PaginationResponse
}

func (c *AssetController) list(ctx *gin.Context) {
	reqBody := service.ListAssetReq{}

	if workspaceId, err := c.BindRequestJson(ctx, &reqBody, "list"); err != nil {
		return
	} else {
		recordTotal, res, err := service.AssetServiceImpl.ListAssetByType(ctx, workspaceId, consts.DATASOURCE, &reqBody)
		if err != nil {
			c.AbortServerError(ctx, "list: "+err.Error())
			return
		}

		c.ResponseJson(ctx, ListAssetResponse{
			Data: res,
			PaginationResponse: shared.PaginationResponse{
				PageIndex:   reqBody.Page,
				PageSize:    reqBody.Size,
				RecordTotal: int(recordTotal),
			},
		})
	}

}

type ReadQuery struct {
	FileId string `form:"fileId" binding:"required"`
}

/*CRUD*/
func (c *AssetController) read(ctx *gin.Context) {
	var query struct {
		FileId string `form:"fileId" binding:"required"`
	}
	if err := ctx.ShouldBindQuery(&query); err != nil {
		c.AbortClientError(ctx, "read: "+err.Error())
		return
	}
	_, workspaceId := getWorkspaceCode(ctx)

	uploadRpc := getUploadRpcService(ctx)

	res, err := service.AssetServiceImpl.GetAssetDetail(ctx, uploadRpc, workspaceId, query.FileId)
	if err != nil {
		c.AbortServerError(ctx, "read: "+err.Error())
		return
	}

	c.ResponseJson(ctx, res)

}

func (c *AssetController) update(ctx *gin.Context) {
	reqBody := service.UpdateAssetReq{}
	if workspaceId, err := c.BindRequestJson(ctx, &reqBody, "update"); err != nil {
		return
	} else {
		if err := service.UpdateName(ctx, workspaceId, &reqBody); err != nil {
			c.AbortServerError(ctx, "update: "+err.Error())
			return
		}
		c.SuccessVoid(ctx)
	}

}
func (c *AssetController) delete(ctx *gin.Context) {
	var req struct {
		Data []string `json:"data" binding:"required"`
	}

	if workspaceId, err := c.BindRequestJson(ctx, &req, "delete"); err != nil {
		return
	} else {

		if err := service.DeleteAssets(ctx, workspaceId, req.Data); err != nil {
			c.AbortServerError(ctx, "delete: "+err.Error())
			return
		}
		c.SuccessVoid(ctx)

	}
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

	// uploadRpc := getUploadRpcService(ctx)
	panic("unimplemented")
	c.SuccessVoid(ctx)

}
func (c *AssetController) replace(ctx *gin.Context) {
}

func (c *AssetController) download(ctx *gin.Context) {
}
func (c *AssetController) _import(ctx *gin.Context) {
}

func (c *AssetController) moveGroup(ctx *gin.Context) {

	var req struct {
		FileIds   []string `json:"fileIds" binding:"required"`
		GroupCode string   `json:"groupCode"`
		GroupName string   `json:"groupName"`
	}
	if workspaceId, err := c.BindRequestJson(ctx, &req, "move"); err != nil {
		return
	} else {

		if req.GroupCode == "" && req.GroupName == "" {
			c.AbortClientError(ctx, "move: groupCode or groupName is required")
			return
		}

		if err := service.AssetServiceImpl.BatchMoveGroup(ctx, workspaceId, req.FileIds, req.GroupName, req.GroupCode); err != nil {
			c.AbortServerError(ctx, "move: "+err.Error())
			return
		}

		c.SuccessVoid(ctx)
	}

}
