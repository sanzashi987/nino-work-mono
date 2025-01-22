package service

import (
	"context"
	"errors"
	"sort"

	"github.com/sanzashi987/nino-work/apps/canvas-pro/consts"
	"github.com/sanzashi987/nino-work/apps/canvas-pro/db/dao"
	"github.com/sanzashi987/nino-work/apps/canvas-pro/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
	"gorm.io/gorm"
)

type GroupService struct{}

var GroupServiceImpl *GroupService = &GroupService{}

var ErrorNameExisted = errors.New("error group name is exist")

func (serv GroupService) Create(ctx context.Context, workspaceId uint64, name, typeTag string) (*model.GroupModel, error) {
	tx := db.NewTx(ctx)
	return createGroup(tx, workspaceId, name, typeTag)
}

// var ErrorGroupNotFound = errors.New("error group is not exist")

func delete(tx *gorm.DB, id uint64) (err error) {
	toDelete := model.GroupModel{}
	toDelete.Id = id

	return tx.Delete(&toDelete).Error

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
	tx := db.NewTx(ctx).Begin()
	groupId, _, _ := consts.GetIdFromCode(groupCode)

	chain, exist := typeTagToChainedHandler[typeTag]
	if !exist {
		return errors.New("not a supported type tag")
	}

	if err = delete(tx, groupId); err != nil {
		tx.Rollback()
		return
	}
	if err = chainDao.DeleleGroupEffect(groupId, workspaceId); err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

var ErrorFailToRename = errors.New("Fail to rename group")

func (serv GroupService) Rename(ctx context.Context, workspaceId uint64, groupCode, groupName, typeTag string) error {

	if err := consts.IsLegalName(groupName); err != nil {
		return err
	}

	tx := db.NewTx(ctx)

	groups, err := dao.FindByNameAndWorkspace(tx, groupName, workspaceId, typeTag)
	if err != nil {
		return err
	}

	// tagedGroups := model.FilterRecordsByTypeTag(groups, typeTag)

	if len(groups) > 0 {
		return ErrorFailToRename
	}

	id, _, _ := consts.GetIdFromCode(groupCode)
	toUpdate := model.GroupModel{}
	toUpdate.Id = id
	if err := tx.Model(&toUpdate).Update(map[string]any{"name": groupName}).Error; err != nil {
		return err
	}
	return nil
}

func createGroup(tx *gorm.DB, workspaceId uint64, groupName, typeTag string) (*model.GroupModel, error) {

	if groupName != "" {
		if err := consts.IsLegalName(groupName); err != nil {
			tx.Rollback()
			return nil, err
		}

		newGroup, err := dao.Create(tx, workspaceId, groupName, typeTag)
		if err != nil {
			tx.Rollback()
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

type ListGroupOutputs []ListGroupOutput

func (p ListGroupOutputs) Len() int {
	return len(p)
}
func (p ListGroupOutputs) Less(i, j int) bool {
	return p[i].Id < p[j].Id
}
func (p ListGroupOutputs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type GetGroupCount interface {
	GetCountFromGroupId(context.Context, uint64, []uint64) ([]*GroupCount, error)
}

var typeTagToGroupCountHandler = map[string]GetGroupCount{
	consts.PROJECT: ProjectServiceImpl,
	consts.DESIGN:  AssetServiceImpl,
}

func (serv GroupService) ListGroups(ctx context.Context, workspaceId uint64, groupName, typeTag string) (output ListGroupOutputs, err error) {
	groupTypeTage, err := consts.GetGroupTypeTagFromBasic(typeTag)

	tx := db.NewTx(ctx)

	records, err := dao.FindByNameAndWorkspace(tx, groupName, workspaceId, groupTypeTage)
	if err != nil {
		return
	}

	impl, exist := typeTagToGroupCountHandler[typeTag]
	if !exist {
		err = errors.New("Not find a corresponding interface related to the give type tag")
		return
	}

	groupIds := []uint64{}

	idToRecord := map[uint64]model.GroupModel{}
	for _, record := range records {
		groupIds = append(groupIds, record.Id)
		idToRecord[record.Id] = record
	}

	groupCounts, err := impl.GetCountFromGroupId(ctx, workspaceId, groupIds)

	for _, groupCount := range groupCounts {
		record := idToRecord[groupCount.Id]
		output = append(output, ListGroupOutput{
			Id:    groupCount.Id,
			Count: uint(groupCount.Count),
			Name:  record.Name,
			Code:  record.Code,
		})
	}

	sort.Sort(output)

	return
}
