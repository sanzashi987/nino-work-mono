package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/pkg/controller"
)

func CanvasUserLoggedIn(loginUrl string) func(*gin.Context) {
	sdkMiddleware := controller.ValidateMiddleware(loginUrl)
	return func(ctx *gin.Context) {
		jwtToken := ctx.GetHeader("auth")
		ctx.Request.Header.Add("Authentication", jwtToken)
		sdkMiddleware(ctx)
	}
}
