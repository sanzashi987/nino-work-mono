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
		Fail:     false,
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

	service.GetUserServiceRpc().UserLogin(ctx, &req, &res)

	if target, shouldRedirect := ctx.GetQuery("redirect"); shouldRedirect && !res.Fail {
		ctx.Redirect(iHttp.StatusSeeOther, fmt.Sprintf("%s?token=%s", target, res.JwtToken))
		return
	}

	controller.ResponseJson(ctx, iHttp.StatusOK, "", &res)
}

func (controller *UserController) UserRegister(ctx *gin.Context) {
	var req = user.UserRegisterRequest{}
	var res = user.UserRegisterResponse{
		Fail:     false,
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
	//TODO extract reason code to response
	service.GetUserServiceRpc().UserRegister(ctx, &req, &res)
	controller.ResponseJson(ctx, iHttp.StatusOK, "", &res)
}
