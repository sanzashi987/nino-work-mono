package dao

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/canvas-pro/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type WorkspaceDao struct {
	db.BaseDao[model.WorkspaceModel]
}

func NewWorkspaceDao(ctx context.Context, chain ...*db.BaseDao[model.WorkspaceModel]) *WorkspaceDao {
	var baseDao db.BaseDao[model.WorkspaceModel]

	if len(chain) > 0 {
		baseDao = *chain[0]
	} else {
		baseDao = db.InitBaseDao[model.WorkspaceModel](ctx)
	}

	return &WorkspaceDao{baseDao}
}

func (dao *WorkspaceDao) CreateWorkspace(canvasUser *model.CanvasUserModel) {
	userId := canvasUser.UnifiedUserId
	newWorkspace := model.WorkspaceModel{Owner: userId}
	newWorkspace.Creator = userId

}
