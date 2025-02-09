package userService

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/user/db/dao"
	"github.com/sanzashi987/nino-work/apps/user/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
	"gorm.io/gorm"
)

func GetUserRolePermission(ctx context.Context, userId uint64) (*model.UserModel, *gorm.DB, error) {
	tx := db.NewTx(ctx)
	user, err := dao.FindUserWithRoles(tx, userId)
	if err != nil {
		return nil, nil, err
	}

	userRoles := []*model.RoleModel{}
	userRoles = append(userRoles, user.Roles...)

	err = dao.FindRolesWithPermissions(tx, userRoles...)

	if err != nil {
		return nil, nil, err
	}

	user.Roles = userRoles

	return user, tx, nil
}

type AppAdminResult struct {
	SuperAdminApps []*model.ApplicationModel
	AdminApps      []*model.ApplicationModel
	Tx             *gorm.DB
}

func (a *AppAdminResult) HasAnyAdmin() bool {
	return (len(a.SuperAdminApps) + len(a.AdminApps)) > 0
}
func (a *AppAdminResult) GetAllAppIds() []uint64 {
	admins := a.AdminApps
	superAdmins := a.SuperAdminApps

	appIdSet := make(map[uint64]bool)

	for _, app := range admins {
		appIdSet[app.Id] = true
	}
	for _, app := range superAdmins {
		appIdSet[app.Id] = true
	}

	appIds := []uint64{}
	for key := range appIdSet {
		appIds = append(appIds, key)
	}
	return appIds
}

func GetUserAdmins(ctx context.Context, userId uint64) (*AppAdminResult, error) {
	user, tx, err := GetUserRolePermission(ctx, userId)
	if err != nil {
		return nil, err
	}

	applications := map[uint64]bool{}
	permissions := map[uint64]bool{}
	for _, role := range user.Roles {
		for _, permission := range role.Permissions {
			applications[permission.AppId] = true
			permissions[permission.Id] = true
		}
	}

	appIds := []uint64{}
	for appId := range applications {
		appIds = append(appIds, appId)
	}
	apps := []model.ApplicationModel{}
	err = tx.Table("applications").Where("id IN ?", appIds).Find(&apps).Error
	if err != nil {
		return nil, err

	}

	superRes := []*model.ApplicationModel{}
	adminRes := []*model.ApplicationModel{}
	superResMap := map[uint64]*model.ApplicationModel{}
	adminResMap := map[uint64]*model.ApplicationModel{}
	for _, app := range apps {
		if _, superExist := permissions[app.SuperAdmin]; superExist {
			superResMap[app.Id] = &app
		}
	}

	for _, app := range apps {
		if _, adminExist := permissions[app.Admin]; adminExist {
			adminResMap[app.Id] = &app
		}
	}

	for _, app := range superResMap {
		superRes = append(superRes, app)
	}
	for _, app := range adminResMap {
		adminRes = append(adminRes, app)
	}
	result := AppAdminResult{
		SuperAdminApps: superRes,
		AdminApps:      adminRes,
		Tx:             tx,
	}

	return &result, nil
}

type AdminResult struct {
	IsAdmin bool `json:"is_admin"`
	IsSuper bool `json:"is_super"`
}

func (a AdminResult) Admin() bool {
	return a.IsAdmin || a.IsSuper
}

type AdminData struct {
	Result *AdminResult
	App    *model.ApplicationModel
	Tx     *gorm.DB
}

func UserIsAdmin(ctx context.Context, userId uint64, appId *uint64) (*AdminData, error) {
	tx := db.NewTx(ctx)
	user, err := dao.FindUserWithRoles(tx, userId)
	if err != nil {
		return nil, err
	}

	userRolesMap := map[uint64]*model.RoleModel{}
	for _, role := range user.Roles {
		userRolesMap[role.Id] = role
	}

	app := model.ApplicationModel{}
	if err := tx.Where("id = ?", *appId).Find(&app).Error; err != nil {
		return nil, err
	}

	result := AdminResult{}

	superRoles, err := dao.GetRolesByPermission(tx, app.SuperAdmin)
	if err != nil {
		return nil, err
	}
	roles, err := dao.GetRolesByPermission(tx, app.Admin)
	if err != nil {
		return nil, err
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

	data := AdminData{
		Tx:     tx,
		Result: &result,
		App:    &app,
	}

	return &data, nil
}
