package dao

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/storage/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type FileDao struct {
	db.BaseDao[model.Object]
}

func NewObjectDao(ctx context.Context, dao ...*db.BaseDao[model.Object]) *FileDao {
	return &FileDao{BaseDao: db.NewDao(ctx, dao...)}
}
