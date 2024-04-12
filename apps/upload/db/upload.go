package model

import (
	"context"

	"github.com/cza14h/nino-work/pkg/db"
)

type File struct {
	db.BaseModel
	Code string `gorm:"unique;index"`
	URI  string `gorm:"type:varchar(255);unique;index"`
	MimeType string
}

type UploadDao struct {
	db.BaseDao[File]
}

func NewUploadDao(ctx context.Context) *UploadDao {
	return &UploadDao{db.InitBaseDao[File](ctx)}
}

func (dao UploadDao) CreateFile(ty, url string) error {

}
