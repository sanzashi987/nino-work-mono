package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/user/service"
	"github.com/sanzashi987/nino-work/pkg/auth"
	"github.com/sanzashi987/nino-work/pkg/controller"
)

type AppController struct {
	controller.BaseController
}

type AppInfoMeta struct {
	AppId       uint64 `json:"app_id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Status      uint   `json:"status"`
	Description string `json:"description"`
}

func (c *AppController) ListApps(ctx *gin.Context) {
	userId := ctx.GetUint64(auth.UserID)

	apps, err := service.AppServiceWebImpl.ListApplications(ctx, userId)
	if err != nil {
		c.AbortServerError(ctx, "[http] list apps: Fail to get app list:"+err.Error())
		return
	}

	res := []AppInfoMeta{}
	for _, app := range apps {
		res = append(res, AppInfoMeta{
			AppId:       app.Id,
			Name:        app.Name,
			Code:        app.Code,
			Status:      app.Status,
			Description: app.Description,
		})
	}

	c.ResponseJson(ctx, res)
}

func (c *AppController) CreateApp(ctx *gin.Context) {
	req := service.CreateAppRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.AbortClientError(ctx, "[http] create app: Fail to read required fields "+err.Error())
		return
	}

	userId := ctx.GetUint64(auth.UserID)
	app, err := service.AppServiceWebImpl.CreateApplication(ctx, userId, req)
	if err != nil {
		c.AbortServerError(ctx, "[http] create app: Fail to create app "+err.Error())
		return
	}

	c.ResponseJson(ctx, app.Id)
}
