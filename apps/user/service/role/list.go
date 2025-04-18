package roleService

import (
	"context"
	"errors"

	"github.com/sanzashi987/nino-work/apps/user/db/model"
	userService "github.com/sanzashi987/nino-work/apps/user/service/user"
	"github.com/sanzashi987/nino-work/pkg/db"
	"github.com/sanzashi987/nino-work/pkg/shared"
	"gorm.io/gorm"
)

var errNopermission = errors.New("user does not have any admin permission")
var errEmptyPermission = errors.New("no permission to bind")

type RoleMeta struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
	// Permissions []*shared.EnumMeta `json:"permissions"`
}

type ListRolesResponse = shared.ResponseWithPagination[[]*RoleMeta]

func makeListRoleQuery(ctx context.Context, userId uint64) (*gorm.DB, []uint64, error) {
	result, err := userService.GetUserAdmins(ctx, userId)
	if err != nil {
		return nil, nil, err
	}

	if !result.HasAnyAdmin() {
		return nil, nil, errNopermission
	}
	tx := result.Tx
	appIds := result.GetAllAppIds()

	permissionIds := []uint64{}

	if err := tx.Table("permissions").Where("app_id IN ?", appIds).Select("id").Scan(&permissionIds).Error; err != nil {
		return nil, nil, err
	}

	roleIds := []uint64{}

	if err := tx.
		Select("role_model_id").Table("role_permissions").Group("role_model_id").
		Having("SUM(permission_model_id IN ? ) > 0 AND SUM(permission_model_id NOT IN ?) = 0", permissionIds, permissionIds).
		Scan(&roleIds).Error; err != nil {
		return nil, nil, err
	}

	return tx, roleIds, nil
}

func ListRoles(ctx context.Context, userId uint64, payload *shared.PaginationRequest) (*ListRolesResponse, error) {

	tx, roleIds, err := makeListRoleQuery(ctx, userId)
	if err != nil {
		return nil, err
	}

	query := tx.Model(&model.RoleModel{}).Where("id IN ?", roleIds)

	var totalCount int64 = 0
	if err := query.Count(&totalCount).Error; err != nil {
		return nil, err
	}
	roles := []*model.RoleModel{}
	if err := query.Order("id DESC").Scopes(db.Paginate(payload.Page, payload.Size)).Find(&roles).Error; err != nil {
		return nil, err
	}

	roleMetas := make([]*RoleMeta, len(roles))
	for i, role := range roles {
		roleMetas[i] = &RoleMeta{
			Id:   role.Id,
			Name: role.Name,
			Code: role.Code,
		}
	}

	res := &ListRolesResponse{}
	res.Init(roleMetas, payload.Page, int(totalCount))
	return res, nil

}

func ListAllRoles(ctx context.Context, userId uint64) ([]*shared.EnumMeta, error) {
	tx, roleIds, err := makeListRoleQuery(ctx, userId)
	if err != nil {
		return nil, err
	}

	var roles []*model.RoleModel

	if err := tx.Where("id IN ?", roleIds).Find(&roles).Error; err != nil {
		return nil, err
	}

	roleEnums := make([]*shared.EnumMeta, len(roles))
	for i, role := range roles {
		roleEnums[i] = &shared.EnumMeta{
			Value: role.Id,
			Name:  role.Name,
		}
	}

	return roleEnums, nil
}
