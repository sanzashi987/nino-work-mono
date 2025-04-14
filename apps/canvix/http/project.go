package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/canvix/service/project"
	"github.com/sanzashi987/nino-work/pkg/shared"
)

type ProjectController struct {
	CanvixController
}

func registerProjectRoutes(router *gin.RouterGroup, loggedMiddleware, workspaceMiddleware gin.HandlerFunc) {
	projectController := ProjectController{}

	projectRoutes := router.Group("project").Use(loggedMiddleware, workspaceMiddleware)

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

	infoList, err := project.List(ctx, workspaceId, req.Page, req.Size, req.Name, req.Group)
	if err != nil {
		c.AbortServerError(ctx, "project list error: "+err.Error())
		return
	}
	c.ResponseJson(ctx, infoList)
}

/*CRUD*/

func (c *ProjectController) create(ctx *gin.Context) {
	req := &project.CreateProjectReq{}

	workspaceId, err := c.BindRequestJson(ctx, &req, "project create")

	if err != nil {
		return
	}

	projectCode, err := project.Create(ctx, workspaceId, req)
	if err != nil {
		c.AbortServerError(ctx, "project create error: "+err.Error())
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
	workspaceId := c.MustGetWorkspaceId(ctx)

	projectDetail, err := project.GetInfoById(ctx, workspaceId, query.Id)
	if err != nil {
		c.AbortServerError(ctx, "project read error: "+err.Error())
		return
	}

	c.ResponseJson(ctx, &projectDetail)
}

func (c *ProjectController) update(ctx *gin.Context) {
	req := project.ProjectUpdateReq{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.AbortClientError(ctx, "project update error: "+err.Error())
		return
	}
	workspaceId := c.MustGetWorkspaceId(ctx)

	if err := project.Update(ctx, workspaceId, &req); err != nil {
		c.AbortServerError(ctx, "project update error: "+err.Error())
		return
	}

	c.SuccessVoid(ctx)
}

type ProjectDeleteRequest struct {
	Ids []string `json:"ids" binding:"required"`
}

func (c *ProjectController) delete(ctx *gin.Context) {
	req := ProjectDeleteRequest{}

	workspaceId, err := c.BindRequestJson(ctx, &req, "project create")
	if err != nil {
		return
	}

	if err := project.Delete(ctx, workspaceId, req.Ids); err != nil {
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

	workspaceId, err := c.BindRequestJson(ctx, &req, "project duplicate")
	if err != nil {
		return
	}

	projectCode, err := project.Duplicate(ctx, workspaceId, req.Id)
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

	if err := project.PublishProject(ctx, req.Code, req.PublishFlag, req.PublishSecretKey); err != nil {
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
