package dao

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/storage/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
	"gorm.io/gorm"
)

type BucketDao struct {
	db.BaseDao[model.Bucket]
}

func NewBucketDao(ctx context.Context, dao ...*db.BaseDao[model.Bucket]) *BucketDao {
	return &BucketDao{BaseDao: db.NewDao[model.Bucket](ctx, dao...)}
}

func (dao BucketDao) CreateBucket(code string) (*model.Bucket, error) {
	dao.BeginTransaction()
	bucket := &model.Bucket{Code: code}
	err := dao.GetOrm().Create(bucket).Error
	if err != nil {
		dao.RollbackTransaction()
		return nil, err
	}

	// tableName := model.DynamicObjectTableName(code)
	// err = dao.GetOrm().Table(tableName).AutoMigrate(&model.Object{})
	// if err != nil {
	// 	dao.RollbackTransaction()
	// 	return nil, err
	// }
	// root folder
	rootDir := model.Object{
		BucketID:  bucket.Id,
		Directory: true,
		ParentId:  0,
	}

	err = dao.GetOrm().Create(&rootDir).Error
	if err != nil {
		dao.RollbackTransaction()
		return nil, err
	}

	dao.CommitTransaction()
	return bucket, err
}

func (dao BucketDao) GetBucket(id uint64) (*model.Bucket, error) {
	var bucket model.Bucket
	err := dao.GetOrm().First(&bucket, id).Error
	return &bucket, err
}

func GetBucketByCode(tx *gorm.DB, code string) (*model.Bucket, error) {
	var bucket model.Bucket
	err := tx.Where("code = ?", code).First(&bucket).Error
	if err != nil {
		return nil, err
	}
	return &bucket, nil
}

func GetBucketWithUsers(tx *gorm.DB, code string) (*model.Bucket, error) {
	var bucket model.Bucket
	err := tx.Preload("Users").Where("code = ?", code).First(&bucket).Error
	if err != nil {
		return nil, err
	}
	return &bucket, nil
}
