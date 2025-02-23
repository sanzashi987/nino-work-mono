package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/user/service"
	"github.com/sanzashi987/nino-work/pkg/controller"
)

type PermissionController struct {
	controller.BaseController
}

func RegisterAppPermissionRoutes(router gin.IRoutes) {
	var permissionController = PermissionController{}
	router.GET("apps/permission/list", permissionController.ListPermissionsByApp)
	router.POST("apps/permission/create", permissionController.CreatePermission)
	router.POST("apps/permission/delete", permissionController.DeletePermission)
	router.POST("apps/permission/delete", permissionController.DeletePermission)
}

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

	result, err := service.ListPermissionsByApp(ctx, userId, req.AppId)

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

	if err := service.CreatePermission(ctx, userId, req); err != nil {
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

	if err := service.DeletePermission(ctx, userId, req); err != nil {
		c.AbortServerError(ctx, "[http] list permissions: service error"+err.Error())
		return
	}

	c.SuccessVoid(ctx)
}

func (c *PermissionController) GetAdministratedPermissions(ctx *gin.Context) {
	userId := ctx.GetUint64(controller.UserID)

	result, err := service.GetAdministratedPermissions(ctx, userId)

	if err != nil {
		c.AbortServerError(ctx, "[http] list permissions: service error"+err.Error())
		return
	}

	c.ResponseJson(ctx, result)
}
