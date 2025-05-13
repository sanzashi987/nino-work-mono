package http

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/apps/canvix/service/asset"
)

type AssetController struct {
	CanvixController
}

func registerAssetRoutes(router *gin.RouterGroup, loggedMiddleware, workspaceMiddleware gin.HandlerFunc) {
	assetController := AssetController{}

	assetGroup := router.Group("assets")
	// assetGroup.Use(loggedMiddleware, workspaceMiddleware)

	assetGroup.POST("list", assetController.list)
	assetGroup.POST("list-all", assetController.list)
	assetGroup.POST("update", assetController.update)
	assetGroup.DELETE("delete", assetController.delete)
	assetGroup.POST("upload", assetController.upload)
	assetGroup.POST("detail", assetController.read)
	assetGroup.POST("replace", assetController.replace)
	assetGroup.POST("download", assetController.download)
	assetGroup.POST("import", assetController._import)
}

func (c *AssetController) list(ctx *gin.Context) {
	req := asset.ListAssetReq{}
	workspaceId, err := c.BindRequestJson(ctx, &req, "list asset")
	if err != nil {
		return
	}

	res, err := asset.ListAssetByType(ctx, workspaceId, consts.DATASOURCE, &req)
	if err != nil {
		c.AbortServerError(ctx, "list: "+err.Error())
		return
	}

	c.ResponseJson(ctx, res)
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
	workspaceId := c.MustGetWorkspaceId(ctx)

	uploadRpc := getUploadRpcService(ctx)

	res, err := asset.GetAssetDetail(ctx, uploadRpc, workspaceId, query.FileId)
	if err != nil {
		c.AbortServerError(ctx, "read: "+err.Error())
		return
	}

	c.ResponseJson(ctx, res)

}

func (c *AssetController) update(ctx *gin.Context) {
	req := asset.UpdateAssetReq{}
	workspaceId, err := c.BindRequestJson(ctx, &req, "asset update")
	if err != nil {
		return
	}

	if err := asset.UpdateName(ctx, workspaceId, &req); err != nil {
		c.AbortServerError(ctx, "update: "+err.Error())
		return
	}
	c.SuccessVoid(ctx)

}

func (c *AssetController) delete(ctx *gin.Context) {
	var req struct {
		Ids []string `json:"ids" binding:"required"`
	}
	workspaceId, err := c.BindRequestJson(ctx, &req, "asset delete")
	if err != nil {
		return
	}

	if err := asset.Delete(ctx, workspaceId, req.Ids); err != nil {
		c.AbortServerError(ctx, "delete: "+err.Error())
		return
	}
	c.SuccessVoid(ctx)

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
