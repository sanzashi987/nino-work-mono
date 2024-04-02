package http

import (
	"github.com/cza14h/nino-work/pkg/controller"
	"github.com/gin-gonic/gin"
)

const common_prefix = "common"

type CommonController struct {
	controller.BaseController
}

var commonController = &CommonController{}

func (c *CommonController) searchComponents(ctx *gin.Context) {

}

func (c *CommonController) getUserInfo(ctx *gin.Context) {

}
