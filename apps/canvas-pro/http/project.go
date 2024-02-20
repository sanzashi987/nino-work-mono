package http

import (
	"net/http"

	"github.com/cza14h/nino-work/apps/canvas-pro/http/request"
	"github.com/cza14h/nino-work/apps/canvas-pro/service"
	"github.com/cza14h/nino-work/pkg/controller"
	"github.com/gin-gonic/gin"
)

const project_prefix = "screen-operation"

type ProjectController struct {
	controller.BaseController
}

type GetProjectListRequest struct {
	request.PaginationRequest
	Name  string
	Group string
}

func (c *ProjectController) list(ctx *gin.Context) {

}

/*CRUD*/
type CreateProjectRequest struct {
	Name string `binding:"required"`
	// Version     string
	GroupCode   string `json:"groupCode"`
	Config      string `json:"rootConfig" binding:"required"`
	UseTemplate string //template Id
}

func (c *ProjectController) create(ctx *gin.Context) {
	param := &CreateProjectRequest{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		c.AbortJson(ctx, http.StatusBadRequest, "Error in create projecet handler: "+err.Error())
		return
	}
	projectCode, err := service.GetProjectService().Create(ctx, param.Name, param.GroupCode, param.Config, param.UseTemplate)
	if err != nil {
		c.AbortJson(ctx, http.StatusInternalServerError, "Error in create projecet handler: "+err.Error())
	}
	c.ResponseJson(ctx, http.StatusOK, "Success", projectCode)
}

func (c *ProjectController) read(ctx *gin.Context) {
	id, err := c.MustGetParam(ctx, "id")
	if err != nil {
		return
	}

}
func (c *ProjectController) update(ctx *gin.Context) {

}
func (c *ProjectController) delete(ctx *gin.Context) {

}

// features
func (c *ProjectController) duplicate(ctx *gin.Context) {
	id, err := c.MustGetParam(ctx, "id")
	if err != nil {
		return
	}

	c.AbortJson(ctx, http.StatusBadRequest, "Id is required")
}

func (c *ProjectController) publish(ctx *gin.Context) {

}

func (c *ProjectController) export(ctx *gin.Context) {

}

func (c *ProjectController) _import(ctx *gin.Context) {

}

func (c *ProjectController) getInteraction(ctx *gin.Context) {

}
