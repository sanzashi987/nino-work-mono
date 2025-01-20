package service

import (
	"context"
	"errors"

	"github.com/sanzashi987/nino-work/apps/user/db/dao"
	"github.com/sanzashi987/nino-work/apps/user/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
	"gorm.io/gorm"
)

type PermissionServiceWeb struct{}

var PermissionServiceWebImpl *PermissionServiceWeb = &PermissionServiceWeb{}

type PermissionPayload struct {
	Name        string `json:"name" binding:"required"`
	Code        string `json:"code" binding:"required"`
	Description string `json:"description"`
}

type CreatePermissionRequest struct {
	AppId       *uint64              `json:"app_id" binding:"required"`
	Permissions []*PermissionPayload `json:"permissions" binding:"required"`
}

func (u *PermissionServiceWeb) CreatePermission(ctx context.Context, userId uint64, payload CreatePermissionRequest) error {
	tx, result, err := userIsAdmin(ctx, userId, payload.AppId)
	if err != nil || !result.Admin() {
		return errors.New("user is not the admin of this app")
	}

	var codes []string
	for _, permission := range payload.Permissions {
		codes = append(codes, permission.Code)
	}
	var count int64
	if err := tx.Model(&model.PermissionModel{}).Where("app_id = ? AND code IN ?", *payload.AppId, codes).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("one or more permission codes already exist")
	}

	permissionModels := []*model.PermissionModel{}
	for _, permission := range payload.Permissions {
		permissionModels = append(permissionModels, &model.PermissionModel{
			AppId:       *payload.AppId,
			Name:        permission.Name,
			Code:        permission.Code,
			Description: permission.Description,
		})
	}

	app := model.ApplicationModel{}
	app.Id = *payload.AppId

	return tx.Model(app).Association("Permissions").Append(permissionModels)
}

type RemovePermissionRequest struct {
	AppId       *uint64  `json:"app_id" binding:"required"`
	Permissions []uint64 `json:"permissions"`
}

func (u *ApplicationServiceWeb) RemovePermission(ctx context.Context, userId uint64, payload RemovePermissionRequest) error {
	app, tx, err := userIsManager(ctx, userId, payload.AppId, false)
	if err != nil {
		return err
	}

	permissions := []*model.PermissionModel{}
	for _, id := range payload.Permissions {
		p := model.PermissionModel{}
		p.Id = id
		permissions = append(permissions, &p)
	}

	return tx.Model(app).Association("Permissions").Delete(permissions)

}

type AdminResult struct {
	IsAdmin bool `json:"is_admin"`
	IsSuper bool `json:"is_super"`
}

func (a AdminResult) Admin() bool {
	return a.IsAdmin || a.IsSuper
}

func userIsAdmin(ctx context.Context, userId uint64, appId *uint64) (*gorm.DB, *AdminResult, error) {
	tx := db.NewTx(ctx)
	user, err := dao.FindUserWithRoles(tx, userId)
	if err != nil {
		return nil, nil, err
	}

	userRolesMap := map[uint64]*model.RoleModel{}
	for _, role := range user.Roles {
		userRolesMap[role.Id] = role
	}

	app := model.ApplicationModel{}
	if err := tx.Where("id = ?", *appId).Find(&app).Error; err != nil {
		return nil, nil, err
	}

	result := AdminResult{}

	superRoles, err := dao.GetRolesByPermission(tx, app.SuperAdmin)
	if err != nil {
		return nil, nil, err
	}
	roles, err := dao.GetRolesByPermission(tx, app.Admin)
	if err != nil {
		return nil, nil, err
	}

	for _, role := range superRoles {
		if _, exist := userRolesMap[role.Id]; exist {
			result.IsSuper = true
			break
		}
	}
	for _, role := range roles {
		if _, exist := userRolesMap[role.Id]; exist {
			result.IsAdmin = true
			break
		}
	}

	return tx, &result, nil
}

type PermissionRecord struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type PermissionsResult struct {
	Permissions []*PermissionRecord `json:"permissions"`
	Admin       uint64              `json:"admin_id"`
	SuperAdmin  uint64              `json:"super_admin_id"`
	AppName     string              `json:"app_name"`
	*AdminResult
}

func (s *PermissionServiceWeb) ListPermissionsByApp(ctx context.Context, userId uint64, appId *uint64) (*PermissionsResult, error) {
	tx, adminResult, err := userIsAdmin(ctx, userId, appId)
	if err != nil {
		return nil, err
	}

	if !adminResult.Admin() {
		return nil, errors.New("user is not the admin of this app")

	}

	app := model.ApplicationModel{}

	if err := tx.Preload("Permissions", "app_id = ?", *appId).Where("id  = ?", *appId).Find(&app).Error; err != nil {
		return nil, err
	}

	res := &PermissionsResult{
		AdminResult: adminResult,
		Admin:       app.Admin,
		SuperAdmin:  app.SuperAdmin,
		AppName:     app.Name,
	}

	for _, p := range app.Permissions {
		res.Permissions = append(res.Permissions, &PermissionRecord{
			Id:   p.Id,
			Name: p.Name,
			Code: p.Code,
		})
	}

	return res, nil
}
