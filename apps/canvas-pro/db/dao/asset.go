package dao

import (
	"context"

	"gorm.io/gorm"
)

type AssetDao struct {
	*gorm.DB
}

func NewAssetDao(ctx context.Context) *AssetDao {
	return &AssetDao{
		DB: newDBSession(ctx),
	}
}

func (a *AssetDao) Create() {

}
