package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
	ErrorPrefix string
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

func (controller *BaseController) AbortClientError(ctx *gin.Context, errMsg string) {
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"msg":  controller.ErrorPrefix + errMsg,
		"data": nil,
		"code": http.StatusBadRequest,
	})
}

func (controller *BaseController) AbortServerError(ctx *gin.Context, errMsg string) {
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"msg":  controller.ErrorPrefix + errMsg,
		"data": nil,
		"code": http.StatusInternalServerError,
	})
}

func (controller *BaseController) AbortJson(ctx *gin.Context, reasonCode int, msg string) {
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"msg":  msg,
		"data": nil,
		"code": reasonCode,
	})
}

func (controller *BaseController) MustGetParam(ctx *gin.Context, key string) (string, error) {
	value := ctx.Param(key)
	if value == "" {
		resultStr := key + " is not provided in url"
		controller.AbortClientError(ctx, resultStr)
		return "", errors.New(resultStr)
	}
	return value, nil
}
