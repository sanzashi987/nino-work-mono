package dao

import (
	"github.com/sanzashi987/nino-work/apps/user/db/model"
	"gorm.io/gorm"
)

func GetRolesFromPermissions(tx *gorm.DB, permissionIds []uint64) ([]*model.RoleModel, error) {

	permissions := []*model.PermissionModel{}

	for _, id := range permissionIds {
		permission := &model.PermissionModel{}
		permission.Id = id
		permissions = append(permissions, permission)
	}

	err := tx.Preload("Roles").Find(&permissions).Error

	roles := []*model.RoleModel{}

	for _, p := range permissions {
		roles = append(roles, p.Roles...)
	}

	return roles, err

}

func GetRolesByPermission(tx *gorm.DB, permissionId uint64) ([]*model.RoleModel, error) {
	roles := []*model.RoleModel{}
	err := tx.Table("role_permissions").
		Where("permission_model_id = ? ", permissionId).
		Joins("INNER JOIN roles ON role_permissions.role_model_id = roles.id").
		Scan(&roles).Error

	return roles, err
}
