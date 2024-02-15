package http

import (
	"github.com/cza14h/nino-work/pkg/controller"
	"github.com/gin-gonic/gin"
)

type TemplateController struct {
	controller.BaseController
}

const template_prefix = "screenTemplate"

func (c *TemplateController) list(ctx *gin.Context) {
	
}
