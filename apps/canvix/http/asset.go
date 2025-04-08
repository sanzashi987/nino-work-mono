package http

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/apps/canvix/service"
	"github.com/sanzashi987/nino-work/pkg/shared"
)

type AssetController struct {
	CanvixController
}

func registerAssetRoutes(router *gin.RouterGroup, loggedMiddleware, workspaceMiddleware gin.HandlerFunc) {
	assetController := AssetController{}

	assetGroup := router.Group("assets")
	// assetGroup.Use(loggedMiddleware, workspaceMiddleware)
	{
		assetGroup.POST("selectMyAssets", assetController.list)
		assetGroup.POST("updateMyAssetsName", assetController.update)
		assetGroup.POST("updateAssetsGroup", assetController.moveGroup)
		assetGroup.DELETE("deleteAssets", assetController.delete)
		assetGroup.POST("upload", assetController.upload)
		assetGroup.POST("detail", assetController.read)
		assetGroup.POST("replace", assetController.replace)
		assetGroup.POST("loadAsset", assetController.download)
		assetGroup.POST("importAsset", assetController._import)
	}
}

type ListAssetResponse struct {
	Data []*service.ListAssetRes `json:"data"`
	shared.PaginationResponse
}

func (c *AssetController) list(ctx *gin.Context) {
	req := service.ListAssetReq{}
	workspaceId, err := c.BindRequestJson(ctx, &req, "list asset")
	if err != nil {
		return
	}

	recordTotal, res, err := service.ListAssetByType(ctx, workspaceId, consts.DATASOURCE, &req)
	if err != nil {
		c.AbortServerError(ctx, "list: "+err.Error())
		return
	}

	c.ResponseJson(ctx, ListAssetResponse{
		Data: res,
		PaginationResponse: shared.PaginationResponse{
			PageIndex:   req.Page,
			PageSize:    req.Size,
			RecordTotal: int(recordTotal),
		},
	})

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

	res, err := service.GetAssetDetail(ctx, uploadRpc, workspaceId, query.FileId)
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

		if err := service.BatchMoveGroup(ctx, workspaceId, req.FileIds, req.GroupName, req.GroupCode); err != nil {
			c.AbortServerError(ctx, "move: "+err.Error())
			return
		}

		c.SuccessVoid(ctx)
	}

}
