package dao

import (
	"context"

	"github.com/cza14h/nino-work/apps/canvas-pro/db/model"
	"github.com/cza14h/nino-work/pkg/db"
)

type UserDao struct {
	db.BaseDao[model.CanvasUserModel]
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{db.InitBaseDao[model.CanvasUserModel](ctx)}
}

func (dao *UserDao) GetUserWorkspaces(userId uint64) (*model.CanvasUserModel, error) {
	canvasUser := model.CanvasUserModel{}
	canvasUser.Id = userId
	if err := dao.GetOrm().Model(&canvasUser).Association("Workspaces").Find(&canvasUser.Workspaces); err != nil {
		return nil, err
	}
	return &canvasUser, nil
}
