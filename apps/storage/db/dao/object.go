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
	return &FileDao{BaseDao: db.NewDao[model.Object](ctx, dao...)}
}

func (dao FileDao) CreateObject(bucketId uint, name, mimeType, uri, fileId, extension string, size int64) error {
	toInsert := model.Object{
		FileId:    fileId,
		URI:       uri,
		Name:      name,
		MimeType:  mimeType,
		Extension: extension,
		Size:      size,
	}
	return dao.GetOrm().Create(&toInsert).Error
}

func (dao FileDao) QueryFile(fileId string) (*model.Object, error) {
	res := model.Object{}
	err := dao.GetOrm().Model(&model.Object{}).Where("file_id = ?", fileId).Take(&res).Error
	return &res, err
}
