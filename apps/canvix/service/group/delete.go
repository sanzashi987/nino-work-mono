package group

import (
	"context"
	"errors"

	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/apps/canvix/db/dao"
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
	"gorm.io/gorm"
)

func delete(tx *gorm.DB, id uint64) (err error) {
	toDelete := model.GroupModel{}
	toDelete.Id = id

	return tx.Delete(&toDelete).Error

}

type DeleleGroupEffect = func(*gorm.DB, uint64, uint64) error

var typeTagToChainedHandler = map[string]DeleleGroupEffect{
	consts.PROJECT: dao.ProjectDeleleGroupEffect,
	consts.DESIGN:  dao.AssetDeleleGroupEffect,
}

type DeleteAssetGroupReq struct {
	GroupCode string `json:"group_code" binding:"required"`
	TypeTag   string `json:"type" binding:"required"`
}

func Delete(ctx context.Context, workspaceId uint64, req *DeleteAssetGroupReq) (err error) {
	groupCode, typeTag := req.GroupCode, req.TypeTag
	tx := db.NewTx(ctx).Begin()
	groupId, _, _ := consts.GetIdFromCode(groupCode)

	deleteEffect, exist := typeTagToChainedHandler[typeTag]
	if !exist {
		return errors.New("not a supported type tag")
	}

	if err = delete(tx, groupId); err != nil {
		tx.Rollback()
		return
	}
	if err = deleteEffect(tx, groupId, workspaceId); err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}
