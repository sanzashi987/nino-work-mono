package http

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/user/db/model"
	"github.com/sanzashi987/nino-work/pkg/controller"
	"github.com/sanzashi987/nino-work/pkg/db"
	"github.com/sanzashi987/nino-work/pkg/utils"
)

func NewRouter(loginPageUrl string) *gin.Engine {
	apiEngine := gin.Default()

	enhancedAuth := func(ctx *gin.Context) {
		claim, err := controller.ValidateFromCtx(ctx)
		if err == nil {
			ctx.Set(controller.UserID, claim.UserID)
			ctx.Set(controller.Username, claim.Username)
			ctx.Next()
		} else {
			method := ctx.Request.Method
			psm := ctx.Request.Header.Get("x-psm")
			signature := ctx.Request.Header.Get("x-signature")
			accessKey := ctx.Request.Header.Get("x-ak")
			contentType := ctx.Request.Header.Get("Content-Type")
			timestamp := ctx.Request.Header.Get("x-timestamp")
			path := ctx.Request.URL.Path

			if method != "" && psm != "" && signature != "" && accessKey != "" && contentType != "" && timestamp != "" && path != "" {
				unixTime, err := strconv.ParseInt(timestamp, 10, 64)
				if err != nil {
					ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid timestamp"})
					return
				}
				requestTime := time.Unix(unixTime, 0)
				if err != nil {
					ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid timestamp"})
					return
				}

				if time.Since(requestTime) > 5*time.Minute {
					ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Outdated request"})
					return
				}

				tx := db.NewTx(ctx)
				app := model.ApplicationModel{}

				if err := tx.Where("psm = ? AND access_key = ?", psm, accessKey).First(&app).Error; err == nil {
					secret := app.SecretKey

					sign := utils.GenerateSignature(psm, method, path, contentType, timestamp, secret)

					if sign != signature {
						ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Signature Validation Failed"})
						return
					}
					ctx.Set(controller.UserID, app.ServiceUser)
					ctx.Set(controller.Username, app.Code)
					ctx.Next()
				}

			} else {
				redirectURL := loginPageUrl + "?redirect=" + url.QueryEscape(ctx.Request.Referer())
				ctx.Redirect(http.StatusFound, redirectURL)
				return
			}

		}

	}

	v1 := apiEngine.Group("/backend/root/v1")
	authed := v1.Use(enhancedAuth)
	{
		v1.GET("misc/importmap", miscController.GetImportMap)

		RegisterUserRoutes(v1, authed)
		RegisterAppRoutes(authed)
		RegisterAppPermissionRoutes(authed)
		RegisterRoleRoutes(authed)

	}

	return apiEngine
}
