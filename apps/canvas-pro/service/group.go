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

func (serv GroupService) Create(ctx context.Context, workspaceId uint64, name, typeTag string) (*model.GroupModel, error) {
	if err := consts.IsLegalName(name); err != nil {
		return nil, err
	}

	groupDao := dao.NewGroupDao(ctx)

	return groupDao.Create(workspaceId, name, typeTag)
}

// var ErrorGroupNotFound = errors.New("error group is not exist")

func delete(groupDao *dao.GroupDao, id uint64) (err error) {
	err = groupDao.Delete(id)
	return
}

type DeleteGroupEffect interface {
	DeleleGroupEffect(uint64, uint64) error
}

type GetChainedDao = func(context.Context, *db.BaseDao[model.GroupModel]) DeleteGroupEffect

var typeTagToChainedHandler = map[string]GetChainedDao{
	consts.PROJECT: func(ctx context.Context, baseDao *db.BaseDao[model.GroupModel]) DeleteGroupEffect {
		return dao.NewProjectDao(ctx, (*db.BaseDao[model.ProjectModel])(baseDao))
	},
	consts.DESIGN: func(ctx context.Context, baseDao *db.BaseDao[model.GroupModel]) DeleteGroupEffect {
		return dao.NewAssetDao(ctx, (*db.BaseDao[model.AssetModel])(baseDao))
	},
}

func (serv GroupService) Delete(ctx context.Context, workspaceId uint64, groupCode, typeTag string) (err error) {
	groupDao := dao.NewGroupDao(ctx)
	groupDao.BeginTransaction()
	groupId, _, _ := consts.GetIdFromCode(groupCode)

	chain, exist := typeTagToChainedHandler[typeTag]
	if !exist {
		return errors.New("Not a supported type tag")
	}

	chainDao := chain(ctx, &groupDao.BaseDao)
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

var ErrorFailToRename = errors.New("Fail to rename group")

func (serv GroupService) Rename(ctx context.Context, workspaceId uint64, groupCode, groupName, typeTag string) (err error) {

	if err = consts.IsLegalName(groupName); err != nil {
		return
	}

	groupDao := dao.NewGroupDao(ctx)

	groups, err := groupDao.FindByNameAndWorkspace(groupName, workspaceId, typeTag)
	if err != nil {
		return err
	}

	// tagedGroups := model.FilterRecordsByTypeTag(groups, typeTag)

	if len(groups) > 0 {
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

func createGroup[T any](ctx context.Context, chainedDao *dao.AnyDao[T], workspaceId uint64, groupName, typeTag string) (*model.GroupModel, error) {

	if groupName != "" {
		if err := consts.IsLegalName(groupName); err != nil {
			chainedDao.RollbackTransaction()
			return nil, err
		}

		groupDao := dao.NewGroupDao(ctx, (*db.BaseDao[model.GroupModel])(&chainedDao.BaseDao))
		newGroup, err := groupDao.Create(workspaceId, groupName, typeTag)
		if err != nil {
			chainedDao.RollbackTransaction()
			return nil, err
		}
		return newGroup, nil
	}

	return nil, nil

}

type ListGroupOutput struct {
	Id    uint64 `json:"id"`
	Name  string `json:"name"`
	Code  string `json:"code"`
	Count uint   `json:"count"`
}

func (serv GroupService) ListGroups(ctx context.Context, workspaceId uint64, groupName, typeTag string) (output []ListGroupOutput, err error) {
	groupTypeTage, err := consts.GetGroupTypeTagFromBasic(typeTag)

	groupDao := dao.NewGroupDao(ctx)

	records, err := groupDao.FindByNameAndWorkspace(groupName, workspaceId, groupTypeTage)
	if err != nil {
		return
	}

	return
}
