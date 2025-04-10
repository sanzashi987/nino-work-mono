package group

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/pkg/db"
)

func commonMoveGroup(codes []string, groupCode string) (uint64, []uint64, error) {
	groupId, _, err := consts.GetIdFromCode(groupCode)
	if err != nil {
		return 0, nil, err
	}

	ids := []uint64{}
	for _, code := range codes {
		id, _, errInside := consts.GetIdFromCode(code)
		if errInside != nil {
			err = errInside
			return 0, nil, err
		}
		ids = append(ids, id)
	}
	return groupId, ids, nil
}

type BatchMoveProjectGroupReq struct {
	GroupName   string   `json:"groupName"`
	GroupCode   string   `json:"groupCode"`
	ToMoveCodes []string `json:"codes" binding:"required"`
	TypeTag     string   `json:"typeTag" binding:"required"`
}

func BatchMoveGroup(ctx context.Context, workspaceId uint64, input *BatchMoveProjectGroupReq) error {
	code := input.GroupCode

	m, exist := typeTagToGroupCountHandler[input.TypeTag]
	if !exist {
		return errTagNotSupported
	}

	tx := db.NewTx(ctx).Begin()

	if newGroup, err := CreateGroup(tx, workspaceId, &CreateAssetGroupReq{
		GroupName: input.GroupName,
		TypeTag:   input.TypeTag,
	}); err != nil {
		return err
	} else if newGroup != nil {
		code = newGroup.Code
	}

	groupId, ids, err := commonMoveGroup(input.ToMoveCodes, code)
	if err != nil {
		return err
	}

	if err := tx.Model(&m).Where("workspace = ? AND id IN ? ", workspaceId, ids).Update("group_id", groupId).Error; err != nil {
		return err
	}

	tx.Commit()
	return nil

}
