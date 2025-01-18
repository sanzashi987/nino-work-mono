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

type ListPermissionResult struct {
	*AppAdminResult
	AppList   []*model.ApplicationModel
	App       *model.ApplicationModel
	FromSuper bool
	FromAdmin bool
}

func (s *PermissionServiceWeb) ListPermissionByApp(ctx context.Context, userId uint64, appId *uint64) (*ListPermissionResult, error) {
	result, err := getUserAdmins(ctx, userId)
	if err != nil {
		return nil, err
	}
	var toQuery *uint64 = nil

	appList := removeRepeat(result)

	fromSuper, fromAdmin := false, false

	if len(result.SuperAdminApps) > 0 {
		if appId == nil {
			toQuery = &result.SuperAdminApps[0].Id
			fromSuper = true
		} else {
			for _, app := range result.SuperAdminApps {
				if app.Id == *appId {
					toQuery = appId
					fromSuper = true
					break
				}
			}

		}
	} else if len(result.AdminApps) > 0 {
		if appId == nil {
			toQuery = &result.AdminApps[0].Id
			fromAdmin = true
		} else {
			for _, app := range result.AdminApps {
				if app.Id == *appId {
					toQuery = appId
					fromAdmin = true
					break
				}
			}

		}
	}

	if toQuery == nil {
		return nil, nil
	}

	tx := db.NewTx(ctx)

	app, err := dao.FindApplicationByIdWithPermission(tx, *toQuery)
	if err != nil {
		return nil, err
	}

	listResult := ListPermissionResult{
		AppAdminResult: result,
		AppList:        appList,
		App:            app,
		FromSuper:      fromSuper,
		FromAdmin:      fromAdmin,
	}
	return &listResult, nil
}

type PermissionPayload struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type AddPermissionRequest struct {
	AppId       *uint64             `json:"app_id" binding:"required"`
	Permissions []PermissionPayload `json:"permissions"`
}

func (u *ApplicationServiceWeb) AddPermission(ctx context.Context, userId uint64, payload AddPermissionRequest) (err error) {
	app, tx, err := userIsManager(ctx, userId, payload.AppId, false)
	if err != nil {
		return
	}

	nextPermissionMap := map[string]bool{}
	for _, permission := range payload.Permissions {
		nextPermissionMap[permission.Code] = true
	}

	// 检查是否存在相同Code的权限
	for _, p := range app.Permissions {
		if _, ok := nextPermissionMap[p.Code]; ok {
			return errors.New("permission code already exists")
		}
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

	err = tx.Model(app).Association("Permissions").Append(permissionModels)
	return
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
	id   uint64
	name string
	code string
}

type PermissionsOfApp struct {
	Permissions []*PermissionRecord `json:"permissions"`
	Admin       uint64              `json:"admin"`
	SuperAdmin  uint64              `json:"super_admin"`
	*AdminResult
}

func (s *PermissionServiceWeb) ListPermissionsByApp(ctx context.Context, userId uint64, appId *uint64) (*PermissionsOfApp, error) {
	tx, adminResult, err := userIsAdmin(ctx, userId, appId)
	if err != nil {
		return nil, err
	}

	if !adminResult.IsAdmin && !adminResult.IsSuper {
		return nil, errors.New("user is not the admin of this app")

	}

	app := model.ApplicationModel{}

	if err := tx.Preload("Permissions", "app_id = ?", *appId).Where("id  = ?", *appId).Find(&app).Error; err != nil {
		return nil, err
	}

	res := &PermissionsOfApp{
		AdminResult: adminResult,
	}

	for _, p := range app.Permissions {
		res.Permissions = append(res.Permissions, &PermissionRecord{
			id:   p.Id,
			name: p.Name,
			code: p.Code,
		})
		res.Admin = app.Admin
		res.SuperAdmin = app.SuperAdmin
	}

	return res, nil
}
