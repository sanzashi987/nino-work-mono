package model

import "github.com/cza14h/nino-work/pkg/db"

type DataSourceModel struct {
	db.BaseModel
	Code      string
	Name      string
	Type      int
	Deleted   int
	Workspace string
	Config    string `gorm:"type:blob"`
}

func (d DataSourceModel) TableName() string {
	return "data_sources"
}
