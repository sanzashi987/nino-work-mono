package model

import (
	"github.com/sanzashi987/nino-work/pkg/db"
	"gorm.io/gorm"
)

type DataSourceVersionModel struct {
	db.BaseTime
	DeleteTime gorm.DeletedAt `gorm:"index"`

	Version    string `gorm:"column:version"`
	SourceInfo string `gorm:"column:source_info;type:blob"`
}
