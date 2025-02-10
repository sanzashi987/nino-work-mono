package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/user/db/model"
	menuService "github.com/sanzashi987/nino-work/apps/user/service/menu"
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
	menuController.ErrorPrefix = "[http] "
	router.POST("menus/create", menuController.createMenu)
	router.GET("menus/:id", menuController.getMenuDetail)
	router.POST("menus/update", menuController.updateMenu)
	router.POST("menus/delete", menuController.deleteMenu)
	router.GET("menus/suggest", menuController.suggestMenus)
}

func (mc *MenuController) createMenu(c *gin.Context) {
	var req menuService.CreateMenuRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		mc.AbortClientError(c, "createMenu: Fail to read required fields: "+err.Error())
		return
	}

	userId := c.GetUint64(controller.UserID)
	if err := menuService.Create(c, userId, &req); err != nil {
		mc.AbortClientError(c, "createMenu: Fail to read required fields: "+err.Error())
		return
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
		mc.AbortClientError(c, "getMenuDetail: Invalid menu id")
		return
	}
	menu := model.MenuModel{}
	tx := db.NewTx(c)
	if err := tx.Preload("Roles", "menu_model_id = ?", menuId).Where("id = ?", menuId).Find(&menu).Error; err != nil {
		mc.AbortServerError(c, "getMenuDetail: Fail to get menu detail: "+err.Error())
		return
	}

	mc.ResponseJson(c, convertToMenuDetail(&menu))
}

func (mc *MenuController) updateMenu(c *gin.Context) {
	var req menuService.UpdateMenuRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		mc.AbortClientError(c, "updateMenu: fail to read required fields: "+err.Error())
		return
	}
	userId := c.GetUint64(controller.UserID)
	if err := menuService.Update(c, userId, &req); err != nil {
		mc.AbortServerError(c, "updateMenu: internal server error "+err.Error())
		return
	}
	mc.SuccessVoid(c)
}

func (mc *MenuController) deleteMenu(c *gin.Context) {
	menuId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		mc.AbortClientError(c, "deleteMenu: Invalid menu id")
		return
	}

	userId := c.GetUint64(controller.UserID)

	if err := menuService.Delete(c, userId, menuId); err != nil {
		mc.AbortServerError(c, "deleteMenu: fail to delete menu: "+err.Error())
		return
	}

	mc.SuccessVoid(c)
}

func (mc *MenuController) suggestMenus(c *gin.Context) {
	keyword := c.Query("keyword")
	menus, err := userService.MenuServiceWebImpl.SuggestMenus(c, keyword)
	if err != nil {
		mc.AbortServerError(c, "suggestMenus: fail to suggest menus: "+err.Error())
		return
	}

	mc.ResponseJson(c, menus)
}
