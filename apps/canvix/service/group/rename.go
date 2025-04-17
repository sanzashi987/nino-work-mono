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

type UpdateAssetGroupReq struct {
	GroupName string `json:"group_name" binding:"required"`
	GroupCode string `json:"group_code" binding:"required"`
	TypeTag   string `json:"type" binding:"required"`
}

func Rename(ctx context.Context, workspaceId uint64, req *UpdateAssetGroupReq) error {
	groupName, groupCode, typeTag := req.GroupName, req.GroupCode, req.TypeTag
	if err := consts.IsLegalName(groupName); err != nil {
		return err
	}

	tx := db.NewTx(ctx)

	groups, err := dao.FindByNameAndWorkspace(tx, workspaceId, groupName, typeTag)
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
