package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
	ErrorPrefix string
}

func (controler *BaseController) ResponseJson(c *gin.Context, data interface{}) {

	c.JSON(http.StatusOK, gin.H{
		"msg":  "Success",
		"data": data,
		"code": 0,
	})

}
func (controler *BaseController) SuccessVoid(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"msg":  "Success",
		"data": nil,
		"code": 0,
	})

}

func (controller *BaseController) AbortClientError(c *gin.Context, errMsg string) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"msg":  controller.ErrorPrefix + errMsg,
		"data": nil,
		"code": http.StatusBadRequest,
	})
}

func (controller *BaseController) AbortServerError(c *gin.Context, errMsg string) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"msg":  controller.ErrorPrefix + errMsg,
		"data": nil,
		"code": http.StatusInternalServerError,
	})
}

func (controller *BaseController) AbortJson(c *gin.Context, reasonCode int, msg string) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"msg":  msg,
		"data": nil,
		"code": reasonCode,
	})
}

func (controller *BaseController) MustGetParam(c *gin.Context, key string) (string, error) {
	value := c.Param(key)
	if value == "" {
		resultStr := key + " is not provided in url"
		controller.AbortClientError(c, resultStr)
		return "", errors.New(resultStr)
	}
	return value, nil
}
