package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/canvix/service"
	"github.com/sanzashi987/nino-work/pkg/controller"
	"github.com/sanzashi987/nino-work/pkg/shared"
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

type ThemeItem struct {
	Name  string `json:"name"`
	Theme string `json:"theme"`
	Id    uint64 `json:"id"`
	shared.DBTime
}

func (c *ThemeController) list(ctx *gin.Context) {
	_, workspaceId := getWorkspaceCode(ctx)
	models, err := service.ThemeServiceImpl.GetThemes(ctx, workspaceId)
	if err != nil {
		c.AbortServerError(ctx, updatePrefix+err.Error())
		return
	}

	res := []ThemeItem{}

	for _, model := range models {
		res = append(res, ThemeItem{
			DBTime: shared.DBTime{
				CreateTime: model.GetCreatedDate(),
				UpdateTime: model.GetUpdatedDate(),
			},
			Name:  model.Name,
			Id:    model.Id,
			Theme: model.Config,
		})
	}

	c.ResponseJson(ctx, res)
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
		c.AbortClientError(ctx, updatePrefix+err.Error())
		return
	}

	if reqBody.Config == nil && reqBody.Name == nil {
		c.AbortClientError(ctx, "update: should at least provide one property")
		return
	}

	if err := service.ThemeServiceImpl.UpdateTheme(ctx, workspaceId, reqBody.Id, reqBody.Name, reqBody.Config); err != nil {
		c.AbortServerError(ctx, updatePrefix+err.Error())
		return
	}

	c.SuccessVoid(ctx)
}

type CreateThemeReq struct {
	Name   string `json:"name" binding:"required"`
	Config string `json:"theme" binding:"required"`
}

func (c *ThemeController) create(ctx *gin.Context) {

	req := CreateThemeReq{}
	if err := ctx.BindJSON(&req); err != nil {
		c.AbortClientError(ctx, createPrefix+err.Error())
		return
	}

	_, workspaceId := getWorkspaceCode(ctx)

	if err := service.ThemeServiceImpl.CreateTheme(ctx, workspaceId, req.Name, req.Config); err != nil {
		c.AbortServerError(ctx, createPrefix+err.Error())
		return
	}

	c.SuccessVoid(ctx)

}

type DeleteThemeReq struct {
	Data []uint64 `json:"data" binding:"required"`
}

func (c *ThemeController) delete(ctx *gin.Context) {

	req := DeleteThemeReq{}
	if err := ctx.BindJSON(&req); err != nil {
		c.AbortClientError(ctx, deletePrefix+err.Error())
		return
	}
	_, workspaceId := getWorkspaceCode(ctx)

	if err := service.ThemeServiceImpl.DeleteThemes(ctx, workspaceId, req.Data); err != nil {
		c.AbortServerError(ctx, deletePrefix+err.Error())
		return
	}

	c.SuccessVoid(ctx)

}
