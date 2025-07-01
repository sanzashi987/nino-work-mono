package dao

import (
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"gorm.io/gorm"
)

func CreateWorkspace(canvasUser *model.CanvixUserModel) {
	userId := canvasUser.UnifiedUserId
	newWorkspace := model.WorkspaceModel{Owner: userId}
	newWorkspace.Creator = userId

}

// CreateAndPersistWorkspace creates a workspace and associates members.
func CreateAndPersistWorkspace(tx *gorm.DB, ownerId uint64, members []*model.CanvixUserModel) (*model.WorkspaceModel, error) {
	workspace := &model.WorkspaceModel{
		Owner:   ownerId,
		Default: 1,
		Members: members,
	}
	if err := tx.Create(workspace).Error; err != nil {
		return nil, err
	}
	if err := tx.Model(workspace).Association("Members").Append(members); err != nil {
		return nil, err
	}
	return workspace, nil
}
