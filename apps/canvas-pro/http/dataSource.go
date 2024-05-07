package http

import (
	"github.com/cza14h/nino-work/apps/canvas-pro/http/request"
	"github.com/cza14h/nino-work/pkg/controller"
	"github.com/gin-gonic/gin"
)

const data_source_prefix = "jdbc-connect-template"

type DataSourceController struct {
	controller.BaseController
}

var dataSourceController = &DataSourceController{
	controller.BaseController{
		ErrorPrefix: "[http] canvas data-source handler ",
	},
}

type QueryDataSourceRequest struct {
	request.PaginationRequest
	SourceName string   `json:"sourceName"`
	SourceType []string `json:"sourceType"`
	Search     string   `json:"search"`
}

func (c *DataSourceController) list(ctx *gin.Context) {
	reqBody := QueryDataSourceRequest{}

	if err := ctx.BindJSON(&reqBody); err != nil {
		c.AbortClientError(ctx, "list "+err.Error())
		return
	}

	// service.AssetServiceImpl.ListDataSources()

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

func (c *DataSourceController) replaceIp(ctx *gin.Context) {

}
