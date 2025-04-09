package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/canvix/service"
)

type ThemeController struct {
	CanvixController
}

func registerThemeRoutes(router *gin.RouterGroup, loggedMiddleware, workspaceMiddleware gin.HandlerFunc) {
	themeController := ThemeController{}

	themeRoutes := router.Group("system-theme").Use(loggedMiddleware, workspaceMiddleware)

	themeRoutes.POST("list", themeController.list)
	themeRoutes.POST("create", themeController.create)
	themeRoutes.POST("update", themeController.update)
	themeRoutes.DELETE("delete", themeController.delete)

}

func (c *ThemeController) list(ctx *gin.Context) {
	_, workspaceId := getWorkspaceCode(ctx)
	res, err := service.GetThemes(ctx, workspaceId)
	if err != nil {
		c.AbortServerError(ctx, "theme list error: "+err.Error())
		return
	}

	c.ResponseJson(ctx, res)
}

func (c *ThemeController) update(ctx *gin.Context) {

	var req service.UpdateThemeReq

	workspaceId, err := c.BindRequestJson(ctx, &req, "theme update")

	if err != nil {
		return
	}

	if req.Config == nil && req.Name == nil {
		c.AbortClientError(ctx, "theme update error: should at least provide one property")
		return
	}

	if err := service.UpdateTheme(ctx, workspaceId, &req); err != nil {
		c.AbortServerError(ctx, "theme update error: "+err.Error())
		return
	}

	c.SuccessVoid(ctx)
}

func (c *ThemeController) create(ctx *gin.Context) {

	var req service.CreateThemeReq

	workspaceId, err := c.BindRequestJson(ctx, &req, "theme create")

	if err != nil {
		return
	}

	if err := service.CreateTheme(ctx, workspaceId, &req); err != nil {
		c.AbortServerError(ctx, "theme create error: "+err.Error())
		return
	}

	c.SuccessVoid(ctx)
}

func (c *ThemeController) delete(ctx *gin.Context) {

	var req service.DeleteThemeReq
	workspaceId, err := c.BindRequestJson(ctx, &req, "theme delete")
	if err != nil {
		return
	}

	if err := service.DeleteThemes(ctx, workspaceId, req.Data); err != nil {
		c.AbortServerError(ctx, "theme delete error: "+err.Error())
		return
	}

	c.SuccessVoid(ctx)
}
