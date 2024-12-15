package model

import (
	"github.com/sanzashi987/nino-work/pkg/db"
)

type Bucket struct {
	db.BaseModel
	Name  string `gorm:"uniqueIndex;not null"`
	Files []File `gorm:"foreignKey:BucketID"`
}

type File struct {
	db.BaseModel
	FileId    string `gorm:"unique;index"`
	URI       string `gorm:"type:varchar(255);unique;index"`
	BucketID  uint   `gorm:"not null"`
	Name      string
	Path      string
	Size      int64
	MimeType  string
	Extension string
} 