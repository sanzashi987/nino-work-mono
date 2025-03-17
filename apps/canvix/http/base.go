package http

import (
	"github.com/sanzashi987/nino-work/pkg/controller"

	"github.com/gin-gonic/gin"
)

type CanvixController struct {
	controller.BaseController
}

/** Also return the workspaceId */
func (c CanvixController) BindRequestJson(ctx *gin.Context, reqBody any, funcName string) (uint64, error) {
	if err := ctx.ShouldBindJSON(reqBody); err != nil {
		c.AbortClientError(ctx, funcName+" "+err.Error())
		return 0, err
	}
	_, workspaceId := getWorkspaceCode(ctx)

	return workspaceId, nil
}

func createCanvixController(errorPrefix string) CanvixController {
	return CanvixController{
		controller.BaseController{
			ErrorPrefix: errorPrefix,
		},
	}
}

/**** common prefix *****/
const listPrefix = "list: "
const readPreix = "read: "
const createPrefix = "create: "
const updatePrefix = "update: "
const deletePrefix = "delete: "
