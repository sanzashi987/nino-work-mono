package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	userService "github.com/sanzashi987/nino-work/apps/user/service/user"
	"github.com/sanzashi987/nino-work/pkg/controller"
)

type MiscController struct {
	controller.BaseController
}

var miscController = MiscController{}

var defaultMap = map[string]map[string]string{
	"imports": {
		"react":            "https://cdn.jsdelivr.net/npm/react@18.3.1/umd/react.production.min.js",
		"react-dom":        "https://cdn.jsdelivr.net/npm/react-dom@18.3.1/umd/react-dom.production.min.js",
		"react-dom/client": "https://cdn.jsdelivr.net/npm/react-dom@18.3.1/umd/react-dom.production.min.js",
		"single-spa":       "https://cdn.jsdelivr.net/npm/single-spa@6.0.3/lib/es2015/system/single-spa.min.js",
	},
}

func (c *MiscController) GetImportMap(ctx *gin.Context) {
	data := defaultMap

	authed, err := controller.ValidateFromCtx(ctx)
	if err == nil {
		id := authed.UserID
		userInfo, e := userService.UserServiceWebImpl.GetUserInfo(ctx, id)
		if e == nil {
			menus := userInfo.Menus
			menuCodes := make([]string, len(menus))
			for i, menu := range menus {
				menuCodes[i] = menu.Code
			}
		}

	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		c.AbortServerError(ctx, "Failed to create JSON")
		return
	}

	ctx.Writer.WriteHeader(http.StatusOK)
	// Set the appropriate headers for the file response
	ctx.Header("Content-Disposition", "attachment; filename=importmap.json")
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Content-Length", fmt.Sprintf("%d", len(jsonData)))

	// Write the JSON content to the response body
	ctx.Writer.Write(jsonData)
}
