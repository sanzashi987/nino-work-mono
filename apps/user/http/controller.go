package http

import (
	"fmt"
	iHttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/user/service"
	"github.com/sanzashi987/nino-work/pkg/controller"
	"github.com/sanzashi987/nino-work/proto/user"
)

type UserController struct {
	controller.BaseController
}

const voidStr = ""

func (controller *UserController) UserLogin(ctx *gin.Context) {
	var req = user.UserLoginRequest{}
	var res = user.UserLoginResponse{
		JwtToken: voidStr,
	}
	if err := ctx.BindJSON(&req); err != nil {
		controller.AbortClientError(ctx, "[http] user login: Fail to read required fields")
		return
	}

	if err := service.GetUserServiceRpc().UserLogin(ctx, &req, &res); err != nil {
		controller.AbortJson(ctx, int(res.Reason), "[rpc] user service: Login Error")
		return
	}

	if target, shouldRedirect := ctx.GetQuery("redirect"); shouldRedirect {
		ctx.Redirect(iHttp.StatusSeeOther, fmt.Sprintf("%s?token=%s", target, res.JwtToken))
		return
	}

	controller.ResponseJson(ctx, &res)
}

func (controller *UserController) UserRegister(ctx *gin.Context) {
	var req = user.UserRegisterRequest{}
	var res = user.UserRegisterResponse{
		JwtToken: voidStr,
	}
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
