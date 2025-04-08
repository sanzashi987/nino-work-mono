package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/canvix/service"
)

const data_source_prefix = "data-source"

type DataSourceController struct {
	CanvixController
}


func (c *DataSourceController) list(ctx *gin.Context) {

	reqBody := service.QueryDataSourceRequest{}
	workspaceId, err := c.BindRequestJson(ctx, &reqBody, "list")
	if err != nil {
		return
	}
	dataSourceList, err := service.DataSourceServiceImpl.ListDataSources(ctx, workspaceId, &reqBody)
	if err != nil {
		c.AbortServerError(ctx, "list "+err.Error())
		return
	}

	c.ResponseJson(ctx, dataSourceList)

}

/*CRUD*/

func (c *DataSourceController) create(ctx *gin.Context) {
	reqBody := service.CreateDataSourceRequest{}
	workspaceId, err := c.BindRequestJson(ctx, &reqBody, "create")
	if err != nil {
		return
	}

	dataSource, err := service.DataSourceServiceImpl.Create(ctx, workspaceId, &reqBody)
	if err != nil {
		c.AbortServerError(ctx, createPrefix+err.Error())
		return
	}

	c.ResponseJson(ctx, dataSource)
}

func (c *DataSourceController) read(ctx *gin.Context) {
	var query struct {
		SourceId string `uri:"sourceId" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&query); err != nil {
		c.AbortClientError(ctx, readPreix+err.Error())
		return
	}
	_, workspaceId := getWorkspaceCode(ctx)

	dataSource, err := service.DataSourceServiceImpl.GetDataSourceById(ctx, workspaceId, query.SourceId)
	if err != nil {
		c.AbortServerError(ctx, readPreix+err.Error())
		return
	}

	c.ResponseJson(ctx, dataSource)
}

func (c *DataSourceController) update(ctx *gin.Context) {
	reqBody := service.UpdateDataSourceRequest{}
	workspaceId, err := c.BindRequestJson(ctx, &reqBody, "udpate")
	if err != nil {
		return
	}
	dataSource, err := service.DataSourceServiceImpl.Update(ctx, workspaceId, &reqBody)
	if err != nil {
		c.AbortServerError(ctx, updatePrefix+err.Error())
		return
	}

	c.ResponseJson(ctx, dataSource)
}

func (c *DataSourceController) delete(ctx *gin.Context) {
	var reqBody struct {
		SourceId []string `json:"sourceId" binding:"required"`
	}
	workspaceId, err := c.BindRequestJson(ctx, &reqBody, "delete")
	if err != nil {
		return
	}

	err = service.DataSourceServiceImpl.Delete(ctx, workspaceId, reqBody.SourceId)
	if err != nil {
		c.AbortServerError(ctx, deletePrefix+err.Error())
		return
	}

	c.SuccessVoid(ctx)
}

func (c *DataSourceController) replaceIp(ctx *gin.Context) {
	var reqBody struct {
		Search   string   `json:"search"`
		Target   string   `json:"target"`
		SourceId []string `json:"sourceId"`
	}
	if workspaceId, err := c.BindRequestJson(ctx, &reqBody, "replaceIp"); err != nil {
		return
	}

}

func (c *DataSourceController) search(ctx *gin.Context) {
	var reqBody struct {
		SourceTypes []string `json:"sourceType"`
		Search      string   `json:"search"`
	}
	if workspaceId, err := c.BindRequestJson(ctx, &reqBody, "search"); err != nil {
		return
	}

}
