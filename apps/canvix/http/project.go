package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/canvix/service"
	"github.com/sanzashi987/nino-work/pkg/shared"
)

type ProjectController struct {
	CanvixController
}

func registerProjectRoutes(router *gin.RouterGroup, loggedMiddleware, workspaceMiddleware gin.HandlerFunc) {
	projectController := ProjectController{}

	projectRoutes := router.Group("screen-operation").Use(loggedMiddleware, workspaceMiddleware)

	projectRoutes.POST("create", projectController.create)
	projectRoutes.POST("createByTemplate", projectController.create)
	projectRoutes.GET("info/:id", projectController.read)
	projectRoutes.POST("update", projectController.update)
	projectRoutes.DELETE("delete", projectController.delete)
	projectRoutes.POST("list", projectController.list)
	projectRoutes.POST("copy", projectController.duplicate)
	projectRoutes.POST("publish", projectController.publish)
	projectRoutes.POST("downloadScreen", projectController.export)
	projectRoutes.POST("downloadApp", projectController.compile)
	projectRoutes.POST("importScreen", projectController._import)
	projectRoutes.POST("move", projectController.moveGroup)
}

type GetProjectListRequest struct {
	shared.PaginationRequest
	// Workspace string
	Name  *string
	Group *string
}

func (c *ProjectController) list(ctx *gin.Context) {
	req := &GetProjectListRequest{}

	workspaceId, err := c.BindRequestJson(ctx, &req, "project list")

	if err != nil {
		return
	}

	infoList, err := service.ProjectServiceImpl.GetList(ctx, workspaceId, req.Page, req.Size, req.Name, req.Group)
	if err != nil {
		c.AbortServerError(ctx, "project list error: "+err.Error())
		return
	}
	c.ResponseJson(ctx, infoList)
}

/*CRUD*/
type CreateProjectRequest struct {
	Name string `json:"name" binding:"required"`
	// Version     string
	Config      string  `json:"root_config" binding:"required"`
	GroupCode   *string `json:"group_code"`
	UseTemplate *string `json:"template_id"` //template Id
}

func (c *ProjectController) create(ctx *gin.Context) {
	req := &CreateProjectRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		c.AbortClientError(ctx, "create read error: "+err.Error())
		return
	}
	projectCode, err := service.ProjectServiceImpl.Create(ctx, req.Name, req.Config, req.GroupCode, req.UseTemplate)
	if err != nil {
		c.AbortServerError(ctx, "create read error: "+err.Error())
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
		c.AbortClientError(ctx, "project read error: "+err.Error())
		return
	}

	projectDetail, err := service.ProjectServiceImpl.GetInfoById(ctx, query.Id)
	if err != nil {
		c.AbortServerError(ctx, "project read error: "+err.Error())
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

	if err := ctx.ShouldBindJSON(&param); err != nil {
		c.AbortClientError(ctx, "project update error: "+err.Error())
		return
	}

	if err := service.ProjectServiceImpl.Update(ctx, param.Code, param.Name, param.Config, param.Thumbnail, param.GroupCode); err != nil {
		c.AbortServerError(ctx, "project update error: "+err.Error())
		return
	}

	c.SuccessVoid(ctx)
}

type BatchMoveProjectGroupRequest struct {
	GroupName string   `json:"groupName"`
	GroupCode string   `json:"groupCode"`
	Ids       []string `json:"ids" binding:"required"`
}

func (c *ProjectController) moveGroup(ctx *gin.Context) {
	req := BatchMoveProjectGroupRequest{}

	workspaceId, err := c.BindRequestJson(ctx, &req, "project move group")

	if err != nil {
		return
	}

	if err := service.ProjectServiceImpl.BatchMoveGroup(ctx, workspaceId, req.Ids, req.GroupName, req.GroupCode); err != nil {
		c.AbortClientError(ctx, "move: "+err.Error())
		return
	}

	c.SuccessVoid(ctx)
}

type ProjectDeleteRequest struct {
	Ids []string `json:"ids" binding:"required"`
}

func (c *ProjectController) delete(ctx *gin.Context) {
	req := ProjectDeleteRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.AbortClientError(ctx, "project delete error: "+err.Error())
		return
	}

	if err := service.ProjectServiceImpl.LogicalDeletion(ctx, req.Ids); err != nil {
		c.AbortServerError(ctx, "project delete error: "+err.Error())
		return
	}
	c.ResponseJson(ctx, nil)
}

// features
func (c *ProjectController) duplicate(ctx *gin.Context) {
	var req struct {
		Id string `json:"id" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.AbortClientError(ctx, "project create error: "+err.Error())
		return
	}
	projectCode, err := service.ProjectServiceImpl.Duplicate(ctx, req.Id)
	if err != nil {
		c.AbortServerError(ctx, "project create error: "+err.Error())
		return
	}
	c.ResponseJson(ctx, projectCode)
}

type ProjectPublishRequest struct {
	Code             string  `json:"id"`
	PublishFlag      int     `json:"publish_flag"`
	PublishSecretKey *string `json:"public_secret_key"`
}

func (c *ProjectController) publish(ctx *gin.Context) {
	req := &ProjectPublishRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		c.AbortClientError(ctx, "project publish error: "+err.Error())
		return
	}

	if err := service.ProjectServiceImpl.PublishProject(ctx, req.Code, req.PublishFlag, req.PublishSecretKey); err != nil {
		c.AbortServerError(ctx, "project publish error: "+err.Error())
		return
	}
	c.ResponseJson(ctx, nil)
}

func (c *ProjectController) export(ctx *gin.Context) {

}

func (c *ProjectController) _import(ctx *gin.Context) {

}

func (c *ProjectController) compile(ctx *gin.Context) {
	// var req struct {
	// 	Code string `json:"id" binding:"required"`
	// 	Type string `json:"type" binding:"required"`
	// }
}

// func (c *ProjectController) getInteraction(ctx *gin.Context) {

// }
