package dao

import (
	"context"

	"github.com/cza14h/nino-work/apps/canvas-pro/db/model"
	"github.com/cza14h/nino-work/pkg/db"
)

type GroupDao struct {
	db.BaseDao[model.BaseModel]
}

func NewGroupDao(ctx context.Context) *GroupDao {
	return &GroupDao{db.InitBaseDao[model.BaseModel](ctx)}
}

func (dao *GroupDao) FindByNameAndWorkspace(name, workspace string) (res *[]model.BaseModel, err error) {
	err = dao.DB.Where("name = ? AND workspace = ?", name, workspace).Find(res).Error
	return
}

// func (dao *GroupDao) UpdateById(id uint64, dbModel DBModel) (res *model.BaseModel, err error) {
// 	err = dao.DB.Table(dbModel.TableName()).Where("id = ?", id).First(res).Error
// 	return
// }

func (dao *GroupDao) Delete(id uint64, table string) (err error) {
	toDelete := model.BaseModel{Deleted: model.Deleted}
	toDelete.Id = id

	if err = dao.UpdateById(toDelete, table); err != nil {
		return
	}


	
	// return dao.DB.Table(table).Create(&newGroup).Error
}
