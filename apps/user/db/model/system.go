package model

import (
	"github.com/sanzashi987/nino-work/pkg/db"
)

const (
	SystemOnline  = 0
	SystemOffline = 1
)

type SystemModel struct {
	db.BaseModel
	Name        string            `gorm:"column:name;type:varchar(255)"`
	Code        string            `gorm:"column:code;type:varchar(255);uniqueIndex"`
	Description string            `gorm:"column:description"`
	Status      uint              `gorm:"column:status"`
	CreateBy    uint64            `gorm:"column:create_by"`
	Permissions []PermissionModel `gorm:"foreignKey:ServiceID"`
}

func (f SystemModel) TableName() string {
	return "services"
}
