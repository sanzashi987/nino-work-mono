package dao

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/canvas-pro/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type UserDao struct {
	db.BaseDao[model.CanvasUserModel]
}

func NewUserDao(ctx context.Context, dao ...*db.BaseDao[model.CanvasUserModel]) *UserDao {
	return &UserDao{BaseDao: db.NewDao[model.CanvasUserModel](ctx, dao...)}
}

func (dao *UserDao) GetUserWorkspaces(userId uint64) (*model.CanvasUserModel, error) {
	canvasUser := model.CanvasUserModel{UnifiedUserId: userId}
	if err := dao.GetOrm().Model(&canvasUser).Association("Workspaces").Find(&canvasUser.Workspaces); err != nil {
		return nil, err
	}
	return &canvasUser, nil
}

func (dao *UserDao) CreateUser(userId uint64) error {
	canvasUser := model.CanvasUserModel{UnifiedUserId: userId}
	return dao.GetOrm().Create(&canvasUser).Error
}
