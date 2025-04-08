package group

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/canvix/consts"
	"github.com/sanzashi987/nino-work/apps/canvix/db/dao"
	"github.com/sanzashi987/nino-work/apps/canvix/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
	"gorm.io/gorm"
)

func CreateGroup(tx *gorm.DB, workspaceId uint64, groupName, typeTag string) (*model.GroupModel, error) {

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

func Create(ctx context.Context, workspaceId uint64, name, typeTag string) (*model.GroupModel, error) {
	tx := db.NewTx(ctx).Begin()
	return CreateGroup(tx, workspaceId, name, typeTag)
}
