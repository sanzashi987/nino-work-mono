package http

import (
	"github.com/gin-gonic/gin"
)

type CommonController struct {
	CanvixController
}

func registerCommonRoutes(router *gin.RouterGroup, loggedMiddleware, workspaceMiddleware gin.HandlerFunc) {

	commonController := CommonController{}
	nonAuthed := router.Group("common")
	authed := nonAuthed.Use(loggedMiddleware)

	authed.GET("user", commonController.getUserInfo)
	authed.GET("workspace", commonController.GetWorkspaceInfo)

}

// func (c *CommonController) searchComponents(ctx *gin.Context) {

// }

func (c *CommonController) getUserInfo(ctx *gin.Context) {

}

func (c *CommonController) GetWorkspaceInfo(ctx *gin.Context) {

}
