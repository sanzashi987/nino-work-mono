package model

import (
	"github.com/sanzashi987/nino-work/pkg/db"
	"gorm.io/gorm"
)

type AssetVersionModel struct {
	db.BaseTime
	DeleteTime gorm.DeletedAt `gorm:"index"`

	Version     string `gorm:"column:version"`
	AssetConfig string `gorm:"column:asset_config;type:blob"`
}
