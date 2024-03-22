package service

import (
	"context"
	"errors"

	"github.com/cza14h/nino-work/apps/canvas-pro/consts"
	"github.com/cza14h/nino-work/apps/canvas-pro/db/dao"
	"github.com/cza14h/nino-work/apps/canvas-pro/db/model"
)

type GroupService struct{}

var groupService *GroupService

func init() {
	groupService = &GroupService{}
}

func GetGroupService() *GroupService {
	return groupService
}

var ErrorNameContainIllegalChar = errors.New("error name contains illegal character")
var ErrorNameExisted = errors.New("error group name is exist")

func create(ctx context.Context, name, workspace string, dbModel dao.DBModel) (err error) {
	if consts.LegalNameReg.FindStringIndex(name) == nil {
		err = ErrorNameContainIllegalChar
		return
	}

	groupDao := dao.NewGroupDao(ctx)

	records, err := groupDao.FindByNameAndWorkspace(name, workspace)
	if records != nil && err == nil {
		groupsInUse := model.FilterRecordsInUse(*records)
		if len(groupsInUse) > 0 {
			err = ErrorNameExisted
			return
		}
	}
	return groupDao.Create(name, workspace, dbModel)
}

func (serv *GroupService) CreateProjectGroup(ctx context.Context, name, workspace string) error {
	return create(ctx, name, workspace, model.ProjectGroupModel{})
}

var ErrorGroupNotFound = errors.New("error group is not exist")

func (serv *GroupService) delete(ctx context.Context, code, workspace string, dbModel dao.DBModel) (err error) {

	groupDao := dao.NewGroupDao(ctx)

	id, _, _ := consts.GetIdFromCode(code)

	record, err := groupDao.FindByKey("id", id, dbModel.TableName())
	if record == nil || err != nil {
		err = ErrorGroupNotFound
		return
	}

	if record.Deleted == model.Deleted {
		return
	}

	record.Deleted = model.Deleted
	groupDao.UpdateById(*record, dbModel.TableName())
	return

}
