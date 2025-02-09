package service

import (
	"context"
	"errors"
	"fmt"

	userService "github.com/sanzashi987/nino-work/apps/user/service/user"

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

func (u *ApplicationServiceWeb) CreateApp(ctx context.Context, userId uint64, payload CreateAppRequest) (*model.ApplicationModel, error) {

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

	superAdminRole, superAdminPermission := model.CreateRoleWithPermission(
		fmt.Sprintf("%s应用超级管理员权限", application.Name),
		fmt.Sprintf("%s.admin.super", application.Code),
	)

	adminRole, adminPermission := model.CreateRoleWithPermission(
		fmt.Sprintf("%s应用管理员权限", application.Name),
		fmt.Sprintf("%s.admin.super", application.Code),
	)

	superAdminPermission.AppId, adminPermission.AppId = application.Id, application.Id

	rolesToCreate := []*model.RoleModel{superAdminRole, adminRole}
	if err := tx.Create(&rolesToCreate).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := dao.InitPermissionForApp(tx, application, superAdminPermission, adminPermission); err != nil {
		tx.Rollback()
		return nil, err
	}

	user := model.UserModel{}
	user.Id = userId
	if err := tx.Model(&user).Association("Roles").Append(&rolesToCreate); err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return application, nil
}

func removeRepeat(result *userService.AppAdminResult) []*model.ApplicationModel {
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
	result, err := userService.GetUserAdmins(ctx, userId)
	if err != nil {
		return nil, err
	}
	apps := removeRepeat(result)
	return apps, err
}

func userIsManager(ctx context.Context, userId uint64, appId *uint64, superOnly bool) (app *model.ApplicationModel, tx *gorm.DB, err error) {
	user, tx, err := userService.GetUserRolePermission(ctx, userId)
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

func (u *ApplicationServiceWeb) RemoveApp(ctx context.Context, userId uint64) {

}
