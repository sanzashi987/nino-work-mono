package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/user/service"
	"github.com/sanzashi987/nino-work/pkg/controller"
)

type RoleController struct {
	controller.BaseController
}

func RegisterRoleRoutes(router gin.IRoutes) {
	var roleController = RoleController{}
	router.POST("roles/create", roleController.createRole)
	router.GET("roles/:id", roleController.getRoleDetail)
	router.POST("roles/update", roleController.updateRole)
	router.POST("roles/delete", roleController.deleteRole)
	router.GET("roles/suggest", roleController.suggestRoles)
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

func (rc *RoleController) getRoleDetail(c *gin.Context) {
	roleId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		rc.AbortClientError(c, "[http] getRoleDetail: Invalid role id")
		return
	}

	role, err := service.RoleServiceWebImpl.GetRoleDetail(c, roleId)
	if err != nil {
		rc.AbortServerError(c, "[http] getRoleDetail: Fail to get role detail: "+err.Error())
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

func (rc *RoleController) suggestRoles(c *gin.Context) {
	keyword := c.Query("keyword")
	roles, err := service.RoleServiceWebImpl.SuggestRoles(c, keyword)
	if err != nil {
		rc.AbortServerError(c, "[http] suggestRoles: Fail to suggest roles: "+err.Error())
		return
	}

	rc.ResponseJson(c, roles)
}
