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
	Name        string
	Code        string            `gorm:"column:name;type:varchar(255);uniqueIndex"`
	Description string            `gorm:"column:description"`
	Status      uint              `gorm:"column:status"`
	Permissions []PermissionModel `gorm:"foreignKey:ServiceID"`
}

func (f SystemModel) TableName() string {
	return "services"
}
