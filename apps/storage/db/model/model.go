package model

import (
	"fmt"

	"github.com/sanzashi987/nino-work/pkg/db"
)

type Bucket struct {
	db.BaseModel
	Code    string    `gorm:"uniqueIndex;not null"`
	AK      string    `gorm:"column:access_key"`
	SK      string    `gorm:"column:secret_key"`
	Objects []*Object `gorm:"foreignKey:BucketID"`
	Users   []*User   `gorm:"many2many:bucket_user"`
}

type Object struct {
	db.BaseModel
	BucketID  uint64 `gorm:"not null;primaryKey;autoIncrement:false"`
	Dir       string `gorm:"column:dir;primaryKey"`
	FileId    string `gorm:"unique;index;column:file_id"`
	URI       string `gorm:"type:varchar(255);unique;index;column:uri"`
	Name      string `gorm:"column:name"`
	Size      int64  `gorm:"column:size"`
	MimeType  string `gorm:"column:mime_type"`
	Extension string `gorm:"column:extension"`
}

const (
	USER uint = 0
	APP  uint = 1
)

type User struct {
	db.BaseModel
	UserId  uint64    `gorm:"column:user_id;index"`
	AppId   uint64    `gorm:"column:app_id;index"`
	Buckets []*Bucket `gorm:"many2many:bucket_user"`
	Type    uint      `gorm:"column:type"`
}

type LargeFile struct {
	BucketID       uint64 `gorm:"not null"`
	Hash           string `gorm:"uniqueIndex"`
	Chunks         int    `gorm:"column:chunks"`
	UploadedChunks string `gorm:"column:uploaded_chunks"`
	FileId         string `gorm:"column:file_id"`
	Done           bool   `gorm:"done"`
}

/** No used **/
func DynamicObjectTableName(bucketCode string) string {
	return fmt.Sprintf("objects_%s", bucketCode)
}
