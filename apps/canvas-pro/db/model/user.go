package model

import "github.com/cza14h/nino-work/pkg/db"

type CanvasUserModel struct {
	db.BaseModel
	// ID         uint64           `gorm:"column:id;primaryKey;index:,unique"`
	Workspaces []WorkspaceModel `gorm:"many2many:canvas_workspace_user;foreignKey:Id;References:Id;joinForeignKey:CanvasUserId;joinReferences:WorkspaceId"`
	Permission uint64
}

func (u CanvasUserModel) TableName() string {
	return "canvas_users"
}
