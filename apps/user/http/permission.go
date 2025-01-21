package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/user/service"
	"github.com/sanzashi987/nino-work/pkg/controller"
)

type PermissionController struct {
	controller.BaseController
}

var permissionController = PermissionController{}

type ListPermissionRequest struct {
	AppId *uint64 `form:"app_id" binding:"required"`
}

func (c *PermissionController) ListPermissionsByApp(ctx *gin.Context) {
	userId := ctx.GetUint64(controller.UserID)

	req := ListPermissionRequest{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		c.AbortClientError(ctx, "[http] list permissions: Fail to read required fields "+err.Error())
		return
	}

	result, err := service.PermissionServiceWebImpl.ListPermissionsByApp(ctx, userId, req.AppId)

	if err != nil {
		c.AbortServerError(ctx, "[http] list permissions: service error"+err.Error())
		return
	}

	c.ResponseJson(ctx, result)
}

func (c *PermissionController) CreatePermission(ctx *gin.Context) {
	userId := ctx.GetUint64(controller.UserID)

	req := service.CreatePermissionRequest{}

	if err := ctx.ShouldBindQuery(&req); err != nil {
		c.AbortClientError(ctx, "[http] list permissions: Fail to read required fields "+err.Error())
		return
	}

	if err := service.PermissionServiceWebImpl.CreatePermission(ctx, userId, req); err != nil {
		c.AbortServerError(ctx, "[http] list permissions: service error"+err.Error())
		return
	}

	c.SuccessVoid(ctx)
}

func (c *PermissionController) DeletePermission(ctx *gin.Context) {
	userId := ctx.GetUint64(controller.UserID)
	req := service.DeletePermissionRequest{}

	if err := ctx.ShouldBindQuery(&req); err != nil {
		c.AbortClientError(ctx, "[http] list permissions: Fail to read required fields "+err.Error())
		return
	}

	if err := service.PermissionServiceWebImpl.DeletePermission(ctx, userId, req); err != nil {
		c.AbortServerError(ctx, "[http] list permissions: service error"+err.Error())
		return
	}

	c.SuccessVoid(ctx)
}
