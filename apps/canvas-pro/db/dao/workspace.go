package dao

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/canvas-pro/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type WorkspaceDao struct {
	db.BaseDao[model.WorkspaceModel]
}

func NewWorkspaceDao(ctx context.Context, dao ...*db.BaseDao[model.WorkspaceModel]) *WorkspaceDao {
	return &WorkspaceDao{BaseDao: db.NewDao[model.WorkspaceModel](ctx, dao...)}
}

func (dao *WorkspaceDao) CreateWorkspace(canvasUser *model.CanvasUserModel) {
	userId := canvasUser.UnifiedUserId
	newWorkspace := model.WorkspaceModel{Owner: userId}
	newWorkspace.Creator = userId

}
