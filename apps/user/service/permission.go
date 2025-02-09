package service

import (
	"context"
	"errors"

	"github.com/sanzashi987/nino-work/apps/user/db/model"
	userService "github.com/sanzashi987/nino-work/apps/user/service/user"
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
	data, err := userService.UserIsAdmin(ctx, userId, payload.AppId)
	if err != nil || !data.Result.Admin() {
		return errors.New("user is not the admin of this app")
	}

	var codes []string
	for _, permission := range payload.Permissions {
		codes = append(codes, permission.Code)
	}
	var count int64
	if err := data.Tx.Model(&model.PermissionModel{}).Where("app_id = ? AND code IN ?", *payload.AppId, codes).Count(&count).Error; err != nil {
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

	return data.Tx.Model(app).Association("Permissions").Append(permissionModels)
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
	*userService.AdminResult
}

func (s *PermissionServiceWeb) ListPermissionsByApp(ctx context.Context, userId uint64, appId *uint64) (*PermissionsResult, error) {
	data, err := userService.UserIsAdmin(ctx, userId, appId)
	if err != nil {
		return nil, err
	}

	if !data.Result.Admin() {
		return nil, errors.New("user is not the admin of this app")

	}

	if err := data.Tx.Preload("Permissions", "app_id = ?", *appId).Where("id  = ?", *appId).Find(data.App).Error; err != nil {
		return nil, err
	}

	res := &PermissionsResult{
		AdminResult: data.Result,
		Admin:       data.App.Admin,
		SuperAdmin:  data.App.SuperAdmin,
		AppName:     data.App.Name,
	}

	for _, p := range data.App.Permissions {
		res.Permissions = append(res.Permissions, &PermissionRecord{
			Id:   p.Id,
			Name: p.Name,
			Code: p.Code,
		})
	}

	return res, nil
}

type DeletePermissionRequest struct {
	Id    uint64 `json:"id"`
	AppId uint64 `json:"app_id"`
}

func (s *PermissionServiceWeb) DeletePermission(ctx context.Context, userId uint64, payload DeletePermissionRequest) error {
	data, err := userService.UserIsAdmin(ctx, userId, &payload.AppId)
	if err != nil {
		return err
	}

	if !data.Result.Admin() {
		return errors.New("user is not the admin of this app")
	}

	app := data.App

	if app.SuperAdmin == payload.Id || app.Admin == payload.Id {
		return errors.New("cannot delete a admin permission, replace it with another one before delete it")
	}

	if err := data.Tx.Delete(&model.PermissionModel{}, "id = ? ", payload.Id).Error; err != nil {
		return err
	}

	return nil
}

type SearchPermissionRequest struct {
	AppId *uint64 `json:"app_id" binding:"required"`
	Code  string  `json:"code" binding:"required"`
}

func (s *PermissionServiceWeb) SearchPermission(ctx context.Context, userId uint64, payload SearchPermissionRequest) (*PermissionRecord, error) {
	data, err := userService.UserIsAdmin(ctx, userId, payload.AppId)

	if err != nil {
		return nil, err
	}

	if !data.Result.Admin() {
		return nil, errors.New("user is not the admin of this app")
	}

	permission := model.PermissionModel{}
	if err := data.Tx.Where("app_id = ? AND code = ?", *payload.AppId, payload.Code).First(&permission).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("permission not found")
		}
		return nil, err
	}

	return &PermissionRecord{
		Id:   permission.Id,
		Name: permission.Name,
		Code: permission.Code,
	}, nil
}
