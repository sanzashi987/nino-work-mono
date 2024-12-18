package service

import (
	"context"
	"errors"

	"github.com/sanzashi987/nino-work/apps/user/db/dao"
	"github.com/sanzashi987/nino-work/apps/user/db/model"
)

type SystemServiceWeb struct{}

var SystemServiceWebImpl *SystemServiceWeb = &SystemServiceWeb{}

func (u *SystemServiceWeb) CreateSystem(ctx context.Context, operator uint16, systemName, systemDescription string) error {
	if operator == 0 {
		return errors.New("user id is required")
	}

	systemDao := dao.NewSystemDao(ctx)
	systemDao.BeginTransaction()

	// 创建系统
	newSystem := &model.SystemModel{
		Name:        systemName,
		Description: systemDescription,
		Status:      model.SystemOnline,
	}

	if err := systemDao.Create(newSystem); err != nil {
		systemDao.RollbackTransaction()
		return err
	}

	// 创建管理员权限
	adminPermission := &model.PermissionModel{
		ServiceID:   uint(newSystem.Id),
		Name:        "系统管理员",
		Code:        "system:admin",
		Description: "系统管理员权限",
	}

	if err := systemDao.CreatePermission(adminPermission); err != nil {
		systemDao.RollbackTransaction()
		return err
	}

	systemDao.CommitTransaction()
	return nil
}
