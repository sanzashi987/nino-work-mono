package model

import (
	"github.com/sanzashi987/nino-work/pkg/db"
)

type PermissionModel struct {
	db.BaseModel
	ServiceID   uint64        `gorm:"column:service_id;index"`
	Name        string      `gorm:"column:name;type:varchar(255)"`
	Code        string      `gorm:"column:code;type:varchar(255);uniqueIndex"`
	Description string      `gorm:"column:description"`
	Roles       []RoleModel `gorm:"many2many:role_permissions;"`
}

func (u PermissionModel) TableName() string {
	return "permissions"
}
