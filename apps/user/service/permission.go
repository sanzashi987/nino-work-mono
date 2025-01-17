package service

import (
	"context"
	"errors"

	"github.com/sanzashi987/nino-work/apps/user/db/dao"
	"github.com/sanzashi987/nino-work/apps/user/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type PermissionServiceWeb struct{}

var PermissionServiceWebImpl *PermissionServiceWeb = &PermissionServiceWeb{}

type ListPermissionResult struct {
	*UserAdminResult
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
		UserAdminResult: result,
		AppList:         appList,
		App:             app,
		FromSuper:       fromSuper,
		FromAdmin:       fromAdmin,
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
