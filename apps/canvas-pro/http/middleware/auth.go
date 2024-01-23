package middleware

import (
	"github.com/cza14h/nino-work/pkg/auth"
	"github.com/gin-gonic/gin"
)

func CanvasMiddleware(loginUrl string) func(*gin.Context) {
	sdkMiddleware := auth.ValidateMiddleware(loginUrl)
	return func(ctx *gin.Context) {
		jwtToken := ctx.GetHeader("auth")
		ctx.Request.Header.Add("Authentication", jwtToken)
		sdkMiddleware(ctx)
	}
}
