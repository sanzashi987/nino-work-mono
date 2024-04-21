package http

import (
	"github.com/cza14h/nino-work/apps/canvas-pro/service"
	"github.com/cza14h/nino-work/pkg/controller"
	"github.com/gin-gonic/gin"
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

}

type UpdateThemeReq struct {
	Id     uint64  `json:"id" binding:"required"`
	Name   *string `json:"name"`
	Config *string `json:"config"`
}

func (c *ThemeController) update(ctx *gin.Context) {
	_, workspaceId := getWorkspaceCode(ctx)

	reqBody := UpdateThemeReq{}

	if err := ctx.BindJSON(&reqBody); err != nil {
		c.AbortClientError(ctx, "update: "+err.Error())
		return
	}

	if reqBody.Config == nil && reqBody.Name == nil {
		c.AbortClientError(ctx, "update: should at least provide one property")
		return
	}

	if err := service.ThemeServiceImpl.UpdateTheme(ctx, workspaceId, reqBody.Id, reqBody.Name, reqBody.Config); err != nil {
		c.AbortServerError(ctx, "update: "+err.Error())
		return
	}

	c.SuccessVoid(ctx)
}
func (c *ThemeController) create(ctx *gin.Context) {
	_, workspaceId := getWorkspaceCode(ctx)

}
func (c *ThemeController) delete(ctx *gin.Context) {

	_, workspaceId := getWorkspaceCode(ctx)

}
