package http

import (
	"math"

	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/canvas-pro/http/request"
	"github.com/sanzashi987/nino-work/apps/user/service"
	"github.com/sanzashi987/nino-work/pkg/auth"
	"github.com/sanzashi987/nino-work/pkg/controller"
)

func ceil(num float64) int {
	intPart := int(num) // 获取整数部分
	if num > float64(intPart) {
		return intPart + 1 // 如果有小数部分，向上取整
	}
	return intPart // 如果本身是整数，直接返回
}

type AppController struct {
	controller.BaseController
}

type AppInfoMeta struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Status      uint   `json:"status"`
	Description string `json:"description"`
}

type ListAppResponse struct {
	Data []AppInfoMeta `json:"data"`
	request.PaginationResponse
}

func (c *AppController) ListApps(ctx *gin.Context) {
	userId := ctx.GetUint64(auth.UserID)

	pagination := request.PaginationRequest{}

	if err := ctx.ShouldBindQuery(&pagination); err != nil {
		c.AbortClientError(ctx, "[http] list apps: Fail to read required fields "+err.Error())
		return
	}

	apps, err := service.AppServiceWebImpl.ListApplications(ctx, userId)
	if err != nil {
		c.AbortServerError(ctx, "[http] list apps: Fail to get app list:"+err.Error())
		return
	}

	metas := []AppInfoMeta{}
	for _, app := range apps {
		metas = append(metas, AppInfoMeta{
			Id:          app.Id,
			Name:        app.Name,
			Code:        app.Code,
			Status:      app.Status,
			Description: app.Description,
		})
	}

	total := len(metas)
	pageIndex := pagination.Page
	pageSize := pagination.Size
	pageTotal := ceil(float64(total) / float64(pageSize))

	println(pagination.Page, pagination.Size, pageTotal)
	// Ensure pageIndex is within valid range
	if pageIndex < 1 {
		pageIndex = 1
	} else if pageIndex > pageTotal {
		pageIndex = pageTotal
	}

	start := (pageIndex - 1) * pageSize
	start = int(math.Max(0.0, float64(start)))
	end := start + pageSize

	if start > total {
		start = total
	}
	if end > total {
		end = total
	}

	pagedMetas := metas[start:end]

	res := ListAppResponse{}
	res.PageIndex = pageIndex
	res.PageSize = pageSize
	res.PageTotal = pageTotal
	res.RecordTotal = total
	res.Data = pagedMetas

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
