package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/user/db/dao"
	"github.com/sanzashi987/nino-work/apps/user/db/model"
	userService "github.com/sanzashi987/nino-work/apps/user/service/user"
	"github.com/sanzashi987/nino-work/pkg/controller"
	"github.com/sanzashi987/nino-work/pkg/db"
	"github.com/sanzashi987/nino-work/pkg/shared"
)

type MenuController struct {
	controller.BaseController
}

func RegisterMenuRoutes(router gin.IRoutes) {
	var menuController = MenuController{}
	router.POST("menus/create", menuController.createMenu)
	router.GET("menus/:id", menuController.getMenuDetail)
	router.POST("menus/update", menuController.updateMenu)
	router.POST("menus/delete", menuController.deleteMenu)
	router.GET("menus/suggest", menuController.suggestMenus)
}

type CreateMenuRequest struct {
	Name   string   `json:"name" binding:"required"`
	Code   string   `json:"code" binding:"required"`
	Type   uint8    `json:"type" binding:"required"`
	Order  int      `json:"order"`
	Status int      `json:"status"`
	Path   string   `json:"path" binding:"required"`
	Roles  []uint64 `json:"roles"`
}

func (mc *MenuController) createMenu(c *gin.Context) {
	var req CreateMenuRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		mc.AbortClientError(c, "[http] createMenu: Fail to read required fields: "+err.Error())
		return
	}

	userId := c.GetUint64(controller.UserID)

	result, err := userService.GetUserAdmins(c, userId)
	if err != nil {
		mc.AbortServerError(c, "[http] createMenu: validate user err "+err.Error())
		return
	}

	if !result.HasAnyAdmin() {
		mc.AbortServerError(c, "[http] createMenu: user does not have any admin permission")
		return
	}

	hasPermission, err := result.ToPermissionSet()
	if err != nil {
		mc.AbortServerError(c, "[http] createMenu: fetch user permission error "+err.Error())
		return
	}

	bindRole := len(req.Roles) > 0

	if bindRole {
		tryToUsePermission, err := dao.FindAllPermissionsWithRoleIds(result.Tx, req.Roles)
		if err != nil {
			mc.AbortServerError(c, "[http] createMenu: read role permissions error "+err.Error())
			return
		}
		if !hasPermission.IsStrictlyContains(tryToUsePermission) {
			mc.AbortServerError(c, "[http] createMenu: user cannot create roles outside the permission range of admined apps")
			return
		}
	}

	toCreateMenu := model.MenuModel{
		Name:   req.Name,
		Code:   req.Code,
		Type:   model.MenuType(req.Type),
		Order:  req.Order,
		Status: req.Status,
		Path:   req.Path,
	}

	tx := result.Tx.Begin()

	if err := tx.Create(&toCreateMenu).Error; err != nil {
		tx.Rollback()
		mc.AbortServerError(c, "[http] createMenu: fail to create menu: "+err.Error())
		return
	}

	if bindRole {
		toBind := make([]*model.RoleModel, len(req.Roles))
		for index, id := range req.Roles {
			role := model.RoleModel{}
			role.Id = id
			toBind[index] = &role
		}

		if err := tx.Model(&toCreateMenu).Association("Roles").Replace(&toBind); err != nil {
			tx.Rollback()
			mc.AbortServerError(c, "[http] createMenu: fail to bind roles: "+err.Error())
			return
		}
	}

	mc.SuccessVoid(c)
}

type MenuDetail struct {
	Id          uint64             `json:"id"`
	Name        string             `json:"name"`
	Code        string             `json:"code"`
	Description string             `json:"description"`
	Type        model.MenuType     `json:"type"`
	Order       int                `json:"order"`
	Status      int                `json:"status"`
	Path        string             `json:"path"`
	Icon        string             `json:"icon"`
	Roles       []*shared.EnumMeta `json:"roles"`
}

func convertToMenuDetail(menu *model.MenuModel) MenuDetail {
	roles := make([]*shared.EnumMeta, len(menu.Roles))
	for i, role := range menu.Roles {
		roles[i] = &shared.EnumMeta{
			Value: role.Id,
			Name:  role.Name,
		}
	}

	return MenuDetail{
		Id:          menu.Id,
		Name:        menu.Name,
		Code:        menu.Code,
		Description: menu.Description,
		Type:        menu.Type,
		Order:       menu.Order,
		Status:      menu.Status,
		Path:        menu.Path,
		Icon:        menu.Icon,
		Roles:       roles,
	}
}

func (mc *MenuController) getMenuDetail(c *gin.Context) {
	menuId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		mc.AbortClientError(c, "[http] getMenuDetail: Invalid menu id")
		return
	}
	menu := model.MenuModel{}
	tx := db.NewTx(c)
	if err := tx.Preload("Roles", "menu_model_id = ?", menuId).Where("id = ?", menuId).Find(&menu).Error; err != nil {
		mc.AbortServerError(c, "[http] getMenuDetail: Fail to get menu detail: "+err.Error())
		return
	}

	mc.ResponseJson(c, convertToMenuDetail(&menu))
}

type UpdateMenuRequest struct {
	Id          uint64    `json:"id" binding:"required"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Order       *int      `json:"order"`
	Status      *int      `json:"status"`
	Path        string    `json:"path"`
	Roles       *[]uint64 `json:"roles"`
}

func (mc *MenuController) updateMenu(c *gin.Context) {
	var req UpdateMenuRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		mc.AbortClientError(c, "[http] updateMenu: fail to read required fields: "+err.Error())
		return
	}

	tx := db.NewTx(c)
	toBind := model.MenuModel{}
	if err := tx.Preload("Roles", "menu_model_id = ?", req.Id).Where("id = ?", req.Id).Find(&toBind).Error; err != nil {
		mc.AbortServerError(c, "[http] updateMenu: fail to get menu detail: "+err.Error())
		return
	}

	payload := model.MenuModel{}
	if req.Name != "" {
		payload.Name = req.Name
	}
	if req.Description != "" {
		payload.Description = req.Description
	}
	if req.Order != nil {
		payload.Order = *req.Order
	}
	if req.Status != nil {
		payload.Status = *req.Status
	}
	if req.Path != "" {
		payload.Path = req.Path
	}

	if req.Roles != nil {
		toReplace := *req.Roles

		currentRoles := db.ToIdList(toBind.Roles)
		// currentRoles := make([]db.IGetId, len(toBind.Roles))
		// for i, role := range toBind.Roles {
		// 	currentRoles[i] = role
		// }
	}

	tx = tx.Begin()
	if err := tx.Model(&toBind).Updates(&payload).Error; err != nil {
		tx.Rollback()
		mc.AbortServerError(c, "[http] updateMenu: fail to update menu: "+err.Error())
		return
	}

	mc.SuccessVoid(c)
}

func (mc *MenuController) deleteMenu(c *gin.Context) {
	menuId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		mc.AbortClientError(c, "[http] deleteMenu: Invalid menu id")
		return
	}

	if err := userService.MenuServiceWebImpl.DeleteMenu(c, menuId); err != nil {
		mc.AbortServerError(c, "[http] deleteMenu: fail to delete menu: "+err.Error())
		return
	}

	mc.SuccessVoid(c)
}

func (mc *MenuController) suggestMenus(c *gin.Context) {
	keyword := c.Query("keyword")
	menus, err := userService.MenuServiceWebImpl.SuggestMenus(c, keyword)
	if err != nil {
		mc.AbortServerError(c, "[http] suggestMenus: fail to suggest menus: "+err.Error())
		return
	}

	mc.ResponseJson(c, menus)
}
