package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/canvix/service/dataSource"
)

type DataSourceController struct {
	CanvixController
}

func registerDataSourceRoutes(router *gin.RouterGroup, loggedMiddleware, workspaceMiddleware gin.HandlerFunc) {
	dataSourceController := DataSourceController{}

	nonAuthed := router.Group("data-source")

	authed := nonAuthed.Use(loggedMiddleware, workspaceMiddleware)

	authed.POST("list", dataSourceController.list)
	authed.POST("create", dataSourceController.create)
	authed.POST("update", dataSourceController.update)
	authed.DELETE("delete", dataSourceController.delete)
	// authed.POST("replace-ip", dataSourceController.replaceIp)
	authed.GET("info/:sourceId", dataSourceController.read)
	authed.POST("search", dataSourceController.search)
	authed.POST("preview-file", dataSourceController.previewFile)

	nonAuthed.POST("file", dataSourceController.getFileFromPublic)
	nonAuthed.POST("api", dataSourceController.getApiFromPublic)
	nonAuthed.GET("static", dataSourceController.getStaticFromPublic)

}

func (c *DataSourceController) list(ctx *gin.Context) {

	req := dataSource.ListReq{}
	workspaceId, err := c.BindRequestJson(ctx, &req, "list")
	if err != nil {
		return
	}
	dataSourceList, err := dataSource.List(ctx, workspaceId, &req)
	if err != nil {
		c.AbortServerError(ctx, "list "+err.Error())
		return
	}

	c.ResponseJson(ctx, dataSourceList)

}

/*CRUD*/
func (c *DataSourceController) create(ctx *gin.Context) {
	req := dataSource.CreateReq{}
	workspaceId, err := c.BindRequestJson(ctx, &req, "create")
	if err != nil {
		return
	}

	dataSource, err := dataSource.Create(ctx, workspaceId, &req)
	if err != nil {
		c.AbortServerError(ctx, "data create error: "+err.Error())
		return
	}

	c.ResponseJson(ctx, dataSource)
}

func (c *DataSourceController) read(ctx *gin.Context) {
	var query struct {
		SourceId string `uri:"sourceId" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&query); err != nil {
		c.AbortClientError(ctx, "data read error: "+err.Error())
		return
	}

	workspaceId := c.MustGetWorkspaceId(ctx)

	dataSource, err := dataSource.GetDataSourceById(ctx, workspaceId, query.SourceId)
	if err != nil {
		c.AbortServerError(ctx, "data read error: "+err.Error())
		return
	}

	c.ResponseJson(ctx, dataSource)
}

func (c *DataSourceController) update(ctx *gin.Context) {
	req := dataSource.UpdateReq{}
	workspaceId, err := c.BindRequestJson(ctx, &req, "udpate")
	if err != nil {
		return
	}
	dataSource, err := dataSource.Update(ctx, workspaceId, &req)
	if err != nil {
		c.AbortServerError(ctx, "data update error: "+err.Error())
		return
	}

	c.ResponseJson(ctx, dataSource)
}

func (c *DataSourceController) delete(ctx *gin.Context) {
	var req struct {
		SourceId []string `json:"sourceId" binding:"required"`
	}
	workspaceId, err := c.BindRequestJson(ctx, &req, "data delete")
	if err != nil {
		return
	}

	err = dataSource.Delete(ctx, workspaceId, req.SourceId)
	if err != nil {
		c.AbortServerError(ctx, "data delete error: "+err.Error())
		return
	}

	c.SuccessVoid(ctx)
}

// func (c *DataSourceController) replaceIp(ctx *gin.Context) {
// 	var reqBody struct {
// 		Search   string   `json:"search"`
// 		Target   string   `json:"target"`
// 		SourceId []string `json:"sourceId"`
// 	}
// 	if workspaceId, err := c.BindRequestJson(ctx, &reqBody, "replaceIp"); err != nil {
// 		return
// 	}

// }

func (c *DataSourceController) search(ctx *gin.Context) {
	// var req struct {
	// 	SourceTypes []string `json:"sourceType"`
	// 	Search      string   `json:"search"`
	// }
	// if workspaceId, err := c.BindRequestJson(ctx, &req, "search"); err != nil {
	// return
	// }

}

func (c *DataSourceController) previewFile(ctx *gin.Context) {
	// var req struct {
	// 	SourceId string `json:"sourceId" binding:"required"`
	// }
	// if workspaceId, err := c.BindRequestJson(ctx, &req, "previewFile"); err != nil {
	// return
	// }

}

func (c *DataSourceController) getFileFromPublic(ctx *gin.Context) {
}
func (c *DataSourceController) getApiFromPublic(ctx *gin.Context) {
}
func (c *DataSourceController) getStaticFromPublic(ctx *gin.Context) {
}
