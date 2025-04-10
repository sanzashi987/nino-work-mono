package dao

import (
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"gorm.io/gorm"
)

func GetUserWorkspaces(tx *gorm.DB, userId uint64) (*model.CanvasUserModel, error) {
	canvasUser := model.CanvasUserModel{UnifiedUserId: userId}
	if err := tx.Model(&canvasUser).Association("Workspaces").Find(&canvasUser.Workspaces); err != nil {
		return nil, err
	}
	return &canvasUser, nil
}

func CreateUser(tx *gorm.DB, userId uint64) error {
	canvasUser := model.CanvasUserModel{UnifiedUserId: userId}
	return tx.Create(&canvasUser).Error
}
