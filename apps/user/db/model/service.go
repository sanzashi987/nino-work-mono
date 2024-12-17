package model

import (
	"github.com/sanzashi987/nino-work/pkg/db"
)

type ServiceModel struct {
	db.BaseModel
	Name        string            `gorm:"column:name;type:varchar(255);uniqueIndex"`
	Description string            `gorm:"column:description"`
	Status      uint              `gorm:"column:status"`
	Permissions []PermissionModel `gorm:"foreignKey:ServiceID"`
}

func (f ServiceModel) TableName() string {
	return "services"
}
