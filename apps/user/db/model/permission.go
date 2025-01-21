package model

import (
	"github.com/sanzashi987/nino-work/pkg/db"
)

type PermissionModel struct {
	db.BaseModel
	AppId       uint64       `gorm:"column:app_id;index"`
	Name        string       `gorm:"column:name;type:varchar(255)"`
	Code        string       `gorm:"column:code;type:varchar(255);uniqueIndex"`
	Description string       `gorm:"column:description"`
	Roles       []*RoleModel `gorm:"many2many:role_permissions"`
}

func (u PermissionModel) TableName() string {
	return "permissions"
}

func CreateRoleWithPermission(name, code string) (*RoleModel, *PermissionModel) {

	permission := PermissionModel{
		Name: name,
		Code: code,
	}
	role := RoleModel{
		Name:        name,
		Code:        code,
		Permissions: []*PermissionModel{&permission},
	}
	return &role, &permission
}
