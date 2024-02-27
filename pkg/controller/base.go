package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct{}

func (controler *BaseController) ResponseJson(c *gin.Context, code int, msg string, data interface{}) {

	c.JSON(http.StatusOK, gin.H{
		"msg":  msg,
		"data": data,
		"code": code,
	})

}

func (controller *BaseController) AbortError(c *gin.Context, code int, err error) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"msg":  err.Error(),
		"data": nil,
		"code": code,
	})
}

func (controller *BaseController) AbortJson(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"msg":  msg,
		"data": nil,
		"code": code,
	})
}

func (controller *BaseController) MustGetParam(c *gin.Context, key string) (string, error) {
	value := c.Param(key)
	if value == "" {
		resultStr := key + " is not provided in url"
		controller.AbortJson(c, http.StatusBadRequest, resultStr)
		return "", errors.New(value)
	}
	return value, nil
}
