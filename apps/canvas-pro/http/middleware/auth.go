package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/pkg/auth"
)

func CanvasUserLoggedIn(loginUrl string) func(*gin.Context) {
	sdkMiddleware := auth.ValidateMiddleware(loginUrl)
	return func(ctx *gin.Context) {
		jwtToken := ctx.GetHeader("auth")
		ctx.Request.Header.Add("Authentication", jwtToken)
		sdkMiddleware(ctx)
	}
}
