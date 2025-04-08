package group

import (
	"context"
	"errors"

	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/apps/canvix/db/dao"
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

var ErrorFailToRename = errors.New("Fail to rename group")

func Rename(ctx context.Context, workspaceId uint64, groupCode, groupName, typeTag string) error {

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
	if err := tx.Model(&toUpdate).Update("name", groupName).Error; err != nil {
		return err
	}
	return nil
}
