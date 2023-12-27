package http

import (
	"fmt"
	iHttp "net/http"

	"github.com/cza14h/nino-work/apps/user/service"
	"github.com/cza14h/nino-work/pkg/controller"
	"github.com/cza14h/nino-work/proto/user"
	"github.com/gin-gonic/gin"
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
		controller.ResponseJson(
			ctx,
			iHttp.StatusBadRequest,
			"Fail to read required fields",
			nil,
		)
		return
	}

	if err := service.GetUserServiceRpc().UserLogin(ctx, &req, &res); err != nil {
		controller.ResponseJson(ctx, int(res.Reason), "Login Service Error", nil)
		return
	}

	if target, shouldRedirect := ctx.GetQuery("redirect"); shouldRedirect {
		ctx.Redirect(iHttp.StatusSeeOther, fmt.Sprintf("%s?token=%s", target, res.JwtToken))
		return
	}

	controller.ResponseJson(ctx, int(res.Reason), "", &res)
}

func (controller *UserController) UserRegister(ctx *gin.Context) {
	var req = user.UserRegisterRequest{}
	var res = user.UserRegisterResponse{
		JwtToken: voidStr,
	}
	if err := ctx.BindJSON(&req); err != nil {
		controller.ResponseJson(
			ctx,
			iHttp.StatusBadRequest,
			"Fail to read required fields",
			&res,
		)
		return
	}

	if err := service.GetUserServiceRpc().UserRegister(ctx, &req, &res); err != nil {
		controller.ResponseJson(ctx, int(res.Reason), "Register service error", nil)
		return
	}

	controller.ResponseJson(ctx, int(res.Reason), "", &res)
}
