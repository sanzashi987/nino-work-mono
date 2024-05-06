package db

import (
	"context"

	"github.com/cza14h/nino-work/pkg/db"
)

type File struct {
	db.BaseModel
	FileId    string `gorm:"unique;index"`
	URI       string `gorm:"type:varchar(255);unique;index"`
	MimeType  string
	Extension string
	Size      int64
}

type UploadDao struct {
	db.BaseDao[File]
}

func NewUploadDao(ctx context.Context) *UploadDao {
	return &UploadDao{db.InitBaseDao[File](ctx)}
}

func (dao UploadDao) CreateFile(mimeType, uri, fileId, extension string, size int64) error {
	toInsert := File{
		FileId:    fileId,
		URI:       uri,
		MimeType:  mimeType,
		Extension: extension,
		Size:      size,
	}

	return dao.GetOrm().Create(&toInsert).Error
}

func (dao UploadDao) QueryFile(fileId string) (*File, error) {
	res := File{}

	err := dao.GetOrm().Model(&File{}).Where("file_id = ?", fileId).Take(&res).Error

	return &res, err
}
