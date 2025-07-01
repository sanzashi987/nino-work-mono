package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanzashi987/nino-work/apps/canvix/service"
	"github.com/sanzashi987/nino-work/pkg/controller"
)

// WorkspaceAccessRequired checks if the user has access to the workspace specified in the 'x-canvix-workspace' header.
func WorkspaceAccessRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		workspaceCode := ctx.GetHeader("x-canvix-workspace")
		if workspaceCode == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing x-canvix-workspace header"})
			return
		}

		userId := ctx.GetUint64(controller.UserID)
		if userId == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			return
		}

		if !service.ValidateUserWorkspace(ctx, userId, workspaceCode) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "No access to workspace"})
			return
		}

		ctx.Next()
	}
}
