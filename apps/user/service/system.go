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
		CreateBy:    payload.Operator,
	}

	if err := systemDao.Create(newSystem); err != nil {
		systemDao.RollbackTransaction()
		return err
	}

	// 创建管理员权限
	superAdminPermission := &model.PermissionModel{
		ServiceID:   newSystem.Id,
		Name:        fmt.Sprintf("%s系统超级管理员权限", newSystem.Name),
		Code:        fmt.Sprintf("%s.super_admin", newSystem.Code),
		SuperAdmin:  true,
		Authorize:   true,
		Description: fmt.Sprintf("%s系统超级管理员权限", newSystem.Name),
	}

	adminPermission := &model.PermissionModel{
		ServiceID:   newSystem.Id,
		Name:        fmt.Sprintf("%s系统管理员权限", newSystem.Name),
		Code:        fmt.Sprintf("%s.admin", newSystem.Code),
		SuperAdmin:  false,
		Authorize:   true,
		Description: fmt.Sprintf("%s系统管理员权限", newSystem.Name),
	}

	permissionDao := dao.NewPermissionDao(ctx, (*db.BaseDao[model.PermissionModel])(&systemDao.BaseDao))

	if err := permissionDao.CreatePermissions(superAdminPermission, adminPermission); err != nil {
		systemDao.RollbackTransaction()
		return err
	}

	systemDao.CommitTransaction()
	return nil
}

type AddPermissionRequest struct {
	SystemId   uint64                 `json:"system_id"`
	OperatorId uint64                 `json:"operator_id"`
	Permission *model.PermissionModel `json:"permission"`
}

func (u *SystemServiceWeb) AddPermission(ctx context.Context, payload AddPermissionRequest) (err error) {
	if userId == 0 {
		return errors.New("用户ID不能为空")
	}
	// 检查用户是否有编辑权限
	userPermissions, err := UserServiceWebImpl.GetUserRoleWithPermissions(ctx, operatorId)
	if err != nil {
		return err
	}

	hasEditPermission := false
	for _, permission := range userPermissions {
		if permission.Authorize {
			hasEditPermission = true
			break
		}
	}

	if !hasEditPermission {
		return errors.New("用户无编辑权限")
	}
	_, err = UserServiceWebImpl.GetUserRoleWithPermissions(ctx, userId)

	return err
}
