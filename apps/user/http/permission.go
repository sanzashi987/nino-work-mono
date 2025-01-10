package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/user/db/model"
	"github.com/sanzashi987/nino-work/apps/user/service"
	"github.com/sanzashi987/nino-work/pkg/controller"
)

type PermissionController struct {
	controller.BaseController
}

type BasicInfo struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

func intoPermissionInfo(permission *model.PermissionModel) BasicInfo {
	return BasicInfo{
		Id:          permission.Id,
		Name:        permission.Name,
		Code:        permission.Code,
		Description: permission.Description,
	}
}

type ListPermissionRequest struct {
	AppId *uint64 `json:"app_id"`
}
type ListPermissionResponse struct {
	Permissions []BasicInfo `json:"permissions"`
	Apps        []AppInfo   `json:"apps"`
	SuperAdmin  bool        `json:"super_admin"`
	Admin       bool        `json:"admin"`
}

func (c *PermissionController) ListPermissionsByApps(ctx *gin.Context) {
	userId := ctx.GetUint64(controller.UserID)

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
		Permissions: []BasicInfo{},
		Apps:        []AppInfo{},
		SuperAdmin:  false,
		Admin:       false,
	}

	if result != nil {
		response.SuperAdmin, response.Admin = result.FromSuper, result.FromAdmin

		apps := []AppInfo{}
		for _, app := range result.AppList {
			apps = append(apps, intoAppInfoMeta(app))
		}
		permissions := []BasicInfo{}
		for _, app := range result.App.Permissions {
			permissions = append(permissions, intoPermissionInfo(&app))
		}
		response.Permissions, response.Apps = permissions, apps
	}

	c.ResponseJson(ctx, response)
}
