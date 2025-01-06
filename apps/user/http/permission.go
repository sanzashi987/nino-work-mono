package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/user/service"
	"github.com/sanzashi987/nino-work/pkg/auth"
	"github.com/sanzashi987/nino-work/pkg/controller"
)

type PermissionController struct {
	controller.BaseController
}

type ListPermissionRequest struct {
	AppId *uint64 `json:"app_id"`
}
type ListPermissionResponse struct {
	Permissions []any         `json:"permissions"`
	Apps        []AppInfoMeta `json:"apps"`
	SuperAdmin  bool          `json:"super_admin"`
	Admin       bool          `json:"admin"`
}

func (c *PermissionController) ListPermissionsByApps(ctx *gin.Context) {
	userId := ctx.GetUint64(auth.UserID)

	req := ListPermissionRequest{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		c.AbortClientError(ctx, "[http] list permissions: Fail to read required fields "+err.Error())
		return
	}

	result, err := service.PermissionServiceWebImpl.ListPermissionByApp(ctx, userId, req.AppId)

	if err != nil {
		c.AbortServerError(ctx, "[http] list permissions: service error"+err.Error())
		return
	}

	response := ListPermissionResponse{
		Permissions: []any{},
		Apps:        []AppInfoMeta{},
		SuperAdmin:  false,
		Admin:       false,
	}

	if result != nil {
		
	}

}
