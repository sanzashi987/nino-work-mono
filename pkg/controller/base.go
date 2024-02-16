package controller

import (
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

func (controller *BaseController) AbortJson(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"msg":  msg,
		"data": nil,
		"code": code,
	})
}
