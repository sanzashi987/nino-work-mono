package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/sanzashi987/nino-work/apps/user/db/dao"
	"github.com/sanzashi987/nino-work/apps/user/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type SystemServiceWeb struct{}

var SystemServiceWebImpl *SystemServiceWeb = &SystemServiceWeb{}

// CreateSystemRequest 创建系统请求参数
type CreateSystemRequest struct {
	Operator          uint64
	SystemName        string
	SystemCode        string
	SystemDescription string
}

func (u *SystemServiceWeb) CreateSystem(ctx context.Context, payload CreateSystemRequest) error {
	if payload.Operator == 0 {
		return errors.New("user id is required")
	}

	if payload.SystemCode == "" {
		return errors.New("system code is required")
	}

	systemDao := dao.NewSystemDao(ctx)
	systemDao.BeginTransaction()

	// 创建系统
	newSystem := &model.SystemModel{
		Name:        payload.SystemName,
		Code:        payload.SystemCode,
		Description: payload.SystemDescription,
		Status:      model.SystemOnline,
	}

	if err := systemDao.Create(newSystem); err != nil {
		systemDao.RollbackTransaction()
		return err
	}

	// 创建管理员权限
	adminPermission := &model.PermissionModel{
		ServiceID:   newSystem.Id,
		Name:        fmt.Sprintf("%s系统管理员权限", newSystem.Name),
		Code:        fmt.Sprintf("%s.admin", newSystem.Code),
		Description: fmt.Sprintf("%s系统管理员权限", newSystem.Name),
	}

	permissionDao := dao.NewPermissionDao(ctx, (*db.BaseDao[model.PermissionModel])(&systemDao.BaseDao))

	if err := permissionDao.CreatePermission(adminPermission); err != nil {
		systemDao.RollbackTransaction()
		return err
	}

	systemDao.CommitTransaction()
	return nil
}
