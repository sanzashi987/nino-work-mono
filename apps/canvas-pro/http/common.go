package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/pkg/controller"
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
