package model

import (
	"github.com/sanzashi987/nino-work/pkg/db"
	"gorm.io/gorm"
)

type ProjectVersionModel struct {
	db.BaseTime
	DeleteTime gorm.DeletedAt `gorm:"index"`

	Version       string `gorm:"column:version"`
	RootConfig    string `gorm:"column:root_config"`
	PublishConfig string `gorm:"column:publish_config"`
}
