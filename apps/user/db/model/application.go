package model

import (
	"github.com/sanzashi987/nino-work/pkg/db"
)

const (
	SystemOnline  = 0
	SystemOffline = 1
)

type ApplicationModel struct {
	db.BaseModel
	Name        string `gorm:"column:name;type:varchar(255)"`
	Code        string `gorm:"column:code;type:varchar(255);uniqueIndex"`
	Description string `gorm:"column:description"`
	Status      uint   `gorm:"column:status"`
	CreateBy    uint64 `gorm:"column:create_by"`
	// store permission Id here
	SuperAdmin uint64 `gorm:"column:super_admin;"`
	// store permission Id here
	Admin       uint64            `gorm:"column:admin;"`
	Permissions []*PermissionModel `gorm:"foreignKey:AppId"`
}

func (f ApplicationModel) TableName() string {
	return "applications"
}
