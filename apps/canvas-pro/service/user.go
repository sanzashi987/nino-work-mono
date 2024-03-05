package service

import (
	"context"

	"github.com/cza14h/nino-work/apps/canvas-pro/db/dao"
)

func ValidateUserWorkspace(ctx context.Context, userId uint64, workspaceCode string) bool {
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserWorkspaces(userId)
	if err != nil {
		return false
	}

	for _, workspace := range user.Workspaces {
		if workspace.Code == workspaceCode {
			return true
		}
	}
	return false
}
