package http

import (
	"net/http"
	"strings"

	"github.com/cza14h/nino-work/pkg/controller"
	"github.com/gin-gonic/gin"
)

const project_prefix = "screen-operation"

type ProjectController struct {
	controller.BaseController
}

func (c *ProjectController) list(ctx *gin.Context) {

}

/*CRUD*/

type ProjectCreateRequest struct {
	Name      string
	Version   string
	GroupCode string `json:"groupCode"`
	Config    string `json:"rootConfig"`
	
}

func (c *ProjectController) create(ctx *gin.Context) {

}
func (c *ProjectController) read(ctx *gin.Context) {

}
func (c *ProjectController) update(ctx *gin.Context) {

}
func (c *ProjectController) delete(ctx *gin.Context) {

}

// features
func (c *ProjectController) duplicate(ctx *gin.Context) {
	id := ctx.Param("id")
	if trimmedId := strings.Trim(id, "/"); trimmedId != "" {

		return
	}

	c.AbortJson(ctx, http.StatusBadRequest, "Id is required", nil)
}

func (c *ProjectController) publish(ctx *gin.Context) {

}

func (c *ProjectController) export(ctx *gin.Context) {

}

func (c *ProjectController) _import(ctx *gin.Context) {

}

func (c *ProjectController) getInteraction(ctx *gin.Context) {

}
