package http

import (
	"github.com/cza14h/nino-work/pkg/controller"
	"github.com/gin-gonic/gin"
)

type ThemeController struct {
	controller.BaseController
}

const theme_prefix = "system-theme"

var themeController = &ThemeController{}

func (c *ThemeController) list(ctx *gin.Context) {

}

func (c *ThemeController) update(ctx *gin.Context) {

}
func (c *ThemeController) create(ctx *gin.Context) {

}
func (c *ThemeController) delete(ctx *gin.Context) {

}
