package http

import (
	"github.com/sanzashi987/nino-work/pkg/controller"

	"github.com/gin-gonic/gin"
)

type CanvasController struct {
	controller.BaseController
}

/** Also return the workspaceId */
func (c CanvasController) BindRequestJson(ctx *gin.Context, reqBody any, funcName string) (uint64, error) {
	if err := ctx.BindJSON(reqBody); err != nil {
		c.AbortClientError(ctx, funcName+" "+err.Error())
		return 0, err
	}
	_, workspaceId := getWorkspaceCode(ctx)

	return workspaceId, nil
}

func createCanvasController(errorPrefix string) CanvasController {
	return CanvasController{
		controller.BaseController{
			ErrorPrefix: errorPrefix,
		},
	}
}
