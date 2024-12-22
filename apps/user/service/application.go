package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/sanzashi987/nino-work/apps/user/db/dao"
	"github.com/sanzashi987/nino-work/apps/user/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type ApplicationServiceWeb struct{}

var SystemServiceWebImpl *ApplicationServiceWeb = &ApplicationServiceWeb{}

// CreateAppRequest 创建系统请求参数
type CreateAppRequest struct {
	Operator    uint64
	Name        string
	Code        string
	Description string
}

func (u *ApplicationServiceWeb) CreateApplication(ctx context.Context, payload CreateAppRequest) error {
	if payload.Operator == 0 {
		return errors.New("user id is required")
	}

	if payload.Code == "" {
		return errors.New("system code is required")
	}

	appDao := dao.NewApplicationDao(ctx)
	appDao.BeginTransaction()

	application := &model.ApplicationModel{
		Name:        payload.Name,
		Code:        payload.Code,
		Description: payload.Description,
		Status:      model.SystemOnline,
		CreateBy:    payload.Operator,
	}

	if err := appDao.Create(application); err != nil {
		appDao.RollbackTransaction()
		return err
	}

	superAdminPermission := &model.PermissionModel{
		AppId:       application.Id,
		Name:        fmt.Sprintf("%s应用超级管理员权限", application.Name),
		Code:        fmt.Sprintf("%s.super_admin", application.Code),
		Description: fmt.Sprintf("%s应用超级管理员权限", application.Name),
	}

	adminPermission := &model.PermissionModel{
		AppId:       application.Id,
		Name:        fmt.Sprintf("%s应用管理员权限", application.Name),
		Code:        fmt.Sprintf("%s.admin", application.Code),
		Description: fmt.Sprintf("%s应用管理员权限", application.Name),
	}

	permissionDao := dao.NewPermissionDao(ctx, (*db.BaseDao[model.PermissionModel])(&appDao.BaseDao))

	if err := permissionDao.CreatePermissions(superAdminPermission, adminPermission); err != nil {
		appDao.RollbackTransaction()
		return err
	}

	if err := appDao.InitPermissionForSystem(application, superAdminPermission, adminPermission); err != nil {
		appDao.RollbackTransaction()
		return err
	}

	appDao.CommitTransaction()
	return nil
}

type AddPermissionRequest struct {
	SystemId   uint64                 `json:"system_id"`
	OperatorId uint64                 `json:"operator_id"`
	Permission *model.PermissionModel `json:"permission"`
}

func (u *ApplicationServiceWeb) AddPermission(ctx context.Context, payload AddPermissionRequest) (err error) {
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
