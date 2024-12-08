package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/pkg/controller"
)

type TemplateController struct {
	controller.BaseController
}

const template_prefix = "screenTemplate"

func (c *TemplateController) list(ctx *gin.Context) {
	
}
