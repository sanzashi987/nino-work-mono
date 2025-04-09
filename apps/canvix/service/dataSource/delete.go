package dataSource

import (
	"context"
	"errors"

	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

func Delete(ctx context.Context, workspaceId uint64, codes []string) error {
	tx := db.NewTx(ctx).Begin()

	ids := []uint64{}
	for _, code := range codes {
		id, tag, err := consts.GetIdFromCode(code)
		if err != nil {
			return err
		}
		if tag != consts.DATASOURCE {
			return errors.New("Has invalid datasource code: " + code)
		}
		ids = append(ids, id)
	}

	if err := tx.Model(&model.DataSourceModel{}).Where("workspace = ? AND id in ?", workspaceId, ids).Delete(&model.DataSourceModel{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
