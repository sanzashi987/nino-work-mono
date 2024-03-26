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

var groupService *GroupService

func init() {
	groupService = &GroupService{}
}

func GetGroupService() *GroupService {
	return groupService
}

var ErrorNameContainIllegalChar = errors.New("error name contains illegal character")
var ErrorNameExisted = errors.New("error group name is exist")

func create(ctx context.Context, name, workspace string, tableName string) (err error) {
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

	record := model.BaseModel{}
	record.Name, record.Workspace, record.TypeTag = name, workspaceId, consts.GROUP
	return groupDao.Create(record, db.TableName(tableName))
}

var projectGroupTableName = model.ProjectGroupModel{}.TableName()
var assetGroupTableName = model.AssetGroupModel{}.TableName()

func (serv *GroupService) CreateProjectGroup(ctx context.Context, name, workspace string) error {
	return create(ctx, name, workspace, projectGroupTableName)
}
func (serv *GroupService) CreateAssetGroup(ctx context.Context, name, workspace string) error {
	return create(ctx, name, workspace, assetGroupTableName)
}

var ErrorGroupNotFound = errors.New("error group is not exist")

func delete(ctx context.Context, code, workspace, tableName string) (err error) {
	groupDao := dao.NewGroupDao(ctx)
	id, _, _ := consts.GetIdFromCode(code)

	record, err := groupDao.FindByKey("id", id, db.TableName(tableName))
	if record == nil || err != nil {
		err = ErrorGroupNotFound
		return
	}

	if record.DeleteTime != nil {
		return
	}
	groupDao.LogicalDelete(*record, db.TableName(tableName))
	return
}

func (serv *GroupService) DeleteProjectGroup(ctx context.Context, code, workspace string) (err error) {
	if err = delete(ctx, code, workspace, projectGroupTableName); err != nil {
		return
	}

	return
}

func (serv *GroupService) DeleteAssetGroup(ctx context.Context, code, workspace string) (err error) {
	if err = delete(ctx, code, workspace, assetGroupTableName); err != nil {
		return
	}

	return
}
