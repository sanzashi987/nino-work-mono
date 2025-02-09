package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
	userService "github.com/sanzashi987/nino-work/apps/user/service/user"
	"github.com/sanzashi987/nino-work/pkg/controller"
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

func (mc *MenuController) createMenu(c *gin.Context) {
	var req userService.CreateMenuRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		mc.AbortClientError(c, "[http] createMenu: Fail to read required fields: "+err.Error())
		return
	}

	userId := c.GetUint64(controller.UserID)

	if err := userService.MenuServiceWebImpl.CreateMenu(c, userId, req); err != nil {
		mc.AbortServerError(c, "[http] createMenu: Fail to create menu: "+err.Error())
		return
	}

	mc.SuccessVoid(c)
}

func (mc *MenuController) getMenuDetail(c *gin.Context) {
	menuId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		mc.AbortClientError(c, "[http] getMenuDetail: Invalid menu id")
		return
	}

	menu, err := userService.MenuServiceWebImpl.GetMenuDetail(c, menuId)
	if err != nil {
		mc.AbortServerError(c, "[http] getMenuDetail: Fail to get menu detail: "+err.Error())
		return
	}

	mc.ResponseJson(c, menu)
}

func (mc *MenuController) updateMenu(c *gin.Context) {
	var req userService.UpdateMenuRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		mc.AbortClientError(c, "[http] updateMenu: Fail to read required fields: "+err.Error())
		return
	}

	if err := userService.MenuServiceWebImpl.UpdateMenu(c, req); err != nil {
		mc.AbortServerError(c, "[http] updateMenu: Fail to update menu: "+err.Error())
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
		mc.AbortServerError(c, "[http] deleteMenu: Fail to delete menu: "+err.Error())
		return
	}

	mc.SuccessVoid(c)
}

func (mc *MenuController) suggestMenus(c *gin.Context) {
	keyword := c.Query("keyword")
	menus, err := userService.MenuServiceWebImpl.SuggestMenus(c, keyword)
	if err != nil {
		mc.AbortServerError(c, "[http] suggestMenus: Fail to suggest menus: "+err.Error())
		return
	}

	mc.ResponseJson(c, menus)
}
