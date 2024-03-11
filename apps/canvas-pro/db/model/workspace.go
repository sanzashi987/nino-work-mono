package model

import (
	"github.com/cza14h/nino-work/pkg/db"
)

type WorkspaceModel struct {
	db.BaseModel
	Name       string `gorm:"column:name"`
	Code       string `gorm:"type:varchar(255);index;unique"`
	Default    int8
	Owner      uint64
	Capacity   uint64
	Permission uint64
	Deleted    int8
	Members    []CanvasUserModel `gorm:"many2many:canvas_workspace_user;foreignKey:Id;joinForeginKey:WorkspaceId;joinReferences:CanvasUserId;References:Id"`
}

func (s WorkspaceModel) TableName() string {
	return "workspaces"
}
