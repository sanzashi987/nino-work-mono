package service

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/canvix/db/dao"
)

type UserService struct{}

var UserServiceImpl *UserService = &UserService{}

func (serv UserService) ValidateUserWorkspace(ctx context.Context, userId uint64, workspaceCode string) bool {
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

func (serv UserService) UserOnBoard(ctx context.Context, userId uint64) {
	userDao := dao.NewUserDao(ctx)
	userDao.BeginTransaction()
	
}
