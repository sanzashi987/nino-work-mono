package dao

import (
	"context"

	"github.com/sanzashi987/nino-work/apps/storage/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
)

type FileDao struct {
	db.BaseDao[model.File]
}

func NewFileDao(ctx context.Context) *FileDao {
	return &FileDao{db.InitBaseDao[model.File](ctx)}
}

func (dao FileDao) CreateFile(bucketId uint, name, mimeType, uri, fileId, extension string, size int64) error {
	toInsert := model.File{
		FileId:    fileId,
		URI:       uri,
		BucketID:  bucketId,
		Name:      name,
		MimeType:  mimeType,
		Extension: extension,
		Size:      size,
	}
	return dao.GetOrm().Create(&toInsert).Error
}

func (dao FileDao) QueryFile(fileId string) (*model.File, error) {
	res := model.File{}
	err := dao.GetOrm().Model(&model.File{}).Where("file_id = ?", fileId).Take(&res).Error
	return &res, err
}
