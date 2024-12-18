package dao

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/storage/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type BucketDao struct {
	db.BaseDao[model.Bucket]
}

func NewBucketDao(ctx context.Context, dao ...*db.BaseDao[model.Bucket]) *BucketDao {
	return &BucketDao{BaseDao: db.NewDao[model.Bucket](ctx, dao...)}
}

func (dao BucketDao) CreateBucket(name string) (*model.Bucket, error) {
	bucket := &model.Bucket{Name: name}
	err := dao.GetOrm().Create(bucket).Error
	return bucket, err
}

func (dao BucketDao) GetBucket(id uint) (*model.Bucket, error) {
	var bucket model.Bucket
	err := dao.GetOrm().First(&bucket, id).Error
	return &bucket, err
}

func (dao BucketDao) GetBucketByName(name string) (*model.Bucket, error) {
	var bucket model.Bucket
	err := dao.GetOrm().Where("name = ?", name).First(&bucket).Error
	return &bucket, err
}
