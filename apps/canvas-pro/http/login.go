package http

import (
	"github.com/cza14h/nino-work/pkg/controller"
	"github.com/gin-gonic/gin"
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

