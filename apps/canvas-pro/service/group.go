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

var ErrorNameExisted = errors.New("error group name is exist")

func (serv GroupService) Create(ctx context.Context, name, workspace, typeTag string) (err error) {
	if err = consts.IsLegalName(name); err != nil {
		return err
	}

	groupDao := dao.NewGroupDao(ctx)

	records, err := groupDao.FindByNameAndWorkspace(name, workspace)
	if records != nil && err == nil {
		if len(records) > 0 {
			err = ErrorNameExisted
			return
		}
	}
	workspaceId, _, _ := consts.GetIdFromCode(workspace)

	record := model.GroupModel{}
	record.Name, record.Workspace, record.TypeTag = name, workspaceId, typeTag
	return groupDao.Create(record)
}

// var ErrorGroupNotFound = errors.New("error group is not exist")

func delete(groupDao *dao.GroupDao, id uint64) (err error) {
	err = groupDao.Delete(id)
	return
}

type DeleteGroupEffect interface {
	DeleleGroupEffect(uint64, uint64) error
}

func deleteEntry(ctx context.Context, groupCode, workspaceCode string, chain func(*db.BaseDao[model.GroupModel]) DeleteGroupEffect) (err error) {
	groupDao := dao.NewGroupDao(ctx)
	groupDao.BeginTransaction()
	groupId, _, _ := consts.GetIdFromCode(groupCode)
	workspaceId, _, _ := consts.GetIdFromCode(workspaceCode)

	chainDao := chain(&groupDao.BaseDao)
	if err = delete(groupDao, groupId); err != nil {
		groupDao.RollbackTransaction()
		return
	}
	if err = chainDao.DeleleGroupEffect(groupId, workspaceId); err != nil {
		groupDao.RollbackTransaction()
		return
	}
	groupDao.CommitTransaction()
	return
}

func (serv GroupService) DeleteProjectGroup(ctx context.Context, groupCode, workspaceCode string) (err error) {
	return deleteEntry(ctx, groupCode, workspaceCode, func(baseDao *db.BaseDao[model.GroupModel]) DeleteGroupEffect {
		return dao.NewProjectDao(ctx, (*db.BaseDao[model.ProjectModel])(baseDao))
	})
}
func (serv GroupService) DeleteAssetGroup(ctx context.Context, groupCode, workspaceCode string) (err error) {
	return deleteEntry(ctx, groupCode, workspaceCode, func(baseDao *db.BaseDao[model.GroupModel]) DeleteGroupEffect {
		return dao.NewAssetDao(ctx, (*db.BaseDao[model.AssetModel])(baseDao))
	})
}

var ErrorFailToRename = errors.New("Fail to rename group")

func (serv GroupService) Rename(ctx context.Context, workspaceCode, groupCode, groupName, typeTag string) (err error) {

	if err = consts.IsLegalName(groupName); err != nil {
		return
	}

	groupDao := dao.NewGroupDao(ctx)

	groups, err := groupDao.FindByNameAndWorkspace(groupName, workspaceCode)
	if err != nil {
		return err
	}

	tagedGroups := model.FilterRecordsByTypeTag(groups, typeTag)

	if len(tagedGroups) > 0 {
		return ErrorFailToRename
	}

	id, _, _ := consts.GetIdFromCode(groupCode)
	toUpdate := model.GroupModel{}
	toUpdate.Id, toUpdate.Name = id, groupName
	if err = groupDao.UpdateById(toUpdate); err != nil {
		return
	}
	return
}
