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

const projectMessage = "Error in create project handler: "

func (c *ProjectController) create(ctx *gin.Context) {
	param := &CreateProjectRequest{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		c.AbortJson(ctx, http.StatusBadRequest, projectMessage+err.Error())
		return
	}
	projectCode, err := service.GetProjectService().Create(ctx, param.Name, param.GroupCode, param.Config, param.UseTemplate)
	if err != nil {
		c.AbortJson(ctx, http.StatusInternalServerError, projectMessage+err.Error())
		return
	}
	c.ResponseJson(ctx, http.StatusOK, "Success", projectCode)
}

func (c *ProjectController) read(ctx *gin.Context) {
	code, err := c.MustGetParam(ctx, "id")
	if err != nil {
		c.AbortJson(ctx, http.StatusBadRequest, projectMessage+err.Error())
		return
	}

	projectDetail, err := service.GetProjectService().GetInfoById(ctx, code)
	if err != nil {
		c.AbortJson(ctx, http.StatusInternalServerError, projectMessage+err.Error())
		return
	}

	c.ResponseJson(ctx, http.StatusOK, "", &projectDetail)

}

type ProjectUpdateRequest struct {
}

func (c *ProjectController) update(ctx *gin.Context) {
	request := ProjectUpdateRequest{}
	ctx.BindJSON(&request)

	service.GetProjectService().Update(ctx)
}
func (c *ProjectController) delete(ctx *gin.Context) {

}

// features
func (c *ProjectController) duplicate(ctx *gin.Context) {
	id, err := c.MustGetParam(ctx, "id")
	if err != nil {
		return
	}

}

func (c *ProjectController) publish(ctx *gin.Context) {

}

func (c *ProjectController) export(ctx *gin.Context) {

}

func (c *ProjectController) _import(ctx *gin.Context) {

}

func (c *ProjectController) getInteraction(ctx *gin.Context) {

}
