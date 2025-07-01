package service

import (
	"context"
	"errors"

	"github.com/sanzashi987/nino-work/apps/canvix/db/dao"
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
	"gorm.io/gorm"
)

func ValidateUserWorkspace(ctx context.Context, userId uint64, workspaceCode string) bool {
	tx := db.NewTx(ctx)

	user, err := dao.GetUserWorkspaces(tx, userId)
	if err != nil {
		return false
	}

	for _, workspace := range user.Workspaces {
		if workspace.Code == workspaceCode {
			return true
		}
	}
	return false
}

func UserOnBoard(ctx context.Context, userId uint64) error {
	tx := db.NewTx(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 1. Create user if not exists
	userModel := &model.CanvixUserModel{UnifiedUserId: userId}
	err := tx.Where(userModel).FirstOrCreate(userModel).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 2. Create default workspace
	workspace := &model.WorkspaceModel{
		Owner:   userId,
		Default: 1,
		Members: []*model.CanvixUserModel{userModel},
	}
	if err := tx.Create(workspace).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 3. Update join table (user as member)
	if err := tx.Model(workspace).Association("Members").Append(userModel); err != nil {
		tx.Rollback()
		return errors.New("failed to associate user with workspace")
	}

	tx.Commit()
	return nil

}

// ValidateOrOnboardCanvixUser checks if a Canvix user exists; if not, it onboards the user.
func ValidateOrOnboardCanvixUser(ctx context.Context, userId uint64) error {
	tx := db.NewTx(ctx)
	var user model.CanvixUserModel
	err := tx.Where(&model.CanvixUserModel{UnifiedUserId: userId}).First(&user).Error
	if err == nil {
		return nil // User exists
	}
	if err == gorm.ErrRecordNotFound {
		// User does not exist, onboard
		return UserOnBoard(ctx, userId)
	}
	return err // Some other error
}
