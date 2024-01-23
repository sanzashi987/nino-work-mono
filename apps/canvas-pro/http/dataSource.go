package http

import (
	"github.com/cza14h/nino-work/pkg/controller"
	"github.com/gin-gonic/gin"
)

const data_source_group = "jdbc-connect-template"

type DataSourceController struct {
	controller.BaseController
}

type QueryDataSourceRequest struct {
	Page       int16    `json:"page"`
	Size       int16    `json:"size"`
	SourceName string   `json:"sourceName"`
	SourceType []string `json:"sourceType"`
	Search     string   `json:"search"`
}

func (c *DataSourceController) queryDataSourceList(ctx *gin.Context) {

}

/*CRUD*/
func (c *DataSourceController) createDataSource(ctx *gin.Context) {

}
func (c *DataSourceController) readDataSource(ctx *gin.Context) {

}
func (c *DataSourceController) updateDataSource(ctx *gin.Context) {

}
func (c *DataSourceController) deleteDataSource(ctx *gin.Context) {

}

func (c *DataSourceController) replaceIp(ctx *gin.Context) {

}
