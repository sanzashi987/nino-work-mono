package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/canvix/service"
	"github.com/sanzashi987/nino-work/pkg/controller"
	"github.com/sanzashi987/nino-work/pkg/shared"
)

const project_prefix = "screen-operation"

type ProjectController struct {
	controller.BaseController
}

var projectController = &ProjectController{
	controller.BaseController{
		ErrorPrefix: "[http] canvas project handler ",
	},
}

type GetProjectListRequest struct {
	shared.PaginationRequest
	// Workspace string
	Name  *string
	Group *string
}

func (c *ProjectController) list(ctx *gin.Context) {
	requestBody := &GetProjectListRequest{}
	if err := ctx.BindJSON(&requestBody); err != nil {
		c.AbortClientError(ctx, listPrefix+err.Error())
		return
	}

	// userId, exists := ctx.Get(auth.UserID)
	// if !exists {
	// 	c.AbortClientError(ctx, listPrefix+" userId does not exist in http context")
	// 	return
	// }
	_, workspaceId := getWorkspaceCode(ctx)
	infoList, err := service.ProjectServiceImpl.GetList(ctx, workspaceId, requestBody.Page, requestBody.Size, requestBody.Name, requestBody.Group)
	if err != nil {
		c.AbortServerError(ctx, listPrefix+err.Error())
		return
	}
	c.ResponseJson(ctx, infoList)
}

/*CRUD*/
type CreateProjectRequest struct {
	Name string `json:"name" binding:"required"`
	// Version     string
	Config      string  `json:"rootConfig" binding:"required"`
	GroupCode   *string `json:"groupCode"`
	UseTemplate *string `json:"useTemplate"` //template Id
}

func (c *ProjectController) create(ctx *gin.Context) {
	param := &CreateProjectRequest{}
	if err := ctx.BindJSON(param); err != nil {
		c.AbortClientError(ctx, createPrefix+err.Error())
		return
	}
	projectCode, err := service.ProjectServiceImpl.Create(ctx, param.Name, param.Config, param.GroupCode, param.UseTemplate)
	if err != nil {
		c.AbortServerError(ctx, createPrefix+err.Error())
		return
	}
	c.ResponseJson(ctx, projectCode)
}

type ReadProjectQuery struct {
	Id string `uri:"id" binding:"required"`
}

func (c *ProjectController) read(ctx *gin.Context) {
	query := ReadProjectQuery{}

	if err := ctx.ShouldBindUri(&query); err != nil {
		c.AbortClientError(ctx, createPrefix+err.Error())
		return
	}

	projectDetail, err := service.ProjectServiceImpl.GetInfoById(ctx, query.Id)
	if err != nil {
		c.AbortServerError(ctx, createPrefix+err.Error())
		return
	}

	c.ResponseJson(ctx, &projectDetail)
}

type ProjectUpdateRequest struct {
	Code      string  `json:"code"`
	Name      *string `json:"name"`
	Config    *string `json:"rootConfig"`
	Thumbnail *string `json:"thumbnail"`
	GroupCode *string `json:"groupCode"`
}

func (c *ProjectController) update(ctx *gin.Context) {
	param := ProjectUpdateRequest{}
	err := ctx.BindJSON(&param)
	if err != nil {
		c.AbortClientError(ctx, updatePrefix+err.Error())
		return
	}

	if err := service.ProjectServiceImpl.Update(ctx, param.Code, param.Name, param.Config, param.Thumbnail, param.GroupCode); err != nil {
		c.AbortServerError(ctx, updatePrefix+err.Error())
		return
	}

	c.SuccessVoid(ctx)
}

type BatchMoveProjectGroupRequest struct {
	GroupName string   `json:"groupName"`
	GroupCode string   `json:"groupCode"`
	Ids       []string `json:"codes" binding:"required"`
}

func (c *ProjectController) moveGroup(ctx *gin.Context) {
	reqBody := BatchMoveProjectGroupRequest{}

	if err := ctx.BindJSON(&reqBody); err != nil {
		c.AbortClientError(ctx, "move: "+err.Error())
		return
	}
	_, workspaceId := getWorkspaceCode(ctx)

	if err := service.ProjectServiceImpl.BatchMoveGroup(ctx, workspaceId, reqBody.Ids, reqBody.GroupName, reqBody.GroupCode); err != nil {
		c.AbortClientError(ctx, "move: "+err.Error())
		return
	}

	c.SuccessVoid(ctx)
}

type ProjectDeleteRequest struct {
	Ids []string `json:"ids" binding:"required"`
}

func (c *ProjectController) delete(ctx *gin.Context) {
	param := ProjectDeleteRequest{}
	if err := ctx.BindJSON(&param); err != nil {
		c.AbortClientError(ctx, deletePrefix+err.Error())
		return
	}

	if err := service.ProjectServiceImpl.LogicalDeletion(ctx, param.Ids); err != nil {
		c.AbortServerError(ctx, deletePrefix+err.Error())
		return
	}
	c.ResponseJson(ctx, nil)
}

// features
func (c *ProjectController) duplicate(ctx *gin.Context) {
	query := ReadProjectQuery{}

	if err := ctx.ShouldBindUri(&query); err != nil {
		return
	}
	projectCode, err := service.ProjectServiceImpl.Duplicate(ctx, query.Id)
	if err != nil {
		c.AbortServerError(ctx, createPrefix+err.Error())
		return
	}
	c.ResponseJson(ctx, projectCode)
}

type ProjectPublishRequest struct {
	Code             string  `json:"code"`
	PublishFlag      int     `json:"publishFlag"`
	PublishSecretKey *string `json:"publishSecretKey"`
}

const publishPrefix = "publish: "

func (c *ProjectController) publish(ctx *gin.Context) {
	req := &ProjectPublishRequest{}
	if err := ctx.BindJSON(req); err != nil {
		c.AbortClientError(ctx, publishPrefix+err.Error())
		return
	}

	if err := service.ProjectServiceImpl.PublishProject(ctx, req.Code, req.PublishFlag, req.PublishSecretKey); err != nil {
		c.AbortServerError(ctx, publishPrefix+err.Error())
		return
	}
	c.ResponseJson(ctx, nil)
}

func (c *ProjectController) export(ctx *gin.Context) {

}

func (c *ProjectController) _import(ctx *gin.Context) {

}

// func (c *ProjectController) getInteraction(ctx *gin.Context) {

// }
