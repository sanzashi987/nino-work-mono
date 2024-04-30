package dao

import (
	"context"

	"github.com/cza14h/nino-work/apps/canvas-pro/db/model"
	"github.com/cza14h/nino-work/pkg/db"
)

type ProjectDao struct {
	db.BaseDao[model.ProjectModel]
}

func NewProjectDao(ctx context.Context, dao ...*db.BaseDao[model.ProjectModel]) *ProjectDao {
	var baseDao db.BaseDao[model.ProjectModel]
	if len(dao) > 0 {
		baseDao = *dao[0]
	} else {
		baseDao = db.InitBaseDao[model.ProjectModel](ctx)
	}
	return &ProjectDao{BaseDao: baseDao}
}

func (dao *ProjectDao) GetList(page, size int, workspace uint64, name *string, groupId *uint64) (projects *[]model.ProjectModel, err error) {

	query := dao.GetOrm().Scopes(db.Paginate(page, size)).Model(&model.ProjectModel{}).Where("workspace = ?", workspace)

	if groupId != nil {
		query = query.Where(" group_id = ?", *groupId)
	}

	if name != nil {
		query = query.Where(" name LIKE ?", *name)
	}
	err = query.Find(projects).Error
	return
}

var projectTableName = model.ProjectModel{}.TableName()

func (dao *ProjectDao) BatchLogicalDelete(ids []uint64) error {
	return dao.GetOrm().Table(projectTableName).Where("id IN ?", ids).Delete(&model.ProjectModel{}).Error

}

func (dao *ProjectDao) DeleleGroupEffect(groupId, workspace uint64) error {
	return dao.GetOrm().Table(projectTableName).Where("group_id = ? AND workspace = ?", groupId, workspace).Updates(map[string]any{"group_id": 0}).Error
}

func (dao *ProjectDao) BatchMoveGroup(groupId, workspace uint64, projectIds []uint64) error {

	orm := dao.GetOrm().Table(projectTableName)

	return orm.Where("id IN ? AND workspace = ?", projectIds, workspace).Update("group_id", groupId).Error

}

func (dao ProjectDao) GetProjectCountByGroup(workspaceId uint64, groupIds []uint64) (res []GroupCount, err error) {

	err = dao.GetOrm().Table(projectTableName).Where("workspace = ?", workspaceId).Where("group_id IN ?", groupIds).Select("id", "COUNT(id) as count").Group("group_id").Find(&res).Error
	return

}
