package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
	ErrorPrefix string
}

func (controler *BaseController) GetErrorPrefix() string {
	if controler.ErrorPrefix == "" {
		return "[http] "
	}
	return controler.ErrorPrefix
}

func (controler *BaseController) ResponseJson(ctx *gin.Context, data interface{}) {

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": data,
		"code": 0,
	})

}
func (controler *BaseController) SuccessVoid(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": nil,
		"code": 0,
	})

}

func (c *BaseController) AbortClientError(ctx *gin.Context, errMsg string) {
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"msg":  c.GetErrorPrefix() + errMsg,
		"data": nil,
		"code": http.StatusBadRequest,
	})
}

func (c *BaseController) AbortServerError(ctx *gin.Context, errMsg string) {
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"msg":  c.GetErrorPrefix() + errMsg,
		"data": nil,
		"code": http.StatusInternalServerError,
	})
}

func (c *BaseController) AbortServerErrorWithCode(ctx *gin.Context, reasonCode int, msg string) {
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"msg":  msg,
		"data": nil,
		"code": reasonCode,
	})
}
