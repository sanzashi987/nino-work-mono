package http

import (
	"github.com/cza14h/nino-work/pkg/controller"
	"github.com/gin-gonic/gin"
)

const data_source_prefix = "jdbc-connect-template"

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

func (c *DataSourceController) list(ctx *gin.Context) {

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
