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

	canvixUser := model.CanvixUserModel{UnifiedUserId: userId}
	if err := tx.Model(&canvixUser).Association("Workspaces").Find(&canvixUser.Workspaces); err != nil {
		return nil, err
	}
	userWorkspaceCache.Store(userId, &canvixUser)
	return &canvixUser, nil
}

func CreateUser(tx *gorm.DB, userId uint64) error {
	canvixUser := model.CanvixUserModel{UnifiedUserId: userId}
	return tx.Create(&canvixUser).Error
}
