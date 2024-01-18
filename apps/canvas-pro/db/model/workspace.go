package model

import (
	"github.com/cza14h/nino-work/pkg/db"
)

type WorkSpaceModel struct {
	db.BaseModel
	Name       string `gorm:"column:name"`
	Code       string `gorm:"type:varchar(255);index;unique"`
	Default    int8
	Owner      uint64
	Capacity   uint64
	Permission uint64
	Deleted    int8
	// Members    []CanvasUserModel `gorm:"foreignKey:Workspace;references:Code"`
}

func (s WorkSpaceModel) TableName() string {
	return "workspaces"
}
