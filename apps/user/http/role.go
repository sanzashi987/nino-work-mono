package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/user/db/model"
	"github.com/sanzashi987/nino-work/apps/user/service"
	"github.com/sanzashi987/nino-work/pkg/controller"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type RoleController struct {
	controller.BaseController
}

func RegisterRoleRoutes(router gin.IRoutes) {
	var roleController = RoleController{}
	router.POST("roles/create", roleController.createRole)
	router.POST("roles/list", roleController.listManagedRoles)
	router.POST("roles/update", roleController.updateRole)
	router.POST("roles/delete", roleController.deleteRole)
	router.GET("roles/suggest", roleController.suggest)
}

func (rc *RoleController) createRole(c *gin.Context) {
	var req service.CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		rc.AbortClientError(c, "[http] createRole: Fail to read required fields: "+err.Error())
		return
	}

	userId := c.GetUint64(controller.UserID)

	if err := service.RoleServiceWebImpl.CreateRole(c, userId, req); err != nil {
		rc.AbortServerError(c, "[http] createRole: Fail to create role: "+err.Error())
		return
	}

	rc.SuccessVoid(c)
}

func (rc *RoleController) listManagedRoles(c *gin.Context) {
	roleId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		rc.AbortClientError(c, "[http] listManagedRoles: Invalid role id")
		return
	}

	role, err := service.RoleServiceWebImpl.GetRoleDetail(c, roleId)
	if err != nil {
		rc.AbortServerError(c, "[http] listManagedRoles: Fail to get role detail: "+err.Error())
		return
	}

	rc.ResponseJson(c, role)
}

func (rc *RoleController) updateRole(c *gin.Context) {
	var req service.UpdateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		rc.AbortClientError(c, "[http] updateRole: Fail to read required fields: "+err.Error())
		return
	}

	if err := service.RoleServiceWebImpl.UpdateRole(c, req); err != nil {
		rc.AbortServerError(c, "[http] updateRole: Fail to update role: "+err.Error())
		return
	}

	rc.SuccessVoid(c)
}

func (rc *RoleController) deleteRole(c *gin.Context) {
	roleId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		rc.AbortClientError(c, "[http] deleteRole: Invalid role id")
		return
	}

	if err := service.RoleServiceWebImpl.DeleteRole(c, roleId); err != nil {
		rc.AbortServerError(c, "[http] deleteRole: Fail to delete role: "+err.Error())
		return
	}

	rc.SuccessVoid(c)
}

func (rc *RoleController) suggest(c *gin.Context) {
	keyword := c.Query("keyword")
	if keyword == "" {
		rc.ResponseJson(c, []model.RoleModel{})
		return
	}

	var roles []model.RoleModel
	if err := db.CommonSuggest[model.RoleModel](c, keyword, &roles); err != nil {
		rc.AbortClientError(c, "fail to suggest: "+err.Error())
	}

	rc.ResponseJson(c, roles)
}
