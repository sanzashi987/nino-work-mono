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

var roleController = RoleController{}

func RegisterRoleRoutes(router gin.IRoutes) {
	router.POST("roles/create", roleController.createRole)
	router.GET("roles/:id", roleController.getRoleDetail)
	router.POST("roles/update", roleController.updateRole)
	router.POST("roles/delete", roleController.deleteRole)
	router.GET("roles/suggest", roleController.suggestRoles)
}

func (rc *RoleController) createRole(c *gin.Context) {
	var req service.CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		rc.AbortClientError(c, "Fail to read required fields: "+err.Error())
		return
	}

	userId := c.GetUint64(controller.UserID)

	if err := service.RoleServiceWebImpl.CreateRole(c.Request.Context(), userId, req); err != nil {
		rc.AbortServerError(c, "Fail to create role: "+err.Error())
		return
	}

	rc.SuccessVoid(c)
}

func (rc *RoleController) getRoleDetail(c *gin.Context) {
	roleId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		rc.AbortClientError(c, "Invalid role id")
		return
	}

	role, err := service.RoleServiceWebImpl.GetRoleDetail(c.Request.Context(), roleId)
	if err != nil {
		rc.AbortServerError(c, "Fail to get role detail: "+err.Error())
		return
	}

	rc.ResponseJson(c, role)
}

func (rc *RoleController) updateRole(c *gin.Context) {
	roleId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		rc.AbortClientError(c, "Invalid role id")
		return
	}

	var req service.CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		rc.AbortClientError(c, "Fail to read required fields: "+err.Error())
		return
	}

	if err := service.RoleServiceWebImpl.UpdateRole(c.Request.Context(), roleId, req); err != nil {
		rc.AbortServerError(c, "Fail to update role: "+err.Error())
		return
	}

	rc.SuccessVoid(c)
}

func (rc *RoleController) deleteRole(c *gin.Context) {
	roleId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		rc.AbortClientError(c, "Invalid role id")
		return
	}

	if err := service.RoleServiceWebImpl.DeleteRole(c.Request.Context(), roleId); err != nil {
		rc.AbortServerError(c, "Fail to delete role: "+err.Error())
		return
	}

	rc.SuccessVoid(c)
}

func (rc *RoleController) suggestRoles(c *gin.Context) {
	keyword := c.Query("keyword")
	roles, err := service.RoleServiceWebImpl.SuggestRoles(c.Request.Context(), keyword)
	if err != nil {
		rc.AbortServerError(c, "Fail to suggest roles: "+err.Error())
		return
	}

	rc.ResponseJson(c, roles)
}
