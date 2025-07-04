package model

import "github.com/sanzashi987/nino-work/pkg/db"

type RoleModel struct {
	db.BaseModel
	Code        string             `gorm:"column:code;type:varchar(255);uniqueIndex"`
	Name        string             `gorm:"column:name"`
	Description string             `gorm:"column:description"`
	Permissions []*PermissionModel `gorm:"many2many:role_permissions;"`
	// Users       []*UserModel       `gorm:"many2many:user_roles;"`
	Users []*UserModel `gorm:"-"`
	Menus []*MenuModel `gorm:"many2many:menu_roles;"`
}

func (u RoleModel) TableName() string {
	return "roles"
}
