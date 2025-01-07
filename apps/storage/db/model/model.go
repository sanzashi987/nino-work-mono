package model

import (
	"time"

	"github.com/sanzashi987/nino-work/pkg/db"
)

type Bucket struct {
	db.BaseModel
	Code        string `gorm:"uniqueIndex;not null"`
	Description string
	Objects     []Object `gorm:"foreignKey:BucketID"`
}

type Object struct {
	db.BaseModel
	FileId    string `gorm:"unique;index"`
	URI       string `gorm:"type:varchar(255);unique;index"`
	BucketID  uint   `gorm:"not null"`
	Name      string
	Key       string `gorm:"column:key"`
	Size      int64
	MimeType  string
	Extension string
}

type Temps struct {
	db.BaseModel
	FileId    string `gorm:"unique;index"`
	Done      bool
	FinshedAt time.Time
}
