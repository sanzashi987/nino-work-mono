package http

import (
	"fmt"
	iHttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/user/service"
	"github.com/sanzashi987/nino-work/pkg/auth"
	"github.com/sanzashi987/nino-work/pkg/controller"
	"github.com/sanzashi987/nino-work/proto/user"
)

type UserController struct {
	controller.BaseController
}

func (controller *UserController) UserLogin(ctx *gin.Context) {
	var req = user.UserLoginRequest{}
	var res = user.UserLoginResponse{}
	if err := ctx.BindJSON(&req); err != nil {
		controller.AbortClientError(ctx, "[http] user login: Fail to read required fields")
		return
	}

	if err := service.GetUserServiceRpc().UserLogin(ctx, &req, &res); err != nil {
		controller.AbortJson(ctx, int(res.Reason), "[rpc] user service: Login Error")
		return
	}

	expiry := int(res.Expiry)
	ctx.SetCookie(auth.CookieName, res.JwtToken, expiry*60*60*24, "/", ".nino.work", false, false)
	if target, shouldRedirect := ctx.GetQuery("redirect"); shouldRedirect {
		ctx.Redirect(iHttp.StatusSeeOther, fmt.Sprintf("%s?token=%s", target, res.JwtToken))
		return
	}

	controller.ResponseJson(ctx, &res)
}

func (controller *UserController) UserRegister(ctx *gin.Context) {
	var req = user.UserRegisterRequest{}
	var res = user.UserLoginResponse{}
	if err := ctx.BindJSON(&req); err != nil {
		controller.AbortClientError(ctx, "[http] user regiser: Fail to read required fields "+err.Error())
		return
	}

	if err := service.GetUserServiceRpc().UserRegister(ctx, &req, &res); err != nil {
		controller.AbortJson(ctx, int(res.Reason), "[rpc] user service: Register error")
		return
	}

	controller.ResponseJson(ctx, &res)
}

func (controller *UserController) Dashboard(ctx *gin.Context) {

}
