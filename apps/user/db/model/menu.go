package model

import "github.com/sanzashi987/nino-work/pkg/db"

const (
	MenuTypeMenu    MenuType = 1
	MenuTypeCatelog MenuType = 2
	MenuTypeButton  MenuType = 3
)

type MenuType uint8

type MenuModel struct {
	db.BaseModel
	Name        string            `gorm:"column:name;type:varchar(255)"`
	Code        string            `gorm:"column:code;type:varchar(255);uniqueIndex"`
	Description string            `gorm:"column:description"`
	Permissions []PermissionModel `gorm:"many2many:menu_permissions;"`
	Type        MenuType          `gorm:"column:type"`
	Order       int               `gorm:"column:order"`
	Status      int               `gorm:"column:status"`
	Hyperlink   bool              `gorm:"column:hyperlink"`
	Path        string            `gorm:"column:path"`
	Icon        string            `gorm:"column:icon"`
	// ParentID    uint64            `gorm:"column:parent_id"`
}

func (m MenuModel) TableName() string {
	return "menus"
}
