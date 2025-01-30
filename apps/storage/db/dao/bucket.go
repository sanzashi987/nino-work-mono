package dao

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/sanzashi987/nino-work/apps/storage/db/model"
	"gorm.io/gorm"
)

// Requires `tx` is a beganned transaction
func AddUsersToBucket(tx *gorm.DB, bucketId uint64, userIds []uint64) error {
	// test if user is in the table
	var existingUsers []*model.User
	if err := tx.Where("user_id IN ?", userIds).Find(&existingUsers).Error; err != nil {
		return err
	}

	existingIDMap := make(map[uint64]bool)
	for _, user := range existingUsers {
		existingIDMap[user.Id] = true
	}

	var missingUsers []*model.User
	for _, id := range userIds {
		if !existingIDMap[id] {
			missingUsers = append(missingUsers, &model.User{UserId: id})
		}
	}

	if len(missingUsers) > 0 {
		if err := tx.Create(&missingUsers).Error; err != nil {
			return err
		}
	}

	// add user to bucket
	bucket := &model.Bucket{}
	bucket.Id = bucketId

	toAppend := append([]*model.User{}, existingUsers...)
	toAppend = append(toAppend, missingUsers...)
	return tx.Model(bucket).Association("Users").Append(toAppend)

}

func CreateBucket(tx *gorm.DB, code, bucketpath string) (*model.Bucket, error) {
	bucket := &model.Bucket{Code: code}
	bucketFullpath := filepath.Join(bucketpath, code)
	if err := os.MkdirAll(bucketFullpath, fs.ModePerm); err != nil {
		return nil, nil
	}

	if err := tx.Create(bucket).Error; err != nil {
		return nil, err
	}

	// tableName := model.DynamicObjectTableName(code)
	// err = tx.Table(tableName).AutoMigrate(&model.Object{})
	// if err != nil {
	// 	tx.Rollback()
	// 	return nil, err
	// }

	// root folder
	rootDir := model.Object{
		BucketID: bucket.Id,
		Dir:      true,
		Name:     "Root",
		ParentId: 0,
	}

	// err = tx.Create(&rootDir).Error
	if err := tx.Create(&rootDir).Error; err != nil {
		return nil, err
	}

	return bucket, nil
}

func GetBucket(tx *gorm.DB, id uint64) (*model.Bucket, error) {
	var bucket model.Bucket
	err := tx.First(&bucket, id).Error
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

func ListObjectsByDir(tx *gorm.DB, bucketId, parentPathId uint64) ([]*model.Object, error) {

	models := []*model.Object{}
	if err := tx.Where("bucket_id = ? AND parent_id = ? ", bucketId, parentPathId).Find(&models).Error; err != nil {
		return nil, err
	}

	return models, nil

}
