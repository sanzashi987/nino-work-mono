package model

import (
	// "github.com/sanzashi987/nino-work/apps/canvix/utils"
	"github.com/sanzashi987/nino-work/pkg/db"
	// "gorm.io/gorm"
)

type CanvixUserModel struct {
	db.BaseModel
	UnifiedUserId uint64           `gorm:"column:unified_user_id;index;unique;"`
	Workspaces    []WorkspaceModel `gorm:"many2many:canvas_workspace_user;foreignKey:Id;References:Id;joinForeignKey:CanvasUserId;joinReferences:WorkspaceId"`
}

func (u CanvixUserModel) TableName() string {
	return "canvix_users"
}

// func (u *CanvasUserModel) BeforeCreate(tx *gorm.DB) (err error) {
// 	tempId := utils.GenerateId()
// 	u.Id = uint64(tempId)
// 	return
// }
