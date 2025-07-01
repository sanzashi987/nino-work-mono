package dao

import (
	"sync"

	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"gorm.io/gorm"
)

var userWorkspaceCache sync.Map // map[uint64]*model.CanvixUserModel

func GetUserWorkspaces(tx *gorm.DB, userId uint64) (*model.CanvixUserModel, error) {
	if cached, ok := userWorkspaceCache.Load(userId); ok {
		if userModel, ok := cached.(*model.CanvixUserModel); ok {
			return userModel, nil
		}
	}

	canvasUser := model.CanvixUserModel{UnifiedUserId: userId}
	if err := tx.Model(&canvasUser).Association("Workspaces").Find(&canvasUser.Workspaces); err != nil {
		return nil, err
	}
	userWorkspaceCache.Store(userId, &canvasUser)
	return &canvasUser, nil
}

func CreateUser(tx *gorm.DB, userId uint64) error {
	canvasUser := model.CanvixUserModel{UnifiedUserId: userId}
	return tx.Create(&canvasUser).Error
}
