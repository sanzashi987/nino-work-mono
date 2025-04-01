package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/canvix/service"
	"github.com/sanzashi987/nino-work/pkg/controller"
)

type ThemeController struct {
	controller.BaseController
}

const theme_prefix = "system-theme"

var themeController = &ThemeController{
	controller.BaseController{
		ErrorPrefix: "[http] canvas theme handler ",
	},
}

func (c *ThemeController) list(ctx *gin.Context) {
	_, workspaceId := getWorkspaceCode(ctx)
	res, err := service.GetThemes(ctx, workspaceId)
	if err != nil {
		c.AbortServerError(ctx, updatePrefix+err.Error())
		return
	}

	c.ResponseJson(ctx, res)
}

func (c *ThemeController) update(ctx *gin.Context) {
	_, workspaceId := getWorkspaceCode(ctx)

	var req service.UpdateThemeReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.AbortClientError(ctx, updatePrefix+err.Error())
		return
	}

	if req.Config == nil && req.Name == nil {
		c.AbortClientError(ctx, "update: should at least provide one property")
		return
	}

	if err := service.UpdateTheme(ctx, workspaceId, &req); err != nil {
		c.AbortServerError(ctx, updatePrefix+err.Error())
		return
	}

	c.SuccessVoid(ctx)
}

func (c *ThemeController) create(ctx *gin.Context) {

	var req service.CreateThemeReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.AbortClientError(ctx, createPrefix+err.Error())
		return
	}

	_, workspaceId := getWorkspaceCode(ctx)

	if err := service.CreateTheme(ctx, workspaceId, &req); err != nil {
		c.AbortServerError(ctx, createPrefix+err.Error())
		return
	}

	c.SuccessVoid(ctx)

}

func (c *ThemeController) delete(ctx *gin.Context) {

	var req service.DeleteThemeReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.AbortClientError(ctx, deletePrefix+err.Error())
		return
	}
	_, workspaceId := getWorkspaceCode(ctx)

	if err := service.DeleteThemes(ctx, workspaceId, req.Data); err != nil {
		c.AbortServerError(ctx, deletePrefix+err.Error())
		return
	}

	c.SuccessVoid(ctx)

}
