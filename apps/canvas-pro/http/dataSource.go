package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/canvas-pro/http/request"
	"github.com/sanzashi987/nino-work/apps/canvas-pro/service"
)

const data_source_prefix = "jdbc-connect-template"

type DataSourceController struct {
	CanvasController
}

var dataSourceController = &DataSourceController{
	CanvasController: createCanvasController("[http] canvas data-source handler "),
}

type QueryDataSourceSearchRequest struct {
	SourceName string   `json:"sourceName"`
	SourceType []string `json:"sourceType"`
	Search     string   `json:"search"`
}

type QueryDataSourceRequest struct {
	request.PaginationRequest
	QueryDataSourceSearchRequest
}

func (c *DataSourceController) list(ctx *gin.Context) {

	reqBody := QueryDataSourceRequest{}
	workspaceId, err := c.BindRequestJson(ctx, &reqBody, "list")
	if err != nil {
		return
	}
	dataSourceList, err := service.DataSourceServiceImpl.ListDataSources(ctx, workspaceId, reqBody.Page, reqBody.Size, reqBody.SourceName, reqBody.SourceType)
	if err != nil {
		c.AbortServerError(ctx, "list "+err.Error())
		return
	}

	c.ResponseJson(ctx, dataSourceList)

}

/*CRUD*/
type CreateDataSourceRequest struct {
	SourceName string `json:"sourceName"`
	SourceType string `json:"sourceType"`
	SourceInfo string `json:"sourceInfo"`
}

func (c *DataSourceController) create(ctx *gin.Context) {

}

type ReadDataSourceQuery struct {
	SourceId string `uri:"sourceId" binding:"required"`
}

func (c *DataSourceController) read(ctx *gin.Context) {
	query := ReadDataSourceQuery{}
	if err := ctx.ShouldBindUri(&query); err != nil {
		c.AbortClientError(ctx, readPreix+err.Error())
		return
	}
	_, workspaceId := getWorkspaceCode(ctx)

	dataSource, err := service.DataSourceServiceImpl.GetDataSourceById(ctx, workspaceId, query.SourceId)
	if err != nil {
		c.AbortClientError(ctx, readPreix+err.Error())
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
	err = service.DataSourceServiceImpl.UpdateDataSourceById(ctx, workspaceId, &reqBody)
	if err != nil {
		c.AbortServerError(ctx, updatePrefix+err.Error())
		return
	}

	c.ResponseJson(ctx, nil)
}

type DeleteDataSourceRequest struct {
	SourceId []string `json:"sourceId"`
}

func (c *DataSourceController) delete(ctx *gin.Context) {

}

type ReplaceIpRequest struct {
	Search   string   `json:"search"`
	Target   string   `json:"target"`
	SourceId []string `json:"sourceId"`
}

func (c *DataSourceController) replaceIp(ctx *gin.Context) {
	reqBody := ReplaceIpRequest{}
	if workspaceId, err := c.BindRequestJson(ctx, &reqBody, "list"); err != nil {
		return
	}

}

type SearchByIpRequest struct {
	SourceTypes []string `json:"sourceType"`
	Search      string   `json:"search"`
}

func (c *DataSourceController) search(ctx *gin.Context) {
	reqBody := SearchByIpRequest{}
	if workspaceId, err := c.BindRequestJson(ctx, &reqBody, "list"); err != nil {
		return
	}

}
