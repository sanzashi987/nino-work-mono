package http

import (
	"github.com/cza14h/nino-work/apps/canvas-pro/http/request"
	"github.com/cza14h/nino-work/apps/canvas-pro/service"
	"github.com/gin-gonic/gin"
)

const data_source_prefix = "jdbc-connect-template"

type DataSourceController struct {
	CanvasController
}

var dataSourceController = &DataSourceController{
	CanvasController: createCanvasController("[http] canvas data-source handler "),
}

type QueryDataSourceRequest struct {
	request.PaginationRequest
	SourceName string `json:"sourceName"`
	SourceType string `json:"sourceType"`
	Search     string `json:"search"`
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
func (c *DataSourceController) create(ctx *gin.Context) {

}
func (c *DataSourceController) read(ctx *gin.Context) {

}
func (c *DataSourceController) update(ctx *gin.Context) {

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
