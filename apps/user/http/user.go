package http

import (
	"fmt"
	iHttp "net/http"

	"github.com/gin-gonic/gin"
	userService "github.com/sanzashi987/nino-work/apps/user/service/user"
	"github.com/sanzashi987/nino-work/pkg/controller"
	"github.com/sanzashi987/nino-work/pkg/shared"
	"github.com/sanzashi987/nino-work/proto/user"
)

type UserController struct {
	controller.BaseController
}

func RegisterUserRoutes(public, authed gin.IRoutes) {
	var userController = UserController{}

	public.POST("users/login", userController.UserLogin)
	authed.GET("users/info", userController.UserInfo)
	authed.GET("users/user-roles", userController.GetUserRoles)
	authed.POST("users/list", userController.ListUser)
	authed.POST("users/bind-roles", userController.BindUserRoles)
	authed.POST("users/create", userController.CreateUserByAdmin)
}

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
		c.AbortServerErrorWithCode(ctx, int(res.Reason), "[rpc] user service: Login Error "+err.Error())
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
// 		c.AbortServerErrorWithCode(ctx, int(res.Reason), "[rpc] user service: Register error")
// 		return
// 	}

// 	c.ResponseJson(ctx, &res)
// }

func (c *UserController) CreateUserByAdmin(ctx *gin.Context) {
	userId := ctx.GetUint64(controller.UserID)

	req := userService.CreateUserRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.AbortClientError(ctx, "[http] user create: Fail to read required fields "+err.Error())
		return
	}

	id, err := userService.CreateUserByAdmin(ctx, userId, &req)
	if err != nil {
		c.AbortServerError(ctx, "[rpc] user service: Create user error "+err.Error())
		return
	}

	c.ResponseJson(ctx, gin.H{"id": id})
}

func (c *UserController) UserInfo(ctx *gin.Context) {

	userId := ctx.GetUint64(controller.UserID)
	info, err := userService.GetUserInfo(ctx, userId)

	if err != nil {
		c.AbortServerError(ctx, "[http] user info: Fail to read user info:"+err.Error())
		return
	}

	c.ResponseJson(ctx, info)

}

func (c *UserController) ListUser(ctx *gin.Context) {
	var req shared.PaginationRequest
	if err := ctx.ShouldBind(&req); err != nil {
		c.AbortClientError(ctx, "[http] user list: Fail to read required fields "+err.Error())
		return
	}

	res, err := userService.ListUser(ctx, &req)
	if err != nil {
		c.AbortServerError(ctx, "[http] user list: Fail to read user list "+err.Error())
		return
	}

	c.ResponseJson(ctx, res)

}

func (c *UserController) BindUserRoles(ctx *gin.Context) {
	userId := ctx.GetUint64(controller.UserID)

	var req userService.BindRoleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.AbortClientError(ctx, "[http] user bind roles: Fail to read required fields "+err.Error())
		return
	}

	if err := userService.BindUserRoles(ctx, userId, &req); err != nil {
		c.AbortServerError(ctx, "[http] user bind roles: Fail to bind user roles "+err.Error())
		return
	}

	c.SuccessVoid(ctx)
}

func (c *UserController) GetUserRoles(ctx *gin.Context) {
	userId := ctx.GetUint64(controller.UserID)

	var req struct {
		UserId uint64 `form:"id" binding:"required"`
	}

	if err := ctx.ShouldBindQuery(&req); err != nil {
		c.AbortClientError(ctx, "[http] user get roles: Fail to read required fields "+err.Error())
		return
	}

	res, err := userService.GetUserRoles(ctx, userId, req.UserId)
	if err != nil {
		c.AbortServerError(ctx, "[http] user get roles: Fail to get user roles "+err.Error())
		return
	}

	c.ResponseJson(ctx, res)
}
