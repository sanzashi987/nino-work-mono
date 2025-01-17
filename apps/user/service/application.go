package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/sanzashi987/nino-work/apps/user/db/dao"
	"github.com/sanzashi987/nino-work/apps/user/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
	"gorm.io/gorm"
)

type ApplicationServiceWeb struct{}

var AppServiceWebImpl *ApplicationServiceWeb = &ApplicationServiceWeb{}

// CreateAppRequest 创建系统请求参数
type CreateAppRequest struct {
	Name        string `json:"name" binding:"required"`
	Code        string `json:"code" binding:"required"`
	Description string `json:"description"`
}

func (u *ApplicationServiceWeb) CreateApplication(ctx context.Context, userId uint64, payload CreateAppRequest) (*model.ApplicationModel, error) {

	tx := db.NewTx(ctx)
	tx = tx.Begin()

	application := &model.ApplicationModel{
		Name:        payload.Name,
		Code:        payload.Code,
		Description: payload.Description,
		Status:      model.SystemOnline,
		CreateBy:    userId,
	}

	if err := dao.CreateApp(tx, application); err != nil {
		tx.Rollback()
		return nil, err
	}

	superAdminPermission := &model.PermissionModel{
		AppId:       application.Id,
		Name:        fmt.Sprintf("%s应用超级管理员权限", application.Name),
		Code:        fmt.Sprintf("%s.admin.super", application.Code),
		Description: fmt.Sprintf("%s应用超级管理员权限", application.Name),
	}

	adminPermission := &model.PermissionModel{
		AppId:       application.Id,
		Name:        fmt.Sprintf("%s应用管理员权限", application.Name),
		Code:        fmt.Sprintf("%s.admin", application.Code),
		Description: fmt.Sprintf("%s应用管理员权限", application.Name),
	}

	if err := tx.Create([]*model.PermissionModel{superAdminPermission, adminPermission}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := dao.InitPermissionForApp(tx, application, superAdminPermission, adminPermission); err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return application, nil
}

func removeRepeat(result *UserAdminResult) []*model.ApplicationModel {
	apps := []*model.ApplicationModel{}
	appMap := map[uint64]*model.ApplicationModel{}

	for _, app := range result.SuperAdminApps {
		appMap[app.Id] = app
	}

	for _, app := range result.AdminApps {
		appMap[app.Id] = app
	}

	for _, app := range appMap {
		apps = append(apps, app)

	}

	return apps
}

func (u *ApplicationServiceWeb) ListApplications(ctx context.Context, userId uint64) ([]*model.ApplicationModel, error) {
	result, err := getUserAdmins(ctx, userId)
	if err != nil {
		return nil, err
	}
	apps := removeRepeat(result)
	return apps, err
}

func userIsManager(ctx context.Context, userId uint64, appId *uint64, superOnly bool) (app *model.ApplicationModel, tx *gorm.DB, err error) {
	user, tx, err := getUserRolePermission(ctx, userId)
	if err != nil {
		return
	}

	app, err = dao.FindApplicationByIdWithPermission(tx, *appId)
	if err != nil {
		return nil, nil, err
	}

	testers := map[uint64]bool{}
	testers[app.SuperAdmin] = true
	if !superOnly {
		testers[app.Admin] = true
	}

	var hasPermission = false
topLoop:
	for _, role := range user.Roles {
		for _, permission := range role.Permissions {

			if permission.AppId == *appId {
				if _, exsit := testers[permission.Id]; exsit {
					hasPermission = true
					break topLoop
				}
			}

		}
	}

	if !hasPermission {
		return nil, nil, errors.New("no permission")
	}

	return app, tx, nil
}
