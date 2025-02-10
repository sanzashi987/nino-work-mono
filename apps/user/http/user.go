package http

import (
	"fmt"
	iHttp "net/http"

	"github.com/gin-gonic/gin"
	userService "github.com/sanzashi987/nino-work/apps/user/service/user"
	"github.com/sanzashi987/nino-work/pkg/controller"
	"github.com/sanzashi987/nino-work/proto/user"
)

type UserController struct {
	controller.BaseController
}

var userController = UserController{}

type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Expiry   int32  `json:"expiry" binding:"required"`
}

func (c *UserController) UserLogin(ctx *gin.Context) {
	var req = UserLoginRequest{}
	var res = user.UserLoginResponse{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.AbortClientError(ctx, "[http] user login: Fail to read required fields"+err.Error())
		return
	}

	rpcReq := user.UserLoginRequest{
		Username: req.Username,
		Password: req.Password,
		Expiry:   &req.Expiry,
	}

	if err := userService.UserServiceRpcImpl.UserLogin(ctx, &rpcReq, &res); err != nil {
		c.AbortJson(ctx, int(res.Reason), "[rpc] user service: Login Error "+err.Error())
		return
	}

	expiry := int(res.Expiry)
	ctx.SetCookie(controller.CookieName, res.JwtToken, expiry*60*60*24, "/", ".nino.work", false, false)
	if target, shouldRedirect := ctx.GetQuery("redirect"); shouldRedirect {
		ctx.Redirect(iHttp.StatusSeeOther, fmt.Sprintf("%s?token=%s", target, res.JwtToken))
		return
	}

	c.ResponseJson(ctx, &res)
}

// func (c *UserController) UserRegister(ctx *gin.Context) {
// 	var req = user.UserRegisterRequest{}
// 	var res = user.UserLoginResponse{}
// 	if err := ctx.BindJSON(&req); err != nil {
// 		c.AbortClientError(ctx, "[http] user regiser: Fail to read required fields "+err.Error())
// 		return
// 	}

// 	if err := userService.UserServiceRpcImpl.UserRegister(ctx, &req, &res); err != nil {
// 		c.AbortJson(ctx, int(res.Reason), "[rpc] user service: Register error")
// 		return
// 	}

// 	c.ResponseJson(ctx, &res)
// }

func (c *UserController) UserInfo(ctx *gin.Context) {

	userId := ctx.GetUint64(controller.UserID)
	info, err := userService.UserServiceWebImpl.GetUserInfo(ctx, userId)

	if err != nil {
		c.AbortServerError(ctx, "[http] user info: Fail to read user info:"+err.Error())
		return
	}

	c.ResponseJson(ctx, info)

}

func (c *UserController) ListServiceUsers(ctx *gin.Context) {
	_ = ctx.GetUint64(controller.UserID)

}

func (c *UserController) TestToken(ctx *gin.Context) {
	return
}
