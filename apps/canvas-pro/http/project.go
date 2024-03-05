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

const projectListMessage = "Error in listing project handler: "

func (c *ProjectController) list(ctx *gin.Context) {
	param := &GetProjectListRequest{}
	if err := ctx.ShouldBindJSON(&param); err != nil {
		c.AbortJson(ctx, http.StatusBadRequest, projectListMessage+err.Error())
	}

}

/*CRUD*/
type CreateProjectRequest struct {
	Name string `binding:"required"`
	// Version     string
	GroupCode   string `json:"groupCode"`
	Config      string `json:"rootConfig" binding:"required"`
	UseTemplate string //template Id
}

const projectCreateMessage = "Error in create project handler: "

func (c *ProjectController) create(ctx *gin.Context) {
	param := &CreateProjectRequest{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		c.AbortJson(ctx, http.StatusBadRequest, projectCreateMessage+err.Error())
		return
	}
	projectCode, err := service.GetProjectService().Create(ctx, param.Name, param.GroupCode, param.Config, param.UseTemplate)
	if err != nil {
		c.AbortJson(ctx, http.StatusInternalServerError, projectCreateMessage+err.Error())
		return
	}
	c.ResponseJson(ctx, http.StatusOK, "Success", projectCode)
}

func (c *ProjectController) read(ctx *gin.Context) {
	code, err := c.MustGetParam(ctx, "id")
	if err != nil {
		c.AbortJson(ctx, http.StatusBadRequest, projectCreateMessage+err.Error())
		return
	}

	projectDetail, err := service.GetProjectService().GetInfoById(ctx, code)
	if err != nil {
		c.AbortJson(ctx, http.StatusInternalServerError, projectCreateMessage+err.Error())
		return
	}

	c.ResponseJson(ctx, http.StatusOK, "", &projectDetail)

}

type ProjectUpdateRequest struct {
	Code      string `json:"code" binding:"required"`
	Name      *string
	Config    *string `json:"rootConfig"`
	Thumbnail *string
	GroupCode *string `json:"groupCode"`
}

const projectUpdateMessage = "Error in update project handler: "

func (c *ProjectController) update(ctx *gin.Context) {
	param := ProjectUpdateRequest{}
	err := ctx.BindJSON(&param)
	if err != nil {
		c.AbortJson(ctx, http.StatusInternalServerError, projectUpdateMessage+err.Error())
		return
	}

	if err := service.GetProjectService().Update(ctx, param.Code, param.Name, param.Config, param.Thumbnail, param.GroupCode); err != nil {
		c.AbortJson(ctx, http.StatusInternalServerError, projectUpdateMessage+err.Error())
		return
	}

	c.ResponseJson(ctx, http.StatusOK, "Success", nil)

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
