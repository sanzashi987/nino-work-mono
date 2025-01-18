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
	AppId *uint64 `json:"app_id"`
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
