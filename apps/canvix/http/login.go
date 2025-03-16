package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/pkg/controller"
)

const login_prefix = "/api/auth/v1/"

type LoginController struct {
	controller.BaseController
}


var loginController = &LoginController{}




func (c *LoginController) getUuid(ctx *gin.Context) {

}

func (c *LoginController) login(ctx *gin.Context) {

}

func (c *LoginController) logout(ctx *gin.Context) {

}

func (c *LoginController) getVerifyCode(ctx *gin.Context) {

}

