package http

import (
	"github.com/cza14h/nino-work/apps/canvas-pro/http/request"
	"github.com/cza14h/nino-work/apps/canvas-pro/service"
	"github.com/cza14h/nino-work/pkg/auth"
	"github.com/cza14h/nino-work/pkg/controller"
	"github.com/gin-gonic/gin"
)

const project_prefix = "screen-operation"

type ProjectController struct {
	controller.BaseController
}

type GetProjectListRequest struct {
	request.PaginationRequest
	Workspace string
	Name      *string
	Group     *string
}

const projectListMessage = "[http] canvas project list: "

func (c *ProjectController) list(ctx *gin.Context) {
	requestBody := &GetProjectListRequest{}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		c.AbortClientError(ctx, projectListMessage+err.Error())
		return
	}

	userId, exists := ctx.Get(auth.UserID)
	if !exists {
		c.AbortClientError(ctx, projectListMessage+" userId does not exist in http context")
		return
	}

	userIdTyped := userId.(uint64)
	infoList, err := service.GetProjectService().GetList(ctx, userIdTyped, requestBody.Page, requestBody.Size, requestBody.Workspace, requestBody.Name, requestBody.Group)
	if err != nil {
		c.AbortServerError(ctx, projectListMessage+err.Error())
		return
	}
	c.ResponseJson(ctx, infoList)

}

/*CRUD*/
type CreateProjectRequest struct {
	Name string `binding:"required"`
	// Version     string
	GroupCode   string `json:"groupCode"`
	Config      string `json:"rootConfig" binding:"required"`
	UseTemplate string //template Id
}

const projectCreateMessage = "[http] canvas project create: "

func (c *ProjectController) create(ctx *gin.Context) {
	param := &CreateProjectRequest{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		c.AbortClientError(ctx, projectCreateMessage+err.Error())
		return
	}
	projectCode, err := service.GetProjectService().Create(ctx, param.Name, param.GroupCode, param.Config, param.UseTemplate)
	if err != nil {
		c.AbortServerError(ctx, projectCreateMessage+err.Error())
		return
	}
	c.ResponseJson(ctx, projectCode)
}

func (c *ProjectController) read(ctx *gin.Context) {
	code, err := c.MustGetParam(ctx, "id")
	if err != nil {
		c.AbortClientError(ctx, projectCreateMessage+err.Error())
		return
	}

	projectDetail, err := service.GetProjectService().GetInfoById(ctx, code)
	if err != nil {
		c.AbortServerError(ctx, projectCreateMessage+err.Error())
		return
	}

	c.ResponseJson(ctx, &projectDetail)

}

type ProjectUpdateRequest struct {
	Code      string `json:"code" binding:"required"`
	Name      *string
	Config    *string `json:"rootConfig"`
	Thumbnail *string
	GroupCode *string `json:"groupCode"`
}

const projectUpdateMessage = "[http] canvas project update"

func (c *ProjectController) update(ctx *gin.Context) {
	param := ProjectUpdateRequest{}
	err := ctx.BindJSON(&param)
	if err != nil {
		c.AbortClientError(ctx, projectUpdateMessage+err.Error())
		return
	}

	if err := service.GetProjectService().Update(ctx, param.Code, param.Name, param.Config, param.Thumbnail, param.GroupCode); err != nil {
		c.AbortServerError(ctx, projectUpdateMessage+err.Error())
		return
	}

	c.ResponseJson(ctx, nil)

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
