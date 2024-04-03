package service

import (
	"context"
	"errors"

	"github.com/cza14h/nino-work/apps/canvas-pro/consts"
	"github.com/cza14h/nino-work/apps/canvas-pro/db/dao"
	"github.com/cza14h/nino-work/apps/canvas-pro/db/model"
	"github.com/cza14h/nino-work/pkg/db"
)

type GroupService struct{}

var GroupServiceImpl *GroupService = &GroupService{}

var ErrorNameContainIllegalChar = errors.New("error name contains illegal character")
var ErrorNameExisted = errors.New("error group name is exist")

func (serv GroupService) Create(ctx context.Context, name, workspace, typeTag string) (err error) {
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
	workspaceId, _, _ := consts.GetIdFromCode(workspace)

	record := model.GroupModel{}
	record.Name, record.Workspace, record.TypeTag = name, workspaceId, typeTag
	return groupDao.Create(record)
}

var ErrorGroupNotFound = errors.New("error group is not exist")

func delete(groupDao *dao.GroupDao, id, workspace uint64, typeTag string) (err error) {

	record, err := groupDao.FindByKey("id", id)
	if record == nil || err != nil {
		err = ErrorGroupNotFound
		return
	}

	if record.DeleteTime != nil {
		return
	}
	groupDao.LogicalDelete(*record)
	return
}

func (serv GroupService) DeleteProjectGroup(ctx context.Context, groupCode, workspaceCode string) (err error) {

	groupDao := dao.NewGroupDao(ctx)
	groupDao.BeginTransaction()
	groupId, _, _ := consts.GetIdFromCode(groupCode)
	workspaceId, _, _ := consts.GetIdFromCode(workspaceCode)

	chainProjectDao := dao.NewProjectDao(ctx, (*db.BaseDao[model.ProjectModel])(&groupDao.BaseDao))

	if err = delete(groupDao, groupId, workspaceId, consts.PROJECT); err != nil {
		groupDao.RollbackTransaction()
		return
	}
	if err = chainProjectDao.DeleleGroupEffect(groupId, workspaceId); err != nil {
		groupDao.RollbackTransaction()
		return
	}

	groupDao.CommitTransaction()
	return
}

func (serv GroupService) Rename(ctx context.Context, userId uint64, workspaceCode, groupCode, groupName string) {
	UserServiceImpl.ValidateUserWorkspace(ctx, userId, workspaceCode)
}
