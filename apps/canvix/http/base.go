package http

import (
	"github.com/sanzashi987/nino-work/pkg/controller"

	"github.com/gin-gonic/gin"
)

type CanvixController struct {
	controller.BaseController
}

/** Also return the workspaceId */
func (c *CanvixController) BindRequestJson(ctx *gin.Context, reqBody any, funcName string) (uint64, error) {
	if err := ctx.ShouldBindJSON(reqBody); err != nil {
		c.AbortClientError(ctx, "[http] "+funcName+" error: "+err.Error())
		return 0, err
	}
	workspaceId := c.MustGetWorkspaceId(ctx)

	return workspaceId, nil
}

func (c *CanvixController) MustGetWorkspaceId(ctx *gin.Context) uint64 {
	workspaceId := ctx.MustGet(WORKSPACE_ID).(uint64)
	return workspaceId
}
